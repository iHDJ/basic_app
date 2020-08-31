package conf

var Conf Config

type Config struct {
	MySQL
	Redis
	Session
}

//mysql数据库配置
type MySQL struct {
	Username     string
	Password     string
	Address      string
	Database     string
	Config       string
	MaxIdleConns int
	MaxOpenConns int
	LogMode      bool
}

//redis服务器配置
type Redis struct {
	Address  string
	Password string
	DB       int
}

//存储用户session的redis服务器配置
type Session struct {
	Address  string
	Password string
	DB       int
}
