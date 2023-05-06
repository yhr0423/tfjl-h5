package protocols

import (
	"bytes"
	"encoding/binary"
	"errors"
)

type T_Fight_Silver_Struct struct {
	Silver           int32
	SilverChange     int32
	Type             int32
	SilverChangeType int32
}

func (p *T_Fight_Silver_Struct) Encode() []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.LittleEndian, p.Silver)
	binary.Write(buffer, binary.LittleEndian, p.SilverChange)
	binary.Write(buffer, binary.LittleEndian, p.Type)
	binary.Write(buffer, binary.LittleEndian, p.SilverChangeType)
	return buffer.Bytes()
}

func (p *T_Fight_Silver_Struct) Decode(buffer *bytes.Buffer) error {
	if buffer.Len() < 16 {
		return errors.New("message length error")
	}
	binary.Read(buffer, binary.LittleEndian, &p.Silver)
	binary.Read(buffer, binary.LittleEndian, &p.SilverChange)
	binary.Read(buffer, binary.LittleEndian, &p.Type)
	binary.Read(buffer, binary.LittleEndian, &p.SilverChangeType)
	return nil
}

type T_Blood_Struct struct {
	DropBlood      int32
	CurBlood       int32
	MaxBlood       int32
	AddBlood       int32
	Type           int32
	HeroId         int32
	ExtraDropMoney int32
	DamageType     int32
	MaxDropBlood   int32
}

func (p *T_Blood_Struct) Encode() []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.LittleEndian, p.DropBlood)
	binary.Write(buffer, binary.LittleEndian, p.CurBlood)
	binary.Write(buffer, binary.LittleEndian, p.MaxBlood)
	binary.Write(buffer, binary.LittleEndian, p.AddBlood)
	binary.Write(buffer, binary.LittleEndian, p.Type)
	binary.Write(buffer, binary.LittleEndian, p.HeroId)
	binary.Write(buffer, binary.LittleEndian, p.ExtraDropMoney)
	binary.Write(buffer, binary.LittleEndian, p.DamageType)
	binary.Write(buffer, binary.LittleEndian, p.MaxDropBlood)
	return buffer.Bytes()
}

func (p *T_Blood_Struct) Decode(buffer *bytes.Buffer) error {
	if buffer.Len() < 36 {
		return errors.New("message length error")
	}
	binary.Read(buffer, binary.LittleEndian, &p.DropBlood)
	binary.Read(buffer, binary.LittleEndian, &p.CurBlood)
	binary.Read(buffer, binary.LittleEndian, &p.MaxBlood)
	binary.Read(buffer, binary.LittleEndian, &p.AddBlood)
	binary.Read(buffer, binary.LittleEndian, &p.Type)
	binary.Read(buffer, binary.LittleEndian, &p.HeroId)
	binary.Read(buffer, binary.LittleEndian, &p.ExtraDropMoney)
	binary.Read(buffer, binary.LittleEndian, &p.DamageType)
	binary.Read(buffer, binary.LittleEndian, &p.MaxDropBlood)
	return nil
}

type T_Monster_Blood_Struct struct {
	MonsterBlood map[int32]T_Blood_Struct
}

func (p *T_Monster_Blood_Struct) Encode() []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.LittleEndian, uint32(len(p.MonsterBlood)))
	for k, v := range p.MonsterBlood {
		binary.Write(buffer, binary.LittleEndian, k)
		buffer.Write(v.Encode())
	}
	return buffer.Bytes()
}

func (p *T_Monster_Blood_Struct) Decode(buffer *bytes.Buffer) error {
	if buffer.Len() < 4 {
		return errors.New("message length error")
	}
	var MonsterBloodLen uint32
	binary.Read(buffer, binary.LittleEndian, &MonsterBloodLen)
	if uint32(buffer.Len()) < MonsterBloodLen*40 {
		return errors.New("message length error")
	}
	p.MonsterBlood = make(map[int32]T_Blood_Struct, MonsterBloodLen)
	for i := uint32(0); i < MonsterBloodLen; i++ {
		var k int32
		var v T_Blood_Struct
		binary.Read(buffer, binary.LittleEndian, &k)
		if err := v.Decode(buffer); err != nil {
			return err
		}
		p.MonsterBlood[k] = v
	}
	return nil
}

