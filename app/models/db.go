package models

import (
	"errors"
	"fmt"
	"go-api-arch-mvc-template/configs"

	"gorm.io/driver/mysql"  // MySQLドライバー追加
	"gorm.io/driver/sqlite" // SQLiteドライバー追加
	"gorm.io/gorm"
)

const (
	InstanceSqlLite int = iota
	InstanceMySQL
)

var (
	DB                            *gorm.DB
	errInvalidSQLDatabaseInstance = errors.New("invalid sql db instance")
)

// モデル一覧をdbから返す関数
// 構造体のポインタを返す
func GetModels() []interface{} {
	return []interface{}{&Album{}, &Category{}}
}

// 引数で指定されたdbのインスタンスに応じて、dbに接続するための*gorm.DB型の値を作成
func NewDatabaseSQLFactory(instance int) (db *gorm.DB, err error) {
	switch instance {
	case InstanceMySQL:
		dsn := fmt.Sprintf(
			"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True",
			configs.Config.DBUser,
			configs.Config.DBPassword,
			configs.Config.DBHost,
			configs.Config.DBPort,
			configs.Config.DBName)
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{}) // mysql接続
	case InstanceSqlLite:
		db, err = gorm.Open(sqlite.Open(configs.Config.DBName), &gorm.Config{}) // sqlite接続
	default:
		return nil, errInvalidSQLDatabaseInstance
	}
	return db, err
}

func SetDatabase(instance int) (err error) {
	db, err := NewDatabaseSQLFactory(instance)
	if err != nil {
		return err
	}
	DB = db
	return err
}
