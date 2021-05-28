package qrcode

import (
	"encoding/base64"
	"fmt"

	qrcode "github.com/skip2/go-qrcode"
)

//GenCode 二维码
func GenCode(content string, size int) (string, error) {
	//data:image/png;base64,
	binaryData, err := qrcode.Encode(content, qrcode.Medium, size)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("data:%s;base64,%s", "image/png", base64.StdEncoding.EncodeToString(binaryData)), nil
}

//WriteFile 写到文件
func WriteFile(content string, size int, fielName string) error {
	return qrcode.WriteFile(content, qrcode.Medium, size, fielName)
}
