package threads

import (
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

type Response struct {
	Url        string
	StatusCode int
	Body       string
	Error      error
}

func fetchURL(url string, wg *sync.WaitGroup, ch chan<- Response) {
	defer wg.Done()

	resp, err := http.Get(url)
	if err != nil {
		ch <- Response{Url: url, Error: err}
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		ch <- Response{Url: url, Error: err}
		return
	}

	ch <- Response{Url: url, StatusCode: resp.StatusCode, Body: string(body)}
}

func fetchAllURLs(urls []string, maxConcurrency int) []Response {
	ch := make(chan Response, len(urls))
	var wg sync.WaitGroup

	for _, url := range urls {
		wg.Add(1)
		go fetchURL(url, &wg, ch)
	}

	wg.Wait()
	close(ch)

	var results []Response
	for resp := range ch {
		results = append(results, resp)
	}

	return results
}

func main() {
	start := time.Now() // 获取当前时间
	urls := []string{
		"https://www.google.com",
		"https://www.facebook.com",
		"https://www.apple.com",
		"https://www.amazon.com",
	}

	//传入url列表，以及并发数量发送http请求
	responses := fetchAllURLs(urls, 1)

	//接口返回的值
	for _, resp := range responses {
		if resp.Error != nil {
			fmt.Printf("Error fetching %s: %s\n", resp.Url, resp.Error)
		} else {
			fmt.Printf("%s returned status %d, body length %d\n", resp.Url, resp.StatusCode, len(resp.Body))
		}
	}

	//获取结束的时间
	elapsed := time.Since(start)
	fmt.Println("该函数执行完成耗时：", elapsed)
}