type T_Player_Monster_Blood_Struct struct {
	MosterBlood map[int64]T_Monster_Blood_Struct
}

func (p *T_Player_Monster_Blood_Struct) Encode() []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.LittleEndian, uint32(len(p.MosterBlood)))
	for k, v := range p.MosterBlood {
		binary.Write(buffer, binary.LittleEndian, k)
		buffer.Write(v.Encode())
	}
	return buffer.Bytes()
}

func (p *T_Player_Monster_Blood_Struct) Decode(buffer *bytes.Buffer) error {
	if buffer.Len() < 4 {
		return errors.New("message length error")
	}
	var MosterBloodLen uint32
	binary.Read(buffer, binary.LittleEndian, &MosterBloodLen)
	p.MosterBlood = make(map[int64]T_Monster_Blood_Struct, MosterBloodLen)
	for i := uint32(0); i < MosterBloodLen; i++ {
		var k int64
		var v T_Monster_Blood_Struct
		binary.Read(buffer, binary.LittleEndian, &k)
		if err := v.Decode(buffer); err != nil {
			return err
		}
		p.MosterBlood[k] = v
	}
	return nil
}

type T_Hero_Attr_Struct struct {
	ID   string
	Attr []int32
}

func (p *T_Hero_Attr_Struct) Encode() []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.LittleEndian, uint32(len(p.ID)))
	buffer.Write([]byte(p.ID))
	binary.Write(buffer, binary.LittleEndian, uint32(len(p.Attr)))
	for _, v := range p.Attr {
		binary.Write(buffer, binary.LittleEndian, v)
	}
	return buffer.Bytes()
}

func (p *T_Hero_Attr_Struct) Decode(buffer *bytes.Buffer) error {
	if buffer.Len() < 4 {
		return errors.New("message length error")
	}
	var IDLen uint32
	binary.Read(buffer, binary.LittleEndian, &IDLen)
	if uint32(buffer.Len()) < IDLen {
		return errors.New("message length error")
	}
	p.ID = string(buffer.Next(int(IDLen)))
	if buffer.Len() < 4 {
		return errors.New("message length error")
	}
	var AttrLen uint32
	binary.Read(buffer, binary.LittleEndian, &AttrLen)
	if uint32(buffer.Len()) < AttrLen*4 {
		return errors.New("message length error")
	}
	p.Attr = make([]int32, AttrLen)
	for i := uint32(0); i < AttrLen; i++ {
		binary.Read(buffer, binary.LittleEndian, &p.Attr[i])
	}
	return nil
}

type T_Update_Hero_Struct struct {
	Camp           int32
	Pos            int32
	HeroTemplateId int32
	HeroLv         int32
	HeroNpcId      int32
	HeroTrainLv    int32
}

func (p *T_Update_Hero_Struct) Encode() []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.LittleEndian, p.Camp)
	binary.Write(buffer, binary.LittleEndian, p.Pos)
	binary.Write(buffer, binary.LittleEndian, p.HeroTemplateId)
	binary.Write(buffer, binary.LittleEndian, p.HeroLv)
	binary.Write(buffer, binary.LittleEndian, p.HeroNpcId)
	binary.Write(buffer, binary.LittleEndian, p.HeroTrainLv)
	return buffer.Bytes()
}

func (p *T_Update_Hero_Struct) Decode(buffer *bytes.Buffer) error {
	if buffer.Len() < 24 {
		return errors.New("message length error")
	}
	binary.Read(buffer, binary.LittleEndian, &p.Camp)
	binary.Read(buffer, binary.LittleEndian, &p.Pos)
	binary.Read(buffer, binary.LittleEndian, &p.HeroTemplateId)
	binary.Read(buffer, binary.LittleEndian, &p.HeroLv)
	binary.Read(buffer, binary.LittleEndian, &p.HeroNpcId)
	binary.Read(buffer, binary.LittleEndian, &p.HeroTrainLv)
	return nil
}

