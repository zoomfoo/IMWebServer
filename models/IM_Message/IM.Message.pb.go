// Code generated by protoc-gen-go.
// source: IM.Message.proto
// DO NOT EDIT!

/*
Package IM_Message is a generated protocol buffer package.

It is generated from these files:
	IM.Message.proto

It has these top-level messages:
	IMMsgData
	IMMsgDataAck
	IMMsgDataReadAck
	IMMsgDataReadNotify
	IMClientTimeReq
	IMClientTimeRsp
	IMUnreadMsgCntReq
	IMUnreadMsgCntRsp
	IMGetMsgListReq
	IMGetMsgListRsp
	IMGetLatestMsgIdReq
	IMGetLatestMsgIdRsp
	IMGetMsgByIdReq
	IMGetMsgByIdRsp
*/
package IM_Message

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import "IMWebServer/models/IM_BaseDefine"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

// service id  0x0003
type IMMsgData struct {
	// cmd id:		0x0301
	FromUserId       *uint32                `protobuf:"varint,1,req,name=from_user_id" json:"from_user_id,omitempty"`
	ToSessionId      *uint32                `protobuf:"varint,2,req,name=to_session_id" json:"to_session_id,omitempty"`
	MsgId            *uint32                `protobuf:"varint,3,req,name=msg_id" json:"msg_id,omitempty"`
	CreateTime       *uint32                `protobuf:"varint,4,req,name=create_time" json:"create_time,omitempty"`
	MsgType          *IM_BaseDefine.MsgType `protobuf:"varint,5,req,name=msg_type,enum=IM.BaseDefine.MsgType" json:"msg_type,omitempty"`
	MsgData          []byte                 `protobuf:"bytes,6,req,name=msg_data" json:"msg_data,omitempty"`
	AttachData       []byte                 `protobuf:"bytes,20,opt,name=attach_data" json:"attach_data,omitempty"`
	XXX_unrecognized []byte                 `json:"-"`
}

func (m *IMMsgData) Reset()                    { *m = IMMsgData{} }
func (m *IMMsgData) String() string            { return proto.CompactTextString(m) }
func (*IMMsgData) ProtoMessage()               {}
func (*IMMsgData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *IMMsgData) GetFromUserId() uint32 {
	if m != nil && m.FromUserId != nil {
		return *m.FromUserId
	}
	return 0
}

func (m *IMMsgData) GetToSessionId() uint32 {
	if m != nil && m.ToSessionId != nil {
		return *m.ToSessionId
	}
	return 0
}

func (m *IMMsgData) GetMsgId() uint32 {
	if m != nil && m.MsgId != nil {
		return *m.MsgId
	}
	return 0
}

func (m *IMMsgData) GetCreateTime() uint32 {
	if m != nil && m.CreateTime != nil {
		return *m.CreateTime
	}
	return 0
}

func (m *IMMsgData) GetMsgType() IM_BaseDefine.MsgType {
	if m != nil && m.MsgType != nil {
		return *m.MsgType
	}
	return IM_BaseDefine.MsgType_MSG_TYPE_SINGLE_TEXT
}

func (m *IMMsgData) GetMsgData() []byte {
	if m != nil {
		return m.MsgData
	}
	return nil
}

func (m *IMMsgData) GetAttachData() []byte {
	if m != nil {
		return m.AttachData
	}
	return nil
}

type IMMsgDataAck struct {
	// cmd id:		0x0302
	UserId           *uint32                    `protobuf:"varint,1,req,name=user_id" json:"user_id,omitempty"`
	SessionId        *uint32                    `protobuf:"varint,2,req,name=session_id" json:"session_id,omitempty"`
	MsgId            *uint32                    `protobuf:"varint,3,req,name=msg_id" json:"msg_id,omitempty"`
	SessionType      *IM_BaseDefine.SessionType `protobuf:"varint,4,req,name=session_type,enum=IM.BaseDefine.SessionType" json:"session_type,omitempty"`
	XXX_unrecognized []byte                     `json:"-"`
}

