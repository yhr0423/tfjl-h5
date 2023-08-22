package protocols

import (
	"bytes"
	"encoding/binary"
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type T_Information_FightTypeData struct {
	MaxRound         int32
	WinNum           int32
	LostNum          int32
	TotalWinNum      int32
	TotalLostNum     int32
	SeriesWinNum     int32
	SeriesLostNum    int32
	WinLostResetNum  int32
	AdditionalDayNum int32
}

func (p *T_Information_FightTypeData) Encode() []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.LittleEndian, p.MaxRound)
	binary.Write(buffer, binary.LittleEndian, p.WinNum)
	binary.Write(buffer, binary.LittleEndian, p.LostNum)
	binary.Write(buffer, binary.LittleEndian, p.TotalWinNum)
	binary.Write(buffer, binary.LittleEndian, p.TotalLostNum)
	binary.Write(buffer, binary.LittleEndian, p.SeriesWinNum)
	binary.Write(buffer, binary.LittleEndian, p.SeriesLostNum)
	binary.Write(buffer, binary.LittleEndian, p.WinLostResetNum)
	binary.Write(buffer, binary.LittleEndian, p.AdditionalDayNum)
	return buffer.Bytes()
}

func (p *T_Information_FightTypeData) Decode(buffer *bytes.Buffer) error {
	if buffer.Len() < 36 {
		return errors.New("message length error")
	}
	binary.Read(buffer, binary.LittleEndian, &p.MaxRound)
	binary.Read(buffer, binary.LittleEndian, &p.WinNum)
	binary.Read(buffer, binary.LittleEndian, &p.LostNum)
	binary.Read(buffer, binary.LittleEndian, &p.TotalWinNum)
	binary.Read(buffer, binary.LittleEndian, &p.TotalLostNum)
	binary.Read(buffer, binary.LittleEndian, &p.SeriesWinNum)
	binary.Read(buffer, binary.LittleEndian, &p.SeriesLostNum)
	binary.Read(buffer, binary.LittleEndian, &p.WinLostResetNum)
	binary.Read(buffer, binary.LittleEndian, &p.AdditionalDayNum)
	return nil
}

type T_Information_FightData struct {
	TypeData map[int32]T_Information_FightTypeData
}

func (p *T_Information_FightData) Encode() []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.LittleEndian, uint32(len(p.TypeData)))
	for k, v := range p.TypeData {
		binary.Write(buffer, binary.LittleEndian, k)
		buffer.Write(v.Encode())
	}
	return buffer.Bytes()
}

func (p *T_Information_FightData) Decode(buffer *bytes.Buffer) error {
	if buffer.Len() < 4 {
		return errors.New("message length error")
	}
	var TypeDataLen uint32
	binary.Read(buffer, binary.LittleEndian, &TypeDataLen)
	if buffer.Len() < int(TypeDataLen)*40 {
		return errors.New("message length error")
	}
	p.TypeData = make(map[int32]T_Information_FightTypeData, TypeDataLen)
	for i := 0; i < int(TypeDataLen); i++ {
		var key int32
		binary.Read(buffer, binary.LittleEndian, &key)
		var value T_Information_FightTypeData
		value.Decode(buffer)
		p.TypeData[key] = value
	}
	return nil
}

type T_Task_Box_Data struct {
	BTake bool
}

func (p *T_Task_Box_Data) Encode() []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.LittleEndian, p.BTake)
	return buffer.Bytes()
}

func (p *T_Task_Box_Data) Decode(buffer *bytes.Buffer) error {
	if buffer.Len() < 1 {
		return errors.New("message length error")
	}
	binary.Read(buffer, binary.LittleEndian, &p.BTake)
	return nil
}

type T_Task_Group_Data struct {
	Boxs map[int32]T_Task_Box_Data
}

func (p *T_Task_Group_Data) Encode() []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.LittleEndian, uint32(len(p.Boxs)))
	for k, v := range p.Boxs {
		binary.Write(buffer, binary.LittleEndian, k)
		buffer.Write(v.Encode())
	}
	return buffer.Bytes()
}

func (p *T_Task_Group_Data) Decode(buffer *bytes.Buffer) error {
	if buffer.Len() < 4 {
		return errors.New("message length error")
	}
	var BoxsLen uint32
	binary.Read(buffer, binary.LittleEndian, &BoxsLen)
	if buffer.Len() < int(BoxsLen)*5 {
		return errors.New("message length error")
	}
	p.Boxs = make(map[int32]T_Task_Box_Data, BoxsLen)
	for i := 0; i < int(BoxsLen); i++ {
		var key int32
		binary.Read(buffer, binary.LittleEndian, &key)
		var value T_Task_Box_Data
		value.Decode(buffer)
		p.Boxs[key] = value
	}
	return nil
}

type T_Task_Rand_Task_Data struct {
	TaskID int32
}

func (p *T_Task_Rand_Task_Data) Encode() []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.LittleEndian, p.TaskID)
	return buffer.Bytes()
}

func (p *T_Task_Rand_Task_Data) Decode(buffer *bytes.Buffer) error {
	if buffer.Len() < 4 {
		return errors.New("message length error")
	}
	binary.Read(buffer, binary.LittleEndian, &p.TaskID)
	return nil
}

type T_Task_Rand_History_Task_Data struct {
}

func (p *T_Task_Rand_History_Task_Data) Encode() []byte {
	buffer := new(bytes.Buffer)
	return buffer.Bytes()
}

func (p *T_Task_Rand_History_Task_Data) Decode(buffer *bytes.Buffer) error {
	return nil
}

type T_Task_Rand_History_Data struct {
	Task map[int32]T_Task_Rand_History_Task_Data
}

func (p *T_Task_Rand_History_Data) Encode() []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.LittleEndian, uint32(len(p.Task)))
	for k, v := range p.Task {
		binary.Write(buffer, binary.LittleEndian, k)
		buffer.Write(v.Encode())
	}
	return buffer.Bytes()
}

func (p *T_Task_Rand_History_Data) Decode(buffer *bytes.Buffer) error {
	if buffer.Len() < 4 {
		return errors.New("message length error")
	}
	var TaskLen uint32
	binary.Read(buffer, binary.LittleEndian, &TaskLen)
	if buffer.Len() < int(TaskLen)*4 {
		return errors.New("message length error")
	}
	p.Task = make(map[int32]T_Task_Rand_History_Task_Data, TaskLen)
	for i := 0; i < int(TaskLen); i++ {
		var key int32
		binary.Read(buffer, binary.LittleEndian, &key)
		var value T_Task_Rand_History_Task_Data
		value.Decode(buffer)
		p.Task[key] = value
	}
	return nil
}

type T_Task_Extra_Data struct {
	TaskGroup          map[int32]T_Task_Group_Data
	RandTask           map[int32]T_Task_Rand_Task_Data
	ReplaceRandTaskNum int32
	RandTaskHistory    []T_Task_Rand_History_Data
}

func (p *T_Task_Extra_Data) Encode() []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.LittleEndian, uint32(len(p.TaskGroup)))
	for k, v := range p.TaskGroup {
		binary.Write(buffer, binary.LittleEndian, k)
		buffer.Write(v.Encode())
	}
	binary.Write(buffer, binary.LittleEndian, uint32(len(p.RandTask)))
	for k, v := range p.RandTask {
		binary.Write(buffer, binary.LittleEndian, k)
		buffer.Write(v.Encode())
	}
	binary.Write(buffer, binary.LittleEndian, p.ReplaceRandTaskNum)
	binary.Write(buffer, binary.LittleEndian, uint32(len(p.RandTaskHistory)))
	for _, v := range p.RandTaskHistory {
		buffer.Write(v.Encode())
	}
	return buffer.Bytes()
}

func (p *T_Task_Extra_Data) Decode(buffer *bytes.Buffer) error {
	if buffer.Len() < 4 {
		return errors.New("message length error")
	}
	var TaskGroupLen uint32
	binary.Read(buffer, binary.LittleEndian, &TaskGroupLen)
	p.TaskGroup = make(map[int32]T_Task_Group_Data, TaskGroupLen)
	for i := 0; i < int(TaskGroupLen); i++ {
		var key int32
		binary.Read(buffer, binary.LittleEndian, &key)
		var value T_Task_Group_Data
		value.Decode(buffer)
		p.TaskGroup[key] = value
	}
	if buffer.Len() < 4 {
		return errors.New("message length error")
	}
	var RandTaskLen uint32
	binary.Read(buffer, binary.LittleEndian, &RandTaskLen)
	if buffer.Len() < int(RandTaskLen)*8 {
		return errors.New("message length error")
	}
	p.RandTask = make(map[int32]T_Task_Rand_Task_Data, RandTaskLen)
	for i := 0; i < int(RandTaskLen); i++ {
		var key int32
		binary.Read(buffer, binary.LittleEndian, &key)
		var value T_Task_Rand_Task_Data
		value.Decode(buffer)
		p.RandTask[key] = value
	}
	if buffer.Len() < 8 {
		return errors.New("message length error")
	}
	binary.Read(buffer, binary.LittleEndian, &p.ReplaceRandTaskNum)
	var RandTaskHistoryLen uint32
	binary.Read(buffer, binary.LittleEndian, &RandTaskHistoryLen)
	p.RandTaskHistory = make([]T_Task_Rand_History_Data, RandTaskHistoryLen)
	for i := 0; i < int(RandTaskHistoryLen); i++ {
		p.RandTaskHistory[i].Decode(buffer)
	}
	return nil
}

type T_Role_SingleTask struct {
	ID_        primitive.ObjectID `bson:"_id"`
	RoleID     int64              `bson:"role_id"`
	TaskID     int32              `bson:"task_id"`
	TaskState  int32              `bson:"task_state"`
	TaskCount  int32              `bson:"task_count"`
	TaskCDTime int32              `bson:"task_cd_time"`
	ExtraState int32              `bson:"extra_state"`
}

func (p *T_Role_SingleTask) Encode() []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.LittleEndian, p.TaskID)
	binary.Write(buffer, binary.LittleEndian, p.TaskState)
	binary.Write(buffer, binary.LittleEndian, p.TaskCount)
	binary.Write(buffer, binary.LittleEndian, p.TaskCDTime)
	binary.Write(buffer, binary.LittleEndian, p.ExtraState)
	return buffer.Bytes()
}

func (p *T_Role_SingleTask) Decode(buffer *bytes.Buffer) error {
	if buffer.Len() < 20 {
		return errors.New("message length error")
	}
	binary.Read(buffer, binary.LittleEndian, &p.TaskID)
	binary.Read(buffer, binary.LittleEndian, &p.TaskState)
	binary.Read(buffer, binary.LittleEndian, &p.TaskCount)
	binary.Read(buffer, binary.LittleEndian, &p.TaskCDTime)
	binary.Read(buffer, binary.LittleEndian, &p.ExtraState)
	return nil
}

type T_TotalWatchADBox struct {
	WatchADRound      int32
	WatchADIndex      int32
	WatchADDay        int32
	WatchADNum        int32
	IsReceive         bool
	IsExtraRecboolive bool
}

func (p *T_TotalWatchADBox) Encode() []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.LittleEndian, p.WatchADRound)
	binary.Write(buffer, binary.LittleEndian, p.WatchADIndex)
	binary.Write(buffer, binary.LittleEndian, p.WatchADDay)
	binary.Write(buffer, binary.LittleEndian, p.WatchADNum)
	binary.Write(buffer, binary.LittleEndian, p.IsReceive)
	binary.Write(buffer, binary.LittleEndian, p.IsExtraRecboolive)
	return buffer.Bytes()
}

func (p *T_TotalWatchADBox) Decode(buffer *bytes.Buffer) error {
	if buffer.Len() < 18 {
		return errors.New("message length error")
	}
	binary.Read(buffer, binary.LittleEndian, &p.WatchADRound)
	binary.Read(buffer, binary.LittleEndian, &p.WatchADIndex)
	binary.Read(buffer, binary.LittleEndian, &p.WatchADDay)
	binary.Read(buffer, binary.LittleEndian, &p.WatchADNum)
	binary.Read(buffer, binary.LittleEndian, &p.IsReceive)
	binary.Read(buffer, binary.LittleEndian, &p.IsExtraRecboolive)
	return nil
}

type T_Game_Time struct {
	Year    int32
	Month   int32
	Day     int32
	Hour    int32
	Minnute int32
	Second  int32
	WeedDay int32
}

func (p *T_Game_Time) Encode() []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.LittleEndian, p.Year)
	binary.Write(buffer, binary.LittleEndian, p.Month)
	binary.Write(buffer, binary.LittleEndian, p.Day)
	binary.Write(buffer, binary.LittleEndian, p.Hour)
	binary.Write(buffer, binary.LittleEndian, p.Minnute)
	binary.Write(buffer, binary.LittleEndian, p.Second)
	binary.Write(buffer, binary.LittleEndian, p.WeedDay)
	return buffer.Bytes()
}

func (p *T_Game_Time) Decode(buffer *bytes.Buffer) error {
	if buffer.Len() < 28 {
		return errors.New("message length error")
	}
	binary.Read(buffer, binary.LittleEndian, &p.Year)
	binary.Read(buffer, binary.LittleEndian, &p.Month)
	binary.Read(buffer, binary.LittleEndian, &p.Day)
	binary.Read(buffer, binary.LittleEndian, &p.Hour)
	binary.Read(buffer, binary.LittleEndian, &p.Minnute)
	binary.Read(buffer, binary.LittleEndian, &p.Second)
	binary.Read(buffer, binary.LittleEndian, &p.WeedDay)
	return nil
}

type T_Information_Data struct {
	ID_                            primitive.ObjectID      `bson:"_id,omitempty"` // 主键
	RoleID                         int64                   `bson:"role_id"`       // 角色ID
	FightData                      T_Information_FightData `bson:"-"`
	EncourageNum                   int32                   `bson:"encourage_num"`
	EncourageDayNum                int32                   `bson:"encourage_day_num"`
	EncourageTime                  int32                   `bson:"encourage_time"`
	CooperationDayLeftNum          int32                   `bson:"cooperation_day_left_num"`
	CooperationDayWatchAdNum       int32                   `bson:"cooperation_day_watch_ad_num"`
	SupportBoxNum                  int32                   `bson:"support_box_num"`
	SupportBoxDayNum               int32                   `bson:"support_box_day_num"`
	SupportBoxLastTime             int32                   `bson:"support_box_last_time"`
	CooperationDayBuyNum           int32                   `bson:"cooperation_day_buy_num"`
	CooperationDayGiveNum          int32                   `bson:"cooperation_day_give_num"`
	CooperationDayNum              int32                   `bson:"cooperation_day_num"`
	FightLoserProtectDayNum        int32                   `bson:"fight_loser_protect_day_num"`
	FightLoserDecScore             int32                   `bson:"fight_loser_dec_score"`
	FightWinDayNum                 int32                   `bson:"fight_win_day_num"`
	FightWinExtra                  bool                    `bson:"fight_win_extra"`
	WeekCardAddCooperationExtraNum int32                   `bson:"week_card_add_cooperation_extra_num"`
	CooperationExtra               bool                    `bson:"cooperation_extra"`
	UsedCooperationLastExtraEnd    bool                    `bson:"used_cooperation_last_extra_end"`
	DayFirstWinNum                 int32                   `bson:"day_first_win_num"`
	DayFirstIsWin                  bool                    `bson:"day_first_is_win"`
	CooperADNum                    int32                   `bson:"cooper_ad_num"`
	CooperADRoundNum               int32                   `bson:"cooper_ad_round_num"`
	IsGetFirstRechargeReward       bool                    `bson:"is_get_first_recharge_reward"`
	FightLookADDoubleCount         int32                   `bson:"fight_look_ad_double_count"`
	CooperationDayHelpNum          int32                   `bson:"cooperation_day_help_num"`
	QQHallBlueVipVector            []int32                 `bson:"-"`
	QQHallRewardsVector            []string                `bson:"-"`
	HeroSkinMap                    map[int64]int32         `bson:"-"`
	SelfSelectMap                  map[int32]int32         `bson:"-"`
	CarLinkMap                     map[int32]int32         `bson:"-"`
}

