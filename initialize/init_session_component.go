package initialize

var (
	DefaultRedisSecret = []byte("basic_app")
)

func InitSessionMiddleware() (err error) {
	// var config = conf.Conf.Redis

	// store, err := sessions.NewStore(10, "tcp", config.Address, config.Password, DefaultRedisSecret)
	// if err != nil {
	// 	return
	// }

	// middleware.Session = sessions.Sessions("basic_app", store)
	return
}