func (m *IMMsgDataAck) Reset()                    { *m = IMMsgDataAck{} }
func (m *IMMsgDataAck) String() string            { return proto.CompactTextString(m) }
func (*IMMsgDataAck) ProtoMessage()               {}
func (*IMMsgDataAck) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *IMMsgDataAck) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *IMMsgDataAck) GetSessionId() uint32 {
	if m != nil && m.SessionId != nil {
		return *m.SessionId
	}
	return 0
}

func (m *IMMsgDataAck) GetMsgId() uint32 {
	if m != nil && m.MsgId != nil {
		return *m.MsgId
	}
	return 0
}

func (m *IMMsgDataAck) GetSessionType() IM_BaseDefine.SessionType {
	if m != nil && m.SessionType != nil {
		return *m.SessionType
	}
	return IM_BaseDefine.SessionType_SESSION_TYPE_SINGLE
}

type IMMsgDataReadAck struct {
	// cmd id:		0x0303
	UserId           *uint32                    `protobuf:"varint,1,req,name=user_id" json:"user_id,omitempty"`
	SessionId        *uint32                    `protobuf:"varint,2,req,name=session_id" json:"session_id,omitempty"`
	MsgId            *uint32                    `protobuf:"varint,3,req,name=msg_id" json:"msg_id,omitempty"`
	SessionType      *IM_BaseDefine.SessionType `protobuf:"varint,4,req,name=session_type,enum=IM.BaseDefine.SessionType" json:"session_type,omitempty"`
	XXX_unrecognized []byte                     `json:"-"`
}

func (m *IMMsgDataReadAck) Reset()                    { *m = IMMsgDataReadAck{} }
func (m *IMMsgDataReadAck) String() string            { return proto.CompactTextString(m) }
func (*IMMsgDataReadAck) ProtoMessage()               {}
func (*IMMsgDataReadAck) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *IMMsgDataReadAck) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *IMMsgDataReadAck) GetSessionId() uint32 {
	if m != nil && m.SessionId != nil {
		return *m.SessionId
	}
	return 0
}

func (m *IMMsgDataReadAck) GetMsgId() uint32 {
	if m != nil && m.MsgId != nil {
		return *m.MsgId
	}
	return 0
}

func (m *IMMsgDataReadAck) GetSessionType() IM_BaseDefine.SessionType {
	if m != nil && m.SessionType != nil {
		return *m.SessionType
	}
	return IM_BaseDefine.SessionType_SESSION_TYPE_SINGLE
}

type IMMsgDataReadNotify struct {
	// cmd id:		0x0304
	UserId           *uint32                    `protobuf:"varint,1,req,name=user_id" json:"user_id,omitempty"`
	SessionId        *uint32                    `protobuf:"varint,2,req,name=session_id" json:"session_id,omitempty"`
	MsgId            *uint32                    `protobuf:"varint,3,req,name=msg_id" json:"msg_id,omitempty"`
	SessionType      *IM_BaseDefine.SessionType `protobuf:"varint,4,req,name=session_type,enum=IM.BaseDefine.SessionType" json:"session_type,omitempty"`
	XXX_unrecognized []byte                     `json:"-"`
}

func (m *IMMsgDataReadNotify) Reset()                    { *m = IMMsgDataReadNotify{} }
func (m *IMMsgDataReadNotify) String() string            { return proto.CompactTextString(m) }
func (*IMMsgDataReadNotify) ProtoMessage()               {}
func (*IMMsgDataReadNotify) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *IMMsgDataReadNotify) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *IMMsgDataReadNotify) GetSessionId() uint32 {
	if m != nil && m.SessionId != nil {
		return *m.SessionId
	}
	return 0
}

func (m *IMMsgDataReadNotify) GetMsgId() uint32 {
	if m != nil && m.MsgId != nil {
		return *m.MsgId
	}
	return 0
}

func (m *IMMsgDataReadNotify) GetSessionType() IM_BaseDefine.SessionType {
	if m != nil && m.SessionType != nil {
		return *m.SessionType
	}
	return IM_BaseDefine.SessionType_SESSION_TYPE_SINGLE
}

