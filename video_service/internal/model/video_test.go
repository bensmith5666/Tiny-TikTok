package model

import (
	"fmt"
	"testing"
	"time"
)

// 测试视频喜欢记录 + 1
func TestVideoModel_AddFavoriteCount(t *testing.T) {
	InitDb()
	GetVideoInstance().AddFavoriteCount(1949953677991936)
}

// 测试视频喜欢记录 - 1
func TestFavoriteModel_DeleteFavorite2(t *testing.T) {
	InitDb()
	GetVideoInstance().DeleteFavoriteCount(1949953677991936)
}

// 测试新增喜欢记录
func TestFavoriteModel_AddFavorite(t *testing.T) {
	InitDb()
	favorite := Favorite{
		UserId:  123,
		VideoId: 456,
	}
	GetFavoriteInstance().AddFavorite(&favorite)
}

// 测试软删除喜欢记录
func TestFavoriteModel_DeleteFavorite(t *testing.T) {
	InitDb()
	favorite := Favorite{
		UserId:  123,
		VideoId: 456,
	}
	GetFavoriteInstance().DeleteFavorite(&favorite)
}

// 测试创建评论
func TestCommentModel_CreateComment(t *testing.T) {
	InitDb()
	comment := Comment{
		UserId:  111,
		VideoId: 222,
		Content: "喜欢",
	}
	GetCommentInstance().CreateComment(&comment)
}

// 测试删除评论
func TestCommentModel_DeleteComment(t *testing.T) {
	InitDb()
	GetCommentInstance().DeleteComment(2289128100995072)
}

func TestTime(t *testing.T) {
	currentTime := time.Now()
	fmt.Println(currentTime)
	timeString := currentTime.Format("2006-01-02 15:04:05")
	fmt.Println("Formatted time:", timeString)
}

// 测试用户获赞数量
func TestVideoModel_GetFavoritedCount(t *testing.T) {
	InitDb()
	count, _ := GetVideoInstance().GetFavoritedCount(812575311663104)
	fmt.Println(count)
}

// 统计作品数量
func TestVideoModel_GetWorkCount(t *testing.T) {
	InitDb()
	count, _ := GetVideoInstance().GetWorkCount(812575311663104)
	fmt.Println(count)
}

// 统计喜欢数量
func TestFavorite_GetFavoriteCount(t *testing.T) {
	InitDb()
	count, _ := GetFavoriteInstance().GetFavoriteCount(812575311663104)
	fmt.Println(count)
}

// 找到喜欢的视频id
func TestFavoriteModel_FavoriteVideoList(t *testing.T) {
	InitDb()
	list, _ := GetFavoriteInstance().FavoriteVideoList(812575311663104)

	videoList, _ := GetVideoInstance().GetVideoList(list)
	fmt.Print(list)
	fmt.Print(videoList)
}
