package env

import "flag"

const (
	Production  = "production"
	Test        = "test"
	Development = "development"
)

var env string

func Init() {
	flag.StringVar(&env, "e", Production, "运行环境， 默认 production")
}