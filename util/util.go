package util

import (
	"math/rand"
	"time"
)

func Randomstring(n int) string {
	var letters = []byte("qwertyuioplkjhgfdsamnbvcxzMNBVCXZLKJHGFDSAPOIUYTREWQ")
	result := make([]byte, n)

	rand.Seed(time.Now().Unix())
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}
