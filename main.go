package main

import (
	"fmt"
	"github.com/gaomugong/go-netdisk/routers"
)

func main() {
	port := 5000
	fmt.Println("go-netdisk begin server at port: ", port)

	r := routers.SetupRouter()
	r.Run(fmt.Sprintf("127.0.0.1:%d", port))
}
