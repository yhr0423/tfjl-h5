package protocols

import (
	"bytes"
	"encoding/binary"
	"errors"
)

type T_Sociaty_Player_Donate struct {
	DayNum int32
}

func (p *T_Sociaty_Player_Donate) Encode() []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.LittleEndian, p.DayNum)
	return buffer.Bytes()
}

func (p *T_Sociaty_Player_Donate) Decode(buffer *bytes.Buffer) error {
	if buffer.Len() < 4 {
		return errors.New("message length error")
	}
	binary.Read(buffer, binary.LittleEndian, &p.DayNum)
	return nil
}

type T_Sociaty_Player struct {
	ActiveLeaveSociaty     int32
	PassiveLeaveSociaty    int32
	SeriesSignInNum        int32
	DaySignInSociatyID     string
	DayReceivePrizeNum     int32
	SociatyID              string
	SociatyName            string
	SociatyLevel           int32
	SociatyFlag            int32
	Job                    int32
	Contribution           int32
	Donate                 map[int32]T_Sociaty_Player_Donate
	DayConvertSociatyMedal int32
	RedEnvelopes           []int64
	DurationTimes          []int32
}

func (p *T_Sociaty_Player) Encode() []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.LittleEndian, p.ActiveLeaveSociaty)
	binary.Write(buffer, binary.LittleEndian, p.PassiveLeaveSociaty)
	binary.Write(buffer, binary.LittleEndian, p.SeriesSignInNum)
	binary.Write(buffer, binary.LittleEndian, uint32(len(p.DaySignInSociatyID)))
	buffer.Write([]byte(p.DaySignInSociatyID))
	binary.Write(buffer, binary.LittleEndian, p.DayReceivePrizeNum)
	binary.Write(buffer, binary.LittleEndian, uint32(len(p.SociatyID)))
	buffer.Write([]byte(p.SociatyID))
	binary.Write(buffer, binary.LittleEndian, uint32(len(p.SociatyName)))
	buffer.Write([]byte(p.SociatyName))
	binary.Write(buffer, binary.LittleEndian, p.SociatyLevel)
	binary.Write(buffer, binary.LittleEndian, p.SociatyFlag)
	binary.Write(buffer, binary.LittleEndian, p.Job)
	binary.Write(buffer, binary.LittleEndian, p.Contribution)
	binary.Write(buffer, binary.LittleEndian, uint32(len(p.Donate)))
	for k, v := range p.Donate {
		binary.Write(buffer, binary.LittleEndian, k)
		buffer.Write(v.Encode())
	}
	binary.Write(buffer, binary.LittleEndian, p.DayConvertSociatyMedal)
	binary.Write(buffer, binary.LittleEndian, uint32(len(p.RedEnvelopes)))
	for _, v := range p.RedEnvelopes {
		binary.Write(buffer, binary.LittleEndian, v)
	}
	binary.Write(buffer, binary.LittleEndian, uint32(len(p.DurationTimes)))
	for _, v := range p.DurationTimes {
		binary.Write(buffer, binary.LittleEndian, v)
	}
	return buffer.Bytes()
}

