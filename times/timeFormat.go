package times

import (
	"fmt"
	"time"
)

func main() {
	//获取当前时间
	now := time.Now()
	fmt.Println(now)

	//获取时间戳
	fmt.Println(now.Unix())
	//格式化时间
	fmt.Println(now.Format("2006-01-02 15:04:05.000 Mon Jan"))
	fmt.Println(now.Format("2006-01-02 15:04:05"))
	fmt.Println(now.Format("20060102"))
}
