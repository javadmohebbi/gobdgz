package gobdgz

import (
	"fmt"
)

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

type GetManagedEndpointDetailsResponse struct {
	ID      *string                                 `json:"id"`
	JSONRPC string                                  `json:"jsonrpc"`
	Result  GetManagedEndpointDetailsResponseResult `json:"result"`
}
type ManagedEndpointStatus int

func (v ManagedEndpointStatus) String() string {
	str := "!!undocumented!!"
	switch v {
	case 1:
		str = "Online"
	case 2:
		str = "Offline"
	case 3:
		str = "Suspended"
	case 0:
		str = "Unknown"
	}
	return fmt.Sprintf("(%d) %v", v, str)
}

type GetManagedEndpointDetailsResponseResult struct {
	ID                            string                                       `json:"id"`
	Name                          string                                       `json:"name"`
	CompanyID                     string                                       `json:"companyId"`
	OperatingSystem               string                                       `json:"operatingSystem"`
	State                         ManagedEndpointStatus                        `json:"state"`
	IP                            string                                       `json:"ip"`
	MachineType                   InventoryObjectMachineType                   `json:"machineType"`
	Agent                         GetManagedEndpointDetailsResponseResultAgent `json:"agent"`
	Group                         AgentGroup                                   `json:"group"`
	MalwareStatus                 AgentMalwareStatus                           `json:"malwareStatus"`
	Policy                        AgentPolicy                                  `json:"policy"`
	HypervisorMemoryIntrospection AgentHypervisorMemoryIntrospection           `json:"hypervisorMemoryIntrospection"`
	Modules                       AgentModules                                 `json:"modules"`
	Label                         string                                       `json:"label"`
	ManagedWithBest               bool                                         `json:"managedWithBest"`
	ManagedExchangeServer         bool                                         `json:"managedExchangeServer"`
	ManagedRelay                  bool                                         `json:"managedRelay"`
	SecurityServer                bool                                         `json:"securityServer"`
	ManagedWithNsx                bool                                         `json:"managedWithNsx"`
	ManagedWithVShield            bool                                         `json:"managedWithVShield"`
	ManagedWithHvi                bool                                         `json:"managedWithHvi"`
	HVIProtectionType             HVIProtectionType                            `json:"hviProtectionType"`
}
type EngineType int

func (v EngineType) String() string {
	str := "!!undocumented!!"
	switch v {
	case 1:
		str = "Central Scanning (Security Server)"
	case 2:
		str = "Hybrid Scanning (Light Engines)"
	case 3:
		str = "Local Scanning (Full Engines)"
	case 0:
		str = "Unknown"
	}
	return fmt.Sprintf("(%d) %v", v, str)
}

type AgentLicense int

func (v AgentLicense) String() string {
	str := "!!undocumented!!"
	switch v {
	case 0:
		str = "Pending authentication"
	case 1:
		str = "Active license"
	case 2:
		str = "Expired license"
	case 6:
		str = "No license or not applicable"
	}
	return fmt.Sprintf("(%d) %v", v, str)
}

type TypeOfAgent int

func (v TypeOfAgent) String() string {
	str := "!!undocumented!!"
	switch v {
	case 1:
		str = "Endpoint Security"
	case 2:
		str = "Bitdefender Tools"
	case 3:
		str = "BEST"
	}
	return fmt.Sprintf("(%d) %v", v, str)
}

