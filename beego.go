package common

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/logs"
)

func convertLogLevel(level string) int {

	switch level {
	case "debug":
		return logs.LevelDebug
	case "warn":
		return logs.LevelWarn
	case "info":
		return logs.LevelInfo
	case "trace":
		return logs.LevelTrace
	}

	return logs.LevelDebug
}

func InitBeegoLogger(path, level string) (err error) {
	config := make(map[string]interface{})
	config["filename"] = path
	config["level"] = convertLogLevel(level)

	configStr, err := json.Marshal(config)
	if err != nil {
		fmt.Println("marshal failed, err:", err)
		return
	}

	logs.SetLogger(logs.AdapterFile, string(configStr))
	return
}
