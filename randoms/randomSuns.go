package randoms

import (
	"fmt"
	"math/rand"
	"time"
)

// 生成随机字符串
const (
	letterBytes  = "abcdefghijklmnopqrstuvwxyz"
	capitalBytes = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numberBytes  = "0123456789"
	symbolBytes  = "!@#$%^&*()"
)

// 生成随机数
// start 10，end 20，将获得这之间的随机数
func RandomSum(start int, end int) int {
	// 以当前时间戳作为种子，确保每次生成的随机数都不同
	rand.Seed(time.Now().Unix())
	// 生成一个指定范围之间的随机整数

	randomInt := rand.Intn(end-start+1) + start
	return randomInt
}

func RandStringBytes(n int, randomString string) string {
	rand.Seed(time.Now().Unix())
	b := make([]byte, n)
	for i := range b {
		b[i] = randomString[rand.Intn(len(randomString))]
	}
	return string(b)
}

func main() {
	fmt.Println(RandomSum(1, 4))
	//传入随机字符串的长度，以及需要随机那些字符串
	randomString := RandStringBytes(10, letterBytes+numberBytes)
	fmt.Println(randomString)
}
