package https

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"io"
	"net/http"
	"time"
)

func main() {
	m := GetFunc("https://baidu.com")
	fmt.Println(m["code"])

	//参数拼接即可
	url := "http://httpbin.org/get"
	//自定义 ua
	headerCustom := map[string]string{
		"User-Agent": "123",
		"name":       "771",
	}

	headerFunc := GetCustomHeaderFunc(url, headerCustom)
	//fmt.Println(headerFunc["body"])
	fmt.Println(headerFunc["headers"])
}

func GetFunc(url string) map[string]interface{} {
	valueMap := make(map[string]interface{})
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error while getting the response:", err)
	}
	defer response.Body.Close()
	code := response.StatusCode
	header := response.Header.Get("Server")
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}
	valueMap["code"] = code
	valueMap["header"] = header
	valueMap["body"] = string(body)

	return valueMap

}

func GetCustomHeaderFunc(url string, headerMap interface{}) map[string]interface{} {
	req, err := http.NewRequest("GET", url, nil)
	respMap := make(map[string]interface{})
	if err != nil {
		fmt.Println(err)
	}

	m := headerMap.(map[string]string)
	var headerKey string
	var headerValue string

	for k, v := range m {
		headerKey = k
		headerValue = v
	}
	//设置header头
	req.Header.Set(headerKey, headerValue)
	client := http.Client{
		Timeout: 3 * time.Second, //超时时间
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	body, _ := io.ReadAll(resp.Body)
	headers := jsoniter.Get(body, "headers")

	respMap["body"] = string(body)
	respMap["headers"] = headers.ToString()
	return respMap

}
