package main

import (
	"flag"

	_ "github.com/harbur/golang-gin-starter/docs"
	"github.com/harbur/golang-gin-starter/pkgs/routers"
)

// @title Golang Starter
// @version 1.0.0
// @description Golang Starter.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.email contact@harbur.io

// @license.name MIT
// @license.url https://github.com/harbur/golang-gin-starter/blob/master/LICENSE

// @BasePath /api/
func main() {
	r := routers.SetupRouter()

	addr := flag.String("address", "localhost:8080", "server address (default: 'localhost:8080')")
	flag.Parse()

	r.Run(*addr)
}
