package env

import (
	"estrim/common"
	"os"
)

func Get(key string) common.Parser {
	return common.ParseString(os.Getenv(key))
}
