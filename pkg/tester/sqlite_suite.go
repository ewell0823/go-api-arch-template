// SQLiteでのテストで前後に実行する処理のコード
package tester

import (
	"go-api-arch-mvc-template/app/models"
	"go-api-arch-mvc-template/configs"
	"os"

	"github.com/stretchr/testify/suite"
)

type DBSQLiteSuite struct {
	suite.Suite
}

// テスト前に自動で実行するメソッド
// ここではnilチェックを行う
func (suite *DBSQLiteSuite) SetupSuite() {
	configs.Config.DBName = "unittest.sqlite"
	err := models.SetDatabase(models.InstanceSqlLite)
	suite.Assert().Nil(err)

	for _, model := range models.GetModels() {
		err := models.DB.AutoMigrate(model) // モデルに応じたテーブル作成
		suite.Assert().Nil(err)
	}
}

// テスト後に実行するメソッド
// dbファイルの削除、エラーが起きないことのチェック
func (suite *DBSQLiteSuite) TearDownSuite() {
	err := os.Remove(configs.Config.DBName)
	suite.Assert().Nil(err)
}
