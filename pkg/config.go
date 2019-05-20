package root

// MssqlConfig mssql数据库连接参数
type MssqlConfig struct {
	Dialects string `json:"dialects"`
	Parm     string `json:"parm"`
}

// HTTPConfig http服务设置
type HTTPConfig struct {
	BaseHost string `json:"baseHost"`
	BasePath string `json:"basePath"`
	Host     string `json:"host"`
}

// BasicConfig 基础信息设置
type BasicConfig struct {
	AppID          string `json:"appid"`
	Secret         string `json:"secret"`
	TokenValidTime int64  `json:"tokenvalidtime"`
	TimeStamp      int64  `json:"timestamp"`
	Token          string `json:"token"`
}

//PathConfig 路径设置
type PathConfig struct {
	ConfigPath string `json:"configPath"`
	LogPath    string `json:"logPath"`
}

// LogConfig 日志对象
type LogConfig struct {
	Dir      string `json:"dir"`
	FileName string `json:"filename"`
	Level    string `json:"level"`
}

// Config 设置
type Config struct {
	Mssql *MssqlConfig `json:"mssql"`
	HTTP  *HTTPConfig  `json:"http"`
	Path  *PathConfig  `json:"-"`
	Basic *BasicConfig `json:"basic"`
	Log   *LogConfig   `json:"log"`
}

// Configer 获取系统配置信息
type Configer interface {
	GetConfig() (*Config, error)
}
