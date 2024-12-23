package erc20

import (
	"context"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/topfreegames/pitaya/v2"
	"github.com/topfreegames/pitaya/v2/logger"
	"github.com/topfreegames/pitaya/v2/timer"
	"go.uber.org/zap"
	"math/big"
	"project/base/cache/rediscache"
	"project/common/utilc"
	"strings"
	"time"
)

var tmErc20 *TimerErc20

type TimerErc20 struct {
	timer *timer.Timer
}

func NewTimerErc20() *TimerErc20 {
	if tmErc20 == nil {
		tmErc20 = &TimerErc20{}
		tmErc20.init()
	}
	return tmErc20
}

func (t *TimerErc20) init() {
	t.timer = pitaya.NewTimer(time.Second*5, t.EventApproval)
}

func (t *TimerErc20) EventApproval() {
	fromBlock := rediscache.RGetErcBlockApprove()
	if fromBlock < 0 {
		logger.Log.Error("EventApproval fromBlock err", zap.Int64("fromBlock", fromBlock))
		return
	}
	header, err := GetIns().Client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		logger.Log.Error("EventApproval HeaderByNumber err", zap.Error(err))
		return
	}
	currBlockId := header.Number.Int64()
	if fromBlock <= 0 {
		fromBlock = currBlockId
	}
	toBlockId := fromBlock + 5000
	if toBlockId > currBlockId-10 {
		toBlockId = currBlockId - 10
	}
	if fromBlock >= toBlockId {
		_ = rediscache.RSetErcBlockApprove(fromBlock)
		return
	}
	//查询
	cAddr := common.HexToAddress("0x337610d27c682E347C9cD60BD4b3b107C9d34dDd")
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(fromBlock),
		ToBlock:   big.NewInt(toBlockId),
		Addresses: []common.Address{
			cAddr,
		},
	}
	logs, err := c.Client.FilterLogs(context.Background(), query)
	if err != nil {
		logger.Log.Error("EventApproval FilterLogs err", zap.Error(err))
		return
	}
	setBlock := toBlockId
	for _, vLog := range logs {
		ev, err := GetIns().AN.ParseApproval(vLog)
		if err != nil {
			//logger.Log.Error("EventApproval ParseApproval err", zap.Error(err))
			continue
		}
		hash := ev.Raw.TxHash.String()      //hash
		block := ev.Raw.BlockNumber         //块
		userAddr := ev.Owner.String()       //玩家钱包地址
		conAddr := ev.Spender.String()      //合约钱包地址(被授权的合约)
		coinAddr := ev.Raw.Address.String() //币地址(usdt)
		if strings.ToLower(conAddr) == strings.ToLower("0x77198c2EC3f07951Da4372a28c3d9B0fceA5Eb8F") {
			logger.Log.Debugf("----block:%d, hash:%s, Value:%d, conAddr:%s, userAddr:%s, coinAddr:%s", block, hash, utilc.CoinToScore(ev.Value), conAddr, userAddr, coinAddr)

		}

		/*if strings.ToLower(ac) == strings.ToLower("0x47B8EcBA4fdaaF1D47cd60920B9834B333548E51") ||
			strings.ToLower(bc) == strings.ToLower("0x47B8EcBA4fdaaF1D47cd60920B9834B333548E51") ||
			strings.ToLower(cc) == strings.ToLower("0x47B8EcBA4fdaaF1D47cd60920B9834B333548E51") ||
			strings.ToLower(ac) == strings.ToLower("0x77198c2EC3f07951Da4372a28c3d9B0fceA5Eb8F") ||
			strings.ToLower(bc) == strings.ToLower("0x77198c2EC3f07951Da4372a28c3d9B0fceA5Eb8F") ||
			strings.ToLower(cc) == strings.ToLower("0x77198c2EC3f07951Da4372a28c3d9B0fceA5Eb8F") {
			println(ac, bc, cc)
		}*/
	}
	_ = rediscache.RSetErcBlockApprove(setBlock + 1)
}
