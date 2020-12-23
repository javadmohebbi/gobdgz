package main

import (
	"flag"
	"fmt"
	"log"

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
		URL: gz.BaseURL + "/network",

		// HTTP MEthods
		HttpMethod: "POST",
	}

	// method kind command line argument
	kind := flag.String("method", "getContainers", "methods to call: possible values are: getContainers, getNetworkInventoryItems, createScanTask, createReconfigureClientTask, getScanTasksList, getEndpointsList, getManagedEndpointDetails, createCustomGroup, deleteCustomGroup, moveCustomGroup, moveEndpoints, deleteEndpoint, setEndpointLabel, createScanTaskByMac, assignPolicy")
	containerService := flag.String("service", "computers", "service to call: possible values are: computers, virtualmachines, mobile (NOT AVAILABLE in getNetworkInventoryItems)")

	// parsing flags
	flag.Parse()

	switch *kind {
	case "getContainers":
		r.URL += "/" + *containerService
		getContainers(&gz, r)
	case "getNetworkInventoryItems":
		r.URL += "/" + *containerService
		getNetworkInventoryItems(&gz, r)
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
