package main

import (
	"fmt"

	"github.com/Gabriel-Newton-dev/server_redis_golang/redis_c"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()
	redis_c.ConectRedius()
	fmt.Println("Inicializando o servidor redius.")
}
