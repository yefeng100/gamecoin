package web

import (
	"context"
	"encoding/json"
	"github.com/topfreegames/pitaya/v2"
	"github.com/topfreegames/pitaya/v2/component"
	"github.com/topfreegames/pitaya/v2/logger"
	"github.com/topfreegames/pitaya/v2/timer"
	"project/constants"
	"project/mods/middleware"
	"project/pb"
	"project/structs"
	"time"
)

type HandlerRemote struct {
	component.Base
	timer *timer.Timer
	app   pitaya.Pitaya
}

func NewHandlerRemote(app pitaya.Pitaya) *HandlerRemote {
	p := &HandlerRemote{
		app: app}
	return p
}

func (p *HandlerRemote) AfterInit() {
	logger.Log.Infof("Handler Now:%d", time.Now().Unix())
}

// MainMsg web消息入口
func (p *HandlerRemote) MainMsg(ctx context.Context, req *pb.MsgRequest) (*pb.MsgResponse, error) {
	logger.Log.Infof("MainMsg http 消息入口 Msg:%v", req.GetMsg())
	res := &pb.MsgResponse{}
	res.Code = constants.ResCodeSuc
	res.Msg = ""
	//消息
	reqHand := new(structs.HttpHandMsg)
	err := json.Unmarshal([]byte(req.GetMsg()), reqHand)
	if err != nil {
		res.Code = constants.ResCodeSysErr
		logger.Log.Errorf("MainMsg Unmarshal err. err:%v", err)
		return res, nil
	}
	//检查jwtToken
	jwtCode, userId := verifySessionId(reqHand.MsgId, reqHand.JwtToken)
	if jwtCode != constants.ResCodeSuc {
		res.Code = jwtCode
		logger.Log.Errorf("MainMsg verifySessionId err. jwtCode:%v, reqHand:%v", jwtCode, reqHand)
		return res, nil
	}

	switch reqHand.MsgId {
	case structs.WebMsgIdNonce:
		//获取nonce
		res.Code, res.Msg = handlerNonce(ctx, req)
	case structs.WebMsgIdRegister:
		//注册
		res.Code, res.Msg = handlerRegister(ctx, req)
	case structs.WebMsgIdLogin:
		//登录
		res.Code, res.Msg = handlerLogin(ctx, req)
	case structs.WebMsgIdUserInfo:
		//用户信息
		res.Code, res.Msg = handlerUserInfo(ctx, req)
	case structs.WebMsgIdHomeRoll:
		//首页滚动数据
		res.Code, res.Msg = handlerHomeRollList(ctx, userId, req)
	default:
		res.Code = constants.ResCodeUndefined
		logger.Log.Errorf("MainMsg msg code err. code:%v, res:%v", res.Code, res)
	}
	return res, nil
}

// 验证sessionId
func verifySessionId(msgId int32, jwtToken string) (int32, int32) {
	if msgId == structs.WebMsgIdNonce ||
		msgId == structs.WebMsgIdRegister ||
		msgId == structs.WebMsgIdLogin {
		return constants.ResCodeSuc, 0
	}
	jwtCode, userId := middleware.JwtVerify(jwtToken)

	return jwtCode, userId
}
