package cases

import (
	"fmt"
	"log"
	"strings"
	"zxb_test_cases/config"
	"zxb_test_cases/util"
)

var authUrl = "https://wlc.nppa.gov.cn/test/authentication/check/"

func (ca Cases) case01() {
	body := map[string]string{
		"ai":    "100000000000000001",
		"name":  "某一一",
		"idNum": "110000190101010001",
	}
	keys := []string{"ai", "name", "idNum"}

	ca.auth(body, keys, config.Cfg.Cases["case01"])
}

func (ca Cases) case02() {
	body := map[string]string{
		"ai":    "200000000000000001",
		"name":  "某二一",
		"idNum": "110000190201010009",
	}
	keys := []string{"ai", "name", "idNum"}

	ca.auth(body, keys, config.Cfg.Cases["case02"])
}

func (ca Cases) case03() {
	body := map[string]string{
		"ai":    "330000000000000001",
		"name":  "张三",
		"idNum": config.Cfg.RealCard,
	}
	keys := []string{"ai", "name", "idNum"}

	ca.auth(body, keys, config.Cfg.Cases["case03"])
}

func (Cases) auth(data map[string]string, sortKeys []string, cases config.Cases) {

	if cases.Pass {
		return
	}

	header := map[string]string{
		"Content-Type": "application/json;charset=utf-8",
		"appId":        config.Cfg.AppId,
		"bizId":        config.Cfg.BizId,
		"timestamps":   fmt.Sprintf("%d", currentTime),
	}

	var strBody string
	for _, key := range sortKeys {
		strBody += fmt.Sprintf("\"%s\":\"%s\",", key, data[key])
	}
	strBody = "{" + strings.TrimSuffix(strBody, ",") + "}"

	encipher, err1 := util.Encipher(strBody, config.Cfg.Secret)
	if err1 != nil {
		log.Fatalln("auth encipher fail", err1)
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
		log.Fatalln("auth sign fail", err2)
	}

	header["sign"] = signature

	response, err3 := util.Request("POST", authUrl+cases.Prefix, header, encipher)
	if err3 != nil {
		log.Fatalln("request send fail", err3)
	}

	log.Println("================")
	log.Println(authUrl + cases.Prefix)
	log.Println(response)
	log.Println("================")
}
