package main

import (
	"fmt"
	"log"
	"os"

	"github.com/AthithyanR/portfolio/internal/routes"
	"github.com/valyala/fasthttp"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	port_addr := fmt.Sprintf(":%s", port)

	r := routes.NewRouter()

	log.Println("Listening on PORT", port)
	log.Fatal(fasthttp.ListenAndServe(port_addr, r.Handler))
}