func (p *T_Sociaty_Player) Decode(buffer *bytes.Buffer) error {
	if buffer.Len() < 16 {
		return errors.New("message length error")
	}
	binary.Read(buffer, binary.LittleEndian, &p.ActiveLeaveSociaty)
	binary.Read(buffer, binary.LittleEndian, &p.PassiveLeaveSociaty)
	binary.Read(buffer, binary.LittleEndian, &p.SeriesSignInNum)
	var DaySignInSociatyIDLen uint32
	binary.Read(buffer, binary.LittleEndian, &DaySignInSociatyIDLen)
	if uint32(buffer.Len()) < DaySignInSociatyIDLen {
		return errors.New("message length error")
	}
	p.DaySignInSociatyID = string(buffer.Next(int(DaySignInSociatyIDLen)))
	binary.Read(buffer, binary.LittleEndian, &p.DayReceivePrizeNum)
	var SociatyIDLen uint32
	binary.Read(buffer, binary.LittleEndian, &SociatyIDLen)
	if uint32(buffer.Len()) < SociatyIDLen {
		return errors.New("message length error")
	}
	p.SociatyID = string(buffer.Next(int(SociatyIDLen)))
	var SociatyNameLen uint32
	binary.Read(buffer, binary.LittleEndian, &SociatyNameLen)
	if uint32(buffer.Len()) < SociatyNameLen {
		return errors.New("message length error")
	}
	p.SociatyName = string(buffer.Next(int(SociatyNameLen)))
	binary.Read(buffer, binary.LittleEndian, &p.SociatyLevel)
	binary.Read(buffer, binary.LittleEndian, &p.SociatyFlag)
	binary.Read(buffer, binary.LittleEndian, &p.Job)
	binary.Read(buffer, binary.LittleEndian, &p.Contribution)
	var DonateLen uint32
	binary.Read(buffer, binary.LittleEndian, &DonateLen)
	if uint32(buffer.Len()) < DonateLen*8 {
		return errors.New("message length error")
	}
	p.Donate = make(map[int32]T_Sociaty_Player_Donate, DonateLen)
	for i := uint32(0); i < DonateLen; i++ {
		var k int32
		var v T_Sociaty_Player_Donate
		binary.Read(buffer, binary.LittleEndian, &k)
		if err := v.Decode(buffer); err != nil {
			return err
		}
		p.Donate[k] = v
	}
	binary.Read(buffer, binary.LittleEndian, &p.DayConvertSociatyMedal)
	var RedEnvelopesLen uint32
	binary.Read(buffer, binary.LittleEndian, &RedEnvelopesLen)
	if uint32(buffer.Len()) < RedEnvelopesLen*8 {
		return errors.New("message length error")
	}
	p.RedEnvelopes = make([]int64, RedEnvelopesLen)
	for i := uint32(0); i < RedEnvelopesLen; i++ {
		binary.Read(buffer, binary.LittleEndian, &p.RedEnvelopes[i])
	}
	var DurationTimesLen uint32
	binary.Read(buffer, binary.LittleEndian, &DurationTimesLen)
	if uint32(buffer.Len()) < DurationTimesLen*4 {
		return errors.New("message length error")
	}
	p.DurationTimes = make([]int32, DurationTimesLen)
	for i := uint32(0); i < DurationTimesLen; i++ {
		binary.Read(buffer, binary.LittleEndian, &p.DurationTimes[i])
	}
	return nil
}

type T_Sociaty_HeadData struct {
	RoleID      int64
	NickName    string
	HeadID      int32
	HeadUrl     string
	HeadFrameID int32
	Score       int32
	DayMaxRound int32
}

func (p *T_Sociaty_HeadData) Encode() []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.LittleEndian, p.RoleID)
	binary.Write(buffer, binary.LittleEndian, uint32(len(p.NickName)))
	buffer.Write([]byte(p.NickName))
	binary.Write(buffer, binary.LittleEndian, p.HeadID)
	binary.Write(buffer, binary.LittleEndian, uint32(len(p.HeadUrl)))
	buffer.Write([]byte(p.HeadUrl))
	binary.Write(buffer, binary.LittleEndian, p.HeadFrameID)
	binary.Write(buffer, binary.LittleEndian, p.Score)
	binary.Write(buffer, binary.LittleEndian, p.DayMaxRound)
	return buffer.Bytes()
}

