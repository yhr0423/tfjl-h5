package apis

import (
	"bytes"
	"fmt"
	"math/rand"
	"tfjl-h5/core"
	"tfjl-h5/db"
	"tfjl-h5/iface"
	"tfjl-h5/models"
	"tfjl-h5/net"
	"tfjl-h5/protocols"
	"time"

	jsoniter "github.com/json-iterator/go"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

const (
	HEAD_LENGTH = 8
)

var KEY uint8
var FIGHT_KEY uint8

// 登录Ping路由
type LoginPingRouter struct {
	net.BaseRouter
}

func (p *LoginPingRouter) Handle(request iface.IRequest) {
	request.GetConnection().SendMessage(request.GetMsgType(), protocols.P_Login_Ping, []byte{})
}

// 登录验证路由
type LoginValidateOnlineRouter struct {
	net.BaseRouter
}

func (p *LoginValidateOnlineRouter) Handle(request iface.IRequest) {
	var cLoginValidateOnline = protocols.C_Login_ValidateOnline{}
	cLoginValidateOnline.Decode(bytes.NewBuffer(request.GetData()))
	logrus.Info(cLoginValidateOnline)

	role := db.DbManager.FindRoleByAccount(cLoginValidateOnline.AcountName)
	if role == (models.Role{}) {
		logrus.Error("role not found")
		return
	}
	player, err := core.NewPlayer(role, request.GetConnection())
	if err != nil {
		logrus.Error("NewPlayer error:", err)
		return
	}
	core.WorldMgrObj.AddPlayer(player)
	request.GetConnection().SetProperty("roleID", role.RoleID)
	request.GetConnection().SetProperty("key", role.Key)

	var sLoginValidate = protocols.S_Login_Validate{IsSucceed: true, ServerUTCTime: int32(time.Now().Unix()), Key: role.Key}
	request.GetConnection().SendMessage(request.GetMsgType(), protocols.P_Login_Validate, sLoginValidate.Encode())

	logrus.Info("========> Player PID =", player.PID, "is arrived <========")
}

// 登录请求角色路由
type LoginRequestRoleRouter struct {
	net.BaseRouter
}

func (p *LoginRequestRoleRouter) Handle(request iface.IRequest) {
	logrus.Info("请求角色")
	roleID, err := request.GetConnection().GetProperty("roleID")
	if err != nil {
		logrus.Error("GetProperty error:", err)
		return
	}
	player := core.WorldMgrObj.GetPlayerByPID(roleID.(int64))

	var sLoginRequestRole = protocols.S_Login_RequestRole{BIndulge: false, Roles: map[int64]protocols.T_Login_Role{
		player.PID: {
			RoleID: player.PID,
			ShowID: player.ShowID,
			Name:   player.Nickname,
			Level:  player.Level,
		},
	}, ForbidLoginTimeRemaining: player.ForbidLoginTimeRemaining}
	request.GetConnection().SendMessage(request.GetMsgType(), protocols.P_Login_RequestRole, sLoginRequestRole.Encode())
}

type LoginChooseRoleRouter struct {
	net.BaseRouter
}

func (p *LoginChooseRoleRouter) Handle(request iface.IRequest) {
	logrus.Info("选择角色")
	roleID, err := request.GetConnection().GetProperty("roleID")
	if err != nil {
		logrus.Error("GetProperty error:", err)
		return
	}
	player := core.WorldMgrObj.GetPlayerByPID(roleID.(int64))

	var cLoginChooseRole = protocols.C_Login_ChooseRole{}
	cLoginChooseRole.Decode(bytes.NewBuffer(request.GetData()), player.Key)
	logrus.Info(cLoginChooseRole)

	var sLoginChooseRole = protocols.S_Login_ChooseRole{Result: true}
	request.GetConnection().SendMessage(request.GetMsgType(), protocols.P_Login_ChooseRole, sLoginChooseRole.Encode())

	// 同步角色属性
	var roleAttr []protocols.S_Role_SynRoleAttrValue
	db.DbManager.FindRoleAttrValueItems(bson.M{"role_id": cLoginChooseRole.RoleID}, &roleAttr)
	for _, v := range roleAttr {
		request.GetConnection().SendMessage(request.GetMsgType(), protocols.P_Role_SynRoleAttrValue, v.Encode())
	}

	// 同步角色章节数据
	jsonData := []byte(`{"ScoreFundData":{"CurrentChapterId":1,"ChapterInfoMap":{"1":{"ChapterId":1,"ChapterProgress":0,"RewardBoxStateMap":{}},"10":{"ChapterId":10,"ChapterProgress":0,"RewardBoxStateMap":{}},"101":{"ChapterId":101,"ChapterProgress":0,"RewardBoxStateMap":{}},"102":{"ChapterId":102,"ChapterProgress":0,"RewardBoxStateMap":{}},"103":{"ChapterId":103,"ChapterProgress":0,"RewardBoxStateMap":{}},"104":{"ChapterId":104,"ChapterProgress":0,"RewardBoxStateMap":{}},"105":{"ChapterId":105,"ChapterProgress":0,"RewardBoxStateMap":{}},"106":{"ChapterId":106,"ChapterProgress":0,"RewardBoxStateMap":{}},"107":{"ChapterId":107,"ChapterProgress":0,"RewardBoxStateMap":{}},"108":{"ChapterId":108,"ChapterProgress":0,"RewardBoxStateMap":{}},"109":{"ChapterId":109,"ChapterProgress":0,"RewardBoxStateMap":{}},"11":{"ChapterId":11,"ChapterProgress":0,"RewardBoxStateMap":{}},"110":{"ChapterId":110,"ChapterProgress":0,"RewardBoxStateMap":{}},"111":{"ChapterId":111,"ChapterProgress":0,"RewardBoxStateMap":{}},"12":{"ChapterId":12,"ChapterProgress":0,"RewardBoxStateMap":{}},"13":{"ChapterId":13,"ChapterProgress":0,"RewardBoxStateMap":{}},"14":{"ChapterId":14,"ChapterProgress":0,"RewardBoxStateMap":{}},"15":{"ChapterId":15,"ChapterProgress":0,"RewardBoxStateMap":{}},"16":{"ChapterId":16,"ChapterProgress":0,"RewardBoxStateMap":{}},"17":{"ChapterId":17,"ChapterProgress":0,"RewardBoxStateMap":{}},"18":{"ChapterId":18,"ChapterProgress":0,"RewardBoxStateMap":{}},"19":{"ChapterId":19,"ChapterProgress":0,"RewardBoxStateMap":{}},"2":{"ChapterId":2,"ChapterProgress":0,"RewardBoxStateMap":{}},"3":{"ChapterId":3,"ChapterProgress":0,"RewardBoxStateMap":{}},"4":{"ChapterId":4,"ChapterProgress":0,"RewardBoxStateMap":{}},"5":{"ChapterId":5,"ChapterProgress":0,"RewardBoxStateMap":{}},"6":{"ChapterId":6,"ChapterProgress":0,"RewardBoxStateMap":{}},"7":{"ChapterId":7,"ChapterProgress":0,"RewardBoxStateMap":{}},"88888":{"ChapterId":0,"ChapterProgress":0,"RewardBoxStateMap":{"0":0,"100":1,"1000":1,"10000":0,"1200":1,"12000":0,"14000":0,"1600":1,"16000":0,"18000":0,"200":1,"2000":1,"20000":0,"22000":0,"2500":1,"2800":1,"300":1,"3000":1,"400":1,"4000":0,"50":1,"500":1,"5000":0,"600":1,"6000":0,"800":1,"8000":0}},"88889":{"ChapterId":0,"ChapterProgress":0,"RewardBoxStateMap":{"1000":1,"10000":1,"1500":1,"2000":1,"3000":1,"4000":1,"500":1,"5000":1,"6000":1,"8000":1}},"88890":{"ChapterId":0,"ChapterProgress":0,"RewardBoxStateMap":{"1000":0,"10000":0,"1500":0,"2000":0,"3000":0,"4000":0,"500":0,"5000":0,"6000":0,"8000":0}},"88989":{"ChapterId":0,"ChapterProgress":0,"RewardBoxStateMap":{"10500":0,"11000":0,"11500":0,"12000":0,"13000":0,"14000":0,"15000":0,"16000":0,"18000":0,"20000":0}},"88990":{"ChapterId":0,"ChapterProgress":0,"RewardBoxStateMap":{"10500":0,"11000":0,"11500":0,"12000":0,"13000":0,"14000":0,"15000":0,"16000":0,"18000":0,"20000":0}},"98889":{"ChapterId":0,"ChapterProgress":0,"RewardBoxStateMap":{"0":0,"100000":0,"1000000":0,"20000":0,"200000":0,"2000000":0,"5000":0,"50000":0,"500000":0}}}}}`)
	var sRoleSynChapterData protocols.S_Role_SynChapterData
	err = json.Unmarshal(jsonData, &sRoleSynChapterData)
	if err != nil {
		logrus.Error(err)
		return
	}
	request.GetConnection().SendMessage(request.GetMsgType(), protocols.P_Role_SynChapterData, sRoleSynChapterData.Encode())

	// 同步角色信息
	var tInformationData = protocols.T_Information_Data{
		FightData: protocols.T_Information_FightData{
			TypeData: map[int32]protocols.T_Information_FightTypeData{
				1:  {MaxRound: 310, WinNum: 999, LostNum: 0, TotalWinNum: 999, TotalLostNum: 0, SeriesWinNum: 999, SeriesLostNum: 0, WinLostResetNum: 0, AdditionalDayNum: 0},
				2:  {MaxRound: 310, WinNum: 999, LostNum: 0, TotalWinNum: 999, TotalLostNum: 0, SeriesWinNum: 999, SeriesLostNum: 0, WinLostResetNum: 0, AdditionalDayNum: 0}, // 合作
				3:  {MaxRound: 310, WinNum: 999, LostNum: 0, TotalWinNum: 999, TotalLostNum: 0, SeriesWinNum: 999, SeriesLostNum: 0, WinLostResetNum: 0, AdditionalDayNum: 0},
				7:  {MaxRound: 310, WinNum: 0, LostNum: 0, TotalWinNum: 0, TotalLostNum: 0, SeriesWinNum: 0, SeriesLostNum: 0, WinLostResetNum: 0, AdditionalDayNum: 0},
				8:  {MaxRound: 310, WinNum: 0, LostNum: 0, TotalWinNum: 0, TotalLostNum: 0, SeriesWinNum: 0, SeriesLostNum: 0, WinLostResetNum: 0, AdditionalDayNum: 0},
				9:  {MaxRound: 310, WinNum: 0, LostNum: 0, TotalWinNum: 0, TotalLostNum: 0, SeriesWinNum: 0, SeriesLostNum: 0, WinLostResetNum: 0, AdditionalDayNum: 0},
				10: {MaxRound: 200, WinNum: 999, LostNum: 0, TotalWinNum: 999, TotalLostNum: 0, SeriesWinNum: 0, SeriesLostNum: 0, WinLostResetNum: 0, AdditionalDayNum: 0}, // 大航海
				11: {MaxRound: 0, WinNum: 0, LostNum: 0, TotalWinNum: 0, TotalLostNum: 0, SeriesWinNum: 0, SeriesLostNum: 0, WinLostResetNum: 0, AdditionalDayNum: 0},
				12: {MaxRound: 130, WinNum: 999, LostNum: 0, TotalWinNum: 999, TotalLostNum: 0, SeriesWinNum: 0, SeriesLostNum: 0, WinLostResetNum: 0, AdditionalDayNum: 0}, // 寒冰堡
				13: {MaxRound: 310, WinNum: 999, LostNum: 0, TotalWinNum: 999, TotalLostNum: 0, SeriesWinNum: 1, SeriesLostNum: 0, WinLostResetNum: 0, AdditionalDayNum: 0},
				14: {MaxRound: 310, WinNum: 999, LostNum: 0, TotalWinNum: 999, TotalLostNum: 0, SeriesWinNum: 0, SeriesLostNum: 0, WinLostResetNum: 0, AdditionalDayNum: 0},
				15: {MaxRound: 200, WinNum: 999, LostNum: 0, TotalWinNum: 999, TotalLostNum: 0, SeriesWinNum: 0, SeriesLostNum: 0, WinLostResetNum: 0, AdditionalDayNum: 0}, // 机械迷城
			},
		},
	}
	request.GetConnection().SendMessage(request.GetMsgType(), protocols.P_Role_SynRoleInformationData, tInformationData.Encode())

	var tTaskExtraData = protocols.T_Task_Extra_Data{
		TaskGroup: map[int32]protocols.T_Task_Group_Data{
			1: {
				Boxs: map[int32]protocols.T_Task_Box_Data{
					1: {
						BTake: true,
					},
				},
			},
		},
		RandTask: map[int32]protocols.T_Task_Rand_Task_Data{
			1: {
				TaskID: 1,
			},
		},
		ReplaceRandTaskNum: 0,
		RandTaskHistory: []protocols.T_Task_Rand_History_Data{
			{
				Task: map[int32]protocols.T_Task_Rand_History_Task_Data{
					1: {},
				},
			},
		},
	}
	request.GetConnection().SendMessage(request.GetMsgType(), protocols.P_Role_SynTaskExtraData, tTaskExtraData.Encode())

	var sRoleSynTaskData = protocols.S_Role_SynTaskData{
		ChangeTask: []protocols.T_Role_SingleTask{
			{
				TaskID:     1,
				TaskState:  1,
				TaskCount:  1,
				TaskCDTime: 1,
				ExtraState: 1,
			},
		},
		DeleteTask: []int32{1},
	}
	request.GetConnection().SendMessage(request.GetMsgType(), protocols.P_Role_SynTaskData, sRoleSynTaskData.Encode())

	var sRoleTotalWatchADBoxData = protocols.S_Role_TotalWatchADBoxData{
		Totalwatchadbox: protocols.T_TotalWatchADBox{
			WatchADRound:      1,
			WatchADIndex:      1,
			WatchADDay:        1,
			WatchADNum:        1,
			IsReceive:         false,
			IsExtraRecboolive: false,
		},
	}
	request.GetConnection().SendMessage(request.GetMsgType(), protocols.P_Role_TotalWatchADBoxData, sRoleTotalWatchADBoxData.Encode())

	var costGetMap = make(map[int32]int32)
	costGetMap[1] = 0
	costGetMap[2] = 0
	costGetMap[3] = 0
	costGetMap[4] = 0
	costGetMap[5] = 0
	costGetMap[6] = 0
	costGetMap[7] = 0
	costGetMap[8] = 0
	costGetMap[9] = 0
	costGetMap[101] = 0
	costGetMap[102] = 0
	costGetMap[103] = 0
	costGetMap[201] = 0
	costGetMap[202] = 0
	costGetMap[203] = 0
	costGetMap[204] = 0
	costGetMap[205] = 0
	costGetMap[206] = 0
	costGetMap[207] = 0
	costGetMap[208] = 0
	costGetMap[209] = 0
	costGetMap[210] = 0
	costGetMap[211] = 0
	costGetMap[212] = 0
	costGetMap[213] = 0
	costGetMap[214] = 0
	costGetMap[1001] = 0
	costGetMap[1002] = 0
	costGetMap[1003] = 0
	costGetMap[1004] = 0
	var sRoleSyncCostGet = protocols.S_Role_SyncCostGet{CostGetMap: costGetMap}
	request.GetConnection().SendMessage(request.GetMsgType(), protocols.P_Role_SyncCostGet, sRoleSyncCostGet.Encode())

	// 同步活动数据
	// 试炼场
	jsonData = []byte(`{"Error":0,"ActivityID":3000,"Status":1,"Count":0,"RemainingNum":0,"AssembleNum":0,"IsLive":true,"Rank":0,"NormalPrize":{"1":true,"10":true,"11":true,"12":true,"13":true,"14":true,"15":true,"16":true,"17":true,"18":true,"19":true,"2":true,"20":true,"21":true,"22":true,"23":true,"24":true,"25":true,"26":true,"27":true,"28":true,"29":true,"3":true,"30":true,"31":true,"32":true,"33":true,"34":true,"35":true,"36":true,"37":true,"38":true,"39":true,"4":true,"40":true,"41":true,"42":true,"43":true,"44":true,"45":true,"5":true,"6":true,"7":true,"8":true,"9":true},"Superprize":{"1":true,"10":true,"11":true,"12":true,"13":true,"14":true,"15":true,"16":true,"17":true,"18":true,"19":true,"2":true,"20":true,"21":true,"22":true,"23":true,"24":true,"25":true,"26":true,"27":true,"28":true,"29":true,"3":true,"30":true,"31":true,"32":true,"33":true,"34":true,"35":true,"36":true,"37":true,"38":true,"39":true,"4":true,"40":true,"41":true,"42":true,"43":true,"44":true,"45":true,"5":true,"6":true,"7":true,"8":true,"9":true}}`)
	var sActivitySyncEatChickenData protocols.S_Activity_SyncEatChickenData
	err = json.Unmarshal(jsonData, &sActivitySyncEatChickenData)
	if err != nil {
		logrus.Error(err)
		return
	}
	request.GetConnection().SendMessage(request.GetMsgType(), protocols.P_Activity_SyncEatChickenData, sActivitySyncEatChickenData.Encode())

	// 大航海
	rand.Seed(time.Now().Unix())
	cardID := rand.Intn(135) + 1
	jsonData = []byte(fmt.Sprintf(`{"Error":0,"ActivityID":5000,"FailCount":0,"IsOpen":true,"HistoryMaxScore":200,"TodayScore":200,"DayFailNum":0,"DayMatchNum":0,"ContinuousWinNum":0,"ContinuousFailNum":0,"WinNum":0,"ReliveNum":0,"MaxContinuousWinNum":0,"CardId":%d,"PrizeReward":{},"RefleshCardNum":0}`, cardID))
	var sActivitySyncGreatSailingData protocols.S_Activity_SyncGreatSailingData
	err = json.Unmarshal(jsonData, &sActivitySyncGreatSailingData)
	if err != nil {
		logrus.Error(err)
		return
	}
	request.GetConnection().SendMessage(request.GetMsgType(), protocols.P_Activity_SyncGreatSailingData, sActivitySyncGreatSailingData.Encode())

	// 寒冰堡
	jsonData = []byte(`{"Error":0,"ActivityID":7000,"IsOpen":true,"FailCount":0,"Score":0,"HistoryMaxScore":0,"DayFailNum":0,"DayMatchNum":0,"ContinuousWinNum":0,"ContinuousFailNum":0,"WinNum":3,"ReliveNum":0,"MaxContinuousWinNum":130,"RefleshRobotId":1,"Prize":{}}`)
	var sActivitySyncWeekCooperationData protocols.S_Activity_SyncWeekCooperationData
	err = json.Unmarshal(jsonData, &sActivitySyncWeekCooperationData)
	if err != nil {
		logrus.Error(err)
		return
	}
	request.GetConnection().SendMessage(request.GetMsgType(), protocols.P_Activity_SyncWeekCooperationData, sActivitySyncWeekCooperationData.Encode())

	// 雾隐
	jsonData = []byte(`{"Error":0,"ActivityID":10000,"IsOpen":false,"FailCount":0,"Score":0,"DayFailNum":0,"DayMatchNum":0,"ContinuousWinNum":0,"ContinuousFailNum":0,"WinNum":26,"ReliveNum":5,"MaxContinuousWinNum":0,"RefleshId":0,"Prize":{},"NormalPrize":null,"Superprize":null}`)
	var sActivitySyncFogHiddenData protocols.S_Activity_SyncFogHiddenData
	err = json.Unmarshal(jsonData, &sActivitySyncFogHiddenData)
	if err != nil {
		logrus.Error(err)
		return
	}
	request.GetConnection().SendMessage(request.GetMsgType(), protocols.P_Activity_SyncFogHiddenData, sActivitySyncFogHiddenData.Encode())

	// 机械迷城数据
	jsonData = []byte(`{"Error":0,"ActivityID":11000,"IsOpen":false,"FailCount":0,"Score":0,"DayFailNum":0,"DayMatchNum":0,"ContinuousWinNum":0,"ContinuousFailNum":0,"WinNum":999,"ReliveNum":0,"MaxContinuousWinNum":0,"RefleshId":1,"DayMaxRound":0,"Prize":{},"NormalPrize":null,"Superprize":null}`)
	var sActivitySyncMachinariumData protocols.S_Activity_SyncMachinariumData
	err = json.Unmarshal(jsonData, &sActivitySyncMachinariumData)
	if err != nil {
		logrus.Error(err)
		return
	}
	request.GetConnection().SendMessage(request.GetMsgType(), protocols.P_Activity_SyncMachinariumData, sActivitySyncMachinariumData.Encode())

	var sRoleRoleEnterLogic = protocols.S_Role_RoleEnterLogic{}
	request.GetConnection().SendMessage(request.GetMsgType(), protocols.P_Role_RoleEnterLogic, sRoleRoleEnterLogic.Encode())

	// 联盟数据
	jsonData = []byte(`{"PlayerData":{"ActiveLeaveSociaty":1,"PassiveLeaveSociaty":1,"SeriesSignInNum":0,"DaySignInSociatyID":"","DayReceivePrizeNum":0,"SociatyID":"12345","SociatyName":"塔防精灵联盟","SociatyLevel":10,"SociatyFlag":1,"Job":1,"Contribution":99999,"Donate":{"1":{"DayNum":0},"2":{"DayNum":0}},"DayConvertSociatyMedal":0,"RedEnvelopes":[],"DurationTimes":[]}}`)
	var sSociatySynData protocols.S_Sociaty_SynData
	err = json.Unmarshal(jsonData, &sSociatySynData)
	if err != nil {
		logrus.Error(err)
		return
	}
	request.GetConnection().SendMessage(request.GetMsgType(), protocols.P_Sociaty_SynData, sSociatySynData.Encode())
}
