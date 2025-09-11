package peers

import (
	"strconv"
)

type PeerInfo struct {
	ID   string
	IP   string
	Port int
}

func (pi *PeerInfo) GenerateAddressString() string {
	return pi.IP + ":" + strconv.Itoa(pi.Port)
}
