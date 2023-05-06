package net

type Message struct {
	MsgType int    //消息类型websocket使用
	MsgID   uint32 //业务消息ID
	Data    []byte //消息的内容
	DataLen uint32 //消息的长度
}

// 创建一个Message消息包
func NewMsg(msgType int, msgID uint32, data []byte) *Message {
	return &Message{
		MsgType: msgType,
		MsgID:   msgID,
		Data:    data,
		DataLen: uint32(len(data)),
	}
}

// 获取消息类型
func (msg *Message) GetMsgType() int {
	return msg.MsgType
}

// 获取消息ID
func (msg *Message) GetMsgID() uint32 {
	return msg.MsgID
}

// 获取消息数据段长度
func (msg *Message) GetDataLen() uint32 {
	return msg.DataLen
}

// 获取消息内容
func (msg *Message) GetData() []byte {
	return msg.Data
}
