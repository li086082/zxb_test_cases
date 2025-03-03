package cases

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
	"zxb_test_cases/config"
	"zxb_test_cases/util"
)

var reportUrl = "https://wlc.nppa.gov.cn/test/collection/loginout/"

type CollectionsList struct {
	Collections []Collections `json:"collections"`
}

type Collections struct {
	No int    `json:"no"`
	Si string `json:"si"`
	Bt int64  `json:"bt"`
	Ot int64  `json:"ot"`
	Ct int64  `json:"ct"`
	Di string `json:"di"`
	Pi string `json:"pi"`
}

func (ca Cases) case07() {
	cc := Collections{
		No: 1,
		Si: "hysdk001",
		Bt: 1,
		Ot: time.Now().Unix(),
		Ct: 2, // 游客模式
		Di: "abc123456789",
		Pi: "1fffbjzos82bs9cnyj1dna7d6d29zg4esnh99u",
	}
	ca.report(cc, config.Cfg.Cases["case07"])
}

func (ca Cases) case08() {
	cc := Collections{
		No: 1,
		Si: "hysdk001",
		Bt: 1,
		Ot: time.Now().Unix(),
		Ct: 0, // 已认证用户
		Di: "abc123456789",
		Pi: "1fffbjzos82bs9cnyj1dna7d6d29zg4esnh99u",
	}
	ca.report(cc, config.Cfg.Cases["case08"])
}

func (Cases) report(col Collections, cases config.Cases) {

	if cases.Pass {
		return
	}

	header := map[string]string{
		"Content-Type": "application/json;charset=utf-8",
		"appId":        config.Cfg.AppId,
		"bizId":        config.Cfg.BizId,
		"timestamps":   fmt.Sprintf("%d", currentTime),
	}

	var cc CollectionsList
	cc.Collections = make([]Collections, 1)
	cc.Collections[0] = col

	marshal, err1 := json.Marshal(cc)
	if err1 != nil {
		log.Fatalln(err1)
	}
	marshalStr := string(marshal)

	encipher, err2 := util.Encipher(marshalStr, config.Cfg.Secret)
	if err2 != nil {
		log.Fatalln(err2)
	}

	var str string
	str += config.Cfg.Secret
	str += "appId"
	str += config.Cfg.AppId
	str += "bizId"
	str += config.Cfg.BizId
	str += "timestamps"
	str += fmt.Sprintf("%d", currentTime)
	str += encipher

	signature, err2 := util.Sign(str)
	if err2 != nil {
		log.Fatalln(err2)
	}
	header["sign"] = signature

	fullUrl := reportUrl + cases.Prefix

	response, err3 := util.Request("POST", fullUrl, header, encipher)
	if err3 != nil {
		log.Fatalln("report request send fail", err3)
	}

	log.Println("================")
	log.Println(fullUrl)
	log.Println(response)
	log.Println("================")
}
