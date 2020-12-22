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
		return "Reporter"
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
