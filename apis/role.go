package apis

import (
	"bytes"
	"tfjl-h5/constants"
	"tfjl-h5/core"
	"tfjl-h5/db"
	"tfjl-h5/iface"
	"tfjl-h5/models"
	"tfjl-h5/net"
	"tfjl-h5/protocols"
	"time"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
)

type RoleSynRoleDataRouter struct {
	net.BaseRouter
}

func (p *RoleSynRoleDataRouter) Handle(request iface.IRequest) {
	logrus.Info("***************************同步角色数据***************************")
	roleID, err := request.GetConnection().GetProperty("roleID")
	if err != nil {
		logrus.Error("GetProperty error:", err)
		return
	}
	player := core.WorldMgrObj.GetPlayerByPID(roleID.(int64))

	var cRoleSynRoleData = protocols.C_Role_SynRoleData{}
	cRoleSynRoleData.Decode(bytes.NewBuffer(request.GetData()), KEY)
	logrus.Info(cRoleSynRoleData)

	var sRoleSynRoleData = protocols.S_Role_SynRoleData{}
	sRoleSynRoleData.CurrTime = int32(time.Now().Unix())
	sRoleSynRoleData.RoleID = player.PID
	sRoleSynRoleData.StrID = player.ShowID
	sRoleSynRoleData.RoleName = player.Nickname
	sRoleSynRoleData.BIndulge = player.BIndulge
	sRoleSynRoleData.IndulgeTime = player.IndulgeTime
	sRoleSynRoleData.IndulgeDayOnlineTime = player.IndulgeDayOnlineTime
	// 同步角色属性
	var roleAttr []protocols.S_Role_SynRoleAttrValue
	db.DbManager.FindRoleAttrValueItems(bson.M{"role_id": roleID}, &roleAttr)
	logrus.Info(roleAttr)
	var roleAttrMap = make(map[int32]int32, len(roleAttr))
	for _, v := range roleAttr {
		roleAttrMap[v.Index] = v.Value
	}
	sRoleSynRoleData.RoleAttrValue = roleAttrMap
	sRoleSynRoleData.GameTime = protocols.T_Game_Time{
		Year:    int32(time.Now().Year()),
		Month:   int32(time.Now().Month()),
		Day:     int32(time.Now().Day()),
		Hour:    int32(time.Now().Hour()),
		Minnute: int32(time.Now().Minute()),
		Second:  int32(time.Now().Second()),
		WeedDay: int32(time.Now().Weekday()),
	}

	var roleInformation = db.DbManager.FindRoleInformationByRoleID(player.PID)
	roleInformation.FightData = protocols.T_Information_FightData{
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
	}

	var roleHeroSkin = []models.RoleHeroSkin{}
	db.DbManager.FindRoleHeroSkinByRoleID(player.PID, &roleHeroSkin)
	var roleHeroSkinMap = make(map[int64]int32, len(roleHeroSkin))
	for _, v := range roleHeroSkin {
		roleHeroSkinMap[v.UUID] = v.ID
	}
	roleInformation.HeroSkinMap = roleHeroSkinMap

	sRoleSynRoleData.Infomation = roleInformation

	sRoleSynRoleData.ClientData = protocols.T_Client_Data{
		IntMap: map[int32]int32{33: 0, 34: 0, 36: 0, 37: 55, 38: 77326, 39: 0, 40: 100, 41: 0, 42: 0, 43: 0, 44: 0, 45: 0, 46: 0, 47: 40, 48: 0, 49: 0, 50: 0, 51: 0, 52: 0, 53: 0, 54: 0, 55: 0, 56: 0, 57: 0, 58: 0, 59: 0, 60: 0, 61: 0, 62: 551, 63: 0, 64: 0, 65: 0, 66: 0, 67: 0, 68: 0, 69: 0, 70: 0, 71: 0, 72: 0, 73: 0, 74: 0},
	}

	sRoleSynRoleData.Recharge = protocols.T_Role_Recharge_Data{
		Recharges: map[int32]protocols.T_Role_Recharge_Single{},
	}

	var roleBagItems = []protocols.T_Role_Item{}
	err = db.DbManager.FindRoleBagItemsByRoleID(player.PID, &roleBagItems)
	if err != nil {
		logrus.Error("FindRoleBagItemsByRoleID err:", err)
		return
	}

	var roleBagItemsMap = make(map[int64]protocols.T_Role_Item, len(roleBagItems))
	for _, v := range roleBagItems {
		roleBagItemsMap[v.ItemUUID] = v
	}
	sRoleSynRoleData.RoleBag = protocols.T_Role_Bag{
		Items: roleBagItemsMap,
	}

	sRoleSynRoleData.RoleMail = protocols.T_Role_Mail{
		Mails: map[int64]protocols.T_Role_SingleMail{},
	}
	sRoleSynRoleData.RoleItemInfo = protocols.T_Role_ItemInfo{
		Day: map[int32]protocols.T_Role_ItemInfo_Day{},
	}

	// 任务items
	var tRoleSingleTask = []protocols.T_Role_SingleTask{}
	db.DbManager.FindRoleTaskItemsByRoleID(player.PID, &tRoleSingleTask)
	var tRoleSingleTaskMap = make(map[int32]protocols.T_Role_SingleTask, len(tRoleSingleTask))
	for _, v := range tRoleSingleTask {
		tRoleSingleTaskMap[v.TaskID] = v
	}
	sRoleSynRoleData.RoleTask = protocols.T_Role_Task{
		Tasks: tRoleSingleTaskMap,
		Extra: protocols.T_Task_Extra_Data{
			TaskGroup:          map[int32]protocols.T_Task_Group_Data{},
			RandTask:           map[int32]protocols.T_Task_Rand_Task_Data{1: {TaskID: 64003}, 2: {TaskID: 66001}, 3: {TaskID: 68002}},
			ReplaceRandTaskNum: 0,
			RandTaskHistory:    []protocols.T_Task_Rand_History_Data{},
		},
	}

	sRoleSynRoleData.Exchange = protocols.T_Role_ExchangeData{
		Groups: map[int32]protocols.T_Role_ExchangeGroup{},
	}

	var roleBattleArray []protocols.T_Role_BattleArrayIndexData
	db.DbManager.FindRoleBattleArray(bson.M{"role_id": player.PID}, &roleBattleArray)
	var roleBattleArrayMap = make(map[int32]protocols.T_Role_BattleArrayIDData, len(roleBattleArray))
	for _, v := range roleBattleArray {
		if roleBattleArrayIDData, ok := roleBattleArrayMap[v.ID]; ok {
			roleBattleArrayIDData.IndexData[v.Index] = v
		} else {
			roleBattleArrayMap[v.ID] = protocols.T_Role_BattleArrayIDData{
				IndexData:     map[int32]protocols.T_Role_BattleArrayIndexData{v.Index: v},
				RuneIndexData: map[int32]protocols.T_Role_BattleRuneIndexData{},
				BattleArray:   string(v.Index),
			}
		}
	}
	sRoleSynRoleData.Battlearray = protocols.T_Role_BattleArrayData{
		DefineID: 1, // 默认阵容
		IDData:   roleBattleArrayMap,
	}

	sRoleSynRoleData.Expressionarray = protocols.T_Role_ExpressionArrayData{
		ArrayData: map[int32]protocols.T_Role_ExpressionArrayIndexData{},
	}
	sRoleSynRoleData.ScoreAchievement = protocols.T_Role_ScoreAchievement{
		ScoreAchievements: map[int32]protocols.T_Role_ScoreAchievementSingle{},
	}
	sRoleSynRoleData.Signin = protocols.T_SignInData{
		Days: map[int32]protocols.T_SignInDayData{},
	}
	sRoleSynRoleData.Timebox = protocols.T_TimeBoxData{
		Boxs: map[int32]protocols.T_TimeBoxSingleData{},
	}
	sRoleSynRoleData.Themebox = protocols.T_ThemeBoxData{
		Entity: map[int32]protocols.T_ThemeBoxEntityData{},
	}

	// 赛季奖励数据
	var tSeasonEntityData = db.DbManager.FindRoleSeasonItemBySeasonID(player.PID, 32)
	var roleSeasonForeverScorePrize []models.RoleSeasonForeverScorePrize
	err = db.DbManager.FindRoleSeasonForeverScorePrizeByRoleID(player.PID, &roleSeasonForeverScorePrize)
	if err != nil {
		logrus.Error("FindRoleSeasonForeverScorePrizeByRoleID err:", err)
		return
	}
	var roleSeasonForeverScorePrizeMap = make(map[int32]protocols.T_SeasonForeverScorePrizeData, len(roleSeasonForeverScorePrize))
	for _, v := range roleSeasonForeverScorePrize {
		roleSeasonForeverScorePrizeMap[v.ID] = protocols.T_SeasonForeverScorePrizeData{
			IsPrize: v.IsPrize,
			IsExtra: v.IsExtra,
		}
	}
	tSeasonEntityData.ForeverScorePrize = roleSeasonForeverScorePrizeMap
	var roleSeasonScorePrize []models.RoleSeasonScorePrize
	err = db.DbManager.FindRoleSeasonScorePrizeByRoleID(player.PID, &roleSeasonScorePrize)
	if err != nil {
		logrus.Error("FindRoleSeasonScorePrizeByRoleID err:", err)
		return
	}
	var roleSeasonScorePrizeMap = make(map[int32]protocols.T_SeasonScorePrizeData, len(roleSeasonScorePrize))
	for _, v := range roleSeasonForeverScorePrize {
		roleSeasonScorePrizeMap[v.ID] = protocols.T_SeasonScorePrizeData{
			IsPrize: v.IsPrize,
			IsExtra: v.IsExtra,
		}
	}
	tSeasonEntityData.ScorePrize = roleSeasonScorePrizeMap
	sRoleSynRoleData.Season = protocols.T_SeasonData{
		Entity: map[int32]protocols.T_SeasonEntityData{0: tSeasonEntityData},
		Last:   map[int32]protocols.T_SeasonLastData{},
	}
	sRoleSynRoleData.Share = protocols.T_ShareData{
		Ranks:            map[int64]protocols.T_SharePlayerRankData{},
		Players:          map[int64]protocols.T_SharePlayerData{},
		DaySharePrizeNum: 0,
	}
	sRoleSynRoleData.Totalsignin = protocols.T_TotalSignInData{
		SigninRound:    1,
		SigninIndex:    1,
		SigninDay:      1,
		IsReceive:      false,
		IsExtraReceive: false,
	}
	sRoleSynRoleData.CDKdata = protocols.T_CDKData{
		Data: map[string]int32{},
	}
	sRoleSynRoleData.Watchadbox = protocols.T_TotalWatchADBox{
		WatchADRound:      1,
		WatchADIndex:      1,
		WatchADDay:        1,
		WatchADNum:        0,
		IsReceive:         false,
		IsExtraRecboolive: false,
	}
	sRoleSynRoleData.Halloffame = protocols.T_HallofFameData{
		Data: map[int32]protocols.T_HallofFameRoleData{},
	}
	sRoleSynRoleData.Condshare = protocols.T_CondShareData{
		Condshares: map[int32]protocols.T_CondShare(nil),
	}
	sRoleSynRoleData.Finalrune = protocols.T_FinalRuneData{
		Runes: map[int32]protocols.T_SingleFinalRuneData(nil),
	}
	sRoleSynRoleData.TimelockBox = protocols.T_TimeLockBoxData{
		Position:              map[int32]protocols.T_TimeLockBoxPositionData{},
		DayFreeFastReceiveNum: 0,
	}
	sRoleSynRoleData.ChapterData = protocols.T_Role_ChapterData{
		CurrentChapterId: 1,
		ChapterInfoMap: map[int32]protocols.T_Role_ChapterInfo{
			1: {
				ChapterId:         1,
				ChapterProgress:   0,
				RewardBoxStateMap: map[int32]int32{},
			},
			2: {
				ChapterId:         2,
				ChapterProgress:   0,
				RewardBoxStateMap: map[int32]int32{},
			},
			3: {
				ChapterId:         3,
				ChapterProgress:   0,
				RewardBoxStateMap: map[int32]int32{},
			},
			4: {
				ChapterId:         4,
				ChapterProgress:   0,
				RewardBoxStateMap: map[int32]int32{},
			},
			5: {
				ChapterId:         5,
				ChapterProgress:   0,
				RewardBoxStateMap: map[int32]int32{},
			},
			6: {
				ChapterId:         6,
				ChapterProgress:   0,
				RewardBoxStateMap: map[int32]int32{},
			},
			7: {
				ChapterId:         7,
				ChapterProgress:   0,
				RewardBoxStateMap: map[int32]int32{},
			},
			10: {
				ChapterId:         10,
				ChapterProgress:   0,
				RewardBoxStateMap: map[int32]int32{},
			},
			11: {
				ChapterId:         11,
				ChapterProgress:   0,
				RewardBoxStateMap: map[int32]int32{},
			},
			12: {
				ChapterId:         12,
				ChapterProgress:   0,
				RewardBoxStateMap: map[int32]int32{},
			},
			13: {
				ChapterId:         13,
				ChapterProgress:   0,
				RewardBoxStateMap: map[int32]int32{},
			},
			14: {
				ChapterId:         14,
				ChapterProgress:   0,
				RewardBoxStateMap: map[int32]int32{},
			},
			15: {
				ChapterId:         15,
				ChapterProgress:   0,
				RewardBoxStateMap: map[int32]int32{},
			},
			16: {
				ChapterId:         16,
				ChapterProgress:   0,
				RewardBoxStateMap: map[int32]int32{},
			},
			17: {
				ChapterId:         17,
				ChapterProgress:   0,
				RewardBoxStateMap: map[int32]int32{},
			},
			18: {
				ChapterId:         18,
				ChapterProgress:   0,
				RewardBoxStateMap: map[int32]int32{},
			},
			19: {
				ChapterId:         19,
				ChapterProgress:   0,
				RewardBoxStateMap: map[int32]int32{},
			},
			101: {
				ChapterId:         101,
				ChapterProgress:   0,
				RewardBoxStateMap: map[int32]int32{},
			},
			102: {
				ChapterId:         102,
				ChapterProgress:   0,
				RewardBoxStateMap: map[int32]int32{},
			},
			103: {
				ChapterId:         103,
				ChapterProgress:   0,
				RewardBoxStateMap: map[int32]int32{},
			},
			104: {
				ChapterId:         104,
				ChapterProgress:   0,
				RewardBoxStateMap: map[int32]int32{},
			},
			105: {
				ChapterId:         105,
				ChapterProgress:   0,
				RewardBoxStateMap: map[int32]int32{},
			},
			106: {
				ChapterId:         106,
				ChapterProgress:   0,
				RewardBoxStateMap: map[int32]int32{},
			},
			107: {
				ChapterId:         107,
				ChapterProgress:   0,
				RewardBoxStateMap: map[int32]int32{},
			},
			108: {
				ChapterId:         108,
				ChapterProgress:   0,
				RewardBoxStateMap: map[int32]int32{},
			},
			109: {
				ChapterId:         109,
				ChapterProgress:   0,
				RewardBoxStateMap: map[int32]int32{},
			},
			110: {
				ChapterId:         110,
				ChapterProgress:   0,
				RewardBoxStateMap: map[int32]int32{},
			},
			111: {
				ChapterId:         111,
				ChapterProgress:   0,
				RewardBoxStateMap: map[int32]int32{},
			},
			88888: {
				ChapterId:       0,
				ChapterProgress: 0,
				RewardBoxStateMap: map[int32]int32{
					0:     0,
					50:    0,
					100:   0,
					200:   0,
					300:   0,
					400:   0,
					500:   0,
					600:   0,
					800:   0,
					1000:  0,
					1200:  0,
					1600:  0,
					2000:  0,
					2500:  0,
					2800:  0,
					3000:  0,
					4000:  0,
					5000:  0,
					6000:  0,
					8000:  0,
					10000: 0,
					12000: 0,
					14000: 0,
					16000: 0,
					18000: 0,
					20000: 0,
					22000: 0,
				},
			},
			88889: {
				ChapterId:       0,
				ChapterProgress: 0,
				RewardBoxStateMap: map[int32]int32{
					500:   0,
					1000:  0,
					1500:  0,
					2000:  0,
					3000:  0,
					4000:  0,
					5000:  0,
					6000:  0,
					8000:  0,
					10000: 0,
				},
			},
			88890: {
				ChapterId:       0,
				ChapterProgress: 0,
				RewardBoxStateMap: map[int32]int32{
					500:   0,
					1000:  0,
					1500:  0,
					2000:  0,
					3000:  0,
					4000:  0,
					5000:  0,
					6000:  0,
					8000:  0,
					10000: 0,
				},
			},
		},
	}
	jsonData2 := []byte(`{"CurrentChapterId":1,"ChapterInfoMap":{"1":{"ChapterId":1,"ChapterProgress":0,"RewardBoxStateMap":{}},"10":{"ChapterId":10,"ChapterProgress":0,"RewardBoxStateMap":{}},"101":{"ChapterId":101,"ChapterProgress":0,"RewardBoxStateMap":{}},"102":{"ChapterId":102,"ChapterProgress":0,"RewardBoxStateMap":{}},"103":{"ChapterId":103,"ChapterProgress":0,"RewardBoxStateMap":{}},"104":{"ChapterId":104,"ChapterProgress":0,"RewardBoxStateMap":{}},"105":{"ChapterId":105,"ChapterProgress":0,"RewardBoxStateMap":{}},"106":{"ChapterId":106,"ChapterProgress":0,"RewardBoxStateMap":{}},"107":{"ChapterId":107,"ChapterProgress":0,"RewardBoxStateMap":{}},"108":{"ChapterId":108,"ChapterProgress":0,"RewardBoxStateMap":{}},"109":{"ChapterId":109,"ChapterProgress":0,"RewardBoxStateMap":{}},"11":{"ChapterId":11,"ChapterProgress":0,"RewardBoxStateMap":{}},"110":{"ChapterId":110,"ChapterProgress":0,"RewardBoxStateMap":{}},"111":{"ChapterId":111,"ChapterProgress":0,"RewardBoxStateMap":{}},"12":{"ChapterId":12,"ChapterProgress":0,"RewardBoxStateMap":{}},"13":{"ChapterId":13,"ChapterProgress":0,"RewardBoxStateMap":{}},"14":{"ChapterId":14,"ChapterProgress":0,"RewardBoxStateMap":{}},"15":{"ChapterId":15,"ChapterProgress":0,"RewardBoxStateMap":{}},"16":{"ChapterId":16,"ChapterProgress":0,"RewardBoxStateMap":{}},"17":{"ChapterId":17,"ChapterProgress":0,"RewardBoxStateMap":{}},"18":{"ChapterId":18,"ChapterProgress":0,"RewardBoxStateMap":{}},"19":{"ChapterId":19,"ChapterProgress":0,"RewardBoxStateMap":{}},"2":{"ChapterId":2,"ChapterProgress":0,"RewardBoxStateMap":{}},"3":{"ChapterId":3,"ChapterProgress":0,"RewardBoxStateMap":{}},"4":{"ChapterId":4,"ChapterProgress":0,"RewardBoxStateMap":{}},"5":{"ChapterId":5,"ChapterProgress":0,"RewardBoxStateMap":{}},"6":{"ChapterId":6,"ChapterProgress":0,"RewardBoxStateMap":{}},"7":{"ChapterId":7,"ChapterProgress":0,"RewardBoxStateMap":{}},"88888":{"ChapterId":0,"ChapterProgress":0,"RewardBoxStateMap":{"0":0,"100":1,"1000":1,"10000":0,"1200":1,"12000":0,"14000":0,"1600":1,"16000":0,"18000":0,"200":1,"2000":1,"20000":0,"22000":0,"2500":1,"2800":1,"300":1,"3000":1,"400":1,"4000":0,"50":1,"500":1,"5000":0,"600":1,"6000":0,"800":1,"8000":0}},"88889":{"ChapterId":0,"ChapterProgress":0,"RewardBoxStateMap":{"1000":1,"10000":1,"1500":1,"2000":1,"3000":1,"4000":1,"500":1,"5000":1,"6000":1,"8000":1}},"88890":{"ChapterId":0,"ChapterProgress":0,"RewardBoxStateMap":{"1000":0,"10000":0,"1500":0,"2000":0,"3000":0,"4000":0,"500":0,"5000":0,"6000":0,"8000":0}},"88989":{"ChapterId":0,"ChapterProgress":0,"RewardBoxStateMap":{"10500":0,"11000":0,"11500":0,"12000":0,"13000":0,"14000":0,"15000":0,"16000":0,"18000":0,"20000":0}},"88990":{"ChapterId":0,"ChapterProgress":0,"RewardBoxStateMap":{"10500":0,"11000":0,"11500":0,"12000":0,"13000":0,"14000":0,"15000":0,"16000":0,"18000":0,"20000":0}},"98889":{"ChapterId":0,"ChapterProgress":0,"RewardBoxStateMap":{"0":0,"100000":0,"1000000":0,"20000":0,"200000":0,"2000000":0,"5000":0,"50000":0,"500000":0}}}}`)
	var sRoleSynChapterData protocols.T_Role_ChapterData
	err = json.Unmarshal(jsonData2, &sRoleSynChapterData)
	if err != nil {
		logrus.Error(err)
		return
	}
	sRoleSynRoleData.ChapterData = sRoleSynChapterData

	sRoleSynRoleData.LegendData = protocols.T_Role_LegendData{
		ChapterInfoMap: map[int32]int32{
			0:     0,
			50:    0,
			100:   0,
			200:   0,
			300:   0,
			400:   0,
			500:   0,
			600:   0,
			800:   0,
			1000:  0,
			1200:  0,
			1600:  0,
			2000:  0,
			2500:  0,
			2800:  0,
			3000:  0,
			4000:  0,
			5000:  0,
			6000:  0,
			8000:  0,
			10000: 0,
			12000: 0,
			14000: 0,
			16000: 0,
			18000: 0,
			20000: 0,
			22000: 0,
		},
	}

	request.GetConnection().SendMessage(request.GetMsgType(), protocols.P_Role_SynRoleData, sRoleSynRoleData.Encode())

	var allActivityData []protocols.T_Activity_Data
	err = db.DbManager.FindActivitys(bson.M{}, &allActivityData)
	if err != nil {
		logrus.Errorln("FindActivitys error:", err)
		return
	}
	var allActivityDataMap = make(map[int32]protocols.T_Activity_Data, len(allActivityData))
	for _, v := range allActivityData {
		allActivityDataMap[v.ActivityID] = v
	}
	// 同步所有活动数据
	var sActivitySynAllActivityData = protocols.S_Activity_SynAllActivityData{ActivityData: allActivityDataMap}
	request.GetConnection().SendMessage(request.GetMsgType(), protocols.P_Activity_SynAllActivityData, sActivitySynAllActivityData.Encode())

	// 开关数据
	jsonData := []byte(`{"Onoff":{"1":true,"10":true,"11":false,"12":false,"13":true,"14":true,"15":true,"16":true,"17":true,"18":true,"19":true,"2":true,"20":true,"21":true,"22":true,"23":true,"24":false,"25":false,"26":false,"27":true,"28":false,"29":false,"3":true,"30":true,"31":true,"32":true,"33":true,"34":true,"35":true,"36":true,"37":true,"38":true,"39":true,"4":true,"40":true,"41":true,"42":true,"43":true,"44":true,"45":true,"46":true,"47":true,"48":true,"49":false,"5":true,"50":false,"51":true,"52":true,"53":true,"54":true,"55":true,"56":false,"57":true,"58":false,"59":true,"6":true,"60":true,"61":true,"62":false,"63":true,"64":true,"65":true,"66":true,"67":true,"68":true,"69":true,"7":true,"70":true,"71":true,"72":false,"73":false,"74":true,"75":true,"76":false,"77":true,"78":true,"79":true,"8":true,"80":false,"9":true}}`)
	var sRoleOnOffDataInfo protocols.S_Role_OnOffDataInfo
	err = json.Unmarshal(jsonData, &sRoleOnOffDataInfo)
	if err != nil {
		logrus.Error(err)
		return
	}
	request.GetConnection().SendMessage(request.GetMsgType(), protocols.P_Role_OnOffDataInfo, sRoleOnOffDataInfo.Encode())
}

