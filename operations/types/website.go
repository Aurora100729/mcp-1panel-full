package types

import "time"

type ListWebsiteRes struct {
	Response
	Data struct {
		PageResult
		Items []WebsiteRes `json:"items"`
	} `json:"data"`
}

type ListWebsiteRequest struct {
	PageRequest
	Order   string `json:"order"`
	OrderBy string `json:"orderBy"`
}

type CreateWebsiteRequest struct {
	Domains []WebsiteDomain `json:"domains"`
	Alias          string `json:"alias"`
	Type           string `json:"type"`
	WebsiteGroupID uint   `json:"websiteGroupId"`
	Proxy          string `json:"proxy"`
	AppType        string `json:"appType"`
}

type WebsiteDomain struct {
	Domain string `json:"domain" validate:"required"`
	Port   int    `json:"port"`
	SSL    bool   `json:"ssl"`
}

type GroupRequest struct {
	Type string `json:"type"`
}

type GroupRes struct {
	Response
	Data []Group `json:"data"`
}

type Group struct {
	ID        uint `json:"id"`
	IsDefault bool `json:"isDefault"`
}

type WebsiteRes struct {
	ID            uint      `json:"id"`
	CreatedAt     time.Time `json:"createdAt"`
	Protocol      string    `json:"protocol"`
	PrimaryDomain string    `json:"primaryDomain"`
	Type          string    `json:"type"`
	Alias         string    `json:"alias"`
	Remark        string    `json:"remark"`
	Status        string    `json:"status"`
	ExpireDate    time.Time `json:"expireDate"`
	AppName       string    `json:"appName"`
	RuntimeName   string    `json:"runtimeName"`
	SSLExpireDate time.Time `json:"sslExpireDate"`
}
