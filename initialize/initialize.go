package initialize

type initializer func() (err error)

func Initialize() (err error) {
	initializers := []initializer{
		InitConfig,
		InitDao,
		InitErrorsMap,
		InitSessionMiddleware,
	}

	for _, initializer := range initializers {
		if err = initializer(); err != nil {
			return
		}
	}

	return
}