type EquipItem struct {
	ItemID int32
	ItemNum int32
}

func (p *EquipItem) Encode() []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.LittleEndian, p.ItemID)
	binary.Write(buffer, binary.LittleEndian, p.ItemNum)
	return buffer.Bytes()
}

func (p *EquipItem) Decode(buffer *bytes.Buffer) error {
	if buffer.Len() < 8 {
		return errors.New("message length error")
	}
	binary.Read(buffer, binary.LittleEndian, &p.ItemID)
	binary.Read(buffer, binary.LittleEndian, &p.ItemNum)
	return nil
}

/**************************************  客户端  **************************************/
// 战斗结果提交服务器（单人）
type C_Fight_Report_Result_To_Logic struct {
	ReportData T_Fight_Detail_Info
}

func (p *C_Fight_Report_Result_To_Logic) Encode() []byte {
	buffer := new(bytes.Buffer)
	buffer.Write(p.ReportData.Encode())
	return buffer.Bytes()
}

func (p *C_Fight_Report_Result_To_Logic) Decode(buffer *bytes.Buffer, key uint8) error {
	if key != 0 {
		for i := 0; i < buffer.Len(); i++ {
			buffer.Bytes()[i] ^= byte(key)
		}
	}
	return p.ReportData.Decode(buffer)
}

type C_Fight_Role_Login struct {
	RoleID       int64
	MatchRoles   map[int64]T_RoleAbstract
	FightToken   string
	FightPattern int32
	FightType    int32
}

func (p *C_Fight_Role_Login) Encode() []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.LittleEndian, p.RoleID)
	binary.Write(buffer, binary.LittleEndian, uint32(len(p.MatchRoles)))
	for k, v := range p.MatchRoles {
		binary.Write(buffer, binary.LittleEndian, k)
		buffer.Write(v.Encode())
	}
	binary.Write(buffer, binary.LittleEndian, uint32(len(p.FightToken)))
	buffer.Write([]byte(p.FightToken))
	binary.Write(buffer, binary.LittleEndian, p.FightPattern)
	binary.Write(buffer, binary.LittleEndian, p.FightType)
	return buffer.Bytes()
}

func (p *C_Fight_Role_Login) Decode(buffer *bytes.Buffer) error {
	if buffer.Len() < 12 {
		return errors.New("message length error")
	}
	binary.Read(buffer, binary.LittleEndian, &p.RoleID)
	var MatchRolesLen uint32
	binary.Read(buffer, binary.LittleEndian, &MatchRolesLen)
	p.MatchRoles = make(map[int64]T_RoleAbstract, MatchRolesLen)
	for i := uint32(0); i < MatchRolesLen; i++ {
		var k int64
		var v T_RoleAbstract
		binary.Read(buffer, binary.LittleEndian, &k)
		if err := v.Decode(buffer); err != nil {
			return err
		}
		p.MatchRoles[k] = v
	}
	if buffer.Len() < 4 {
		return errors.New("message length error")
	}
	var FightTokenLen uint32
	binary.Read(buffer, binary.LittleEndian, &FightTokenLen)
	if uint32(buffer.Len()) < FightTokenLen {
		return errors.New("message length error")
	}
	p.FightToken = string(buffer.Next(int(FightTokenLen)))
	if buffer.Len() < 8 {
		return errors.New("message length error")
	}
	binary.Read(buffer, binary.LittleEndian, &p.FightPattern)
	binary.Read(buffer, binary.LittleEndian, &p.FightType)
	return nil
}

type C_Fight_Silver_SYNC struct {
	RoleID    int64
	Silver    T_Fight_Silver_Struct
	TimeFrame int32
}

