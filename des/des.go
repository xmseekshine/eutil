package util

import (
	"bytes"
	"crypto/des"
	"encoding/base64"
	"fmt"
)

// EncryptDESECB func
func EncryptDESECB(src []byte, key []byte) []byte {
	desBlockEncrypter, err := des.NewCipher(key[0:8])
	if err != nil {
		panic(err)
	}

	bs := desBlockEncrypter.BlockSize()
	srcPadding := PKCS5Padding(src, bs)
	if len(srcPadding)%bs != 0 {
		//return nil, errors.New("Need a multiple of the blocksize")
		fmt.Printf("Need a multiple of the blocksize")
	}
	out := make([]byte, len(srcPadding))
	dst := out
	for len(srcPadding) > 0 {
		desBlockEncrypter.Encrypt(dst, srcPadding[:bs])
		srcPadding = srcPadding[bs:]
		dst = dst[bs:]
	}

	return out
}

// DecryptDESECB func
func DecryptDESECB(src []byte, key []byte) []byte {
	desBlockDecrypter, err := des.NewCipher(key[0:8])
	if err != nil {
		panic(err)
	}

	bs := desBlockDecrypter.BlockSize()
	out := make([]byte, len(src))
	dst := out
	for len(src) > 0 {
		desBlockDecrypter.Decrypt(dst, src[:bs])
		src = src[bs:]
		dst = dst[bs:]
	}

	out = PKCS5UnPadding(out)

	return out

}

//EntryptDesECBBase64 ...
func EntryptDesECBBase64(data, key string) string {
	out := EncryptDESECB([]byte(data), []byte(key))
	return base64.StdEncoding.EncodeToString(out)
}

//DecryptDesECBBase64 ...
func DecryptDesECBBase64(data, key string) string {
	tmp, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return ""
	}
	out := DecryptDESECB(tmp, []byte(key))
	return string(out)
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	// 去掉最后一个字节 unpadding 次
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
