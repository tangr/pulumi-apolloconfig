//go:build debug

package apolloconfigapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) printAsCurl(req *http.Request) {
    curl := fmt.Sprintf("curl -X %s '%s'", req.Method, req.URL.String())

    // 添加请求头
    for key, values := range req.Header {
        for _, value := range values {
            curl += fmt.Sprintf(" -H '%s: %s'", key, value)
        }
    }

    // 添加请求体
    if req.Body != nil {
        // 因为 Body 是一个 io.ReadCloser，我们需要复制它的内容
        bodyBytes, err := io.ReadAll(req.Body)
        if err == nil {
            // 重新设置请求体，因为它已经被读取了
            req.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
            // 如果是 JSON 数据，美化输出
            var prettyJSON bytes.Buffer
            err = json.Indent(&prettyJSON, bodyBytes, "", "  ")
            if err == nil {
                curl += fmt.Sprintf(" -d '%s'", prettyJSON.String())
            } else {
                curl += fmt.Sprintf(" -d '%s'", string(bodyBytes))
            }
        }
    }

    fmt.Printf("\nCurl command:\n%s\n\n", curl)
}
