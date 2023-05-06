package protocols

import (
	"bytes"
	"encoding/binary"
	"errors"
)

type S_Chat_ToClient struct {
	Data T_Chat_Data
}

func (p *S_Chat_ToClient) Encode() []byte {
	buffer := new(bytes.Buffer)
	buffer.Write(p.Data.Encode())
	return buffer.Bytes()
}

func (p *S_Chat_ToClient) Decode(buffer *bytes.Buffer) error {
	return p.Data.Decode(buffer)
}

type S_Chat_CloseFightRoom struct {
	RoomID string
}

func (p *S_Chat_CloseFightRoom) Encode() []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.LittleEndian, uint32(len(p.RoomID)))
	buffer.Write([]byte(p.RoomID))
	return buffer.Bytes()
}

func (p *S_Chat_CloseFightRoom) Decode(buffer *bytes.Buffer) error {
	if buffer.Len() < 4 {
		return errors.New("message length error")
	}
	var RoomIDLen uint32
	binary.Read(buffer, binary.LittleEndian, &RoomIDLen)
	if uint32(buffer.Len()) < RoomIDLen {
		return errors.New("message length error")
	}
	p.RoomID = string(buffer.Next(int(RoomIDLen)))
	return nil
}
