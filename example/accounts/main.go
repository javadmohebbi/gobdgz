package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
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
		// Debug: true,

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

	// method kind command line argument
	kind := flag.String("method", "getAccountsList", "methods to call: possible values are: getAccountsList, deleteAccount, createAccount, updateAccount, configureNotificationsSettings, getNotificationsSettings")

	// parsing flags
	flag.Parse()

	switch *kind {
	case "getAccountsList":
		getAccountsList(&gz, r)
	case "createAccount":
		createAccount(&gz, r)
	case "updateAccount":
		updateAccount(&gz, r)
	case "getNotificationsSettings":
		getNotificationsSettings(&gz, r)
	case "configureNotificationsSettings":
		configureNotificationsSettings(&gz, r)
	case "deleteAccount":
		deleteAccount(&gz, r)
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
		// "targetIds": "" // []string // ids for targets to be manage by this user
	}

	gz.Accounts.SetRequest(r)
	resp, err := gz.Accounts.CreateAccount()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("\n\tUser '%v' with ID '%v' has created!\n", usr, resp.Result)

}

// update an account
func updateAccount(gz *gobdgz.GravityZoneAPI, rq gobdgz.Request) {
	r := rq
	r.Method = "updateAccount"

	// read from input
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Account ID: ")
	accountID, _ := reader.ReadString('\n')
	accountID = strings.Trim(accountID, " \n")

	rand.Seed(time.Now().UnixNano())
	usr := fmt.Sprintf("user-%v", rand.Intn(20000))
	r.Params = map[string]interface{}{
		"accountId": accountID,
		"email":     usr + "@example-update.org",
		"profile": map[string]interface{}{
			"fullName": "full name updated",
			"timezone": "Asia/Tehran",
			"language": "en_US",
		},
		"password": "1234!@#$qwerQWER" + fmt.Sprintf("-%v", rand.Intn(40000000000)),
		"role":     3, // 1 = company admin, 2 = net admin, 3 = Security Analyst, 5 = custom
		"rights": map[string]interface{}{
			"manageCompanies": false,
			"manageNetworks":  false,
			"manageUsers":     false,
			"manageReports":   false,
			"companyManager":  false,
		},
		// "targetIds": ""  // []string // ids for targets to be manage by this user
	}

	gz.Accounts.SetRequest(r)
	resp, err := gz.Accounts.UpdateAccount()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("\n\tUser ID '%v' has updated! The result is '%v'\n", accountID, resp.Result)

}

// Get Notifications Settings
func getNotificationsSettings(gz *gobdgz.GravityZoneAPI, rq gobdgz.Request) {
	r := rq
	r.Method = "getNotificationsSettings"

	// read from input
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Account ID: ")
	accountID, _ := reader.ReadString('\n')
	accountID = strings.Trim(accountID, " \n")

	r.Params = map[string]interface{}{
		"accountId": accountID,
	}

	gz.Accounts.SetRequest(r)
	resp, err := gz.Accounts.GetNotificationsSettings()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\nNotifications for User ID '%v' has got!", accountID)
	fmt.Printf("\n\tNotificationStruct: %v\n\n", resp)

}

// configure Notifications Settings
func configureNotificationsSettings(gz *gobdgz.GravityZoneAPI, rq gobdgz.Request) {
	r := rq
	r.Method = "configureNotificationsSettings"

	// read from input
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Account ID: ")
	accountID, _ := reader.ReadString('\n')
	accountID = strings.Trim(accountID, " \n")

	r.Params = map[string]interface{}{
		"accountId":         accountID,
		"deleteAfter":       60,
		"includeDeviceName": true,
		"includeDeviceFQDN": true,
		"emailAddresses":    []string{"example1@example.com"},
		"notificationsSettings": []map[string]interface{}{
			map[string]interface{}{

				// for type value this guide should be considered
				// http://download.bitdefender.com/business/API/Bitdefender_GravityZone_On-Premises_APIGuide_enUS.pdf#page=25&zoom=100,33,438
				"type":    23,
				"enabled": true,
				"visibilitySettings": map[string]interface{}{
					"sendPerEmail":               true,
					"showInConsole":              true,
					"useCustomEmailDistribution": false,
					"emails":                     []string{},
					"logToServer":                true,
				},

				// to understand relation between type & notification
				// settings you should read this guide
				// http://download.bitdefender.com/business/API/Bitdefender_GravityZone_On-Premises_APIGuide_enUS.pdf#page=27&zoom=100,33,409
				"configurationSettings": map[string]interface{}{
					// "threshold":    15,
					// "useThreshold": false,
					"notUpdated": true,
					"reboot":     true,
					"stopped":    true,
				},
			},
		},
	}

	gz.Accounts.SetRequest(r)
	resp, err := gz.Accounts.ConfigureNotificationsSettings()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("\n\tNotification settings for User ID '%v' has updated! The result is '%v'\n", accountID, resp.Result)

}

// delete an Account
func deleteAccount(gz *gobdgz.GravityZoneAPI, rq gobdgz.Request) {
	r := rq
	r.Method = "deleteAccount"

	// read from input
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Account ID: ")
	accountID, _ := reader.ReadString('\n')
	accountID = strings.Trim(accountID, " \n")

	r.Params = map[string]interface{}{
		"accountId": accountID,
	}

	gz.Accounts.SetRequest(r)
	resp, err := gz.Accounts.GetNotificationsSettings()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\nUser ID '%v' has deleted with result '%v'!\n", accountID, resp.Result)

}
