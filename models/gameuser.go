package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID_                primitive.ObjectID `bson:"_id"`
	Account            string             `bson:"account"`
	PasswordCiphertext string             `bson:"password_ciphertext"`
	Authorization      string             `bson:"authorization"`
}
type LoginResult struct {
	Error       int         `json:"error"`
	SdkType     int         `json:"sdktype"`
	SdkId       int         `json:"sdkid"`
	AccountName string      `json:"accountname"`
	OpenId      interface{} `json:"openid"`
	Zone        int         `json:"zone"`
	WebName     string      `json:"webname"`
	WebPort     string      `json:"webport"`
	WanIp       string      `json:"wanip"`
	WanPort     string      `json:"wanport"`
	Sign        string      `json:"sign"`
	Examine     int         `json:"examine"`
	AdToShare   int         `json:"adtoshare"`
}
