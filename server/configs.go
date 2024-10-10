package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/spf13/viper"
)

func initConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	viper.ReadInConfig()
}

func initTimeZone() {
	ict, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
	time.Local = ict
}
