package setting

type AllConfig struct {
	App struct {
		Version string
		Name    string
		Log     struct {
			Level string
			File  struct {
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
}