func (p *C_Fight_Silver_SYNC) Encode() []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.LittleEndian, p.RoleID)
	buffer.Write(p.Silver.Encode())
	binary.Write(buffer, binary.LittleEndian, p.TimeFrame)
	return buffer.Bytes()
}

func (p *C_Fight_Silver_SYNC) Decode(buffer *bytes.Buffer, key uint8) error {
	if key != 0 {
		for i := 0; i < buffer.Len(); i++ {
			buffer.Bytes()[i] ^= byte(key)
		}
	}
	if buffer.Len() < 28 {
		return errors.New("message length error")
	}
	binary.Read(buffer, binary.LittleEndian, &p.RoleID)
	if err := p.Silver.Decode(buffer); err != nil {
		return err
	}
	binary.Read(buffer, binary.LittleEndian, &p.TimeFrame)
	return nil
}

type C_Fight_Loading_Ready struct {
	RoleID     int64
	FightToken string
}

func (p *C_Fight_Loading_Ready) Encode() []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.LittleEndian, p.RoleID)
	binary.Write(buffer, binary.LittleEndian, uint32(len(p.FightToken)))
	buffer.Write([]byte(p.FightToken))
	return buffer.Bytes()
}

func (p *C_Fight_Loading_Ready) Decode(buffer *bytes.Buffer, key uint8) error {
	if key != 0 {
		for i := 0; i < buffer.Len(); i++ {
			buffer.Bytes()[i] ^= byte(key)
		}
	}
	if buffer.Len() < 12 {
		return errors.New("message length error")
	}
	binary.Read(buffer, binary.LittleEndian, &p.RoleID)
	var FightTokenLen uint32
	binary.Read(buffer, binary.LittleEndian, &FightTokenLen)
	if uint32(buffer.Len()) < FightTokenLen {
		return errors.New("message length error")
	}
	p.FightToken = string(buffer.Next(int(FightTokenLen)))
	return nil
}

type C_Fight_Refresh_Card_Count_SYNC struct {
	RoleID     int64
	RefreshNum int32
	CardIdVec  []int32
}

func (p *C_Fight_Refresh_Card_Count_SYNC) Encode() []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.LittleEndian, p.RoleID)
	binary.Write(buffer, binary.LittleEndian, p.RefreshNum)
	binary.Write(buffer, binary.LittleEndian, uint32(len(p.CardIdVec)))
	for _, v := range p.CardIdVec {
		binary.Write(buffer, binary.LittleEndian, v)
	}
	return buffer.Bytes()
}

func (p *C_Fight_Refresh_Card_Count_SYNC) Decode(buffer *bytes.Buffer, key uint8) error {
	if key != 0 {
		for i := 0; i < buffer.Len(); i++ {
			buffer.Bytes()[i] ^= byte(key)
		}
	}
	if buffer.Len() < 16 {
		return errors.New("message length error")
	}
	binary.Read(buffer, binary.LittleEndian, &p.RoleID)
	binary.Read(buffer, binary.LittleEndian, &p.RefreshNum)
	var CardIdVecLen uint32
	binary.Read(buffer, binary.LittleEndian, &CardIdVecLen)
	if uint32(buffer.Len()) < CardIdVecLen*4 {
		return errors.New("message length error")
	}
	p.CardIdVec = make([]int32, CardIdVecLen)
	for i := uint32(0); i < CardIdVecLen; i++ {
		binary.Read(buffer, binary.LittleEndian, &p.CardIdVec[i])
	}
	return nil
}

type C_Fight_Monster_Blood_SYNC struct {
	TimeFrame          int32
	PlayerMonsterBlood map[int64]T_Player_Monster_Blood_Struct
}

func (p *C_Fight_Monster_Blood_SYNC) Encode() []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.LittleEndian, p.TimeFrame)
	binary.Write(buffer, binary.LittleEndian, uint32(len(p.PlayerMonsterBlood)))
	for k, v := range p.PlayerMonsterBlood {
		binary.Write(buffer, binary.LittleEndian, k)
		buffer.Write(v.Encode())
	}
	return buffer.Bytes()
}

