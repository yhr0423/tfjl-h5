package protocols

import (
	"bytes"
	"encoding/binary"
	"errors"
)

type T_RoleFightTypeAbstract struct {
	MaxRound     int32
	WinNum       int32
	LostNum      int32
	SeriesWinNum int32
}

func (p *T_RoleFightTypeAbstract) Encode() []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.LittleEndian, p.MaxRound)
	binary.Write(buffer, binary.LittleEndian, p.WinNum)
	binary.Write(buffer, binary.LittleEndian, p.LostNum)
	binary.Write(buffer, binary.LittleEndian, p.SeriesWinNum)
	return buffer.Bytes()
}

func (p *T_RoleFightTypeAbstract) Decode(buffer *bytes.Buffer) error {
	if buffer.Len() < 16 {
		return errors.New("message length error")
	}
	binary.Read(buffer, binary.LittleEndian, &p.MaxRound)
	binary.Read(buffer, binary.LittleEndian, &p.WinNum)
	binary.Read(buffer, binary.LittleEndian, &p.LostNum)
	binary.Read(buffer, binary.LittleEndian, &p.SeriesWinNum)
	return nil
}

type T_Attr struct {
	Type  int32
	Value int32
}

func (p *T_Attr) Encode() []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.LittleEndian, p.Type)
	binary.Write(buffer, binary.LittleEndian, p.Value)
	return buffer.Bytes()
}

func (p *T_Attr) Decode(buffer *bytes.Buffer) error {
	if buffer.Len() < 8 {
		return errors.New("message length error")
	}
	binary.Read(buffer, binary.LittleEndian, &p.Type)
	binary.Read(buffer, binary.LittleEndian, &p.Value)
	return nil
}

type T_HeroAbstract struct {
	HeroUUID  int64
	HeroID    int32
	HeroLevel int32
	Attr      []T_Attr
	SkinID    int32
}

func (p *T_HeroAbstract) Encode() []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.LittleEndian, p.HeroUUID)
	binary.Write(buffer, binary.LittleEndian, p.HeroID)
	binary.Write(buffer, binary.LittleEndian, p.HeroLevel)
	binary.Write(buffer, binary.LittleEndian, uint32(len(p.Attr)))
	for _, v := range p.Attr {
		buffer.Write(v.Encode())
	}
	binary.Write(buffer, binary.LittleEndian, p.SkinID)
	return buffer.Bytes()
}

func (p *T_HeroAbstract) Decode(buffer *bytes.Buffer) error {
	if buffer.Len() < 20 {
		return errors.New("message length error")
	}
	binary.Read(buffer, binary.LittleEndian, &p.HeroUUID)
	binary.Read(buffer, binary.LittleEndian, &p.HeroID)
	binary.Read(buffer, binary.LittleEndian, &p.HeroLevel)
	var AttrLen uint32
	binary.Read(buffer, binary.LittleEndian, &AttrLen)
	if uint32(buffer.Len()) < AttrLen*8 {
		return errors.New("message length error")
	}
	p.Attr = make([]T_Attr, AttrLen)
	for i := uint32(0); i < AttrLen; i++ {
		p.Attr[i].Decode(buffer)
	}
	if buffer.Len() < 4 {
		return errors.New("message length error")
	}
	binary.Read(buffer, binary.LittleEndian, &p.SkinID)
	return nil
}

type T_ExpressionAbstract struct {
	ExpressionUUID int64
	ExpressionID   int32
}

func (p *T_ExpressionAbstract) Encode() []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.LittleEndian, p.ExpressionUUID)
	binary.Write(buffer, binary.LittleEndian, p.ExpressionID)
	return buffer.Bytes()
}

func (p *T_ExpressionAbstract) Decode(buffer *bytes.Buffer) error {
	if buffer.Len() < 12 {
		return errors.New("message length error")
	}
	binary.Read(buffer, binary.LittleEndian, &p.ExpressionUUID)
	binary.Read(buffer, binary.LittleEndian, &p.ExpressionID)
	return nil
}

type T_RuneAbstract struct {
	ItemID int32
}

func (p *T_RuneAbstract) Encode() []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.LittleEndian, p.ItemID)
	return buffer.Bytes()
}

func (p *T_RuneAbstract) Decode(buffer *bytes.Buffer) error {
	if buffer.Len() < 4 {
		return errors.New("message length error")
	}
	binary.Read(buffer, binary.LittleEndian, &p.ItemID)
	return nil
}

type T_FinalRuneAbstract struct {
	Level int32
}

func (p *T_FinalRuneAbstract) Encode() []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.LittleEndian, p.Level)
	return buffer.Bytes()
}

func (p *T_FinalRuneAbstract) Decode(buffer *bytes.Buffer) error {
	if buffer.Len() < 4 {
		return errors.New("message length error")
	}
	binary.Read(buffer, binary.LittleEndian, &p.Level)
	return nil
}

type T_Fight_SeasonData struct {
	SeasonID  int32
	BeginTime int32
}