func (p *T_Information_Data) Encode() []byte {
	buffer := new(bytes.Buffer)
	buffer.Write(p.FightData.Encode())
	binary.Write(buffer, binary.LittleEndian, p.EncourageNum)
	binary.Write(buffer, binary.LittleEndian, p.EncourageDayNum)
	binary.Write(buffer, binary.LittleEndian, p.EncourageTime)
	binary.Write(buffer, binary.LittleEndian, p.CooperationDayLeftNum)
	binary.Write(buffer, binary.LittleEndian, p.CooperationDayWatchAdNum)
	binary.Write(buffer, binary.LittleEndian, p.SupportBoxNum)
	binary.Write(buffer, binary.LittleEndian, p.SupportBoxDayNum)
	binary.Write(buffer, binary.LittleEndian, p.SupportBoxLastTime)
	binary.Write(buffer, binary.LittleEndian, p.CooperationDayBuyNum)
	binary.Write(buffer, binary.LittleEndian, p.CooperationDayGiveNum)
	binary.Write(buffer, binary.LittleEndian, p.CooperationDayNum)
	binary.Write(buffer, binary.LittleEndian, p.FightLoserProtectDayNum)
	binary.Write(buffer, binary.LittleEndian, p.FightLoserDecScore)
	binary.Write(buffer, binary.LittleEndian, p.FightWinDayNum)
	binary.Write(buffer, binary.LittleEndian, p.FightWinExtra)
	binary.Write(buffer, binary.LittleEndian, p.WeekCardAddCooperationExtraNum)
	binary.Write(buffer, binary.LittleEndian, p.CooperationExtra)
	binary.Write(buffer, binary.LittleEndian, p.UsedCooperationLastExtraEnd)
	binary.Write(buffer, binary.LittleEndian, p.DayFirstWinNum)
	binary.Write(buffer, binary.LittleEndian, p.DayFirstIsWin)
	binary.Write(buffer, binary.LittleEndian, p.CooperADNum)
	binary.Write(buffer, binary.LittleEndian, p.CooperADRoundNum)
	binary.Write(buffer, binary.LittleEndian, p.IsGetFirstRechargeReward)
	binary.Write(buffer, binary.LittleEndian, p.FightLookADDoubleCount)
	binary.Write(buffer, binary.LittleEndian, p.CooperationDayHelpNum)
	binary.Write(buffer, binary.LittleEndian, uint32(len(p.QQHallBlueVipVector)))
	for _, v := range p.QQHallBlueVipVector {
		binary.Write(buffer, binary.LittleEndian, v)
	}
	binary.Write(buffer, binary.LittleEndian, uint32(len(p.QQHallRewardsVector)))
	for _, v := range p.QQHallRewardsVector {
		binary.Write(buffer, binary.LittleEndian, []byte(v))
	}
	binary.Write(buffer, binary.LittleEndian, uint32(len(p.HeroSkinMap)))
	for k, v := range p.HeroSkinMap {
		binary.Write(buffer, binary.LittleEndian, k)
		binary.Write(buffer, binary.LittleEndian, v)
	}
	binary.Write(buffer, binary.LittleEndian, uint32(len(p.SelfSelectMap)))
	for k, v := range p.SelfSelectMap {
		binary.Write(buffer, binary.LittleEndian, k)
		binary.Write(buffer, binary.LittleEndian, v)
	}
	binary.Write(buffer, binary.LittleEndian, uint32(len(p.CarLinkMap)))
	for k, v := range p.CarLinkMap {
		binary.Write(buffer, binary.LittleEndian, k)
		binary.Write(buffer, binary.LittleEndian, v)
	}
	return buffer.Bytes()
}

func (p *T_Information_Data) Decode(buffer *bytes.Buffer) error {
	if buffer.Len() < 89 {
		return errors.New("message length error")
	}
	p.FightData.Decode(buffer)
	binary.Read(buffer, binary.LittleEndian, &p.EncourageNum)
	binary.Read(buffer, binary.LittleEndian, &p.EncourageDayNum)
	binary.Read(buffer, binary.LittleEndian, &p.EncourageTime)
	binary.Read(buffer, binary.LittleEndian, &p.CooperationDayLeftNum)
	binary.Read(buffer, binary.LittleEndian, &p.CooperationDayWatchAdNum)
	binary.Read(buffer, binary.LittleEndian, &p.SupportBoxNum)
	binary.Read(buffer, binary.LittleEndian, &p.SupportBoxDayNum)
	binary.Read(buffer, binary.LittleEndian, &p.SupportBoxLastTime)
	binary.Read(buffer, binary.LittleEndian, &p.CooperationDayBuyNum)
	binary.Read(buffer, binary.LittleEndian, &p.CooperationDayGiveNum)
	binary.Read(buffer, binary.LittleEndian, &p.CooperationDayNum)
	binary.Read(buffer, binary.LittleEndian, &p.FightLoserProtectDayNum)
	binary.Read(buffer, binary.LittleEndian, &p.FightLoserDecScore)
	binary.Read(buffer, binary.LittleEndian, &p.FightWinDayNum)
	binary.Read(buffer, binary.LittleEndian, &p.FightWinExtra)
	binary.Read(buffer, binary.LittleEndian, &p.WeekCardAddCooperationExtraNum)
	binary.Read(buffer, binary.LittleEndian, &p.CooperationExtra)
	binary.Read(buffer, binary.LittleEndian, &p.UsedCooperationLastExtraEnd)
	binary.Read(buffer, binary.LittleEndian, &p.DayFirstWinNum)
	binary.Read(buffer, binary.LittleEndian, &p.DayFirstIsWin)
	binary.Read(buffer, binary.LittleEndian, &p.CooperADNum)
	binary.Read(buffer, binary.LittleEndian, &p.CooperADRoundNum)
	binary.Read(buffer, binary.LittleEndian, &p.IsGetFirstRechargeReward)
	binary.Read(buffer, binary.LittleEndian, &p.FightLookADDoubleCount)
	binary.Read(buffer, binary.LittleEndian, &p.CooperationDayHelpNum)
	if buffer.Len() < 4 {
		return errors.New("message length error")
	}
	var QQHallBlueVipVectorLen uint32
	binary.Read(buffer, binary.LittleEndian, &QQHallBlueVipVectorLen)
	p.QQHallBlueVipVector = make([]int32, QQHallBlueVipVectorLen)
	for i := uint32(0); i < QQHallBlueVipVectorLen; i++ {
		binary.Read(buffer, binary.LittleEndian, &p.QQHallBlueVipVector[i])
	}
	if buffer.Len() < 4 {
		return errors.New("message length error")
	}
	var QQHallRewardsVectorLen uint32
	binary.Read(buffer, binary.LittleEndian, &QQHallRewardsVectorLen)
	for i := uint32(0); i < QQHallRewardsVectorLen; i++ {
		if buffer.Len() < 4 {
			return errors.New("message length error")
		}
		var length uint32
		binary.Read(buffer, binary.LittleEndian, &length)
		if uint32(buffer.Len()) < length {
			return errors.New("message length error")
		}
		p.QQHallRewardsVector = append(p.QQHallRewardsVector, string(buffer.Next(int(length))))
	}
	if buffer.Len() < 4 {
		return errors.New("message length error")
	}
	var HeroSkinMapLen uint32
	binary.Read(buffer, binary.LittleEndian, &HeroSkinMapLen)
	if uint32(buffer.Len()) < HeroSkinMapLen*12 {
		return errors.New("message length error")
	}
	p.HeroSkinMap = make(map[int64]int32, HeroSkinMapLen)
	for i := uint32(0); i < HeroSkinMapLen; i++ {
		var key int64
		binary.Read(buffer, binary.LittleEndian, &key)
		var value int32
		binary.Read(buffer, binary.LittleEndian, &value)
		p.HeroSkinMap[key] = value
	}
	if buffer.Len() < 4 {
		return errors.New("message length error")
	}
	var SelfSelectMapLen uint32
	binary.Read(buffer, binary.LittleEndian, &SelfSelectMapLen)
	if uint32(buffer.Len()) < SelfSelectMapLen*8 {
		return errors.New("message length error")
	}
	p.SelfSelectMap = make(map[int32]int32, SelfSelectMapLen)
	for i := uint32(0); i < SelfSelectMapLen; i++ {
		var key int32
		binary.Read(buffer, binary.LittleEndian, &key)
		var value int32
		binary.Read(buffer, binary.LittleEndian, &value)
		p.SelfSelectMap[key] = value
	}
	if buffer.Len() < 4 {
		return errors.New("message length error")
	}
	var CarLinkMapLen uint32
	binary.Read(buffer, binary.LittleEndian, &CarLinkMapLen)
	if uint32(buffer.Len()) < CarLinkMapLen*8 {
		return errors.New("message length error")
	}
	p.CarLinkMap = make(map[int32]int32, CarLinkMapLen)
	for i := uint32(0); i < CarLinkMapLen; i++ {
		var key int32
		binary.Read(buffer, binary.LittleEndian, &key)
		var value int32
		binary.Read(buffer, binary.LittleEndian, &value)
		p.CarLinkMap[key] = value
	}
	return nil
}

type T_Client_Data struct {
	IntMap map[int32]int32
}

func (p *T_Client_Data) Encode() []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.LittleEndian, uint32(len(p.IntMap)))
	for k, v := range p.IntMap {
		binary.Write(buffer, binary.LittleEndian, k)
		binary.Write(buffer, binary.LittleEndian, v)
	}
	return buffer.Bytes()
}

func (p *T_Client_Data) Decode(buffer *bytes.Buffer) error {
	if buffer.Len() < 4 {
		return errors.New("message length error")
	}
	var IntMapLen uint32
	binary.Read(buffer, binary.LittleEndian, &IntMapLen)
	if buffer.Len() < int(IntMapLen*8) {
		return errors.New("message length error")
	}
	p.IntMap = make(map[int32]int32, IntMapLen)
	for i := 0; i < int(IntMapLen); i++ {
		var key int32
		binary.Read(buffer, binary.LittleEndian, &key)
		var value int32
		binary.Read(buffer, binary.LittleEndian, &value)
		p.IntMap[key] = value
	}
	return nil
}

type T_Role_Recharge_Single struct {
	BFirstPrize bool
	TotalNum    int32
	DayNum      int32
}

func (p *T_Role_Recharge_Single) Encode() []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.LittleEndian, p.BFirstPrize)
	binary.Write(buffer, binary.LittleEndian, p.TotalNum)
	binary.Write(buffer, binary.LittleEndian, p.DayNum)
	return buffer.Bytes()
}

func (p *T_Role_Recharge_Single) Decode(buffer *bytes.Buffer) error {
	if buffer.Len() < 9 {
		return errors.New("message length error")
	}
	binary.Read(buffer, binary.LittleEndian, &p.BFirstPrize)
	binary.Read(buffer, binary.LittleEndian, &p.TotalNum)
	binary.Read(buffer, binary.LittleEndian, &p.DayNum)
	return nil
}

type T_Role_Recharge_Data struct {
	Recharges map[int32]T_Role_Recharge_Single
}

func (p *T_Role_Recharge_Data) Encode() []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.LittleEndian, uint32(len(p.Recharges)))
	for k, v := range p.Recharges {
		binary.Write(buffer, binary.LittleEndian, k)
		buffer.Write(v.Encode())
	}
	return buffer.Bytes()
}

func (p *T_Role_Recharge_Data) Decode(buffer *bytes.Buffer) error {
	if buffer.Len() < 4 {
		return errors.New("message length error")
	}
	var RechargesLen uint32
	binary.Read(buffer, binary.LittleEndian, &RechargesLen)
	if buffer.Len() < int(RechargesLen*13) {
		return errors.New("message length error")
	}
	p.Recharges = make(map[int32]T_Role_Recharge_Single, RechargesLen)
	for i := 0; i < int(RechargesLen); i++ {
		var key int32
		binary.Read(buffer, binary.LittleEndian, &key)
		var value T_Role_Recharge_Single
		value.Decode(buffer)
		p.Recharges[key] = value
	}
	return nil
}

type T_Role_Item struct {
	ID_          primitive.ObjectID `bson:"_id"`
	RoleID       int64              `bson:"role_id"`
	ItemUUID     int64              `bson:"uuid"`
	ItemID       int32              `bson:"id"`
	OwnUUID      int64              `bson:"own_uuid"`
	CreateTime   int32              `bson:"create_time"`
	ExpiryCDTime int32              `bson:"expiry_cd_time"`
	ItemName     string             `bson:"name"`
	ItemNum      int32              `bson:"num"`
	ItemPos      int32              `bson:"pos"`
	ItemLevel    int32              `bson:"level"`
	ItemExp      int32              `bson:"exp"`
	ItemType     int32              `bson:"type"`
}

func (p *T_Role_Item) Encode() []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.LittleEndian, p.ItemUUID)
	binary.Write(buffer, binary.LittleEndian, p.ItemID)
	binary.Write(buffer, binary.LittleEndian, p.OwnUUID)
	binary.Write(buffer, binary.LittleEndian, p.CreateTime)
	binary.Write(buffer, binary.LittleEndian, p.ExpiryCDTime)
	binary.Write(buffer, binary.LittleEndian, uint32(len(p.ItemName)))
	buffer.Write([]byte(p.ItemName))
	binary.Write(buffer, binary.LittleEndian, p.ItemNum)
	binary.Write(buffer, binary.LittleEndian, p.ItemPos)
	binary.Write(buffer, binary.LittleEndian, p.ItemLevel)
	binary.Write(buffer, binary.LittleEndian, p.ItemExp)
	return buffer.Bytes()
}

func (p *T_Role_Item) Decode(buffer *bytes.Buffer) error {
	if buffer.Len() < 32 {
		return errors.New("message length error")
	}
	binary.Read(buffer, binary.LittleEndian, &p.ItemUUID)
	binary.Read(buffer, binary.LittleEndian, &p.ItemID)
	binary.Read(buffer, binary.LittleEndian, &p.OwnUUID)
	binary.Read(buffer, binary.LittleEndian, &p.CreateTime)
	binary.Read(buffer, binary.LittleEndian, &p.ExpiryCDTime)
	var ItemNameLen uint32
	binary.Read(buffer, binary.LittleEndian, &ItemNameLen)
	if uint32(buffer.Len()) < ItemNameLen {
		return errors.New("message length error")
	}
	p.ItemName = string(buffer.Next(int(ItemNameLen)))
	if buffer.Len() < 16 {
		return errors.New("message length error")
	}
	binary.Read(buffer, binary.LittleEndian, &p.ItemNum)
	binary.Read(buffer, binary.LittleEndian, &p.ItemPos)
	binary.Read(buffer, binary.LittleEndian, &p.ItemLevel)
	binary.Read(buffer, binary.LittleEndian, &p.ItemExp)
	return nil
}

type T_Role_Bag struct {
	Items map[int64]T_Role_Item
}

func (p *T_Role_Bag) Encode() []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.LittleEndian, uint32(len(p.Items)))
	for k, v := range p.Items {
		binary.Write(buffer, binary.LittleEndian, k)
		buffer.Write(v.Encode())
	}
	return buffer.Bytes()
}

func (p *T_Role_Bag) Decode(buffer *bytes.Buffer) error {
	if buffer.Len() < 4 {
		return errors.New("message length error")
	}
	var ItemsLen uint32
	binary.Read(buffer, binary.LittleEndian, &ItemsLen)
	p.Items = make(map[int64]T_Role_Item, ItemsLen)
	for i := 0; i < int(ItemsLen); i++ {
		var key int64
		binary.Read(buffer, binary.LittleEndian, &key)
		var value T_Role_Item
		value.Decode(buffer)
		p.Items[key] = value
	}
	return nil
}

type T_Role_MailAnnex struct {
	Annex []T_Reward
}

func (p *T_Role_MailAnnex) Encode() []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.LittleEndian, uint32(len(p.Annex)))
	for _, v := range p.Annex {
		buffer.Write(v.Encode())
	}
	return buffer.Bytes()
}

func (p *T_Role_MailAnnex) Decode(buffer *bytes.Buffer) error {
	if buffer.Len() < 4 {
		return errors.New("message length error")
	}
	var AnnexLen uint32
	binary.Read(buffer, binary.LittleEndian, &AnnexLen)
	p.Annex = make([]T_Reward, AnnexLen)
	for i := 0; i < int(AnnexLen); i++ {
		p.Annex[i].Decode(buffer)
	}
	return nil
}

