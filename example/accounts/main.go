package main

import (
	"fmt"
	"log"

	"github.com/javadmohebbi/gobdgz"
)

func main() {
	gz := gobdgz.GravityZoneAPI{
		BaseURL: "https://172.16.50.105/api/v1.0/jsonrpc/",
	}
	r := gobdgz.Request{
		APIKey: "c24428a799d5c255e7d88ea828c66e42921a0f09ca8c352c05bf75bc005be721",
		// Debug:  true,
		// ID:         "0",
		JSONRPC:    "2.0",
		URL:        gz.BaseURL + "/accounts",
		HttpMethod: "POST",
	}

	getAccountsList(&gz, r)

}

func getAccountsList(gz *gobdgz.GravityZoneAPI, rq gobdgz.Request) {
	r := rq
	r.Method = "getAccountsList"
	gz.Accounts.SetRequest(r)
	resp, err := gz.Accounts.GetAccounList()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Accuonts List:")
	for k, v := range resp.Result.Items {
		fmt.Printf("\tRow: %v -> %v\n", k+1, v)
	}
}
