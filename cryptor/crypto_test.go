package cryptor

import (
	"fmt"
	"testing"
)

/*
   @File: crypto_test.go
   @Author: khaosles
   @Time: 2023/8/13 13:34
   @Desc:
*/

func TestAesEcbEncrypt(t *testing.T) {
	data := "hello"
	key := "abcdefghijklmnop"

	encrypted := AesEcbEncrypt([]byte(data), []byte(key))

	decrypted := AesEcbDecrypt(encrypted, []byte(key))

	fmt.Println(string(decrypted))

	// Output:
	// hello
}

func TestAesEcbDecrypt(t *testing.T) {
	data := "hello"
	key := "abcdefghijklmnop"

	encrypted := AesEcbEncrypt([]byte(data), []byte(key))

	decrypted := AesEcbDecrypt(encrypted, []byte(key))

	fmt.Println(string(decrypted))

	// Output:
	// hello
}

func TestAesCbcEncrypt(t *testing.T) {
	data := "hello"
	key := "abcdefghijklmnop"

	encrypted := AesCbcEncrypt([]byte(data), []byte(key))

	decrypted := AesCbcDecrypt(encrypted, []byte(key))

	fmt.Println(string(decrypted))

	// Output:
	// hello
}

func TestAesCbcDecrypt(t *testing.T) {
	data := "hello"
	key := "abcdefghijklmnop"

	encrypted := AesCbcEncrypt([]byte(data), []byte(key))

	decrypted := AesCbcDecrypt(encrypted, []byte(key))

	fmt.Println(string(decrypted))

	// Output:
	// hello
}

func TestAesCtrCrypt(t *testing.T) {
	data := "hello"
	key := "abcdefghijklmnop"

	encrypted := AesCtrCrypt([]byte(data), []byte(key))
	decrypted := AesCtrCrypt(encrypted, []byte(key))

	fmt.Println(string(decrypted))

	// Output:
	// hello
}

func TestAesCfbEncrypt(t *testing.T) {
	data := "hello"
	key := "abcdefghijklmnop"

	encrypted := AesCfbEncrypt([]byte(data), []byte(key))
	decrypted := AesCfbDecrypt(encrypted, []byte(key))

	fmt.Println(string(decrypted))

	// Output:
	// hello
}

func TestAesCfbDecrypt(t *testing.T) {
	data := "hello"
	key := "abcdefghijklmnop"

	encrypted := AesCfbEncrypt([]byte(data), []byte(key))

	decrypted := AesCfbDecrypt(encrypted, []byte(key))

	fmt.Println(string(decrypted))

	// Output:
	// hello
}

func TestAesOfbEncrypt(t *testing.T) {
	data := "hello"
	key := "abcdefghijklmnop"

	encrypted := AesOfbEncrypt([]byte(data), []byte(key))

	decrypted := AesOfbDecrypt(encrypted, []byte(key))

	fmt.Println(string(decrypted))

	// Output:
	// hello
}

func TestAesOfbDecrypt(t *testing.T) {
	data := "hello"
	key := "abcdefghijklmnop"

	encrypted := AesOfbEncrypt([]byte(data), []byte(key))

	decrypted := AesOfbDecrypt(encrypted, []byte(key))

	fmt.Println(string(decrypted))

	// Output:
	// hello
}

func TestDesEcbEncrypt(t *testing.T) {
	data := "hello"
	key := "abcdefgh"

	encrypted := DesEcbEncrypt([]byte(data), []byte(key))

	decrypted := DesEcbDecrypt(encrypted, []byte(key))

	fmt.Println(string(decrypted))

	// Output:
	// hello
}

func TestDesEcbDecrypt(t *testing.T) {
	data := "hello"
	key := "abcdefgh"

	encrypted := DesEcbEncrypt([]byte(data), []byte(key))

	decrypted := DesEcbDecrypt(encrypted, []byte(key))

	fmt.Println(string(decrypted))

	// Output:
	// hello
}

func TestDesCbcEncrypt(t *testing.T) {
	data := "hello"
	key := "abcdefgh"

	encrypted := DesCbcEncrypt([]byte(data), []byte(key))

	decrypted := DesCbcDecrypt(encrypted, []byte(key))

	fmt.Println(string(decrypted))

	// Output:
	// hello
}

func TestDesCbcDecrypt(t *testing.T) {
	data := "hello"
	key := "abcdefgh"

	encrypted := DesCbcEncrypt([]byte(data), []byte(key))

	decrypted := DesCbcDecrypt(encrypted, []byte(key))

	fmt.Println(string(decrypted))

	// Output:
	// hello
}

func TestDesCtrCrypt(t *testing.T) {
	data := "hello"
	key := "abcdefgh"

	encrypted := DesCtrCrypt([]byte(data), []byte(key))
	decrypted := DesCtrCrypt(encrypted, []byte(key))

	fmt.Println(string(decrypted))

	// Output:
	// hello
}

func TestDesCfbEncrypt(t *testing.T) {
	data := "hello"
	key := "abcdefgh"

	encrypted := DesCfbEncrypt([]byte(data), []byte(key))

	decrypted := DesCfbDecrypt(encrypted, []byte(key))

	fmt.Println(string(decrypted))

	// Output:
	// hello
}

func TestDesCfbDecrypt(t *testing.T) {
	data := "hello"
	key := "abcdefgh"

	encrypted := DesCfbEncrypt([]byte(data), []byte(key))

	decrypted := DesCfbDecrypt(encrypted, []byte(key))

	fmt.Println(string(decrypted))

	// Output:
	// hello
}

