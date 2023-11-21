package dto

type EndpointInfoDTO struct {
	IPAddress         string `json:"ipAddress"`
	ServerName        string `json:"serverName"`
	StatusMessage     string `json:"statusMessage"`
	Grade             string `json:"grade"`
	GradeTrustIgnored string `json:"gradeTrustIgnored"`
	HasWarnings       bool   `json:"hasWarnings"`
	IsExceptional     bool   `json:"isExceptional"`
	Progress          int    `json:"progress"`
	Duration          int    `json:"duration"`
	Delegation        int    `json:"delegation"`
}

type SSLInfoResultDTO struct {
	Host            string            `json:"host"`
	Port            int               `json:"port"`
	Protocol        string            `json:"protocol"`
	IsPublic        bool              `json:"isPublic"`
	Status          string            `json:"status"`
	StartTime       int               `json:"startTime"`
	TestTime        int               `json:"testTime"`
	EngineVersion   string            `json:"engineVersion"`
	CriteriaVersion string            `json:"criteriaVersion"`
	Endpoints       []EndpointInfoDTO `json:"endpoints"`
}
