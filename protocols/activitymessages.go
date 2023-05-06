package protocols

import (
	"bytes"
	"encoding/binary"
	"errors"
)

type T_Activity_Data struct {
	IsRuning   bool  `bson:"is_runing"`
	ActivityID int32 `bson:"activity_id"`
	BeginTime  int32 `bson:"begin_time"`
	EndTime    int32 `bson:"end_time"`
}

func (p *T_Activity_Data) Encode() []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.LittleEndian, p.IsRuning)
	binary.Write(buffer, binary.LittleEndian, p.ActivityID)
	binary.Write(buffer, binary.LittleEndian, p.BeginTime)
	binary.Write(buffer, binary.LittleEndian, p.EndTime)
	return buffer.Bytes()
}

func (p *T_Activity_Data) Decode(buffer *bytes.Buffer) error {
	if buffer.Len() < 13 {
		return errors.New("message length error")
	}
	binary.Read(buffer, binary.LittleEndian, &p.IsRuning)
	binary.Read(buffer, binary.LittleEndian, &p.ActivityID)
	binary.Read(buffer, binary.LittleEndian, &p.BeginTime)
	binary.Read(buffer, binary.LittleEndian, &p.EndTime)
	return nil
}

/***********************************  客户端  ***********************************/
type C_Activity_GetGreatSailingData struct {
	ActivityID int32
}

func (p *C_Activity_GetGreatSailingData) Encode() []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.LittleEndian, p.ActivityID)
	return buffer.Bytes()
}

func (p *C_Activity_GetGreatSailingData) Decode(buffer *bytes.Buffer, key uint8) error {
	if key != 0 {
		for i := 0; i < buffer.Len(); i++ {
			buffer.Bytes()[i] ^= byte(key)
		}
	}
	if buffer.Len() < 4 {
		return errors.New("message length error")
	}
	binary.Read(buffer, binary.LittleEndian, &p.ActivityID)
	return nil
}

type C_Activity_GreatSailingRefleshCard struct {
	RefleshCardNum int32
}

func (p *C_Activity_GreatSailingRefleshCard) Encode() []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.LittleEndian, p.RefleshCardNum)
	return buffer.Bytes()
}

func (p *C_Activity_GreatSailingRefleshCard) Decode(buffer *bytes.Buffer, key uint8) error {
	if key != 0 {
		for i := 0; i < buffer.Len(); i++ {
			buffer.Bytes()[i] ^= byte(key)
		}
	}
	if buffer.Len() < 4 {
		return errors.New("message length error")
	}
	binary.Read(buffer, binary.LittleEndian, &p.RefleshCardNum)
	return nil
}

/***********************************  服务端  ***********************************/
type S_Activity_SynAllActivityData struct {
	ActivityData map[int32]T_Activity_Data
}

func (p *S_Activity_SynAllActivityData) Encode() []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.LittleEndian, uint32(len(p.ActivityData)))
	for k, v := range p.ActivityData {
		binary.Write(buffer, binary.LittleEndian, k)
		buffer.Write(v.Encode())
	}
	return buffer.Bytes()
}

func (p *S_Activity_SynAllActivityData) Decode(buffer *bytes.Buffer) error {
	if buffer.Len() < 4 {
		return errors.New("message length error")
	}
	var ActivityDataLen uint32
	binary.Read(buffer, binary.LittleEndian, &ActivityDataLen)
	if uint32(buffer.Len()) < ActivityDataLen*17 {
		return errors.New("message length error")
	}
	p.ActivityData = make(map[int32]T_Activity_Data, ActivityDataLen)
	for i := uint32(0); i < ActivityDataLen; i++ {
		var k int32
		binary.Read(buffer, binary.LittleEndian, &k)
		var v T_Activity_Data
		if err := v.Decode(buffer); err != nil {
			return err
		}
		p.ActivityData[k] = v
	}
	return nil
}

type S_Activity_SyncEatChickenData struct {
	Error        int32
	ActivityID   int32
	Status       int32
	Count        int32
	RemainingNum int32
	AssembleNum  int32
	IsLive       bool
	Rank         int32
	NormalPrize  map[int32]bool
	Superprize   map[int32]bool
}

