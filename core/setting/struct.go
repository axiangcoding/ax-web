package setting

type AllConfig struct {
	App    App
	Server Server
}

type App struct {
	Version string
	Name    string
	Log     Log
	Token   Token
	Swagger Swagger
}

type Log struct {
	Level   string
	FileLog FileLog
}

type FileLog struct {
	Enable bool
	Path   string
}

type Token struct {
	Secret         string
	ExpireDuration string
}

type Swagger struct {
	Enable bool
}

type Server struct {
	RunMode string `json:"run_mode"`
	Port    string
}


