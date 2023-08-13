package gen

import (
	crand "crypto/rand"
	"io"
	"math/big"
	"math/rand"
	"strings"

	"github.com/lithammer/shortuuid/v3"
	gouuid "github.com/satori/go.uuid"
)

/*
   @File: gen.go
   @Author: khaosles
   @Time: 2023/8/13 13:03
   @Desc:
*/

func Uuid() string {
	return gouuid.NewV4().String()
}

func UuidShort() string {
	return shortuuid.New()
}

func UuidNoSeparator() string {
	return strings.Replace(Uuid(), "-", "", -1)
}

func RandBytes(length int) []byte {
	if length < 1 {
		return []byte{}
	}
	b := make([]byte, length)
	if _, err := io.ReadFull(crand.Reader, b); err != nil {
		return nil
	}
	return b
}

func RandInt(min, max int64) int64 {
	// calculate the max we will be using
	bg := big.NewInt(max - min)

	// get big.Int between 0 and bg
	// in this case 0 to 20
	n, err := crand.Int(crand.Reader, bg)
	if err != nil {
		panic(err)
	}

	// add n to min to support the passed in range
	return n.Int64() + min
}

const (
	LettersLetter          = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	LettersUpperCaseLetter = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	LettersLowwerLetters   = "abcdefghijklmnopqrstuvwxyz"
	LettersNumber          = "0123456789"
	LettersNumberNoZero    = "123456789"
	LettersSymbol          = "~`!@#$%^&*()_-+={[}]|\\:;\"'<,>.?/"
)

func RandString(n int, letters ...string) (string, error) {
	lettersDefaultValue := LettersLetter + LettersNumber + LettersSymbol
	if len(letters) > 0 {
		lettersDefaultValue = ""
		for _, letter := range letters {
			lettersDefaultValue = lettersDefaultValue + letter
		}
	}
	bytes := make([]byte, n)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	for i, b := range bytes {
		bytes[i] = lettersDefaultValue[b%byte(len(lettersDefaultValue))]
	}

	return string(bytes), nil
}

func RandUniqueIntSlice(n, min, max int) []int {
	if min > max {
		return []int{}
	}
	if n > max-min {
		n = max - min
	}

	nums := make([]int, n)
	used := make(map[int]struct{}, n)
	for i := 0; i < n; {
		r := RandInt(int64(min), int64(max))
		if _, use := used[int(r)]; use {
			continue
		}
		used[int(r)] = struct{}{}
		nums[i] = int(r)
		i++
	}

	return nums
}
