package main

import "GinGo/Routers"

func main() {
	var PORT = ":300"

	Routers.StartServer().Run(PORT)
}
