package main

import "github.com/vinigracindo/gowitter/cmd/api"

func main() {
	apiServer := api.NewServer(8080)
	apiServer.Run()
}