func (p *S_Activity_SyncEatChickenData) Encode() []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.LittleEndian, p.Error)
	binary.Write(buffer, binary.LittleEndian, p.ActivityID)
	binary.Write(buffer, binary.LittleEndian, p.Status)
	binary.Write(buffer, binary.LittleEndian, p.Count)
	binary.Write(buffer, binary.LittleEndian, p.RemainingNum)
	binary.Write(buffer, binary.LittleEndian, p.AssembleNum)
	binary.Write(buffer, binary.LittleEndian, p.IsLive)
	binary.Write(buffer, binary.LittleEndian, p.Rank)
	binary.Write(buffer, binary.LittleEndian, uint32(len(p.NormalPrize)))
	for k, v := range p.NormalPrize {
		binary.Write(buffer, binary.LittleEndian, k)
		binary.Write(buffer, binary.LittleEndian, v)
	}
	binary.Write(buffer, binary.LittleEndian, uint32(len(p.Superprize)))
	for k, v := range p.Superprize {
		binary.Write(buffer, binary.LittleEndian, k)
		binary.Write(buffer, binary.LittleEndian, v)
	}
	return buffer.Bytes()
}

func (p *S_Activity_SyncEatChickenData) Decode(buffer *bytes.Buffer) error {
	if buffer.Len() < 33 {
		return errors.New("message length error")
	}
	binary.Read(buffer, binary.LittleEndian, &p.Error)
	binary.Read(buffer, binary.LittleEndian, &p.ActivityID)
	binary.Read(buffer, binary.LittleEndian, &p.Status)
	binary.Read(buffer, binary.LittleEndian, &p.Count)
	binary.Read(buffer, binary.LittleEndian, &p.RemainingNum)
	binary.Read(buffer, binary.LittleEndian, &p.AssembleNum)
	binary.Read(buffer, binary.LittleEndian, &p.IsLive)
	binary.Read(buffer, binary.LittleEndian, &p.Rank)
	if buffer.Len() < 4 {
		return errors.New("message length error")
	}
	var NormalPrizeLen uint32
	binary.Read(buffer, binary.LittleEndian, &NormalPrizeLen)
	if uint32(buffer.Len()) < NormalPrizeLen*5 {
		return errors.New("message length error")
	}
	p.NormalPrize = make(map[int32]bool, NormalPrizeLen)
	for i := uint32(0); i < NormalPrizeLen; i++ {
		var k int32
		binary.Read(buffer, binary.LittleEndian, &k)
		var v bool
		binary.Read(buffer, binary.LittleEndian, &v)
		p.NormalPrize[k] = v
	}
	if buffer.Len() < 4 {
		return errors.New("message length error")
	}
	var SuperprizeLen uint32
	binary.Read(buffer, binary.LittleEndian, &SuperprizeLen)
	if uint32(buffer.Len()) < SuperprizeLen*5 {
		return errors.New("message length error")
	}
	p.Superprize = make(map[int32]bool, SuperprizeLen)
	for i := uint32(0); i < SuperprizeLen; i++ {
		var k int32
		binary.Read(buffer, binary.LittleEndian, &k)
		var v bool
		binary.Read(buffer, binary.LittleEndian, &v)
		p.Superprize[k] = v
	}
	return nil
}

type S_Activity_SyncGreatSailingData struct {
	Error               int32
	ActivityID          int32
	FailCount           int32
	IsOpen              bool
	HistoryMaxScore     int32
	TodayScore          int32
	DayFailNum          int32
	DayMatchNum         int32
	ContinuousWinNum    int32
	ContinuousFailNum   int32
	WinNum              int32
	ReliveNum           int32
	MaxContinuousWinNum int32
	CardId              int32
	PrizeReward         map[int32]bool
	RefleshCardNum      int32
}