type IMClientTimeReq struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *IMClientTimeReq) Reset()                    { *m = IMClientTimeReq{} }
func (m *IMClientTimeReq) String() string            { return proto.CompactTextString(m) }
func (*IMClientTimeReq) ProtoMessage()               {}
func (*IMClientTimeReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type IMClientTimeRsp struct {
	// cmd id:		0x0306
	ServerTime       *uint32 `protobuf:"varint,1,req,name=server_time" json:"server_time,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *IMClientTimeRsp) Reset()                    { *m = IMClientTimeRsp{} }
func (m *IMClientTimeRsp) String() string            { return proto.CompactTextString(m) }
func (*IMClientTimeRsp) ProtoMessage()               {}
func (*IMClientTimeRsp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *IMClientTimeRsp) GetServerTime() uint32 {
	if m != nil && m.ServerTime != nil {
		return *m.ServerTime
	}
	return 0
}

type IMUnreadMsgCntReq struct {
	// cmd id:		0x0307
	UserId           *uint32 `protobuf:"varint,1,req,name=user_id" json:"user_id,omitempty"`
	AttachData       []byte  `protobuf:"bytes,20,opt,name=attach_data" json:"attach_data,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *IMUnreadMsgCntReq) Reset()                    { *m = IMUnreadMsgCntReq{} }
func (m *IMUnreadMsgCntReq) String() string            { return proto.CompactTextString(m) }
func (*IMUnreadMsgCntReq) ProtoMessage()               {}
func (*IMUnreadMsgCntReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *IMUnreadMsgCntReq) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *IMUnreadMsgCntReq) GetAttachData() []byte {
	if m != nil {
		return m.AttachData
	}
	return nil
}

type IMUnreadMsgCntRsp struct {
	// cmd id:		0x0308
	UserId           *uint32                     `protobuf:"varint,1,req,name=user_id" json:"user_id,omitempty"`
	TotalCnt         *uint32                     `protobuf:"varint,2,req,name=total_cnt" json:"total_cnt,omitempty"`
	UnreadinfoList   []*IM_BaseDefine.UnreadInfo `protobuf:"bytes,3,rep,name=unreadinfo_list" json:"unreadinfo_list,omitempty"`
	AttachData       []byte                      `protobuf:"bytes,20,opt,name=attach_data" json:"attach_data,omitempty"`
	XXX_unrecognized []byte                      `json:"-"`
}

func (m *IMUnreadMsgCntRsp) Reset()                    { *m = IMUnreadMsgCntRsp{} }
func (m *IMUnreadMsgCntRsp) String() string            { return proto.CompactTextString(m) }
func (*IMUnreadMsgCntRsp) ProtoMessage()               {}
func (*IMUnreadMsgCntRsp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *IMUnreadMsgCntRsp) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *IMUnreadMsgCntRsp) GetTotalCnt() uint32 {
	if m != nil && m.TotalCnt != nil {
		return *m.TotalCnt
	}
	return 0
}

func (m *IMUnreadMsgCntRsp) GetUnreadinfoList() []*IM_BaseDefine.UnreadInfo {
	if m != nil {
		return m.UnreadinfoList
	}
	return nil
}

func (m *IMUnreadMsgCntRsp) GetAttachData() []byte {
	if m != nil {
		return m.AttachData
	}
	return nil
}

type IMGetMsgListReq struct {
	// cmd id:		0x0309
	UserId           *uint32                    `protobuf:"varint,1,req,name=user_id" json:"user_id,omitempty"`
	SessionType      *IM_BaseDefine.SessionType `protobuf:"varint,2,req,name=session_type,enum=IM.BaseDefine.SessionType" json:"session_type,omitempty"`
	SessionId        *uint32                    `protobuf:"varint,3,req,name=session_id" json:"session_id,omitempty"`
	MsgIdBegin       *uint32                    `protobuf:"varint,4,req,name=msg_id_begin" json:"msg_id_begin,omitempty"`
	MsgCnt           *uint32                    `protobuf:"varint,5,req,name=msg_cnt" json:"msg_cnt,omitempty"`
	MsgType          *uint32                    `protobuf:"varint,6,opt,name=msg_type" json:"msg_type,omitempty"`
	AttachData       []byte                     `protobuf:"bytes,20,opt,name=attach_data" json:"attach_data,omitempty"`
	XXX_unrecognized []byte                     `json:"-"`
}

