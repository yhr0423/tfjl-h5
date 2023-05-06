package net

import "tfjl-h5/iface"

type Request struct {
	conn iface.IConnection //已经和客户端建立好的 链接
	msg  iface.IMessage    //客户端请求的数据
}

// 获取请求连接信息
func (r *Request) GetConnection() iface.IConnection {
	return r.conn
}

// 获取请求消息的数据
func (r *Request) GetData() []byte {
	return r.msg.GetData()
}

// 获取请求的消息的ID
func (r *Request) GetMsgType() int {
	return r.msg.GetMsgType()
}

// 获取请求的消息的ID
func (r *Request) GetMsgID() uint32 {
	return r.msg.GetMsgID()
}
