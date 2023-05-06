package utils

import (
	"encoding/base64"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

func GetRandomNumber(n int) string {
	letters := []rune("0123456789")

	rand.Seed(time.Now().UnixNano())
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func GetRandomString(n int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	rand.Seed(time.Now().UnixNano())
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func GetShowID(num int64) string {
	return strings.Replace(strconv.FormatInt(num, 10), "0000000", "", -1)
}

func GetRoleName(num int64) string {
	return fmt.Sprintf("塔防精灵%d", num)
}

func GetRandomKey() uint8 {
	rand.Seed(time.Now().UnixNano())
	return uint8(rand.Intn(255) + 1)
}

func GetRandomMachinariumcarID() int32 {
	rand.Seed(time.Now().UnixNano())
	return int32(rand.Intn(13) + 1)
}

func GetFightToken() string {
	// 生成32字节的随机字节片
	randomBytes := make([]byte, 32)
	if _, err := rand.Read(randomBytes); err != nil {
		logrus.Error("rand.Read: ", err)
		return ""
	}

	// 将随机字节片转换为可读的字符串格式
	return base64.StdEncoding.EncodeToString(randomBytes)
}