type GetManagedEndpointDetailsResponseResultAgent struct {
	EngineVersion           string       `json:"engineVersion"`
	PrimaryEngine           EngineType   `json:"primaryEngine"`
	FallbackEngine          EngineType   `json:"fallbackEngine"`
	LastUpdate              string       `json:"lastUpdate"`
	Licensed                AgentLicense `json:"licensed"`
	ProductOutdated         bool         `json:"productOutdated"`
	ProductUpdateDisabled   bool         `json:"productUpdateDisabled"`
	ProductVersion          string       `json:"productVersion"`
	SignatureOutdated       bool         `json:"signatureOutdated"`
	SignatureUpdateDisabled bool         `json:"signatureUpdateDisabled"`
	Type                    TypeOfAgent  `json:"type"`
}
type AgentGroup struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
type AgentMalwareStatus struct {
	Detection bool `json:"detection"`
	Infected  bool `json:"infected"`
}
type AgentPolicy struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Applied bool   `json:"applied"`
}
type AgentHypervisorMemoryIntrospection struct {
	Status         bool                                             `json:"status"`
	ActiveModules  AgentHypervisorMemoryIntrospectionActiveModules  `json:"activeModules"`
	SecurityServer AgentHypervisorMemoryIntrospectionSecurityServer `json:"securityServer"`
	IsLicensed     bool                                             `json:"isLicensed"`
}
type AgentHypervisorMemoryIntrospectionActiveModules struct {
	UserMode   bool `json:"userMode"`
	KernelMode bool `json:"kernelMode"`
}
type AgentHypervisorMemoryIntrospectionSecurityServer struct {
	Name  string `json:"name"`
	IP    string `json:"ip"`
	Label string `json:"label"`
}
type AgentModules struct {
	AdvancedThreatControl bool `json:"advancedThreatControl"`
	Antimalware           bool `json:"antimalware"`
	ContentControl        bool `json:"contentControl"`
	DeviceControl         bool `json:"deviceControl"`
	Firewall              bool `json:"firewall"`
	PowerUser             bool `json:"powerUser"`
	Encryption            bool `json:"encryption"`
	HyperDetect           bool `json:"hyperDetect"`
	PatchManagement       bool `json:"patchManagement"`
	Relay                 bool `json:"relay"`
	Exchange              bool `json:"exchange"`
	SandboxAnalyzer       bool `json:"sandboxAnalyzer"`
	AdvancedAntiExploit   bool `json:"advancedAntiExploit"`
	NetworkAttackDefense  bool `json:"networkAttackDefense"`
}

type CreateCustomGroupResponse struct {
	ID      *string `json:"id"`
	JSONRPC string  `json:"jsonrpc"`
	Result  string  `json:"result"`
}

type DeleteCustomGroupResponse struct {
	ID      *string `json:"id"`
	JSONRPC string  `json:"jsonrpc"`
	Result  *string `json:"result"`
}

type MoveCustomGroupResponse struct {
	ID      *string `json:"id"`
	JSONRPC string  `json:"jsonrpc"`
	Result  *string `json:"result"`
}

type MoveEndpointGroupResponse struct {
	ID      *string `json:"id"`
	JSONRPC string  `json:"jsonrpc"`
	Result  *string `json:"result"`
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

// This method returns detailed information, such as: details to identify the endpoint
// and the security agent, the status of installed protection modules.
// services are: computers, for "Computers and Virtual Machines"
// and virtualmachines, for "Virtual Machines"
func (n *Network) GetManagedEndpointDetails() (GetManagedEndpointDetailsResponse, error) {
	var resp GetManagedEndpointDetailsResponse
	err := n.request.SendRequest(&resp)
	return resp, err
}

// This method creates a new custom group.
// services are: computers, for "Computers and Virtual Machines"
// and virtualmachines, for "Virtual Machines"
func (n *Network) CreateCustomGroup() (CreateCustomGroupResponse, error) {
	var resp CreateCustomGroupResponse
	err := n.request.SendRequest(&resp)
	return resp, err
}

// This method deletes a custom group.
// services are: computers, for "Computers and Virtual Machines"
// and virtualmachines, for "Virtual Machines"
func (n *Network) DeleteCustomGroup() (DeleteCustomGroupResponse, error) {
	var resp DeleteCustomGroupResponse
	err := n.request.SendRequest(&resp)
	return resp, err
}

// This method deletes a custom group.
// services are: computers, for "Computers and Virtual Machines"
// and virtualmachines, for "Virtual Machines"
func (n *Network) MoveCustomGroup() (MoveCustomGroupResponse, error) {
	var resp MoveCustomGroupResponse
	err := n.request.SendRequest(&resp)
	return resp, err
}

// This method moves a list of endpoints to a custom group.
// services are: computers, for "Computers and Virtual Machines"
// and virtualmachines, for "Virtual Machines"
func (n *Network) MoveEndpoints() (MoveEndpointGroupResponse, error) {
	var resp MoveEndpointGroupResponse
	err := n.request.SendRequest(&resp)
	return resp, err
}
