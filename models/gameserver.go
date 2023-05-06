package models

type Data struct {
	OpenID  string `json:"open_id"`
	Tishen  int    `json:"tishen"`
	ExtJSON string `json:"extjson"`
	SdkType int    `json:"sdktype"`
	SdkID   int    `json:"sdkid"`
}

type Server struct {
	GroupName string `json:"groupname"`
	GroupID   int    `json:"groupid"`
	State     int    `json:"state"`
	Roles     []int  `json:"roles"`
}

type Result struct {
	Data    Data     `json:"data"`
	Servers []Server `json:"servers"`
}
