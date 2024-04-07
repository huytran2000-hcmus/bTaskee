package httpapi

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"time"
)

type HTTPRequest struct {
	Method string            // HTTP method
	URL    string            // URL to send the request to
	Header map[string]string // HTTP headers
	Body   interface{}       // Body of the request
}

type HTTPOptions struct {
	Timeout time.Duration // Timeout for the request
}

func SendHTTPRequest[O any](ctx context.Context, httpReq HTTPRequest, options HTTPOptions) (O, error) {
	var err error

	var res O
	var bodyBytes []byte
	if httpReq.Body != nil {
		bodyBytes, err = json.Marshal(httpReq.Body)
		if err != nil {
			return res, err
		}
	}

	req, err := http.NewRequestWithContext(ctx, httpReq.Method, httpReq.URL, bytes.NewReader(bodyBytes))
	if err != nil {
		return res, err
	}

	req.Header.Set("Content-Type", "application/json")
	for k, v := range httpReq.Header {
		req.Header.Set(k, v)
	}

	client := &http.Client{Timeout: options.Timeout}
	resp, err := client.Do(req)
	if err != nil {
		return res, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		return res, err
	}

	return res, nil
}