type RoleBattleArrayUpRouter struct {
	net.BaseRouter
}

func (p *RoleBattleArrayUpRouter) Handle(request iface.IRequest) {
	logrus.Info("***********************************英雄阵容修改***********************************")
	roleID, err := request.GetConnection().GetProperty("roleID")
	if err != nil {
		logrus.Error("GetProperty error:", err)
		return
	}
	player := core.WorldMgrObj.GetPlayerByPID(roleID.(int64))

	var cRoleBattleArrayUp = protocols.C_Role_BattleArrayUp{}
	cRoleBattleArrayUp.Decode(bytes.NewBuffer(request.GetData()), player.Key)
	logrus.Infof("%#v", cRoleBattleArrayUp)

	db.DbManager.UpdateOneRoleBattleArrayByID(player.PID, cRoleBattleArrayUp.ArrayID, cRoleBattleArrayUp.ArrayIndex, cRoleBattleArrayUp.ItemUUID)

	var roleBattleArray []protocols.T_Role_BattleArrayIndexData
	db.DbManager.FindRoleBattleArray(bson.M{"role_id": player.PID}, &roleBattleArray)
	var roleBattleArrayMap = make(map[int32]protocols.T_Role_BattleArrayIDData, len(roleBattleArray))
	for _, v := range roleBattleArray {
		if roleBattleArrayIDData, ok := roleBattleArrayMap[v.ID]; ok {
			roleBattleArrayIDData.IndexData[v.Index] = v
		} else {
			roleBattleArrayMap[v.ID] = protocols.T_Role_BattleArrayIDData{
				IndexData:     map[int32]protocols.T_Role_BattleArrayIndexData{v.Index: v},
				RuneIndexData: map[int32]protocols.T_Role_BattleRuneIndexData{},
				BattleArray:   string(v.Index),
			}
		}
	}
	var battleArrayData protocols.S_Role_SynBattleArrayData
	battleArrayData.Battlearray = protocols.T_Role_BattleArrayData{
		DefineID: 1, // 默认阵容
		IDData:   roleBattleArrayMap,
	}
	request.GetConnection().SendMessage(request.GetMsgType(), protocols.P_Role_SynBattleArrayData, battleArrayData.Encode())
}

