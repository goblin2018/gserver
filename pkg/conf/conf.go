package conf

import "time"

var (
	Server server
	DB     db
	Redis  redis
)

type server struct {
	Port         int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type db struct {
	Host            string
	Port            int
	User            string
	Password        string
	Name            string
	MaxOpenConns    int
	MaxIdleConns    int
	ConnMaxLifetime time.Duration
}

type redis struct {
	Host string
	Port int
}

func init() {
	Server = server{
		Port:         9000,
		ReadTimeout:  time.Minute,
		WriteTimeout: time.Minute,
	}
	DB = db{
		Host:            "localhost",
		Port:            5432,
		User:            "postgres",
		Password:        "postgres",
		Name:            "goblin",
		MaxOpenConns:    400,
		MaxIdleConns:    40,
		ConnMaxLifetime: time.Hour,
	}

	Redis = redis{
		Host: "localhost",
		Port: 6379,
	}
}
