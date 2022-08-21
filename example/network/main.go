package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/javadmohebbi/gobdgz"
)

func main() {
	// create new instance of GravityZone API
	gz := gobdgz.GravityZoneAPI{
		BaseURL: "https://192.168.59.113/api/v1.0/jsonrpc/",
	}

	// Create Default Request struct
	r := gobdgz.Request{
		// API KEY
		// must be enabled in GravityZone Control Center
		// http://download.bitdefender.com/business/API/Bitdefender_GravityZone_On-Premises_APIGuide_enUS.pdf#page=7&zoom=100,33,85
		APIKey: "64db6b1d6addb60d7d6fc09b538c3a0e2f3b1925f26e649f0a68b0c92d7af481",

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
	case "getScanTasksList":
		r.URL += "/" + *containerService
		getScanTasksList(&gz, r)
	case "getEndpointsList":
		r.URL += "/" + *containerService
		getEndpointsList(&gz, r)
	case "getManagedEndpointDetails":
		r.URL += "/" + *containerService
		getManagedEndpointDetails(&gz, r)
	case "createCustomGroup":
		r.URL += "/" + *containerService
		createCustomGroup(&gz, r)
	case "deleteCustomGroup":
		r.URL += "/" + *containerService
		deleteCustomGroup(&gz, r)
	case "moveCustomGroup":
		r.URL += "/" + *containerService
		moveCustomGroup(&gz, r)
	case "moveEndpoints":
		r.URL += "/" + *containerService
		moveEndpoints(&gz, r)
	case "deleteEndpoint":
		r.URL += "/" + *containerService
		deleteEndpoint(&gz, r)
	case "setEndpointLabel":
		r.URL += "/" + *containerService
		setEndpointLabel(&gz, r)
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
		// "parentId": "5fc4b4403212d72d98240606",

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

// This method creates a new Reconfigure Client task. With this task you can choose
// This method requires you to place the {service} name in the APIURL. The allowed
// services are: computers, for "Computers and Virtual Machines"
// and virtualmachines, for "Virtual Machines"
func getScanTasksList(gz *gobdgz.GravityZoneAPI, rq gobdgz.Request) {
	r := rq
	r.Method = "getScanTasksList"

	r.Params = map[string]interface{}{
		// Task name
		// * means all
		"name": "*",

		// 1 = Pending
		// 2 = In progress
		// 3 = Finished
		"status": 3,
		// "page": 1,
		// "perPage": 30
	}

	gz.Network.SetRequest(r)
	resp, err := gz.Network.GetScanTasksList()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("List of scan tasks")
	for _, v := range resp.Result.Items {
		fmt.Printf("\t ID:'%v', Name:'%v', StartDate:'%v', Status:'%v'  \n\n", v.ID, v.Name, v.StartDate, v.Status)
	}

}

// This method returns the list of the endpoints.
// To find the parentId, you must do several recursive calls to getContainers
// untilthe container with the endpoints is reached. The containerID from the response
// of getContainers should be used as parentId in this call. The same viewType
// used in getContainers should be used in this call.
// services are: computers, for "Computers and Virtual Machines"
// and virtualmachines, for "Virtual Machines"
func getEndpointsList(gz *gobdgz.GravityZoneAPI, rq gobdgz.Request) {
	r := rq
	r.Method = "getEndpointsList"

	r.Params = map[string]interface{}{

		// parent id container
		// "parentId": "",

		// isManaged or not
		"isManaged": false,

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

		//"filters": map[string]interface{}{} //http://download.bitdefender.com/business/API/Bitdefender_GravityZone_On-Premises_APIGuide_enUS.pdf#available%20filters
		// "filters": map[string]interface{}{
		// 	"depth": map[string]interface{}{
		// 		"allItemsRecursively": true,
		// 	},
		// 	"security": map[string]interface{}{
		// 		"management": map[string]interface{}{
		// 			"managedWithBest": true,
		// 			"managedRelays":   true,
		// 			"securityServers": true,
		// 		},
		// 	},
		// },

		// "page": 1,
		// "perPage": 30
	}

	gz.Network.SetRequest(r)
	resp, err := gz.Network.GetEndpointsList()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Endpoint Items:")
	for k, v := range resp.Result.Items {
		fmt.Printf("\tRow: %v -> %v='%v',type='%v',IP='%v'\n", k+1, v.ID, v.Name, v.MachineType, v.IP)
	}

}

// This method returns detailed information, such as: details to identify the endpoint
// and the security agent, the status of installed protection modules.
// services are: computers, for "Computers and Virtual Machines"
// and virtualmachines, for "Virtual Machines"
func getManagedEndpointDetails(gz *gobdgz.GravityZoneAPI, rq gobdgz.Request) {
	r := rq
	r.Method = "getManagedEndpointDetails"

	// read from input
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter endpointId: ")
	objectID, _ := reader.ReadString('\n')
	objectID = strings.Trim(objectID, " \n")

	r.Params = map[string]interface{}{
		"endpointId": objectID,
	}

	gz.Network.SetRequest(r)
	resp, err := gz.Network.GetManagedEndpointDetails()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Endpint '%v (%v)' under group '%v' is protected via '%v' policy\n\n", objectID, resp.Result.Name, resp.Result.Group.Name, resp.Result.Policy.Name)

}

// This method creates a new custom group.
// services are: computers, for "Computers and Virtual Machines"
// and virtualmachines, for "Virtual Machines"
func createCustomGroup(gz *gobdgz.GravityZoneAPI, rq gobdgz.Request) {
	r := rq
	r.Method = "createCustomGroup"

	// read from input
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter parentId: ")
	objectID, _ := reader.ReadString('\n')
	objectID = strings.Trim(objectID, " \n")

	// read from input
	fmt.Print("Enter group name: ")
	reader = bufio.NewReader(os.Stdin)
	groupName, _ := reader.ReadString('\n')
	groupName = strings.Trim(groupName, " \n")

	r.Params = map[string]interface{}{
		"parentId":  objectID,
		"groupName": groupName,
	}

	gz.Network.SetRequest(r)
	resp, err := gz.Network.CreateCustomGroup()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Group '%v' with '%v' ID has created \n\n", groupName, resp.Result)

}

// This method deletes a custom group.
// services are: computers, for "Computers and Virtual Machines"
// and virtualmachines, for "Virtual Machines"
func deleteCustomGroup(gz *gobdgz.GravityZoneAPI, rq gobdgz.Request) {
	r := rq
	r.Method = "deleteCustomGroup"

	// read from input
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter groupId: ")
	objectID, _ := reader.ReadString('\n')
	objectID = strings.Trim(objectID, " \n")

	r.Params = map[string]interface{}{
		"groupId": objectID,
		// "force": true,
	}

	gz.Network.SetRequest(r)
	resp, err := gz.Network.DeleteCustomGroup()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Group '%v' deletion request with result '%v' lanuched \n\n", objectID, resp.Result)

}

// This methodmovesa customgroupto anothercustomgroup
// services are: computers, for "Computers and Virtual Machines"
// and virtualmachines, for "Virtual Machines"
func moveCustomGroup(gz *gobdgz.GravityZoneAPI, rq gobdgz.Request) {
	r := rq
	r.Method = "moveCustomGroup"

	// read from input
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter The ID of the custom group to be moved: ")
	objectID, _ := reader.ReadString('\n')
	objectID = strings.Trim(objectID, " \n")

	// read from input
	reader = bufio.NewReader(os.Stdin)
	fmt.Print("Enter The ID of the destination custom group: ")
	destinationID, _ := reader.ReadString('\n')
	destinationID = strings.Trim(destinationID, " \n")

	r.Params = map[string]interface{}{

		// The ID of the custom group to be moved
		"groupId": objectID,

		// The ID of the destination custom group
		"parentId": destinationID,
	}

	gz.Network.SetRequest(r)
	resp, err := gz.Network.MoveCustomGroup()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Group with ID '%v' & with result '%v' has moved to Group ID '%v' \n\n", objectID, resp.Result, destinationID)

}

// This method moves a list of endpoints to a custom group.
// services are: computers, for "Computers and Virtual Machines"
// and virtualmachines, for "Virtual Machines"
func moveEndpoints(gz *gobdgz.GravityZoneAPI, rq gobdgz.Request) {
	r := rq
	r.Method = "moveEndpoints"

	// read from input
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter The ID of the destination custom group: ")
	destinationID, _ := reader.ReadString('\n')
	destinationID = strings.Trim(destinationID, " \n")

	// read from input
	reader = bufio.NewReader(os.Stdin)
	fmt.Print("How many endpoints you want to move to: ")
	countStr, _ := reader.ReadString('\n')
	countStr = strings.Trim(countStr, " \n")

	// convert to Int
	count, err := strconv.Atoi(countStr)

	// if err, returns error
	if err != nil {
		log.Fatal(err)
	}

	if count <= 0 {
		log.Fatal("Endpoints count can not be ZERO. Exiting....")
	}

	// endpointIDs
	var objectIDs []string

	for i := 0; i < count; i++ {
		// read from input
		reader = bufio.NewReader(os.Stdin)
		fmt.Printf("Enter The ID of the Endpoint to be moved (%v/%v): ", i+1, count)
		objectID, _ := reader.ReadString('\n')
		objectID = strings.Trim(objectID, " \n")

		// append endpoint ID to list
		objectIDs = append(objectIDs, objectID)
	}

	r.Params = map[string]interface{}{

		// The list of endpointsIDs to be moved
		"endpointIds": objectIDs,

		// The ID of the destination custom group
		"groupId": destinationID,
	}

	gz.Network.SetRequest(r)
	resp, err := gz.Network.MoveEndpoints()
	if err != nil {
		log.Fatal(err)
	}

	for k, objID := range objectIDs {
		fmt.Printf("%v -> Endpoint with ID '%v' & with result '%v' has moved to Group ID '%v' \n", k+1, objID, resp.Result, destinationID)
	}
	fmt.Printf("\n")

}

// This method deletes an endpoint.
// services are: computers, for "Computers and Virtual Machines"
// and virtualmachines, for "Virtual Machines"
// ** Deleting an endpoint under CustomGroups moves it to the Deleted group.
// ** For managed endpoints,
// ** an Uninstall task is automatically generated.
// ** To permanently remove an endpoint, call the method twice using the sameID.
func deleteEndpoint(gz *gobdgz.GravityZoneAPI, rq gobdgz.Request) {
	r := rq
	r.Method = "deleteEndpoint"

	// read from input
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter The ID of the Endpoint to be removed: ")
	objectID, _ := reader.ReadString('\n')
	objectID = strings.Trim(objectID, " \n")

	r.Params = map[string]interface{}{
		// The ID of the endpoint
		"endpointId": objectID,
	}

	gz.Network.SetRequest(r)
	resp, err := gz.Network.DeleteEndpoint()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Endpoint with ID '%v' & with result '%v' has removed\n\n", objectID, resp.Result)

}

// This method sets a new label to an endpoint.
// This method returns a Boolean which is True, when the label was successfully set
// ** A string representing the label. The maximum
// ** allowed length is 64 characters. Enter an empty
// ** string to reset a previously set label
func setEndpointLabel(gz *gobdgz.GravityZoneAPI, rq gobdgz.Request) {
	r := rq
	r.Method = "setEndpointLabel"

	// read from input
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter The ID of the Endpoint you want to change it's label: ")
	objectID, _ := reader.ReadString('\n')
	objectID = strings.Trim(objectID, " \n")

	// read from input
	reader = bufio.NewReader(os.Stdin)
	fmt.Printf("Enter Endpoint '%v' label: ", objectID)
	label, _ := reader.ReadString('\n')
	label = strings.Trim(label, " \n")

	r.Params = map[string]interface{}{
		// The ID of the endpoint
		"endpointId": objectID,

		// The ID of the endpoint
		"label": label,
	}

	gz.Network.SetRequest(r)
	resp, err := gz.Network.SetEndpointLabel()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Endpoint's label with ID '%v' & with result '%v' has changed to '%v'\n\n", objectID, resp.Result, label)

}
