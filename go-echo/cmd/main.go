package main

import (
	"boilerplates/pkg/adapters"
	"boilerplates/pkg/infrastructures/server"
)

func init() {
	// This will be executed first
	adapters.IocConfigs()
	adapters.IoCMapper()
	adapters.IoCCache()
	adapters.IoCLogger()
	adapters.IoCCrypto()
	adapters.IoCDatabase()
}

func main() {
	// This will be executed second
	server.StartEchoServer()
}
