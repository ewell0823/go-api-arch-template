// テストコード
// go test ./config/...で実行できる
package configs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitEnv(t *testing.T) {
	err := LoadEnv()
	assert.Nil(t, err)
	assert.Equal(t, "development", Config.Env)
}
