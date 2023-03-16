package redis_c

import (
	_ "github.com/redis/go-redis"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

func ConectRedius() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "viper.Db_Password", // no password set
		DB:       "viper.Db_database",  // use default DB
	})
}

//Ping o servidor Redis para verificar se está acessível. 
er := rdb.Ping(ctx).Err()
if err != nil{
	panic(err)
}

// Definir um valor para chave "foo" - set a value for the key "foo"
err = rdb.Set(ctx, "foo", "bar", 0).Err()
if err != nil{
	panic(err)
}


