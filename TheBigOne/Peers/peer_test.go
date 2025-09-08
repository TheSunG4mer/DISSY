package peers_test

import (
	peers "P2P_block_chain/Peers"
	"testing"
)

func TestPeer_GetName(t *testing.T) {
	tests := []struct {
		name string
		id   string
		ip   string
		port int
		want string
	}{
		{name: "John_is_John", id: "John", ip: "127.0.0.1", port: 1234, want: "John"},
		{name: "Steve_is_Steve", id: "Steve", ip: "127.0.0.1", port: 1234, want: "Steve"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := peers.MakePeer(tt.id, tt.ip, tt.port)
			got := p.GetName()
			if got != tt.want {
				t.Errorf("GetName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPeer_GetMessageCounter(t *testing.T) {
	tests := []struct {
		name string
		id   string
		ip   string
		port int
		want int
	}{
		{name: "New account has message count 0", id: "John", ip: "127.0.0.1", port: 1234, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := peers.MakePeer(tt.id, tt.ip, tt.port)
			got := p.GetMessageCounter()
			if got != tt.want {
				t.Errorf("GetMessageCounter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPeer_BumpMessageCounter(t *testing.T) {
	p := peers.MakePeer("Andrew", "127.0.0.1", 1234)
	mc := p.BumpMessageCounter()
	if mc != 0 {
		t.Errorf("BumpMessageCounter() = %v, want %v", mc, 0)
	}
	mc = p.BumpMessageCounter()
	if mc != 1 {
		t.Errorf("BumpMessageCounter() = %v, want %v", mc, 1)
	}
	mc = p.BumpMessageCounter()
	if mc != 2 {
		t.Errorf("BumpMessageCounter() = %v, want %v", mc, 2)
	}
}
