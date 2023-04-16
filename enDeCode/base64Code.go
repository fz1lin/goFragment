package enDeCode

import "encoding/base64"

// base64编码
func base64Encode(encodeStr string) string {
	return base64.StdEncoding.EncodeToString([]byte(encodeStr))
}

// base64解码
func Base64Decode(decodeStr string) string {
	decoded, err := base64.StdEncoding.DecodeString(decodeStr)
	if err != nil {
		panic(err)
	}
	return string(decoded)
}