func (p *T_Fight_SeasonData) Encode() []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.LittleEndian, p.SeasonID)
	binary.Write(buffer, binary.LittleEndian, p.BeginTime)
	return buffer.Bytes()
}

func (p *T_Fight_SeasonData) Decode(buffer *bytes.Buffer) error {
	if buffer.Len() < 8 {
		return errors.New("message length error")
	}
	binary.Read(buffer, binary.LittleEndian, &p.SeasonID)
	binary.Read(buffer, binary.LittleEndian, &p.BeginTime)
	return nil
}

type T_Fight_Activity_Cooperation struct {
	ScriptID int32
	BuffId   []int32
}

func (p *T_Fight_Activity_Cooperation) Encode() []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.LittleEndian, p.ScriptID)
	binary.Write(buffer, binary.LittleEndian, uint32(len(p.BuffId)))
	for _, v := range p.BuffId {
		binary.Write(buffer, binary.LittleEndian, v)
	}
	return buffer.Bytes()
}

func (p *T_Fight_Activity_Cooperation) Decode(buffer *bytes.Buffer) error {
	if buffer.Len() < 8 {
		return errors.New("message length error")
	}
	binary.Read(buffer, binary.LittleEndian, &p.ScriptID)
	var BuffIdLen uint32
	binary.Read(buffer, binary.LittleEndian, &BuffIdLen)
	if uint32(buffer.Len()) < BuffIdLen*4 {
		return errors.New("message length error")
	}
	p.BuffId = make([]int32, BuffIdLen)
	for i := uint32(0); i < BuffIdLen; i++ {
		binary.Read(buffer, binary.LittleEndian, &p.BuffId[i])
	}
	return nil
}

type T_Fight_Extra_Data struct {
	CooperationData map[int32]T_Fight_Activity_Cooperation
}

func (p *T_Fight_Extra_Data) Encode() []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.LittleEndian, uint32(len(p.CooperationData)))
	for k, v := range p.CooperationData {
		binary.Write(buffer, binary.LittleEndian, k)
		buffer.Write(v.Encode())
	}
	return buffer.Bytes()
}

func (p *T_Fight_Extra_Data) Decode(buffer *bytes.Buffer) error {
	if buffer.Len() < 4 {
		return errors.New("message length error")
	}
	var CooperationDataLen uint32
	binary.Read(buffer, binary.LittleEndian, &CooperationDataLen)
	p.CooperationData = make(map[int32]T_Fight_Activity_Cooperation, CooperationDataLen)
	for i := uint32(0); i < CooperationDataLen; i++ {
		var CooperationDataKey int32
		binary.Read(buffer, binary.LittleEndian, &CooperationDataKey)
		var CooperationDataValue T_Fight_Activity_Cooperation
		if err := CooperationDataValue.Decode(buffer); err != nil {
			return err
		}
		p.CooperationData[CooperationDataKey] = CooperationDataValue
	}
	return nil
}

type T_RoleAbstract struct {
	RoleID           int64
	ShowID           string
	BRobot           bool
	Aiid             int32
	LoginSDKType     int32
	LoginClientType  int32
	LoginSystematic  int32
	NickName         string
	HeadID           int32
	HeadUrl          string
	HeadFrameID      int32
	MapSkinID        int32
	AutoExpressionID int32
	Score            int32
	MaxScore         int32
	MaxSeasonScore   int32
	EncourageNum     int32
	ExtraCooperation int32
	Level            int32
	LevelExp         int32
	SociatyID        string
	SociatyName      string
	SociatyFlag      int32
	SociatyJob       int32
	SociatyLevel     int32
	FightType        map[int32]T_RoleFightTypeAbstract
	Heros            map[int32]T_HeroAbstract
	Expressions      map[int32]T_ExpressionAbstract
	Runes            map[int32]T_RuneAbstract
	Finalrunes       map[int32]T_FinalRuneAbstract
	FightSeasonData  map[int32]T_Fight_SeasonData
	PetId            int32
	QQHallVipInfo    []int32
	QQMiniGiftGetMap map[int32]int32
}

