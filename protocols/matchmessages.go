package protocols

import (
	"bytes"
	"encoding/binary"
	"errors"
)

/************************************  客户端  *********************************/

type C_Match_Fight struct {
	FightType int32
	Params    []int32
}

func (p *C_Match_Fight) Encode() []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.LittleEndian, p.FightType)
	binary.Write(buffer, binary.LittleEndian, uint32(len(p.Params)))
	for _, v := range p.Params {
		binary.Write(buffer, binary.LittleEndian, v)
	}
	return buffer.Bytes()
}

func (p *C_Match_Fight) Decode(buffer *bytes.Buffer, key uint8) error {
	if key != 0 {
		for i := 0; i < buffer.Len(); i++ {
			buffer.Bytes()[i] ^= byte(key)
		}
	}
	binary.Read(buffer, binary.LittleEndian, &p.FightType)
	var ParamsLen uint32
	binary.Read(buffer, binary.LittleEndian, &ParamsLen)
	if uint32(buffer.Len()) < ParamsLen*4 {
		return errors.New("message length error")
	}
	p.Params = make([]int32, ParamsLen)
	for i := uint32(0); i < ParamsLen; i++ {
		binary.Read(buffer, binary.LittleEndian, &p.Params[i])
	}
	return nil
}

type C_Match_Duel_Fight struct {
	FightType  int32
	UseRoom    int32
	RoomID     string
	LongRoomID string
	Params     []int32
}

func (p *C_Match_Duel_Fight) Encode() []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.LittleEndian, p.FightType)
	binary.Write(buffer, binary.LittleEndian, p.UseRoom)
	binary.Write(buffer, binary.LittleEndian, uint32(len(p.RoomID)))
	buffer.Write([]byte(p.RoomID))
	binary.Write(buffer, binary.LittleEndian, uint32(len(p.LongRoomID)))
	buffer.Write([]byte(p.LongRoomID))
	binary.Write(buffer, binary.LittleEndian, uint32(len(p.Params)))
	for _, v := range p.Params {
		binary.Write(buffer, binary.LittleEndian, v)
	}
	return buffer.Bytes()
}

func (p *C_Match_Duel_Fight) Decode(buffer *bytes.Buffer, key uint8) error {
	if key != 0 {
		for i := 0; i < buffer.Len(); i++ {
			buffer.Bytes()[i] ^= byte(key)
		}
	}
	if uint32(buffer.Len()) < 12 {
		return errors.New("message length error")
	}
	binary.Read(buffer, binary.LittleEndian, &p.FightType)
	binary.Read(buffer, binary.LittleEndian, &p.UseRoom)
	var RoomIDLen uint32
	binary.Read(buffer, binary.LittleEndian, &RoomIDLen)
	if uint32(buffer.Len()) < RoomIDLen {
		return errors.New("message length error")
	}
	p.RoomID = string(buffer.Next(int(RoomIDLen)))
	var LongRoomIDLen uint32
	binary.Read(buffer, binary.LittleEndian, &LongRoomIDLen)
	if uint32(buffer.Len()) < LongRoomIDLen {
		return errors.New("message length error")
	}
	p.LongRoomID = string(buffer.Next(int(LongRoomIDLen)))
	var ParamsLen uint32
	binary.Read(buffer, binary.LittleEndian, &ParamsLen)
	if uint32(buffer.Len()) < ParamsLen*4 {
		return errors.New("message length error")
	}
	p.Params = make([]int32, ParamsLen)
	for i := uint32(0); i < ParamsLen; i++ {
		binary.Read(buffer, binary.LittleEndian, &p.Params[i])
	}
	return nil
}

/*************************************  服务端  ********************************/

type S_Match_Result struct {
	Ret                int32
	FightPattern       int32
	FightType          int32
	FightModule        int32
	RoleIdentity       int32
	MatchRoles         map[int64]T_RoleAbstract
	FightServerIp      string
	FightServerPort    int32
	FightServerIpSSL   string
	FightServerPortSSL int32
	FightToken         string
	SeedId             []int32
	ExtraData          map[int64]T_Fight_Extra_Data
	BossIdIndexs       []int32
}