func (p *S_Activity_SyncGreatSailingData) Encode() []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.LittleEndian, p.Error)
	binary.Write(buffer, binary.LittleEndian, p.ActivityID)
	binary.Write(buffer, binary.LittleEndian, p.FailCount)
	binary.Write(buffer, binary.LittleEndian, p.IsOpen)
	binary.Write(buffer, binary.LittleEndian, p.HistoryMaxScore)
	binary.Write(buffer, binary.LittleEndian, p.TodayScore)
	binary.Write(buffer, binary.LittleEndian, p.DayFailNum)
	binary.Write(buffer, binary.LittleEndian, p.DayMatchNum)
	binary.Write(buffer, binary.LittleEndian, p.ContinuousWinNum)
	binary.Write(buffer, binary.LittleEndian, p.ContinuousFailNum)
	binary.Write(buffer, binary.LittleEndian, p.WinNum)
	binary.Write(buffer, binary.LittleEndian, p.ReliveNum)
	binary.Write(buffer, binary.LittleEndian, p.MaxContinuousWinNum)
	binary.Write(buffer, binary.LittleEndian, p.CardId)
	binary.Write(buffer, binary.LittleEndian, uint32(len(p.PrizeReward)))
	for k, v := range p.PrizeReward {
		binary.Write(buffer, binary.LittleEndian, k)
		binary.Write(buffer, binary.LittleEndian, v)
	}
	binary.Write(buffer, binary.LittleEndian, p.RefleshCardNum)
	return buffer.Bytes()
}

func (p *S_Activity_SyncGreatSailingData) Decode(buffer *bytes.Buffer) error {
	if buffer.Len() < 57 {
		return errors.New("message length error")
	}
	binary.Read(buffer, binary.LittleEndian, &p.Error)
	binary.Read(buffer, binary.LittleEndian, &p.ActivityID)
	binary.Read(buffer, binary.LittleEndian, &p.FailCount)
	binary.Read(buffer, binary.LittleEndian, &p.IsOpen)
	binary.Read(buffer, binary.LittleEndian, &p.HistoryMaxScore)
	binary.Read(buffer, binary.LittleEndian, &p.TodayScore)
	binary.Read(buffer, binary.LittleEndian, &p.DayFailNum)
	binary.Read(buffer, binary.LittleEndian, &p.DayMatchNum)
	binary.Read(buffer, binary.LittleEndian, &p.ContinuousWinNum)
	binary.Read(buffer, binary.LittleEndian, &p.ContinuousFailNum)
	binary.Read(buffer, binary.LittleEndian, &p.WinNum)
	binary.Read(buffer, binary.LittleEndian, &p.ReliveNum)
	binary.Read(buffer, binary.LittleEndian, &p.MaxContinuousWinNum)
	binary.Read(buffer, binary.LittleEndian, &p.CardId)
	var PrizeRewardLen uint32
	binary.Read(buffer, binary.LittleEndian, &PrizeRewardLen)
	if uint32(buffer.Len()) < PrizeRewardLen*5 {
		return errors.New("message length error")
	}
	p.PrizeReward = make(map[int32]bool, PrizeRewardLen)
	for i := uint32(0); i < PrizeRewardLen; i++ {
		var k int32
		binary.Read(buffer, binary.LittleEndian, &k)
		var v bool
		binary.Read(buffer, binary.LittleEndian, &v)
		p.PrizeReward[k] = v
	}
	binary.Read(buffer, binary.LittleEndian, &p.RefleshCardNum)
	return nil
}

type S_Activity_SyncWeekCooperationData struct {
	Error               int32
	ActivityID          int32
	IsOpen              bool
	FailCount           int32
	Score               int32
	HistoryMaxScore     int32
	DayFailNum          int32
	DayMatchNum         int32
	ContinuousWinNum    int32
	ContinuousFailNum   int32
	WinNum              int32
	ReliveNum           int32
	MaxContinuousWinNum int32
	RefleshRobotId      int32
	Prize               map[int32]bool
}

