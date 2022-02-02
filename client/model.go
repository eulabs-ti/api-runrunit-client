package client

import gotoolboxtime "github.com/eucatur/go-toolbox/time"

const (
	// GatewayType tipo do gateway
	GatewayType = "api-runrunit"

	headerAppKey    = "App-Key"
	headerUserToken = "User-Token"
)

// Client ...
type Client struct {
	Host      string
	AppKey    string
	UserToken string
}

type OffDay struct {
	ID          int64                 `json:"id"`
	Day         gotoolboxtime.TimeEUA `json:"day"`
	Description string                `json:"description"`
}
