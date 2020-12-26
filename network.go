package gobdgz

import "fmt"

type Network struct {
	request Request
}

type GetContainersResponse struct {
	ID      *string               `json:"id"`
	JSONRPC string                `json:"jsonrpc"`
	Result  []GetContainersResult `json:"result"`
}
type GetContainersResult struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type GetNetworkInventoryItemsResponse struct {
	ID      *string                                `json:"id"`
	JSONRPC string                                 `json:"jsonrpc"`
	Result  GetNetworkInventoryItemsResponseResult `json:"result"`
}
type GetNetworkInventoryItemsResponseResult struct {
	Total      int `json:"total"`
	Page       int `json:"page"`
	PerPage    int `json:"perPage"`
	PagesCount int `json:"pagesPage"`

	Items []GetNetworkInventoryItemsResponseResultItems `json:"items"`
}
type InventoryObjectType int

func (v InventoryObjectType) String() string {
	str := "!!undocumented!!"
	switch v {
	case 4:
		str = "Group"
	case 5:
		str = "Computer"
	case 6:
		str = "Virtual Machine"
	case 8:
		str = "Virtual Host"
	case 9:
		str = "vShield App"
	case 10:
		str = "Virtualization Cluster"
	case 11:
		str = "Virtualization Datacenter"
	case 12:
		str = "Resource Pool"
	case 13:
		str = "Virtualization Pool"
	}

	return fmt.Sprintf("(%d) %v", v, str)
}

type GetNetworkInventoryItemsResponseResultItems struct {
	ID       string              `json:"id"`
	Name     string              `json:"name"`
	Type     InventoryObjectType `json:"type"`
	ParentID string              `json:"parentId"`

	// details available only for type 5 , 6
	Details GetNetworkInventoryItemsResponseResultItemsDetails `json:"details"`
}
type InventoryObjectMachineType int

func (v InventoryObjectMachineType) String() string {
	str := "!!undocumented!!"
	switch v {
	case 1:
		str = "computer"
	case 2:
		str = "virtual machine"
	case 0:
		str = "other"
	}
	return fmt.Sprintf("(%d) %v", v, str)
}

type GetNetworkInventoryItemsResponseResultItemsDetails struct {
	Label                  string                     `json:"label"`
	FQDN                   string                     `json:"fqdn"`
	GroupID                string                     `json:"groupId"`
	IsManaged              bool                       `json:"isManaged"`
	MachineType            InventoryObjectMachineType `json:"machineType"`
	OperatingSystemVersion string                     `json:"operatingSystemVersion"`
	IP                     string                     `json:"ip"`
	MACs                   []string                   `json:"macs"`
}

type CreateScanTaskResponse struct {
	ID      *string `json:"id"`
	JSONRPC string  `json:"jsonrpc"`
	Result  bool    `json:"result"`
}

type CreateReconfigureClientTaskResponse struct {
	ID      *string `json:"id"`
	JSONRPC string  `json:"jsonrpc"`
	Result  bool    `json:"result"`
}

type GetScanTasksListResponse struct {
	ID      *string                        `json:"id"`
	JSONRPC string                         `json:"jsonrpc"`
	Result  GetScanTasksListResponseResult `json:"result"`
}
type GetScanTasksListResponseResult struct {
	Total      int `json:"total"`
	Page       int `json:"page"`
	PerPage    int `json:"perPage"`
	PagesCount int `json:"pagesPage"`

	Items []GetScanTasksListResponseResultItems `json:"items"`
}
type ScanTaskStatus int

func (v ScanTaskStatus) String() string {
	str := "!!undocumented!!"
	switch v {
	case 1:
		str = "Pending"
	case 2:
		str = "In Progress"
	case 3:
		str = "Finished"
	}
	return fmt.Sprintf("(%d) %v", v, str)
}

type GetScanTasksListResponseResultItems struct {
	ID        string         `json:"id"`
	Name      string         `json:"name"`
	Status    ScanTaskStatus `json:"status"`
	StartDate string         `json:"startDate"`
}

