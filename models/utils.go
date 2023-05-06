package models

type WebsocketDataDecode struct {
	ClientType  int    `json:"clienttype"`
	ProtocolNum int    `json:"protocolnum"`
	Bytes       string `json:"bytes"`
}
