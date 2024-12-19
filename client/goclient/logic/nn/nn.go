package nn

var nn *LogicNN

type LogicNN struct {
}

func NewLogicNN() *LogicNN {
	if nn == nil {
		nn = &LogicNN{}
	}
	return nn
}

func (t *LogicNN) HandlerData(data []byte) {

}
