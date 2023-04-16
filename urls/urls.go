package urls

import (
	"fmt"
	"net/url"
)

func main() {
	// 要解析的 URL 字符串
	urlString := "https://www.example.com/path/to/resource?key1=value1&key2=value2#fragment"

	// 解析 URL 字符串
	parsedUrl, err := url.Parse(urlString)
	if err != nil {
		panic(err)
	}

	// 输出解析后的 URL 的各个部分
	fmt.Println("Scheme:", parsedUrl.Scheme)
	fmt.Println("Host:", parsedUrl.Host)
	fmt.Println("Path:", parsedUrl.Path)
	fmt.Println("RawQuery:", parsedUrl.RawQuery)
	fmt.Println("Fragment:", parsedUrl.Fragment)

	// 解析 URL 的查询参数
	query := parsedUrl.Query()
	fmt.Println("key1:", query.Get("key1"))
	fmt.Println("key2:", query.Get("key2"))
}
