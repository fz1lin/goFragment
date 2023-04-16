package https

import (
	"bytes"
	"fmt"
	"net/http"
)

func SendPostRequest(url string, payload []byte, headers map[string]string) (*http.Response, error) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}

	// 设置请求头部
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// 调用
func main() {
	url := "https://example.com/api"
	payload := []byte(`{"key": "value"}`)
	headers := map[string]string{
		"Content-Type":  "application/json",
		"Authorization": "Bearer abc123",
	}

	resp, err := SendPostRequest(url, payload, headers)
	if err != nil {
		fmt.Println(err)
		return
	}
	body := resp.Body
	fmt.Println(body)
	// 处理响应
}
