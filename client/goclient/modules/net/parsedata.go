package net

import "encoding/json"

var prettyJSON = false

func parseData(data []byte) []byte {
	prettyJSON = false
	if prettyJSON {
		var m interface{}
		_ = json.Unmarshal(data, &m)
		data, _ = json.MarshalIndent(m, "", "\t")
	}
	return data
}