func (p *C_Fight_Monster_Blood_SYNC) Decode(buffer *bytes.Buffer, key uint8) error {
	if key != 0 {
		for i := 0; i < buffer.Len(); i++ {
			buffer.Bytes()[i] ^= byte(key)
		}
	}
	if buffer.Len() < 8 {
		return errors.New("message length error")
	}
	binary.Read(buffer, binary.LittleEndian, &p.TimeFrame)
	var PlayerMonsterBloodLen uint32
	binary.Read(buffer, binary.LittleEndian, &PlayerMonsterBloodLen)
	p.PlayerMonsterBlood = make(map[int64]T_Player_Monster_Blood_Struct, PlayerMonsterBloodLen)
	for i := uint32(0); i < PlayerMonsterBloodLen; i++ {
		var k int64
		binary.Read(buffer, binary.LittleEndian, &k)
		var v T_Player_Monster_Blood_Struct
		if err := v.Decode(buffer); err != nil {
			return err
		}
		p.PlayerMonsterBlood[k] = v
	}
	return nil
}

type C_Fight_Hero_Attr_SYNC struct {
	RoleID      int64
	HeroAttrVec []T_Hero_Attr_Struct
}

func (p *C_Fight_Hero_Attr_SYNC) Encode() []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.LittleEndian, p.RoleID)
	binary.Write(buffer, binary.LittleEndian, uint32(len(p.HeroAttrVec)))
	for _, v := range p.HeroAttrVec {
		buffer.Write(v.Encode())
	}
	return buffer.Bytes()
}

func (p *C_Fight_Hero_Attr_SYNC) Decode(buffer *bytes.Buffer, key uint8) error {
	if key != 0 {
		for i := 0; i < buffer.Len(); i++ {
			buffer.Bytes()[i] ^= byte(key)
		}
	}
	if buffer.Len() < 12 {
		return errors.New("message length error")
	}
	binary.Read(buffer, binary.LittleEndian, &p.RoleID)
	var HeroAttrVecLen uint32
	binary.Read(buffer, binary.LittleEndian, &HeroAttrVecLen)
	p.HeroAttrVec = make([]T_Hero_Attr_Struct, HeroAttrVecLen)
	for i := uint32(0); i < HeroAttrVecLen; i++ {
		if err := p.HeroAttrVec[i].Decode(buffer); err != nil {
			return err
		}
	}
	return nil
}

type C_Fight_Update_Hero_SYNC struct {
	RoleID     int64
	UpdateHero T_Update_Hero_Struct
}

func (p *C_Fight_Update_Hero_SYNC) Encode() []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.LittleEndian, p.RoleID)
	buffer.Write(p.UpdateHero.Encode())
	return buffer.Bytes()
}

func (p *C_Fight_Update_Hero_SYNC) Decode(buffer *bytes.Buffer, key uint8) error {
	if key != 0 {
		for i := 0; i < buffer.Len(); i++ {
			buffer.Bytes()[i] ^= byte(key)
		}
	}
	if buffer.Len() < 32 {
		return errors.New("message length error")
	}
	binary.Read(buffer, binary.LittleEndian, &p.RoleID)
	if err := p.UpdateHero.Decode(buffer); err != nil {
		return err
	}
	return nil
}

type C_Fight_Operate_Equip_SYNC struct {
	RoleID      int64
	OperateType int32
	CarId       int32
	Target      int32
}

func (p *C_Fight_Operate_Equip_SYNC) Encode() []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.LittleEndian, p.RoleID)
	binary.Write(buffer, binary.LittleEndian, p.OperateType)
	binary.Write(buffer, binary.LittleEndian, p.CarId)
	binary.Write(buffer, binary.LittleEndian, p.Target)
	return buffer.Bytes()
}

