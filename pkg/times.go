package pkg

import "time"

// 時間を取得するためのインターフェース
// 時間に関連する操作を抽象化してる
type Clock interface {
	// Now は現在時刻を返すメソッド
	Now() time.Time
}

// RealClock は実際のシステム時計を使用するClockインターフェースの実装
// フィールドを持たない空の構造体として定義されている
type RealClock struct{}

// NowメソッドはRealClock型に対して実装されており、
// 標準ライブラリのtime.Now()を呼び出して現在の実時間を返す
func (RealClock) Now() time.Time {
	return time.Now()
}

func isLeap(date time.Time) bool {
	year := date.Year()
	if year%400 == 0 {
		return true
	} else if year%100 == 0 {
		return false
	} else if year%4 == 0 {
		return true
	}
	return false
}

// うるう年考慮したアルバムのリリース日の年内の経過日数を調整するメソッド
func GetAdjustedReleaseDay(releaseDate time.Time, now time.Time) int {
	releaseDay := releaseDate.YearDay() // リリースから年内で経ってる経過日数
	currentDay := now.YearDay()         // 現在年内で経ってる経過日数

	// リリース日がうるう年
	// 現在がうるう年じゃない
	// リリース日が年内で60日以上経ってる
	if isLeap(releaseDate) && !isLeap(now) && releaseDay >= 60 {
		return releaseDay - 1
	}
	// リリース日がうるう年
	// 現在がうるう年じゃない
	// リリース日が年内で60日以上経ってる
	if isLeap(now) && isLeap(releaseDate) && currentDay >= 60 {
		return releaseDay + 1
	}
	return releaseDay
}
