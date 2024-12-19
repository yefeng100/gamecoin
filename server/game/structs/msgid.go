package structs

// 服务器消息主ID
const (
	ServerWebMainID   = 1000 //web
	ServerGwMainID    = 2000 //网关
	ServerDbMainID    = 3000 //db
	ServerLobbyMainID = 4000 //大厅
	ServerPdkMainID   = 5000 //跑得快
)

// 服务器消息子ID
const (
	GwMsgIdTest = ServerGwMainID + iota //网关
)

// Web服务器消息子ID
const (
	WebMsgIdNonce    = ServerWebMainID + iota //nonce
	WebMsgIdRegister = ServerWebMainID + iota //注册
	WebMsgIdLogin    = ServerWebMainID + iota //登录
	WebMsgIdUserInfo = ServerWebMainID + iota //用户信息
	WebMsgIdHomeRoll = ServerWebMainID + iota //获取主页滚动列表数据
)
