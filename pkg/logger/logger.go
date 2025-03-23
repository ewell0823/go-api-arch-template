// ロギング設定
// Zapという高速で構造化されたロギングパッケージを使用
// zap.Loggerとzap.SuggerLoggerという型がある
// SuggerLoggerの方がLoggerより簡単にかける
// 参考記事 https://qiita.com/a-kym/items/1b9718da836f08ba8497

package logger

import (
	"os"

	"go.uber.org/zap"
)

var (
	ZapLogger        *zap.Logger        // zap.Logger型の変数宣言
	zapSugaredLogger *zap.SugaredLogger // zap.SuggerLogger型の変数宣言
)

// 上記変数の初期化をする
func init() {
	// 本番環境用の設定構築
	cfg := zap.NewProductionConfig()
	// 環境変数取得
	logFile := os.Getenv("APP_LOG_FILE")
	if logFile != "" {
		// ログ出力先
		cfg.OutputPaths = []string{"stderr", logFile}
	}

	// ロガー作成（失敗したらpanic）
	ZapLogger = zap.Must(cfg.Build())
	if os.Getenv("APP_ENV") == "development" {
		// 開発環境用のロガー作成
		ZapLogger = zap.Must(zap.NewDevelopment())
	}
	zapSugaredLogger = ZapLogger.Sugar()
}

// Sync使用することでメモリ内のバッファに残ってるログをファイルに出力
func Sync() {
	err := zapSugaredLogger.Sync()
	if err != nil {
		zap.Error(err)
	}
}

// 各ログレベルのログ出力

func Info(msg string, keysAndValues ...interface{}) {
	zapSugaredLogger.Infow(msg, keysAndValues...)
}

func Debug(msg string, keysAndValues ...interface{}) {
	zapSugaredLogger.Debugw(msg, keysAndValues...)
}

func Warn(msg string, keysAndValues ...interface{}) {
	zapSugaredLogger.Warnw(msg, keysAndValues...)
}

func Error(msg string, keysAndValues ...interface{}) {
	zapSugaredLogger.Errorw(msg, keysAndValues...)
}

func Fatal(msg string, keysAndValues ...interface{}) {
	zapSugaredLogger.Fatalw(msg, keysAndValues...)
}

func Panic(msg string, keysAndValues ...interface{}) {
	zapSugaredLogger.Panicw(msg, keysAndValues...)
}
