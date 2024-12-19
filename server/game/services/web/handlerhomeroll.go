package web

import (
	"context"
	"encoding/json"
	"project/base/dao/daobase"
	"project/constants"
	"project/pb"
	"project/structs"
)

// 获取首页滚动数据
func handlerHomeRollList(ctx context.Context, userId int32, msg *pb.MsgRequest) (int32, string) {
	idxRollDataList := daobase.GetHomeRollList()
	resList := make([]*structs.HttpHomeRollRes, 0)
	for _, item := range idxRollDataList {
		data := &structs.HttpHomeRollRes{}
		data.ID = item.ID
		data.Img = item.Img
		data.Url = item.Url
		data.UrlOpenType = item.UrlOpenType
		resList = append(resList, data)
	}
	buf, _ := json.Marshal(resList)
	return constants.ResCodeSuc, string(buf)
}