func (m *IMGetMsgListReq) Reset()                    { *m = IMGetMsgListReq{} }
func (m *IMGetMsgListReq) String() string            { return proto.CompactTextString(m) }
func (*IMGetMsgListReq) ProtoMessage()               {}
func (*IMGetMsgListReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *IMGetMsgListReq) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *IMGetMsgListReq) GetSessionType() IM_BaseDefine.SessionType {
	if m != nil && m.SessionType != nil {
		return *m.SessionType
	}
	return IM_BaseDefine.SessionType_SESSION_TYPE_SINGLE
}

func (m *IMGetMsgListReq) GetSessionId() uint32 {
	if m != nil && m.SessionId != nil {
		return *m.SessionId
	}
	return 0
}

func (m *IMGetMsgListReq) GetMsgIdBegin() uint32 {
	if m != nil && m.MsgIdBegin != nil {
		return *m.MsgIdBegin
	}
	return 0
}

func (m *IMGetMsgListReq) GetMsgCnt() uint32 {
	if m != nil && m.MsgCnt != nil {
		return *m.MsgCnt
	}
	return 0
}

func (m *IMGetMsgListReq) GetMsgType() uint32 {
	if m != nil && m.MsgType != nil {
		return *m.MsgType
	}
	return 0
}

func (m *IMGetMsgListReq) GetAttachData() []byte {
	if m != nil {
		return m.AttachData
	}
	return nil
}

// 对于群而言，如果消息数目返回的数值小于请求的cnt,则表示群的消息能拉取的到头了，更早的消息没有权限拉取。
// 如果msg_cnt 和 msg_id_begin计算得到的最早消息id与实际返回的最早消息id不一致，说明服务器消息有缺失，需要
// 客户端做一个缺失标记，避免下次再次拉取。
type IMGetMsgListRsp struct {
	// cmd id:		0x030a
	UserId           *uint32                    `protobuf:"varint,1,req,name=user_id" json:"user_id,omitempty"`
	SessionType      *IM_BaseDefine.SessionType `protobuf:"varint,2,req,name=session_type,enum=IM.BaseDefine.SessionType" json:"session_type,omitempty"`
	SessionId        *uint32                    `protobuf:"varint,3,req,name=session_id" json:"session_id,omitempty"`
	MsgIdBegin       *uint32                    `protobuf:"varint,4,req,name=msg_id_begin" json:"msg_id_begin,omitempty"`
	MsgList          []*IM_BaseDefine.MsgInfo   `protobuf:"bytes,5,rep,name=msg_list" json:"msg_list,omitempty"`
	AttachData       []byte                     `protobuf:"bytes,20,opt,name=attach_data" json:"attach_data,omitempty"`
	XXX_unrecognized []byte                     `json:"-"`
}

func (m *IMGetMsgListRsp) Reset()                    { *m = IMGetMsgListRsp{} }
func (m *IMGetMsgListRsp) String() string            { return proto.CompactTextString(m) }
func (*IMGetMsgListRsp) ProtoMessage()               {}
func (*IMGetMsgListRsp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *IMGetMsgListRsp) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *IMGetMsgListRsp) GetSessionType() IM_BaseDefine.SessionType {
	if m != nil && m.SessionType != nil {
		return *m.SessionType
	}
	return IM_BaseDefine.SessionType_SESSION_TYPE_SINGLE
}

func (m *IMGetMsgListRsp) GetSessionId() uint32 {
	if m != nil && m.SessionId != nil {
		return *m.SessionId
	}
	return 0
}

func (m *IMGetMsgListRsp) GetMsgIdBegin() uint32 {
	if m != nil && m.MsgIdBegin != nil {
		return *m.MsgIdBegin
	}
	return 0
}

func (m *IMGetMsgListRsp) GetMsgList() []*IM_BaseDefine.MsgInfo {
	if m != nil {
		return m.MsgList
	}
	return nil
}

func (m *IMGetMsgListRsp) GetAttachData() []byte {
	if m != nil {
		return m.AttachData
	}
	return nil
}

