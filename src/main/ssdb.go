package main

import (
	"log"

	"github.com/seefan/gossdb"
	"github.com/seefan/gossdb/conf"
)

func main() {
	pool, err := gossdb.NewPool(&conf.Config{
		Host:             "127.0.0.1",
		Port:             8888,
		MinPoolSize:      5,
		MaxPoolSize:      50,
		AcquireIncrement: 5,
	})

	if err != nil {
		log.Fatal(err)
		return
	}

	c, err := pool.NewClient()
	if err != nil {
		log.Println(err.Error())
	}

	defer c.Close()

	setErr := c.Hset("test", "test1", "hello world")
	if setErr != nil {
		log.Println(setErr)
	}
	re, err := c.HgetAll("test")
	if err != nil {
		log.Println(err)
	} else {
		log.Println(re)
	}
}
