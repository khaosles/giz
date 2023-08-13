package cryptor

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
)

/*
   @File: rsa.go
   @Author: khaosles
   @Time: 2023/8/13 13:24
   @Desc:
*/

// RsaEncrypt encrypt data with ras algorithm to match java
func RsaEncrypt(data, publicKey string) (string, error) {

	publicKey = fmt.Sprintf("-----BEGIN PUBLIC KEY-----\n%s\n-----END PUBLIC KEY-----", publicKey)
	block, _ := pem.Decode([]byte(publicKey))

	if block == nil {
		return "", errors.New("public key error")
	}

	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return "", err
	}

	pub := pubInterface.(*rsa.PublicKey)
	encryptPKCS1v15, err := rsa.EncryptPKCS1v15(rand.Reader, pub, []byte(data))
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(encryptPKCS1v15), nil
}

// RsaDecrypt decrypt data with ras algorithm to match java
func RsaDecrypt(data, privateKey string) (string, error) {

	privateKey = fmt.Sprintf("-----BEGIN PRIVATE KEY-----\n%s\n-----END PRIVATE KEY-----", privateKey)
	block, _ := pem.Decode([]byte(privateKey))
	if block == nil {
		return "", errors.New("private key error")
	}

	privInterface, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return "", err
	}

	priv := privInterface.(*rsa.PrivateKey)
	decodeString, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return "", err
	}

	decryptPKCS1v15, err := rsa.DecryptPKCS1v15(rand.Reader, priv, decodeString)
	if err != nil {
		return "", err
	}

	return string(decryptPKCS1v15), nil
}

// RsaSign sign to match java
func RsaSign(data, privateKey string) (string, error) {

	privateKey = fmt.Sprintf("-----BEGIN PRIVATE KEY-----\n%s\n-----END PRIVATE KEY-----", privateKey)
	block, _ := pem.Decode([]byte(privateKey))
	if block == nil {
		return "", errors.New("private key error")
	}

	privInterface, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return "", err
	}
	priv := privInterface.(*rsa.PrivateKey)

	h := crypto.SHA256
	hn := h.New()
	hn.Write([]byte(data))
	sum := hn.Sum(nil)

	signPKCS1v15, err := rsa.SignPKCS1v15(rand.Reader, priv, h, sum)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(signPKCS1v15), nil
}

// RsaVerify verify sign to match java
func RsaVerify(data, sign, publicKey string) error {

	publicKey = fmt.Sprintf("-----BEGIN PUBLIC KEY-----\n%s\n-----END PUBLIC KEY-----", publicKey)
	block, _ := pem.Decode([]byte(publicKey))
	if block == nil {
		return errors.New("public key error!")
	}

	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return err
	}

	pub := pubInterface.(*rsa.PublicKey)
	h := crypto.SHA256
	hn := h.New()
	hn.Write([]byte(data))
	sum := hn.Sum(nil)

	signData, err := base64.StdEncoding.DecodeString(sign)
	if err != nil {
		return err
	}

	return rsa.VerifyPKCS1v15(pub, h, sum, signData)
}

// GenKey gen key
func GenKey(bits int) (string, string) {
	privateKey, _ := rsa.GenerateKey(rand.Reader, bits)
	publicKey := &privateKey.PublicKey

	bytePri, _ := x509.MarshalPKCS8PrivateKey(privateKey)
	pri := base64.StdEncoding.EncodeToString(bytePri)

	bytePub, _ := x509.MarshalPKIXPublicKey(publicKey)
	pub := base64.StdEncoding.EncodeToString(bytePub)

	fmt.Println("Public key: ", pri)
	fmt.Println()
	fmt.Println("Private key: ", pub)
	return pub, pri
}