type T_Role_SingleMail struct {
	MailUUID   int64
	MailID     int32
	CreateTime int32
	State      int32
	Title      string
	FromName   string
	Msg        string
	Annex      T_Role_MailAnnex
}

func (p *T_Role_SingleMail) Encode() []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.LittleEndian, p.MailUUID)
	binary.Write(buffer, binary.LittleEndian, p.MailID)
	binary.Write(buffer, binary.LittleEndian, p.CreateTime)
	binary.Write(buffer, binary.LittleEndian, p.State)
	binary.Write(buffer, binary.LittleEndian, int32(len(p.Title)))
	buffer.Write([]byte(p.Title))
	binary.Write(buffer, binary.LittleEndian, int32(len(p.FromName)))
	buffer.Write([]byte(p.FromName))
	binary.Write(buffer, binary.LittleEndian, int32(len(p.Msg)))
	buffer.Write([]byte(p.Msg))
	buffer.Write(p.Annex.Encode())
	return buffer.Bytes()
}

func (p *T_Role_SingleMail) Decode(buffer *bytes.Buffer) error {
	if buffer.Len() < 24 {
		return errors.New("message length error")
	}
	binary.Read(buffer, binary.LittleEndian, &p.MailUUID)
	binary.Read(buffer, binary.LittleEndian, &p.MailID)
	binary.Read(buffer, binary.LittleEndian, &p.CreateTime)
	binary.Read(buffer, binary.LittleEndian, &p.State)
	var TitleLen uint32
	binary.Read(buffer, binary.LittleEndian, &TitleLen)
	if uint32(buffer.Len()) < TitleLen {
		return errors.New("message length error")
	}
	p.Title = string(buffer.Next(int(TitleLen)))
	if buffer.Len() < 4 {
		return errors.New("message length error")
	}
	var FromNameLen uint32
	binary.Read(buffer, binary.LittleEndian, &FromNameLen)
	if uint32(buffer.Len()) < FromNameLen {
		return errors.New("message length error")
	}
	p.FromName = string(buffer.Next(int(FromNameLen)))
	if buffer.Len() < 4 {
		return errors.New("message length error")
	}
	var MsgLen uint32
	binary.Read(buffer, binary.LittleEndian, &MsgLen)
	if uint32(buffer.Len()) < MsgLen {
		return errors.New("message length error")
	}
	p.Msg = string(buffer.Next(int(MsgLen)))
	p.Annex.Decode(buffer)
	return nil
}

type T_Role_Mail struct {
	Mails map[int64]T_Role_SingleMail
}

func (p *T_Role_Mail) Encode() []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.LittleEndian, uint32(len(p.Mails)))
	for k, v := range p.Mails {
		binary.Write(buffer, binary.LittleEndian, k)
		buffer.Write(v.Encode())
	}
	return buffer.Bytes()
}

func (p *T_Role_Mail) Decode(buffer *bytes.Buffer) error {
	if buffer.Len() < 4 {
		return errors.New("message length error")
	}
	var MailsLen uint32
	binary.Read(buffer, binary.LittleEndian, &MailsLen)
	p.Mails = make(map[int64]T_Role_SingleMail, MailsLen)
	for i := 0; i < int(MailsLen); i++ {
		var key int64
		binary.Read(buffer, binary.LittleEndian, &key)
		var value T_Role_SingleMail
		value.Decode(buffer)
		p.Mails[key] = value
	}
	return nil
}

type T_Role_ItemInfo_Day struct {
	DelNum int32
}

func (p *T_Role_ItemInfo_Day) Encode() []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.LittleEndian, p.DelNum)
	return buffer.Bytes()
}

func (p *T_Role_ItemInfo_Day) Decode(buffer *bytes.Buffer) error {
	if buffer.Len() < 4 {
		return errors.New("message length error")
	}
	binary.Read(buffer, binary.LittleEndian, &p.DelNum)
	return nil
}

type T_Role_ItemInfo struct {
	Day map[int32]T_Role_ItemInfo_Day
}

func (p *T_Role_ItemInfo) Encode() []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.LittleEndian, uint32(len(p.Day)))
	for k, v := range p.Day {
		binary.Write(buffer, binary.LittleEndian, k)
		buffer.Write(v.Encode())
	}
	return buffer.Bytes()
}

func (p *T_Role_ItemInfo) Decode(buffer *bytes.Buffer) error {
	if buffer.Len() < 4 {
		return errors.New("message length error")
	}
	var DayLen uint32
	binary.Read(buffer, binary.LittleEndian, &DayLen)
	if uint32(buffer.Len()) < DayLen*8 {
		return errors.New("message length error")
	}
	p.Day = make(map[int32]T_Role_ItemInfo_Day, DayLen)
	for i := 0; i < int(DayLen); i++ {
		var key int32
		binary.Read(buffer, binary.LittleEndian, &key)
		var value T_Role_ItemInfo_Day
		value.Decode(buffer)
		p.Day[key] = value
	}
	return nil
}

type T_Role_Task struct {
	Tasks map[int32]T_Role_SingleTask
	Extra T_Task_Extra_Data
}

func (p *T_Role_Task) Encode() []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.LittleEndian, uint32(len(p.Tasks)))
	for k, v := range p.Tasks {
		binary.Write(buffer, binary.LittleEndian, k)
		buffer.Write(v.Encode())
	}
	buffer.Write(p.Extra.Encode())
	return buffer.Bytes()
}

func (p *T_Role_Task) Decode(buffer *bytes.Buffer) error {
	if buffer.Len() < 4 {
		return errors.New("message length error")
	}
	var TasksLen uint32
	binary.Read(buffer, binary.LittleEndian, &TasksLen)
	if uint32(buffer.Len()) < TasksLen*24 {
		return errors.New("message length error")
	}
	p.Tasks = make(map[int32]T_Role_SingleTask, TasksLen)
	for i := 0; i < int(TasksLen); i++ {
		var key int32
		binary.Read(buffer, binary.LittleEndian, &key)
		var value T_Role_SingleTask
		value.Decode(buffer)
		p.Tasks[key] = value
	}
	p.Extra.Decode(buffer)
	return nil
}

type T_Role_ExchangeItem struct {
	ExchangeID int32
	TotleNum   int32
	DayNum     int32
}

func (p *T_Role_ExchangeItem) Encode() []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.LittleEndian, p.ExchangeID)
	binary.Write(buffer, binary.LittleEndian, p.TotleNum)
	binary.Write(buffer, binary.LittleEndian, p.DayNum)
	return buffer.Bytes()
}

func (p *T_Role_ExchangeItem) Decode(buffer *bytes.Buffer) error {
	if buffer.Len() < 12 {
		return errors.New("message length error")
	}
	binary.Read(buffer, binary.LittleEndian, &p.ExchangeID)
	binary.Read(buffer, binary.LittleEndian, &p.TotleNum)
	binary.Read(buffer, binary.LittleEndian, &p.DayNum)
	return nil
}

type T_Role_ExchangeGroup struct {
	GroupID   int32
	Exchanges map[int32]T_Role_ExchangeItem
}

func (p *T_Role_ExchangeGroup) Encode() []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.LittleEndian, p.GroupID)
	binary.Write(buffer, binary.LittleEndian, uint32(len(p.Exchanges)))
	for k, v := range p.Exchanges {
		binary.Write(buffer, binary.LittleEndian, k)
		buffer.Write(v.Encode())
	}
	return buffer.Bytes()
}

func (p *T_Role_ExchangeGroup) Decode(buffer *bytes.Buffer) error {
	if buffer.Len() < 8 {
		return errors.New("message length error")
	}
	binary.Read(buffer, binary.LittleEndian, &p.GroupID)
	var ExchangesLen uint32
	binary.Read(buffer, binary.LittleEndian, &ExchangesLen)
	if uint32(buffer.Len()) < ExchangesLen*16 {
		return errors.New("message length error")
	}
	p.Exchanges = make(map[int32]T_Role_ExchangeItem, ExchangesLen)
	for i := 0; i < int(ExchangesLen); i++ {
		var key int32
		binary.Read(buffer, binary.LittleEndian, &key)
		var value T_Role_ExchangeItem
		value.Decode(buffer)
		p.Exchanges[key] = value
	}
	return nil
}

type T_Role_ExchangeData struct {
	Groups map[int32]T_Role_ExchangeGroup
}

func (p *T_Role_ExchangeData) Encode() []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.LittleEndian, uint32(len(p.Groups)))
	for k, v := range p.Groups {
		binary.Write(buffer, binary.LittleEndian, k)
		buffer.Write(v.Encode())
	}
	return buffer.Bytes()
}

func (p *T_Role_ExchangeData) Decode(buffer *bytes.Buffer) error {
	if buffer.Len() < 4 {
		return errors.New("message length error")
	}
	var GroupsLen uint32
	binary.Read(buffer, binary.LittleEndian, &GroupsLen)
	p.Groups = make(map[int32]T_Role_ExchangeGroup, GroupsLen)
	for i := 0; i < int(GroupsLen); i++ {
		var key int32
		binary.Read(buffer, binary.LittleEndian, &key)
		var value T_Role_ExchangeGroup
		value.Decode(buffer)
		p.Groups[key] = value
	}
	return nil
}

type T_Role_BattleArrayIndexData struct {
	HeroUUID int64
}

func (p *T_Role_BattleArrayIndexData) Encode() []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.LittleEndian, p.HeroUUID)
	return buffer.Bytes()
}

func (p *T_Role_BattleArrayIndexData) Decode(buffer *bytes.Buffer) error {
	if buffer.Len() < 8 {
		return errors.New("message length error")
	}
	binary.Read(buffer, binary.LittleEndian, &p.HeroUUID)
	return nil
}

type T_Role_BattleRuneIndexData struct {
	ItemID int32
}

func (p *T_Role_BattleRuneIndexData) Encode() []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.LittleEndian, p.ItemID)
	return buffer.Bytes()
}

func (p *T_Role_BattleRuneIndexData) Decode(buffer *bytes.Buffer) error {
	if buffer.Len() < 4 {
		return errors.New("message length error")
	}
	binary.Read(buffer, binary.LittleEndian, &p.ItemID)
	return nil
}

type T_Role_BattleArrayIDData struct {
	IndexData     map[int32]T_Role_BattleArrayIndexData
	RuneIndexData map[int32]T_Role_BattleRuneIndexData
	BattleArray   string
}

func (p *T_Role_BattleArrayIDData) Encode() []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.LittleEndian, uint32(len(p.IndexData)))
	for k, v := range p.IndexData {
		binary.Write(buffer, binary.LittleEndian, k)
		buffer.Write(v.Encode())
	}
	binary.Write(buffer, binary.LittleEndian, uint32(len(p.RuneIndexData)))
	for k, v := range p.RuneIndexData {
		binary.Write(buffer, binary.LittleEndian, k)
		buffer.Write(v.Encode())
	}
	binary.Write(buffer, binary.LittleEndian, uint32(len(p.BattleArray)))
	buffer.Write([]byte(p.BattleArray))
	return buffer.Bytes()
}

func (p *T_Role_BattleArrayIDData) Decode(buffer *bytes.Buffer) error {
	if buffer.Len() < 4 {
		return errors.New("message length error")
	}
	var IndexDataLen uint32
	binary.Read(buffer, binary.LittleEndian, &IndexDataLen)
	if buffer.Len() < int(IndexDataLen)*12 {
		return errors.New("message length error")
	}
	p.IndexData = make(map[int32]T_Role_BattleArrayIndexData, IndexDataLen)
	for i := 0; i < int(IndexDataLen); i++ {
		var key int32
		binary.Read(buffer, binary.LittleEndian, &key)
		var value T_Role_BattleArrayIndexData
		value.Decode(buffer)
		p.IndexData[key] = value
	}
	if buffer.Len() < 4 {
		return errors.New("message length error")
	}
	var RuneIndexDataLen uint32
	binary.Read(buffer, binary.LittleEndian, &RuneIndexDataLen)
	if buffer.Len() < int(RuneIndexDataLen)*8 {
		return errors.New("message length error")
	}
	p.RuneIndexData = make(map[int32]T_Role_BattleRuneIndexData, RuneIndexDataLen)
	for i := 0; i < int(RuneIndexDataLen); i++ {
		var key int32
		binary.Read(buffer, binary.LittleEndian, &key)
		var value T_Role_BattleRuneIndexData
		value.Decode(buffer)
		p.RuneIndexData[key] = value
	}
	if buffer.Len() < 4 {
		return errors.New("message length error")
	}
	var BattleArrayLen uint32
	binary.Read(buffer, binary.LittleEndian, &BattleArrayLen)
	if buffer.Len() < int(BattleArrayLen) {
		return errors.New("message length error")
	}
	p.BattleArray = string(buffer.Next(int(BattleArrayLen)))
	return nil
}

type T_Role_BattleArrayData struct {
	DefineID int32
	IDData   map[int32]T_Role_BattleArrayIDData
}

func (p *T_Role_BattleArrayData) Encode() []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.LittleEndian, p.DefineID)
	binary.Write(buffer, binary.LittleEndian, uint32(len(p.IDData)))
	for k, v := range p.IDData {
		binary.Write(buffer, binary.LittleEndian, k)
		buffer.Write(v.Encode())
	}
	return buffer.Bytes()
}

func (p *T_Role_BattleArrayData) Decode(buffer *bytes.Buffer) error {
	if buffer.Len() < 8 {
		return errors.New("message length error")
	}
	binary.Read(buffer, binary.LittleEndian, &p.DefineID)
	var IDDataLen uint32
	binary.Read(buffer, binary.LittleEndian, &IDDataLen)
	p.IDData = make(map[int32]T_Role_BattleArrayIDData, IDDataLen)
	for i := 0; i < int(IDDataLen); i++ {
		var key int32
		binary.Read(buffer, binary.LittleEndian, &key)
		var value T_Role_BattleArrayIDData
		value.Decode(buffer)
		p.IDData[key] = value
	}
	return nil
}

type T_Role_ExpressionArrayIndexData struct {
	ExpressionUUID int64
}

func (p *T_Role_ExpressionArrayIndexData) Encode() []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.LittleEndian, p.ExpressionUUID)
	return buffer.Bytes()
}

func (p *T_Role_ExpressionArrayIndexData) Decode(buffer *bytes.Buffer) error {
	if buffer.Len() < 8 {
		return errors.New("message length error")
	}
	binary.Read(buffer, binary.LittleEndian, &p.ExpressionUUID)
	return nil
}

type T_Role_ExpressionArrayData struct {
	ArrayData map[int32]T_Role_ExpressionArrayIndexData
}

func (p *T_Role_ExpressionArrayData) Encode() []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.LittleEndian, uint32(len(p.ArrayData)))
	for k, v := range p.ArrayData {
		binary.Write(buffer, binary.LittleEndian, k)
		buffer.Write(v.Encode())
	}
	return buffer.Bytes()
}

func (p *T_Role_ExpressionArrayData) Decode(buffer *bytes.Buffer) error {
	if buffer.Len() < 4 {
		return errors.New("message length error")
	}
	var ArrayDataLen uint32
	binary.Read(buffer, binary.LittleEndian, &ArrayDataLen)
	if buffer.Len() < int(ArrayDataLen)*12 {
		return errors.New("message length error")
	}
	p.ArrayData = make(map[int32]T_Role_ExpressionArrayIndexData, ArrayDataLen)
	for i := 0; i < int(ArrayDataLen); i++ {
		var key int32
		binary.Read(buffer, binary.LittleEndian, &key)
		var value T_Role_ExpressionArrayIndexData
		value.Decode(buffer)
		p.ArrayData[key] = value
	}
	return nil
}

type T_Role_ScoreAchievementSingle struct {
	TakedInfo map[int32]int32
}

func (p *T_Role_ScoreAchievementSingle) Encode() []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.LittleEndian, uint32(len(p.TakedInfo)))
	for k, v := range p.TakedInfo {
		binary.Write(buffer, binary.LittleEndian, k)
		binary.Write(buffer, binary.LittleEndian, v)
	}
	return buffer.Bytes()
}

