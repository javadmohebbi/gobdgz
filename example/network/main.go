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
		APIKey: "373fb67c5f28940929beb7f9cfa3504d310ea71fa7d90b2a13bda448f861fa9c",

		// set it to true if need debug information
		Debug: true,

		// JSON-RPC needs id, set it to the needed ID, UUID, NULL or leave it empty
		// and we will take care of everythings
		// ID:         "0",

		// JSON RPC version
		JSONRPC: "2.0",

		// URL
		URL: gz.BaseURL + "/network",

		// HTTP MEthods
		HttpMethod: "POST",
	}

	// method kind command line argument
	kind := flag.String("method", "getContainers", "methods to call: possible values are: getContainers, getNetworkInventoryItems, createScanTask, createReconfigureClientTask, getScanTasksList, getEndpointsList, getManagedEndpointDetails, createCustomGroup, deleteCustomGroup, moveCustomGroup, moveEndpoints, deleteEndpoint, setEndpointLabel, createScanTaskByMac, assignPolicy")
	containerService := flag.String("service", "computers", "service to call: possible values are: computers, virtualmachines, mobile (mobile avaialble ONLY for getContainers)")

	// parsing flags
	flag.Parse()

	switch *kind {
	case "getContainers":
		r.URL += "/" + *containerService
		getContainers(&gz, r)
	case "getNetworkInventoryItems":
		r.URL += "/" + *containerService
		getNetworkInventoryItems(&gz, r)
	case "createScanTask":
		r.URL += "/" + *containerService
		createScanTask(&gz, r)
	case "createReconfigureClientTask":
		r.URL += "/" + *containerService
		createReconfigureClientTask(&gz, r)
	default:
		flag.PrintDefaults()
	}

}

// Returns the network containers.
func getContainers(gz *gobdgz.GravityZoneAPI, rq gobdgz.Request) {
	r := rq
	r.Method = "getContainers"

	r.Params = map[string]interface{}{

		// The ID of the container. If null, the top containers of the specified service type will be returned.
		// "parentId": 4,

		// The ID of the view type for the virtual environment inventory. The view type depends on the
		// virtualization platform. In VMWare integrations,
		// the available options are:
		//		● 1 - Hosts and Clusters view (default)
		//		● 2 - Virtual Machines view.
		// In Citrix, XenServer integrations, the available
		// options are:
		// 		● 3 - Server view (default)
		//		● 4 - Folder view.
		// "viewType": 4,

	}

	gz.Network.SetRequest(r)
	resp, err := gz.Network.GetContainers()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Containers List:")
	for k, v := range resp.Result {
		fmt.Printf("\tRow: %v -> %v=%v\n", k+1, v.ID, v.Name)
	}
}

// This method returns network inventory items.
func getNetworkInventoryItems(gz *gobdgz.GravityZoneAPI, rq gobdgz.Request) {
	r := rq
	r.Method = "getNetworkInventoryItems"

	r.Params = map[string]interface{}{
		// "parentId": "5fc4b4403212d72d9824060b",

		//"filters": map[string]interface{}{} //http://download.bitdefender.com/business/API/Bitdefender_GravityZone_On-Premises_APIGuide_enUS.pdf#available%20filters
		"filters": map[string]interface{}{
			"depth": map[string]interface{}{
				"allItemsRecursively": true,
			},
			"security": map[string]interface{}{
				"management": map[string]interface{}{
					"managedWithBest": true,
					"managedRelays":   true,
					"securityServers": true,
				},
			},
		},

		//"viewType": 1, 2, 3, 4
		//"page": 1,
		//"perPage": 30
	}

	gz.Network.SetRequest(r)
	resp, err := gz.Network.GetNetworkInventoryItems()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inventory Items:")
	for k, v := range resp.Result.Items {
		fmt.Printf("\tRow: %v -> %v='%v',type='%v',parentId='%v'\n", k+1, v.ID, v.Name, v.Type, v.ParentID)
		if v.Type == 5 || v.Type == 6 {
			fmt.Printf("\n\t\t  --> (Details): label='%v',fqdn='%v',groupId='%v',isManaged='%v',machineType='%v',os='%v',ip='%v'\n\n",
				v.Details.Label, v.Details.FQDN, v.Details.GroupID, v.Details.IsManaged, v.Details.MachineType, v.Details.OperatingSystemVersion, v.Details.IP,
			)
		}
	}

}

