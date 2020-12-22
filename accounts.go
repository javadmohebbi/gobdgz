package gobdgz

import "fmt"

type Accounts struct {
	request Request
}

type Role int

func (r Role) String() string {
	switch r {
	case 1:
		return "Company Administrator"
	case 2:
		return "Network Administrator"
	case 3:
		return "Security Analyst"
	case 5:
		return "Custom"
	default:
		return "Not-Documented"
	}
}

type GetAccountsListResponse struct {
	ID      *string               `json:"id"`
	JSONRPC string                `json:"jsonrpc"`
	Result  GetAccountsListResult `json:"result"`
}
type GetAccountsListResult struct {
	Total      int `json:"total"`
	Page       int `json:"page"`
	PerPage    int `json:"perPage"`
	PagesCount int `json:"pagesPage"`

	Items []GetAccountsListResponseItems `json:"items"`
}
type GetAccountsListResponseItems struct {
	ID    string `json:"id"`
	Email string `json:"email"`

	Profile GetAccountsListResponseItemsProfile `json:"profile"`

	// Role: 1 - Company Administrator, 2 - Network Administrator, 3 - Reporter, 5 - Custom
	Role Role `json:"role"`

	Rights GetAccountsListResponseItemsRights `json:"rights"`

	Username string `json:"userName"`
}

func (v GetAccountsListResponseItems) String() string {
	return fmt.Sprintf("Id: '%v', email: '%v', fullname: '%v', username: '%v', role: '%v'\n\t\t Rights: %v\n", v.ID, v.Email, v.Profile.FullName, v.Username, v.Role, v.Rights)
}

type GetAccountsListResponseItemsProfile struct {
	FullName string `json:"fullName"`
	TimeZone string `json:"timezone"`
	Language string `json:"language"`
}
type GetAccountsListResponseItemsRights struct {
	ManageCompanies bool `json:"manageCompanies"`
	ManageNetworks  bool `json:"manageNetworks"`
	ManageUsers     bool `json:"manageUsers"`
	ManageReports   bool `json:"manageReports"`
	ComapnyManager  bool `json:"companyManager"`
}

func (v GetAccountsListResponseItemsRights) String() string {
	return fmt.Sprintf("(rights) -> Manage Networks: %v, Manage Users: %v, , View & Analyze data: %v, Manage Solution: %v, , Manage Company: %v",
		v.ManageNetworks, v.ManageUsers, v.ManageReports, v.ComapnyManager, v.ManageCompanies,
	)
}

type CreateAccountResponse struct {
	ID      *string `json:"id"`
	JSONRPC string  `json:"jsonrpc"`
	Result  string  `json:"result"`
}

type UpdateAccountResponse struct {
	ID      *string `json:"id"`
	JSONRPC string  `json:"jsonrpc"`
	Result  bool    `json:"result"`
}

type GetNotificationsSettingsResponse struct {
	ID      *string                                `json:"id"`
	JSONRPC string                                 `json:"jsonrpc"`
	Result  GetNotificationsSettingsResponseResult `json:"result"`
}
type GetNotificationsSettingsResponseResult struct {
	DeleteAfter           int                                      `json:"deleteAfter"`
	IncludeDeviceName     bool                                     `json:"includeDeviceName"`
	IncludeDeviceFQDN     bool                                     `json:"includeDeviceFQDN"`
	EmailAddresses        []string                                 `json:"emailAddresses"`
	NotificationsSettings []GetNotificationsSettingsResponseDetail `json:"notificationsSettings"`
}
type NotificationType int

func (n NotificationType) String() string {
	str := ""
	switch n {
	case 1:
		str = "Malware Outbreak"
	case 2:
		str = "License Expires"
	case 3:
		str = "License Usage Limit Has Been Reached"
	case 4:
		str = "License Limit Is About To Be Reached"
	case 5:
		str = "Update Available"
	case 6:
		str = "Internet Connection"
	case 7:
		str = "SMTP Connection"
	case 8:
		str = "Database Backup"
	case 9:
		str = "Exchange License Usage Limit Has Been Reached"
	case 10:
		str = "Invalid Exchange User Credentials"
	case 11:
		str = "Upgrade Status"
	case 12:
		str = "Exchange Malware Detected"
	case 13:
		str = "Authentication Audit"
	case 14:
		str = "Certificate Expires"
	case 15:
		str = "GravityZone Update"
	case 16:
		str = "Antimalware Event"
	case 17:
		str = "Antipshising Event"
	case 18:
		str = "Firewall Event"
	case 19:
		str = "ATC/IDS Event"
	case 20:
		str = "User Control Event"
	case 21:
		str = "Data Protection Event"
	case 22:
		str = "Product Modules Event"
	case 23:
		str = "Security Server Status Event"
	case 24:
		str = "Product Registration Event"
	case 26:
		str = "Task Status"
	case 27:
		str = "Outdated Update Server"
	case 28:
		str = "New Application In Application Inventory"
	case 29:
		str = "Blocked Application"
	case 30:
		str = "Detected Memory Violation"
	case 31:
		str = "Mobile Device Users Without EmailAddress"
	case 38:
		str = "Blocked Devices"
	default:
		str = "Undocumented!"
	}

	return fmt.Sprintf("%v (%d)", str, n)
}

type GetNotificationsSettingsResponseDetail struct {
	Type                  NotificationType                                    `json:"type"`
	Enabled               bool                                                `json:"enabled"`
	VisibilitySettings    GetNotificationsSettingsResponseDetailVisibility    `json:"visibilitySettings"`
	ConfigurationSettings GetNotificationsSettingsResponseDetailConfiguration `json:"configurationSettings"`
}
type GetNotificationsSettingsResponseDetailVisibility struct {
	SendPerEmail               bool     `json:"sendPerEmail"`
	ShowInConsole              bool     `json:"showInConsole"`
	UseCustomEmailDistribution bool     `json:"UseCustomEmailDistribution"`
	Emails                     []string `json:"emails"`
	LogToServer                bool     `json:"logToServer"`
}
type GetNotificationsSettingsResponseDetailConfiguration struct {
	Threshold    int  `json:"threshold"`
	UseThreshold bool `json:"useThreshold"`
}

type ConfigureNotificationsSettingsResponse struct {
	ID      *string `json:"id"`
	JSONRPC string  `json:"jsonrpc"`
	Result  bool    `json:"result"`
}

// prepare request before sening the request
func (a *Accounts) SetRequest(r Request) {
	a.request = r
}

// get account list
func (a *Accounts) GetAccounList() (GetAccountsListResponse, error) {
	var resp GetAccountsListResponse
	err := a.request.SendRequest(&resp)
	return resp, err
}

// create account
func (a *Accounts) CreateAccount() (CreateAccountResponse, error) {
	var resp CreateAccountResponse
	err := a.request.SendRequest(&resp)
	return resp, err
}

// Update account
func (a *Accounts) UpdateAccount() (UpdateAccountResponse, error) {
	var resp UpdateAccountResponse
	err := a.request.SendRequest(&resp)
	return resp, err
}

// Get Notifications Settings
func (a *Accounts) GetNotificationsSettings() (GetNotificationsSettingsResponse, error) {
	var resp GetNotificationsSettingsResponse
	err := a.request.SendRequest(&resp)
	return resp, err
}

// Configure Notifications Settings
func (a *Accounts) ConfigureNotificationsSettings() (ConfigureNotificationsSettingsResponse, error) {
	var resp ConfigureNotificationsSettingsResponse
	err := a.request.SendRequest(&resp)
	return resp, err
}
