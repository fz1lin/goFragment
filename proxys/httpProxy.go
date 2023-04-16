package proxys

import (
	"fmt"
	"net/http"
	"net/url"
)

// 设置本程序全局代理
func SetGlobalHttpProxy(proxyServer string) error {
	if proxyServer == "" {
		return nil
	}

	proxyURL, err := url.Parse(proxyServer)
	if err != nil {
		return fmt.Errorf("Error parsing proxy server address: %s", err)
	}

	http.DefaultTransport = &http.Transport{Proxy: http.ProxyURL(proxyURL)}
	return nil
}

func main() {
	//调用代理,可以判断结合命令行参数进行调用
	proxyServer := "http://127.0.0.1:123"
	if err := SetGlobalHttpProxy(proxyServer); err != nil {
		fmt.Println(err)
		return
	}
}
