package gobdgz

type Policy struct {
	request Request
}

// prepare request before sening the request
func (p *Policy) SetRequest(r Request) {
	p.request = r
}

type ExportPolicyResponse struct {
	ID      *string            `json:"id"`
	JSONRPC string             `json:"jsonrpc"`
	Result  ExportPolicyResult `json:"result"`
}
type ExportPolicyResult struct {
	Total         int           `json:"total"`
	Page          int           `json:"page"`
	PerPage       int           `json:"perPage"`
	PagesCount    int           `json:"pagesPage"`
	Items         []PolicyItems `json:"items"`
	CompatVersion string        `json:"compatVersion"`
}

type ImportPolicyResponse struct {
	ID      *string            `json:"id"`
	JSONRPC string             `json:"jsonrpc"`
	Result  ImportPolicyResult `json:"result"`
}

type ImportPolicyResult struct {
	Success bool `json:"success"`
}

// ExportPolicies
func (p *Policy) ExportPolicies() (ExportPolicyResponse, error) {
	var resp ExportPolicyResponse
	err := p.request.SendRequest(&resp)
	return resp, err
}

// ImportPolicies
func (p *Policy) ImportPolicies() (ImportPolicyResponse, error) {
	var resp ImportPolicyResponse
	err := p.request.SendRequest(&resp)
	return resp, err
}
