package main

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	MyVar string `mapstructure:"MY_VAR"`
}

func main() {
	// cwd, err := os.Getwd()
    // if err != nil {
    //     return nil, apperror.NewError(apperror.InternalError, err)
    // }

    // envPath := filepath.Join(cwd, ".env")

	viper.AutomaticEnv()
	// 実際のプロジェクトでは、引数にenvPathを渡す
	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		panic(err)
	}

	fmt.Println("MY_VAR:", config.MyVar)
}
