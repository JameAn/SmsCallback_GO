package libs

import (
	"os"
	"github.com/rifflock/lfshook"
	log "github.com/sirupsen/logrus"
	"fmt"
//	"log"
)

type LogMsg struct {
	Msg string
	Fields log.Fields
}
type LogGer struct {
	slog *log.Logger
	logPath string
}

var Nlog *LogGer

func init(){
	Nlog = new(LogGer) 
}

func NewSlog(logPath string) *LogGer{
	pathMap := lfshook.PathMap{
		log.InfoLevel: fmt.Sprintf("%s/info.log", logPath),
		log.ErrorLevel: fmt.Sprintf("%s/error.log", logPath),
	}
	Nlog.slog = log.New()
    //Nlog.slog.SetFormatter(&log.JSONFormatter{})
	//Nlog.slog.SetLevel(log.InfoLevel)

    file, err := os.OpenFile("/dev/null", os.O_CREATE|os.O_WRONLY, 0666)
    if err == nil {
        Nlog.slog.Out = file
    } else {
        log.Info("Failed to log to file, using default stderr")
    }

	Nlog.slog.Hooks.Add(lfshook.NewHook(
		pathMap,
		&log.JSONFormatter{},
	))

	return Nlog 
}

//func (nl *LogGer) TestLog() {
//	nl.slog.WithFields(log.Fields{
//		"animal": "walrus",
//		"size":   10,
//	}).Info("A group of walrus emerges from the ocean")
//}

func (nl *LogGer) WriteInfo(params log.Fields, msg string) {
	nl.slog.WithFields(params).Info(msg)
}




