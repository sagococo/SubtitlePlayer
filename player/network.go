package main

import (
	"fmt"
	"io"
	"net/http"
)

func getBilibiliVideo(videoID string) ([]byte, error) {
	videoURL := fmt.Sprintf("https://bili.zhouql.vip/meta/%v", videoID)
	client := &http.Client{}
	req, err := http.NewRequest("GET", videoURL, nil)
	if err != nil {
		fmt.Println("创建请求失败:", err)
		return nil, err
	}
	// 发送请求并获取响应
	rsp, err := client.Do(req)
	if err != nil {
		fmt.Println("发送请求失败:", err)
		return nil, err
	}
	defer rsp.Body.Close()
	return io.ReadAll(rsp.Body)
}
