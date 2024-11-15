//go:build !debug

package apolloconfigapi

import "net/http"

func (c *Client) printAsCurl(req *http.Request) {
}