func (p *T_Sociaty_HeadData) Decode(buffer *bytes.Buffer) error {
	if buffer.Len() < 12 {
		return errors.New("message length error")
	}
	binary.Read(buffer, binary.LittleEndian, &p.RoleID)
	var NickNameLen uint32
	binary.Read(buffer, binary.LittleEndian, &NickNameLen)
	if uint32(buffer.Len()) < NickNameLen {
		return errors.New("message length error")
	}
	p.NickName = string(buffer.Next(int(NickNameLen)))
	binary.Read(buffer, binary.LittleEndian, &p.HeadID)
	var HeadUrlLen uint32
	binary.Read(buffer, binary.LittleEndian, &HeadUrlLen)
	if uint32(buffer.Len()) < HeadUrlLen {
		return errors.New("message length error")
	}
	p.HeadUrl = string(buffer.Next(int(HeadUrlLen)))
	binary.Read(buffer, binary.LittleEndian, &p.HeadFrameID)
	binary.Read(buffer, binary.LittleEndian, &p.Score)
	binary.Read(buffer, binary.LittleEndian, &p.DayMaxRound)
	return nil
}

type T_Sociaty_MachinariumRankListData struct {
	RankId     int32
	Roles      map[int64]T_Sociaty_HeadData
	Round      int32
	TotalTime  int32
	UpdateTime int32
}

func (p *T_Sociaty_MachinariumRankListData) Encode() []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.LittleEndian, p.RankId)
	binary.Write(buffer, binary.LittleEndian, uint32(len(p.Roles)))
	for k, v := range p.Roles {
		binary.Write(buffer, binary.LittleEndian, k)
		buffer.Write(v.Encode())
	}
	binary.Write(buffer, binary.LittleEndian, p.Round)
	binary.Write(buffer, binary.LittleEndian, p.TotalTime)
	binary.Write(buffer, binary.LittleEndian, p.UpdateTime)
	return buffer.Bytes()
}

func (p *T_Sociaty_MachinariumRankListData) Decode(buffer *bytes.Buffer) error {
	if buffer.Len() < 8 {
		return errors.New("message length error")
	}
	binary.Read(buffer, binary.LittleEndian, &p.RankId)
	var RolesLen uint32
	binary.Read(buffer, binary.LittleEndian, &RolesLen)
	p.Roles = make(map[int64]T_Sociaty_HeadData, RolesLen)
	for i := uint32(0); i < RolesLen; i++ {
		var k int64
		var v T_Sociaty_HeadData
		binary.Read(buffer, binary.LittleEndian, &k)
		if err := v.Decode(buffer); err != nil {
			return err
		}
		p.Roles[k] = v
	}
	binary.Read(buffer, binary.LittleEndian, &p.Round)
	binary.Read(buffer, binary.LittleEndian, &p.TotalTime)
	binary.Read(buffer, binary.LittleEndian, &p.UpdateTime)
	return nil
}

/**********************************  客户端  **********************************/
type C_Sociaty_RoleGetMachinariumData struct {
	RoleID int64
}

func (p *C_Sociaty_RoleGetMachinariumData) Encode() []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.LittleEndian, p.RoleID)
	return buffer.Bytes()
}

func (p *C_Sociaty_RoleGetMachinariumData) Decode(buffer *bytes.Buffer, key uint8) error {
	if key != 0 {
		for i := 0; i < buffer.Len(); i++ {
			buffer.Bytes()[i] ^= byte(key)
		}
	}
	if buffer.Len() < 8 {
		return errors.New("message length error")
	}
	binary.Read(buffer, binary.LittleEndian, &p.RoleID)
	return nil
}

type C_Sociaty_RoleMachinariumSelectCard struct {
	CardID int32
}

func (p *C_Sociaty_RoleMachinariumSelectCard) Encode() []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.LittleEndian, p.CardID)
	return buffer.Bytes()
}

func (p *C_Sociaty_RoleMachinariumSelectCard) Decode(buffer *bytes.Buffer, key uint8) error {
	if key != 0 {
		for i := 0; i < buffer.Len(); i++ {
			buffer.Bytes()[i] ^= byte(key)
		}
	}
	if buffer.Len() < 4 {
		return errors.New("message length error")
	}
	binary.Read(buffer, binary.LittleEndian, &p.CardID)
	return nil
}

/**********************************  服务器  **********************************/
type S_Sociaty_SynData struct {
	PlayerData T_Sociaty_Player
}

