package client

import "encoding/json"

func MustEncodeJSON(v interface{}) string {
	b, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		panic(err)
	}
	return string(b)
}
