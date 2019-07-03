package cryptography

import (
	"encoding/base64"
	"fmt"
	"testing"
)

func TestDes(t *testing.T) {
	// 加密
	key := []byte("12345678") // 密码长度 8 byte
	res := DESEncryptCBC([]byte("远处蔚蓝天空下涌动着金色的麦浪"), key)
	fmt.Println(base64.StdEncoding.EncodeToString(res))

	// 解密
	res = DESDecryptCBC(res, key)
	fmt.Println("解密后的数据:", string(res))

}

func TestTripleDes(t *testing.T) {
	// 加密
	key := []byte("123456781234567812345678") // 密码长度 24 byte
	res := TripleDESEncrypt([]byte("远处蔚蓝天空下涌动着金色的麦浪"), key)
	fmt.Println(base64.StdEncoding.EncodeToString(res))

	// 解密
	res = TripleDESDecrypt(res, key)
	fmt.Println("解密后的数据:", string(res))

}

func TestAes(t *testing.T) {
	// 加密
	key := []byte("1234567812345678")
	res := AESEncrypt([]byte("远处蔚蓝天空下涌动着金色的麦浪"), key)
	fmt.Println(base64.StdEncoding.EncodeToString(res))

	// 解密
	res = AESDecrypt(res, key)
	fmt.Println("解密后的数据:", string(res))

}