type GetEndpointsListResponse struct {
	ID      *string                        `json:"id"`
	JSONRPC string                         `json:"jsonrpc"`
	Result  GetEndpointsListResponseResult `json:"result"`
}

type GetEndpointsListResponseResult struct {
	Total      int `json:"total"`
	Page       int `json:"page"`
	PerPage    int `json:"perPage"`
	PagesCount int `json:"pagesPage"`

	Items []GetEndpointsListResponseResultItems `json:"items"`
}

type HVIProtectionType int

func (v HVIProtectionType) String() string {
	str := "!!undocumented!!"
	switch v {
	case 1:
		str = "Security Server Multi-platform"
	case 2:
		str = "BEST"
	}
	return fmt.Sprintf("(%d) %v", v, str)
}

type GetEndpointsListResponseResultItems struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Label string `json:"label"`
	FQDN  string `json:"fqdn"`

	GroupID     string                     `json:"groupId"`
	IsManaged   bool                       `json:"isManaged"`
	MachineType InventoryObjectMachineType `json:"machineType"`

	OperatingSystemVersion string   `json:"operatingSystemVersion"`
	IP                     string   `json:"ip"`
	MACs                   []string `json:"macs"`
	SSID                   string   `json:"ssid"`

	ManagedWithHvi        bool              `json:"managedWithHvi"`
	HVIProtectionType     HVIProtectionType `json:"hviProtectionType"`
	ManagedWithBest       bool              `json:"managedWithBest"`
	ManagedExchangeServer bool              `json:"managedExchangeServer"`
	ManagedRelay          bool              `json:"managedRelay"`
	SecurityServer        bool              `json:"securityServer"`
	ManagedWithNsx        bool              `json:"managedWithNsx"`
	ManagedWithVShield    bool              `json:"managedWithVShield"`
}

// prepare request before sening the request
func (n *Network) SetRequest(r Request) {
	n.request = r
}

// This method returns network containers. It will return an empty list if the parentId
// is not a container or does not contain any other container within it.
func (n *Network) GetContainers() (GetContainersResponse, error) {
	var resp GetContainersResponse
	err := n.request.SendRequest(&resp)
	return resp, err
}

// This method returns network inventory items.
func (n *Network) GetNetworkInventoryItems() (GetNetworkInventoryItemsResponse, error) {
	var resp GetNetworkInventoryItemsResponse
	err := n.request.SendRequest(&resp)
	return resp, err
}

// This method creates a new scan task.
// Please note that the managed endpoints from virtualmachines service are also
// displayed in computers service under Custom Group To avoid launching duplicate
// scan tasks we recommend you to use the endpoints from the computers service.
func (n *Network) CreateScanTask() (CreateScanTaskResponse, error) {
	var resp CreateScanTaskResponse
	err := n.request.SendRequest(&resp)
	return resp, err
}

// This method creates a new Reconfigure Client task. With this task you can choose
// which modules to install on target agents.
// * The networkMonitor module is deprecated. It is recommended to use networkAttackDefense instead.
func (n *Network) CreateReconfigureClientTask() (CreateReconfigureClientTaskResponse, error) {
	var resp CreateReconfigureClientTaskResponse
	err := n.request.SendRequest(&resp)
	return resp, err
}

// This method returns the list of scan tasks.
// This method requires you to place the {service} name in the APIURL. The allowed
// services are:
// 		● computers, for "Computers and Virtual Machines"
// 		● virtualmachines, for "Virtual Machines"
func (n *Network) GetScanTasksList() (GetScanTasksListResponse, error) {
	var resp GetScanTasksListResponse
	err := n.request.SendRequest(&resp)
	return resp, err
}

// This method returns the list of the endpoints.
// To find the parentId, you must do several recursive calls to getContainers
// untilthe container with the endpoints is reached. The containerID from the response
// of getContainers should be used as parentId in this call. The same viewType
// used in getContainers should be used in this call.
// services are: computers, for "Computers and Virtual Machines"
// and virtualmachines, for "Virtual Machines"
func (n *Network) GetEndpointsList() (GetEndpointsListResponse, error) {
	var resp GetEndpointsListResponse
	err := n.request.SendRequest(&resp)
	return resp, err
}
