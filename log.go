package common

import (
	"log"
	"os"
)

// 初始化go默认的日志
func InitGoLog(path string) (err error) {
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		return
	}

	log.SetFlags(log.Lshortfile | log.LstdFlags)
	log.SetOutput(f)
	//log.SetOutput(bufio.NewWriterSize(f, 1024 * 16)) // 设置带缓冲的writer，要等到缓冲满了才会写入到文件，收不到信号
	return
}
