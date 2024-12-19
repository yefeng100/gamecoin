package constants

// 返回错误码
const (
	ResCodeSuc              int32 = 0    //返回成功
	ResCodeNotServerErr     int32 = 9    //未找到服务器
	ResCodeSysErr           int32 = 10   //系统错误
	ResCodeUndefined        int32 = 11   //消息未定义
	ResCodeBindErr          int32 = 12   //绑定网络错误,需要重新连接
	ResCodeJwtTokenErr      int32 = 13   //令牌错误
	ResCodeReqParamErr      int32 = 14   //参数错误
	ResCodeAccExistErr      int32 = 1000 //账号已存在
	ResCodeAccNotExistErr   int32 = 1001 //账号不存在
	ResCodeMachineOverLimit int32 = 1002 //机器码超过限制
	ResCodeUserPwdErr       int32 = 1003 //密码错误
	ResCodeCreateUserErr    int32 = 1004 //注册失败
	ResCodeNonceErr         int32 = 1005 //nonce错误
	ResCodeEcbErr           int32 = 1006 //解密失败
)
