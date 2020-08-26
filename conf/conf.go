package conf

var Conf Config

type Config struct {
	MySQL
	Redis
}

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

type Redis struct {
	Address  string
	Password string
	DB       int
}
