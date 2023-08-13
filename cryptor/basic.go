package cryptor

import (
	"bufio"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
	"os"
)

/*
   @File: basic.go
   @Author: khaosles
   @Time: 2023/8/13 13:20
   @Desc:
*/

// Base64StdEncode encode string with base64 encoding.
func Base64StdEncode(s string) string {
	return base64.StdEncoding.EncodeToString([]byte(s))
}

// Base64StdDecode decode a base64 encoded string.
func Base64StdDecode(s string) string {
	b, _ := base64.StdEncoding.DecodeString(s)
	return string(b)
}

// Md5String return the md5 value of string.
func Md5String(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

// Md5StringWithBase64 return the md5 value of string with base64.
func Md5StringWithBase64(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

// Md5Byte return the md5 string of byte slice.
func Md5Byte(data []byte) string {
	h := md5.New()
	h.Write(data)
	return hex.EncodeToString(h.Sum(nil))
}

// Md5ByteWithBase64 return the md5 string of byte slice with base64.
func Md5ByteWithBase64(data []byte) string {
	h := md5.New()
	h.Write(data)
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

// Md5File return the md5 value of file.
func Md5File(filename string) (string, error) {
	if fileInfo, err := os.Stat(filename); err != nil {
		return "", err
	} else if fileInfo.IsDir() {
		return "", nil
	}

	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hash := md5.New()

	chunkSize := 65536
	for buf, reader := make([]byte, chunkSize), bufio.NewReader(file); ; {
		n, err := reader.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			return "", err
		}
		hash.Write(buf[:n])
	}

	checksum := fmt.Sprintf("%x", hash.Sum(nil))
	return checksum, nil
}

// HmacMd5 return the hmac hash of string use md5.
func HmacMd5(str, key string) string {
	h := hmac.New(md5.New, []byte(key))
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum([]byte("")))
}

// HmacMd5WithBase64 return the hmac hash of string use md5 with base64.
func HmacMd5WithBase64(data, key string) string {
	h := hmac.New(md5.New, []byte(key))
	h.Write([]byte(data))
	return base64.StdEncoding.EncodeToString(h.Sum([]byte("")))
}

// HmacSha1 return the hmac hash of string use sha1.
func HmacSha1(str, key string) string {
	h := hmac.New(sha1.New, []byte(key))
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum([]byte("")))
}

// HmacSha1WithBase64 return the hmac hash of string use sha1 with base64.
func HmacSha1WithBase64(str, key string) string {
	h := hmac.New(sha1.New, []byte(key))
	h.Write([]byte(str))
	return base64.StdEncoding.EncodeToString(h.Sum([]byte("")))
}

// HmacSha256 return the hmac hash of string use sha256.
func HmacSha256(str, key string) string {
	h := hmac.New(sha256.New, []byte(key))
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum([]byte("")))
}

// HmacSha256WithBase64 return the hmac hash of string use sha256 with base64.
func HmacSha256WithBase64(str, key string) string {
	h := hmac.New(sha256.New, []byte(key))
	h.Write([]byte(str))
	return base64.StdEncoding.EncodeToString(h.Sum([]byte("")))
}

// HmacSha512 return the hmac hash of string use sha512.
func HmacSha512(str, key string) string {
	h := hmac.New(sha512.New, []byte(key))
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum([]byte("")))
}

// HmacSha512WithBase64 return the hmac hash of string use sha512 with base64.
func HmacSha512WithBase64(str, key string) string {
	h := hmac.New(sha512.New, []byte(key))
	h.Write([]byte(str))
	return base64.StdEncoding.EncodeToString(h.Sum([]byte("")))
}

// Sha1 return the sha1 value (SHA-1 hash algorithm) of string.
func Sha1(str string) string {
	sha1 := sha1.New()
	sha1.Write([]byte(str))
	return hex.EncodeToString(sha1.Sum([]byte("")))
}

// Sha1WithBase64 return the sha1 value (SHA-1 hash algorithm) of base64 string.
func Sha1WithBase64(str string) string {
	sha1 := sha1.New()
	sha1.Write([]byte(str))
	return base64.StdEncoding.EncodeToString(sha1.Sum([]byte("")))
}

// Sha256 return the sha256 value (SHA256 hash algorithm) of string.
func Sha256(str string) string {
	sha256 := sha256.New()
	sha256.Write([]byte(str))
	return hex.EncodeToString(sha256.Sum([]byte("")))
}

// Sha256WithBase64 return the sha256 value (SHA256 hash algorithm) of base64 string.
func Sha256WithBase64(str string) string {
	sha256 := sha256.New()
	sha256.Write([]byte(str))
	return base64.StdEncoding.EncodeToString(sha256.Sum([]byte("")))
}

// Sha512 return the sha512 value (SHA512 hash algorithm) of string.
func Sha512(str string) string {
	sha512 := sha512.New()
	sha512.Write([]byte(str))
	return hex.EncodeToString(sha512.Sum([]byte("")))
}

// Sha512WithBase64 return the sha512 value (SHA512 hash algorithm) of base64 string.
func Sha512WithBase64(str string) string {
	sha512 := sha512.New()
	sha512.Write([]byte(str))
	return base64.StdEncoding.EncodeToString(sha512.Sum([]byte("")))
}