func (p *S_Match_Result) Encode() []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.LittleEndian, p.Ret)
	binary.Write(buffer, binary.LittleEndian, p.FightPattern)
	binary.Write(buffer, binary.LittleEndian, p.FightType)
	binary.Write(buffer, binary.LittleEndian, p.FightModule)
	binary.Write(buffer, binary.LittleEndian, p.RoleIdentity)
	binary.Write(buffer, binary.LittleEndian, uint32(len(p.MatchRoles)))
	for k, v := range p.MatchRoles {
		binary.Write(buffer, binary.LittleEndian, k)
		buffer.Write(v.Encode())
	}
	binary.Write(buffer, binary.LittleEndian, uint32(len(p.FightServerIp)))
	buffer.Write([]byte(p.FightServerIp))
	binary.Write(buffer, binary.LittleEndian, p.FightServerPort)
	binary.Write(buffer, binary.LittleEndian, uint32(len(p.FightServerIpSSL)))
	buffer.Write([]byte(p.FightServerIpSSL))
	binary.Write(buffer, binary.LittleEndian, p.FightServerPortSSL)
	binary.Write(buffer, binary.LittleEndian, uint32(len(p.FightToken)))
	buffer.Write([]byte(p.FightToken))
	binary.Write(buffer, binary.LittleEndian, uint32(len(p.SeedId)))
	for _, v := range p.SeedId {
		binary.Write(buffer, binary.LittleEndian, v)
	}
	binary.Write(buffer, binary.LittleEndian, uint32(len(p.ExtraData)))
	for k, v := range p.ExtraData {
		binary.Write(buffer, binary.LittleEndian, k)
		buffer.Write(v.Encode())
	}
	binary.Write(buffer, binary.LittleEndian, uint32(len(p.BossIdIndexs)))
	for _, v := range p.BossIdIndexs {
		binary.Write(buffer, binary.LittleEndian, v)
	}
	return buffer.Bytes()
}

