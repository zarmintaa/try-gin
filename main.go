package main

import "belajar-golang-restful-api/routers"

func main() {
	var PORT = ":8080"

	routers.StartServer().Run(PORT)
}