func (p *S_Activity_SyncWeekCooperationData) Encode() []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.LittleEndian, p.Error)
	binary.Write(buffer, binary.LittleEndian, p.ActivityID)
	binary.Write(buffer, binary.LittleEndian, p.IsOpen)
	binary.Write(buffer, binary.LittleEndian, p.FailCount)
	binary.Write(buffer, binary.LittleEndian, p.Score)
	binary.Write(buffer, binary.LittleEndian, p.HistoryMaxScore)
	binary.Write(buffer, binary.LittleEndian, p.DayFailNum)
	binary.Write(buffer, binary.LittleEndian, p.DayMatchNum)
	binary.Write(buffer, binary.LittleEndian, p.ContinuousWinNum)
	binary.Write(buffer, binary.LittleEndian, p.ContinuousFailNum)
	binary.Write(buffer, binary.LittleEndian, p.WinNum)
	binary.Write(buffer, binary.LittleEndian, p.ReliveNum)
	binary.Write(buffer, binary.LittleEndian, p.MaxContinuousWinNum)
	binary.Write(buffer, binary.LittleEndian, p.RefleshRobotId)
	binary.Write(buffer, binary.LittleEndian, uint32(len(p.Prize)))
	for k, v := range p.Prize {
		binary.Write(buffer, binary.LittleEndian, k)
		binary.Write(buffer, binary.LittleEndian, v)
	}
	return buffer.Bytes()
}

func (p *S_Activity_SyncWeekCooperationData) Decode(buffer *bytes.Buffer) error {
	if buffer.Len() < 57 {
		return errors.New("message length error")
	}
	binary.Read(buffer, binary.LittleEndian, &p.Error)
	binary.Read(buffer, binary.LittleEndian, &p.ActivityID)
	binary.Read(buffer, binary.LittleEndian, &p.IsOpen)
	binary.Read(buffer, binary.LittleEndian, &p.FailCount)
	binary.Read(buffer, binary.LittleEndian, &p.Score)
	binary.Read(buffer, binary.LittleEndian, &p.HistoryMaxScore)
	binary.Read(buffer, binary.LittleEndian, &p.DayFailNum)
	binary.Read(buffer, binary.LittleEndian, &p.DayMatchNum)
	binary.Read(buffer, binary.LittleEndian, &p.ContinuousWinNum)
	binary.Read(buffer, binary.LittleEndian, &p.ContinuousFailNum)
	binary.Read(buffer, binary.LittleEndian, &p.WinNum)
	binary.Read(buffer, binary.LittleEndian, &p.ReliveNum)
	binary.Read(buffer, binary.LittleEndian, &p.MaxContinuousWinNum)
	binary.Read(buffer, binary.LittleEndian, &p.RefleshRobotId)
	var PrizeLen uint32
	binary.Read(buffer, binary.LittleEndian, &PrizeLen)
	if uint32(buffer.Len()) < PrizeLen*5 {
		return errors.New("message length error")
	}
	p.Prize = make(map[int32]bool, PrizeLen)
	for i := uint32(0); i < PrizeLen; i++ {
		var k int32
		binary.Read(buffer, binary.LittleEndian, &k)
		var v bool
		binary.Read(buffer, binary.LittleEndian, &v)
		p.Prize[k] = v
	}
	return nil
}

type S_Activity_SyncMachinariumData struct {
	Error               int32
	ActivityID          int32
	IsOpen              bool
	FailCount           int32
	Score               int32
	DayFailNum          int32
	DayMatchNum         int32
	ContinuousWinNum    int32
	ContinuousFailNum   int32
	WinNum              int32
	ReliveNum           int32
	MaxContinuousWinNum int32
	RefleshId           int32
	DayMaxRound         int32
	Prize               map[int32]bool
	NormalPrize         map[int32]bool
	Superprize          map[int32]bool
}

