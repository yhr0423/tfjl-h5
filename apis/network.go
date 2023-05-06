package apis

import (
	"bytes"
	"tfjl-h5/core"
	"tfjl-h5/iface"
	"tfjl-h5/net"
	"tfjl-h5/protocols"

	"github.com/sirupsen/logrus"
)

// 网络-对战websocket服务到客户端服务
type NetworkFightToClientRouter struct {
	net.BaseRouter
}

func (p *NetworkFightToClientRouter) Handle(request iface.IRequest) {
	logrus.Info("************************对战服务到客户端服务************************")

	var cNetworkFightToClientRoleFightBalance protocols.C_Network_Fight_To_Logic_Role_FightBalance
	cNetworkFightToClientRoleFightBalance.Decode(bytes.NewBuffer(request.GetData()))

	playerClient := core.WorldMgrObj.GetPlayerByPID(cNetworkFightToClientRoleFightBalance.RoleID)
	playerClient.Conn.SendMessage(request.GetMsgType(), protocols.P_Role_FightBalance, cNetworkFightToClientRoleFightBalance.SRoleFightBalance.Encode())
}
