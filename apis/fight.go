package apis

import (
	"bytes"
	"tfjl-h5/constants"
	"tfjl-h5/core"
	"tfjl-h5/iface"
	"tfjl-h5/net"
	"tfjl-h5/protocols"

	"github.com/sirupsen/logrus"
)

// 对战-提交结束结果到主逻辑服务器（单人）
type FightReportResultToLogicRouter struct {
	net.BaseRouter
}

func (p *FightReportResultToLogicRouter) Handle(request iface.IRequest) {
	logrus.Info("*************************** 对战-提交结束结果到主逻辑服务器（单人） ***************************")
	roleID, err := request.GetConnection().GetProperty("roleID")
	if err != nil {
		logrus.Error("GetProperty error:", err)
		return
	}
	player := core.WorldMgrObj.GetPlayerByPID(roleID.(int64))

	var cFightReportResultToLogic = protocols.C_Fight_Report_Result_To_Logic{}
	cFightReportResultToLogic.Decode(bytes.NewBuffer(request.GetData()), player.Key)
	logrus.Infof("cFightReportResultToLogic: %#v", cFightReportResultToLogic)

	var sRoleFightBalance = protocols.S_Role_FightBalance{
		Type: cFightReportResultToLogic.ReportData.FightType,
		BWin: true,
		Roles: map[int64]protocols.T_FightBalance_Role{
			1: {
				RoleAbstract: cFightReportResultToLogic.ReportData.FightRoleInfo[1].RoleAbstract,
			},
			player.PID: {
				RoleAbstract: cFightReportResultToLogic.ReportData.FightRoleInfo[player.PID].RoleAbstract,
			},
		},
		Round:  cFightReportResultToLogic.ReportData.FightRoleInfo[player.PID].Round,
		Battle: map[int32]protocols.T_FightBalance_Battle{},
		Coopration: map[int32]protocols.T_FightBalance_CoopRation{
			0: {
				Prize: []protocols.T_Reward{
					{DropType: 1, DropID: 1001, DropNum: 9999},
					{DropType: 1, DropID: 1002, DropNum: 999},
					{DropType: 2, DropID: 28, DropNum: 999},
					{DropType: 2, DropID: 9, DropNum: 999},
					{DropType: 2, DropID: 24, DropNum: 999},
					{DropType: 1, DropID: 1006, DropNum: 999},
					{DropType: 1, DropID: 1008, DropNum: 999},
					{DropType: 2, DropID: 18, DropNum: 999},
				},
				Extraprize: []protocols.T_Reward{},
			},
		},
		RandomArena:        map[int64]protocols.T_FightBalance_RandomArena{},
		GoldenLeague:       map[int32]protocols.T_FightBlance_GoldenLeague{},
		ActivityCoopration: map[int32]protocols.T_FightBalance_Activity_CoopRation{},
		ExtraData:          map[int32]protocols.T_FightBalance_ExtraData{},
	}
	if cFightReportResultToLogic.ReportData.FightType == constants.FIGHT_TYPE_BATTLE {
		// 对战结算数据
	} else if cFightReportResultToLogic.ReportData.FightType == constants.FIGHT_TYPE_COOPERATION {
		// 合作
	} else if cFightReportResultToLogic.ReportData.FightType == constants.FIGHT_TYPE_BATTLE_GREAT_SAILING {
		// 大航海
		var sFightReportPhaseResultToLogic = protocols.S_Fight_Report_Result_To_Logic{Errorcode: 0}
		request.GetConnection().SendMessage(request.GetMsgType(), protocols.P_Fight_Report_Phase_Result_To_Logic, sFightReportPhaseResultToLogic.Encode())
	} else if cFightReportResultToLogic.ReportData.FightType == constants.FIGHT_TYPE_WEEK_COOPERATION {
		// 寒冰堡
	}
	request.GetConnection().SendMessage(request.GetMsgType(), protocols.P_Role_FightBalance, sRoleFightBalance.Encode())
}

// 对战-每阶段结果提交到主逻辑服务器（单人）
type FightReportPhaseResultToLogicRouter struct {
	net.BaseRouter
}

func (p *FightReportPhaseResultToLogicRouter) Handle(request iface.IRequest) {
	logrus.Info("*************************** 对战-每阶段结果提交到主逻辑服务器（单人） ***************************")
	roleID, err := request.GetConnection().GetProperty("roleID")
	if err != nil {
		logrus.Error("GetProperty error:", err)
		return
	}
	player := core.WorldMgrObj.GetPlayerByPID(roleID.(int64))

	var cFightReportResultToLogic = protocols.C_Fight_Report_Result_To_Logic{}
	cFightReportResultToLogic.Decode(bytes.NewBuffer(request.GetData()), player.Key)
	logrus.Infof("cFightReportResultToLogic: %#v", cFightReportResultToLogic)

	if cFightReportResultToLogic.ReportData.FightType == constants.FIGHT_TYPE_BATTLE {
		// 对战
		var sFightReportPhaseResultToLogic = protocols.S_Fight_Report_Result_To_Logic{Errorcode: 0}
		request.GetConnection().SendMessage(request.GetMsgType(), protocols.P_Fight_Report_Phase_Result_To_Logic, sFightReportPhaseResultToLogic.Encode())
	} else if cFightReportResultToLogic.ReportData.FightType == constants.FIGHT_TYPE_COOPERATION {
		// 合作
		var sFightReportPhaseResultToLogic = protocols.S_Fight_Report_Result_To_Logic{Errorcode: 0}
		request.GetConnection().SendMessage(request.GetMsgType(), protocols.P_Fight_Report_Phase_Result_To_Logic, sFightReportPhaseResultToLogic.Encode())
	} else if cFightReportResultToLogic.ReportData.FightType == constants.FIGHT_TYPE_BATTLE_GREAT_SAILING {
		// 大航海
		var sFightReportPhaseResultToLogic = protocols.S_Fight_Report_Result_To_Logic{Errorcode: 0}
		request.GetConnection().SendMessage(request.GetMsgType(), protocols.P_Fight_Report_Phase_Result_To_Logic, sFightReportPhaseResultToLogic.Encode())
	}
}
