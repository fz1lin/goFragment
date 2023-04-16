package times

import (
	"fmt"
	"time"
)

func TimeSince(f func()) string {
	start := time.Now() // 获取当前时间
	f()                 // 执行传入的函数
	// 获取结束的时间
	elapsed := time.Since(start)
	str := fmt.Sprintf("该函数执行完成耗时：%v", elapsed)
	return str
}

func main() {
	result := TimeSince(timeTest)
	fmt.Println(result)
}

func timeTest() {
	// 运行的函数
	sum := 0
	for i := 0; i < 1111111111; i++ {
		sum++
	}
}