func (p *T_RoleAbstract) Encode() []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.LittleEndian, p.RoleID)
	binary.Write(buffer, binary.LittleEndian, uint32(len(p.ShowID)))
	buffer.Write([]byte(p.ShowID))
	binary.Write(buffer, binary.LittleEndian, p.BRobot)
	binary.Write(buffer, binary.LittleEndian, p.Aiid)
	binary.Write(buffer, binary.LittleEndian, p.LoginSDKType)
	binary.Write(buffer, binary.LittleEndian, p.LoginClientType)
	binary.Write(buffer, binary.LittleEndian, p.LoginSystematic)
	binary.Write(buffer, binary.LittleEndian, uint32(len(p.NickName)))
	buffer.Write([]byte(p.NickName))
	binary.Write(buffer, binary.LittleEndian, p.HeadID)
	binary.Write(buffer, binary.LittleEndian, uint32(len(p.HeadUrl)))
	buffer.Write([]byte(p.HeadUrl))
	binary.Write(buffer, binary.LittleEndian, p.HeadFrameID)
	binary.Write(buffer, binary.LittleEndian, p.MapSkinID)
	binary.Write(buffer, binary.LittleEndian, p.AutoExpressionID)
	binary.Write(buffer, binary.LittleEndian, p.Score)
	binary.Write(buffer, binary.LittleEndian, p.MaxScore)
	binary.Write(buffer, binary.LittleEndian, p.MaxSeasonScore)
	binary.Write(buffer, binary.LittleEndian, p.EncourageNum)
	binary.Write(buffer, binary.LittleEndian, p.ExtraCooperation)
	binary.Write(buffer, binary.LittleEndian, p.Level)
	binary.Write(buffer, binary.LittleEndian, p.LevelExp)
	binary.Write(buffer, binary.LittleEndian, uint32(len(p.SociatyID)))
	buffer.Write([]byte(p.SociatyID))
	binary.Write(buffer, binary.LittleEndian, uint32(len(p.SociatyName)))
	buffer.Write([]byte(p.SociatyName))
	binary.Write(buffer, binary.LittleEndian, p.SociatyFlag)
	binary.Write(buffer, binary.LittleEndian, p.SociatyJob)
	binary.Write(buffer, binary.LittleEndian, p.SociatyLevel)
	binary.Write(buffer, binary.LittleEndian, uint32(len(p.FightType)))
	for k, v := range p.FightType {
		binary.Write(buffer, binary.LittleEndian, k)
		buffer.Write(v.Encode())
	}
	binary.Write(buffer, binary.LittleEndian, uint32(len(p.Heros)))
	for k, v := range p.Heros {
		binary.Write(buffer, binary.LittleEndian, k)
		buffer.Write(v.Encode())
	}
	binary.Write(buffer, binary.LittleEndian, uint32(len(p.Expressions)))
	for k, v := range p.Expressions {
		binary.Write(buffer, binary.LittleEndian, k)
		buffer.Write(v.Encode())
	}
	binary.Write(buffer, binary.LittleEndian, uint32(len(p.Runes)))
	for k, v := range p.Runes {
		binary.Write(buffer, binary.LittleEndian, k)
		buffer.Write(v.Encode())
	}
	binary.Write(buffer, binary.LittleEndian, uint32(len(p.Finalrunes)))
	for k, v := range p.Finalrunes {
		binary.Write(buffer, binary.LittleEndian, k)
		buffer.Write(v.Encode())
	}
	binary.Write(buffer, binary.LittleEndian, uint32(len(p.FightSeasonData)))
	for k, v := range p.FightSeasonData {
		binary.Write(buffer, binary.LittleEndian, k)
		buffer.Write(v.Encode())
	}
	binary.Write(buffer, binary.LittleEndian, p.PetId)
	binary.Write(buffer, binary.LittleEndian, uint32(len(p.QQHallVipInfo)))
	for _, v := range p.QQHallVipInfo {
		binary.Write(buffer, binary.LittleEndian, v)
	}
	binary.Write(buffer, binary.LittleEndian, uint32(len(p.QQMiniGiftGetMap)))
	for k, v := range p.QQMiniGiftGetMap {
		binary.Write(buffer, binary.LittleEndian, k)
		binary.Write(buffer, binary.LittleEndian, v)
	}
	return buffer.Bytes()
}

