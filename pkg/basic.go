package root

// Basic 基础信息
type Basic struct {
	AppID          string `json:"appid"`
	Secret         string `json:"secret"`
	TokenValidTime int64  `json:"tokenvalidtime"`
	TimeStamp      int64  `json:"timestamp"`
	Token          string `json:"token"`
}

// Basicer 基础信息接口
type Basicer interface {
	GetBasic() (Basic, error)
}
