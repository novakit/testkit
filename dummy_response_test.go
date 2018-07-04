package testkit

import (
	"net/http"
	"net/url"
	"testing"

	"github.com/novakit/nova"
)

func TestDummyResponse_All(t *testing.T) {
	d := NewDummyResponse()
	r := &http.Request{Method: http.MethodGet, Host: "example.com"}
	r.URL, _ = url.Parse("/hello/world")
	n := nova.New()
	n.Use(func(c *nova.Context) error {
		c.Res.WriteHeader(http.StatusFound)
		c.Res.Write([]byte(c.Req.URL.Path))
		return nil
	})
	n.ServeHTTP(d, r)
	if d.StatusCode != http.StatusFound {
		t.Error("bad status code")
	}
	if d.String() != "/hello/world" {
		t.Error("bad content")
	}
}
