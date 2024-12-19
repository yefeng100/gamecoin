package structs

//----------🠋🠋🠋🠋🠋🠋客户端/服务端通信消息Login🠋🠋🠋🠋🠋🠋----------

// MsgResErr 返回失败
type MsgResErr struct {
	Code    int32  //错误码
	Content string //错误说明
}

// MsgUserRegReq 客户端注册消息
type MsgUserRegReq struct {
	AccName string `json:"accname"` //账号（游客登录为机器码）
	AccPwd  string `json:"accpwd"`  //密码（游客登录为字符串空）
	Machine string `json:"machine"` //机器码
	RegPlat int8   `json:"regplat"` //注册平台(0:未知,1:安卓,2:IOS,3:WEB,4:PC)
	RegType int8   `json:"regtype"` //注册方式(0:游客,1:账号注册,3:微信,...)
}

// MsgUserLoginReq 客户端登录消息
type MsgUserLoginReq struct {
	AccName   string `json:"accname"`   //账号（游客登录为机器码）
	AccPwd    string `json:"accpwd"`    //密码（游客登录为字符串空）
	LoginType int8   `json:"logintype"` //登录类型(0:账号密码登录,1:手机号登录)
}

// MsgUserRegRes 客户端注册返回消息
type MsgUserRegRes struct {
	UserId   int32  `json:"userid"`   //ID
	AccName  string `json:"accname"`  //账号
	NickName string `json:"nickname"` //昵称
	FaceUrl  string `json:"faceurl"`  //头像
	Machine  string `json:"machine"`  //机器码
	AccType  int8   `json:"acctype"`  //账号类型(0:普通玩家,1:机器人,2:测试账号)
	RegType  int8   `json:"regtype"`  //注册方式(0:游客,1:账号注册,3:微信,...)
	Token    string `json:"token"`    //令牌(所有http请求都需要带令牌,配合UserId验证)
}

// ----------🠋🠋🠋🠋🠋🠋客户端/服务端通信消息Lobby🠋🠋🠋🠋🠋🠋----------

// MsgUpdFaceReq 修改头像消息
type MsgUpdFaceReq struct {
	UserId  int64  //用户ID
	FaceUrl string //头像
}

// MsgUpdFaceRes 修改头像结果消息
type MsgUpdFaceRes struct {
	Code    int32  //错误码
	FaceUrl string //头像
}