func (p *C_Fight_Operate_Equip_SYNC) Decode(buffer *bytes.Buffer, key uint8) error {
	if key != 0 {
		for i := 0; i < buffer.Len(); i++ {
			buffer.Bytes()[i] ^= byte(key)
		}
	}
	if buffer.Len() < 20 {
		return errors.New("message length error")
	}
	binary.Read(buffer, binary.LittleEndian, &p.RoleID)
	binary.Read(buffer, binary.LittleEndian, &p.OperateType)
	binary.Read(buffer, binary.LittleEndian, &p.CarId)
	binary.Read(buffer, binary.LittleEndian, &p.Target)
	return nil
}

// 战斗结果提交服务器（多人）
type C_Fight_Report_Result_To_Fight struct {
	ReportData T_Fight_Detail_Info
}

func (p *C_Fight_Report_Result_To_Fight) Encode() []byte {
	buffer := new(bytes.Buffer)
	buffer.Write(p.ReportData.Encode())
	return buffer.Bytes()
}

func (p *C_Fight_Report_Result_To_Fight) Decode(buffer *bytes.Buffer, key uint8) error {
	if key != 0 {
		for i := 0; i < buffer.Len(); i++ {
			buffer.Bytes()[i] ^= byte(key)
		}
	}
	return p.ReportData.Decode(buffer)
}

/**************************************  服务器  **************************************/
// 服务器返回战斗结果错误码（单人）
type S_Fight_Report_Result_To_Logic struct {
	Errorcode int32
}

func (p *S_Fight_Report_Result_To_Logic) Encode() []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.LittleEndian, p.Errorcode)
	return buffer.Bytes()
}

func (p *S_Fight_Report_Result_To_Logic) Decode(buffer *bytes.Buffer) error {
	if buffer.Len() < 4 {
		return errors.New("message length error")
	}
	binary.Read(buffer, binary.LittleEndian, &p.Errorcode)
	return nil
}

type S_Fight_Role_Login struct {
	Errorcode int32
	Key       uint8
}

func (p *S_Fight_Role_Login) Encode() []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.LittleEndian, p.Errorcode)
	binary.Write(buffer, binary.LittleEndian, p.Key)
	return buffer.Bytes()
}

func (p *S_Fight_Role_Login) Decode(buffer *bytes.Buffer) error {
	if buffer.Len() < 5 {
		return errors.New("message length error")
	}
	binary.Read(buffer, binary.LittleEndian, &p.Errorcode)
	binary.Read(buffer, binary.LittleEndian, &p.Key)
	return nil
}

// 银币同步
type S_Fight_Silver_SYNC struct {
	RoleID    int64
	Silver    T_Fight_Silver_Struct
	TimeFrame int32
}

func (p *S_Fight_Silver_SYNC) Encode() []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.LittleEndian, p.RoleID)
	buffer.Write(p.Silver.Encode())
	binary.Write(buffer, binary.LittleEndian, p.TimeFrame)
	return buffer.Bytes()
}

func (p *S_Fight_Silver_SYNC) Decode(buffer *bytes.Buffer) error {
	if buffer.Len() < 28 {
		return errors.New("message length error")
	}
	binary.Read(buffer, binary.LittleEndian, &p.RoleID)
	if err := p.Silver.Decode(buffer); err != nil {
		return err
	}
	binary.Read(buffer, binary.LittleEndian, &p.TimeFrame)
	return nil
}

type S_Fight_Loading_Ready struct {
	Errorcode int32
}

func (p *S_Fight_Loading_Ready) Encode() []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.LittleEndian, p.Errorcode)
	return buffer.Bytes()
}

func (p *S_Fight_Loading_Ready) Decode(buffer *bytes.Buffer) error {
	if buffer.Len() < 4 {
		return errors.New("message length error")
	}
	binary.Read(buffer, binary.LittleEndian, &p.Errorcode)
	return nil
}

type S_Fight_FightStart struct{}

func (p *S_Fight_FightStart) Encode() []byte {
	buffer := new(bytes.Buffer)
	return buffer.Bytes()
}

func (p *S_Fight_FightStart) Decode(buffer *bytes.Buffer) error {
	return nil
}