func (p *T_RoleAbstract) Decode(buffer *bytes.Buffer) error {
	if buffer.Len() < 12 {
		return errors.New("message length error")
	}
	binary.Read(buffer, binary.LittleEndian, &p.RoleID)
	var ShowIDLen uint32
	binary.Read(buffer, binary.LittleEndian, &ShowIDLen)
	if uint32(buffer.Len()) < ShowIDLen {
		return errors.New("message length error")
	}
	p.ShowID = string(buffer.Next(int(ShowIDLen)))
	if buffer.Len() < 21 {
		return errors.New("message length error")
	}
	binary.Read(buffer, binary.LittleEndian, &p.BRobot)
	binary.Read(buffer, binary.LittleEndian, &p.Aiid)
	binary.Read(buffer, binary.LittleEndian, &p.LoginSDKType)
	binary.Read(buffer, binary.LittleEndian, &p.LoginClientType)
	binary.Read(buffer, binary.LittleEndian, &p.LoginSystematic)
	var NickNameLen uint32
	binary.Read(buffer, binary.LittleEndian, &NickNameLen)
	if uint32(buffer.Len()) < NickNameLen {
		return errors.New("message length error")
	}
	p.NickName = string(buffer.Next(int(NickNameLen)))
	if buffer.Len() < 8 {
		return errors.New("message length error")
	}
	binary.Read(buffer, binary.LittleEndian, &p.HeadID)
	var HeadUrlLen uint32
	binary.Read(buffer, binary.LittleEndian, &HeadUrlLen)
	if uint32(buffer.Len()) < HeadUrlLen {
		return errors.New("message length error")
	}
	p.HeadUrl = string(buffer.Next(int(HeadUrlLen)))
	if buffer.Len() < 44 {
		return errors.New("message length error")
	}
	binary.Read(buffer, binary.LittleEndian, &p.HeadFrameID)
	binary.Read(buffer, binary.LittleEndian, &p.MapSkinID)
	binary.Read(buffer, binary.LittleEndian, &p.AutoExpressionID)
	binary.Read(buffer, binary.LittleEndian, &p.Score)
	binary.Read(buffer, binary.LittleEndian, &p.MaxScore)
	binary.Read(buffer, binary.LittleEndian, &p.MaxSeasonScore)
	binary.Read(buffer, binary.LittleEndian, &p.EncourageNum)
	binary.Read(buffer, binary.LittleEndian, &p.ExtraCooperation)
	binary.Read(buffer, binary.LittleEndian, &p.Level)
	binary.Read(buffer, binary.LittleEndian, &p.LevelExp)
	var SociatyIDLen uint32
	binary.Read(buffer, binary.LittleEndian, &SociatyIDLen)
	if uint32(buffer.Len()) < SociatyIDLen {
		return errors.New("message length error")
	}
	p.SociatyID = string(buffer.Next(int(SociatyIDLen)))
	if buffer.Len() < 4 {
		return errors.New("message length error")
	}
	var SociatyNameLen uint32
	binary.Read(buffer, binary.LittleEndian, &SociatyNameLen)
	if uint32(buffer.Len()) < SociatyNameLen {
		return errors.New("message length error")
	}
	p.SociatyName = string(buffer.Next(int(SociatyNameLen)))
	if buffer.Len() < 16 {
		return errors.New("message length error")
	}
	binary.Read(buffer, binary.LittleEndian, &p.SociatyFlag)
	binary.Read(buffer, binary.LittleEndian, &p.SociatyJob)
	binary.Read(buffer, binary.LittleEndian, &p.SociatyLevel)
	var FightTypeLen uint32
	binary.Read(buffer, binary.LittleEndian, &FightTypeLen)
	if uint32(buffer.Len()) < FightTypeLen*20 {
		return errors.New("message length error")
	}
	p.FightType = make(map[int32]T_RoleFightTypeAbstract, FightTypeLen)
	for i := uint32(0); i < FightTypeLen; i++ {
		var key int32
		binary.Read(buffer, binary.LittleEndian, &key)
		var value T_RoleFightTypeAbstract
		value.Decode(buffer)
		p.FightType[key] = value
	}
	if buffer.Len() < 4 {
		return errors.New("message length error")
	}
	var HerosLen uint32
	binary.Read(buffer, binary.LittleEndian, &HerosLen)
	p.Heros = make(map[int32]T_HeroAbstract, HerosLen)
	for i := uint32(0); i < HerosLen; i++ {
		var key int32
		binary.Read(buffer, binary.LittleEndian, &key)
		var value T_HeroAbstract
		value.Decode(buffer)
		p.Heros[key] = value
	}
	if buffer.Len() < 4 {
		return errors.New("message length error")
	}
	var ExpressionsLen uint32
	binary.Read(buffer, binary.LittleEndian, &ExpressionsLen)
	if uint32(buffer.Len()) < ExpressionsLen*16 {
		return errors.New("message length error")
	}
	p.Expressions = make(map[int32]T_ExpressionAbstract, ExpressionsLen)
	for i := uint32(0); i < ExpressionsLen; i++ {
		var key int32
		binary.Read(buffer, binary.LittleEndian, &key)
		var value T_ExpressionAbstract
		value.Decode(buffer)
		p.Expressions[key] = value
	}
	if buffer.Len() < 4 {
		return errors.New("message length error")
	}
	var RunesLen uint32
	binary.Read(buffer, binary.LittleEndian, &RunesLen)
	if uint32(buffer.Len()) < RunesLen*8 {
		return errors.New("message length error")
	}
	p.Runes = make(map[int32]T_RuneAbstract, RunesLen)
	for i := uint32(0); i < RunesLen; i++ {
		var key int32
		binary.Read(buffer, binary.LittleEndian, &key)
		var value T_RuneAbstract
		value.Decode(buffer)
		p.Runes[key] = value
	}
	if buffer.Len() < 4 {
		return errors.New("message length error")
	}
	var FinalrunesLen uint32
	binary.Read(buffer, binary.LittleEndian, &FinalrunesLen)
	if uint32(buffer.Len()) < FinalrunesLen*8 {
		return errors.New("message length error")
	}
	p.Finalrunes = make(map[int32]T_FinalRuneAbstract, FinalrunesLen)
	for i := uint32(0); i < FinalrunesLen; i++ {
		var key int32
		binary.Read(buffer, binary.LittleEndian, &key)
		var value T_FinalRuneAbstract
		value.Decode(buffer)
		p.Finalrunes[key] = value
	}
	if buffer.Len() < 4 {
		return errors.New("message length error")
	}
	var FightSeasonDataLen uint32
	binary.Read(buffer, binary.LittleEndian, &FightSeasonDataLen)
	if uint32(buffer.Len()) < FightSeasonDataLen*12 {
		return errors.New("message length error")
	}
	p.FightSeasonData = make(map[int32]T_Fight_SeasonData, FightSeasonDataLen)
	for i := uint32(0); i < FightSeasonDataLen; i++ {
		var key int32
		binary.Read(buffer, binary.LittleEndian, &key)
		var value T_Fight_SeasonData
		value.Decode(buffer)
		p.FightSeasonData[key] = value
	}
	if buffer.Len() < 8 {
		return errors.New("message length error")
	}
	binary.Read(buffer, binary.LittleEndian, &p.PetId)
	var QQHallVipInfoLen uint32
	binary.Read(buffer, binary.LittleEndian, &QQHallVipInfoLen)
	if uint32(buffer.Len()) < QQHallVipInfoLen*4 {
		return errors.New("message length error")
	}
	p.QQHallVipInfo = make([]int32, QQHallVipInfoLen)
	for i := uint32(0); i < QQHallVipInfoLen; i++ {
		binary.Read(buffer, binary.LittleEndian, &p.QQHallVipInfo[i])
	}
	if buffer.Len() < 4 {
		return errors.New("message length error")
	}
	var QQMiniGiftGetMapLen uint32
	binary.Read(buffer, binary.LittleEndian, &QQMiniGiftGetMapLen)
	if uint32(buffer.Len()) < QQMiniGiftGetMapLen*8 {
		return errors.New("message length error")
	}
	p.QQMiniGiftGetMap = make(map[int32]int32, QQMiniGiftGetMapLen)
	for i := uint32(0); i < QQMiniGiftGetMapLen; i++ {
		var key int32
		binary.Read(buffer, binary.LittleEndian, &key)
		var value int32
		binary.Read(buffer, binary.LittleEndian, &value)
		p.QQMiniGiftGetMap[key] = value
	}
	return nil
}

