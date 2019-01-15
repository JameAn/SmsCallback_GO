package libs

import (
	"encoding/json"
	"SmsCallback_Go/conf"
)
//大汉三通callback
type Deliver struct {
	Result string
	Desc string
	Delivers []Delivers
}

type Delivers struct {
	Phone string
	Content string
	Subcode string
	Delivertime string
}

type BsType string

type DhstS struct {
	LogHandler *LogGer
}

var Dhst DhstS

func init(){
    Dhst = DhstS{}
}

//解析subCode前两位
func (ds Delivers) DeSubcode() (Delivers, BsType){
	var bsType BsType
	bsType = BsType(ds.Subcode[0:2]) 

	return ds, bsType
}

func (d Deliver) CallToBs(cb_url conf.JxhzSt) {
	for _,value := range d.Delivers {
		delivers, bsType := value.DeSubcode()
		if bsType == "42" {
			data, err :=  json.Marshal(delivers)
			if err != nil {
				break
			}
			SendInfo := &SendData{
                JsonData: data,
				RemoteUrl: cb_url.Jxhz,
				LogHandler: Dhst.LogHandler,
			}
		    SendInfo.PostTo()
		}
	}
}
