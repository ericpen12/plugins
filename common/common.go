package common

import "encoding/json"

func MarshalBind(src, dsc interface{}) error {
	data, err := json.Marshal(src)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, dsc)
}
