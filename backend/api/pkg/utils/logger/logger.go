package logger

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/sirupsen/logrus"
)

var Log = logrus.New()

func init() {
	logDir := "logs"
	if err := os.MkdirAll(logDir, os.ModePerm); err != nil {
		log.Fatalf("logs folder create failed: %v", err)
	}
	logFileName := fmt.Sprintf("log-%s.log", time.Now().Format("02.01.2006-15-04-05"))
	logFilePath := filepath.Join(logDir, logFileName)
	file, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("log file open failed: %v", err)
	}
	multiWriter := io.MultiWriter(file, os.Stdout)
	Log.SetOutput(multiWriter)

	Log.SetLevel(logrus.InfoLevel)
	Log.SetFormatter(&logrus.TextFormatter{
		EnvironmentOverrideColors: true,
		ForceColors:               true,
		FullTimestamp:             true,
		TimestampFormat:           "02.01.2006-15-04-05",
	})
}
