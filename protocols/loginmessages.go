package protocols

import (
	"bytes"
	"encoding/binary"
	"errors"
)

type T_Login_Role struct {
	RoleID int64
	ShowID string
	Name   string
	Level  int32
}

func (p *T_Login_Role) Encode() []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.LittleEndian, p.RoleID)
	binary.Write(buffer, binary.LittleEndian, uint32(len(p.ShowID)))
	buffer.Write([]byte(p.ShowID))
	binary.Write(buffer, binary.LittleEndian, uint32(len(p.Name)))
	buffer.Write([]byte(p.Name))
	binary.Write(buffer, binary.LittleEndian, p.Level)
	return buffer.Bytes()
}

func (p *T_Login_Role) Decode(buffer *bytes.Buffer) error {
	if buffer.Len() < 12 {
		return errors.New("message length error")
	}
	binary.Read(buffer, binary.LittleEndian, &p.RoleID)
	var length uint32
	binary.Read(buffer, binary.LittleEndian, &length)
	if buffer.Len() < int(length) {
		return errors.New("showID length error")
	}
	p.ShowID = string(buffer.Next(int(length)))
	binary.Read(buffer, binary.LittleEndian, &length)
	if buffer.Len() < int(length) {
		return errors.New("name length error")
	}
	p.Name = string(buffer.Next(int(length)))
	binary.Read(buffer, binary.LittleEndian, &p.Level)
	return nil
}

/****************************  客户端  *******************************/

type C_Login_ValidateOnline struct {
	SdkType    int32
	SdkID      int32
	PackageID  int32
	EntranceID int32
	ClientType int32
	AcountName string
	Sign       string
	HeadUrl    string
	Extra      []string
}

func (p *C_Login_ValidateOnline) Encode() []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.LittleEndian, p.SdkType)
	binary.Write(buffer, binary.LittleEndian, p.SdkID)
	binary.Write(buffer, binary.LittleEndian, p.PackageID)
	binary.Write(buffer, binary.LittleEndian, p.EntranceID)
	binary.Write(buffer, binary.LittleEndian, p.ClientType)
	binary.Write(buffer, binary.LittleEndian, uint32(len(p.AcountName)))
	buffer.Write([]byte(p.AcountName))
	binary.Write(buffer, binary.LittleEndian, uint32(len(p.Sign)))
	buffer.Write([]byte(p.Sign))
	binary.Write(buffer, binary.LittleEndian, uint32(len(p.HeadUrl)))
	buffer.Write([]byte(p.HeadUrl))
	binary.Write(buffer, binary.LittleEndian, uint32(len(p.Extra)))
	for _, value := range p.Extra {
		binary.Write(buffer, binary.LittleEndian, uint32(len(value)))
		buffer.Write([]byte(value))
	}
	return buffer.Bytes()
}

func (p *C_Login_ValidateOnline) Decode(buffer *bytes.Buffer) error {
	if buffer.Len() < 24 {
		return errors.New("message length error")
	}
	binary.Read(buffer, binary.LittleEndian, &p.SdkType)
	binary.Read(buffer, binary.LittleEndian, &p.SdkID)
	binary.Read(buffer, binary.LittleEndian, &p.PackageID)
	binary.Read(buffer, binary.LittleEndian, &p.EntranceID)
	binary.Read(buffer, binary.LittleEndian, &p.ClientType)
	var length uint32
	binary.Read(buffer, binary.LittleEndian, &length)
	if buffer.Len() < int(length) {
		return errors.New("accountName length error")
	}
	p.AcountName = string(buffer.Next(int(length)))
	binary.Read(buffer, binary.LittleEndian, &length)
	if buffer.Len() < int(length) {
		return errors.New("sign length error")
	}
	p.Sign = string(buffer.Next(int(length)))
	binary.Read(buffer, binary.LittleEndian, &length)
	if buffer.Len() < int(length) {
		return errors.New("headUrl length error")
	}
	p.HeadUrl = string(buffer.Next(int(length)))
	var count uint32
	binary.Read(buffer, binary.LittleEndian, &count)
	for i := 0; i < int(count); i++ {
		binary.Read(buffer, binary.LittleEndian, &length)
		value := string(buffer.Next(int(length)))
		p.Extra = append(p.Extra, value)
	}
	return nil
}

type C_Login_ChooseRole struct {
	RoleID int64
}

func (p *C_Login_ChooseRole) Encode() []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.LittleEndian, p.RoleID)
	return buffer.Bytes()
}

func (p *C_Login_ChooseRole) Decode(buffer *bytes.Buffer, key uint8) error {
	if key != 0 {
		for i := 0; i < buffer.Len(); i++ {
			buffer.Bytes()[i] ^= byte(key)
		}
	}
	binary.Read(buffer, binary.LittleEndian, &p.RoleID)
	return nil
}

/****************************  服务端  *******************************/

type S_Login_Validate struct {
	IsSucceed     bool
	ServerUTCTime int32
	Key           uint8
}

func (p *S_Login_Validate) Encode() []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.LittleEndian, p.IsSucceed)
	binary.Write(buffer, binary.LittleEndian, p.ServerUTCTime)
	binary.Write(buffer, binary.LittleEndian, p.Key)
	return buffer.Bytes()
}

func (p *S_Login_Validate) Decode(buffer *bytes.Buffer) error {
	if buffer.Len() < 6 {
		return errors.New("message length error")
	}
	binary.Read(buffer, binary.LittleEndian, &p.IsSucceed)
	binary.Read(buffer, binary.LittleEndian, &p.ServerUTCTime)
	binary.Read(buffer, binary.LittleEndian, &p.Key)
	return nil
}

type S_Login_RequestRole struct {
	BIndulge                 bool
	Roles                    map[int64]T_Login_Role
	ForbidLoginTimeRemaining int32
}

func (p *S_Login_RequestRole) Encode() []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.LittleEndian, p.BIndulge)
	binary.Write(buffer, binary.LittleEndian, uint32(len(p.Roles)))
	for k, v := range p.Roles {
		binary.Write(buffer, binary.LittleEndian, k)
		buffer.Write(v.Encode())
	}
	binary.Write(buffer, binary.LittleEndian, p.ForbidLoginTimeRemaining)
	return buffer.Bytes()
}

func (p *S_Login_RequestRole) Decode(buffer *bytes.Buffer) error {
	if buffer.Len() < 5 {
		return errors.New("message length error")
	}
	binary.Read(buffer, binary.LittleEndian, &p.BIndulge)
	var RolesLen uint32
	binary.Read(buffer, binary.LittleEndian, &RolesLen)
	p.Roles = make(map[int64]T_Login_Role, RolesLen)
	for i := 0; i < int(RolesLen); i++ {
		var key int64
		var value T_Login_Role
		binary.Read(buffer, binary.LittleEndian, &key)
		value.Decode(buffer)
		p.Roles[key] = value
	}
	binary.Read(buffer, binary.LittleEndian, &p.ForbidLoginTimeRemaining)
	return nil
}

type S_Login_ChooseRole struct {
	Result bool
}

func (p *S_Login_ChooseRole) Encode() []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.LittleEndian, p.Result)
	return buffer.Bytes()
}

func (p *S_Login_ChooseRole) Decode(buffer *bytes.Buffer) error {
	if buffer.Len() < 1 {
		return errors.New("message length error")
	}
	binary.Read(buffer, binary.LittleEndian, &p.Result)
	return nil
}