type RoleCarSkinChangeRouter struct {
	net.BaseRouter
}

func (p *RoleCarSkinChangeRouter) Handle(request iface.IRequest) {
	logrus.Info("*****************************修改车皮肤********************************")
	roleID, err := request.GetConnection().GetProperty("roleID")
	if err != nil {
		logrus.Error("GetProperty error:", err)
		return
	}
	player := core.WorldMgrObj.GetPlayerByPID(roleID.(int64))

	var cRoleCarSkinChange = protocols.C_Role_Car_Skin_Change{}
	cRoleCarSkinChange.Decode(bytes.NewBuffer(request.GetData()), player.Key)
	logrus.Infof("%#v", cRoleCarSkinChange)

	db.DbManager.UpdateRoleAttrValueByAttrID(player.PID, constants.ROLE_ATTR_CARSKINID, cRoleCarSkinChange.SkinId)
	var sRoleSynRoleAttrValue = protocols.S_Role_SynRoleAttrValue{
		Index: constants.ROLE_ATTR_CARSKINID,
		Value: cRoleCarSkinChange.SkinId,
	}
	request.GetConnection().SendMessage(request.GetMsgType(), protocols.P_Role_SynRoleAttrValue, sRoleSynRoleAttrValue.Encode())
	var sRoleCarSkinChange = protocols.S_Role_Car_Skin_Change{
		Errorcode: 0,
	}
	request.GetConnection().SendMessage(request.GetMsgType(), protocols.P_Role_Car_Skin_Change, sRoleCarSkinChange.Encode())
}

