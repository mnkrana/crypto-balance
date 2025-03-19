package domain

type Chain struct {
	Network string `json:"network"`
	RpcUrl  string `json:"rpc"`
	PKey    string `json:"key"`
}
