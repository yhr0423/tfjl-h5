package core

import "sync"

/*
当前游戏世界总管理模块
*/
type WorldManager struct {
	// 当前全部在线Players集合
	Players map[int64]*Player
	// 保护Players集合的锁
	pLock sync.RWMutex
}

// 提供一个对外的世界管理模块引用（全局）
var WorldMgrObj *WorldManager

func init() {
	WorldMgrObj = &WorldManager{
		Players: make(map[int64]*Player),
	}
}

// 添加一个玩家
func (wm *WorldManager) AddPlayer(player *Player) {
	wm.pLock.Lock()
	wm.Players[player.PID] = player
	wm.pLock.Unlock()
}

// 删除一个玩家
func (wm *WorldManager) RemovePlayerByPID(pid int64) {
	wm.pLock.Lock()
	delete(wm.Players, pid)
	wm.pLock.Unlock()
}

// 通过玩家ID查询Player对象
func (wm *WorldManager) GetPlayerByPID(pid int64) *Player {
	wm.pLock.RLock()
	defer wm.pLock.RUnlock()

	return wm.Players[pid]
}

// 获取全部在线玩家
func (wm *WorldManager) GetAllPlayers() []*Player {
	wm.pLock.RLock()
	defer wm.pLock.RUnlock()

	players := make([]*Player, 0)

	for _, player := range wm.Players {
		players = append(players, player)
	}

	return players
}