func TestDesOfbEncrypt(t *testing.T) {
	data := "hello"
	key := "abcdefgh"

	encrypted := DesOfbEncrypt([]byte(data), []byte(key))

	decrypted := DesOfbDecrypt(encrypted, []byte(key))

	fmt.Println(string(decrypted))

	// Output:
	// hello
}

func TestDesOfbDecrypt(t *testing.T) {
	data := "hello"
	key := "abcdefgh"

	encrypted := DesOfbEncrypt([]byte(data), []byte(key))

	decrypted := DesOfbDecrypt(encrypted, []byte(key))

	fmt.Println(string(decrypted))

	// Output:
	// hello
}

func TestBase64StdEncode(t *testing.T) {
	base64Str := Base64StdEncode("hello")

	fmt.Println(base64Str)

	// Output:
	// aGVsbG8=
}

func TestBase64StdDecode(t *testing.T) {
	str := Base64StdDecode("aGVsbG8=")

	fmt.Println(str)

	// Output:
	// hello
}

func TestHmacMd5(t *testing.T) {
	str := "hello"
	key := "12345"

	hms := HmacMd5(str, key)
	fmt.Println(hms)

	// Output:
	// e834306eab892d872525d4918a7a639a
}

func TestHmacMd5WithBase64(t *testing.T) {
	str := "hello"
	key := "12345"

	hms := HmacMd5WithBase64(str, key)
	fmt.Println(hms)

	// Output:
	// 6DQwbquJLYclJdSRinpjmg==
}

func TestHmacSha1(t *testing.T) {
	str := "hello"
	key := "12345"

	hms := HmacSha1(str, key)
	fmt.Println(hms)

	// Output:
	// 5c6a9db0cccb92e36ed0323fd09b7f936de9ace0
}

func TestHmacSha1WithBase64(t *testing.T) {
	str := "hello"
	key := "12345"

	hms := HmacSha1WithBase64(str, key)
	fmt.Println(hms)

	// Output:
	// XGqdsMzLkuNu0DI/0Jt/k23prOA=
}

func TestHmacSha256(t *testing.T) {
	str := "hello"
	key := "12345"

	hms := HmacSha256(str, key)
	fmt.Println(hms)

	// Output:
	// 315bb93c4e989862ba09cb62e05d73a5f376cb36f0d786edab0c320d059fde75
}

func TestHmacSha256WithBase64(t *testing.T) {
	str := "hello"
	key := "12345"

	hms := HmacSha256WithBase64(str, key)
	fmt.Println(hms)

	// Output:
	// MVu5PE6YmGK6Ccti4F1zpfN2yzbw14btqwwyDQWf3nU=
}

func TestHmacSha512(t *testing.T) {
	str := "hello"
	key := "12345"

	hms := HmacSha512(str, key)
	fmt.Println(hms)

	// Output:
	// dd8f1290a9dd23d354e2526d9a2e9ce8cffffdd37cb320800d1c6c13d2efc363288376a196c5458daf53f8e1aa6b45a6d856303d5c0a2064bff9785861d48cfc
}

func TestHmacSha512WithBase64(t *testing.T) {
	str := "hello"
	key := "12345"

	hms := HmacSha512WithBase64(str, key)
	fmt.Println(hms)

	// Output:
	// 3Y8SkKndI9NU4lJtmi6c6M///dN8syCADRxsE9Lvw2Mog3ahlsVFja9T+OGqa0Wm2FYwPVwKIGS/+XhYYdSM/A==
}

func TestMd5String(t *testing.T) {
	md5Str := Md5String("hello")
	fmt.Println(md5Str)

	// Output:
	// 5d41402abc4b2a76b9719d911017c592
}

func TestMd5StringWithBase64(t *testing.T) {
	md5Str := Md5StringWithBase64("hello")
	fmt.Println(md5Str)

	// Output:
	// XUFAKrxLKna5cZ2REBfFkg==
}

func TestMd5Byte(t *testing.T) {
	md5Str := Md5Byte([]byte{'a'})
	fmt.Println(md5Str)

	// Output:
	// 0cc175b9c0f1b6a831c399e269772661
}

func TestMd5ByteWithBase64(t *testing.T) {
	md5Str := Md5ByteWithBase64([]byte("hello"))
	fmt.Println(md5Str)

	// Output:
	// XUFAKrxLKna5cZ2REBfFkg==
}

func TestSha1(t *testing.T) {
	result := Sha1("hello")
	fmt.Println(result)

	// Output:
	// aaf4c61ddcc5e8a2dabede0f3b482cd9aea9434d
}

func TestSha1WithBase64(t *testing.T) {
	result := Sha1WithBase64("hello")
	fmt.Println(result)

	// Output:
	// qvTGHdzF6KLavt4PO0gs2a6pQ00=
}

func TestSha256(t *testing.T) {
	result := Sha256("hello")
	fmt.Println(result)

	// Output:
	// 2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824
}

func TestSha256WithBase64(t *testing.T) {
	result := Sha256WithBase64("hello")
	fmt.Println(result)

	// Output:
	// LPJNul+wow4m6DsqxbninhsWHlwfp0JecwQzYpOLmCQ=
}

func TestSha512(t *testing.T) {
	result := Sha512("hello")
	fmt.Println(result)

	// Output:
	// 9b71d224bd62f3785d96d46ad3ea3d73319bfbc2890caadae2dff72519673ca72323c3d99ba5c11d7c7acc6e14b8c5da0c4663475c2e5c3adef46f73bcdec043
}

func TestSha512WithBase64(t *testing.T) {
	result := Sha512WithBase64("hello")
	fmt.Println(result)

	// Output:
	// m3HSJL1i83hdltRq0+o9czGb+8KJDKra4t/3JRlnPKcjI8PZm6XBHXx6zG4UuMXaDEZjR1wuXDre9G9zvN7AQw==
}
