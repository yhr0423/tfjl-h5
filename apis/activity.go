package apis

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"tfjl-h5/iface"
	"tfjl-h5/protocols"
	"tfjl-h5/net"
	"time"

	"github.com/sirupsen/logrus"
)

type ActivityGetGreatSailingDataRouter struct {
	net.BaseRouter
}

func (p *ActivityGetGreatSailingDataRouter) Handle(request iface.IRequest) {

	var cActivityGetGreatSailingData = protocols.C_Activity_GetGreatSailingData{}
	cActivityGetGreatSailingData.Decode(bytes.NewBuffer(request.GetData()), KEY)
	logrus.Infof("%#v", cActivityGetGreatSailingData)

	// 同步大航海数据
	rand.Seed(time.Now().Unix())
	cardID := rand.Intn(135) + 1
	jsonData := []byte(fmt.Sprintf(`{"Error":0,"ActivityID":5000,"FailCount":0,"IsOpen":true,"HistoryMaxScore":200,"TodayScore":200,"DayFailNum":0,"DayMatchNum":0,"ContinuousWinNum":0,"ContinuousFailNum":0,"WinNum":0,"ReliveNum":0,"MaxContinuousWinNum":0,"CardId":%d,"PrizeReward":{},"RefleshCardNum":0}`, cardID))
	var sActivitySyncGreatSailingData protocols.S_Activity_SyncGreatSailingData
	err := json.Unmarshal(jsonData, &sActivitySyncGreatSailingData)
	if err != nil {
		logrus.Error(err)
		return
	}
	logrus.Infof("%#v", sActivitySyncGreatSailingData)
	request.GetConnection().SendMessage(request.GetMsgType(), protocols.P_Activity_SyncGreatSailingData, sActivitySyncGreatSailingData.Encode())
}

type ActivityGreatSailingRefleshCardRouter struct {
	net.BaseRouter
}

func (p *ActivityGreatSailingRefleshCardRouter) Handle(request iface.IRequest) {

	// 刷新大海航卡组
	var cActivityGreatSailingRefleshCard = protocols.C_Activity_GreatSailingRefleshCard{}
	cActivityGreatSailingRefleshCard.Decode(bytes.NewBuffer(request.GetData()), KEY)
	logrus.Infof("%#v", cActivityGreatSailingRefleshCard)

	rand.Seed(time.Now().Unix())
	cardID := rand.Intn(135) + 1
	jsonData := []byte(fmt.Sprintf(`{"Error":0,"ActivityID":5000,"FailCount":0,"IsOpen":true,"HistoryMaxScore":200,"TodayScore":200,"DayFailNum":0,"DayMatchNum":0,"ContinuousWinNum":0,"ContinuousFailNum":0,"WinNum":0,"ReliveNum":0,"MaxContinuousWinNum":0,"CardId":%d,"PrizeReward":{},"RefleshCardNum":0}`, cardID))
	var sActivitySyncGreatSailingData protocols.S_Activity_SyncGreatSailingData
	err := json.Unmarshal(jsonData, &sActivitySyncGreatSailingData)
	if err != nil {
		logrus.Error(err)
		return
	}
	logrus.Infof("%#v", sActivitySyncGreatSailingData)
	request.GetConnection().SendMessage(request.GetMsgType(), protocols.P_Activity_SyncGreatSailingData, sActivitySyncGreatSailingData.Encode())
}