func (p *T_Role_ScoreAchievementSingle) Decode(buffer *bytes.Buffer) error {
	if buffer.Len() < 4 {
		return errors.New("message length error")
	}
	var TakedInfoLen uint32
	binary.Read(buffer, binary.LittleEndian, &TakedInfoLen)
	if buffer.Len() < int(TakedInfoLen)*8 {
		return errors.New("message length error")
	}
	p.TakedInfo = make(map[int32]int32, TakedInfoLen)
	for i := 0; i < int(TakedInfoLen); i++ {
		var key int32
		binary.Read(buffer, binary.LittleEndian, &key)
		var value int32
		binary.Read(buffer, binary.LittleEndian, &value)
		p.TakedInfo[key] = value
	}
	return nil
}

type T_Role_ScoreAchievement struct {
	ScoreAchievements map[int32]T_Role_ScoreAchievementSingle
}

func (p *T_Role_ScoreAchievement) Encode() []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.LittleEndian, uint32(len(p.ScoreAchievements)))
	for k, v := range p.ScoreAchievements {
		binary.Write(buffer, binary.LittleEndian, k)
		buffer.Write(v.Encode())
	}
	return buffer.Bytes()
}

func (p *T_Role_ScoreAchievement) Decode(buffer *bytes.Buffer) error {
	if buffer.Len() < 4 {
		return errors.New("message length error")
	}
	var ScoreAchievementsLen uint32
	binary.Read(buffer, binary.LittleEndian, &ScoreAchievementsLen)
	p.ScoreAchievements = make(map[int32]T_Role_ScoreAchievementSingle, ScoreAchievementsLen)
	for i := 0; i < int(ScoreAchievementsLen); i++ {
		var key int32
		binary.Read(buffer, binary.LittleEndian, &key)
		var value T_Role_ScoreAchievementSingle
		value.Decode(buffer)
		p.ScoreAchievements[key] = value
	}
	return nil
}

type T_SignInDayData struct {
	BReceivePrize bool
	BReceiveExtra bool
}

func (p *T_SignInDayData) Encode() []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.LittleEndian, p.BReceivePrize)
	binary.Write(buffer, binary.LittleEndian, p.BReceiveExtra)
	return buffer.Bytes()
}

func (p *T_SignInDayData) Decode(buffer *bytes.Buffer) error {
	if buffer.Len() < 2 {
		return errors.New("message length error")
	}
	binary.Read(buffer, binary.LittleEndian, &p.BReceivePrize)
	binary.Read(buffer, binary.LittleEndian, &p.BReceiveExtra)
	return nil
}

type T_SignInData struct {
	Days map[int32]T_SignInDayData
}

func (p *T_SignInData) Encode() []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.LittleEndian, uint32(len(p.Days)))
	for k, v := range p.Days {
		binary.Write(buffer, binary.LittleEndian, k)
		buffer.Write(v.Encode())
	}
	return buffer.Bytes()
}

func (p *T_SignInData) Decode(buffer *bytes.Buffer) error {
	if buffer.Len() < 4 {
		return errors.New("message length error")
	}
	var DaysLen uint32
	binary.Read(buffer, binary.LittleEndian, &DaysLen)
	if buffer.Len() < int(DaysLen)*6 {
		return errors.New("message length error")
	}
	p.Days = make(map[int32]T_SignInDayData, DaysLen)
	for i := 0; i < int(DaysLen); i++ {
		var key int32
		binary.Read(buffer, binary.LittleEndian, &key)
		var value T_SignInDayData
		value.Decode(buffer)
		p.Days[key] = value
	}
	return nil
}

type T_TimeBoxSingleData struct {
	BoxNum    int32
	CDTime    int32
	DayNum    int32
	ReceiveCD int32
}

func (p *T_TimeBoxSingleData) Encode() []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.LittleEndian, p.BoxNum)
	binary.Write(buffer, binary.LittleEndian, p.CDTime)
	binary.Write(buffer, binary.LittleEndian, p.DayNum)
	binary.Write(buffer, binary.LittleEndian, p.ReceiveCD)
	return buffer.Bytes()
}

func (p *T_TimeBoxSingleData) Decode(buffer *bytes.Buffer) error {
	if buffer.Len() < 16 {
		return errors.New("message length error")
	}
	binary.Read(buffer, binary.LittleEndian, &p.BoxNum)
	binary.Read(buffer, binary.LittleEndian, &p.CDTime)
	binary.Read(buffer, binary.LittleEndian, &p.DayNum)
	binary.Read(buffer, binary.LittleEndian, &p.ReceiveCD)
	return nil
}

type T_TimeBoxData struct {
	Boxs map[int32]T_TimeBoxSingleData
}

func (p *T_TimeBoxData) Encode() []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.LittleEndian, uint32(len(p.Boxs)))
	for k, v := range p.Boxs {
		binary.Write(buffer, binary.LittleEndian, k)
		buffer.Write(v.Encode())
	}
	return buffer.Bytes()
}

func (p *T_TimeBoxData) Decode(buffer *bytes.Buffer) error {
	if buffer.Len() < 4 {
		return errors.New("message length error")
	}
	var BoxsLen uint32
	binary.Read(buffer, binary.LittleEndian, &BoxsLen)
	if buffer.Len() < int(BoxsLen)*20 {
		return errors.New("message length error")
	}
	p.Boxs = make(map[int32]T_TimeBoxSingleData, BoxsLen)
	for i := 0; i < int(BoxsLen); i++ {
		var key int32
		binary.Read(buffer, binary.LittleEndian, &key)
		var value T_TimeBoxSingleData
		value.Decode(buffer)
		p.Boxs[key] = value
	}
	return nil
}

type T_ThemeBoxEntityLevelData struct {
	IsPrize         bool
	IsRechargePrize bool
}

func (p *T_ThemeBoxEntityLevelData) Encode() []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.LittleEndian, p.IsPrize)
	binary.Write(buffer, binary.LittleEndian, p.IsRechargePrize)
	return buffer.Bytes()
}

func (p *T_ThemeBoxEntityLevelData) Decode(buffer *bytes.Buffer) error {
	if buffer.Len() < 2 {
		return errors.New("message length error")
	}
	binary.Read(buffer, binary.LittleEndian, &p.IsPrize)
	binary.Read(buffer, binary.LittleEndian, &p.IsRechargePrize)
	return nil
}

type T_ThemeBoxEntityData struct {
	BoxID        int32
	BoxBeginTime int32
	BoxEndTime   int32
	BoxLevel     int32
	BoxExp       int32
	IsRecharge   bool
	LevelData    map[int32]T_ThemeBoxEntityLevelData
}

func (p *T_ThemeBoxEntityData) Encode() []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.LittleEndian, p.BoxID)
	binary.Write(buffer, binary.LittleEndian, p.BoxBeginTime)
	binary.Write(buffer, binary.LittleEndian, p.BoxEndTime)
	binary.Write(buffer, binary.LittleEndian, p.BoxLevel)
	binary.Write(buffer, binary.LittleEndian, p.BoxExp)
	binary.Write(buffer, binary.LittleEndian, p.IsRecharge)
	binary.Write(buffer, binary.LittleEndian, uint32(len(p.LevelData)))
	for k, v := range p.LevelData {
		binary.Write(buffer, binary.LittleEndian, k)
		buffer.Write(v.Encode())
	}
	return buffer.Bytes()
}

func (p *T_ThemeBoxEntityData) Decode(buffer *bytes.Buffer) error {
	if buffer.Len() < 25 {
		return errors.New("message length error")
	}
	binary.Read(buffer, binary.LittleEndian, &p.BoxID)
	binary.Read(buffer, binary.LittleEndian, &p.BoxBeginTime)
	binary.Read(buffer, binary.LittleEndian, &p.BoxEndTime)
	binary.Read(buffer, binary.LittleEndian, &p.BoxLevel)
	binary.Read(buffer, binary.LittleEndian, &p.BoxExp)
	binary.Read(buffer, binary.LittleEndian, &p.IsRecharge)
	var LevelDataLen uint32
	binary.Read(buffer, binary.LittleEndian, &LevelDataLen)
	if buffer.Len() < int(LevelDataLen)*6 {
		return errors.New("message length error")
	}
	p.LevelData = make(map[int32]T_ThemeBoxEntityLevelData, LevelDataLen)
	for i := 0; i < int(LevelDataLen); i++ {
		var key int32
		binary.Read(buffer, binary.LittleEndian, &key)
		var value T_ThemeBoxEntityLevelData
		value.Decode(buffer)
		p.LevelData[key] = value
	}
	return nil
}

type T_ThemeBoxData struct {
	Entity map[int32]T_ThemeBoxEntityData
}

func (p *T_ThemeBoxData) Encode() []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.LittleEndian, uint32(len(p.Entity)))
	for k, v := range p.Entity {
		binary.Write(buffer, binary.LittleEndian, k)
		buffer.Write(v.Encode())
	}
	return buffer.Bytes()
}

func (p *T_ThemeBoxData) Decode(buffer *bytes.Buffer) error {
	if buffer.Len() < 4 {
		return errors.New("message length error")
	}
	var EntityLen uint32
	binary.Read(buffer, binary.LittleEndian, &EntityLen)
	p.Entity = make(map[int32]T_ThemeBoxEntityData, EntityLen)
	for i := 0; i < int(EntityLen); i++ {
		var key int32
		binary.Read(buffer, binary.LittleEndian, &key)
		var value T_ThemeBoxEntityData
		value.Decode(buffer)
		p.Entity[key] = value
	}
	return nil
}

type T_SeasonForeverScorePrizeData struct {
	IsPrize bool
	IsExtra bool
}

func (p *T_SeasonForeverScorePrizeData) Encode() []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.LittleEndian, p.IsPrize)
	binary.Write(buffer, binary.LittleEndian, p.IsExtra)
	return buffer.Bytes()
}

func (p *T_SeasonForeverScorePrizeData) Decode(buffer *bytes.Buffer) error {
	if buffer.Len() < 2 {
		return errors.New("message length error")
	}
	binary.Read(buffer, binary.LittleEndian, &p.IsPrize)
	binary.Read(buffer, binary.LittleEndian, &p.IsExtra)
	return nil
}

type T_SeasonScorePrizeData struct {
	IsPrize bool
	IsExtra bool
}

func (p *T_SeasonScorePrizeData) Encode() []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.LittleEndian, p.IsPrize)
	binary.Write(buffer, binary.LittleEndian, p.IsExtra)
	return buffer.Bytes()
}

func (p *T_SeasonScorePrizeData) Decode(buffer *bytes.Buffer) error {
	if buffer.Len() < 2 {
		return errors.New("message length error")
	}
	binary.Read(buffer, binary.LittleEndian, &p.IsPrize)
	binary.Read(buffer, binary.LittleEndian, &p.IsExtra)
	return nil
}

type T_SeasonEntityData struct {
	ID_               primitive.ObjectID `bson:"_id"`
	RoleID            int64              `bson:"role_id"`
	SeasonID          int32              `bson:"season_id"`
	SeasonBeginTime   int32              `bson:"season_begin_time"`
	SeasonEndTime     int32              `bson:"season_end_time"`
	ForeverScorePrize map[int32]T_SeasonForeverScorePrizeData
	ScorePrize        map[int32]T_SeasonScorePrizeData
}

func (p *T_SeasonEntityData) Encode() []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.LittleEndian, p.SeasonID)
	binary.Write(buffer, binary.LittleEndian, p.SeasonBeginTime)
	binary.Write(buffer, binary.LittleEndian, p.SeasonEndTime)
	binary.Write(buffer, binary.LittleEndian, uint32(len(p.ForeverScorePrize)))
	for k, v := range p.ForeverScorePrize {
		binary.Write(buffer, binary.LittleEndian, k)
		buffer.Write(v.Encode())
	}
	binary.Write(buffer, binary.LittleEndian, uint32(len(p.ScorePrize)))
	for k, v := range p.ScorePrize {
		binary.Write(buffer, binary.LittleEndian, k)
		buffer.Write(v.Encode())
	}
	return buffer.Bytes()
}

func (p *T_SeasonEntityData) Decode(buffer *bytes.Buffer) error {
	if buffer.Len() < 16 {
		return errors.New("message length error")
	}
	binary.Read(buffer, binary.LittleEndian, &p.SeasonID)
	binary.Read(buffer, binary.LittleEndian, &p.SeasonBeginTime)
	binary.Read(buffer, binary.LittleEndian, &p.SeasonEndTime)
	var ForeverScorePrizeLen uint32
	binary.Read(buffer, binary.LittleEndian, &ForeverScorePrizeLen)
	if buffer.Len() < int(ForeverScorePrizeLen)*6 {
		return errors.New("message length error")
	}
	p.ForeverScorePrize = make(map[int32]T_SeasonForeverScorePrizeData, ForeverScorePrizeLen)
	for i := 0; i < int(ForeverScorePrizeLen); i++ {
		var key int32
		binary.Read(buffer, binary.LittleEndian, &key)
		var value T_SeasonForeverScorePrizeData
		value.Decode(buffer)
		p.ForeverScorePrize[key] = value
	}
	if buffer.Len() < 4 {
		return errors.New("message length error")
	}
	var ScorePrizeLen uint32
	binary.Read(buffer, binary.LittleEndian, &ScorePrizeLen)
	if buffer.Len() < int(ScorePrizeLen)*6 {
		return errors.New("message length error")
	}
	p.ScorePrize = make(map[int32]T_SeasonScorePrizeData, ScorePrizeLen)
	for i := 0; i < int(ScorePrizeLen); i++ {
		var key int32
		binary.Read(buffer, binary.LittleEndian, &key)
		var value T_SeasonScorePrizeData
		value.Decode(buffer)
		p.ScorePrize[key] = value
	}
	return nil
}

type T_SeasonLastData struct {
	SeasonID       int32
	Score          int32
	SeasonMinScore int32
}

func (p *T_SeasonLastData) Encode() []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.LittleEndian, p.SeasonID)
	binary.Write(buffer, binary.LittleEndian, p.Score)
	binary.Write(buffer, binary.LittleEndian, p.SeasonMinScore)
	return buffer.Bytes()
}

func (p *T_SeasonLastData) Decode(buffer *bytes.Buffer) error {
	if buffer.Len() < 12 {
		return errors.New("message length error")
	}
	binary.Read(buffer, binary.LittleEndian, &p.SeasonID)
	binary.Read(buffer, binary.LittleEndian, &p.Score)
	binary.Read(buffer, binary.LittleEndian, &p.SeasonMinScore)
	return nil
}

type T_SeasonData struct {
	Entity map[int32]T_SeasonEntityData
	Last   map[int32]T_SeasonLastData
}

func (p *T_SeasonData) Encode() []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.LittleEndian, uint32(len(p.Entity)))
	for k, v := range p.Entity {
		binary.Write(buffer, binary.LittleEndian, k)
		buffer.Write(v.Encode())
	}
	binary.Write(buffer, binary.LittleEndian, uint32(len(p.Last)))
	for k, v := range p.Last {
		binary.Write(buffer, binary.LittleEndian, k)
		buffer.Write(v.Encode())
	}
	return buffer.Bytes()
}

func (p *T_SeasonData) Decode(buffer *bytes.Buffer) error {
	if buffer.Len() < 4 {
		return errors.New("message length error")
	}
	var EntityLen uint32
	binary.Read(buffer, binary.LittleEndian, &EntityLen)
	p.Entity = make(map[int32]T_SeasonEntityData, EntityLen)
	for i := 0; i < int(EntityLen); i++ {
		var key int32
		binary.Read(buffer, binary.LittleEndian, &key)
		var value T_SeasonEntityData
		value.Decode(buffer)
		p.Entity[key] = value
	}
	if buffer.Len() < 4 {
		return errors.New("message length error")
	}
	var LastLen uint32
	binary.Read(buffer, binary.LittleEndian, &LastLen)
	if buffer.Len() < int(LastLen)*16 {
		return errors.New("message length error")
	}
	p.Last = make(map[int32]T_SeasonLastData, LastLen)
	for i := 0; i < int(LastLen); i++ {
		var key int32
		binary.Read(buffer, binary.LittleEndian, &key)
		var value T_SeasonLastData
		value.Decode(buffer)
		p.Last[key] = value
	}
	return nil
}

