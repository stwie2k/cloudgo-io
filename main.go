// cloadgo-io project main.go
package main

import (
	"os"
	"cloudgo-io/service"
)

func main() {
	PORT := os.Getenv("PORT")
	if len(PORT) == 0 {
		PORT = "8080"
	}

	server := service.NewServer()
	server.Run(":" + PORT)


}