type T_Chat_Data struct {
}

func (p *T_Chat_Data) Encode() []byte {
	buffer := new(bytes.Buffer)
	return buffer.Bytes()
}

func (p *T_Chat_Data) Decode(buffer *bytes.Buffer) error {
	return nil
}

type T_Fight_Role_Detail_Info_HeroMixLevelData struct {
	BuildNum int32
	MixNum   int32
}

func (p *T_Fight_Role_Detail_Info_HeroMixLevelData) Encode() []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.LittleEndian, p.BuildNum)
	binary.Write(buffer, binary.LittleEndian, p.MixNum)
	return buffer.Bytes()
}

func (p *T_Fight_Role_Detail_Info_HeroMixLevelData) Decode(buffer *bytes.Buffer) error {
	if buffer.Len() < 8 {
		return errors.New("message length error")
	}
	binary.Read(buffer, binary.LittleEndian, &p.BuildNum)
	binary.Read(buffer, binary.LittleEndian, &p.MixNum)
	return nil
}

type T_Fight_Role_Detail_Info_Heros struct {
	HeroID        int32
	MaxFightLevel int32
	MixLevel      map[int32]T_Fight_Role_Detail_Info_HeroMixLevelData
	DamageTotal   int64
	MaxDamage     int32
	CardLevel     int32
	HealTotal     int64
}

func (p *T_Fight_Role_Detail_Info_Heros) Encode() []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.LittleEndian, p.HeroID)
	binary.Write(buffer, binary.LittleEndian, p.MaxFightLevel)
	binary.Write(buffer, binary.LittleEndian, uint32(len(p.MixLevel)))
	for key, value := range p.MixLevel {
		binary.Write(buffer, binary.LittleEndian, key)
		buffer.Write(value.Encode())
	}
	binary.Write(buffer, binary.LittleEndian, p.DamageTotal)
	binary.Write(buffer, binary.LittleEndian, p.MaxDamage)
	binary.Write(buffer, binary.LittleEndian, p.CardLevel)
	binary.Write(buffer, binary.LittleEndian, p.HealTotal)
	return buffer.Bytes()
}

func (p *T_Fight_Role_Detail_Info_Heros) Decode(buffer *bytes.Buffer) error {
	if buffer.Len() < 12 {
		return errors.New("message length error")
	}
	binary.Read(buffer, binary.LittleEndian, &p.HeroID)
	binary.Read(buffer, binary.LittleEndian, &p.MaxFightLevel)
	var MixLevelLen uint32
	binary.Read(buffer, binary.LittleEndian, &MixLevelLen)
	if uint32(buffer.Len()) < MixLevelLen*12 {
		return errors.New("message length error")
	}
	p.MixLevel = make(map[int32]T_Fight_Role_Detail_Info_HeroMixLevelData, MixLevelLen)
	for i := uint32(0); i < MixLevelLen; i++ {
		var key int32
		binary.Read(buffer, binary.LittleEndian, &key)
		var value T_Fight_Role_Detail_Info_HeroMixLevelData
		value.Decode(buffer)
		p.MixLevel[key] = value
	}
	if buffer.Len() < 24 {
		return errors.New("message length error")
	}
	binary.Read(buffer, binary.LittleEndian, &p.DamageTotal)
	binary.Read(buffer, binary.LittleEndian, &p.MaxDamage)
	binary.Read(buffer, binary.LittleEndian, &p.CardLevel)
	binary.Read(buffer, binary.LittleEndian, &p.HealTotal)
	return nil
}

