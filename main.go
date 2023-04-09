package main

import "k-style-test/router"

func main() {
	router.StartServer().Run(":3000")
}
