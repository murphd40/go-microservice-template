package utils

import "encoding/json"

func Convert[S, T any](source S, target *T) {
	bs, _ := json.Marshal(source)
	json.Unmarshal(bs, target)
}