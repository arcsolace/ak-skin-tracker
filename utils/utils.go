package utils

import (
	"math/rand"
	"time"
)

func GenerateRandomUserCode(length int) string {
    const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
    seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))

    userCode := make([]byte, length)
    for i := 0; i < length; i++ {
        userCode[i] = charset[seededRand.Intn(len(charset))]
    }

    return string(userCode)
}