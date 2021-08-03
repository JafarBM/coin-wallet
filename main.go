package main

import (
	"fmt"
	"net/http"
	"request_test/tamrin1/core"
	"request_test/tamrin1/transport"
)

func main() {

	walletService := core.GetNewWalletService()
	_ = transport.GetNewRestService(walletService)

	if err := http.ListenAndServe("127.0.0.1:8000", nil); err != nil {
		fmt.Println(err)
		panic(err)
	}
}
