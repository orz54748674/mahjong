package message

import (
	"mahjong/goim/config"
	"mahjong/goim/core"
	"mahjong/goim/response"
	messageService "mahjong/goim/service/message"
	"net"

	"mahjong/protocal"
)

// PrivateMessage 私聊消息
func PrivateMessage(id string, conn *net.TCPConn, impacket *protocal.ImPacket) {
	if err := messageService.PrivateMessage(id, impacket); err != nil {
		core.Logger.Error("[PrivateMessage]%v", err.Error())
		response.GetErrorData(config.MESSAGE_ID_PRIVATE_MESSAGE_RESPONSE, protocal.MSG_TYPE_RESPONSE, impacket.GetMessageNumber(), err, nil).Send(conn)
	}
}

// RoomMessage 房间消息
func RoomMessage(id string, conn *net.TCPConn, impacket *protocal.ImPacket) {
	if err := messageService.RoomMessage(id, impacket); err != nil {
		core.Logger.Error("[RoomMessage]%v", err.Error())
		response.GetErrorData(config.MESSAGE_ID_ROOM_MESSAGE_RESPONSE, protocal.MSG_TYPE_RESPONSE, impacket.GetMessageNumber(), err, nil).Send(conn)
	}
}

// BroadcastMessage 广播消息
func BroadcastMessage(id string, conn *net.TCPConn, impacket *protocal.ImPacket) {
	if err := messageService.BroadcastMessage(id, impacket); err != nil {
		core.Logger.Error("[BroadcastMessage]%v", err.Error())
		response.GetErrorData(config.MESSAGE_ID_BROADCAST_MESSAGE_RESPONSE, protocal.MSG_TYPE_RESPONSE, impacket.GetMessageNumber(), err, nil).Send(conn)
	}
}
