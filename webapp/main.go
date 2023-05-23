package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Rodando webApp!")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
