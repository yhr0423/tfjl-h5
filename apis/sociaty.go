package apis

import (
	"bytes"
	"encoding/json"
	"tfjl-h5/core"
	"tfjl-h5/iface"
	"tfjl-h5/protocols"
	"tfjl-h5/net"
	"tfjl-h5/utils"

	"github.com/sirupsen/logrus"
)

// 联盟-机械迷城
type SociatyRoleGetMachinariumDataRouter struct {
	net.BaseRouter
}

func (p *SociatyRoleGetMachinariumDataRouter) Handle(request iface.IRequest) {
	roleID, err := request.GetConnection().GetProperty("roleID")
	if err != nil {
		logrus.Error("GetProperty error:", err)
		return
	}
	player := core.WorldMgrObj.GetPlayerByPID(roleID.(int64))

	var cSociatyRoleGetMachinariumData = protocols.C_Sociaty_RoleGetMachinariumData{}
	cSociatyRoleGetMachinariumData.Decode(bytes.NewBuffer(request.GetData()), player.Key)
	logrus.Infof("%#v", cSociatyRoleGetMachinariumData)

	// 机械迷城数据
	jsonData := []byte(`{"SociatyId":12345,"CardId":13,"DayMaxRound":0,"DayFailNum":0,"RankListData":[],"AllRound":999,"RewardStatus":{"10000":0,"100000":0,"120000":0,"160000":0,"20000":0,"200000":0,"40000":0,"5000":0,"60000":0,"80000":0}}`)
	var sSociatySyncMachinariumData protocols.S_Sociaty_SyncMachinariumData
	err = json.Unmarshal(jsonData, &sSociatySyncMachinariumData)
	if err != nil {
		logrus.Error(err)
		return
	}
	sSociatySyncMachinariumData.RoleID = player.PID
	sSociatySyncMachinariumData.CardID = utils.GetRandomMachinariumcarID()
	request.GetConnection().SendMessage(request.GetMsgType(), protocols.P_Sociaty_SyncMachinariumData, sSociatySyncMachinariumData.Encode())
}

// 联盟-机械迷城-获取卡组
type SociatyRoleMachinariumSelectCardRouter struct {
	net.BaseRouter
}

func (p *SociatyRoleMachinariumSelectCardRouter) Handle(request iface.IRequest) {
	roleID, err := request.GetConnection().GetProperty("roleID")
	if err != nil {
		logrus.Error("GetProperty error:", err)
		return
	}
	player := core.WorldMgrObj.GetPlayerByPID(roleID.(int64))

	var cSociatyRoleMachinariumSelectCard = protocols.C_Sociaty_RoleMachinariumSelectCard{}
	cSociatyRoleMachinariumSelectCard.Decode(bytes.NewBuffer(request.GetData()), player.Key)
	logrus.Infof("%#v", cSociatyRoleMachinariumSelectCard)

	
	randomMachinariumCarID := utils.GetRandomMachinariumcarID()
	// 机械迷城数据
	jsonData := []byte(`{"SociatyId":12345,"CardId":13,"DayMaxRound":0,"DayFailNum":0,"RankListData":[],"AllRound":999,"RewardStatus":{"10000":0,"100000":0,"120000":0,"160000":0,"20000":0,"200000":0,"40000":0,"5000":0,"60000":0,"80000":0}}`)
	var sSociatySyncMachinariumData protocols.S_Sociaty_SyncMachinariumData
	err = json.Unmarshal(jsonData, &sSociatySyncMachinariumData)
	if err != nil {
		logrus.Error(err)
		return
	}
	sSociatySyncMachinariumData.RoleID = player.PID
	sSociatySyncMachinariumData.CardID = randomMachinariumCarID
	request.GetConnection().SendMessage(request.GetMsgType(), protocols.P_Sociaty_SyncMachinariumData, sSociatySyncMachinariumData.Encode())

	// 机械迷城选择卡组数据
	var sSociatyRoleMachinariumSelectCard = protocols.S_Sociaty_RoleMachinariumSelectCard{
		Errorcode: 0,
		CardID:   randomMachinariumCarID,
	}
	request.GetConnection().SendMessage(request.GetMsgType(), protocols.P_Sociaty_RoleMachinariumSelectCard, sSociatyRoleMachinariumSelectCard.Encode())
}
