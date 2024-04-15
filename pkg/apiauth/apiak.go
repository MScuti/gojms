package apiauth

import (
	"bytes"
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/bytedance/sonic"
	"gopkg.in/twindagger/httpsig.v1"
	"io"
	"net/http"
	"net/url"
	"os"
)

type JmsAKConfig struct {
	Endpoints string `json:"endpoints"`
	Debug     bool   `json:"debug"`
}

func (j *JmsAKConfig) SignReq(r *http.Request) error {
	// read token
	data, err := os.ReadFile("/run/conjur/access-token")
	if err != nil {
		return err
	}
	token := base64.StdEncoding.EncodeToString(data)

	// read endpoint
	endpoint, exists := os.LookupEnv("CONJUR_APPLIANCE_URL")
	if !exists {
		return errors.New("CONJUR_APPLIANCE_URL not found")
	}

	// ak endpoint
	akApi := fmt.Sprintf("%s/secrets/lixiang/variable/%s", endpoint, "Prd_Vault/authn/App_JMS-Tools_prd/IT_JumpServer_JMS-Tools/username")
	req, err := http.NewRequest("GET", akApi, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Token token=\"%s\"", token))

	// do request
	client := &http.Client{Transport: &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("read access token body error: %s", err)
	}
	if resp.StatusCode != 200 {
		return fmt.Errorf("get access token error,status code:%d : %s", resp.StatusCode, string(body))
	}
	ak := string(body)

	// sk endpoint
	skApi := fmt.Sprintf("%s/secrets/lixiang/variable/%s", endpoint, "Prd_Vault/authn/App_JMS-Tools_prd/IT_JumpServer_JMS-Tools/password")
	req, err = http.NewRequest("GET", skApi, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Token token=\"%s\"", token))

	// do request
	resp, err = client.Do(req)
	if err != nil {
		return err
	}
	body, err = io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("read secrect token body error: %s", err)
	}
	if resp.StatusCode != 200 {
		return fmt.Errorf("get secrect token error,status code:%d : %s", resp.StatusCode, string(body))
	}
	sk := string(body)

	// sign request
	signer, err := httpsig.NewRequestSigner(ak, sk, "hmac-sha256")
	if err != nil {
		return err
	}
	return signer.SignRequest(r, []string{"(request-target)", "date"}, nil)
}

func (j *JmsAKConfig) MakeRequest(method, endpoint string, body interface{}) (*http.Request, error) {
	// process body data
	var bodyReader io.Reader
	if body != nil {
		b, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("make request error marshal body error: %s", err)
		}
		bodyReader = bytes.NewReader(b)
	}

	// make request
	req, err := http.NewRequest(method, endpoint, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("make new request error: %s", err)
	}

	// set header
	req.Header.Set("Content-Type", "application/json")
	err = j.SignReq(req)
	if err != nil {
		return nil, fmt.Errorf("sign request error: %s", err)
	}
	return req, nil
}

func (j *JmsAKConfig) DoRequest(req *http.Request, result interface{}) error {
	// do request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	// read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
		}
	}(resp.Body)

	// set debug
	if j.Debug {
		fmt.Printf("response body: %s\n", body)
	}

	// check response status code
	if resp.StatusCode < 200 || resp.StatusCode >= 400 {
		return fmt.Errorf("server response code is not ok, code:%d, content:%s", resp.StatusCode, body)
	}

	// check if result is nil
	if result == nil {
		return nil
	}

	// unmarshal response body
	err = sonic.Unmarshal(body, result)
	return err

}

func (j *JmsAKConfig) SetQuery(req *http.Request, v url.Values) *http.Request {
	// set query
	req.URL.RawQuery = v.Encode()
	return req
}

func (j *JmsAKConfig) GetEndpoint() string {
	return j.Endpoints
}
