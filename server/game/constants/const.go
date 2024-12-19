package constants

// 服务器名称(Route:main)
const (
	ServerNameGw    = "gw"    //网关
	ServerNameLobby = "lobby" //大厅
	ServerNameDb    = "db"    //db
	ServerNameWeb   = "web"   //web
)

// 功能模块(Route:module)
const (
	HandlerModuleGw    = "gw"    //网关服，
	HandlerModuleLobby = "lobby" //大厅服，
	HandlerModuleDb    = "db"    //DB服，
	HandlerModuleWeb   = "web"   //web
)

const (
	HandlerSubMainMsg = "mainmsg" //主消息
)

// 模块名称
const (
	ModuleMysqlStorage = "mysqlStorage" //mysql模块名称
	ModuleRedisStorage = "redisStorage" //redis模块名称
)

// ScoreRatio 金币比例
const ScoreRatio = 10000 //1元 = 10000积分

// 登录/注册固定参数
const (
	RegisterMachineMaxNum = 5             //机器码最大注册数量
	AesPasswordKey        = "Key-123^456" //Aes加密机密key
)

// 语言
const (
	LanguageCodeCn = "cn" //中文
	LanguageCodeEn = "en" //英文
)

// 账号类型
const (
	UserAccountTypeNone   = 0 //正常玩家
	UserAccountTypeRobot  = 1 //机器人
	UserAccountTypeTest   = 2 //测试账号
	UserAccountTypeAdmin  = 3 //管理员账号
	UserAccountTypeKiller = 4 //杀手账号
)
