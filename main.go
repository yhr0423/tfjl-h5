package main

import (
	"bytes"
	"context"
	"crypto/md5"
	"encoding/base64"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"tfjl-h5/apis"
	"tfjl-h5/configs"
	"tfjl-h5/core"
	"tfjl-h5/db"
	"tfjl-h5/iface"
	"tfjl-h5/models"
	"tfjl-h5/net"
	"tfjl-h5/protocols"
	"tfjl-h5/utils"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/securecookie"
	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	configFile string
)

func init() {
	logrus.SetOutput(os.Stdout)
	db.DbManager.SetUserCollection("user")
	db.DbManager.SetRoleCollection("role")
	db.DbManager.SetRoleInformationCollection("role_information")
	db.DbManager.SetRoleBagItemsCollection("role_bag_items")
	db.DbManager.SetRoleAttrValueItemsCollection("role_attr_value_items")
	db.DbManager.SetRoleBattleArrayCollection("role_battle_array")
	db.DbManager.SetRoleHeroSkinCollection("role_hero_skin")
	db.DbManager.SetRoleTaskItemsCollection("role_task_items")
	db.DbManager.SetRoleSeasonItemsCollection("role_season_items")
	db.DbManager.SetRoleSeasonForeverScorePrizeCollection("role_season_forever_score_prize")
	db.DbManager.SetRoleSeasonScorePrizeCollection("role_season_score_prize")
	db.DbManager.SetActivityCollection("activity")
	db.DbManager.SetRoomCollection("room")
	db.DbManager.SetFightItemsCollection("fight_items")
}

// 用户认证中间件
func Authorize() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Header.Get("Authorization") == "e756795a-1245-458f-ae1c-8f1e2ccf5e28" {
			c.Next()
			return
		}

		// 获取token
		session := sessions.Default(c)
		token := session.Get("Authorization")
		if token == nil || token == "" {
			// 如果用户未登录，则重定向到登录页面
			c.Redirect(http.StatusMovedPermanently, "/tfjlh5/")
			c.Abort()
			return
		}

		user := db.DbManager.FindUserByToken(token.(string))
		if user == (models.User{}) {
			c.Redirect(http.StatusMovedPermanently, "/tfjlh5/")
			c.Abort()
			return
		}

		// 将token附加到上下文中，以便后续使用
		c.Set("token", token)

		// 执行下一个中间件或处理程序
		c.Next()
	}
}

func initCmd() {
	flag.StringVar(&configFile, "config", "./config.json", "where load config json")
	flag.Parse()
}

func OnConnectionLost(conn iface.IConnection) {
	roleID, err := conn.GetProperty("roleID")
	if err != nil {
		logrus.Info("conn GetProperty roleID error:", err)
		return
	}
	player := core.WorldMgrObj.GetPlayerByPID(roleID.(int64))
	// 玩家下线
	player.Offline()
	logrus.Info("========> Player roleID =", roleID, "offline... <========")
}

