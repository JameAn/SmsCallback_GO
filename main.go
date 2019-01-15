package main

import (
	"log"
	"github.com/gin-gonic/gin"
	"SmsCallback_Go/libs"
	"SmsCallback_Go/conf"
	"net/http"
	"encoding/json"
	"os"
	"os/signal"
	"time"
	"context"
	"flag"
	"syscall"
)
var (
	srv *http.Server
    r *gin.Engine
	cf *conf.Conf
	slog *libs.LogGer 
)

func init () {
	gin.SetMode(gin.ReleaseMode)
    r = gin.New()
	r.Use(gin.Logger(), gin.Recovery())
	//r = gin.Default()
	cf = conf.New()
	addr := flag.String("addr", cf.GetAddr(), "server addr")
	flag.Parse()
	srv = &http.Server{
		Addr: *addr,
		Handler: r,
	}

	slog = libs.NewSlog(cf.ServerConf.Conf.LogPath)

	libs.Dhst.LogHandler = slog
}

func main() {
	r.POST("/sms/dh_callback", func(c *gin.Context) {
		deliver := c.PostForm("deliver")
		if deliver == "" {
            c.JSON(http.StatusOK, gin.H{
		    	"status": "params empty",
		    })
			return 
		}

		var callData libs.Deliver
		err := json.Unmarshal([]byte(deliver), &callData)
		if err != nil {
            if deliver == "" {
                c.JSON(http.StatusOK, gin.H{
		        	"status": "params json error",
		        })
		    	return 
		    }
         
		}

		go callData.CallToBs(cf.ServerConf.Callback_url)
		
		c.JSON(http.StatusOK, gin.H{
			"status": "success",
		})
	})

	//测试接收
	r.POST("/rectest", func(c *gin.Context) {
		//data := c.PostForm("data")
		c.JSON(http.StatusOK, gin.H{
			"message": "OK",
		})
	})

	go func () {
		if err := srv.ListenAndServe(); err != nil {
			log.Println("listen: ", err)
		}
	}()


	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGUSR1, syscall.SIGUSR2)
	<-quit
	log.Println("shutingdown Server ...")
	// waiting for request finish
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("server shutdown error:", err)
	}
	log.Println("server shutdown finish!")
}