type T_SharePlayerRankData struct {
}

func (p *T_SharePlayerRankData) Encode() []byte {
	buffer := new(bytes.Buffer)
	return buffer.Bytes()
}

func (p *T_SharePlayerRankData) Decode(buffer *bytes.Buffer) error {
	return nil
}

type T_ShareRankPrizeData struct {
	IsReceive bool
}

func (p *T_ShareRankPrizeData) Encode() []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.LittleEndian, p.IsReceive)
	return buffer.Bytes()
}

func (p *T_ShareRankPrizeData) Decode(buffer *bytes.Buffer) error {
	if buffer.Len() < 1 {
		return errors.New("message length error")
	}
	binary.Read(buffer, binary.LittleEndian, &p.IsReceive)
	return nil
}

type T_SharePlayerData struct {
	BindTime     int32
	RefreshTime  int32
	NickName     string
	HeadID       int32
	HeadUrl      string
	HeadFrameID  int32
	MaxScore     int32
	IsSharePrize bool
	RankPrize    map[int32]T_ShareRankPrizeData
}

func (p *T_SharePlayerData) Encode() []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.LittleEndian, p.BindTime)
	binary.Write(buffer, binary.LittleEndian, p.RefreshTime)
	binary.Write(buffer, binary.LittleEndian, uint32(len(p.NickName)))
	buffer.WriteString(p.NickName)
	binary.Write(buffer, binary.LittleEndian, p.HeadID)
	binary.Write(buffer, binary.LittleEndian, uint32(len(p.HeadUrl)))
	buffer.WriteString(p.HeadUrl)
	binary.Write(buffer, binary.LittleEndian, p.HeadFrameID)
	binary.Write(buffer, binary.LittleEndian, p.MaxScore)
	binary.Write(buffer, binary.LittleEndian, p.IsSharePrize)
	binary.Write(buffer, binary.LittleEndian, uint32(len(p.RankPrize)))
	for k, v := range p.RankPrize {
		binary.Write(buffer, binary.LittleEndian, k)
		buffer.Write(v.Encode())
	}
	return buffer.Bytes()
}

func (p *T_SharePlayerData) Decode(buffer *bytes.Buffer) error {
	if buffer.Len() < 12 {
		return errors.New("message length error")
	}
	binary.Read(buffer, binary.LittleEndian, &p.BindTime)
	binary.Read(buffer, binary.LittleEndian, &p.RefreshTime)
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
	if buffer.Len() < 13 {
		return errors.New("message length error")
	}
	binary.Read(buffer, binary.LittleEndian, &p.HeadFrameID)
	binary.Read(buffer, binary.LittleEndian, &p.MaxScore)
	binary.Read(buffer, binary.LittleEndian, &p.IsSharePrize)
	var RankPrizeLen uint32
	binary.Read(buffer, binary.LittleEndian, &RankPrizeLen)
	if buffer.Len() < int(RankPrizeLen)*5 {
		return errors.New("message length error")
	}
	p.RankPrize = make(map[int32]T_ShareRankPrizeData, RankPrizeLen)
	for i := 0; i < int(RankPrizeLen); i++ {
		var key int32
		binary.Read(buffer, binary.LittleEndian, &key)
		var value T_ShareRankPrizeData
		value.Decode(buffer)
		p.RankPrize[key] = value
	}
	return nil
}

type T_ShareData struct {
	Ranks            map[int64]T_SharePlayerRankData
	Players          map[int64]T_SharePlayerData
	DaySharePrizeNum int32
}

func (p *T_ShareData) Encode() []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.LittleEndian, uint32(len(p.Ranks)))
	for k, v := range p.Ranks {
		binary.Write(buffer, binary.LittleEndian, k)
		buffer.Write(v.Encode())
	}
	binary.Write(buffer, binary.LittleEndian, uint32(len(p.Players)))
	for k, v := range p.Players {
		binary.Write(buffer, binary.LittleEndian, k)
		buffer.Write(v.Encode())
	}
	binary.Write(buffer, binary.LittleEndian, p.DaySharePrizeNum)
	return buffer.Bytes()
}

func (p *T_ShareData) Decode(buffer *bytes.Buffer) error {
	if buffer.Len() < 4 {
		return errors.New("message length error")
	}
	var RanksLen uint32
	binary.Read(buffer, binary.LittleEndian, &RanksLen)
	if buffer.Len() < int(RanksLen)*8 {
		return errors.New("message length error")
	}
	p.Ranks = make(map[int64]T_SharePlayerRankData, RanksLen)
	for i := 0; i < int(RanksLen); i++ {
		var key int64
		binary.Read(buffer, binary.LittleEndian, &key)
		var value T_SharePlayerRankData
		value.Decode(buffer)
		p.Ranks[key] = value
	}
	if buffer.Len() < 4 {
		return errors.New("message length error")
	}
	var PlayersLen uint32
	binary.Read(buffer, binary.LittleEndian, &PlayersLen)
	p.Players = make(map[int64]T_SharePlayerData, PlayersLen)
	for i := 0; i < int(PlayersLen); i++ {
		var key int64
		binary.Read(buffer, binary.LittleEndian, &key)
		var value T_SharePlayerData
		value.Decode(buffer)
		p.Players[key] = value
	}
	if buffer.Len() < 4 {
		return errors.New("message length error")
	}
	binary.Read(buffer, binary.LittleEndian, &p.DaySharePrizeNum)
	return nil
}

type T_TotalSignInData struct {
	SigninRound    int32
	SigninIndex    int32
	SigninDay      int32
	IsReceive      bool
	IsExtraReceive bool
}

func (p *T_TotalSignInData) Encode() []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.LittleEndian, p.SigninRound)
	binary.Write(buffer, binary.LittleEndian, p.SigninIndex)
	binary.Write(buffer, binary.LittleEndian, p.SigninDay)
	binary.Write(buffer, binary.LittleEndian, p.IsReceive)
	binary.Write(buffer, binary.LittleEndian, p.IsExtraReceive)
	return buffer.Bytes()
}

func (p *T_TotalSignInData) Decode(buffer *bytes.Buffer) error {
	if buffer.Len() < 14 {
		return errors.New("message length error")
	}
	binary.Read(buffer, binary.LittleEndian, &p.SigninRound)
	binary.Read(buffer, binary.LittleEndian, &p.SigninIndex)
	binary.Read(buffer, binary.LittleEndian, &p.SigninDay)
	binary.Read(buffer, binary.LittleEndian, &p.IsReceive)
	binary.Read(buffer, binary.LittleEndian, &p.IsExtraReceive)
	return nil
}

type T_CDKData struct {
	Data map[string]int32
}

func (p *T_CDKData) Encode() []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.LittleEndian, uint32(len(p.Data)))
	for k, v := range p.Data {
		binary.Write(buffer, binary.LittleEndian, uint32(len(k)))
		binary.Write(buffer, binary.LittleEndian, []byte(k))
		binary.Write(buffer, binary.LittleEndian, v)
	}
	return buffer.Bytes()
}

func (p *T_CDKData) Decode(buffer *bytes.Buffer) error {
	if buffer.Len() < 4 {
		return errors.New("message length error")
	}
	var DataLen uint32
	binary.Read(buffer, binary.LittleEndian, &DataLen)
	p.Data = make(map[string]int32, DataLen)
	for i := 0; i < int(DataLen); i++ {
		if buffer.Len() < 4 {
			return errors.New("message length error")
		}
		var keyLen uint32
		binary.Read(buffer, binary.LittleEndian, &keyLen)
		if buffer.Len() < int(keyLen) {
			return errors.New("message length error")
		}
		key := make([]byte, keyLen)
		binary.Read(buffer, binary.LittleEndian, &key)
		var value int32
		binary.Read(buffer, binary.LittleEndian, &value)
		p.Data[string(key)] = value
	}
	return nil
}

type T_HallofFameRoleData struct {
	Num int32
}

func (p *T_HallofFameRoleData) Encode() []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.LittleEndian, p.Num)
	return buffer.Bytes()
}

func (p *T_HallofFameRoleData) Decode(buffer *bytes.Buffer) error {
	if buffer.Len() < 4 {
		return errors.New("message length error")
	}
	binary.Read(buffer, binary.LittleEndian, &p.Num)
	return nil
}

type T_HallofFameData struct {
	Data map[int32]T_HallofFameRoleData
}

func (p *T_HallofFameData) Encode() []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.LittleEndian, uint32(len(p.Data)))
	for k, v := range p.Data {
		binary.Write(buffer, binary.LittleEndian, k)
		binary.Write(buffer, binary.LittleEndian, v.Encode())
	}
	return buffer.Bytes()
}

func (p *T_HallofFameData) Decode(buffer *bytes.Buffer) error {
	if buffer.Len() < 4 {
		return errors.New("message length error")
	}
	var DataLen uint32
	binary.Read(buffer, binary.LittleEndian, &DataLen)
	if buffer.Len() < int(DataLen)*8 {
		return errors.New("message length error")
	}
	p.Data = make(map[int32]T_HallofFameRoleData, DataLen)
	for i := 0; i < int(DataLen); i++ {
		var key int32
		binary.Read(buffer, binary.LittleEndian, &key)
		var value T_HallofFameRoleData
		value.Decode(buffer)
		p.Data[key] = value
	}
	return nil
}

type T_CondSharePlayer struct {
	HeadID  int32
	HeadUrl string
}

func (p *T_CondSharePlayer) Encode() []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.LittleEndian, p.HeadID)
	binary.Write(buffer, binary.LittleEndian, uint32(len(p.HeadUrl)))
	binary.Write(buffer, binary.LittleEndian, []byte(p.HeadUrl))
	return buffer.Bytes()
}

func (p *T_CondSharePlayer) Decode(buffer *bytes.Buffer) error {
	if buffer.Len() < 8 {
		return errors.New("message length error")
	}
	binary.Read(buffer, binary.LittleEndian, &p.HeadID)
	var HeadUrlLen uint32
	binary.Read(buffer, binary.LittleEndian, &HeadUrlLen)
	if buffer.Len() < int(HeadUrlLen) {
		return errors.New("message length error")
	}
	HeadUrl := make([]byte, HeadUrlLen)
	binary.Read(buffer, binary.LittleEndian, &HeadUrl)
	p.HeadUrl = string(HeadUrl)
	return nil
}

type T_CondShare struct {
	BeginTime int32
	EndTime   int32
	Players   map[int64]T_CondSharePlayer
	BReceive  bool
}

func (p *T_CondShare) Encode() []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.LittleEndian, p.BeginTime)
	binary.Write(buffer, binary.LittleEndian, p.EndTime)
	binary.Write(buffer, binary.LittleEndian, uint32(len(p.Players)))
	for k, v := range p.Players {
		binary.Write(buffer, binary.LittleEndian, k)
		binary.Write(buffer, binary.LittleEndian, v.Encode())
	}
	binary.Write(buffer, binary.LittleEndian, p.BReceive)
	return buffer.Bytes()
}

func (p *T_CondShare) Decode(buffer *bytes.Buffer) error {
	if buffer.Len() < 12 {
		return errors.New("message length error")
	}
	binary.Read(buffer, binary.LittleEndian, &p.BeginTime)
	binary.Read(buffer, binary.LittleEndian, &p.EndTime)
	var PlayersLen uint32
	binary.Read(buffer, binary.LittleEndian, &PlayersLen)
	for i := 0; i < int(PlayersLen); i++ {
		var key int64
		binary.Read(buffer, binary.LittleEndian, &key)
		var value T_CondSharePlayer
		value.Decode(buffer)
		p.Players[key] = value
	}
	if buffer.Len() < 1 {
		return errors.New("message length error")
	}
	binary.Read(buffer, binary.LittleEndian, &p.BReceive)
	return nil
}

type T_CondShareData struct {
	Condshares map[int32]T_CondShare
}

func (p *T_CondShareData) Encode() []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.LittleEndian, uint32(len(p.Condshares)))
	for k, v := range p.Condshares {
		binary.Write(buffer, binary.LittleEndian, k)
		binary.Write(buffer, binary.LittleEndian, v.Encode())
	}
	return buffer.Bytes()
}

func (p *T_CondShareData) Decode(buffer *bytes.Buffer) error {
	if buffer.Len() < 4 {
		return errors.New("message length error")
	}
	var CondsharesLen uint32
	binary.Read(buffer, binary.LittleEndian, &CondsharesLen)
	p.Condshares = make(map[int32]T_CondShare, CondsharesLen)
	for i := 0; i < int(CondsharesLen); i++ {
		var key int32
		binary.Read(buffer, binary.LittleEndian, &key)
		var value T_CondShare
		if value.Decode(buffer) != nil {
			return errors.New("message length error")
		}
		p.Condshares[key] = value
	}
	return nil
}

type T_SingleFinalRuneData struct {
	Level int32
}

func (p *T_SingleFinalRuneData) Encode() []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.LittleEndian, p.Level)
	return buffer.Bytes()
}

func (p *T_SingleFinalRuneData) Decode(buffer *bytes.Buffer) error {
	if buffer.Len() < 4 {
		return errors.New("message length error")
	}
	binary.Read(buffer, binary.LittleEndian, &p.Level)
	return nil
}

type T_FinalRuneData struct {
	Runes map[int32]T_SingleFinalRuneData
}

func (p *T_FinalRuneData) Encode() []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.LittleEndian, uint32(len(p.Runes)))
	for k, v := range p.Runes {
		binary.Write(buffer, binary.LittleEndian, k)
		binary.Write(buffer, binary.LittleEndian, v.Encode())
	}
	return buffer.Bytes()
}

func (p *T_FinalRuneData) Decode(buffer *bytes.Buffer) error {
	if buffer.Len() < 4 {
		return errors.New("message length error")
	}
	var RunesLen uint32
	binary.Read(buffer, binary.LittleEndian, &RunesLen)
	if buffer.Len() < int(RunesLen)*8 {
		return errors.New("message length error")
	}
	for i := 0; i < int(RunesLen); i++ {
		var key int32
		binary.Read(buffer, binary.LittleEndian, &key)
		var value T_SingleFinalRuneData
		if value.Decode(buffer) != nil {
			return errors.New("message length error")
		}
		p.Runes[key] = value
	}
	return nil
}

type T_TimeLockBoxPositionData struct {
	BoxID     int32
	BoxState  int32
	StateTime int32
}

func (p *T_TimeLockBoxPositionData) Encode() []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.LittleEndian, p.BoxID)
	binary.Write(buffer, binary.LittleEndian, p.BoxState)
	binary.Write(buffer, binary.LittleEndian, p.StateTime)
	return buffer.Bytes()
}

func (p *T_TimeLockBoxPositionData) Decode(buffer *bytes.Buffer) error {
	if buffer.Len() < 12 {
		return errors.New("message length error")
	}
	binary.Read(buffer, binary.LittleEndian, &p.BoxID)
	binary.Read(buffer, binary.LittleEndian, &p.BoxState)
	binary.Read(buffer, binary.LittleEndian, &p.StateTime)
	return nil
}

type T_TimeLockBoxData struct {
	Position              map[int32]T_TimeLockBoxPositionData
	DayFreeFastReceiveNum int32
}

func (p *T_TimeLockBoxData) Encode() []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.LittleEndian, uint32(len(p.Position)))
	for k, v := range p.Position {
		binary.Write(buffer, binary.LittleEndian, k)
		binary.Write(buffer, binary.LittleEndian, v.Encode())
	}
	binary.Write(buffer, binary.LittleEndian, p.DayFreeFastReceiveNum)
	return buffer.Bytes()
}

func (p *T_TimeLockBoxData) Decode(buffer *bytes.Buffer) error {
	if buffer.Len() < 4 {
		return errors.New("message length error")
	}
	var PositionLen uint32
	binary.Read(buffer, binary.LittleEndian, &PositionLen)
	if buffer.Len() < int(PositionLen)*16 {
		return errors.New("message length error")
	}
	p.Position = make(map[int32]T_TimeLockBoxPositionData)
	for i := 0; i < int(PositionLen); i++ {
		var key int32
		binary.Read(buffer, binary.LittleEndian, &key)
		var value T_TimeLockBoxPositionData
		if value.Decode(buffer) != nil {
			return errors.New("message length error")
		}
		p.Position[key] = value
	}
	if buffer.Len() < 4 {
		return errors.New("message length error")
	}
	binary.Read(buffer, binary.LittleEndian, &p.DayFreeFastReceiveNum)
	return nil
}