func main() {
	initCmd()
	if err := configs.LoadConfig(configFile); err != nil {
		fmt.Println("Load config json error:", err)
	}
	net.GWServer = net.NewServer()
	net.GWServer.SetOnConnStop(OnConnectionLost)

	// Ping
	net.GWServer.AddRouter(protocols.P_Login_Ping, &apis.LoginPingRouter{})
	// 登录验证在线
	net.GWServer.AddRouter(protocols.P_Login_ValidateOnline, &apis.LoginValidateOnlineRouter{})
	// 重连验证在线
	net.GWServer.AddRouter(protocols.P_Login_Validate, &apis.LoginValidateOnlineRouter{})
	// 登录请求角色
	net.GWServer.AddRouter(protocols.P_Login_RequestRole, &apis.LoginRequestRoleRouter{})
	// 登录选择角色
	net.GWServer.AddRouter(protocols.P_Login_ChooseRole, &apis.LoginChooseRoleRouter{})

	// 角色同步数据
	net.GWServer.AddRouter(protocols.P_Role_SynRoleData, &apis.RoleSynRoleDataRouter{})
	// 阵容上阵更新
	net.GWServer.AddRouter(protocols.P_Role_BattleArrayUp, &apis.RoleBattleArrayUpRouter{})
	// 角色车皮修改
	net.GWServer.AddRouter(protocols.P_Role_Car_Skin_Change, &apis.RoleCarSkinChangeRouter{})
	// 角色英雄皮肤修改
	net.GWServer.AddRouter(protocols.P_Role_HeroChangeSkin, &apis.RoleHeroChangeSkinRouter{})
	// 角色获取简要信息（头像框点击）
	net.GWServer.AddRouter(protocols.P_Role_GetRoleSimpleInfo, &apis.RoleGetRoleSimpleInfoRouter{})

	// 活动-大航海数据获得
	net.GWServer.AddRouter(protocols.P_Activity_GetGreatSailingData, &apis.ActivityGetGreatSailingDataRouter{})
	// 活动-大航海刷新卡牌
	net.GWServer.AddRouter(protocols.P_Activity_GreatSailingRefleshCard, &apis.ActivityGreatSailingRefleshCardRouter{})

	// 匹配-快速匹配
	net.GWServer.AddRouter(protocols.P_Match_Fight, &apis.MatchFightRouter{})
	// 匹配-房间匹配
	net.GWServer.AddRouter(protocols.P_Match_Duel_Fight, &apis.MatchDuelFightRouter{})
	// 对战-提交结束的战斗结果到逻辑服务器
	net.GWServer.AddRouter(protocols.P_Fight_Report_Result_To_Logic, &apis.FightReportResultToLogicRouter{})
	// 对战-提交每阶段的战斗结果到逻辑服务器
	net.GWServer.AddRouter(protocols.P_Fight_Report_Phase_Result_To_Logic, &apis.FightReportPhaseResultToLogicRouter{})

	// 联盟-获取机械迷城数据
	net.GWServer.AddRouter(protocols.P_Sociaty_RoleGetMachinariumData, &apis.SociatyRoleGetMachinariumDataRouter{})
	// 联盟-机械迷城卡组数据
	net.GWServer.AddRouter(protocols.P_Sociaty_RoleMachinariumSelectCard, &apis.SociatyRoleMachinariumSelectCardRouter{})

	// 对战网络与主逻辑服务器通信
	net.GWServer.AddRouter(protocols.P_Network_Fight_To_Logic, &apis.NetworkFightToClientRouter{})

	router := gin.Default()
	// 初始化基于Cookie的Session中间件
	store := cookie.NewStore(securecookie.GenerateRandomKey(64), securecookie.GenerateRandomKey(32))
	store.Options(sessions.Options{
		Path:   "/",
		MaxAge: 86400 * 7, // 一周
	})
	router.Use(sessions.Sessions("tfjlh5session", store))

	router.Use(static.Serve("/tfjlh5", static.LocalFile("static", false)))
	// 中间件处理
	router.Use(func(c *gin.Context) {
		// 尝试获取静态资源
		c.Next()

		// 如果静态资源不存在，则返回 404 Not Found 错误
		if c.Writer.Status() == http.StatusNotFound {
			// 获取 URL
			url := c.Request.URL.String()

			logrus.Info("url: ", url)

			if strings.Contains(url, "/tfjlh5/assets/") || strings.Contains(url, "/tfjlh5/resources/") || strings.Contains(url, "/tfjlh5/src/") {
				url = strings.Replace(url, "/tfjlh5/", "", 1)
				// 下载文件
				err := downloadFile(url)
				if err != nil {
					logrus.Error("文件下载失败:", err)
					c.AbortWithStatus(http.StatusInternalServerError)
					return
				}

				// 返回刚刚下载的文件
				c.File("static/" + url)
			}
		}
	})
	router.LoadHTMLGlob("templates/*")
	// 解密websocket数据接口
	router.POST("/tfjlh5/decode", Authorize(), decode)
	// 创建用户接口
	router.GET("/tfjlh5/create", Authorize(), create)
	// 删除用户接口
	router.GET("/tfjlh5/delete", Authorize(), delete)

	router.GET("/tfjlh5/", func(c *gin.Context) {
		session := sessions.Default(c)
		token := session.Get("Authorization")
		if token == nil || token == "" {
			c.HTML(http.StatusOK, "login.html", nil)
			return
		}

		user := db.DbManager.FindUserByToken(token.(string))
		if user == (models.User{}) {
			c.HTML(http.StatusOK, "login.html", nil)
			return
		}

		c.Redirect(http.StatusMovedPermanently, "/tfjlh5/index")
	})
	router.POST("/tfjlh5/login", func(c *gin.Context) {
		username := c.PostForm("username")
		user := db.DbManager.FindUserByAccount(username)
		if user == (models.User{}) {
			c.JSON(http.StatusOK, gin.H{"info": "账号或者密码错误"})
			return
		}
		password := c.PostForm("password")
		passwordCiphertext := fmt.Sprintf("%x", md5.Sum([]byte(password)))
		logrus.Info(passwordCiphertext)
		if passwordCiphertext != user.PasswordCiphertext {
			c.JSON(http.StatusOK, gin.H{"info": "账号或者密码错误"})
			return
		}
		token := uuid.NewV4()
		db.DbManager.UpdateTokenByAccount(username, token.String())
		// 在会话中设置值
		session := sessions.Default(c)
		session.Set("Authorization", token.String())
		session.Save()
		c.JSON(http.StatusOK, gin.H{"redirect": "index"})
	})
	router.GET("/tfjlh5/index", Authorize(), func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	router.GET("/tfjlh5/gameservergroup", Authorize(), func(c *gin.Context) {
		// 假设result包含解析后的JSON数据
		result := models.Result{
			Data: models.Data{
				OpenID:  "3514510181",
				Tishen:  0,
				ExtJSON: "null",
				SdkType: 10,
				SdkID:   0,
			},
			Servers: []models.Server{
				{
					GroupName: "\u5854\u9632\u7cbe\u7075",
					GroupID:   1,
					State:     1,
					Roles:     []int{},
				},
			},
		}
		c.JSON(http.StatusOK, result)
	})
	router.GET("/tfjlh5/pcloginbygroup", Authorize(), func(c *gin.Context) {
		token := c.GetString("token")
		user := db.DbManager.FindUserByToken(token)
		if user == (models.User{}) {
			c.JSON(http.StatusOK, gin.H{"error": 1})
			return
		}
		result := models.LoginResult{
			Error:       0,
			SdkType:     10,
			SdkId:       0,
			AccountName: user.Account,
			OpenId:      nil,
			Zone:        0,
			WebName:     "127.0.0.1:8080/tfjlh5/ws",
			WebPort:     "443",
			WanIp:       "",
			WanPort:     "",
			Sign:        "",
			Examine:     0,
			AdToShare:   0,
		}
		c.JSON(http.StatusOK, result)
	})
	router.GET("/tfjlh5/ws", Authorize(), net.WsHandler)
	bindAddress := fmt.Sprintf("%s:%d", configs.GConf.Ip, configs.GConf.Port)
	srv := &http.Server{
		Addr:    bindAddress,
		Handler: router,
	}

	go func() {
		// 服务连接
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logrus.Fatalf("listen: %s\n", err)
		}
	}()

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	signal := <-quit
	log.Println("Shutdown net ..., Signal:", signal)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// 关闭gin的http服务器
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("net Shutdown:", err)
	}
	log.Println("net exiting")
	// 关闭数据库
	db.DbManager.CloseDB()
}

// downloadFile 下载文件（从官方下载文件到本地）
func downloadFile(url string) error {
	req, err := http.NewRequest("GET", "https://mszctest-1300944069.file.myqcloud.com/miyaVideoH5/web-mobile/"+url, nil)
	if err != nil {
		return err
	}
	// 添加自定义标头
	req.Header.Set("User-Agent", "Mozilla/5.0 (Linux; Android 5.1.1; SM-G977N Build/LMY48Z; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/74.0.3729.136 Mobile Safari/537.36")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	filePath := "static/" + url
	dirPath := filepath.Dir(filePath)
	err = os.MkdirAll(dirPath, os.ModePerm)
	if err != nil {
		return err
	}

	out, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	logrus.Info("文件已下载:", url)

	return nil
}

