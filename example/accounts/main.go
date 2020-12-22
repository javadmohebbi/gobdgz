package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/javadmohebbi/gobdgz"
)

func main() {
	// create new instance of GravityZone API
	gz := gobdgz.GravityZoneAPI{
		BaseURL: "https://172.16.50.105/api/v1.0/jsonrpc/",
	}

	// Create Default Request struct
	r := gobdgz.Request{
		// API KEY
		// must be enabled in GravityZone Control Center
		// http://download.bitdefender.com/business/API/Bitdefender_GravityZone_On-Premises_APIGuide_enUS.pdf#page=7&zoom=100,33,85
		APIKey: "c24428a799d5c255e7d88ea828c66e42921a0f09ca8c352c05bf75bc005be721",

		// set it to true if need debug information
		Debug: true,

		// JSON-RPC needs id, set it to the needed ID, UUID, NULL or leave it empty
		// and we will take care of everythings
		// ID:         "0",

		// JSON RPC version
		JSONRPC: "2.0",

		// URL
		URL: gz.BaseURL + "/accounts",

		// HTTP MEthods
		HttpMethod: "POST",
	}

	kind := flag.String("kind", "getAccountsList", "methods to call: possible values are: getAccountsList, deleteAccount, createAccount, updateAccount, configureNotificationsSettings, getNotificationsSettings")

	flag.Parse()

	switch *kind {
	case "getAccountsList":
		getAccountsList(&gz, r)
	case "createAccount":
		createAccount(&gz, r)
	default:
		flag.PrintDefaults()
	}

}

// get all accounts & print them to console
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

// create a new account
func createAccount(gz *gobdgz.GravityZoneAPI, rq gobdgz.Request) {
	r := rq
	r.Method = "createAccount"

	rand.Seed(time.Now().UnixNano())
	usr := fmt.Sprintf("user-%v", rand.Intn(20000))
	r.Params = map[string]interface{}{
		"email":    usr + "@example.org",
		"userName": usr,
		"profile": map[string]interface{}{
			"fullName": "full name " + usr,
			"timezone": "Asia/Tehran",
			"language": "en_US",
		},
		"password": "1234!@#$qwerQWER",
		"role":     5, // 1 = company admin, 2 = net admin, 3 = Security Analyst, 5 = custom
		"rights": map[string]interface{}{
			"manageCompanies": false,
			"manageNetworks":  false,
			"manageUsers":     false,
			"manageReports":   true,
			"companyManager":  false,
		},
		// "targetIds": "" // id for targets to be manage by this user
	}

	gz.Accounts.SetRequest(r)
	resp, err := gz.Accounts.CreateAccount()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("\n\tUser '%v' with ID '%v' has created!\n", usr, resp.Result)

}
