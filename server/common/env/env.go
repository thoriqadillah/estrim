package env

import (
	"estrim/common"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	wd, _ := os.Getwd()
	path := fmt.Sprintf("%s/../.env", wd)
	godotenv.Load(path)
}

func Get(key string) common.Parser {
	return common.ParseString(os.Getenv(key))
}