type IMGetLatestMsgIdReq struct {
	// cmd id: 		0x030b
	UserId           *uint32                    `protobuf:"varint,1,req,name=user_id" json:"user_id,omitempty"`
	SessionType      *IM_BaseDefine.SessionType `protobuf:"varint,2,req,name=session_type,enum=IM.BaseDefine.SessionType" json:"session_type,omitempty"`
	SessionId        *uint32                    `protobuf:"varint,3,req,name=session_id" json:"session_id,omitempty"`
	AttachData       []byte                     `protobuf:"bytes,20,opt,name=attach_data" json:"attach_data,omitempty"`
	XXX_unrecognized []byte                     `json:"-"`
}

func (m *IMGetLatestMsgIdReq) Reset()                    { *m = IMGetLatestMsgIdReq{} }
func (m *IMGetLatestMsgIdReq) String() string            { return proto.CompactTextString(m) }
func (*IMGetLatestMsgIdReq) ProtoMessage()               {}
func (*IMGetLatestMsgIdReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *IMGetLatestMsgIdReq) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *IMGetLatestMsgIdReq) GetSessionType() IM_BaseDefine.SessionType {
	if m != nil && m.SessionType != nil {
		return *m.SessionType
	}
	return IM_BaseDefine.SessionType_SESSION_TYPE_SINGLE
}

func (m *IMGetLatestMsgIdReq) GetSessionId() uint32 {
	if m != nil && m.SessionId != nil {
		return *m.SessionId
	}
	return 0
}

func (m *IMGetLatestMsgIdReq) GetAttachData() []byte {
	if m != nil {
		return m.AttachData
	}
	return nil
}

type IMGetLatestMsgIdRsp struct {
	// cmd id:		0x030c
	UserId           *uint32                    `protobuf:"varint,1,req,name=user_id" json:"user_id,omitempty"`
	SessionType      *IM_BaseDefine.SessionType `protobuf:"varint,2,req,name=session_type,enum=IM.BaseDefine.SessionType" json:"session_type,omitempty"`
	SessionId        *uint32                    `protobuf:"varint,3,req,name=session_id" json:"session_id,omitempty"`
	LatestMsgId      *uint32                    `protobuf:"varint,4,req,name=latest_msg_id" json:"latest_msg_id,omitempty"`
	AttachData       []byte                     `protobuf:"bytes,20,opt,name=attach_data" json:"attach_data,omitempty"`
	XXX_unrecognized []byte                     `json:"-"`
}

func (m *IMGetLatestMsgIdRsp) Reset()                    { *m = IMGetLatestMsgIdRsp{} }
func (m *IMGetLatestMsgIdRsp) String() string            { return proto.CompactTextString(m) }
func (*IMGetLatestMsgIdRsp) ProtoMessage()               {}
func (*IMGetLatestMsgIdRsp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *IMGetLatestMsgIdRsp) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *IMGetLatestMsgIdRsp) GetSessionType() IM_BaseDefine.SessionType {
	if m != nil && m.SessionType != nil {
		return *m.SessionType
	}
	return IM_BaseDefine.SessionType_SESSION_TYPE_SINGLE
}

func (m *IMGetLatestMsgIdRsp) GetSessionId() uint32 {
	if m != nil && m.SessionId != nil {
		return *m.SessionId
	}
	return 0
}

func (m *IMGetLatestMsgIdRsp) GetLatestMsgId() uint32 {
	if m != nil && m.LatestMsgId != nil {
		return *m.LatestMsgId
	}
	return 0
}

func (m *IMGetLatestMsgIdRsp) GetAttachData() []byte {
	if m != nil {
		return m.AttachData
	}
	return nil
}

type IMGetMsgByIdReq struct {
	// cmd id: 		0x030d
	UserId           *uint32                    `protobuf:"varint,1,req,name=user_id" json:"user_id,omitempty"`
	SessionType      *IM_BaseDefine.SessionType `protobuf:"varint,2,req,name=session_type,enum=IM.BaseDefine.SessionType" json:"session_type,omitempty"`
	SessionId        *uint32                    `protobuf:"varint,3,req,name=session_id" json:"session_id,omitempty"`
	MsgIdList        []uint32                   `protobuf:"varint,4,rep,name=msg_id_list" json:"msg_id_list,omitempty"`
	AttachData       []byte                     `protobuf:"bytes,20,opt,name=attach_data" json:"attach_data,omitempty"`
	XXX_unrecognized []byte                     `json:"-"`
}

