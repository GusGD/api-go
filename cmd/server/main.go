package main

import "github.com/gusgd/apigo/configs"

func main() {
	configs, _ := configs.LoadConfig(".")
	println(configs.DBDriver)
}
