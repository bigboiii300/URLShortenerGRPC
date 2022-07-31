package utils

import (
	"github.com/spf13/viper"
	"math/rand"
	"time"
)

var chars string
var lenShortURL int

func GenerateShortURL() string {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, lenShortURL)
	for i := 0; i < lenShortURL; i++ {
		b[i] = chars[rand.Intn(len(chars))]
	}
	return string(b)
}

func InitViper() error {
	viper.SetConfigFile("./config.yml")
	err := viper.ReadInConfig()
	chars = viper.GetString("shortUrlGenerate.chars")
	lenShortURL = viper.GetInt("shortUrlGenerate.len")
	return err
}
