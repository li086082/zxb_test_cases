package cases

import (
	"fmt"
	"log"
	"zxb_test_cases/config"
	"zxb_test_cases/util"
)

var queryUrl = "https://wlc.nppa.gov.cn/test/authentication/query/"

func (ca Cases) case04() {
	ca.query("100000000000000001", config.Cfg.Cases["case04"])
}

func (ca Cases) case05() {
	ca.query("200000000000000001", config.Cfg.Cases["case05"])
}

func (ca Cases) case06() {
	ca.query("300000000000000001", config.Cfg.Cases["case06"])
}

func (Cases) query(ai string, cases config.Cases) {

	if cases.Pass {
		return
	}

	header := map[string]string{
		"appId":      config.Cfg.AppId,
		"bizId":      config.Cfg.BizId,
		"timestamps": fmt.Sprintf("%d", currentTime),
	}

	var str string
	str += config.Cfg.Secret
	str += "ai"
	str += ai
	str += "appId"
	str += config.Cfg.AppId
	str += "bizId"
	str += config.Cfg.BizId
	str += "timestamps"
	str += fmt.Sprintf("%d", currentTime)

	sign, err1 := util.Sign(str)
	if err1 != nil {
		log.Fatalln("query sign fail", err1)
	}

	header["sign"] = sign

	fullUrl := queryUrl + cases.Prefix + "?ai=" + ai

	response, err2 := util.Request("GET", fullUrl, header, "")
	if err2 != nil {
		log.Fatalln("query request send fail", err2)
	}

	log.Println("================")
	log.Println(fullUrl)
	log.Println(response)
	log.Println("================")
}
