package structs

import "encoding/json"

type MsgHttp struct {
	Code    int32           `json:"code"`
	CodeMsg string          `json:"code_msg"`
	Data    json.RawMessage `json:"data"`
}

//----------🠋🠋🠋🠋🠋🠋Http消息🠋🠋🠋🠋🠋🠋----------

// HttpHandMsg 消息头
type HttpHandMsg struct {
	MsgId    int32  `json:"msg_id"`    //消息ID
	Language string `json:"language"`  //语言
	JwtToken string `json:"jwt_token"` //Token
}

//--------------------------

// HttpNonceReq 获取nonce
// 注册：先获取nonce，返回nonce后再通过nonce给注册密码加密，再发注册消息（加密方式Ecb, 内容: ecb密钥_nonce 如: "Key-123^456_12345"）
// 登录：先获取nonce，返回nonce后再通过nonce给登录密码加密，再发登录消息（加密方式Ecb, 内容: ecb密钥_nonce 如: "Key-123^456_12345"）
type HttpNonceReq struct {
	AccName string `json:"acc_name"` //用户名
}

// HttpNonceRes 返回nonce
type HttpNonceRes struct {
	Nonce string `json:"nonce"` //nonce, 用来给密码加密
}

// HttpRegisterReq 注册
type HttpRegisterReq struct {
	AccName  string `json:"acc_name"` //用户名
	Password string `json:"password"` //密码
	Nickname string `json:"nickname"` //昵称
	Machine  string `json:"machine"`  //机器码
}

// HttpLoginReq 登录
type HttpLoginReq struct {
	AccName  string `json:"acc_name"` //用户名
	Password string `json:"password"` //密码
}

// HttpUserInfoReq 请求用户信息
type HttpUserInfoReq struct {
	AccName string `json:"acc_name"` //用户名
	UserId  int32  `json:"user_id"`  //用户ID
}

// HttpUserInfoRes 登录
type HttpUserInfoRes struct {
	JwtToken string `json:"jwt_token"` //token
	UserId   int32  `json:"user_id"`   //ID
	AccName  string `json:"acc_name"`  //账号
	NickName string `json:"nick_name"` //昵称
	FaceUrl  string `json:"face_url"`  //头像
	PhoneNum string `json:"phone_num"` //手机号
	Score    int64  `json:"score"`     //金币
	ScoreBox int64  `json:"score_box"` //保险金币
}

// HttpHomeRollRes 首页滚动数据返回
type HttpHomeRollRes struct {
	ID          int64  //主键
	Img         string //图片(有http为url地址,否则为本地图片名字)
	Url         string //跳转连接,点击图片打开连接(url为空,点击图片不做任何反应)
	UrlOpenType int8   //跳转连接打开方式(0默认跳转App页面,1浏览器打开URL)
}
