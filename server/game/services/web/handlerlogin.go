package web

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/topfreegames/pitaya/v2/logger"
	"gorm.io/gorm"
	"math/rand"
	"project/base/cache/rediscache"
	"project/base/dao/daobase"
	"project/base/entity"
	"project/constants"
	"project/mods/aes"
	"project/mods/middleware"
	"project/modules/mysqlstorage"
	"project/pb"
	"project/structs"
)

// 获取nonce，用户给密码加密
func handlerNonce(ctx context.Context, msg *pb.MsgRequest) (int32, string) {
	req := &structs.HttpNonceReq{}
	err := json.Unmarshal([]byte(msg.GetMsg()), req)
	if err != nil {
		logger.Log.Errorf("handlerNonce Unmarshal error. msg:%v, err:%v", msg.GetMsg(), err.Error())
		return constants.ResCodeReqParamErr, ""
	}
	if req.AccName == "" {
		logger.Log.Errorf("handlerNonce UserKey error. req:%v", req)
		return constants.ResCodeReqParamErr, ""
	}
	//随机nonce
	nonce := fmt.Sprintf("%d", rand.Int31n(90000)+10000)
	//保存redis
	rediscache.RSetNonce(ctx, req.AccName, nonce)

	res := &structs.HttpNonceRes{}
	res.Nonce = nonce
	buf, _ := json.Marshal(res)
	return constants.ResCodeSuc, string(buf)
}

func handlerRegister(ctx context.Context, msg *pb.MsgRequest) (int32, string) {
	req := &structs.HttpRegisterReq{}
	err := json.Unmarshal([]byte(msg.GetMsg()), req)
	if err != nil {
		logger.Log.Errorf("handlerRegister Unmarshal error. msg:%v, err:%v", msg.GetMsg(), err)
		return constants.ResCodeReqParamErr, ""
	}
	if req.AccName == "" {
		logger.Log.Errorf("handlerRegister Username error. req:%v", req)
		return constants.ResCodeReqParamErr, ""
	}
	//账号是否存在
	existAccName := daobase.GetUserAccByAccName(req.AccName)
	if existAccName != nil {
		logger.Log.Errorf("handlerRegister Username not exist. req:%v", req)
		return constants.ResCodeAccExistErr, ""
	}
	//机器码注册个数
	machineRegCount := daobase.GetUserAccCountByMachine(req.Machine)
	if machineRegCount > constants.RegisterMachineMaxNum {
		logger.Log.Errorf("handlerRegister machine over limit err. req:%v", req)
		return constants.ResCodeMachineOverLimit, ""
	}
	//获取nonce
	nonce := rediscache.RGetNonce(ctx, req.AccName)
	if nonce == "" {
		logger.Log.Errorf("handlerRegister RGetNonce err. req:%v", req)
		return constants.ResCodeNonceErr, ""
	}
	//解密
	ecbKey := fmt.Sprintf("%s_%s", constants.AesPasswordKey, nonce)
	pwd := aes.AesDecryptECB(req.Password, []byte(ecbKey))
	if string(pwd) == "" {
		logger.Log.Errorf("handlerRegister AesDecryptECB err. req:%v", req)
		return constants.ResCodeEcbErr, ""
	}
	//注册
	conn := mysqlstorage.GetAcc() //获取mysql实例
	if conn == nil {
		logger.Log.Errorf("handlerRegister mysqlstorage err. req:%v", req)
		return constants.ResCodeSysErr, ""
	}
	err = conn.WithContext(ctx).Transaction(func(tDB *gorm.DB) error {
		//获取用户ID
		userId, errT := daobase.GetNewUserId(tDB)
		if errT != nil {
			logger.Log.Errorf("handlerRegister GetNewUserId err. req:%v, err:%v", req, err)
			return errT
		}
		//账号
		userAcc := &entity.UserAccount{}
		userAcc.UserId = userId
		userAcc.AccName = req.AccName
		userAcc.Pwd = string(pwd)
		userAcc.NickName = req.Nickname
		userAcc.FaceUrl = fmt.Sprintf("%d", rand.Int31n(10)+1)
		userAcc.Machine = req.Machine
		userAcc.AccType = constants.UserAccountTypeNone
		errT = daobase.CreateUserAccount(tDB, userAcc)
		if errT != nil {
			logger.Log.Errorf("handlerRegister CreateUserAccount err. userAcc:%v, err:%v", userAcc, err)
			return errT
		}
		//用户金币
		userScore := &entity.UserScore{}
		userScore.UserId = userId
		errT = daobase.CreateUserScore(tDB, userScore)
		if errT != nil {
			logger.Log.Errorf("handlerRegister CreateUserScore err. userScore:%v, err:%v", userScore, err)
			return errT
		}
		return nil
	})
	if err != nil {
		logger.Log.Errorf("handlerRegister Transaction err. req:%v, err:%v", req, err)
		return constants.ResCodeCreateUserErr, ""
	}
	//登录
	reqUserInfo := &structs.HttpLoginReq{}
	reqUserInfo.AccName = req.AccName
	reqUserInfo.Password = req.Password
	data, _ := json.Marshal(reqUserInfo)

	pbMsg := &pb.MsgRequest{}
	pbMsg.Msg = string(data)
	return handlerLogin(ctx, pbMsg)
}

