package utils

import "strings"

// CombineURL takes a base URL and a path and combines them into a single URL string.
//
// Parameters:
//   - baseURL: The base URL as a string. It may or may not end with a "/".
//   - path: A string representing the path to be appended to the base URL. It may or may not start with a "/".
//
// Returns:
//   - A single string that represents the combination of the base URL and the path.
//
// Implementation:
// The function first checks if the base URL ends with a "/".
// If it does, the function then checks if the path begins with a "/". If it does, the function removes the "/"
// from the end of the baseURL before appending the path. If it does not, the function simply appends the path.
//
// If the baseURL does not end with a "/", the function checks if the path begins with a "/".
// If it does, the function simply appends the path. If it does not, the function adds a "/" to the baseURL before appending the path.
// This is to ensure that there is exactly one "/" between the baseURL and the path.
func CombineURL(baseURL, path string) string {
	if strings.HasSuffix(baseURL, "/") {
		if strings.HasPrefix(path, "/") {
			return baseURL[0:len(baseURL)-1] + path
		} else {
			return baseURL + path
		}
	} else {
		if strings.HasPrefix(path, "/") {
			return baseURL + path
		} else {
			return baseURL + "/" + path
		}

	}
}