func (m *IMGetMsgByIdReq) Reset()                    { *m = IMGetMsgByIdReq{} }
func (m *IMGetMsgByIdReq) String() string            { return proto.CompactTextString(m) }
func (*IMGetMsgByIdReq) ProtoMessage()               {}
func (*IMGetMsgByIdReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *IMGetMsgByIdReq) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *IMGetMsgByIdReq) GetSessionType() IM_BaseDefine.SessionType {
	if m != nil && m.SessionType != nil {
		return *m.SessionType
	}
	return IM_BaseDefine.SessionType_SESSION_TYPE_SINGLE
}

func (m *IMGetMsgByIdReq) GetSessionId() uint32 {
	if m != nil && m.SessionId != nil {
		return *m.SessionId
	}
	return 0
}

func (m *IMGetMsgByIdReq) GetMsgIdList() []uint32 {
	if m != nil {
		return m.MsgIdList
	}
	return nil
}

func (m *IMGetMsgByIdReq) GetAttachData() []byte {
	if m != nil {
		return m.AttachData
	}
	return nil
}

type IMGetMsgByIdRsp struct {
	// cmd id:		0x030e
	UserId           *uint32                    `protobuf:"varint,1,req,name=user_id" json:"user_id,omitempty"`
	SessionType      *IM_BaseDefine.SessionType `protobuf:"varint,2,req,name=session_type,enum=IM.BaseDefine.SessionType" json:"session_type,omitempty"`
	SessionId        *uint32                    `protobuf:"varint,3,req,name=session_id" json:"session_id,omitempty"`
	MsgList          []*IM_BaseDefine.MsgInfo   `protobuf:"bytes,4,rep,name=msg_list" json:"msg_list,omitempty"`
	AttachData       []byte                     `protobuf:"bytes,20,opt,name=attach_data" json:"attach_data,omitempty"`
	XXX_unrecognized []byte                     `json:"-"`
}

func (m *IMGetMsgByIdRsp) Reset()                    { *m = IMGetMsgByIdRsp{} }
func (m *IMGetMsgByIdRsp) String() string            { return proto.CompactTextString(m) }
func (*IMGetMsgByIdRsp) ProtoMessage()               {}
func (*IMGetMsgByIdRsp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *IMGetMsgByIdRsp) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *IMGetMsgByIdRsp) GetSessionType() IM_BaseDefine.SessionType {
	if m != nil && m.SessionType != nil {
		return *m.SessionType
	}
	return IM_BaseDefine.SessionType_SESSION_TYPE_SINGLE
}

func (m *IMGetMsgByIdRsp) GetSessionId() uint32 {
	if m != nil && m.SessionId != nil {
		return *m.SessionId
	}
	return 0
}

func (m *IMGetMsgByIdRsp) GetMsgList() []*IM_BaseDefine.MsgInfo {
	if m != nil {
		return m.MsgList
	}
	return nil
}

func (m *IMGetMsgByIdRsp) GetAttachData() []byte {
	if m != nil {
		return m.AttachData
	}
	return nil
}

