package setting

type Config struct {
	Postgresql PostgresqlSettings `mapstructure:"postgresql"`
	Security   SecuritySettings   `mapstructure:"security"`
	Logger     LoggerSettings     `mapstructure:"logger"`
	Redis      RedisSettings      `mapstructure:"redis"`
	Server     ServerSettings     `mapstructure:"server"`
}

type ServerSettings struct {
	Port int    `mapstructure:"port"`
	Mode string `mapstructure:"mode"`
}

type RedisSettings struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Db       int    `mapstructure:"db"`
	Password string `mapstructure:"password"`
	PoolSize int    `mapstructure:"pool_size"`
}

type PostgresqlSettings struct {
	Host            string `mapstructure:"host"`
	Username        string `mapstructure:"username"`
	Password        string `mapstructure:"password"`
	Port            int    `mapstructure:"port"`
	Dbname          string `mapstructure:"dbname"`
	MaxIdleConns    int    `mapstructure:"maxIdleConns"`
	MaxOpenConns    int    `mapstructure:"maxOpenConns"`
	ConnMaxLifeTime int    `mapstructure:"connMaxLifeTime"`
}

type SecuritySettings struct {
	Jwt struct {
		Secret string `mapstructure:"secret"`
	} `mapstructure:"jwt"`
}

type LoggerSettings struct {
	File_name   string `mapstructure:"file_name"`
	Max_size    int    `mapstructure:"max_size"`
	Max_backups int    `mapstructure:"max_backups"`
	Max_age     int    `mapstructure:"max_age"`
	Compress    bool   `mapstructure:"compress"`
	Loglevel    string `mapstructure:"loglevel"`
}
