package service

import (
	"context"
	"encoding/json"
	"fmt"
	"livewallpaper/db"
	"livewallpaper/model"
	"log"
	"time"
)

var ctx = context.Background()

func VideoList(page int, pageSize int) ([]model.Video, error) {
	var videoList []model.Video
	key := fmt.Sprintf("VideoList_%d_%d", page, pageSize)
	val, err := db.Redis.Get(ctx, key).Result()
	if err != nil {
		db.Mysql.
			Preload("Category").
			Preload("Tags").
			Offset(pageSize * (page - 1)).
			Limit(pageSize).
			Order("updated_at DESC").
			Find(&videoList)

		val, err := json.Marshal(videoList)
		if err == nil {
			db.Redis.Set(ctx, key, val, time.Minute*5)
		}

		log.Printf("VideoList Mysql %d %d %d", page, pageSize, len(videoList))
	} else {
		err = json.Unmarshal([]byte(val), &videoList)
		if err != nil {
			return nil, err
		}
		log.Printf("VideoList Redis %d %d %d", page, pageSize, len(videoList))
	}
	return videoList, nil
}
