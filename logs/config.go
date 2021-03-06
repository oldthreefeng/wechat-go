package logs

import (
	"encoding/json"
	"runtime"
)

//二次开发logger
func CfgDbg() {
	config := logConfig{
		TimeFormat: "15:04:05",
		Console: &consoleLogger{
			LogLevel: LevelDebug,
			Colorful: runtime.GOOS != "windows",
		},
	}
	cfg, _ := json.Marshal(config)
	SetLogger(string(cfg))
	SetLogPath(true)
}

func CfgInfo() {
	config := logConfig{
		TimeFormat: "15:04:05",
		Console: &consoleLogger{
			LogLevel: LevelInformational,
			Colorful: runtime.GOOS != "windows",
		},
	}
	cfg, _ := json.Marshal(config)
	SetLogger(string(cfg))
	SetLogPath(true)
}
