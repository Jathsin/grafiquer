package utils

import "net/http"

// utils
func IsHTMX(r *http.Request) bool {
	return r.Header.Get("HX-Request") == "true"
}
