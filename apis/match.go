package apis

import (
	"bytes"
	"fmt"
	"tfjl-h5/constants"
	"tfjl-h5/core"
	"tfjl-h5/db"
	"tfjl-h5/iface"
	"tfjl-h5/models"
	"tfjl-h5/net"
	"tfjl-h5/protocols"
	"time"

	"tfjl-h5/utils"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
)

// 快速匹配-默认机器人
type MatchFightRouter struct {
	net.BaseRouter
}

func (p *MatchFightRouter) Handle(request iface.IRequest) {
	logrus.Info("************************* 快速匹配 *************************")
	roleID, err := request.GetConnection().GetProperty("roleID")
	if err != nil {
		logrus.Error("GetProperty error:", err)
		return
	}
	player := core.WorldMgrObj.GetPlayerByPID(roleID.(int64))
	var fightData protocols.C_Match_Fight
	fightData.Decode(bytes.NewBuffer(request.GetData()), player.Key)
	logrus.Info("fightData: ", fightData)

	// 快速匹配结果
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

	var roleHeroSkin = []models.RoleHeroSkin{}
	db.DbManager.FindRoleHeroSkinByRoleID(player.PID, &roleHeroSkin)
	var roleHeroSkinMap = make(map[int64]int32, len(roleHeroSkin))
	for _, v := range roleHeroSkin {
		roleHeroSkinMap[v.UUID] = v.ID
	}
	var roleBagItems = []protocols.T_Role_Item{}
	logrus.Error(db.DbManager.FindRoleBagItemsByType(player.PID, 2, &roleBagItems))
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

	carSkinAttrValueItem := db.DbManager.FindRoleAttrValueItemByAttrID(player.PID, 48)
	carSkinBagItem := db.DbManager.FindCarSkinByItemID(player.PID, carSkinAttrValueItem.Value)
	// 战车皮肤
	var tRuneAbstract = protocols.T_RuneAbstract{ItemID: carSkinBagItem.ItemID*1000 + carSkinBagItem.ItemNum}
	var attrValue = db.DbManager.FindRoleAttrValueItemByAttrID(player.PID, 40)

	var matchResult = protocols.S_Match_Result{
		Ret:          0,
		FightPattern: fightData.FightType,
		FightType:    fightData.FightType,
		FightModule:  fightData.FightType,
		RoleIdentity: fightData.FightType,
		MatchRoles: map[int64]protocols.T_RoleAbstract{
			1: {
				RoleID:   1,
				ShowID:   "1",
				BRobot:   true,
				NickName: player.Nickname,
				FightType: map[int32]protocols.T_RoleFightTypeAbstract{
					1: {MaxRound: 0, WinNum: 0, LostNum: 0, SeriesWinNum: 0},
				},
				Heros:           tHeroAbstract,
				Expressions:     map[int32]protocols.T_ExpressionAbstract{},
				Runes:           map[int32]protocols.T_RuneAbstract{48: tRuneAbstract},
				Finalrunes:      map[int32]protocols.T_FinalRuneAbstract{},
				FightSeasonData: map[int32]protocols.T_Fight_SeasonData{},
				PetId:           attrValue.Value,
			},
			player.PID: {
				RoleID:   player.PID,
				ShowID:   player.ShowID,
				BRobot:   false,
				NickName: player.Nickname,
				FightType: map[int32]protocols.T_RoleFightTypeAbstract{
					1:  {MaxRound: 999, WinNum: 999, LostNum: 0, SeriesWinNum: 999},
					2:  {MaxRound: 0, WinNum: 0, LostNum: 0, SeriesWinNum: 0},
					7:  {MaxRound: 0, WinNum: 0, LostNum: 0, SeriesWinNum: 0},
					8:  {MaxRound: 0, WinNum: 0, LostNum: 0, SeriesWinNum: 0},
					9:  {MaxRound: 0, WinNum: 0, LostNum: 0, SeriesWinNum: 0},
					10: {MaxRound: 0, WinNum: 0, LostNum: 0, SeriesWinNum: 0},
					11: {MaxRound: 0, WinNum: 0, LostNum: 0, SeriesWinNum: 0},
					12: {MaxRound: 0, WinNum: 0, LostNum: 0, SeriesWinNum: 0},
					14: {MaxRound: 0, WinNum: 0, LostNum: 0, SeriesWinNum: 0},
					15: {MaxRound: 0, WinNum: 0, LostNum: 0, SeriesWinNum: 0}},
				Heros:       tHeroAbstract,
				Expressions: map[int32]protocols.T_ExpressionAbstract{},
				Runes: map[int32]protocols.T_RuneAbstract{
					40: {ItemID: attrValue.Value},
					48: tRuneAbstract},
				Finalrunes:      map[int32]protocols.T_FinalRuneAbstract{},
				FightSeasonData: map[int32]protocols.T_Fight_SeasonData{},
				PetId:           attrValue.Value,
			},
		},
		FightServerIp:      "",
		FightServerPort:    0,
		FightServerIpSSL:   "",
		FightServerPortSSL: 0,
		FightToken:         "",
	}

	if fightData.FightType == constants.FIGHT_TYPE_BATTLE {
		// 对战
		matchResult.SeedId = []int32{625, 867, 296, 299, 75, 115, 294, 268, 235, 775, 346, 588, 849, 223, 161, 987, 158, 917, 509, 498, 604, 529, 162, 419, 636, 730, 844, 659, 783, 312, 265, 987, 315, 724, 184, 390, 448, 982, 327, 709, 280, 696, 932, 821, 400, 567, 380, 63, 593, 504, 69, 741, 206, 708, 724, 227, 576, 859, 201, 710, 844, 556, 424, 22, 545, 526, 528, 717, 186, 143, 82, 813, 465, 656, 31, 575, 435, 343, 558, 640, 639, 83, 35, 471, 710, 142, 170, 658, 594, 390, 609, 949, 773, 965, 57, 852, 855, 288, 385, 81}
		matchResult.ExtraData = map[int64]protocols.T_Fight_Extra_Data{}
		matchResult.BossIdIndexs = []int32{625, 867, 296, 299, 75, 115, 294, 268, 235, 775, 346, 588, 849, 223, 161, 987, 158, 917, 509, 498, 604, 529, 162, 419, 636, 730, 844, 659, 783, 312, 265, 987, 315, 724, 184, 390, 448, 982, 327, 709, 280, 696, 932, 821, 400, 567, 380, 63, 593, 504, 69, 741, 206, 708, 724, 227, 576, 859, 201, 710, 844, 556, 424, 22, 545, 526, 528, 717, 186, 143, 82, 813, 465, 656, 31, 575, 435, 343, 558, 640, 639, 83, 35, 471, 710, 142, 170, 658, 594, 390, 609, 949, 773, 965, 57, 852, 855, 288, 385, 81}
	} else if fightData.FightType == constants.FIGHT_TYPE_COOPERATION {
		// 合作
		matchResult.SeedId = []int32{3, 2, 1, 3, 7, 4, 6, 3, 2, 3, 3, 9, 4, 2, 6, 8, 6, 9, 3, 1, 9, 2, 8, 5, 3, 5, 6, 2, 7, 6, 4}
		matchResult.ExtraData = map[int64]protocols.T_Fight_Extra_Data{}
		matchResult.BossIdIndexs = []int32{3, 2, 1, 3, 7, 4, 6, 3, 2, 3, 3, 9, 4, 2, 6, 8, 6, 9, 3, 1, 9, 2, 8, 5, 3, 5, 6, 2, 7, 6, 4}
	} else if fightData.FightType == constants.FIGHT_TYPE_BATTLE_GREAT_SAILING {
		// 大航海
		matchResult.SeedId = []int32{3, 5, 6, 2, 5, 1, 7, 5, 9, 1, 4, 2, 1, 1, 6, 4, 7, 1, 4, 9, 1, 4, 7, 6, 1, 9, 6, 2, 3, 8, 3}
		matchResult.ExtraData = map[int64]protocols.T_Fight_Extra_Data{}
		matchResult.BossIdIndexs = []int32{3, 5, 6, 2, 5, 1, 7, 5, 9, 1, 4, 2, 1, 1, 6, 4, 7, 1, 4, 9, 1, 4, 7, 6, 1, 9, 6, 2, 3, 8, 3}
	} else if fightData.FightType == constants.FIGHT_TYPE_WEEK_COOPERATION {
		// 寒冰堡
		matchResult.SeedId = []int32{7, 2, 4, 2, 3, 8, 4, 8, 8, 9, 3, 2, 8, 1, 8, 7, 5, 2, 7, 8, 1, 5, 1, 3, 5, 7, 8, 1, 6, 1, 2}
		matchResult.ExtraData = map[int64]protocols.T_Fight_Extra_Data{}
		matchResult.BossIdIndexs = []int32{7, 2, 4, 2, 3, 8, 4, 8, 8, 9, 3, 2, 8, 1, 8, 7, 5, 2, 7, 8, 1, 5, 1, 3, 5, 7, 8, 1, 6, 1, 2}
	}
	request.GetConnection().SendMessage(request.GetMsgType(), protocols.P_Match_Result, matchResult.Encode())
}

