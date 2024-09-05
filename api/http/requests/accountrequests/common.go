package accountrequests

type Create struct {
	Token      string `json:"token" validate:"required"`
	ProxyType  string `json:"proxy_type" validate:""`
	ProxyLogin string `json:"proxy_login" validate:""`
	ProxyPass  string `json:"proxy_pass" validate:""`
	ProxyIP    string `json:"proxy_ip" validate:""`
	ProxyPort  string `json:"proxy_port" validate:""`
}
