package dbserver

import (
	"context"
	"encoding/json"
	"github.com/topfreegames/pitaya/v2"
	"github.com/topfreegames/pitaya/v2/component"
	"github.com/topfreegames/pitaya/v2/logger"
	"github.com/topfreegames/pitaya/v2/timer"
	"project/base/dao/daobase"
	"project/constants"
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

func (r *HandlerRemote) AfterInit() {
	logger.Log.Infof("Handler Now:%d", time.Now().Unix())
}

func (r *HandlerRemote) UpdFaceUrl(ctx context.Context, msg *pb.MsgRequest) (*pb.MsgResponse, error) {
	logger.Log.Infof("HandlerRemote UpdFaceUrl Now:%d", time.Now().Unix())
	res := &pb.MsgResponse{}
	res.Code = constants.ResCodeSuc

	tmpMsg := new(structs.MsgUpdFaceReq)
	err := json.Unmarshal([]byte(msg.GetMsg()), tmpMsg)
	if err != nil {
		res.Code = constants.ResCodeSysErr
		return res, nil
	}
	//修改头像
	daobase.UpdUserAccByFaceUrl(tmpMsg.UserId, tmpMsg.FaceUrl)
	return res, nil
}
