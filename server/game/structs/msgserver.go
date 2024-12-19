package structs

//----------🠋🠋🠋🠋🠋🠋服务端/服务端消息🠋🠋🠋🠋🠋🠋----------

// MsgUpdScore 修改金币消息
type MsgUpdScore struct {
	UserId      int32
	Score       int64
	TypeId      int32
	Description string
}
