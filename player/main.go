package main

import (
	"encoding/json"
	"os"
	"player/entity"
)

func main() {
	data, err := getBilibiliVideo("BV1A94y1f7su")
	if err != nil {
		println(err.Error())
		return
	}
	videoListResponse := &entity.VideoListResponse{}
	if err := json.Unmarshal(data, videoListResponse); err != nil {
		println(err.Error())
		return
	}
	for _,v := range videoListResponse.Data.Pages {
		println(v.)
	}
	os.WriteFile("file.txt", data, 0777)
}
