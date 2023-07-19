package helper

import (
	"os"
	"path/filepath"
	"runtime"

	"github.com/joho/godotenv"
)

func GetEnv(key string) string {
	_, basePath, _, _ := runtime.Caller(0)
	rootDir := filepath.Join(filepath.Dir(basePath), "../")

	err := godotenv.Load(rootDir + "/.env")
	PanicIfError(err)

	return os.Getenv(key)
}
