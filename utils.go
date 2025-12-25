package main

import "net/http"

// utils
func isHTMX(r *http.Request) bool {
	return r.Header.Get("HX-Request") == "true"
}
