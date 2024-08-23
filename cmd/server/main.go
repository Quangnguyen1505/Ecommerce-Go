package main

import (
	"github.com/ntquang/ecommerce/internal/routers"
)

func main() {
	r := routers.NewRouter()

	r.Run(":8081")
}
