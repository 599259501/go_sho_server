package logging

import (
	log "github.com/sirupsen/logrus"
	"path"
	"utils"
	"github.com/lestrrat/go-file-rotatelogs"
	"time"
	"os"
	"github.com/rifflock/lfshook"
	"strings"
)

func InitLogger(){
	logDir := utils.GetEnv("LOG_DIR", GetServerDir())
	logFileName := utils.GetEnv("LOG_FILE_NAME", "server_log.log")
	logLevel := utils.GetEnv("LOG_LEVEL", "info")
	baseLogPath := path.Join(logDir, logFileName)

	// logDir 初始化
	MakeDir(logDir)

	writer, err := rotatelogs.New(
		baseLogPath+".%Y%m%d%H%M",
		rotatelogs.WithLinkName(baseLogPath),      // 生成软链，指向最新日志文件
		rotatelogs.WithMaxAge(7*24*time.Hour),     // 文件最大保存时间
		rotatelogs.WithRotationTime(24*time.Hour), // 日志切割时间间隔
	)

	if err != nil {
		log.Errorf("config local file system logger error. %v", err)
	}

	switch logLevel {
	case "debug":
		log.SetOutput(os.Stdout)
		log.SetLevel(log.DebugLevel)
	}

	lfHook := lfshook.NewHook(lfshook.WriterMap{
		log.DebugLevel: os.Stdout, // 为不同级别设置不同的输出目的
		log.InfoLevel:  writer,
		log.WarnLevel:  writer,
		log.ErrorLevel: writer,
		log.FatalLevel: writer,
		log.PanicLevel: writer,
	},&log.JSONFormatter{})

	log.AddHook(lfHook)
}

func GetServerDir()string{
	currentDir, _ := os.Getwd()
	tmpDir := strings.Replace(currentDir, "\\", "/", -1)
	logDir := path.Join(tmpDir, "logs")

	return logDir
}

func MakeDir(dir string){
	_, err := os.Stat(dir)
	if os.IsNotExist(err){
		os.MkdirAll(dir, 666)
	}
}