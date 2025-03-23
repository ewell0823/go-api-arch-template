// 環境変数などの設定ファイル
package configs

import (
	"go-api-arch-mvc-template/pkg/logger"
	"os"
	"strconv"

	"go.uber.org/zap"
)

// 環境変数の値の取得と存在確認
func GetEnvDefault(key, defVal string) string {
	val, err := os.LookupEnv(key)
	if !err {
		return defVal
	}
	return val
}

// 環境変数を保持する構造体
type ConfigList struct {
	Env                 string
	DBHost              string
	DBPort              int
	DBDriver            string
	DBName              string
	DBUser              string
	DBPassword          string
	APICorsAllowOrigins []string
}

func (c *ConfigList) IsDevelopment() bool {
	return c.Env == "development"
}

var Config ConfigList

// 環境変数の設定を読み込んで、Configに代入する
func LoadEnv() error {
	// ポートの値を文字列->整数にまず変換
	DBPort, err := strconv.Atoi(GetEnvDefault("MYSQL_PORT", "3306"))
	if err != nil {
		return err
	}

	Config = ConfigList{
		Env:                 GetEnvDefault("APP_ENV", "development"),
		DBDriver:            GetEnvDefault("DB_DRIVER", "mysql"),
		DBHost:              GetEnvDefault("DB_HOST", "localhost"),
		DBPort:              DBPort,
		DBUser:              GetEnvDefault("DB_USER", "app"),
		DBPassword:          GetEnvDefault("DB_PASSWORD", "password"),
		DBName:              GetEnvDefault("DB_NAME", "api_database"),
		APICorsAllowOrigins: []string{"http://localhost:8001"},
	}
	return nil
}

func init() {
	if err := LoadEnv(); err != nil {
		logger.Error("Failed to load env:", zap.Error(err))
		panic(err)
	}
}
