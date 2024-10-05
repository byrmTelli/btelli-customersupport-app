package main

import (
	"btelli-customersupport-api/api"
)

func main() {

	server := api.NewAPIServer(":3000")
	server.Run()
}