type S_Fight_FightEnd struct {
	Reason    int32
	WinRoleID int64
}

func (p *S_Fight_FightEnd) Encode() []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.LittleEndian, p.Reason)
	binary.Write(buffer, binary.LittleEndian, p.WinRoleID)
	return buffer.Bytes()
}

func (p *S_Fight_FightEnd) Decode(buffer *bytes.Buffer) error {
	if buffer.Len() < 12 {
		return errors.New("buffer.Len() < 12")
	}
	binary.Read(buffer, binary.LittleEndian, &p.Reason)
	binary.Read(buffer, binary.LittleEndian, &p.WinRoleID)
	return nil
}

type S_Fight_Refresh_Card_Count_SYNC struct {
	RoleID     int64
	RefreshNum int32
	CardIdVec  []int32
}

func (p *S_Fight_Refresh_Card_Count_SYNC) Encode() []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.LittleEndian, p.RoleID)
	binary.Write(buffer, binary.LittleEndian, p.RefreshNum)
	binary.Write(buffer, binary.LittleEndian, uint32(len(p.CardIdVec)))
	for _, v := range p.CardIdVec {
		binary.Write(buffer, binary.LittleEndian, v)
	}
	return buffer.Bytes()
}

func (p *S_Fight_Refresh_Card_Count_SYNC) Decode(buffer *bytes.Buffer) error {
	if buffer.Len() < 16 {
		return errors.New("message length error")
	}
	binary.Read(buffer, binary.LittleEndian, &p.RoleID)
	binary.Read(buffer, binary.LittleEndian, &p.RefreshNum)
	var CardIdVecLen uint32
	binary.Read(buffer, binary.LittleEndian, &CardIdVecLen)
	if uint32(buffer.Len()) < CardIdVecLen*4 {
		return errors.New("message length error")
	}
	p.CardIdVec = make([]int32, CardIdVecLen)
	for i := uint32(0); i < CardIdVecLen; i++ {
		binary.Read(buffer, binary.LittleEndian, &p.CardIdVec[i])
	}
	return nil
}

type S_Fight_Monster_Blood_SYNC struct {
	TimeFrame    int32
	SyncRoleID   int64
	MonsterBlood map[int64]T_Monster_Blood_Struct
}

func (p *S_Fight_Monster_Blood_SYNC) Encode() []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.LittleEndian, p.TimeFrame)
	binary.Write(buffer, binary.LittleEndian, p.SyncRoleID)
	binary.Write(buffer, binary.LittleEndian, uint32(len(p.MonsterBlood)))
	for k, v := range p.MonsterBlood {
		binary.Write(buffer, binary.LittleEndian, k)
		buffer.Write(v.Encode())
	}
	return buffer.Bytes()
}

func (p *S_Fight_Monster_Blood_SYNC) Decode(buffer *bytes.Buffer) error {
	if buffer.Len() < 16 {
		return errors.New("message length error")
	}
	binary.Read(buffer, binary.LittleEndian, &p.TimeFrame)
	binary.Read(buffer, binary.LittleEndian, &p.SyncRoleID)
	var MonsterBloodLen uint32
	binary.Read(buffer, binary.LittleEndian, &MonsterBloodLen)
	p.MonsterBlood = make(map[int64]T_Monster_Blood_Struct, MonsterBloodLen)
	for i := uint32(0); i < MonsterBloodLen; i++ {
		var k int64
		var v T_Monster_Blood_Struct
		binary.Read(buffer, binary.LittleEndian, &k)
		if err := v.Decode(buffer); err != nil {
			return err
		}
		p.MonsterBlood[k] = v
	}
	return nil
}

type S_Fight_Hero_Attr_SYNC struct {
	RoleID      int64
	HeroAttrVec []T_Hero_Attr_Struct
}

func (p *S_Fight_Hero_Attr_SYNC) Encode() []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.LittleEndian, p.RoleID)
	binary.Write(buffer, binary.LittleEndian, uint32(len(p.HeroAttrVec)))
	for _, v := range p.HeroAttrVec {
		buffer.Write(v.Encode())
	}
	return buffer.Bytes()
}