func (p *S_Sociaty_SynData) Encode() []byte {
	buffer := new(bytes.Buffer)
	buffer.Write(p.PlayerData.Encode())
	return buffer.Bytes()
}

func (p *S_Sociaty_SynData) Decode(buffer *bytes.Buffer) error {
	if err := p.PlayerData.Decode(buffer); err != nil {
		return err
	}
	return nil
}

type S_Sociaty_SyncMachinariumData struct {
	SociatyId    int32
	RoleID       int64
	CardID       int32
	DayMaxRound  int32
	DayFailNum   int32
	RankListData []T_Sociaty_MachinariumRankListData
	AllRound     int32
	RewardStatus map[int32]int32
}

func (p *S_Sociaty_SyncMachinariumData) Encode() []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.LittleEndian, p.SociatyId)
	binary.Write(buffer, binary.LittleEndian, p.RoleID)
	binary.Write(buffer, binary.LittleEndian, p.CardID)
	binary.Write(buffer, binary.LittleEndian, p.DayMaxRound)
	binary.Write(buffer, binary.LittleEndian, p.DayFailNum)
	binary.Write(buffer, binary.LittleEndian, uint32(len(p.RankListData)))
	for _, v := range p.RankListData {
		buffer.Write(v.Encode())
	}
	binary.Write(buffer, binary.LittleEndian, p.AllRound)
	binary.Write(buffer, binary.LittleEndian, uint32(len(p.RewardStatus)))
	for k, v := range p.RewardStatus {
		binary.Write(buffer, binary.LittleEndian, k)
		binary.Write(buffer, binary.LittleEndian, v)
	}
	return buffer.Bytes()
}

func (p *S_Sociaty_SyncMachinariumData) Decode(buffer *bytes.Buffer) error {
	if buffer.Len() < 28 {
		return errors.New("message length error")
	}
	binary.Read(buffer, binary.LittleEndian, &p.SociatyId)
	binary.Read(buffer, binary.LittleEndian, &p.RoleID)
	binary.Read(buffer, binary.LittleEndian, &p.CardID)

	binary.Read(buffer, binary.LittleEndian, &p.DayMaxRound)
	binary.Read(buffer, binary.LittleEndian, &p.DayFailNum)
	var RankListDataLen uint32
	binary.Read(buffer, binary.LittleEndian, &RankListDataLen)
	p.RankListData = make([]T_Sociaty_MachinariumRankListData, RankListDataLen)
	for i := uint32(0); i < RankListDataLen; i++ {
		if err := p.RankListData[i].Decode(buffer); err != nil {
			return err
		}
	}
	binary.Read(buffer, binary.LittleEndian, &p.AllRound)
	var RewardStatusLen uint32
	binary.Read(buffer, binary.LittleEndian, &RewardStatusLen)
	if uint32(buffer.Len()) < RewardStatusLen*8 {
		return errors.New("message length error")
	}
	p.RewardStatus = make(map[int32]int32, RewardStatusLen)
	for i := uint32(0); i < RewardStatusLen; i++ {
		var k int32
		var v int32
		binary.Read(buffer, binary.LittleEndian, &k)
		binary.Read(buffer, binary.LittleEndian, &v)
		p.RewardStatus[k] = v
	}
	return nil
}

type S_Sociaty_RoleMachinariumSelectCard struct {
	Errorcode int32
	CardID    int32
}

func (p *S_Sociaty_RoleMachinariumSelectCard) Encode() []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.LittleEndian, p.Errorcode)
	binary.Write(buffer, binary.LittleEndian, p.CardID)
	return buffer.Bytes()
}

func (p *S_Sociaty_RoleMachinariumSelectCard) Decode(buffer *bytes.Buffer) error {
	if buffer.Len() < 8 {
		return errors.New("message length error")
	}
	binary.Read(buffer, binary.LittleEndian, &p.Errorcode)
	binary.Read(buffer, binary.LittleEndian, &p.CardID)
	return nil
}
