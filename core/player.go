package core

import (
	"tfjl-h5/iface"
	"tfjl-h5/models"
)

type Player struct {
	PID                      int64
	ShowID                   string
	Nickname                 string
	Conn                     iface.IConnection
	Key                      uint8
	Level                    int32
	BIndulge                 bool
	IndulgeTime              int32
	IndulgeDayOnlineTime     int32
	ForbidLoginTimeRemaining int32
}

func NewPlayer(role models.Role, conn iface.IConnection) (*Player, error) {
	p := &Player{
		PID:                      role.RoleID,
		ShowID:                   role.StrID,
		Nickname:                 role.RoleName,
		Conn:                     conn,
		Key:                      role.Key,
		Level:                    role.Level,
		BIndulge:                 role.BIndulge,
		IndulgeTime:              role.IndulgeTime,
		IndulgeDayOnlineTime:     role.IndulgeDayOnlineTime,
		ForbidLoginTimeRemaining: role.ForbidLoginTimeRemaining,
	}
	return p, nil
}

// func (p *Player) SendMsg(msgID uint32, data proto.Message) {
// 	msg, err := proto.Marshal(data)
// 	if err != nil {
// 		log.Println("marshal msg error:", err)
// 		return
// 	}

// 	if p.Conn == nil {
// 		log.Println("Connection in player is nil!")
// 		p.Offline()
// 		return
// 	}

// 	if err := p.Conn.SendMsg(msgID, msg); err != nil {
// 		log.Println("Player SendMsg error!")
// 		return
// 	}
// }

// // 同步开始玩家信息
// func (p *Player) SyncPlayer() {
// 	protoMsg := &pb.SyncPlayer{
// 		PID:      p.PID,
// 		Nickname: p.Nickname,
// 		Pos: &pb.Position{
// 			X: p.X,
// 			Y: p.Y,
// 		},
// 	}
// 	p.SendMsg(2, protoMsg)
// }

// // 广播聊天消息
// func (p *Player) Talk(content string) {
// 	// 封装proto消息
// 	protoMsg := &pb.BroadCast{
// 		PID:  p.PID,
// 		Type: 1,
// 		Data: &pb.BroadCast_Chat{
// 			Chat: &pb.WorldChat{
// 				Nickname: p.Nickname,
// 				Content:  content,
// 			},
// 		},
// 	}
// 	// 得到当前世界所有在线玩家
// 	players := WorldMgrObj.GetAllPlayers()
// 	// 向所有玩家发送广播消息
// 	for _, player := range players {
// 		player.SendMsg(200, protoMsg)
// 	}
// }

// 玩家下线
func (p *Player) Offline() {
	WorldMgrObj.RemovePlayerByPID(p.PID)
}
