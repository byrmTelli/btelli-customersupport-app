package main

import (
	"btelli-customersupport-app/api"
)

func main() {

	server := api.NewAPIServer(":3000")
	server.Run()
}
