package cryptography

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/des"
	"fmt"
)

// DESEncryptCBC DES 加密
func DESEncryptCBC(src []byte, key []byte) (dst []byte) {
	block, err := des.NewCipher(key)
	if err != nil {
		panic(err)
	}

	src = PKCS5Padding(src, block.BlockSize())
	tmp := []byte("Artifact")
	blockMode := cipher.NewCBCEncrypter(block, tmp)

	dst = make([]byte, len(src))
	blockMode.CryptBlocks(dst, src)
	fmt.Println("加密后的数据:", dst)
	return
}

// DESDecryptCBC  DES 解密
func DESDecryptCBC(src []byte, key []byte) (dst []byte) {
	block, err := des.NewCipher(key)
	if err != nil {
		panic(err)
	}

	tmp := []byte("Artifact")
	blockMode := cipher.NewCBCDecrypter(block, tmp)

	dst = src
	blockMode.CryptBlocks(src, dst)
	dst = PKCS5UnPadding(dst)
	return
}

// TripleDESEncrypt 3DES 加密
func TripleDESEncrypt(src, key []byte) (dst []byte) {
	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		panic(err)
	}

	src = PKCS5Padding(src, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, key[:8])

	dst = src
	blockMode.CryptBlocks(dst, src)
	return dst
}

// TripleDESDecrypt 3DES 解密
func TripleDESDecrypt(src, key []byte) (dst []byte) {
	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		panic(err)
	}

	blockMode := cipher.NewCBCDecrypter(block, key[:8])

	dst = src
	blockMode.CryptBlocks(dst, src)

	dst = PKCS5UnPadding(dst)
	return
}

// AESEncrypt AES 加密
func AESEncrypt(src, key []byte) (dst []byte) {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	src = PKCS5Padding(src, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, key[:block.BlockSize()])

	dst = src
	blockMode.CryptBlocks(dst, src)
	return
}

// AESDecrypt AES 解密
func AESDecrypt(src, key []byte) (dst []byte) {
	block, err := aes.NewCipher(key)

	if err != nil {
		panic(err)
	}

	blockMode := cipher.NewCBCDecrypter(block, key[:block.BlockSize()])
	dst = src
	blockMode.CryptBlocks(dst, src)

	dst = PKCS5UnPadding(dst)
	return
}

// PKCS5Padding pkcs5 方式填充
func PKCS5Padding(cipher []byte, blockSize int) []byte {
	padding := blockSize - (len(cipher) % blockSize)
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	newText := append(cipher, padText...)
	return newText
}

// PKCS5UnPadding pkcs5 删除尾部数据
func PKCS5UnPadding(origData []byte) []byte {
	len := len(origData)
	num := int(origData[len-1])
	return origData[:(len - num)]
}
