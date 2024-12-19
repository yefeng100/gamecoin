package gw

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/topfreegames/pitaya/v2"
	"github.com/topfreegames/pitaya/v2/logger"
	"io"
	"net/http"
	"project/constants"
	"project/mods/ipinfo"
	"project/pb"
	"project/structs"
	"strings"
	"time"
)

// WebEnter web消息入口
func WebEnter(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	w.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	w.Header().Add("Access-Control-Allow-Credentials", "true")       // 允许发送 Cookie
	w.Header().Add("Content-Type", "application/json;charset=utf-8") // 响应类型为 JSON
	rpcHttp(w, r)
}

func rpcHttp(w http.ResponseWriter, r *http.Request) {
	bufBy, _ := io.ReadAll(r.Body)
	if strings.EqualFold(string(bufBy), "") {
		//在您的特定情况下,它的要点是Content-Type: application/json您的代码添加的请求标头是触发浏览器执行该预检OPTIONS请求的内容.
		return
	}
	var m map[string]interface{}
	errBy := json.Unmarshal(bufBy, &m)
	if errBy != nil {
		logger.Log.Errorf("rpcHttp json.Unmarshal err. buf:%v, err:%v", string(bufBy), errBy)
		return
	}
	language := r.Header.Get("language")
	if language == "" {
		m["language"] = "cn"
	}
	m["language"] = language
	m["ip"] = ipinfo.GetClientIP(r)
	m["jwt_token"] = r.Header.Get("jwt_token")
	data, _ := json.Marshal(m)
	buf := string(data)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	//请求web
	route := fmt.Sprintf("%s.%s.%s", constants.ServerNameWeb, constants.HandlerModuleWeb, constants.HandlerSubMainMsg)
	rpcReq := &pb.MsgRequest{}
	rpcReq.Msg = buf
	//rpc web
	rpcRes := &pb.MsgResponse{}
	err := pitaya.DefaultApp.RPC(ctx, route, rpcRes, rpcReq)
	if err != nil {
		logger.Log.Errorf("rpc failed err: %v", err.Error())
		rpcRes.Code = constants.ResCodeNotServerErr
	}
	msg := rpcRes.GetMsg()
	if msg == "" {
		msg = "{}"
	}
	resMsg := &structs.MsgHttp{}
	resMsg.Code = rpcRes.GetCode()
	resMsg.CodeMsg = constants.GetCodeMsg(language, rpcRes.GetCode())
	resMsg.Data = []byte(msg)
	resJson, _ := json.Marshal(resMsg)
	_, _ = w.Write(resJson)
	logger.Log.Debugf("MainMsg Body:%v, resCode:%v, resMsg:%v", buf, rpcRes.GetCode(), msg)
}