// This method creates a new scan task.
// Please note that the managed endpoints from virtualmachines service are also
// displayed in computers service under Custom Group To avoid launching duplicate
// scan tasks we recommend you to use the endpoints from the computers service.
func createScanTask(gz *gobdgz.GravityZoneAPI, rq gobdgz.Request) {
	r := rq
	r.Method = "createScanTask"

	rand.Seed(time.Now().UnixNano())
	taskRand := fmt.Sprintf("task-rnd-number-%v", rand.Intn(20000))

	r.Params = map[string]interface{}{
		"targetIds": []string{
			"5fc4b4403212d72d98240606",
			// "another ID",
			// "one another ID",
			// ...
		},
		"type": 1, // 1 = quick, 2 = full, 3 = memory, 4 = custom
		"name": taskRand,

		// when type is 4
		// "customScanSettings": map[string]interface{}{
		// 	"scanDepth": 1, // 1=Aggresive, 2=normal, 3=permissive
		// 	"scanPath": []string{
		// 		"LocalDrives",
		// 	},
		// },
	}

	gz.Network.SetRequest(r)
	resp, err := gz.Network.CreateScanTask()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("\n\tTask '%v' has created! The result is '%v'\n", taskRand, resp.Result)
}

// This method creates a new Reconfigure Client task. With this task you can choose
// which modules to install on target agents.
// * The networkMonitor module is deprecated. It is recommended to use networkAttackDefense instead.
func createReconfigureClientTask(gz *gobdgz.GravityZoneAPI, rq gobdgz.Request) {
	r := rq
	r.Method = "createReconfigureClientTask"

	// read from input
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter ObjectID: ")
	objectID, _ := reader.ReadString('\n')
	objectID = strings.Trim(objectID, " \n")

	r.Params = map[string]interface{}{
		"targetIds": []string{
			objectID,
		},
		"scheduler": map[string]interface{}{
			"type": 1, // 1=immediate, 2=scheduled
			// if 1 other fields not needed

			// "recurrence": 2, // 1 hourly = everyHour required
			// 2 daily  = startTime required
			// 3 weekly = everyHour & startTime required
			// "everyHour": 1, // An integer betweeb 1 to 23
			// "startTime": "13:23", // a string with the following formaat: HH:mm
			// onWeekDay: 3, // an Integer between 1 to 7 where 1 is Monday & 7 is Sunday
		},

		"modules": map[string]interface{}{
			// // 1 = enabled, 0 = disabled
			"antimalware":           1,
			"advancedThreatControl": 1,
			"firewall":              1,
			"contentControl":        1,
			"deviceControl":         1,
			"powerUser":             1,
			"applicationControl":    1,
			// "encryption":            1,
			"advancedAntiExploit":  1,
			"patchManagement":      1,
			"networkAttackDefense": 1,
		},
		// "scanMode": map[string]interface{}{ // scan engine mode
		// 	"type": 1, // 1 = automatic, 2 = custom (required vms and computers)

		// 	"computers": ma[string]interface{}{
		// 		"main": 1, // 1 = central, 2 = hybrid, 3 = local
		// 		"fallback": 2, // 2 = hybrid, 3 = local (main must be 1 to set fallback)
		// 	},

		// 	"vms": ma[string]interface{}{
		// 		"main": 1, // 1 = central, 2 = hybrid, 3 = local
		// 		"fallback": 2, // 2 = hybrid, 3 = local (main must be 1 to set fallback)
		// 	},
		// },

		// "roles": map[string]interface{}{
		// 	"relay": 1, // 1 = enabled, 2 = disabled
		// 	"exchange": 1, // 1 = enabled, 2 = disabled
		// },
	}

	gz.Network.SetRequest(r)
	resp, err := gz.Network.CreateReconfigureClientTask()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\nCreate Configure Client Task with result'%v' has created\n", resp.Result)

}
