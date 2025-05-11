package bootstrap

import "database/sql"

type Application struct {
	Config *Config
	DB     *sql.DB
}

func NewApp() *Application {
	return &Application{
		Config: NewConfigFromEnv(),
		DB:     NewSqlDatabase(),
	}
}

func (app *Application) CloseDatabase() {
	CloseDatabase(app.DB)
}
