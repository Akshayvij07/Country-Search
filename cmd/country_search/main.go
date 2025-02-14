package main

import (
	"github.com/Akshayvij07/country-search/internals/di"
)

func main() {
	server := di.ConfigureServer()
	server.Serve()
}
