package conf

type (
	AllConfig struct {
		App
		Server
		Data
	}
	App struct {
		Version string
		Name    string
		Log     struct {
			Level   string
			FileLog struct {
				Enable bool
				Path   string
			}
		}
		Token struct {
			Secret         string
			ExpireDuration string `mapstructure:"expire_duration"`
		}
		Swagger struct {
			Enable bool
		}
	}
	Server struct {
		RunMode string `mapstructure:"run_mode"`
		Port    string
	}

	Data struct {
		Database
		Redis
	}
	Database struct {
		Driver string
		Source string
	}

	Redis struct {
		Network      string
		Addr         string
		Password     string
		Db           int32
		DialTimeout  int32
		ReadTimeout  int32
		WriteTimeout int32
	}
)