type RoleHeroChangeSkinRouter struct {
	net.BaseRouter
}

func (p *RoleHeroChangeSkinRouter) Handle(request iface.IRequest) {
	logrus.Info("***********************************修改英雄皮肤***********************************")
	roleID, err := request.GetConnection().GetProperty("roleID")
	if err != nil {
		logrus.Error("GetProperty error:", err)
		return
	}
	player := core.WorldMgrObj.GetPlayerByPID(roleID.(int64))

	var cRoleHeroChangeSkin = protocols.C_Role_HeroChangeSkin{}
	cRoleHeroChangeSkin.Decode(bytes.NewBuffer(request.GetData()), player.Key)
	db.DbManager.UpdateRoleHeroSkinByItemUUID(player.PID, cRoleHeroChangeSkin.HeroUUID, cRoleHeroChangeSkin.SkinId)

	// 同步角色信息
	var tInformationData = db.DbManager.FindRoleInformationByRoleID(player.PID)
	tInformationData.FightData = protocols.T_Information_FightData{
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
	}
	var roleHeroSkin = []models.RoleHeroSkin{}
	db.DbManager.FindRoleHeroSkinByRoleID(player.PID, &roleHeroSkin)
	var roleHeroSkinMap = make(map[int64]int32, len(roleHeroSkin))
	for _, v := range roleHeroSkin {
		roleHeroSkinMap[v.UUID] = v.ID
	}
	tInformationData.HeroSkinMap = roleHeroSkinMap
	request.GetConnection().SendMessage(request.GetMsgType(), protocols.P_Role_SynRoleInformationData, tInformationData.Encode())

	var sRoleHeroChanageSkin = protocols.S_Role_HeroChangeSkin{Errorcode: 0}
	request.GetConnection().SendMessage(request.GetMsgType(), protocols.P_Role_HeroChangeSkin, sRoleHeroChanageSkin.Encode())
}

