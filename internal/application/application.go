package application

type Application struct {
	config *Config
}

func New() *Application {
	return &Application{
		config: NewConfigFromEnv(),
	}
}
