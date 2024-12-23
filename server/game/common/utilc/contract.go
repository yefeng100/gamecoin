package utilc

import "math/big"

const CoinExchangeRate = 10000 //汇率，

func ScoreToCoin(val int64) *big.Int {
	secPrice := big.NewFloat(float64(val) / float64(CoinExchangeRate))
	secPrice.Mul(secPrice, big.NewFloat(1000000000000000000))
	f := new(big.Int)
	secPrice.Int(f)
	return f
}

func CoinToScore(val *big.Int) int64 {
	f := new(big.Float).SetInt(val)
	baseValue := big.NewFloat(1000000000000000000)
	amountValue := new(big.Float).Quo(f, baseValue)
	amountTem, _ := amountValue.Float64()
	amount := int64(amountTem * CoinExchangeRate)
	return amount
}
