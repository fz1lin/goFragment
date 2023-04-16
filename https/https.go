package https

import (
	"crypto/tls"
	"fmt"
	"net/http"
)

func main() {
	// 创建一个 http.Transport，设置忽略 SSL 证书验证
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	// 创建一个 http.Client，指定上面创建的 http.Transport
	client := &http.Client{Transport: tr}

	// 发送一个 GET 请求
	resp, err := client.Get("https://example.com")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	// 输出响应结果
	fmt.Println(resp.Status)
}