// 解密抓包数据接口（测试用）
func decode(c *gin.Context) {
	token := c.GetString("token")
	curUser := db.DbManager.FindUserByToken(token)
	if curUser == (models.User{}) && curUser.Account != "test" {
		c.JSON(http.StatusOK, gin.H{"error": 1})
		return
	}
	var websocketDataDecode models.WebsocketDataDecode
	if err := c.ShouldBindJSON(&websocketDataDecode); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": err.Error()})
		return
	}
	byteArr, err := base64.StdEncoding.DecodeString(websocketDataDecode.Bytes)
	if err != nil {
		logrus.Error(err)
		c.JSON(http.StatusOK, gin.H{"data": "bytes base64字符串解码失败"})
		return
	}
	logrus.Info(byteArr)
	buffer := bytes.NewBuffer(byteArr)
	var res interface{}
	switch websocketDataDecode.ProtocolNum {

	/***************************** 主逻辑服务器相关协议 ********************************/
	case protocols.P_Login_ValidateOnline:
		if websocketDataDecode.ClientType == 1 {
			message := protocols.C_Login_ValidateOnline{}
			err = message.Decode(buffer)
			if err != nil {
				logrus.Error(err)
				c.JSON(http.StatusOK, gin.H{"data": err.Error()})
				return
			}
			res = message
		} else if websocketDataDecode.ClientType == 2 {
		}
	case protocols.P_Login_Validate:
		if websocketDataDecode.ClientType == 1 {
		} else if websocketDataDecode.ClientType == 2 {
			message := protocols.S_Login_Validate{}
			err = message.Decode(buffer)
			if err != nil {
				logrus.Error(err)
				c.JSON(http.StatusOK, gin.H{"data": err.Error()})
				return
			}
			apis.KEY = message.Key
			res = message
		}
	case protocols.P_Login_RequestRole:
		if websocketDataDecode.ClientType == 1 {
		} else if websocketDataDecode.ClientType == 2 {
			message := protocols.S_Login_RequestRole{}
			err = message.Decode(buffer)
			if err != nil {
				logrus.Error(err)
				c.JSON(http.StatusOK, gin.H{"data": err.Error()})
				return
			}
			res = message
		}
	case protocols.P_Role_SynRoleAttrValue:
		if websocketDataDecode.ClientType == 1 {

		} else if websocketDataDecode.ClientType == 2 {
			message := protocols.S_Role_SynRoleAttrValue{}
			err = message.Decode(buffer)
			if err != nil {
				logrus.Error(err)
				c.JSON(http.StatusOK, gin.H{"data": err.Error()})
				return
			}
			res = message
		}
	case protocols.P_Role_SynRoleData:
		if websocketDataDecode.ClientType == 1 {

		} else if websocketDataDecode.ClientType == 2 {
			message := protocols.S_Role_SynRoleData{}
			err = message.Decode(buffer)
			if err != nil {
				logrus.Error(err)
				c.JSON(http.StatusOK, gin.H{"data": err.Error()})
				return
			}
			res = message
		}
	case protocols.P_Activity_SyncEatChickenData:
		// 试炼场
		if websocketDataDecode.ClientType == 1 {
		} else if websocketDataDecode.ClientType == 2 {
			message := protocols.S_Activity_SyncEatChickenData{}
			err = message.Decode(buffer)
			if err != nil {
				logrus.Error(err)
				c.JSON(http.StatusOK, gin.H{"data": err.Error()})
				return
			}
			res = message
		}
	case protocols.P_Activity_SyncGreatSailingData:
		// 大航海
		if websocketDataDecode.ClientType == 1 {
		} else if websocketDataDecode.ClientType == 2 {
			message := protocols.S_Activity_SyncGreatSailingData{}
			err = message.Decode(buffer)
			if err != nil {
				logrus.Error(err)
				c.JSON(http.StatusOK, gin.H{"data": err.Error()})
				return
			}
			res = message
		}
	case protocols.P_Activity_SyncWeekCooperationData:
		// 寒冰堡（每周合作）
		if websocketDataDecode.ClientType == 1 {
		} else if websocketDataDecode.ClientType == 2 {
			message := protocols.S_Activity_SyncWeekCooperationData{}
			err = message.Decode(buffer)
			if err != nil {
				logrus.Error(err)
				c.JSON(http.StatusOK, gin.H{"data": err.Error()})
				return
			}
			res = message
		}
	case protocols.P_Activity_SyncMachinariumData:
		// 机械迷城数据
		if websocketDataDecode.ClientType == 1 {
		} else if websocketDataDecode.ClientType == 2 {
			message := protocols.S_Activity_SyncMachinariumData{}
			err = message.Decode(buffer)
			if err != nil {
				logrus.Error(err)
				c.JSON(http.StatusOK, gin.H{"data": err.Error()})
				return
			}
			res = message
		}
	case protocols.P_Role_OnOffDataInfo:
		// 开关数据
		if websocketDataDecode.ClientType == 1 {
		} else if websocketDataDecode.ClientType == 2 {
			message := protocols.S_Role_OnOffDataInfo{}
			err = message.Decode(buffer)
			if err != nil {
				logrus.Error(err)
				c.JSON(http.StatusOK, gin.H{"data": err.Error()})
				return
			}
			res = message
		}
	case protocols.P_Match_Result:
		// 匹配结果
		if websocketDataDecode.ClientType == 1 {
		} else if websocketDataDecode.ClientType == 2 {
			message := protocols.S_Match_Result{}
			err = message.Decode(buffer)
			if err != nil {
				logrus.Error(err)
				c.JSON(http.StatusOK, gin.H{"data": err.Error()})
				return
			}
			res = message
		}
	case protocols.P_Match_Duel_Fight:
		// 匹配对战
		if websocketDataDecode.ClientType == 1 {
			message := protocols.C_Match_Duel_Fight{}
			err = message.Decode(buffer, apis.KEY)
			if err != nil {
				logrus.Error(err)
				c.JSON(http.StatusOK, gin.H{"data": err.Error()})
				return
			}
			res = message
		} else if websocketDataDecode.ClientType == 2 {
			message := protocols.S_Match_Duel_Fight{}
			err = message.Decode(buffer)
			if err != nil {
				logrus.Error(err)
				c.JSON(http.StatusOK, gin.H{"data": err.Error()})
				return
			}
			res = message
		}
	case protocols.P_Chat_CloseFightRoom:
		// 匹配对战
		if websocketDataDecode.ClientType == 1 {
		} else if websocketDataDecode.ClientType == 2 {
			message := protocols.S_Chat_CloseFightRoom{}
			err = message.Decode(buffer)
			if err != nil {
				logrus.Error(err)
				c.JSON(http.StatusOK, gin.H{"data": err.Error()})
				return
			}
			res = message
		}
	case protocols.P_Fight_Report_Phase_Result_To_Logic:
		// 战斗阶段结果提交
		if websocketDataDecode.ClientType == 1 {
			message := protocols.C_Fight_Report_Result_To_Logic{}
			err = message.Decode(buffer, apis.KEY)
			if err != nil {
				logrus.Error(err)
				c.JSON(http.StatusOK, gin.H{"data": err.Error()})
				return
			}
			res = message
		} else if websocketDataDecode.ClientType == 2 {
			message := protocols.S_Fight_Report_Result_To_Logic{}
			err = message.Decode(buffer)
			if err != nil {
				logrus.Error(err)
				c.JSON(http.StatusOK, gin.H{"data": err.Error()})
				return
			}
			res = message
		}
	case protocols.P_Fight_Report_Result_To_Logic:
		// 战斗结束提交
		if websocketDataDecode.ClientType == 1 {
			message := protocols.C_Fight_Report_Result_To_Logic{}
			err = message.Decode(buffer, apis.KEY)
			if err != nil {
				logrus.Error(err)
				c.JSON(http.StatusOK, gin.H{"data": err.Error()})
				return
			}
			res = message
		} else if websocketDataDecode.ClientType == 2 {
			message := protocols.S_Fight_Report_Result_To_Logic{}
			err = message.Decode(buffer)
			if err != nil {
				logrus.Error(err)
				c.JSON(http.StatusOK, gin.H{"data": err.Error()})
				return
			}
			res = message
		}
	case protocols.P_Role_FightBalance:
		// 角色对战结算数据
		if websocketDataDecode.ClientType == 1 {
		} else if websocketDataDecode.ClientType == 2 {
			message := protocols.S_Role_FightBalance{}
			err = message.Decode(buffer)
			if err != nil {
				logrus.Error(err)
				c.JSON(http.StatusOK, gin.H{"data": err.Error()})
				return
			}
			res = message
		}
	case protocols.P_Activity_GetGreatSailingData:
		// 获取大航海数据
		if websocketDataDecode.ClientType == 1 {
		} else if websocketDataDecode.ClientType == 2 {
			message := protocols.S_Activity_SyncGreatSailingData{}
			err = message.Decode(buffer)
			if err != nil {
				logrus.Error(err)
				c.JSON(http.StatusOK, gin.H{"data": err.Error()})
				return
			}
			res = message
		}
	case protocols.P_Activity_GreatSailingRefleshCard:
		// 刷新大航海
		if websocketDataDecode.ClientType == 1 {
			message := protocols.C_Activity_GreatSailingRefleshCard{}
			err = message.Decode(buffer, apis.KEY)
			if err != nil {
				logrus.Error(err)
				c.JSON(http.StatusOK, gin.H{"data": err.Error()})
				return
			}
			res = message
		} else if websocketDataDecode.ClientType == 2 {
		}
	case protocols.P_Role_GetRoleSimpleInfo:
		if websocketDataDecode.ClientType == 1 {
		} else if websocketDataDecode.ClientType == 2 {
			message := protocols.S_Role_GetRoleSimpleInfo{}
			err = message.Decode(buffer)
			if err != nil {
				logrus.Error(err)
				c.JSON(http.StatusOK, gin.H{"data": err.Error()})
				return
			}
			res = message
		}
	case protocols.P_Sociaty_SynData:
		if websocketDataDecode.ClientType == 1 {
		} else if websocketDataDecode.ClientType == 2 {
			message := protocols.S_Sociaty_SynData{}
			err = message.Decode(buffer)
			if err != nil {
				logrus.Error(err)
				c.JSON(http.StatusOK, gin.H{"data": err.Error()})
				return
			}
			res = message
		}
	case protocols.P_Sociaty_RoleGetMachinariumData:
		if websocketDataDecode.ClientType == 1 {
			message := protocols.C_Sociaty_RoleGetMachinariumData{}
			err = message.Decode(buffer, apis.KEY)
			if err != nil {
				logrus.Error(err)
				c.JSON(http.StatusOK, gin.H{"data": err.Error()})
				return
			}
			res = message
		} else if websocketDataDecode.ClientType == 2 {
		}
	case protocols.P_Sociaty_SyncMachinariumData:
		if websocketDataDecode.ClientType == 1 {
		} else if websocketDataDecode.ClientType == 2 {
			message := protocols.S_Sociaty_SyncMachinariumData{}
			err = message.Decode(buffer)
			if err != nil {
				logrus.Error(err)
				c.JSON(http.StatusOK, gin.H{"data": err.Error()})
				return
			}
			res = message
		}
	case protocols.P_Sociaty_RoleMachinariumSelectCard:
		if websocketDataDecode.ClientType == 1 {
			message := protocols.C_Sociaty_RoleMachinariumSelectCard{}
			err = message.Decode(buffer, apis.KEY)
			if err != nil {
				logrus.Error(err)
				c.JSON(http.StatusOK, gin.H{"data": err.Error()})
				return
			}
			res = message
		} else if websocketDataDecode.ClientType == 2 {
			message := protocols.S_Sociaty_RoleMachinariumSelectCard{}
			err = message.Decode(buffer)
			if err != nil {
				logrus.Error(err)
				c.JSON(http.StatusOK, gin.H{"data": err.Error()})
				return
			}
			res = message
		}

	/***************************** fight相关协议 ********************************/
	case protocols.P_Fight_Role_Login:
		if websocketDataDecode.ClientType == 1 {
			message := protocols.C_Fight_Role_Login{}
			err = message.Decode(buffer)
			if err != nil {
				logrus.Error(err)
				c.JSON(http.StatusOK, gin.H{"data": err.Error()})
				return
			}
			res = message
		} else if websocketDataDecode.ClientType == 2 {
			message := protocols.S_Fight_Role_Login{}
			err = message.Decode(buffer)
			if err != nil {
				logrus.Error(err)
				c.JSON(http.StatusOK, gin.H{"data": err.Error()})
				return
			}
			apis.FIGHT_KEY = message.Key
			res = message
		}
	case protocols.P_Fight_Silver_SYNC:
		if websocketDataDecode.ClientType == 1 {
			message := protocols.C_Fight_Silver_SYNC{}
			err = message.Decode(buffer, apis.FIGHT_KEY)
			if err != nil {
				logrus.Error(err)
				c.JSON(http.StatusOK, gin.H{"data": err.Error()})
				return
			}
			res = message
		} else if websocketDataDecode.ClientType == 2 {
			message := protocols.S_Fight_Silver_SYNC{}
			err = message.Decode(buffer)
			if err != nil {
				logrus.Error(err)
				c.JSON(http.StatusOK, gin.H{"data": err.Error()})
				return
			}
			res = message
		}
	case protocols.P_Fight_Loading_Ready:
		if websocketDataDecode.ClientType == 1 {
			message := protocols.C_Fight_Loading_Ready{}
			err = message.Decode(buffer, apis.FIGHT_KEY)
			if err != nil {
				logrus.Error(err)
				c.JSON(http.StatusOK, gin.H{"data": err.Error()})
				return
			}
			res = message
		} else if websocketDataDecode.ClientType == 2 {
			message := protocols.S_Fight_Loading_Ready{}
			err = message.Decode(buffer)
			if err != nil {
				logrus.Error(err)
				c.JSON(http.StatusOK, gin.H{"data": err.Error()})
				return
			}
			res = message
		}
	case protocols.P_Fight_FightStart:
		if websocketDataDecode.ClientType == 1 {
		} else if websocketDataDecode.ClientType == 2 {
			message := protocols.S_Fight_FightStart{}
			err = message.Decode(buffer)
			if err != nil {
				logrus.Error(err)
				c.JSON(http.StatusOK, gin.H{"data": err.Error()})
				return
			}
			res = message
		}
	case protocols.P_Fight_FightEnd:
		if websocketDataDecode.ClientType == 1 {
		} else if websocketDataDecode.ClientType == 2 {
			message := protocols.S_Fight_FightEnd{}
			err = message.Decode(buffer)
			if err != nil {
				logrus.Error(err)
				c.JSON(http.StatusOK, gin.H{"data": err.Error()})
				return
			}
			res = message
		}
	case protocols.P_Fight_Refresh_Card_Count_SYNC:
		if websocketDataDecode.ClientType == 1 {
			message := protocols.C_Fight_Refresh_Card_Count_SYNC{}
			err = message.Decode(buffer, apis.FIGHT_KEY)
			if err != nil {
				logrus.Error(err)
				c.JSON(http.StatusOK, gin.H{"data": err.Error()})
				return
			}
			res = message
		} else if websocketDataDecode.ClientType == 2 {
			message := protocols.S_Fight_Refresh_Card_Count_SYNC{}
			err = message.Decode(buffer)
			if err != nil {
				logrus.Error(err)
				c.JSON(http.StatusOK, gin.H{"data": err.Error()})
				return
			}
			res = message
		}
	case protocols.P_Fight_Monster_Blood_SYNC:
		if websocketDataDecode.ClientType == 1 {
			message := protocols.C_Fight_Monster_Blood_SYNC{}
			err = message.Decode(buffer, apis.FIGHT_KEY)
			if err != nil {
				logrus.Error(err)
				c.JSON(http.StatusOK, gin.H{"data": err.Error()})
				return
			}
			res = message
		} else if websocketDataDecode.ClientType == 2 {
			message := protocols.S_Fight_Monster_Blood_SYNC{}
			err = message.Decode(buffer)
			if err != nil {
				logrus.Error(err)
				c.JSON(http.StatusOK, gin.H{"data": err.Error()})
				return
			}
			res = message
		}
	case protocols.P_Fight_Hero_Attr_SYNC:
		if websocketDataDecode.ClientType == 1 {
			message := protocols.C_Fight_Hero_Attr_SYNC{}
			err = message.Decode(buffer, apis.FIGHT_KEY)
			if err != nil {
				logrus.Error(err)
				c.JSON(http.StatusOK, gin.H{"data": err.Error()})
				return
			}
			res = message
		} else if websocketDataDecode.ClientType == 2 {
			message := protocols.S_Fight_Hero_Attr_SYNC{}
			err = message.Decode(buffer)
			if err != nil {
				logrus.Error(err)
				c.JSON(http.StatusOK, gin.H{"data": err.Error()})
				return
			}
			res = message
		}
	case protocols.P_Fight_Report_Result_To_Fight:
		if websocketDataDecode.ClientType == 1 {
			message := protocols.C_Fight_Report_Result_To_Fight{}
			err = message.Decode(buffer, apis.FIGHT_KEY)
			if err != nil {
				logrus.Error(err)
				c.JSON(http.StatusOK, gin.H{"data": err.Error()})
				return
			}
			res = message
		} else if websocketDataDecode.ClientType == 2 {
			message := protocols.S_Fight_Report_Result_To_Fight{}
			err = message.Decode(buffer)
			if err != nil {
				logrus.Error(err)
				c.JSON(http.StatusOK, gin.H{"data": err.Error()})
				return
			}
			res = message
		}
	case protocols.P_Fight_Report_Phase_Result_To_Fight:
		if websocketDataDecode.ClientType == 1 {
			message := protocols.C_Fight_Report_Result_To_Fight{}
			err = message.Decode(buffer, apis.FIGHT_KEY)
			if err != nil {
				logrus.Error(err)
				c.JSON(http.StatusOK, gin.H{"data": err.Error()})
				return
			}
			res = message
		} else if websocketDataDecode.ClientType == 2 {
			message := protocols.S_Fight_Report_Result_To_Fight{}
			err = message.Decode(buffer)
			if err != nil {
				logrus.Error(err)
				c.JSON(http.StatusOK, gin.H{"data": err.Error()})
				return
			}
			res = message
		}
	case protocols.P_Fight_Update_Hero_SYNC:
		if websocketDataDecode.ClientType == 1 {
			message := protocols.C_Fight_Update_Hero_SYNC{}
			err = message.Decode(buffer, apis.FIGHT_KEY)
			if err != nil {
				logrus.Error(err)
				c.JSON(http.StatusOK, gin.H{"data": err.Error()})
				return
			}
			res = message
		} else if websocketDataDecode.ClientType == 2 {
			message := protocols.S_Fight_Update_Hero_SYNC{}
			err = message.Decode(buffer)
			if err != nil {
				logrus.Error(err)
				c.JSON(http.StatusOK, gin.H{"data": err.Error()})
				return
			}
			res = message
		}
	case protocols.P_Fight_Operate_Equip_SYNC:
		if websocketDataDecode.ClientType == 1 {
			message := protocols.C_Fight_Operate_Equip_SYNC{}
			err = message.Decode(buffer, apis.FIGHT_KEY)
			if err != nil {
				logrus.Error(err)
				c.JSON(http.StatusOK, gin.H{"data": err.Error()})
				return
			}
			res = message
		} else if websocketDataDecode.ClientType == 2 {
			message := protocols.S_Fight_Operate_Equip_SYNC{}
			err = message.Decode(buffer)
			if err != nil {
				logrus.Error(err)
				c.JSON(http.StatusOK, gin.H{"data": err.Error()})
				return
			}
			res = message
		}
	default:
		logrus.Info("protocolNum default: ", websocketDataDecode.ProtocolNum)
	}

	byteRes, err := json.Marshal(&res)
	if err != nil {
		logrus.Error(err)
		c.JSON(http.StatusOK, gin.H{"data": err.Error()})
		return
	}
	logrus.Info(string(byteRes))
	c.JSON(http.StatusOK, gin.H{"data": string(byteRes)})
}

