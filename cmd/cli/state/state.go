package state

import "encoding/json"

type State struct {
	// RootServer - root server configurations
	RootServer RootServer `json:"root_server"`
}

// ToJson - converts the state to a JSON string.
func (s State) ToJson() string {
	res, err := json.Marshal(s)
	if err != nil {
		panic(err)
	}
	return string(res)
}
