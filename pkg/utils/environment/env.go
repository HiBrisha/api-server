package env

import (
	Logger "module/pkg/utils/logs"
	"os"
)

func Get(key string, log *Logger.SysLog) string {
	log.Write("Info", "Lấy giá trị %s thành công: %v", key, os.Getenv(key))
	return os.Getenv(key)
}