// 创建用户角色数据（需要登录test账号）
func create(c *gin.Context) {
	token := c.GetString("token")
	curUser := db.DbManager.FindUserByToken(token)
	if curUser == (models.User{}) && curUser.Account != "test" {
		c.JSON(http.StatusOK, gin.H{"error": 1})
		return
	}
	account := utils.GetRandomString(6)
	var user = models.User{
		ID_:                primitive.NewObjectID(),
		Account:            account,
		PasswordCiphertext: fmt.Sprintf("%x", md5.Sum([]byte(account))),
	}
	err := db.DbManager.CreateUser(user)
	if err != nil {
		logrus.Error("db.DbManager.CreateUser: ", err)
		c.JSON(http.StatusOK, gin.H{"data": "创建用户失败！"})
		return
	}
	var role models.Role
	if err := db.DbManager.FindOneRole(bson.M{"account": "test"}, &role); err != nil {
		logrus.Error("db.DbManager.FindOneRole: ", err)
		c.JSON(http.StatusOK, gin.H{"data": "创建角色失败！"})
		return
	}
	if role != (models.Role{}) {
		roleID := role.RoleID
		role.ID_ = primitive.NewObjectID()
		role.RoleID = role.RoleID + 1
		role.Account = account
		role.StrID = utils.GetShowID(role.RoleID)
		role.RoleName = utils.GetRoleName(role.RoleID % 10000)
		role.Key = utils.GetRandomKey()
		err = db.DbManager.CreateRole(role)
		if err != nil {
			logrus.Error("db.DbManager.CreateRole: ", err)
			c.JSON(http.StatusOK, gin.H{"data": "创建角色失败！"})
			return
		}

		// 复制role_attr_value_items
		var roleAttrValueItems = []protocols.S_Role_SynRoleAttrValue{}
		err = db.DbManager.FindRoleAttrValueItemsByRoleID(roleID, &roleAttrValueItems)
		if err != nil {
			logrus.Error("db.DbManager.FindRoleAttrValueItems: ", err)
			c.JSON(http.StatusOK, gin.H{"data": "复制AttrValue失败！"})
			return
		}
		for _, roleAttrValueItem := range roleAttrValueItems {
			roleAttrValueItem.ID_ = primitive.NewObjectID()
			roleAttrValueItem.RoleID = role.RoleID
			_, err = db.DbManager.CreateRoleAttrValueItem(roleAttrValueItem)
			if err != nil {
				logrus.Error("db.DbManager.CreateRoleAttrValueItem: ", err)
				c.JSON(http.StatusOK, gin.H{"data": "创建AttrValue失败！"})
				return
			}
		}

		// 复制role_information
		var roleInformation = db.DbManager.FindRoleInformationByRoleID(roleID)
		roleInformation.ID_ = primitive.NewObjectID()
		roleInformation.RoleID = role.RoleID
		_, err = db.DbManager.InsertOneRoleInformation(roleInformation)
		if err != nil {
			logrus.Error("db.DbManager.InsertOneRoleInformation: ", err)
			c.JSON(http.StatusOK, gin.H{"data": "复制RoleInformation失败！"})
			return
		}

		// 复制role_bag_items
		var roleBagItems = []protocols.T_Role_Item{}
		err = db.DbManager.FindRoleBagItemsByRoleID(roleID, &roleBagItems)
		if err != nil {
			logrus.Error("db.DbManager.FindRoleBagItems: ", err)
			c.JSON(http.StatusOK, gin.H{"data": "复制RoleBag失败！"})
			return
		}
		for _, roleBagItem := range roleBagItems {
			roleBagItem.ID_ = primitive.NewObjectID()
			roleBagItem.RoleID = role.RoleID
			_, err = db.DbManager.CreateRoleBagItem(roleBagItem)
			if err != nil {
				logrus.Error("db.DbManager.CreateRoleBagItem: ", err)
				c.JSON(http.StatusOK, gin.H{"data": "复制RoleBag失败！"})
				return
			}
		}

		// 复制role_battle_array
		var roleBattleArray = []protocols.T_Role_BattleArrayIndexData{}
		err = db.DbManager.FindRoleBattleArrayByRoleID(roleID, &roleBattleArray)
		if err != nil {
			logrus.Error("db.DbManager.FindRoleBattleArrayByRoleID: ", err)
			c.JSON(http.StatusOK, gin.H{"data": "复制RoleBattleArray失败！"})
			return
		}
		for _, roleBattle := range roleBattleArray {
			roleBattle.ID_ = primitive.NewObjectID()
			roleBattle.RoleID = role.RoleID
			_, err = db.DbManager.InsertOneRoleBattleArray(roleBattle)
			if err != nil {
				logrus.Error("db.DbManager.InsertOneRoleBattleArray: ", err)
				c.JSON(http.StatusOK, gin.H{"data": "复制RoleBattleArray失败！"})
				return
			}
		}

		// 复制role_season_items
		var seasonEntityArray = []protocols.T_SeasonEntityData{}
		err = db.DbManager.FindRoleSeasonItemsByRoleID(roleID, &seasonEntityArray)
		if err != nil {
			logrus.Error("db.DbManager.FindRoleSeasonItemsByRoleID: ", err)
			c.JSON(http.StatusOK, gin.H{"data": "复制RoleSeasonItems失败！"})
			return
		}
		for _, seasonEntity := range seasonEntityArray {
			seasonEntity.ID_ = primitive.NewObjectID()
			seasonEntity.RoleID = role.RoleID
			_, err = db.DbManager.InsertOneSeason(seasonEntity)
			if err != nil {
				logrus.Error("db.DbManager.InsertOneSeason: ", err)
				c.JSON(http.StatusOK, gin.H{"data": "复制RoleSeasonItems失败！"})
				return
			}
		}

		// 复制role_season_forever_score_prize
		var seasonForeverScorePrizeEntityArray = []models.RoleSeasonForeverScorePrize{}
		err = db.DbManager.FindRoleSeasonForeverScorePrizeByRoleID(roleID, &seasonForeverScorePrizeEntityArray)
		if err != nil {
			logrus.Error("db.DbManager.FindRoleSeasonForeverScorePrizeByRoleID: ", err)
			c.JSON(http.StatusOK, gin.H{"data": "复制RoleSeasonForeverScorePrize失败！"})
			return
		}
		for _, seasonForeverScorePrizeEntity := range seasonForeverScorePrizeEntityArray {
			seasonForeverScorePrizeEntity.ID_ = primitive.NewObjectID()
			seasonForeverScorePrizeEntity.RoleID = role.RoleID
			_, err = db.DbManager.InsertOneRoleSeasonForeverScorePrize(seasonForeverScorePrizeEntity)
			if err != nil {
				logrus.Error("db.DbManager.InsertOneRoleSeasonForeverScorePrize: ", err)
				c.JSON(http.StatusOK, gin.H{"data": "复制RoleSeasonForeverScorePrize失败！"})
				return
			}
		}

		// 复制role_season_score_prize
		var seasonScorePrizeEntityArray = []models.RoleSeasonScorePrize{}
		err = db.DbManager.FindRoleSeasonScorePrizeByRoleID(roleID, &seasonScorePrizeEntityArray)
		if err != nil {
			logrus.Error("db.DbManager.FindRoleSeasonScorePrizeByRoleID: ", err)
			c.JSON(http.StatusOK, gin.H{"data": "复制RoleSeasonScorePrize失败！"})
			return
		}
		for _, seasonScorePrizeEntity := range seasonScorePrizeEntityArray {
			seasonScorePrizeEntity.ID_ = primitive.NewObjectID()
			seasonScorePrizeEntity.RoleID = role.RoleID
			_, err = db.DbManager.InsertOneRoleSeasonScorePrize(seasonScorePrizeEntity)
			if err != nil {
				logrus.Error("db.DbManager.InsertOneRoleSeasonScorePrize: ", err)
				c.JSON(http.StatusOK, gin.H{"data": "复制RoleSeasonScorePrize失败！"})
				return
			}
		}

		// 复制role_task_items
		var roleSingleTaskArray = []protocols.T_Role_SingleTask{}
		err = db.DbManager.FindRoleTaskItemsByRoleID(roleID, &roleSingleTaskArray)
		if err != nil {
			logrus.Error("db.DbManager.FindRoleTaskItemsByRoleID: ", err)
			c.JSON(http.StatusOK, gin.H{"data": "复制RoleTaskItems失败！"})
			return
		}
		for _, roleSingleTask := range roleSingleTaskArray {
			roleSingleTask.ID_ = primitive.NewObjectID()
			roleSingleTask.RoleID = role.RoleID
			_, err = db.DbManager.InsertOneRoleSingleTask(roleSingleTask)
			if err != nil {
				logrus.Error("db.DbManager.InsertOneRoleSingleTask: ", err)
				c.JSON(http.StatusOK, gin.H{"data": "复制RoleTaskItems失败！"})
				return
			}
		}

		// 复制role_hero_skin
		var roleHeroSkins = []models.RoleHeroSkin{}
		err = db.DbManager.FindRoleHeroSkinByRoleID(roleID, &roleHeroSkins)
		if err != nil {
			logrus.Error("db.DbManager.FindRoleHeroSkinByRoleID: ", err)
			c.JSON(http.StatusOK, gin.H{"data": "复制RoleHeroSkin失败！"})
			return
		}
		for _, roleHeroSkin := range roleHeroSkins {
			roleHeroSkin.ID_ = primitive.NewObjectID()
			roleHeroSkin.RoleID = role.RoleID
			_, err = db.DbManager.InsertOneRoleHeroSkin(roleHeroSkin)
			if err != nil {
				logrus.Error("db.DbManager.InsertOneRoleHeroSkin: ", err)
				c.JSON(http.StatusOK, gin.H{"data": "复制RoleHeroSkin失败！"})
				return
			}
		}
		c.JSON(http.StatusOK, gin.H{"data": "复制成功！"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": "复制失败！"})
}

// 删除用户角色数据（需要登录test账号）
func delete(c *gin.Context) {
	token := c.GetString("token")
	curUser := db.DbManager.FindUserByToken(token)
	if curUser == (models.User{}) && curUser.Account != "test" {
		c.JSON(http.StatusOK, gin.H{"error": 1})
		return
	}
	account := c.Query("account")
	if account == "" {
		c.JSON(http.StatusOK, gin.H{"error": 1})
		return
	}
	deleteResult, err := db.DbManager.DeleteUserByAccount(account)
	if err != nil {
		logrus.Error("db.DbManager.DeleteUserByAccount: ", err)
		c.JSON(http.StatusOK, gin.H{"data": "删除用户失败！"})
		return
	}
	if deleteResult.DeletedCount == 0 {
		c.JSON(http.StatusOK, gin.H{"data": "删除用户失败！"})
		return
	}
	var role models.Role
	if err := db.DbManager.FindOneRole(bson.M{"account": account}, &role); err != nil {
		logrus.Error("db.DbManager.FindOneRole: ", err)
		c.JSON(http.StatusOK, gin.H{"data": "删除角色失败！"})
		return
	}
	if role != (models.Role{}) {
		roleID := role.RoleID
		deleteCount, err := db.DbManager.DeleteRoleByAccount(account)
		if err != nil {
			logrus.Error("db.DbManager.DeleteRoleByAccount: ", err)
			c.JSON(http.StatusOK, gin.H{"data": "删除角色失败！"})
			return
		}
		if deleteCount == 0 {
			c.JSON(http.StatusOK, gin.H{"data": "删除角色失败！"})
			return
		}
		// 删除role_attr_value_items
		deleteResult, err = db.DbManager.DeleteRoleAttrValueItemsByRoleID(roleID)
		if err != nil {
			logrus.Error("db.DbManager.DeleteRoleAttrValueItemsByRoleID: ", err)
			c.JSON(http.StatusOK, gin.H{"data": "删除RoleAttrValueItems失败！"})
			return
		}
		if deleteResult.DeletedCount == 0 {
			c.JSON(http.StatusOK, gin.H{"data": "删除RoleAttrValueItems失败！"})
			return
		}
		// 删除role_information
		deleteResult, err = db.DbManager.DeleteRoleInformationByRoleID(roleID)
		if err != nil {
			logrus.Error("db.DbManager.DeleteRoleInformationByRoleID: ", err)
			c.JSON(http.StatusOK, gin.H{"data": "删除RoleInformation失败！"})
			return
		}
		if deleteResult.DeletedCount == 0 {
			c.JSON(http.StatusOK, gin.H{"data": "删除RoleInformation失败！"})
			return
		}
		// 删除role_bag_items
		deleteResult, err = db.DbManager.DeleteRoleBagItemByRoleID(roleID)
		if err != nil {
			logrus.Error("db.DbManager.DeleteRoleBagItemByRoleID: ", err)
			c.JSON(http.StatusOK, gin.H{"data": "删除RoleBagItem失败！"})
			return
		}
		if deleteResult.DeletedCount == 0 {
			c.JSON(http.StatusOK, gin.H{"data": "删除RoleBagItem失败！"})
			return
		}
		// 删除role_battle_array
		deleteResult, err = db.DbManager.DeleteRoleBattleArrayByRoleID(roleID)
		if err != nil {
			logrus.Error("db.DbManager.DeleteRoleBattleArrayByRoleID: ", err)
			c.JSON(http.StatusOK, gin.H{"data": "删除RoleBattleArray失败！"})
			return
		}
		if deleteResult.DeletedCount == 0 {
			c.JSON(http.StatusOK, gin.H{"data": "删除RoleBattleArray失败！"})
			return
		}
		// 删除role_season_items
		deleteResult, err = db.DbManager.DeleteRoleSeasonItemsByRoleID(roleID)
		if err != nil {
			logrus.Error("db.DbManager.DeleteRoleSeasonItemsByRoleID: ", err)
			c.JSON(http.StatusOK, gin.H{"data": "删除RoleSeasonItems失败！"})
			return
		}
		if deleteResult.DeletedCount == 0 {
			c.JSON(http.StatusOK, gin.H{"data": "删除RoleSeasonItems失败！"})
			return
		}
		// 删除role_season_forever_score_prize
		deleteResult, err = db.DbManager.DeleteRoleSeasonForeverScorePrizeByRoleID(roleID)
		if err != nil {
			logrus.Error("db.DbManager.DeleteRoleSeasonForeverScorePrizeByRoleID: ", err)
			c.JSON(http.StatusOK, gin.H{"data": "删除RoleSeasonForeverScorePrize失败！"})
			return
		}
		if deleteResult.DeletedCount == 0 {
			c.JSON(http.StatusOK, gin.H{"data": "删除RoleSeasonForeverScorePrize失败！"})
			return
		}
		// 删除role_season_score_prize
		deleteResult, err = db.DbManager.DeleteRoleSeasonScorePrizeByRoleID(roleID)
		if err != nil {
			logrus.Error("db.DbManager.DeleteRoleSeasonScorePrizeByRoleID: ", err)
			c.JSON(http.StatusOK, gin.H{"data": "删除RoleSeasonScorePrize失败！"})
			return
		}
		if deleteResult.DeletedCount == 0 {
			c.JSON(http.StatusOK, gin.H{"data": "删除RoleSeasonScorePrize失败！"})
			return
		}
		// 删除role_task_items
		deleteResult, err = db.DbManager.DeleteRoleTaskItemsByRoleID(roleID)
		if err != nil {
			logrus.Error("db.DbManager.DeleteRoleTaskItemsByRoleID: ", err)
			c.JSON(http.StatusOK, gin.H{"data": "删除RoleTaskItems失败！"})
			return
		}
		if deleteResult.DeletedCount == 0 {
			c.JSON(http.StatusOK, gin.H{"data": "删除RoleTaskItems失败！"})
			return
		}
		// 删除role_hero_skin
		deleteResult, err = db.DbManager.DeleteRoleHeroSkinByRoleID(roleID)
		if err != nil {
			logrus.Error("db.DbManager.DeleteRoleHeroSkinByRoleID: ", err)
			c.JSON(http.StatusOK, gin.H{"data": "删除RoleHeroSkin失败！"})
			return
		}
		if deleteResult.DeletedCount == 0 {
			c.JSON(http.StatusOK, gin.H{"data": "删除RoleHeroSkin失败！"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": "删除角色成功！"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": "删除角色失败！"})
}
