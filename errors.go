package gobdgz

const (
	ErrorParse          = -32700 // Parse error
	ErrorInvalidRequest = -32600 // Invalid request
	ErrorMethodNotFound = -32601 // Method not found
	ErrorInvalidParams  = -32602 // Invalid params
	ErrorServerError    = -32000 // Server error
)

const (
	ErrorParseText          = "Parse error"
	ErrorInvalidRequestText = "Invalid request"
	ErrorMethodNotFoundText = "Method not found"
	ErrorInvalidParamsText  = "Invalid params"
	ErrorServerErrorText    = "Server error"
)

const (
	ErrorHTTP401 = "Unauthorized access"
	ErrorHTTP403 = "Resource forbidden"
	ErrorHTTP405 = "Method not allowed"
	ErrorHTTP429 = "Too many requests"
)

type ResponseError struct {
	ID      *string `json:"id"`
	JSONRPC string  `json:"jsonrpc"`
	Error   Error   `json:"error"`
}
type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    Data   `json:"data"`
}
type Data struct {
	Details string `json:"details"`
}
