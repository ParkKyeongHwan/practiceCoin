package main

import (
	"net/http"
	"text/template"

	"github.com/ParkKyeongHwan/practiceCoin/blockchain"
)

const port string = ":4000"

type homeData struct {
	PageTitle string
	Blocks    []*blockchain.Block
}

func home(rw http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/home.gohtml"))
	data := homeData{"Home", blockchain.GetBlockChain().Blocks}
	tmpl.Execute(rw, data)
}

func main() {
	http.HandleFunc("/", home)
	http.ListenAndServe(port, nil)
}
