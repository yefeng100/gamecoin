package gw

import (
	"context"
	"github.com/topfreegames/pitaya/v2"
	"github.com/topfreegames/pitaya/v2/component"
	"github.com/topfreegames/pitaya/v2/logger"
	"github.com/topfreegames/pitaya/v2/timer"
	"time"
)

type TestMessage struct {
	Name    string `json:"name"`
	Content string `json:"content"`
}

// Handler struct
type Handler struct {
	component.Base
	app   pitaya.Pitaya
	timer *timer.Timer
}

// NewHandler ctor
func NewHandler(app pitaya.Pitaya) *Handler {
	return &Handler{app: app}
}

// AfterInit component lifetime callback
func (r *Handler) AfterInit() {
	r.timer = pitaya.NewTimer(time.Minute*30, func() {
		logger.Log.Debugf("定时器 Now:%d", time.Now().Unix())
	})
}

// TestMessage sync last message to all members
func (r *Handler) TestMessage(ctx context.Context, msg *TestMessage) {
	logger.Log.Infof("Message %v", msg)
}