func (p *S_Fight_Hero_Attr_SYNC) Decode(buffer *bytes.Buffer) error {
	if buffer.Len() < 12 {
		return errors.New("message length error")
	}
	binary.Read(buffer, binary.LittleEndian, &p.RoleID)
	var HeroAttrVecLen uint32
	binary.Read(buffer, binary.LittleEndian, &HeroAttrVecLen)
	p.HeroAttrVec = make([]T_Hero_Attr_Struct, HeroAttrVecLen)
	for i := uint32(0); i < HeroAttrVecLen; i++ {
		if err := p.HeroAttrVec[i].Decode(buffer); err != nil {
			return err
		}
	}
	return nil
}

type S_Fight_Update_Hero_SYNC struct {
	RoleID     int64
	UpdateHero T_Update_Hero_Struct
}

func (p *S_Fight_Update_Hero_SYNC) Encode() []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.LittleEndian, p.RoleID)
	buffer.Write(p.UpdateHero.Encode())
	return buffer.Bytes()
}

func (p *S_Fight_Update_Hero_SYNC) Decode(buffer *bytes.Buffer) error {
	if buffer.Len() < 32 {
		return errors.New("message length error")
	}
	binary.Read(buffer, binary.LittleEndian, &p.RoleID)
	if err := p.UpdateHero.Decode(buffer); err != nil {
		return err
	}
	return nil
}

type S_Fight_Operate_Equip_SYNC struct {
	Errorcode int32
	OperateType int32
	CarId int32
	Target int32
	EquipItems map[int64]EquipItem
}

func (p *S_Fight_Operate_Equip_SYNC) Encode() []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.LittleEndian, p.Errorcode)
	binary.Write(buffer, binary.LittleEndian, p.OperateType)
	binary.Write(buffer, binary.LittleEndian, p.CarId)
	binary.Write(buffer, binary.LittleEndian, p.Target)
	binary.Write(buffer, binary.LittleEndian, uint32(len(p.EquipItems)))
	for k, v := range p.EquipItems{
		binary.Write(buffer, binary.LittleEndian, k)
		buffer.Write(v.Encode())
	}
	return buffer.Bytes()
}

func (p *S_Fight_Operate_Equip_SYNC) Decode(buffer *bytes.Buffer) error {
	if buffer.Len() < 20 {
		return errors.New("message length error")
	}
	binary.Read(buffer, binary.LittleEndian, &p.Errorcode)
	binary.Read(buffer, binary.LittleEndian, &p.OperateType)
	binary.Read(buffer, binary.LittleEndian, &p.CarId)
	binary.Read(buffer, binary.LittleEndian, &p.Target)
	var EquipItemsLen uint32
	binary.Read(buffer, binary.LittleEndian, &EquipItemsLen)
	if uint32(buffer.Len()) < EquipItemsLen * 16 {
		return errors.New("message length error")
	}
	p.EquipItems = make(map[int64]EquipItem, EquipItemsLen)
	for i := uint32(0); i < EquipItemsLen; i++ {
		var k int64
		var v EquipItem
		binary.Read(buffer, binary.LittleEndian, &k)
		if err := v.Decode(buffer); err != nil {
			return err
		}
		p.EquipItems[k] = v
	}
	return nil
}

// 服务器返回战斗结果错误码（多人）
type S_Fight_Report_Result_To_Fight struct {
	Errorcode int32
}

func (p *S_Fight_Report_Result_To_Fight) Encode() []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.LittleEndian, p.Errorcode)
	return buffer.Bytes()
}

func (p *S_Fight_Report_Result_To_Fight) Decode(buffer *bytes.Buffer) error {
	if buffer.Len() < 4 {
		return errors.New("message length error")
	}
	binary.Read(buffer, binary.LittleEndian, &p.Errorcode)
	return nil
}
