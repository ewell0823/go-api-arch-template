package models

import (
	"encoding/json"
	"go-api-arch-mvc-template/api"
	"go-api-arch-mvc-template/pkg"
	"time"
)

type Album struct {
	ID          int
	Title       string
	ReleaseDate time.Time
	CategoryID  int
	Category    *Category // 関連エンティティはポインタとして定義するのが一般的らしい
}

// アルバムリリースからの経過年数を表すメソッド
func (a *Album) Anniversary(clock pkg.Clock) int {
	now := clock.Now()
	years := now.Year() - a.ReleaseDate.Year()
	releaseDay := pkg.GetAdjustedReleaseDay(a.ReleaseDate, now)

	if now.YearDay() < releaseDay {
		years -= 1
	}
	return years
}

// 構造体をJSONに変換する処理
// api.AlbumResponse型の構造体に詰め替えて、json.Marshal関数でJSONに変換して返却している
func (a *Album) MarshalJSON() ([]byte, error) {
	return json.Marshal(&api.AlbumResponse{
		Id:          a.ID,
		Title:       a.Title,
		Anniversary: a.Anniversary(pkg.RealClock{}), // アルバムリリースからの経過年数
		ReleaseDate: &api.ReleaseDate{Time: a.ReleaseDate},
		Category: api.Category{
			Id:   &a.Category.ID,
			Name: api.CategoryName(a.Category.Name),
		},
	})
}

func (a *Album) AnniVersary(clock pkg.RealClock) {
	panic("unimplemented")
}

// アルバム作成メソッド
func CreateAlbum(title string, releaseDate time.Time, categoryName string) (*Album, error) {
	category, err := GetOrCreateCategory(categoryName)
	if err != nil {
		return nil, err
	}

	album := &Album{
		ReleaseDate: releaseDate,
		Title:       title,
		Category:    category,
		CategoryID:  category.ID,
	}
	if err := DB.Create(album).Error; err != nil {
		return nil, err
	}
	return album, err
}

// アルバム情報取得メソッド
func GetAlbum(ID int) (*Album, error) {
	var album = Album{}
	// Preloadでカテゴリー情報取得
	// Firstで検索した最初のレコード取得
	if err := DB.Preload("Category").First(&album, ID).Error; err != nil {
		return nil, err
	}
	return &album, nil
}

// アルバム保存メソッド
func (a *Album) Save() error {
	category, err := GetOrCreateCategory(a.Category.Name)
	if err != nil {
		return err
	}
	a.Category = category
	a.CategoryID = category.ID

	if err := DB.Save(&a).Error; err != nil {
		return err
	}
	return nil
}

// アルバム削除
func (a *Album) Delete() error {
	if err := DB.Where("id = ?", &a.ID).Delete(&a).Error; err != nil {
		return err
	}
	return nil
}
