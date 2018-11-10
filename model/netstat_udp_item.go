package model

// NetstatUDPItem struct represent one UDP item from netstat output
type NetstatUDPItem struct {
	Protocol       string
	LocalAddress   string
	ForeignAddress string
}
