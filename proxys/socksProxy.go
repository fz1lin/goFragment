package proxys

import (
	"fmt"
	"golang.org/x/net/proxy"
	"net/http"
	"net/url"
	"time"
)

func SetGlobalSocksProxy(proxyServer string) error {
	if proxyServer == "" {
		return nil
	}

	proxyURL, err := url.Parse(proxyServer)
	if err != nil {
		return fmt.Errorf("Error parsing proxy server address: %s", err)
	}

	dialer, err := proxy.FromURL(proxyURL, proxy.Direct)
	if err != nil {
		return fmt.Errorf("Error creating SOCKS5 dialer: %s", err)
	}

	http.DefaultTransport = &http.Transport{
		Dial:                dialer.Dial,
		DialContext:         dialer.(proxy.ContextDialer).DialContext,
		DialTLS:             dialer.Dial,
		DialTLSContext:      dialer.(proxy.ContextDialer).DialContext,
		TLSHandshakeTimeout: 10 * time.Second,
	}

	return nil
}

func main() {
	//调用代理,可以判断结合命令行参数进行调用
	if err := SetGlobalSocksProxy("127.0.0.1:123"); err != nil {
		fmt.Println(err)
		return
	}
}
