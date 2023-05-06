package protocols

// 枚举协议号
const (
	/**********************  登录  **********************/
	// 登录验证
	P_Login_Validate = 1000001
	// 登录请求角色列表
	P_Login_RequestRole = 1000002
	// 登录请求创建角色
	P_Login_CreateRole = 1000003
	// 登录请求选择角色
	P_Login_ChooseRole = 1000004
	// 登录验证在线
	P_Login_ValidateOnline = 1000005
	// 登录验证Ping
	P_Login_Ping = 1000006

	/**********************  角色  **********************/
	// 角色进入游戏
	P_Role_RoleEnterLogic = 2000001
	// 同步角色数据
	P_Role_SynRoleData = 2000002
	// 同步角色属性
	P_Role_SynRoleAttrValue = 2000003
	// 同步角色信息
	P_Role_SynRoleInformationData = 2000004
	// 同步角色任务数据
	P_Role_SynTaskData = 2000009
	// 同步战斗数据
	P_Role_SynBattleArrayData = 2000011
	// 设置默认战斗阵容
	P_Role_BattleArraySetDefine = 2000025
	// 设置战斗阵容
	P_Role_BattleArrayUp = 2000026
	// 角色对战结算数据
	P_Role_FightBalance = 2000030
	// 角色简要信息
	P_Role_GetRoleSimpleInfo = 2000039
	// 同步外挂数据
	P_Role_SynIndulge = 2000052

	// 角色总看广告数据
	P_Role_TotalWatchADBoxData = 2000054
	// 同步角色任务额外数据
	P_Role_SynTaskExtraData = 2000071

	// 同步角色条件数据
	P_Role_SynCondShareData = 2000105

	// 同步角色开关数据
	P_Role_OnOffDataInfo = 2000121

	// 同步角色消耗数据
	P_Role_SyncCostGet = 2000131
	// 修改角色战车皮肤
	P_Role_Car_Skin_Change = 2000138
	// 修改角色英雄皮肤
	P_Role_HeroChangeSkin = 2000151

	/**********************  活动  **********************/
	// 同步角色活动数据
	P_Activity_SynAllActivityData = 4000001

	// 同步角色试炼场数据
	P_Activity_SyncEatChickenData = 4000047

	// 同步角色公会战数据
	P_Activity_SyncSociatyWarData = 4000050

	// 获取角色航海数据
	P_Activity_GetGreatSailingData = 4000052

	// 同步角色航海数据
	P_Activity_SyncGreatSailingData = 4000053

	// 大航海刷新卡组
	P_Activity_GreatSailingRefleshCard = 4000055

	// 同步角色天空之城数据
	P_Activity_SyncSkyCastleData = 4000057

	// 同步角色寒冰堡数据
	P_Activity_SyncWeekCooperationData = 4000060

	// 同步角色周年庆数据
	P_Activity_SyncAnniversaryTrialData = 4000063

	// 同步角色回归数据
	P_Activity_SyncReturnBackData = 4000065

	// 同步角色迷雾数据
	P_Activity_SyncFogHiddenData = 4000067

	// 机械迷城数据
	P_Activity_SyncMachinariumData = 4000069

	/**********************  聊天  **********************/
	// 聊天请求
	P_Chat_ToClient = 6000001
	// 战斗匹配房间关闭
	P_Chat_CloseFightRoom = 6000003

	// 同步朋友数据
	P_Friend_SynLoginData = 7000001

	/**********************  战斗  **********************/
	// 战斗匹配
	P_Match_Fight = 9000001
	// 战斗匹配取消
	P_Match_Cancel = 9000002
	// 战斗匹配结果
	P_Match_Result = 9000003
	// 战斗匹配竞争战斗
	P_Match_Duel_Fight = 9000004
	// 战斗匹配对战取消
	P_Match_Duel_Cancel = 9000005

	// 对战登录
	P_Fight_Role_Login = 10000002
	// 对战加载准备
	P_Fight_Loading_Ready = 10000003
	// 对战开始
	P_Fight_FightStart = 10000004
	// 对战结束
	P_Fight_FightEnd = 10000005
	// 怪物血量同步
	P_Fight_Monster_Blood_SYNC = 10000007
	// 战斗对战结束提交（多人，结束）
	P_Fight_Report_Result_To_Fight = 10000014
	// 战斗匹配对战结束(单人，结束)
	P_Fight_Report_Result_To_Logic = 10000015
	// 提交对战每阶段逻辑数据（多人，过程）
	P_Fight_Report_Phase_Result_To_Fight = 10000020
	// 提交对战每阶段逻辑数据（单人，过程）
	P_Fight_Report_Phase_Result_To_Logic = 10000021
	// 更新英雄
	P_Fight_Update_Hero_SYNC = 10000101
	// 战斗银币同步
	P_Fight_Silver_SYNC = 10000104
	// 战斗卖出英雄同步
	P_Fight_Sell_Hero_SYNC = 10000105
	// 战斗英雄属性同步
	P_Fight_Hero_Attr_SYNC = 10000107
	// 卡牌刷新次数同步
	P_Fight_Refresh_Card_Count_SYNC = 10000108
	// 操作装备同步
	P_Fight_Operate_Equip_SYNC = 10000109

	/**********************  联盟  **********************/

	// 同步联盟数据
	P_Sociaty_SynData = 12000001
	// 获取机械迷城数据
	P_Sociaty_RoleGetMachinariumData = 12000033
	// 同步机械迷城数据
	P_Sociaty_SyncMachinariumData = 12000034
	// 机械迷城选择卡组
	P_Sociaty_RoleMachinariumSelectCard = 12000037

	/**********************  其他  **********************/
	// 对战网络到主逻辑服务
	P_Network_Fight_To_Logic = 99999999
)
