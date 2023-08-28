package utils

import (
	"encoding/base64"
	"fmt"
	"math/rand"
	cryptorand "crypto/rand"
	"strconv"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

func GetRandomInt(min int, max int) int {
	return r.Intn(max-min) + min
}

func GetRandomNumber(n int) string {
	letters := []rune("0123456789")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[r.Intn(len(letters))]
	}
	return string(b)
}

func GetRandomString(n int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[r.Intn(len(letters))]
	}
	return string(b)
}

func GetRandomHeroID() int32 {
	return int32(r.Intn(85) + 1)
}

func GetShowID(num int64) string {
	return strings.Replace(strconv.FormatInt(num, 10), "0000000", "", -1)
}

func GetRoleName(num int64) string {
	return fmt.Sprintf("塔防精灵%d", num)
}

func GetRandomKey() uint8 {
	return uint8(r.Intn(255) + 1)
}

func GetRandomMachinariumcarID() int32 {
	return int32(r.Intn(13) + 1)
}

func GetFightToken() string {
	// 生成32字节的随机字节片
	randomBytes := make([]byte, 32)
	if _, err := cryptorand.Read(randomBytes); err != nil {
		logrus.Error("rand.Read: ", err)
		return ""
	}

	// 将随机字节片转换为可读的字符串格式
	return base64.StdEncoding.EncodeToString(randomBytes)
}
