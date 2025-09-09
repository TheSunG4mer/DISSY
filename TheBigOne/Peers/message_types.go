package peers

import (
	"encoding/json"
	"net"
)

type MessageType int

const (
	MessageString      MessageType = iota // Contains message of type String
	MessageTransaction                    // Contains message of type Transaction
	MessagePeerInfo                       // Contains message of type PeerInfo
)

type MessageContent interface {
}

type Message struct {
	Type    MessageType
	Content json.RawMessage
}

func tryMarshal(v interface{}) json.RawMessage {
	b, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	return b
}

func MarshalStringToMessage(s string) Message {
	msg := Message{
		Type:    MessageString,
		Content: tryMarshal(s),
	}
	return msg
}

func MarshalTransactionToMessage(t *Transaction) Message {
	msg := Message{
		Type:    MessageTransaction,
		Content: tryMarshal(t),
	}
	return msg
}

func DemarshalToString(rawMessage json.RawMessage) string {
	var s string
	json.Unmarshal(rawMessage, &s)
	return s
}

func DemarshalToTransaction(rawMessage json.RawMessage) *Transaction {
	var t Transaction
	json.Unmarshal(rawMessage, &t)
	return &t
}

func DemarshalToPeerInfo(rawMessage json.RawMessage) *PeerInfo {
	var pi PeerInfo
	json.Unmarshal(rawMessage, &pi)
	return &pi
}

func sendMessage(conn net.Conn, msg Message) error {
	b, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	_, err = conn.Write(b)
	return err
}
