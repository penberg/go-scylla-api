package scylla

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

type Client struct {
	baseURL string
}

func NewClient(host string, port string) *Client {
	baseURL := fmt.Sprintf("http://%s:%s", host, port)
	return &Client{
		baseURL: baseURL,
	}
}

func (c *Client) Drain() error {
	return c.post("/storage_service/drain", nil)
}

func (c *Client) Flush(keyspace string, columnFamilies []string) error {
	target := "/storage_service/keyspace_flush/" + keyspace
	if len(columnFamilies) > 0 {
		target += "?cf=" + strings.Join(columnFamilies, ",")
	}
	return c.post(target, nil)
}

func (c *Client) Rebuild(sourceDC string) error {
	target := "/storage_service/rebuild/"
	if len(sourceDC) > 0 {
		target += "?source_dc=" + sourceDC
	}
	return c.post(target, nil)
}

func (c *Client) post(target string, body io.Reader) error {
	if _, err := http.Post(c.baseURL+target, "application/json", nil); err != nil {
		return err
	}
	return nil
}