func (p *S_Match_Result) Decode(buffer *bytes.Buffer) error {
	if buffer.Len() < 24 {
		return errors.New("message length error")
	}
	binary.Read(buffer, binary.LittleEndian, &p.Ret)
	binary.Read(buffer, binary.LittleEndian, &p.FightPattern)
	binary.Read(buffer, binary.LittleEndian, &p.FightType)
	binary.Read(buffer, binary.LittleEndian, &p.FightModule)
	binary.Read(buffer, binary.LittleEndian, &p.RoleIdentity)
	var MatchRolesLen uint32
	binary.Read(buffer, binary.LittleEndian, &MatchRolesLen)
	p.MatchRoles = make(map[int64]T_RoleAbstract, MatchRolesLen)
	for i := uint32(0); i < MatchRolesLen; i++ {
		var key int64
		binary.Read(buffer, binary.LittleEndian, &key)
		var value T_RoleAbstract
		value.Decode(buffer)
		p.MatchRoles[key] = value
	}
	if buffer.Len() < 4 {
		return errors.New("message length error")
	}
	var FightServerIpLen uint32
	binary.Read(buffer, binary.LittleEndian, &FightServerIpLen)
	if uint32(buffer.Len()) < FightServerIpLen {
		return errors.New("message length error")
	}
	p.FightServerIp = string(buffer.Next(int(FightServerIpLen)))
	if buffer.Len() < 8 {
		return errors.New("message length error")
	}
	binary.Read(buffer, binary.LittleEndian, &p.FightServerPort)
	var FightServerIpSSLLen uint32
	binary.Read(buffer, binary.LittleEndian, &FightServerIpSSLLen)
	if uint32(buffer.Len()) < FightServerIpSSLLen {
		return errors.New("message length error")
	}
	p.FightServerIpSSL = string(buffer.Next(int(FightServerIpSSLLen)))
	if buffer.Len() < 8 {
		return errors.New("message length error")
	}
	binary.Read(buffer, binary.LittleEndian, &p.FightServerPortSSL)
	var FightTokenLen uint32
	binary.Read(buffer, binary.LittleEndian, &FightTokenLen)
	if uint32(buffer.Len()) < FightTokenLen {
		return errors.New("message length error")
	}
	p.FightToken = string(buffer.Next(int(FightTokenLen)))
	if buffer.Len() < 4 {
		return errors.New("message length error")
	}
	var SeedIdLen uint32
	binary.Read(buffer, binary.LittleEndian, &SeedIdLen)
	if uint32(buffer.Len()) < SeedIdLen*4 {
		return errors.New("message length error")
	}
	p.SeedId = make([]int32, SeedIdLen)
	for i := uint32(0); i < SeedIdLen; i++ {
		binary.Read(buffer, binary.LittleEndian, &p.SeedId[i])
	}
	if buffer.Len() < 4 {
		return errors.New("message length error")
	}
	var ExtraDataLen uint32
	binary.Read(buffer, binary.LittleEndian, &ExtraDataLen)
	p.ExtraData = make(map[int64]T_Fight_Extra_Data, ExtraDataLen)
	for i := uint32(0); i < ExtraDataLen; i++ {
		var key int64
		binary.Read(buffer, binary.LittleEndian, &key)
		var value T_Fight_Extra_Data
		value.Decode(buffer)
		p.ExtraData[key] = value
	}
	if buffer.Len() < 4 {
		return errors.New("message length error")
	}
	var BossIdIndexsLen uint32
	binary.Read(buffer, binary.LittleEndian, &BossIdIndexsLen)
	if uint32(buffer.Len()) < BossIdIndexsLen*4 {
		return errors.New("message length error")
	}
	p.BossIdIndexs = make([]int32, BossIdIndexsLen)
	for i := uint32(0); i < BossIdIndexsLen; i++ {
		binary.Read(buffer, binary.LittleEndian, &p.BossIdIndexs[i])
	}
	return nil
}

type S_Match_Duel_Fight struct {
	Errorcode    int32
	FightType    int32
	RoomID       string
	LongRoomID   string
	RoleAbstract T_RoleAbstract
}

func (p *S_Match_Duel_Fight) Encode() []byte {
	buffer := bytes.NewBuffer([]byte{})
	binary.Write(buffer, binary.LittleEndian, p.Errorcode)
	binary.Write(buffer, binary.LittleEndian, p.FightType)
	binary.Write(buffer, binary.LittleEndian, uint32(len(p.RoomID)))
	buffer.Write([]byte(p.RoomID))
	binary.Write(buffer, binary.LittleEndian, uint32(len(p.LongRoomID)))
	buffer.Write([]byte(p.LongRoomID))
	buffer.Write(p.RoleAbstract.Encode())
	return buffer.Bytes()
}

func (p *S_Match_Duel_Fight) Decode(buffer *bytes.Buffer) error {
	if buffer.Len() < 12 {
		return errors.New("message length error")
	}
	binary.Read(buffer, binary.LittleEndian, &p.Errorcode)
	binary.Read(buffer, binary.LittleEndian, &p.FightType)
	if buffer.Len() < 4 {
		return errors.New("message length error")
	}
	var RoomIDLen uint32
	binary.Read(buffer, binary.LittleEndian, &RoomIDLen)
	if uint32(buffer.Len()) < RoomIDLen {
		return errors.New("message length error")
	}
	p.RoomID = string(buffer.Next(int(RoomIDLen)))
	if buffer.Len() < 4 {
		return errors.New("message length error")
	}
	var LongRoomIDLen uint32
	binary.Read(buffer, binary.LittleEndian, &LongRoomIDLen)
	if uint32(buffer.Len()) < LongRoomIDLen {
		return errors.New("message length error")
	}
	p.LongRoomID = string(buffer.Next(int(LongRoomIDLen)))
	p.RoleAbstract.Decode(buffer)
	return nil
}
