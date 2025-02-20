package application

type Application struct {
	Config *Config
}

func New() *Application {
	return &Application{
		Config: NewConfigFromEnv(),
	}
}