type RoleGetRoleSimpleInfoRouter struct {
	net.BaseRouter
}

func (p *RoleGetRoleSimpleInfoRouter) Handle(request iface.IRequest) {
	logrus.Info("*******************************获取头像简要信息***********************************")
	roleID, err := request.GetConnection().GetProperty("roleID")
	if err != nil {
		logrus.Error("GetProperty error:", err)
		return
	}
	player := core.WorldMgrObj.GetPlayerByPID(roleID.(int64))

	// 获取简要数据
	var cRoleGetRoleSimpleInfo = protocols.C_Role_GetRoleSimpleInfo{}
	cRoleGetRoleSimpleInfo.Decode(bytes.NewBuffer(request.GetData()), player.Key)
	logrus.Infof("%#v", cRoleGetRoleSimpleInfo)

	var roleHeroSkin = []models.RoleHeroSkin{}
	db.DbManager.FindRoleHeroSkinByRoleID(player.PID, &roleHeroSkin)
	var roleHeroSkinMap = make(map[int64]int32, len(roleHeroSkin))
	for _, v := range roleHeroSkin {
		roleHeroSkinMap[v.UUID] = v.ID
	}
	var roleBagItems = []protocols.T_Role_Item{}
	if err := db.DbManager.FindRoleBagItemsByType(player.PID, 2, &roleBagItems); err != nil {
		logrus.Error("FindRoleBagItemsByType error:", err)
		return
	}
	var tHeroAbstract = make(map[int32]protocols.T_HeroAbstract, len(roleBagItems))
	for k, v := range roleBagItems {
		var skinID int32
		if value, ok := roleHeroSkinMap[v.ItemUUID]; ok {
			skinID = value
		}
		tHeroAbstract[int32(k)] = protocols.T_HeroAbstract{
			HeroUUID:  v.ItemUUID,
			HeroID:    v.ItemID,
			HeroLevel: 24,
			Attr:      []protocols.T_Attr{},
			SkinID:    skinID,
		}
	}
	carSkinAttrValueItem := db.DbManager.FindRoleAttrValueItemByAttrID(player.PID, constants.ROLE_ATTR_CARSKINID)
	carSkinBagItem := db.DbManager.FindCarSkinByItemID(player.PID, carSkinAttrValueItem.Value)
	// 战车皮肤
	var tRuneAbstract = protocols.T_RuneAbstract{ItemID: carSkinBagItem.ItemID*1000 + carSkinBagItem.ItemNum}
	var attrValue = db.DbManager.FindRoleAttrValueItemByAttrID(player.PID, 40)
	var sRoleGetRoleSimpleInfo = protocols.S_Role_GetRoleSimpleInfo{
		Errorcode: 0,
		RoleAbstract: protocols.T_RoleAbstract{
			RoleID:      player.PID,
			ShowID:      player.ShowID,
			BRobot:      false,
			NickName:    player.Nickname,
			Heros:       tHeroAbstract,
			Expressions: map[int32]protocols.T_ExpressionAbstract{},
			Runes: map[int32]protocols.T_RuneAbstract{
				40: {ItemID: attrValue.Value},
				constants.ROLE_ATTR_CARSKINID: tRuneAbstract},
			PetId: attrValue.Value,
		},
		RoleProficiency: protocols.T_RoleProficiency{},
	}
	request.GetConnection().SendMessage(request.GetMsgType(), protocols.P_Role_GetRoleSimpleInfo, sRoleGetRoleSimpleInfo.Encode())
}

