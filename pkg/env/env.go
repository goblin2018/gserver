package env

import "flag"

const (
	Production  = "production"
	Test        = "test"
	Development = "dev"
)

var ENV string

func Init() {
	flag.StringVar(&ENV, "e", Production, "运行环境， 默认 production")
	flag.Parse()
}
