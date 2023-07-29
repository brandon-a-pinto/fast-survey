package main

import (
	"github.com/brandon-a-pinto/fast-survey/internal/router"
	"github.com/brandon-a-pinto/fast-survey/pkg/mongodb"
)

func main() {
	mongodb.ConnectMongoDB()
	router.Start()
}