type RoleSetGuideRouter struct {
	net.BaseRouter
}

func (p *RoleSetGuideRouter) Handle(request iface.IRequest) {
	logrus.Info("******************************* 设置 guide ***********************************")
	roleID, err := request.GetConnection().GetProperty("roleID")
	if err != nil {
		logrus.Error("GetProperty error:", err)
		return
	}
	player := core.WorldMgrObj.GetPlayerByPID(roleID.(int64))

	// 获取简要数据
	var cRoleSetGuide = protocols.C_Role_SetGuide{}
	cRoleSetGuide.Decode(bytes.NewBuffer(request.GetData()), player.Key)
	logrus.Infof("%#v", cRoleSetGuide)

	var sRoleSetGuide = protocols.S_Role_SetGuide{Errorcode: 0}
	updateResult, err := db.DbManager.UpdateRoleAttrValueByAttrID(player.PID, constants.ROLE_ATTR_GUIDE, cRoleSetGuide.Guide)
	if err != nil || updateResult.MatchedCount == 0 || updateResult.ModifiedCount == 0 {
		logrus.Error("UpdateRoleAttrValueByAttrID error:", err)
		sRoleSetGuide.Errorcode = 1
	} else {
		var sRoleSynRoleAttrValue = protocols.S_Role_SynRoleAttrValue{
			Index: constants.ROLE_ATTR_GUIDE,
			Value: cRoleSetGuide.Guide,
		}
		request.GetConnection().SendMessage(request.GetMsgType(), protocols.P_Role_SynRoleAttrValue, sRoleSynRoleAttrValue.Encode())
	}
	request.GetConnection().SendMessage(request.GetMsgType(), protocols.P_Role_SetGuide, sRoleSetGuide.Encode())
}