type T_Role_ChapterInfo struct {
	ChapterId         int32
	ChapterProgress   int32
	RewardBoxStateMap map[int32]int32
}

func (p *T_Role_ChapterInfo) Encode() []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.LittleEndian, p.ChapterId)
	binary.Write(buffer, binary.LittleEndian, p.ChapterProgress)
	binary.Write(buffer, binary.LittleEndian, uint32(len(p.RewardBoxStateMap)))
	for k, v := range p.RewardBoxStateMap {
		binary.Write(buffer, binary.LittleEndian, k)
		binary.Write(buffer, binary.LittleEndian, v)
	}
	return buffer.Bytes()
}

func (p *T_Role_ChapterInfo) Decode(buffer *bytes.Buffer) error {
	if buffer.Len() < 12 {
		return errors.New("message length error")
	}
	binary.Read(buffer, binary.LittleEndian, &p.ChapterId)
	binary.Read(buffer, binary.LittleEndian, &p.ChapterProgress)
	var RewardBoxStateMapLen uint32
	binary.Read(buffer, binary.LittleEndian, &RewardBoxStateMapLen)
	if buffer.Len() < int(RewardBoxStateMapLen)*8 {
		return errors.New("message length error")
	}
	p.RewardBoxStateMap = make(map[int32]int32, RewardBoxStateMapLen)
	for i := 0; i < int(RewardBoxStateMapLen); i++ {
		var key int32
		binary.Read(buffer, binary.LittleEndian, &key)
		var value int32
		binary.Read(buffer, binary.LittleEndian, &value)
		p.RewardBoxStateMap[key] = value
	}
	return nil
}

type T_Role_ChapterData struct {
	CurrentChapterId int32
	ChapterInfoMap   map[int32]T_Role_ChapterInfo
}

func (p *T_Role_ChapterData) Encode() []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.LittleEndian, p.CurrentChapterId)
	binary.Write(buffer, binary.LittleEndian, uint32(len(p.ChapterInfoMap)))
	for k, v := range p.ChapterInfoMap {
		binary.Write(buffer, binary.LittleEndian, k)
		binary.Write(buffer, binary.LittleEndian, v.Encode())
	}
	return buffer.Bytes()
}

func (p *T_Role_ChapterData) Decode(buffer *bytes.Buffer) error {
	if buffer.Len() < 8 {
		return errors.New("message length error")
	}
	binary.Read(buffer, binary.LittleEndian, &p.CurrentChapterId)
	var ChapterInfoMapLen uint32
	binary.Read(buffer, binary.LittleEndian, &ChapterInfoMapLen)
	p.ChapterInfoMap = make(map[int32]T_Role_ChapterInfo, ChapterInfoMapLen)
	for i := 0; i < int(ChapterInfoMapLen); i++ {
		var key int32
		binary.Read(buffer, binary.LittleEndian, &key)
		var value T_Role_ChapterInfo
		if value.Decode(buffer) != nil {
			return errors.New("message length error")
		}
		p.ChapterInfoMap[key] = value
	}
	return nil
}

type T_Role_LegendData struct {
	ChapterInfoMap map[int32]int32
}

func (p *T_Role_LegendData) Encode() []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.LittleEndian, uint32(len(p.ChapterInfoMap)))
	for k, v := range p.ChapterInfoMap {
		binary.Write(buffer, binary.LittleEndian, k)
		binary.Write(buffer, binary.LittleEndian, v)
	}
	return buffer.Bytes()
}

func (p *T_Role_LegendData) Decode(buffer *bytes.Buffer) error {
	if buffer.Len() < 4 {
		return errors.New("message length error")
	}
	var ChapterInfoMapLen uint32
	binary.Read(buffer, binary.LittleEndian, &ChapterInfoMapLen)
	if buffer.Len() < int(ChapterInfoMapLen)*8 {
		return errors.New("message length error")
	}
	p.ChapterInfoMap = make(map[int32]int32, ChapterInfoMapLen)
	for i := 0; i < int(ChapterInfoMapLen); i++ {
		var key int32
		binary.Read(buffer, binary.LittleEndian, &key)
		var value int32
		binary.Read(buffer, binary.LittleEndian, &value)
		p.ChapterInfoMap[key] = value
	}
	return nil
}

type T_FightBalance_Hero struct {
	HeroID      int32
	DamageTotal int64
}

func (p *T_FightBalance_Hero) Encode() []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.LittleEndian, p.HeroID)
	binary.Write(buffer, binary.LittleEndian, p.DamageTotal)
	return buffer.Bytes()
}

func (p *T_FightBalance_Hero) Decode(buffer *bytes.Buffer) error {
	if buffer.Len() < 12 {
		return errors.New("message length error")
	}
	binary.Read(buffer, binary.LittleEndian, &p.HeroID)
	binary.Read(buffer, binary.LittleEndian, &p.DamageTotal)
	return nil
}

type T_FightBalance_TypMonster struct {
	DeathNum int32
}

func (p *T_FightBalance_TypMonster) Encode() []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.LittleEndian, p.DeathNum)
	return buffer.Bytes()
}

func (p *T_FightBalance_TypMonster) Decode(buffer *bytes.Buffer) error {
	if buffer.Len() < 4 {
		return errors.New("message length error")
	}
	binary.Read(buffer, binary.LittleEndian, &p.DeathNum)
	return nil
}

type T_FightBalance_Role struct {
	RoleAbstract  T_RoleAbstract
	Heros         map[int64]T_FightBalance_Hero
	TypeMonster   map[int32]T_FightBalance_TypMonster
	TotalHarm     int64
	IsGuard       bool
	StatisticInfo T_Fight_Statistic_info
}

func (p *T_FightBalance_Role) Encode() []byte {
	buffer := new(bytes.Buffer)
	buffer.Write(p.RoleAbstract.Encode())
	var HerosLen uint32 = uint32(len(p.Heros))
	binary.Write(buffer, binary.LittleEndian, HerosLen)
	for k, v := range p.Heros {
		binary.Write(buffer, binary.LittleEndian, k)
		buffer.Write(v.Encode())
	}
	var TypeMonsterLen uint32 = uint32(len(p.TypeMonster))
	binary.Write(buffer, binary.LittleEndian, TypeMonsterLen)
	for k, v := range p.TypeMonster {
		binary.Write(buffer, binary.LittleEndian, k)
		buffer.Write(v.Encode())
	}
	binary.Write(buffer, binary.LittleEndian, p.TotalHarm)
	binary.Write(buffer, binary.LittleEndian, p.IsGuard)
	buffer.Write(p.StatisticInfo.Encode())
	return buffer.Bytes()
}

func (p *T_FightBalance_Role) Decode(buffer *bytes.Buffer) error {
	p.RoleAbstract.Decode(buffer)
	if buffer.Len() < 4 {
		return errors.New("message length error")
	}
	var HerosLen uint32
	binary.Read(buffer, binary.LittleEndian, &HerosLen)
	if uint32(buffer.Len()) < HerosLen*20 {
		return errors.New("message length error")
	}
	p.Heros = make(map[int64]T_FightBalance_Hero, HerosLen)
	for i := uint32(0); i < HerosLen; i++ {
		var k int64
		var v T_FightBalance_Hero
		binary.Read(buffer, binary.LittleEndian, &k)
		v.Decode(buffer)
		p.Heros[k] = v
	}
	if buffer.Len() < 4 {
		return errors.New("message length error")
	}
	var TypeMonsterLen uint32
	binary.Read(buffer, binary.LittleEndian, &TypeMonsterLen)
	if uint32(buffer.Len()) < TypeMonsterLen*8 {
		return errors.New("message length error")
	}
	p.TypeMonster = make(map[int32]T_FightBalance_TypMonster, TypeMonsterLen)
	for i := uint32(0); i < TypeMonsterLen; i++ {
		var k int32
		var v T_FightBalance_TypMonster
		binary.Read(buffer, binary.LittleEndian, &k)
		v.Decode(buffer)
		p.TypeMonster[k] = v
	}
	if buffer.Len() < 9 {
		return errors.New("message length error")
	}
	binary.Read(buffer, binary.LittleEndian, &p.TotalHarm)
	binary.Read(buffer, binary.LittleEndian, &p.IsGuard)
	p.StatisticInfo.Decode(buffer)
	return nil
}

type T_FightBalance_Battle_Winner struct {
	Prize          []T_Reward
	EncouragePrize []T_Reward
	SeriesWinPrize []T_Reward
}

func (p *T_FightBalance_Battle_Winner) Encode() []byte {
	buffer := new(bytes.Buffer)
	var PrizeLen uint32 = uint32(len(p.Prize))
	binary.Write(buffer, binary.LittleEndian, PrizeLen)
	for _, v := range p.Prize {
		buffer.Write(v.Encode())
	}
	var EncouragePrizeLen uint32 = uint32(len(p.EncouragePrize))
	binary.Write(buffer, binary.LittleEndian, EncouragePrizeLen)
	for _, v := range p.EncouragePrize {
		buffer.Write(v.Encode())
	}
	var SeriesWinPrizeLen uint32 = uint32(len(p.SeriesWinPrize))
	binary.Write(buffer, binary.LittleEndian, SeriesWinPrizeLen)
	for _, v := range p.SeriesWinPrize {
		buffer.Write(v.Encode())
	}
	return buffer.Bytes()
}

func (p *T_FightBalance_Battle_Winner) Decode(buffer *bytes.Buffer) error {
	if buffer.Len() < 4 {
		return errors.New("message length error")
	}
	var PrizeLen uint32
	binary.Read(buffer, binary.LittleEndian, &PrizeLen)
	p.Prize = make([]T_Reward, PrizeLen)
	for i := uint32(0); i < PrizeLen; i++ {
		p.Prize[i].Decode(buffer)
	}
	if buffer.Len() < 4 {
		return errors.New("message length error")
	}
	var EncouragePrizeLen uint32
	binary.Read(buffer, binary.LittleEndian, &EncouragePrizeLen)
	p.EncouragePrize = make([]T_Reward, EncouragePrizeLen)
	for i := uint32(0); i < EncouragePrizeLen; i++ {
		p.EncouragePrize[i].Decode(buffer)
	}
	if buffer.Len() < 4 {
		return errors.New("message length error")
	}
	var SeriesWinPrizeLen uint32
	binary.Read(buffer, binary.LittleEndian, &SeriesWinPrizeLen)
	p.SeriesWinPrize = make([]T_Reward, SeriesWinPrizeLen)
	for i := uint32(0); i < SeriesWinPrizeLen; i++ {
		p.SeriesWinPrize[i].Decode(buffer)
	}
	return nil
}

type T_FightBalance_Battle_Loster struct {
	LostDec   []T_Reward
	LostPrize []T_Reward
}

func (p *T_FightBalance_Battle_Loster) Encode() []byte {
	buffer := new(bytes.Buffer)
	var LostDecLen uint32 = uint32(len(p.LostDec))
	binary.Write(buffer, binary.LittleEndian, LostDecLen)
	for _, v := range p.LostDec {
		buffer.Write(v.Encode())
	}
	var LostPrizeLen uint32 = uint32(len(p.LostPrize))
	binary.Write(buffer, binary.LittleEndian, LostPrizeLen)
	for _, v := range p.LostPrize {
		buffer.Write(v.Encode())
	}
	return buffer.Bytes()
}

func (p *T_FightBalance_Battle_Loster) Decode(buffer *bytes.Buffer) error {
	if buffer.Len() < 4 {
		return errors.New("message length error")
	}
	var LostDecLen uint32
	binary.Read(buffer, binary.LittleEndian, &LostDecLen)
	p.LostDec = make([]T_Reward, LostDecLen)
	for i := uint32(0); i < LostDecLen; i++ {
		p.LostDec[i].Decode(buffer)
	}
	if buffer.Len() < 4 {
		return errors.New("message length error")
	}
	var LostPrizeLen uint32
	binary.Read(buffer, binary.LittleEndian, &LostPrizeLen)
	p.LostPrize = make([]T_Reward, LostPrizeLen)
	for i := uint32(0); i < LostPrizeLen; i++ {
		p.LostPrize[i].Decode(buffer)
	}
	return nil
}

type T_FightBalance_Battle struct {
	Winner T_FightBalance_Battle_Winner
	Loster T_FightBalance_Battle_Loster
}

func (p *T_FightBalance_Battle) Encode() []byte {
	buffer := new(bytes.Buffer)
	buffer.Write(p.Winner.Encode())
	buffer.Write(p.Loster.Encode())
	return buffer.Bytes()
}

func (p *T_FightBalance_Battle) Decode(buffer *bytes.Buffer) error {
	p.Winner.Decode(buffer)
	p.Loster.Decode(buffer)
	return nil
}

type T_FightBalance_CoopRation struct {
	Prize      []T_Reward
	Extraprize []T_Reward
}

func (p *T_FightBalance_CoopRation) Encode() []byte {
	buffer := new(bytes.Buffer)
	var PrizeLen uint32 = uint32(len(p.Prize))
	binary.Write(buffer, binary.LittleEndian, PrizeLen)
	for _, v := range p.Prize {
		buffer.Write(v.Encode())
	}
	var ExtraprizeLen uint32 = uint32(len(p.Extraprize))
	binary.Write(buffer, binary.LittleEndian, ExtraprizeLen)
	for _, v := range p.Extraprize {
		buffer.Write(v.Encode())
	}
	return buffer.Bytes()
}

func (p *T_FightBalance_CoopRation) Decode(buffer *bytes.Buffer) error {
	if buffer.Len() < 4 {
		return errors.New("message length error")
	}
	var PrizeLen uint32
	binary.Read(buffer, binary.LittleEndian, &PrizeLen)
	p.Prize = make([]T_Reward, PrizeLen)
	for i := uint32(0); i < PrizeLen; i++ {
		p.Prize[i].Decode(buffer)
	}
	if buffer.Len() < 4 {
		return errors.New("message length error")
	}
	var ExtraprizeLen uint32
	binary.Read(buffer, binary.LittleEndian, &ExtraprizeLen)
	p.Extraprize = make([]T_Reward, ExtraprizeLen)
	for i := uint32(0); i < ExtraprizeLen; i++ {
		p.Extraprize[i].Decode(buffer)
	}
	return nil
}

type T_FightBalance_RandomArena struct {
	WinNum  int32
	LostNum int32
}

func (p *T_FightBalance_RandomArena) Encode() []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.LittleEndian, p.WinNum)
	binary.Write(buffer, binary.LittleEndian, p.LostNum)
	return buffer.Bytes()
}

func (p *T_FightBalance_RandomArena) Decode(buffer *bytes.Buffer) error {
	if buffer.Len() < 8 {
		return errors.New("message length error")
	}
	binary.Read(buffer, binary.LittleEndian, &p.WinNum)
	binary.Read(buffer, binary.LittleEndian, &p.LostNum)
	return nil
}

type T_FightBlance_GoldenLeague struct {
	Score int32
}

func (p *T_FightBlance_GoldenLeague) Encode() []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.LittleEndian, p.Score)
	return buffer.Bytes()
}

func (p *T_FightBlance_GoldenLeague) Decode(buffer *bytes.Buffer) error {
	if buffer.Len() < 4 {
		return errors.New("message length error")
	}
	binary.Read(buffer, binary.LittleEndian, &p.Score)
	return nil
}

type T_FightBalance_Activity_CoopRation struct {
	Prize []T_Reward
}

func (p *T_FightBalance_Activity_CoopRation) Encode() []byte {
	buffer := new(bytes.Buffer)
	var PrizeLen uint32 = uint32(len(p.Prize))
	binary.Write(buffer, binary.LittleEndian, PrizeLen)
	for _, v := range p.Prize {
		buffer.Write(v.Encode())
	}
	return buffer.Bytes()
}