// 房间匹配
type MatchDuelFightRouter struct {
	net.BaseRouter
}

func (p *MatchDuelFightRouter) Handle(request iface.IRequest) {
	logrus.Info("***************************** 房间匹配 *******************************")
	roleID, err := request.GetConnection().GetProperty("roleID")
	if err != nil {
		logrus.Error("GetProperty error:", err)
		return
	}
	player := core.WorldMgrObj.GetPlayerByPID(roleID.(int64))

	var cMatchDuelFight protocols.C_Match_Duel_Fight
	cMatchDuelFight.Decode(bytes.NewBuffer(request.GetData()), player.Key)
	logrus.Info("房间匹配", cMatchDuelFight)

	if cMatchDuelFight.RoomID != "" {
		room := db.DbManager.FindRoomByRoomIDFightType(cMatchDuelFight.RoomID, cMatchDuelFight.FightType)
		if room == (models.Room{}) {
			logrus.Error("FindRoomByRoomIDFightType error:", err)
			sMatchDuelFight := protocols.S_Match_Duel_Fight{
				Errorcode: 1,
				FightType: cMatchDuelFight.FightType,
			}
			request.GetConnection().SendMessage(request.GetMsgType(), protocols.P_Match_Duel_Fight, sMatchDuelFight.Encode())
			return
		}
		otherPlayer := core.WorldMgrObj.GetPlayerByPID(room.CreatorRoleID)
		if otherPlayer == nil {
			logrus.Error("房间创建者不在线！")
			sMatchDuelFight := protocols.S_Match_Duel_Fight{
				Errorcode: 1,
				FightType: cMatchDuelFight.FightType,
			}
			request.GetConnection().SendMessage(request.GetMsgType(), protocols.P_Match_Duel_Fight, sMatchDuelFight.Encode())
			return
		}

		sMatchDuelFight := protocols.S_Match_Duel_Fight{
			Errorcode: 0,
			FightType: cMatchDuelFight.FightType,
			RoomID:    cMatchDuelFight.RoomID,
		}
		request.GetConnection().SendMessage(request.GetMsgType(), protocols.P_Match_Duel_Fight, sMatchDuelFight.Encode())
		otherPlayer.Conn.SendMessage(request.GetMsgType(), protocols.P_Match_Duel_Fight, sMatchDuelFight.Encode())

		fightToken := utils.GetFightToken()
		var fightItem models.FightItem
		fightItem.FightToken = fightToken
		fightItem.Roles = []int64{player.PID, otherPlayer.PID}
		fightItem.FightStatus = 0
		_, err = db.DbManager.CreateFightItem(fightItem)
		if err != nil {
			logrus.Error("CreateFightItem error:", err)
			return
		}

		// 战斗匹配结果
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

		/************************  加入房间者信息  ************************/
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
		carSkinAttrValueItem := db.DbManager.FindRoleAttrValueItemByAttrID(player.PID, 48)
		carSkinBagItem := db.DbManager.FindCarSkinByItemID(player.PID, carSkinAttrValueItem.Value)
		// 战车皮肤
		var tRuneAbstract = protocols.T_RuneAbstract{ItemID: carSkinBagItem.ItemID*1000 + carSkinBagItem.ItemNum}
		var attrValue = db.DbManager.FindRoleAttrValueItemByAttrID(player.PID, 40)

		/**************************  创建者信息  **************************/
		var creatorRoleHeroSkin = []models.RoleHeroSkin{}
		db.DbManager.FindRoleHeroSkinByRoleID(room.CreatorRoleID, &creatorRoleHeroSkin)
		var creatorRoleHeroSkinMap = make(map[int64]int32, len(creatorRoleHeroSkin))
		for _, v := range creatorRoleHeroSkin {
			creatorRoleHeroSkinMap[v.UUID] = v.ID
		}
		var createorRoleBagItems = []protocols.T_Role_Item{}
		if err := db.DbManager.FindRoleBagItemsByType(room.CreatorRoleID, 2, &createorRoleBagItems); err != nil {
			logrus.Error("FindRoleBagItemsByType error:", err)
			return
		}
		var creatorTHeroAbstract = make(map[int32]protocols.T_HeroAbstract, len(createorRoleBagItems))
		for k, v := range createorRoleBagItems {
			var skinID int32
			if value, ok := creatorRoleHeroSkinMap[v.ItemUUID]; ok {
				skinID = value
			}
			creatorTHeroAbstract[int32(k)] = protocols.T_HeroAbstract{
				HeroUUID:  v.ItemUUID,
				HeroID:    v.ItemID,
				HeroLevel: 24,
				Attr:      []protocols.T_Attr{},
				SkinID:    skinID,
			}
		}
		creatorCarSkinAttrValueItem := db.DbManager.FindRoleAttrValueItemByAttrID(room.CreatorRoleID, 48)
		creatorCarSkinBagItem := db.DbManager.FindCarSkinByItemID(room.CreatorRoleID, creatorCarSkinAttrValueItem.Value)
		// 战车皮肤
		var creatorTRuneAbstract = protocols.T_RuneAbstract{ItemID: creatorCarSkinBagItem.ItemID*1000 + creatorCarSkinBagItem.ItemNum}
		var creatorAttrValue = db.DbManager.FindRoleAttrValueItemByAttrID(room.CreatorRoleID, 40)

		var matchResult = protocols.S_Match_Result{
			Ret:          0,
			FightPattern: cMatchDuelFight.FightType,
			FightType:    cMatchDuelFight.FightType,
			FightModule:  cMatchDuelFight.FightType,
			RoleIdentity: 2,
			MatchRoles: map[int64]protocols.T_RoleAbstract{
				room.CreatorRoleID: {
					RoleID:      room.CreatorRoleID,
					ShowID:      otherPlayer.ShowID,
					BRobot:      false,
					NickName:    otherPlayer.Nickname,
					Heros:       creatorTHeroAbstract,
					Expressions: map[int32]protocols.T_ExpressionAbstract{},
					Runes: map[int32]protocols.T_RuneAbstract{
						40: {ItemID: creatorAttrValue.Value},
						48: creatorTRuneAbstract},
					PetId: creatorAttrValue.Value,
				},
				player.PID: {
					RoleID:      player.PID,
					ShowID:      player.ShowID,
					BRobot:      false,
					NickName:    player.Nickname,
					Heros:       tHeroAbstract,
					Expressions: map[int32]protocols.T_ExpressionAbstract{},
					Runes: map[int32]protocols.T_RuneAbstract{
						40: {ItemID: attrValue.Value},
						48: tRuneAbstract},
					PetId: attrValue.Value,
				},
			},
			FightServerIp:      "127.0.0.1:8081/tfjlh5/fight/ws",
			FightServerPort:    8081,
			FightServerIpSSL:   "127.0.0.1:8081/tfjlh5/fight/ws",
			FightServerPortSSL: 443,
			FightToken:         fightToken,
		}
		if cMatchDuelFight.FightType == constants.FIGHT_TYPE_COOPERATION {
			// 合作
			matchResult.SeedId = []int32{3, 2, 1, 3, 7, 4, 6, 3, 2, 3, 3, 9, 4, 2, 6, 8, 6, 9, 3, 1, 9, 2, 8, 5, 3, 5, 6, 2, 7, 6, 4}
			matchResult.ExtraData = map[int64]protocols.T_Fight_Extra_Data{}
			matchResult.BossIdIndexs = []int32{3, 2, 1, 3, 7, 4, 6, 3, 2, 3, 3, 9, 4, 2, 6, 8, 6, 9, 3, 1, 9, 2, 8, 5, 3, 5, 6, 2, 7, 6, 4}
		} else if cMatchDuelFight.FightType == constants.FIGHT_TYPE_WEEK_COOPERATION {
			// 寒冰堡
			matchResult.SeedId = []int32{7, 2, 4, 2, 3, 8, 4, 8, 8, 9, 3, 2, 8, 1, 8, 7, 5, 2, 7, 8, 1, 5, 1, 3, 5, 7, 8, 1, 6, 1, 2}
			matchResult.ExtraData = map[int64]protocols.T_Fight_Extra_Data{}
			matchResult.BossIdIndexs = []int32{7, 2, 4, 2, 3, 8, 4, 8, 8, 9, 3, 2, 8, 1, 8, 7, 5, 2, 7, 8, 1, 5, 1, 3, 5, 7, 8, 1, 6, 1, 2}
		} else if cMatchDuelFight.FightType == constants.FIGHT_TYPE_FOG_HIDDEN {
			// 雾隐/镖人
			matchResult.SeedId = []int32{6, 6, 9, 7, 1, 9, 3, 6, 4, 5, 3, 3, 5, 5, 3, 3, 9, 3, 7, 1, 8, 9, 6, 6, 5, 4, 1, 6, 1, 6, 3}
			matchResult.ExtraData = map[int64]protocols.T_Fight_Extra_Data{}
			matchResult.BossIdIndexs = []int32{6, 6, 9, 7, 1, 9, 3, 6, 4, 5, 3, 3, 5, 5, 3, 3, 9, 3, 7, 1, 8, 9, 6, 6, 5, 4, 1, 6, 1, 6, 3}
		} else if cMatchDuelFight.FightType == constants.FIGHT_TYPE_MACHINARIUM {
			// 机械迷城
			matchResult.SeedId = []int32{8, 8, 8, 6, 9, 5, 5, 6, 3, 1, 4, 6, 3, 1, 9, 2, 5, 2, 9, 5, 8, 2, 7, 8, 1, 7, 8, 2, 6, 4, 2}
			matchResult.ExtraData = map[int64]protocols.T_Fight_Extra_Data{}
			matchResult.BossIdIndexs = []int32{8, 8, 8, 6, 9, 5, 5, 6, 3, 1, 4, 6, 3, 1, 9, 2, 5, 2, 9, 5, 8, 2, 7, 8, 1, 7, 8, 2, 6, 4, 2}
		}
		request.GetConnection().SendMessage(request.GetMsgType(), protocols.P_Match_Result, matchResult.Encode())
		matchResult.RoleIdentity = 1
		otherPlayer.Conn.SendMessage(request.GetMsgType(), protocols.P_Match_Result, matchResult.Encode())
	} else {
		// 数据库查看是否有可用房间号，可用就取消所有可用房间号的匹配，最后创建一个房间号
		var room []models.Room
		err = db.DbManager.FindRoomsByStatus(player.PID, 0, &room)
		if err != nil {
			logrus.Error("FindRoomsByStatus error:", err)
			return
		}
		if len(room) > 0 {
			// 取消所有可用房间号的匹配
			for _, v := range room {
				_, err = db.DbManager.UpdateRoomStatus(v.ID_, 3)
				if err != nil {
					logrus.Error("UpdateRoomStatus error:", err)
					return
				}
			}
		}
		// 创建一个房间号
		roomID := utils.GetRandomNumber(4)
		longRoomID := fmt.Sprintf("%d%s", time.Now().Unix(), roomID)
		err = db.DbManager.CreateRoom(roomID, longRoomID, player.PID, cMatchDuelFight.FightType, 0)
		if err != nil {
			logrus.Error("CreateRoom error:", err)
			return
		}
		sMatchDuelFight := protocols.S_Match_Duel_Fight{
			Errorcode:  0,
			FightType:  cMatchDuelFight.FightType,
			RoomID:     roomID,
			LongRoomID: longRoomID,
		}
		var roleAbstract = protocols.T_RoleAbstract{
			RoleID:   player.PID,
			ShowID:   player.ShowID,
			BRobot:   false,
			Aiid:     0,
			NickName: player.Nickname,
		}
		sMatchDuelFight.RoleAbstract = roleAbstract
		request.GetConnection().SendMessage(request.GetMsgType(), protocols.P_Match_Duel_Fight, sMatchDuelFight.Encode())
	}
}
