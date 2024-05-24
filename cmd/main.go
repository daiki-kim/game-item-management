package main

import (
	"game-item-management/routes"
)

func main() {
	r := routes.ItemRouter()
	r.Run("localhost:8080")
}
