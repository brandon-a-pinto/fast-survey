package main

import (
	"github.com/brandon-a-pinto/fast-survey/configs/env"
	"github.com/brandon-a-pinto/fast-survey/internal/router"
	"github.com/brandon-a-pinto/fast-survey/pkg/mongodb"
)

func main() {
	env.SetupEnv()
	mongodb.ConnectMongoDB()
	router.Start()
}