type T_Fight_Role_Detail_Info_Monsters struct {
	DeathNum int32
}

func (p *T_Fight_Role_Detail_Info_Monsters) Encode() []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.LittleEndian, p.DeathNum)
	return buffer.Bytes()
}

func (p *T_Fight_Role_Detail_Info_Monsters) Decode(buffer *bytes.Buffer) error {
	if buffer.Len() < 4 {
		return errors.New("message length error")
	}
	binary.Read(buffer, binary.LittleEndian, &p.DeathNum)
	return nil
}

type T_Fight_Statistic_Team_Hero_Info struct {
	HeroID    int32
	Level     int32
	StarLevel int32
	SkinID    int32
}

func (p *T_Fight_Statistic_Team_Hero_Info) Encode() []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.LittleEndian, p.HeroID)
	binary.Write(buffer, binary.LittleEndian, p.Level)
	binary.Write(buffer, binary.LittleEndian, p.StarLevel)
	binary.Write(buffer, binary.LittleEndian, p.SkinID)
	return buffer.Bytes()
}

func (p *T_Fight_Statistic_Team_Hero_Info) Decode(buffer *bytes.Buffer) error {
	if buffer.Len() < 16 {
		return errors.New("message length error")
	}
	binary.Read(buffer, binary.LittleEndian, &p.HeroID)
	binary.Read(buffer, binary.LittleEndian, &p.Level)
	binary.Read(buffer, binary.LittleEndian, &p.StarLevel)
	binary.Read(buffer, binary.LittleEndian, &p.SkinID)
	return nil
}

type T_Fight_Statistic_info struct {
	TotalDamage   int64
	TotalHeal     int64
	MaxDamage     int64
	HeroID        int32
	CumulativeSP  int32
	CumulativeRef int32
	Team          []T_Fight_Statistic_Team_Hero_Info
	CarLevel      int32
	OtherInfo     string
}

func (p *T_Fight_Statistic_info) Encode() []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.LittleEndian, p.TotalDamage)
	binary.Write(buffer, binary.LittleEndian, p.TotalHeal)
	binary.Write(buffer, binary.LittleEndian, p.MaxDamage)
	binary.Write(buffer, binary.LittleEndian, p.HeroID)
	binary.Write(buffer, binary.LittleEndian, p.CumulativeSP)
	binary.Write(buffer, binary.LittleEndian, p.CumulativeRef)
	binary.Write(buffer, binary.LittleEndian, uint32(len(p.Team)))
	for _, value := range p.Team {
		buffer.Write(value.Encode())
	}
	binary.Write(buffer, binary.LittleEndian, p.CarLevel)
	binary.Write(buffer, binary.LittleEndian, uint32(len(p.OtherInfo)))
	buffer.Write([]byte(p.OtherInfo))
	return buffer.Bytes()
}

func (p *T_Fight_Statistic_info) Decode(buffer *bytes.Buffer) error {
	if buffer.Len() < 40 {
		return errors.New("message length error")
	}
	binary.Read(buffer, binary.LittleEndian, &p.TotalDamage)
	binary.Read(buffer, binary.LittleEndian, &p.TotalHeal)
	binary.Read(buffer, binary.LittleEndian, &p.MaxDamage)
	binary.Read(buffer, binary.LittleEndian, &p.HeroID)
	binary.Read(buffer, binary.LittleEndian, &p.CumulativeSP)
	binary.Read(buffer, binary.LittleEndian, &p.CumulativeRef)
	var TeamLen uint32
	binary.Read(buffer, binary.LittleEndian, &TeamLen)
	if uint32(buffer.Len()) < TeamLen*16 {
		return errors.New("message length error")
	}
	p.Team = make([]T_Fight_Statistic_Team_Hero_Info, TeamLen)
	for i := uint32(0); i < TeamLen; i++ {
		p.Team[i].Decode(buffer)
	}
	binary.Read(buffer, binary.LittleEndian, &p.CarLevel)
	var OtherInfoLen uint32
	binary.Read(buffer, binary.LittleEndian, &OtherInfoLen)
	if uint32(buffer.Len()) < OtherInfoLen {
		return errors.New("message length error")
	}
	p.OtherInfo = string(buffer.Next(int(OtherInfoLen)))
	return nil
}

