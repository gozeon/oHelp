package utils

import (
	"net/http"
	"net/url"
)

func RequestURL(r *http.Request) (*url.URL, error) {
	u, err := url.Parse(r.URL.String())
	if err != nil {
		return nil, err
	}

	u.Host = r.Host
	if host := r.Header.Get("X-Forwarded-Host"); host != "" {
		u.Host = host
	}

	u.Scheme = "http"
	if scheme := r.Header.Get("X-Forwarded-Proto"); scheme != "" {
		u.Scheme = scheme
	} else if r.TLS != nil {
		u.Scheme = "https"
	}

	return u, nil
}