func (p *T_FightBalance_Activity_CoopRation) Decode(buffer *bytes.Buffer) error {
	if buffer.Len() < 4 {
		return errors.New("message length error")
	}
	var PrizeLen uint32
	binary.Read(buffer, binary.LittleEndian, &PrizeLen)
	p.Prize = make([]T_Reward, PrizeLen)
	for i := uint32(0); i < PrizeLen; i++ {
		p.Prize[i].Decode(buffer)
	}
	return nil
}

type T_FightBalance_ExtraData struct {
	ThemeMedal int32
}

func (p *T_FightBalance_ExtraData) Encode() []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.LittleEndian, p.ThemeMedal)
	return buffer.Bytes()
}

func (p *T_FightBalance_ExtraData) Decode(buffer *bytes.Buffer) error {
	if buffer.Len() < 4 {
		return errors.New("message length error")
	}
	binary.Read(buffer, binary.LittleEndian, &p.ThemeMedal)
	return nil
}

/************************************  客户端  *********************************/

type C_Role_SynRoleData struct {
	ClientVer string
}

func (p *C_Role_SynRoleData) Encode() []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.LittleEndian, uint32(len(p.ClientVer)))
	buffer.Write([]byte(p.ClientVer))
	return buffer.Bytes()
}

func (p *C_Role_SynRoleData) Decode(buffer *bytes.Buffer, key uint8) error {
	if key != 0 {
		for i := 0; i < buffer.Len(); i++ {
			buffer.Bytes()[i] ^= byte(key)
		}
	}
	var ClientVerLen uint32
	binary.Read(buffer, binary.LittleEndian, &ClientVerLen)
	if uint32(buffer.Len()) < ClientVerLen {
		return errors.New("message length error")
	}
	p.ClientVer = string(buffer.Next(int(ClientVerLen)))
	return nil
}

// 客户端请求数据结构-设置战斗阵容
type C_Role_BattleArraySetDefine struct {
	ArrayID int32
}

func (p *C_Role_BattleArraySetDefine) Encode() []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.LittleEndian, p.ArrayID)
	return buffer.Bytes()
}

func (p *C_Role_BattleArraySetDefine) Decode(buffer *bytes.Buffer, key uint8) error {
	if key != 0 {
		for i := 0; i < buffer.Len(); i++ {
			buffer.Bytes()[i] ^= byte(key)
		}
	}
	if buffer.Len() < 4 {
		return errors.New("message length error")
	}
	binary.Read(buffer, binary.LittleEndian, &p.ArrayID)
	return nil
}

// 客户端请求数据结构-阵容上阵修改
type C_Role_BattleArrayUp struct {
	ArrayID    int32
	ArrayIndex int32
	ItemUUID   int64
}

func (p *C_Role_BattleArrayUp) Encode() []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.LittleEndian, p.ArrayID)
	binary.Write(buffer, binary.LittleEndian, p.ArrayIndex)
	binary.Write(buffer, binary.LittleEndian, p.ItemUUID)
	return buffer.Bytes()
}

func (p *C_Role_BattleArrayUp) Decode(buffer *bytes.Buffer, key uint8) error {
	if key != 0 {
		for i := 0; i < buffer.Len(); i++ {
			buffer.Bytes()[i] ^= byte(key)
		}
	}
	if buffer.Len() < 16 {
		return errors.New("message length error")
	}
	binary.Read(buffer, binary.LittleEndian, &p.ArrayID)
	binary.Read(buffer, binary.LittleEndian, &p.ArrayIndex)
	binary.Read(buffer, binary.LittleEndian, &p.ItemUUID)
	return nil
}

// 客户端请求数据结构-战车皮肤更换
type C_Role_Car_Skin_Change struct {
	SkinId int32
}

func (p *C_Role_Car_Skin_Change) Encode() []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.LittleEndian, p.SkinId)
	return buffer.Bytes()
}

func (p *C_Role_Car_Skin_Change) Decode(buffer *bytes.Buffer, key uint8) error {
	if key != 0 {
		for i := 0; i < buffer.Len(); i++ {
			buffer.Bytes()[i] ^= byte(key)
		}
	}
	if buffer.Len() < 4 {
		return errors.New("message length error")
	}
	binary.Read(buffer, binary.LittleEndian, &p.SkinId)
	return nil
}

// 客户端请求数据结构-英雄皮肤更换
type C_Role_HeroChangeSkin struct {
	HeroUUID int64
	SkinId   int32
}

func (p *C_Role_HeroChangeSkin) Encode() []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.LittleEndian, p.HeroUUID)
	binary.Write(buffer, binary.LittleEndian, p.SkinId)
	return buffer.Bytes()
}

func (p *C_Role_HeroChangeSkin) Decode(buffer *bytes.Buffer, key uint8) error {
	if key != 0 {
		for i := 0; i < buffer.Len(); i++ {
			buffer.Bytes()[i] ^= byte(key)
		}
	}
	if buffer.Len() < 12 {
		return errors.New("message length error")
	}
	binary.Read(buffer, binary.LittleEndian, &p.HeroUUID)
	binary.Read(buffer, binary.LittleEndian, &p.SkinId)
	return nil
}

// 客户端请求数据结构-阵容名称设置
type C_Role_SetBattleArrayName struct {
	BattleArrayIndex int32
	BattleArrayName  string
}

func (p *C_Role_SetBattleArrayName) Encode() []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.LittleEndian, p.BattleArrayIndex)
	binary.Write(buffer, binary.LittleEndian, uint32(len(p.BattleArrayName)))
	buffer.Write([]byte(p.BattleArrayName))
	return buffer.Bytes()
}

func (p *C_Role_SetBattleArrayName) Decode(buffer *bytes.Buffer, key uint8) error {
	if key != 0 {
		for i := 0; i < buffer.Len(); i++ {
			buffer.Bytes()[i] ^= byte(key)
		}

	}
	if buffer.Len() < 8 {
		return errors.New("message length error")
	}
	binary.Read(buffer, binary.LittleEndian, &p.BattleArrayIndex)
	var BattleArrayNameLen uint32
	binary.Read(buffer, binary.LittleEndian, &BattleArrayNameLen)
	if uint32(buffer.Len()) < BattleArrayNameLen {
		return errors.New("message length error")
	}
	p.BattleArrayName = string(buffer.Next(int(BattleArrayNameLen)))
	return nil
}

// 客户端请求数据结构-获取角色简要信息
type C_Role_GetRoleSimpleInfo struct {
	ShowID string
}

func (p *C_Role_GetRoleSimpleInfo) Encode() []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.LittleEndian, uint32(len(p.ShowID)))
	buffer.Write([]byte(p.ShowID))
	return buffer.Bytes()
}

func (p *C_Role_GetRoleSimpleInfo) Decode(buffer *bytes.Buffer, key uint8) error {
	if key != 0 {
		for i := 0; i < buffer.Len(); i++ {
			buffer.Bytes()[i] ^= byte(key)
		}
	}
	var ShowIDLen uint32
	binary.Read(buffer, binary.LittleEndian, &ShowIDLen)
	if uint32(buffer.Len()) < ShowIDLen {
		return errors.New("message length error")
	}
	p.ShowID = string(buffer.Next(int(ShowIDLen)))
	return nil
}

type C_Role_SetGuide struct {
	Guide int32
}

func (p *C_Role_SetGuide) Encode() []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.LittleEndian, p.Guide)
	return buffer.Bytes()
}

func (p *C_Role_SetGuide) Decode(buffer *bytes.Buffer, key uint8) error {
	if key != 0 {
		for i := 0; i < buffer.Len(); i++ {
			buffer.Bytes()[i] ^= byte(key)
		}
	}
	if buffer.Len() < 4 {
		return errors.New("message length error")
	}
	binary.Read(buffer, binary.LittleEndian, &p.Guide)
	return nil
}

/************************************  服务端  *********************************/

type S_Role_SynRoleAttrValue struct {
	ID_    primitive.ObjectID `bson:"_id"`     // 唯一ID
	RoleID int64              `bson:"role_id"` // 角色ID
	Index  int32              `bson:"attr_id"` // 属性ID
	Value  int32              `bson:"value"`   // 属性值
}

func (p *S_Role_SynRoleAttrValue) Encode() []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.LittleEndian, p.Index)
	binary.Write(buffer, binary.LittleEndian, p.Value)
	return buffer.Bytes()
}

func (p *S_Role_SynRoleAttrValue) Decode(buffer *bytes.Buffer) error {
	if buffer.Len() < 8 {
		return errors.New("message length error")
	}
	binary.Read(buffer, binary.LittleEndian, &p.Index)
	binary.Read(buffer, binary.LittleEndian, &p.Value)
	return nil
}

type S_Role_SynTaskData struct {
	ChangeTask []T_Role_SingleTask
	DeleteTask []int32
}

func (p *S_Role_SynTaskData) Encode() []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.LittleEndian, uint32(len(p.ChangeTask)))
	for _, v := range p.ChangeTask {
		buffer.Write(v.Encode())
	}
	binary.Write(buffer, binary.LittleEndian, uint32(len(p.DeleteTask)))
	for _, v := range p.DeleteTask {
		binary.Write(buffer, binary.LittleEndian, v)
	}
	return buffer.Bytes()
}

func (p *S_Role_SynTaskData) Decode(buffer *bytes.Buffer) error {
	if buffer.Len() < 4 {
		return errors.New("message length error")
	}
	var ChangeTaskLen uint32
	binary.Read(buffer, binary.LittleEndian, &ChangeTaskLen)
	p.ChangeTask = make([]T_Role_SingleTask, ChangeTaskLen)
	for i := uint32(0); i < ChangeTaskLen; i++ {
		var value T_Role_SingleTask
		value.Decode(buffer)
		p.ChangeTask[i] = value
	}
	if buffer.Len() < 4 {
		return errors.New("message length error")
	}
	var DeleteTaskLen uint32
	binary.Read(buffer, binary.LittleEndian, &DeleteTaskLen)
	p.DeleteTask = make([]int32, DeleteTaskLen)
	for i := uint32(0); i < DeleteTaskLen; i++ {
		binary.Read(buffer, binary.LittleEndian, &p.DeleteTask[i])
	}
	return nil
}

type S_Role_TotalWatchADBoxData struct {
	Totalwatchadbox T_TotalWatchADBox
}

func (p *S_Role_TotalWatchADBoxData) Encode() []byte {
	buffer := new(bytes.Buffer)
	buffer.Write(p.Totalwatchadbox.Encode())
	return buffer.Bytes()
}

func (p *S_Role_TotalWatchADBoxData) Decode(buffer *bytes.Buffer) error {
	if buffer.Len() < 18 {
		return errors.New("message length error")
	}
	p.Totalwatchadbox.Decode(buffer)
	return nil
}

type S_Role_SyncCostGet struct {
	CostGetMap map[int32]int32
}

func (p *S_Role_SyncCostGet) Encode() []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.LittleEndian, uint32(len(p.CostGetMap)))
	for k, v := range p.CostGetMap {
		binary.Write(buffer, binary.LittleEndian, k)
		binary.Write(buffer, binary.LittleEndian, v)
	}
	return buffer.Bytes()
}

func (p *S_Role_SyncCostGet) Decode(buffer *bytes.Buffer) error {
	if buffer.Len() < 4 {
		return errors.New("message length error")
	}
	var CostGetMapLen uint32
	binary.Read(buffer, binary.LittleEndian, &CostGetMapLen)
	if uint32(buffer.Len()) < CostGetMapLen*8 {
		return errors.New("message length error")
	}
	p.CostGetMap = make(map[int32]int32, CostGetMapLen)
	for i := uint32(0); i < CostGetMapLen; i++ {
		var key int32
		var value int32
		binary.Read(buffer, binary.LittleEndian, &key)
		binary.Read(buffer, binary.LittleEndian, &value)
		p.CostGetMap[key] = value
	}
	return nil
}

type S_Role_RoleEnterLogic struct {
}

func (p *S_Role_RoleEnterLogic) Encode() []byte {
	buffer := new(bytes.Buffer)
	return buffer.Bytes()
}
func (p *S_Role_RoleEnterLogic) Decode(buffer *bytes.Buffer) error {
	return nil
}

type S_Role_SynRoleData struct {
	CurrTime             int32
	RoleID               int64  // 角色ID
	StrID                string // 显示ID
	RoleName             string // 角色名字
	BIndulge             bool   // 是否处于防沉迷状态
	IndulgeTime          int32  // 防沉迷时间
	IndulgeDayOnlineTime int32  // 防沉迷当天在线时间
	RoleAttrValue        map[int32]int32
	GameTime             T_Game_Time
	Infomation           T_Information_Data
	ClientData           T_Client_Data
	Recharge             T_Role_Recharge_Data
	RoleBag              T_Role_Bag
	RoleMail             T_Role_Mail
	RoleItemInfo         T_Role_ItemInfo
	RoleTask             T_Role_Task
	Exchange             T_Role_ExchangeData
	Battlearray          T_Role_BattleArrayData
	Expressionarray      T_Role_ExpressionArrayData
	ScoreAchievement     T_Role_ScoreAchievement
	Signin               T_SignInData
	Timebox              T_TimeBoxData
	Themebox             T_ThemeBoxData
	Season               T_SeasonData
	Share                T_ShareData
	Totalsignin          T_TotalSignInData
	CDKdata              T_CDKData
	BCheckFightRoom      bool // 是否检查战斗房间
	Watchadbox           T_TotalWatchADBox
	Halloffame           T_HallofFameData
	Condshare            T_CondShareData
	Finalrune            T_FinalRuneData
	TimelockBox          T_TimeLockBoxData
	ChapterData          T_Role_ChapterData
	LegendData           T_Role_LegendData
}

func (p *S_Role_SynRoleData) Encode() []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.LittleEndian, p.CurrTime)
	binary.Write(buffer, binary.LittleEndian, p.RoleID)
	binary.Write(buffer, binary.LittleEndian, uint32(len(p.StrID)))
	buffer.WriteString(p.StrID)
	binary.Write(buffer, binary.LittleEndian, uint32(len(p.RoleName)))
	buffer.WriteString(p.RoleName)
	binary.Write(buffer, binary.LittleEndian, p.BIndulge)
	binary.Write(buffer, binary.LittleEndian, p.IndulgeTime)
	binary.Write(buffer, binary.LittleEndian, p.IndulgeDayOnlineTime)
	binary.Write(buffer, binary.LittleEndian, uint32(len(p.RoleAttrValue)))
	for k, v := range p.RoleAttrValue {
		binary.Write(buffer, binary.LittleEndian, k)
		binary.Write(buffer, binary.LittleEndian, v)
	}
	buffer.Write(p.GameTime.Encode())
	buffer.Write(p.Infomation.Encode())
	buffer.Write(p.ClientData.Encode())
	buffer.Write(p.Recharge.Encode())
	buffer.Write(p.RoleBag.Encode())
	buffer.Write(p.RoleMail.Encode())
	buffer.Write(p.RoleItemInfo.Encode())
	buffer.Write(p.RoleTask.Encode())
	buffer.Write(p.Exchange.Encode())
	buffer.Write(p.Battlearray.Encode())
	buffer.Write(p.Expressionarray.Encode())
	buffer.Write(p.ScoreAchievement.Encode())
	buffer.Write(p.Signin.Encode())
	buffer.Write(p.Timebox.Encode())
	buffer.Write(p.Themebox.Encode())
	buffer.Write(p.Season.Encode())
	buffer.Write(p.Share.Encode())
	buffer.Write(p.Totalsignin.Encode())
	buffer.Write(p.CDKdata.Encode())
	binary.Write(buffer, binary.LittleEndian, p.BCheckFightRoom)
	buffer.Write(p.Watchadbox.Encode())
	buffer.Write(p.Halloffame.Encode())
	buffer.Write(p.Condshare.Encode())
	buffer.Write(p.Finalrune.Encode())
	buffer.Write(p.TimelockBox.Encode())
	buffer.Write(p.ChapterData.Encode())
	buffer.Write(p.LegendData.Encode())
	return buffer.Bytes()
}