func (p *S_Activity_SyncMachinariumData) Encode() []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.LittleEndian, p.Error)
	binary.Write(buffer, binary.LittleEndian, p.ActivityID)
	binary.Write(buffer, binary.LittleEndian, p.IsOpen)
	binary.Write(buffer, binary.LittleEndian, p.FailCount)
	binary.Write(buffer, binary.LittleEndian, p.Score)
	binary.Write(buffer, binary.LittleEndian, p.DayFailNum)
	binary.Write(buffer, binary.LittleEndian, p.DayMatchNum)
	binary.Write(buffer, binary.LittleEndian, p.ContinuousWinNum)
	binary.Write(buffer, binary.LittleEndian, p.ContinuousFailNum)
	binary.Write(buffer, binary.LittleEndian, p.WinNum)
	binary.Write(buffer, binary.LittleEndian, p.ReliveNum)
	binary.Write(buffer, binary.LittleEndian, p.MaxContinuousWinNum)
	binary.Write(buffer, binary.LittleEndian, p.RefleshId)
	binary.Write(buffer, binary.LittleEndian, p.DayMaxRound)
	binary.Write(buffer, binary.LittleEndian, uint32(len(p.Prize)))
	for k, v := range p.Prize {
		binary.Write(buffer, binary.LittleEndian, k)
		binary.Write(buffer, binary.LittleEndian, v)
	}
	binary.Write(buffer, binary.LittleEndian, uint32(len(p.NormalPrize)))
	for k, v := range p.NormalPrize {
		binary.Write(buffer, binary.LittleEndian, k)
		binary.Write(buffer, binary.LittleEndian, v)
	}
	binary.Write(buffer, binary.LittleEndian, uint32(len(p.Superprize)))
	for k, v := range p.Superprize {
		binary.Write(buffer, binary.LittleEndian, k)
		binary.Write(buffer, binary.LittleEndian, v)
	}
	return buffer.Bytes()
}

func (p *S_Activity_SyncMachinariumData) Decode(buffer *bytes.Buffer) error {
	if buffer.Len() < 57 {
		return errors.New("message length error")
	}
	binary.Read(buffer, binary.LittleEndian, &p.Error)
	binary.Read(buffer, binary.LittleEndian, &p.ActivityID)
	binary.Read(buffer, binary.LittleEndian, &p.IsOpen)
	binary.Read(buffer, binary.LittleEndian, &p.FailCount)
	binary.Read(buffer, binary.LittleEndian, &p.Score)
	binary.Read(buffer, binary.LittleEndian, &p.DayFailNum)
	binary.Read(buffer, binary.LittleEndian, &p.DayMatchNum)
	binary.Read(buffer, binary.LittleEndian, &p.ContinuousWinNum)
	binary.Read(buffer, binary.LittleEndian, &p.ContinuousFailNum)
	binary.Read(buffer, binary.LittleEndian, &p.WinNum)
	binary.Read(buffer, binary.LittleEndian, &p.ReliveNum)
	binary.Read(buffer, binary.LittleEndian, &p.MaxContinuousWinNum)
	binary.Read(buffer, binary.LittleEndian, &p.RefleshId)
	binary.Read(buffer, binary.LittleEndian, &p.DayMaxRound)
	var PrizeLen uint32
	binary.Read(buffer, binary.LittleEndian, &PrizeLen)
	if uint32(buffer.Len()) < PrizeLen*5 {
		return errors.New("message length error")
	}
	p.Prize = make(map[int32]bool, PrizeLen)
	for i := uint32(0); i < PrizeLen; i++ {
		var k int32
		binary.Read(buffer, binary.LittleEndian, &k)
		var v bool
		binary.Read(buffer, binary.LittleEndian, &v)
		p.Prize[k] = v
	}
	var NormalPrizeLen uint32
	binary.Read(buffer, binary.LittleEndian, &NormalPrizeLen)
	if uint32(buffer.Len()) < NormalPrizeLen*5 {
		return errors.New("message length error")
	}
	p.Prize = make(map[int32]bool, NormalPrizeLen)
	for i := uint32(0); i < NormalPrizeLen; i++ {
		var k int32
		binary.Read(buffer, binary.LittleEndian, &k)
		var v bool
		binary.Read(buffer, binary.LittleEndian, &v)
		p.Prize[k] = v
	}
	var SuperprizeLen uint32
	binary.Read(buffer, binary.LittleEndian, &SuperprizeLen)
	if uint32(buffer.Len()) < SuperprizeLen*5 {
		return errors.New("message length error")
	}
	p.Prize = make(map[int32]bool, SuperprizeLen)
	for i := uint32(0); i < SuperprizeLen; i++ {
		var k int32
		binary.Read(buffer, binary.LittleEndian, &k)
		var v bool
		binary.Read(buffer, binary.LittleEndian, &v)
		p.Prize[k] = v
	}
	return nil
}
