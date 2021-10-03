package v1

type ConnectionCheckRequest struct {
	// Addresses to check the connection
	//
	// required: true
	// example: ["http://google.com"]
	Addresses []string `json:"addresses"`

	// timeout
	//
	// required: true
	// example: 3
	Timeout int `json:"timeout"`
}

type ConnectionCheckResponse struct {
	// Status
	//
	// required: true
	// example: up
	Status string `json:"status"`
}
