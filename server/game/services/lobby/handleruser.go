package lobby

import (
	"context"
	"encoding/json"
	"github.com/topfreegames/pitaya/v2"
	"github.com/topfreegames/pitaya/v2/component"
	"github.com/topfreegames/pitaya/v2/logger"
	"github.com/topfreegames/pitaya/v2/timer"
	"project/constants"
	"project/pb"
	"project/structs"
	"time"
)

type HandlerUser struct {
	component.Base
	timer *timer.Timer
	app   pitaya.Pitaya
}

func NewHandlerUser(app pitaya.Pitaya) *HandlerUser {
	p := &HandlerUser{
		app: app}
	return p
}

func (r *HandlerUser) AfterInit() {
	logger.Log.Infof("Handler Now:%d", time.Now().Unix())
}

// UpdFaceUrl 修改头像
func (r *HandlerUser) UpdFaceUrl(ctx context.Context, req *structs.MsgUpdFaceReq) {
	logger.Log.Infof("UpdFaceUrl req:%v", req)
	s := r.app.GetSessionFromCtx(ctx)
	//返回客户端
	res := new(structs.MsgUpdFaceRes)
	res.Code = constants.ResCodeSuc
	res.FaceUrl = req.FaceUrl

	//请求db
	route := "db.dbremote.updfaceurl" // 修改头像
	rpcRes := &pb.MsgResponse{}
	buf, _ := json.Marshal(req)
	rpcReq := &pb.MsgRequest{}
	rpcReq.Msg = string(buf)
	//rpc db
	err := r.app.RPC(ctx, route, rpcRes, rpcReq)
	if err != nil {
		logger.Log.Errorf("UpdFaceUrl failed to enqueue rpc: %v", err.Error())
		res.Code = constants.ResCodeSysErr
	} else {
		res.Code = rpcRes.GetCode()
	}

	_ = s.Push("resUpdFaceUrl", res)
}