type T_Fight_Role_Detail_Info struct {
	RoleID          int64
	Round           int32
	RoleAbstract    T_RoleAbstract
	MaxSp           int32
	CumulativeSp    int32
	LastMonsterTime int32
	Heros           map[int64]T_Fight_Role_Detail_Info_Heros
	TypeMonster     map[int32]T_Fight_Role_Detail_Info_Monsters
	StatisticInfo   T_Fight_Statistic_info
}

func (p *T_Fight_Role_Detail_Info) Encode() []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.LittleEndian, p.RoleID)
	binary.Write(buffer, binary.LittleEndian, p.Round)
	buffer.Write(p.RoleAbstract.Encode())
	binary.Write(buffer, binary.LittleEndian, p.MaxSp)
	binary.Write(buffer, binary.LittleEndian, p.CumulativeSp)
	binary.Write(buffer, binary.LittleEndian, p.LastMonsterTime)
	binary.Write(buffer, binary.LittleEndian, uint32(len(p.Heros)))
	for key, value := range p.Heros {
		binary.Write(buffer, binary.LittleEndian, key)
		buffer.Write(value.Encode())
	}
	binary.Write(buffer, binary.LittleEndian, uint32(len(p.TypeMonster)))
	for key, value := range p.TypeMonster {
		binary.Write(buffer, binary.LittleEndian, key)
		buffer.Write(value.Encode())
	}
	buffer.Write(p.StatisticInfo.Encode())
	return buffer.Bytes()
}

func (p *T_Fight_Role_Detail_Info) Decode(buffer *bytes.Buffer) error {
	if buffer.Len() < 12 {
		return errors.New("message length error")
	}
	binary.Read(buffer, binary.LittleEndian, &p.RoleID)
	binary.Read(buffer, binary.LittleEndian, &p.Round)
	if err := p.RoleAbstract.Decode(buffer); err != nil {
		return err
	}
	if buffer.Len() < 16 {
		return errors.New("message length error")
	}
	binary.Read(buffer, binary.LittleEndian, &p.MaxSp)
	binary.Read(buffer, binary.LittleEndian, &p.CumulativeSp)
	binary.Read(buffer, binary.LittleEndian, &p.LastMonsterTime)
	var HerosLen uint32
	binary.Read(buffer, binary.LittleEndian, &HerosLen)
	p.Heros = make(map[int64]T_Fight_Role_Detail_Info_Heros, HerosLen)
	for i := uint32(0); i < HerosLen; i++ {
		var key int64
		binary.Read(buffer, binary.LittleEndian, &key)
		var value T_Fight_Role_Detail_Info_Heros
		if err := value.Decode(buffer); err != nil {
			return err
		}
		p.Heros[key] = value
	}
	if buffer.Len() < 4 {
		return errors.New("message length error")
	}
	var TypeMonsterLen uint32
	binary.Read(buffer, binary.LittleEndian, &TypeMonsterLen)
	if uint32(buffer.Len()) < TypeMonsterLen*8 {
		return errors.New("message length error")
	}
	p.TypeMonster = make(map[int32]T_Fight_Role_Detail_Info_Monsters, TypeMonsterLen)
	for i := uint32(0); i < TypeMonsterLen; i++ {
		var key int32
		binary.Read(buffer, binary.LittleEndian, &key)
		var value T_Fight_Role_Detail_Info_Monsters
		if err := value.Decode(buffer); err != nil {
			return err
		}
		p.TypeMonster[key] = value
	}
	if err := p.StatisticInfo.Decode(buffer); err != nil {
		return err
	}
	return nil
}

type T_Fight_Detail_Info struct {
	IsWin         bool
	GiveupVsRobot int32
	FightType     int32
	FightModule   int32
	FightRoleInfo map[int64]T_Fight_Role_Detail_Info
	Params        []int32
}

func (p *T_Fight_Detail_Info) Encode() []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.LittleEndian, p.IsWin)
	binary.Write(buffer, binary.LittleEndian, p.GiveupVsRobot)
	binary.Write(buffer, binary.LittleEndian, p.FightType)
	binary.Write(buffer, binary.LittleEndian, p.FightModule)
	binary.Write(buffer, binary.LittleEndian, uint32(len(p.FightRoleInfo)))
	for key, value := range p.FightRoleInfo {
		binary.Write(buffer, binary.LittleEndian, key)
		buffer.Write(value.Encode())
	}
	binary.Write(buffer, binary.LittleEndian, uint32(len(p.Params)))
	for _, value := range p.Params {
		binary.Write(buffer, binary.LittleEndian, value)
	}
	return buffer.Bytes()
}