func init() {
	proto.RegisterType((*IMMsgData)(nil), "IM.Message.IMMsgData")
	proto.RegisterType((*IMMsgDataAck)(nil), "IM.Message.IMMsgDataAck")
	proto.RegisterType((*IMMsgDataReadAck)(nil), "IM.Message.IMMsgDataReadAck")
	proto.RegisterType((*IMMsgDataReadNotify)(nil), "IM.Message.IMMsgDataReadNotify")
	proto.RegisterType((*IMClientTimeReq)(nil), "IM.Message.IMClientTimeReq")
	proto.RegisterType((*IMClientTimeRsp)(nil), "IM.Message.IMClientTimeRsp")
	proto.RegisterType((*IMUnreadMsgCntReq)(nil), "IM.Message.IMUnreadMsgCntReq")
	proto.RegisterType((*IMUnreadMsgCntRsp)(nil), "IM.Message.IMUnreadMsgCntRsp")
	proto.RegisterType((*IMGetMsgListReq)(nil), "IM.Message.IMGetMsgListReq")
	proto.RegisterType((*IMGetMsgListRsp)(nil), "IM.Message.IMGetMsgListRsp")
	proto.RegisterType((*IMGetLatestMsgIdReq)(nil), "IM.Message.IMGetLatestMsgIdReq")
	proto.RegisterType((*IMGetLatestMsgIdRsp)(nil), "IM.Message.IMGetLatestMsgIdRsp")
	proto.RegisterType((*IMGetMsgByIdReq)(nil), "IM.Message.IMGetMsgByIdReq")
	proto.RegisterType((*IMGetMsgByIdRsp)(nil), "IM.Message.IMGetMsgByIdRsp")
}