func handlerLogin(ctx context.Context, msg *pb.MsgRequest) (int32, string) {
	req := &structs.HttpLoginReq{}
	err := json.Unmarshal([]byte(msg.GetMsg()), req)
	if err != nil {
		logger.Log.Errorf("handlerLogin Unmarshal error. msg:%v, err:%v", msg.GetMsg(), err.Error())
		return constants.ResCodeReqParamErr, ""
	}
	if req.AccName == "" {
		logger.Log.Errorf("handlerLogin Username error. req:%v", req)
		return constants.ResCodeAccNotExistErr, ""
	}
	//账号信息
	userAcc := daobase.GetUserAccByAccName(req.AccName)
	if userAcc == nil {
		logger.Log.Errorf("handlerLogin Username not exist. req:%v", req)
		return constants.ResCodeAccNotExistErr, ""
	}
	//获取nonce
	nonce := rediscache.RGetNonce(ctx, req.AccName)
	if nonce == "" {
		logger.Log.Errorf("handlerRegister RGetNonce err. req:%v", req)
		return constants.ResCodeNonceErr, ""
	}
	//解密
	ecbKey := fmt.Sprintf("%s_%s", constants.AesPasswordKey, nonce)
	pwd := aes.AesDecryptECB(req.Password, []byte(ecbKey))
	if string(pwd) == "" {
		logger.Log.Errorf("handlerRegister AesDecryptECB err. req:%v", req)
		return constants.ResCodeEcbErr, ""
	}
	if userAcc.Pwd != string(pwd) {
		logger.Log.Errorf("handlerLogin Password err. req:%v", req)
		return constants.ResCodeUserPwdErr, ""
	}
	userScore := daobase.GetUserScoreById(userAcc.UserId)
	if userScore == nil {
		logger.Log.Errorf("handlerLogin GetUserScoreById err. req:%v", req)
		return constants.ResCodeAccNotExistErr, ""
	}
	// 生成新的jwt token
	jwtToken, err := middleware.GenerateToken(userAcc.UserId)
	if err != nil {
		logger.Log.Errorf("handlerLogin GenerateToken err. req:%v", req)
		return constants.ResCodeSysErr, ""
	}
	//返回数据
	res := &structs.HttpUserInfoRes{}
	res.JwtToken = jwtToken           //token
	res.UserId = userAcc.UserId       //ID
	res.AccName = userAcc.AccName     //账号
	res.NickName = userAcc.NickName   //昵称
	res.FaceUrl = userAcc.FaceUrl     //头像
	res.PhoneNum = userAcc.PhoneNum   //手机号
	res.Score = userScore.Score       //金币
	res.ScoreBox = userScore.ScoreBox //保险金币

	buf, _ := json.Marshal(res)
	return constants.ResCodeSuc, string(buf)
}

func handlerUserInfo(ctx context.Context, msg *pb.MsgRequest) (int32, string) {
	req := &structs.HttpUserInfoReq{}
	err := json.Unmarshal([]byte(msg.GetMsg()), req)
	if err != nil {
		logger.Log.Errorf("handlerLogin Unmarshal error. msg:%v, err:%v", msg.GetMsg(), err.Error())
		return constants.ResCodeReqParamErr, ""
	}
	//账号信息
	var userAcc *entity.UserAccount = nil
	if req.AccName != "" {
		userAcc = daobase.GetUserAccByAccName(req.AccName)
	} else if req.UserId > 0 {
		userAcc = daobase.GetUserAccById(req.UserId)
	} else {
		logger.Log.Errorf("handlerLogin req param error. msg:%v", msg.GetMsg())
		return constants.ResCodeReqParamErr, ""
	}
	if userAcc == nil {
		logger.Log.Errorf("handlerLogin Username not exist. req:%v", req)
		return constants.ResCodeAccNotExistErr, ""
	}
	//金币信息
	userScore := daobase.GetUserScoreById(userAcc.UserId)
	if userScore == nil {
		logger.Log.Errorf("handlerLogin GetUserScoreById err. req:%v", req)
		return constants.ResCodeAccNotExistErr, ""
	}

	//返回数据
	res := &structs.HttpUserInfoRes{}
	res.UserId = userAcc.UserId       //ID
	res.AccName = userAcc.AccName     //账号
	res.NickName = userAcc.NickName   //昵称
	res.FaceUrl = userAcc.FaceUrl     //头像
	res.PhoneNum = userAcc.PhoneNum   //手机号
	res.Score = userScore.Score       //金币
	res.ScoreBox = userScore.ScoreBox //保险金币

	buf, _ := json.Marshal(res)
	return constants.ResCodeSuc, string(buf)
}
