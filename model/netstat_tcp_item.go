package model

// NetstatTCPItem struct represent one TCP item from netstat output
type NetstatTCPItem struct {
	Protocol       string
	LocalAddress   string
	ForeignAddress string
	State          string
}
