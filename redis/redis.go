package redis

import (
	_ "github.com/redis/go-redis"
	"github.com/redis/go-redis/v9"
)

func ConectRedius() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}

//Ping o servidor Redis para verificar se está acessível. 
err := rdb.Ping(ctx).Err()
if err != nil{
	panic(err)
}

// Definir um valor para chave "foo" - set a value for the key "foo"
err = rdb.Set(ctx, "foo", "bar", 0).Err()
if err != nil{
	panic(err)
}


