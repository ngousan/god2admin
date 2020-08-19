package main

import "god2admin/core"

func main() {

	defer core.StopServer()
	core.StartServer()

}