var fileDescriptor0 = []byte{
	// 529 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xbc, 0x95, 0xc1, 0x6e, 0xd3, 0x4c,
	0x10, 0xc7, 0xe5, 0xc4, 0xc9, 0xf7, 0x75, 0xea, 0x36, 0x89, 0x53, 0x20, 0x94, 0x4b, 0xe5, 0x03,
	0xca, 0xc9, 0x42, 0xb9, 0x71, 0x24, 0xad, 0x04, 0x96, 0x6a, 0x0e, 0xa5, 0x9c, 0x2d, 0xd7, 0x5e,
	0x9b, 0x85, 0xd8, 0x1b, 0xb2, 0x63, 0xa4, 0x48, 0x1c, 0xe0, 0x15, 0x38, 0xf1, 0x06, 0xdc, 0x38,
	0x72, 0xe3, 0xdd, 0x98, 0x5d, 0x1b, 0x48, 0x1c, 0x47, 0x15, 0x07, 0xf7, 0x16, 0x4f, 0x66, 0x77,
	0x7e, 0xf3, 0xff, 0xcf, 0xd8, 0x30, 0xf4, 0x7c, 0xd7, 0x67, 0x52, 0x86, 0x29, 0x73, 0x97, 0x2b,
	0x81, 0xc2, 0x86, 0xbf, 0x91, 0xd3, 0x31, 0xfd, 0x9e, 0x87, 0x92, 0x5d, 0xb0, 0x84, 0xe7, 0x55,
	0x82, 0xf3, 0xdd, 0x80, 0x03, 0xcf, 0xf7, 0x65, 0x7a, 0x11, 0x62, 0x68, 0x9f, 0x80, 0x95, 0xac,
	0x44, 0x16, 0x14, 0x92, 0xad, 0x02, 0x1e, 0x4f, 0x8c, 0xb3, 0xce, 0xf4, 0xc8, 0xbe, 0x07, 0x47,
	0x28, 0x02, 0x49, 0xd7, 0x70, 0x91, 0xab, 0x70, 0x47, 0x87, 0x8f, 0xa1, 0x9f, 0xc9, 0x54, 0x3d,
	0x77, 0xf5, 0xf3, 0x18, 0x0e, 0xa3, 0x15, 0x0b, 0x91, 0x05, 0xc8, 0x33, 0x36, 0x31, 0x75, 0x70,
	0x0a, 0xff, 0xab, 0x24, 0x5c, 0x2f, 0xd9, 0xa4, 0x47, 0x91, 0xe3, 0xd9, 0x7d, 0x77, 0x9b, 0x83,
	0x6a, 0x5f, 0xd3, 0xbf, 0xf6, 0xb0, 0xcc, 0x8c, 0x89, 0x63, 0xd2, 0xa7, 0x4c, 0x4b, 0x5d, 0x18,
	0x22, 0x86, 0xd1, 0x9b, 0x32, 0x78, 0x72, 0x66, 0x4c, 0x2d, 0xa7, 0x00, 0xeb, 0x0f, 0xef, 0xb3,
	0xe8, 0x9d, 0x3d, 0x80, 0xff, 0xb6, 0x69, 0x6d, 0x80, 0x5b, 0x51, 0x9f, 0x80, 0xf5, 0x3b, 0x47,
	0x93, 0x99, 0x9a, 0xec, 0xb4, 0x46, 0xf6, 0xaa, 0x4c, 0x51, 0x74, 0xce, 0x5a, 0x89, 0x5b, 0x95,
	0xbd, 0x62, 0x61, 0x7c, 0x87, 0xa5, 0x3f, 0xc2, 0x78, 0xab, 0xf4, 0x4b, 0x81, 0x3c, 0x59, 0xdf,
	0x55, 0xf5, 0x11, 0x0c, 0x3c, 0xff, 0x7c, 0xc1, 0x59, 0x8e, 0xd7, 0x64, 0xeb, 0x15, 0x7b, 0xef,
	0x3c, 0xae, 0x85, 0xe4, 0x52, 0x59, 0x45, 0x2c, 0x1f, 0x08, 0x47, 0x7b, 0xaf, 0x81, 0x9c, 0xa7,
	0x30, 0xf2, 0xfc, 0xd7, 0x39, 0xcd, 0x44, 0x4c, 0xf8, 0xe7, 0x39, 0xd2, 0xe1, 0x5d, 0xec, 0x46,
	0x97, 0x3f, 0x19, 0x3b, 0x67, 0xa9, 0xca, 0xce, 0xd9, 0x11, 0x1c, 0xa0, 0xc0, 0x70, 0x11, 0x44,
	0x39, 0x56, 0x1d, 0xcf, 0x60, 0x50, 0xe8, 0x63, 0x3c, 0x4f, 0x44, 0xb0, 0xe0, 0x12, 0xa9, 0xf5,
	0xee, 0xf4, 0x70, 0xf6, 0xb0, 0xd6, 0x64, 0x79, 0xb9, 0x47, 0x59, 0xcd, 0x08, 0x3f, 0x0c, 0xd5,
	0xe6, 0x73, 0x86, 0x54, 0xff, 0x92, 0xae, 0x69, 0x84, 0xaf, 0xeb, 0xd9, 0xb9, 0x4d, 0xcf, 0x9a,
	0x4b, 0xa5, 0x2b, 0xb4, 0x76, 0xa5, 0x4b, 0xc1, 0x0d, 0x4b, 0x79, 0x5e, 0xad, 0x0e, 0x15, 0x53,
	0x51, 0xd5, 0x5a, 0x4f, 0x07, 0x86, 0x1b, 0xbb, 0xd4, 0x27, 0xc6, 0x3d, 0xda, 0xfd, 0xac, 0x83,
	0x37, 0x29, 0xd7, 0x26, 0x78, 0xb5, 0xf3, 0x5a, 0xfb, 0x9e, 0xd6, 0xbe, 0x61, 0xe7, 0xf7, 0x0b,
	0xff, 0xd9, 0x50, 0x03, 0x4f, 0xfc, 0x97, 0xf4, 0x2e, 0x91, 0xaa, 0x0b, 0x2f, 0x6e, 0x51, 0xfc,
	0x46, 0x86, 0xaf, 0x4d, 0x0c, 0xed, 0xe9, 0x48, 0x6f, 0xd8, 0x85, 0x2e, 0x14, 0x54, 0xdb, 0x6a,
	0xee, 0x47, 0xfb, 0xb2, 0x61, 0xef, 0x7c, 0xdd, 0xb6, 0x34, 0x95, 0xbd, 0xda, 0x4b, 0x93, 0xbc,
	0xdc, 0x03, 0xf5, 0xad, 0x0e, 0xd5, 0x9e, 0x56, 0x9b, 0xd3, 0x65, 0xfe, 0xf3, 0x74, 0xcd, 0x1f,
	0xc1, 0x83, 0x48, 0x64, 0x6e, 0x26, 0xd2, 0xe2, 0x2d, 0x67, 0x2e, 0x62, 0xf9, 0x21, 0xbc, 0x29,
	0x92, 0x17, 0xdd, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x5f, 0x33, 0xae, 0x3f, 0x41, 0x07, 0x00,
	0x00,
}