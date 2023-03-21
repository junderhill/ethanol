package main

import (
	"net/http"

	"github.com/junderhill/ethanol"
)

func main() {

	eth := &ethanol.EthanolRouter{}
	http.ListenAndServe(":8080", eth)
}
