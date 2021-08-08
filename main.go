package main

import (
	explorer "github.com/ParkKyeongHwan/practiceCoin/explorer/templates"
	"github.com/ParkKyeongHwan/practiceCoin/rest"
)

func main() {
	go explorer.Start(3000)
	rest.Start(4000)
}
