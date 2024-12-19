我的笔记

创建组 / 创建自动超时删除组
    GroupCreate / GroupCreateWithTTL
    比如聊天群 / 自动过期聊天群
    里面可以添加成员, 
    GroupRenewTTL刷新并延长时间 

Message types 消息类型
    const (
        Request  Type = 0x00
        Notify   Type = 0x01
        Response Type = 0x02
        Push     Type = 0x03
    )
    var types = map[Type]string{
        Request:  "Request",
        Notify:   "Notify",
        Response: "Response",
        Push:     "Push",
    }