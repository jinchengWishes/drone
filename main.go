package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
)

func main() {
	// 我是master注释
	x, y := 5, 6
	for i := 0; i < 10; i++ {
		fmt.Println(x, "-", y)
		x, y = y, x+y
	}

	key := "Swibk3rUs06eJHNu"
	var data []byte = []byte(`{"name":"lijincheng","age":12}`)
	text, err := Encrypt(key, data)
	if err != nil {
		fmt.Println(err)
		return
	}
	// gYDvcj9QWbGIMHrD3RAgNVgi4w3q5UxcnwY0y81RgDFhTriRuMq6COgtRVQokE/C
	fmt.Println(text)

}

//Encrypt 加密
func Encrypt(keyStr string, src []byte) (res string, err error) {
	key := []byte(keyStr)
	block, err := aes.NewCipher(key)
	if err != nil {
		return
	}
	src = PKCS5Padding(src, block.BlockSize())
	dst, err := ECBEncrypt(block, src)
	if err != nil {
		return
	}
	return base64.StdEncoding.EncodeToString(dst), nil
}

func PKCS5Padding(src []byte, blockSize int) []byte {
	padding := blockSize - len(src)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, padtext...)
}

func ECBEncrypt(block cipher.Block, src []byte) ([]byte, error) {
	blockSize := block.BlockSize()
	encryptData := make([]byte, 0, len(src))
	tmpData := make([]byte, blockSize)
	for index := 0; index < len(src); index += blockSize {
		block.Encrypt(tmpData, src[index:index+blockSize])
		encryptData = append(encryptData, tmpData...)
	}
	return encryptData, nil
}