func (p *T_Fight_Detail_Info) Decode(buffer *bytes.Buffer) error {
	if buffer.Len() < 17 {
		return errors.New("message length error")
	}
	binary.Read(buffer, binary.LittleEndian, &p.IsWin)
	binary.Read(buffer, binary.LittleEndian, &p.GiveupVsRobot)
	binary.Read(buffer, binary.LittleEndian, &p.FightType)
	binary.Read(buffer, binary.LittleEndian, &p.FightModule)
	var FightRoleInfoLen uint32
	binary.Read(buffer, binary.LittleEndian, &FightRoleInfoLen)
	if uint32(buffer.Len()) < FightRoleInfoLen {
		return errors.New("map or array length too long")
	}
	p.FightRoleInfo = make(map[int64]T_Fight_Role_Detail_Info, FightRoleInfoLen)
	for i := uint32(0); i < FightRoleInfoLen; i++ {
		var key int64
		binary.Read(buffer, binary.LittleEndian, &key)
		var value T_Fight_Role_Detail_Info
		if err := value.Decode(buffer); err != nil {
			return err
		}
		p.FightRoleInfo[key] = value
	}
	if buffer.Len() < 4 {
		return errors.New("message length error")
	}
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

type T_Reward struct {
	DropType int32
	DropID   int32
	DropNum  int32
}

func (p *T_Reward) Encode() []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.LittleEndian, p.DropType)
	binary.Write(buffer, binary.LittleEndian, p.DropID)
	binary.Write(buffer, binary.LittleEndian, p.DropNum)
	return buffer.Bytes()
}

func (p *T_Reward) Decode(buffer *bytes.Buffer) error {
	if buffer.Len() < 12 {
		return errors.New("message length error")
	}
	binary.Read(buffer, binary.LittleEndian, &p.DropType)
	binary.Read(buffer, binary.LittleEndian, &p.DropID)
	binary.Read(buffer, binary.LittleEndian, &p.DropNum)
	return nil
}

type T_ProficiencyInfo struct {
	Proficiency int32
	WorldRank   int32
	LocalRank   int32
	IPName      string
	ShareData   map[int32]int32
}

func (p *T_ProficiencyInfo) Encode() []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.LittleEndian, p.Proficiency)
	binary.Write(buffer, binary.LittleEndian, p.WorldRank)
	binary.Write(buffer, binary.LittleEndian, p.LocalRank)
	binary.Write(buffer, binary.LittleEndian, uint32(len(p.IPName)))
	buffer.WriteString(p.IPName)
	binary.Write(buffer, binary.LittleEndian, uint32(len(p.ShareData)))
	for key, value := range p.ShareData {
		binary.Write(buffer, binary.LittleEndian, key)
		binary.Write(buffer, binary.LittleEndian, value)
	}
	return buffer.Bytes()
}

func (p *T_ProficiencyInfo) Decode(buffer *bytes.Buffer) error {
	if buffer.Len() < 16 {
		return errors.New("message length error")
	}
	binary.Read(buffer, binary.LittleEndian, &p.Proficiency)
	binary.Read(buffer, binary.LittleEndian, &p.WorldRank)
	binary.Read(buffer, binary.LittleEndian, &p.LocalRank)
	var IPNameLen uint32
	binary.Read(buffer, binary.LittleEndian, &IPNameLen)
	if uint32(buffer.Len()) < IPNameLen {
		return errors.New("message length error")
	}
	p.IPName = string(buffer.Next(int(IPNameLen)))
	if buffer.Len() < 4 {
		return errors.New("message length error")
	}
	var ShareDataLen uint32
	binary.Read(buffer, binary.LittleEndian, &ShareDataLen)
	if uint32(buffer.Len()) < ShareDataLen*8 {
		return errors.New("message length error")
	}
	p.ShareData = make(map[int32]int32, ShareDataLen)
	for i := uint32(0); i < ShareDataLen; i++ {
		var key int32
		binary.Read(buffer, binary.LittleEndian, &key)
		var value int32
		binary.Read(buffer, binary.LittleEndian, &value)
		p.ShareData[key] = value
	}
	return nil
}

type T_RoleProficiency struct {
	ProficiencyInfoMap map[int32]T_ProficiencyInfo
}

func (p *T_RoleProficiency) Encode() []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.LittleEndian, uint32(len(p.ProficiencyInfoMap)))
	for key, value := range p.ProficiencyInfoMap {
		binary.Write(buffer, binary.LittleEndian, key)
		buffer.Write(value.Encode())
	}
	return buffer.Bytes()
}

func (p *T_RoleProficiency) Decode(buffer *bytes.Buffer) error {
	if buffer.Len() < 4 {
		return errors.New("message length error")
	}
	var ProficiencyInfoMapLen uint32
	binary.Read(buffer, binary.LittleEndian, &ProficiencyInfoMapLen)
	p.ProficiencyInfoMap = make(map[int32]T_ProficiencyInfo, ProficiencyInfoMapLen)
	for i := uint32(0); i < ProficiencyInfoMapLen; i++ {
		var key int32
		binary.Read(buffer, binary.LittleEndian, &key)
		var value T_ProficiencyInfo
		if err := value.Decode(buffer); err != nil {
			return err
		}
		p.ProficiencyInfoMap[key] = value
	}
	return nil
}
