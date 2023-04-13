package main

import (
	"fmt"
	"time"

	"github.com/tidwall/gjson"
	"github.com/yomorun/yomo/core/frame"
)

type NoiseSensorData struct {
	Noise float32 `json:"noise"` // Noise level
	Time  int64   `json:"time"`  // Timestamp (ms)
	From  string  `json:"from"`  // Source IP
}

// Handler will handle the raw data
func Handler(data []byte) (frame.Tag, []byte) {
	fmt.Printf("sfn received %d bytes: %s\n", len(data), string(data))
	level := gjson.Get(string(data), "noise")
	timestamp := gjson.Get(string(data), "time")

	now := time.Now().UnixNano() / int64(time.Millisecond)
	fmt.Println(now)
	fmt.Println(timestamp.Int())

	fmt.Printf("data: %v, ⚡️: %dms\n", level.String(), now-timestamp.Int())

	return frame.Tag(0x34), nil
}

func DataTags() []frame.Tag {
	return []frame.Tag{0x33}
}
