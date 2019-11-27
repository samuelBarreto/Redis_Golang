package main

import (
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis"
)

type Pessoal struct {
	Nome  string `json:"nome"`
	Idade int    `json:"idade"`
}

func main() {
	n := "Samuel"
	i := 20
	fmt.Println("Meu nome é ", n, "e tenho ", i)

	Cliente()

}

func Cliente() {

	cliente := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0, //  usaremos o banco de dados padrão, que é indicado pelo valor 0
	})

	pong, err := cliente.Ping().Result()
	fmt.Println(pong, err)

	json, err := json.Marshal(Pessoal{Nome: "Samuel", Idade: 22})
	if err != nil {
		fmt.Println(err)
	}

	err = cliente.Set("id12", json, 0).Err()
	if err != nil {
		fmt.Println(err)
	}
	val, err := cliente.Get("id12").Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(val)
}