func (p *S_Role_SynRoleData) Decode(buffer *bytes.Buffer) error {
	if buffer.Len() < 16 {
		return errors.New("message length error")
	}
	binary.Read(buffer, binary.LittleEndian, &p.CurrTime)
	binary.Read(buffer, binary.LittleEndian, &p.RoleID)
	var StrIDLen uint32
	binary.Read(buffer, binary.LittleEndian, &StrIDLen)
	if uint32(buffer.Len()) < StrIDLen {
		return errors.New("message length error")
	}
	p.StrID = string(buffer.Next(int(StrIDLen)))
	if buffer.Len() < 4 {
		return errors.New("message length error")
	}
	var RoleNameLen uint32
	binary.Read(buffer, binary.LittleEndian, &RoleNameLen)
	if uint32(buffer.Len()) < RoleNameLen {
		return errors.New("message length error")
	}
	p.RoleName = string(buffer.Next(int(RoleNameLen)))
	if buffer.Len() < 13 {
		return errors.New("message length error")
	}
	binary.Read(buffer, binary.LittleEndian, &p.BIndulge)
	binary.Read(buffer, binary.LittleEndian, &p.IndulgeTime)
	binary.Read(buffer, binary.LittleEndian, &p.IndulgeDayOnlineTime)
	var RoleAttrValueLen uint32
	binary.Read(buffer, binary.LittleEndian, &RoleAttrValueLen)
	if uint32(buffer.Len()) < RoleAttrValueLen*8 {
		return errors.New("message length error")
	}
	p.RoleAttrValue = make(map[int32]int32, RoleAttrValueLen)
	for i := uint32(0); i < RoleAttrValueLen; i++ {
		var k int32
		var v int32
		binary.Read(buffer, binary.LittleEndian, &k)
		binary.Read(buffer, binary.LittleEndian, &v)
		p.RoleAttrValue[k] = v
	}
	if buffer.Len() < 28 {
		return errors.New("message length error")
	}
	p.GameTime.Decode(buffer)
	p.Infomation.Decode(buffer)
	p.ClientData.Decode(buffer)
	p.Recharge.Decode(buffer)
	p.RoleBag.Decode(buffer)
	p.RoleMail.Decode(buffer)
	p.RoleItemInfo.Decode(buffer)
	p.RoleTask.Decode(buffer)
	p.Exchange.Decode(buffer)
	if p.Battlearray.Decode(buffer) != nil {
		return errors.New("battlearray decode error")
	}
	p.Expressionarray.Decode(buffer)
	p.ScoreAchievement.Decode(buffer)
	p.Signin.Decode(buffer)
	p.Timebox.Decode(buffer)
	p.Themebox.Decode(buffer)
	p.Season.Decode(buffer)
	p.Share.Decode(buffer)
	if buffer.Len() < 14 {
		return errors.New("message length error")
	}
	p.Totalsignin.Decode(buffer)
	p.CDKdata.Decode(buffer)
	if buffer.Len() < 19 {
		return errors.New("message length error")
	}
	binary.Read(buffer, binary.LittleEndian, &p.BCheckFightRoom)
	p.Watchadbox.Decode(buffer)
	p.Halloffame.Decode(buffer)
	p.Condshare.Decode(buffer)
	p.Finalrune.Decode(buffer)
	p.TimelockBox.Decode(buffer)
	p.ChapterData.Decode(buffer)
	p.LegendData.Decode(buffer)
	return nil
}

// 服务器返回-同步角色开关数据
type S_Role_OnOffDataInfo struct {
	Onoff map[int32]bool
}

func (p *S_Role_OnOffDataInfo) Encode() []byte {
	buffer := new(bytes.Buffer)
	var OnoffLen uint32 = uint32(len(p.Onoff))
	binary.Write(buffer, binary.LittleEndian, OnoffLen)
	for k, v := range p.Onoff {
		binary.Write(buffer, binary.LittleEndian, k)
		binary.Write(buffer, binary.LittleEndian, v)
	}
	return buffer.Bytes()
}

func (p *S_Role_OnOffDataInfo) Decode(buffer *bytes.Buffer) error {
	if buffer.Len() < 4 {
		return errors.New("message length error")
	}
	var OnoffLen uint32
	binary.Read(buffer, binary.LittleEndian, &OnoffLen)
	if uint32(buffer.Len()) < OnoffLen*5 {
		return errors.New("message length error")
	}
	p.Onoff = make(map[int32]bool, OnoffLen)
	for i := uint32(0); i < OnoffLen; i++ {
		var k int32
		var v bool
		binary.Read(buffer, binary.LittleEndian, &k)
		binary.Read(buffer, binary.LittleEndian, &v)
		p.Onoff[k] = v
	}
	return nil
}

// 服务器返回-同步战斗阵容数据
type S_Role_SynBattleArrayData struct {
	Battlearray T_Role_BattleArrayData
}

func (p *S_Role_SynBattleArrayData) Encode() []byte {
	buffer := new(bytes.Buffer)
	buffer.Write(p.Battlearray.Encode())
	return buffer.Bytes()
}

func (p *S_Role_SynBattleArrayData) Decode(buffer *bytes.Buffer) error {
	if buffer.Len() < 0 {
		return errors.New("message length error")
	}
	p.Battlearray.Decode(buffer)
	return nil
}

type S_Role_FightBalance struct {
	Type               int32 // 战斗类型
	BWin               bool  // 是否胜利
	Roles              map[int64]T_FightBalance_Role
	Round              int32 // 战斗回合数
	Battle             map[int32]T_FightBalance_Battle
	Coopration         map[int32]T_FightBalance_CoopRation
	RandomArena        map[int64]T_FightBalance_RandomArena
	GoldenLeague       map[int32]T_FightBlance_GoldenLeague
	ActivityCoopration map[int32]T_FightBalance_Activity_CoopRation
	ExtraData          map[int32]T_FightBalance_ExtraData
}

func (p *S_Role_FightBalance) Encode() []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.LittleEndian, p.Type)
	binary.Write(buffer, binary.LittleEndian, p.BWin)
	binary.Write(buffer, binary.LittleEndian, int32(len(p.Roles)))
	for k, v := range p.Roles {
		binary.Write(buffer, binary.LittleEndian, k)
		buffer.Write(v.Encode())
	}
	binary.Write(buffer, binary.LittleEndian, p.Round)
	binary.Write(buffer, binary.LittleEndian, int32(len(p.Battle)))
	for k, v := range p.Battle {
		binary.Write(buffer, binary.LittleEndian, k)
		buffer.Write(v.Encode())
	}
	binary.Write(buffer, binary.LittleEndian, int32(len(p.Coopration)))
	for k, v := range p.Coopration {
		binary.Write(buffer, binary.LittleEndian, k)
		buffer.Write(v.Encode())
	}
	binary.Write(buffer, binary.LittleEndian, int32(len(p.RandomArena)))
	for k, v := range p.RandomArena {
		binary.Write(buffer, binary.LittleEndian, k)
		buffer.Write(v.Encode())
	}
	binary.Write(buffer, binary.LittleEndian, int32(len(p.GoldenLeague)))
	for k, v := range p.GoldenLeague {
		binary.Write(buffer, binary.LittleEndian, k)
		buffer.Write(v.Encode())
	}
	binary.Write(buffer, binary.LittleEndian, int32(len(p.ActivityCoopration)))
	for k, v := range p.ActivityCoopration {
		binary.Write(buffer, binary.LittleEndian, k)
		buffer.Write(v.Encode())
	}
	binary.Write(buffer, binary.LittleEndian, int32(len(p.ExtraData)))
	for k, v := range p.ExtraData {
		binary.Write(buffer, binary.LittleEndian, k)
		buffer.Write(v.Encode())
	}
	return buffer.Bytes()
}

func (p *S_Role_FightBalance) Decode(buffer *bytes.Buffer) error {
	if buffer.Len() < 9 {
		return errors.New("message length error")
	}
	binary.Read(buffer, binary.LittleEndian, &p.Type)
	binary.Read(buffer, binary.LittleEndian, &p.BWin)
	var RolesLen uint32
	binary.Read(buffer, binary.LittleEndian, &RolesLen)
	p.Roles = make(map[int64]T_FightBalance_Role, RolesLen)
	for i := uint32(0); i < RolesLen; i++ {
		var k int64
		var v T_FightBalance_Role
		binary.Read(buffer, binary.LittleEndian, &k)
		v.Decode(buffer)
		p.Roles[k] = v
	}
	if buffer.Len() < 8 {
		return errors.New("message length error")
	}
	binary.Read(buffer, binary.LittleEndian, &p.Round)
	var BattleLen uint32
	binary.Read(buffer, binary.LittleEndian, &BattleLen)
	p.Battle = make(map[int32]T_FightBalance_Battle, BattleLen)
	for i := uint32(0); i < BattleLen; i++ {
		var k int32
		var v T_FightBalance_Battle
		binary.Read(buffer, binary.LittleEndian, &k)
		v.Decode(buffer)
		p.Battle[k] = v
	}
	if buffer.Len() < 4 {
		return errors.New("message length error")
	}
	var CooprationLen uint32
	binary.Read(buffer, binary.LittleEndian, &CooprationLen)
	p.Coopration = make(map[int32]T_FightBalance_CoopRation, CooprationLen)
	for i := uint32(0); i < CooprationLen; i++ {
		var k int32
		var v T_FightBalance_CoopRation
		binary.Read(buffer, binary.LittleEndian, &k)
		v.Decode(buffer)
		p.Coopration[k] = v
	}
	if buffer.Len() < 4 {
		return errors.New("message length error")
	}
	var RandomArenaLen uint32
	binary.Read(buffer, binary.LittleEndian, &RandomArenaLen)
	if uint32(buffer.Len()) < RandomArenaLen*16 {
		return errors.New("message length error")
	}
	p.RandomArena = make(map[int64]T_FightBalance_RandomArena, RandomArenaLen)
	for i := uint32(0); i < RandomArenaLen; i++ {
		var k int64
		var v T_FightBalance_RandomArena
		binary.Read(buffer, binary.LittleEndian, &k)
		v.Decode(buffer)
		p.RandomArena[k] = v
	}
	if buffer.Len() < 4 {
		return errors.New("message length error")
	}
	var GoldenLeagueLen uint32
	binary.Read(buffer, binary.LittleEndian, &GoldenLeagueLen)
	if uint32(buffer.Len()) < GoldenLeagueLen*8 {
		return errors.New("message length error")
	}
	p.GoldenLeague = make(map[int32]T_FightBlance_GoldenLeague, GoldenLeagueLen)
	for i := uint32(0); i < GoldenLeagueLen; i++ {
		var k int32
		var v T_FightBlance_GoldenLeague
		binary.Read(buffer, binary.LittleEndian, &k)
		v.Decode(buffer)
		p.GoldenLeague[k] = v
	}
	if buffer.Len() < 4 {
		return errors.New("message length error")
	}
	var ActivityCooprationLen uint32
	binary.Read(buffer, binary.LittleEndian, &ActivityCooprationLen)
	p.ActivityCoopration = make(map[int32]T_FightBalance_Activity_CoopRation, ActivityCooprationLen)
	for i := uint32(0); i < ActivityCooprationLen; i++ {
		var k int32
		var v T_FightBalance_Activity_CoopRation
		binary.Read(buffer, binary.LittleEndian, &k)
		v.Decode(buffer)
		p.ActivityCoopration[k] = v
	}
	if buffer.Len() < 4 {
		return errors.New("message length error")
	}
	var ExtraDataLen uint32
	binary.Read(buffer, binary.LittleEndian, &ExtraDataLen)
	if uint32(buffer.Len()) < ExtraDataLen*8 {
		return errors.New("message length error")
	}
	p.ExtraData = make(map[int32]T_FightBalance_ExtraData, ExtraDataLen)
	for i := uint32(0); i < ExtraDataLen; i++ {
		var k int32
		var v T_FightBalance_ExtraData
		binary.Read(buffer, binary.LittleEndian, &k)
		v.Decode(buffer)
		p.ExtraData[k] = v
	}
	return nil
}

// 服务器返回数据结构-设置默认阵容
type S_Role_BattleArraySetDefine struct {
	Errorcode int32
}

func (p *S_Role_BattleArraySetDefine) Encode() []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.LittleEndian, p.Errorcode)
	return buffer.Bytes()
}

func (p *S_Role_BattleArraySetDefine) Decode(buffer *bytes.Buffer) error {
	if buffer.Len() < 4 {
		return errors.New("message length error")
	}
	binary.Read(buffer, binary.LittleEndian, &p.Errorcode)
	return nil
}

// 服务器返回数据结构-战车皮肤更换
type S_Role_Car_Skin_Change struct {
	Errorcode int32
}

func (p *S_Role_Car_Skin_Change) Encode() []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.LittleEndian, p.Errorcode)
	return buffer.Bytes()
}

func (p *S_Role_Car_Skin_Change) Decode(buffer *bytes.Buffer) error {
	if buffer.Len() < 4 {
		return errors.New("message length error")
	}
	binary.Read(buffer, binary.LittleEndian, &p.Errorcode)
	return nil
}

// 服务器返回数据结构-英雄皮肤更换
type S_Role_HeroChangeSkin struct {
	Errorcode int32
}

func (p *S_Role_HeroChangeSkin) Encode() []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.LittleEndian, p.Errorcode)
	return buffer.Bytes()
}

func (p *S_Role_HeroChangeSkin) Decode(buffer *bytes.Buffer) error {
	if buffer.Len() < 4 {
		return errors.New("message length error")
	}
	binary.Read(buffer, binary.LittleEndian, &p.Errorcode)
	return nil
}

// 服务器返回数据结构-阵容名称设置
type S_Role_SetBattleArrayName struct {
	Errorcode       int32
	BattleArrayName string
}

func (p *S_Role_SetBattleArrayName) Encode() []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.LittleEndian, p.Errorcode)
	binary.Write(buffer, binary.LittleEndian, uint32(len(p.BattleArrayName)))
	buffer.WriteString(p.BattleArrayName)
	return buffer.Bytes()
}

func (p *S_Role_SetBattleArrayName) Decode(buffer *bytes.Buffer) error {
	if buffer.Len() < 8 {
		return errors.New("message length error")
	}
	binary.Read(buffer, binary.LittleEndian, &p.Errorcode)
	var BattleArrayNameLen uint32
	binary.Read(buffer, binary.LittleEndian, &BattleArrayNameLen)
	if uint32(buffer.Len()) < BattleArrayNameLen {
		return errors.New("message length error")
	}
	p.BattleArrayName = string(buffer.Next(int(BattleArrayNameLen)))
	return nil
}

// 服务器返回数据结构-获取角色简要信息
type S_Role_GetRoleSimpleInfo struct {
	Errorcode       int32
	RoleAbstract    T_RoleAbstract
	RoleProficiency T_RoleProficiency
}

func (p *S_Role_GetRoleSimpleInfo) Encode() []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.LittleEndian, p.Errorcode)
	buffer.Write(p.RoleAbstract.Encode())
	buffer.Write(p.RoleProficiency.Encode())
	return buffer.Bytes()
}

func (p *S_Role_GetRoleSimpleInfo) Decode(buffer *bytes.Buffer) error {
	if buffer.Len() < 4 {
		return errors.New("message length error")
	}
	binary.Read(buffer, binary.LittleEndian, &p.Errorcode)
	p.RoleAbstract.Decode(buffer)
	p.RoleProficiency.Decode(buffer)
	return nil
}

type S_Role_SynChapterData struct {
	ScoreFundData T_Role_ChapterData
}

func (p *S_Role_SynChapterData) Encode() []byte {
	buffer := new(bytes.Buffer)
	buffer.Write(p.ScoreFundData.Encode())
	return buffer.Bytes()
}

func (p *S_Role_SynChapterData) Decode(buffer *bytes.Buffer) error {
	if err := p.ScoreFundData.Decode(buffer); err != nil {
		return err
	}
	return nil
}

type S_Role_SetGuide struct {
	Errorcode int32
}

func (p *S_Role_SetGuide) Encode() []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.LittleEndian, p.Errorcode)
	return buffer.Bytes()
}

func (p *S_Role_SetGuide) Decode(buffer *bytes.Buffer) error {
	if buffer.Len() < 4 {
		return errors.New("message length error")
	}
	binary.Read(buffer, binary.LittleEndian, &p.Errorcode)
	return nil
}
