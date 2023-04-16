package readWrite

import (
	"fmt"
	"io"
	"os"
)

// filename 输入文件读取文件的名字
// 返回读取内容
func ReadFileFunc(fileName string) string {
	//打开文件
	file, err := os.Open(fileName)
	if err != nil {
		return "文件打开失败，没有" + fileName + "文件"
	}
	//及时关闭 file 句柄，否则会有内存泄漏
	defer file.Close()

	//读取内容
	buf := make([]byte, 1024)
	var res string
	for {
		count, err := file.Read(buf)
		if err == io.EOF {
			break
		} else {
			currBytes := buf[:count]
			res += string(currBytes)
		}
	}

	return res
}

// writeString 写入的字符串
// newfileName 创建新的文件名字
func WriteFileFunc(writeString string, newfileName string) bool {
	//创建文件
	file, err := os.Create(newfileName)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	// 写入文件内容
	_, err = io.WriteString(file, writeString)
	if err != nil {
		fmt.Println(err)
	}
	return true
}
