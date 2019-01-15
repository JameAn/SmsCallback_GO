package libs 

import (
	"net/http"
	"io/ioutil"
	"net/url"
	"fmt"
	log "github.com/sirupsen/logrus"
)

type SendData struct {
	JsonData []byte 
	RemoteUrl string
	LogHandler *LogGer
}

func (s *SendData) PostTo(){
	res, err := http.PostForm(s.RemoteUrl, url.Values{"data": {string(s.JsonData)}})
	if err != nil {
	    logMsg := fmt.Sprintf("Fatal error ", err.Error())
		s.LogHandler.WriteInfo(log.Fields{}, logMsg)
		return
	}
	defer res.Body.Close()
	
	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
	    logMsg := fmt.Sprintf("Fatal error ", err.Error())
		s.LogHandler.WriteInfo(log.Fields{}, logMsg)
		return
	}
	logMsg := fmt.Sprintf("SendSuccessful ", string(content))
    s.LogHandler.WriteInfo(log.Fields{}, logMsg)
}
