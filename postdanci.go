package main

import (
	"./utils/logs"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/sdwolfe32/anirip/anirip"
	"github.com/widuu/gojson"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"
	"unicode"
)

const (
	login_url      string = "http://langeasy.com.cn/denglu.action"
	post_login_url string = "http://langeasy.com.cn/login.action"
	uname          string = "root@7jdg.com"
	pwd            string = "fuckyou"
)

func GetCookies() (result []*http.Cookie) {
	formData := url.Values{
		"name":   {uname},
		"passwd": {pwd},
	}
	loginReqHeaders := http.Header{}
	loginReqHeaders.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.11; rv:47.0) Gecko/20100101 Firefox/47.0")
	loginReqHeaders.Add("referer", login_url)
	loginReqHeaders.Add("content-type", "application/x-www-form-urlencoded")
	loginResponse, err := anirip.GetHTTPResponse("POST",
		post_login_url,
		bytes.NewBufferString(formData.Encode()),
		loginReqHeaders,
		nil)
	if err != nil {
		fmt.Println("[anirip] GetHTTPResponse Error ...")
		return
	}

	var logincookie = loginResponse.Cookies()
	return logincookie
}

func post_danci(danci, jieshi string, mycookie []*http.Cookie) string {
	formData := url.Values{
		"newwordlist": {"{\"word\":\"" + danci + "\",\"course\":\"*\",\"wordidx\":\"*\",\"infoidx\":\"100\",\"selection\":\"*\",\"info\":\"" + jieshi + "\",\"opcode\":\"1\"}"},
	}
	PostReqHeaders := http.Header{}
	PostReqHeaders.Add("Accept", "*/*")
	PostReqHeaders.Add("Origin", "chrome-extension://cklfipcjofdnmdolnfngpmokdaejidim")
	PostReqHeaders.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.11; rv:47.0) Gecko/20100101 Firefox/47.0")
	PostReqHeaders.Add("content-type", "application/x-www-form-urlencoded")
	PostReqHeaders.Add("DNT", "1")
	PostResponse, err := anirip.GetHTTPResponse("POST",
		"http://langeasy.com.cn/insertNewWord.action",
		bytes.NewBufferString(formData.Encode()),
		PostReqHeaders,
		mycookie)
	if err != nil {
		fmt.Println("[anirip] GetHTTPResponse Error ...")
		//return
	}
	defer PostResponse.Body.Close()
	body, err := ioutil.ReadAll(PostResponse.Body)
	if err != nil {
		fmt.Println("ReadAll Response Err:", err)
		//return
	}
	var rBody = string(body)
	//fmt.Println(rBody)
	if rBody == "null" {
		return danci + "单词添加成功"
	} else {
		return danci + "单词添加失败"
	}

}

func geturl(myurl string, mycookie []*http.Cookie) (result string) {
	showResponse, err := anirip.GetHTTPResponse("GET",
		myurl,
		nil,
		nil,
		mycookie)
	if err != nil {
		fmt.Println("[anirip] GetHTTPResponse Error ...")
		return
	}
	defer showResponse.Body.Close()
	body, err := ioutil.ReadAll(showResponse.Body)
	if err != nil {
		fmt.Println("ReadAll Response Err:", err)
		return
	}
	var rBody = string(body)
	return rBody
}

//取单词解释
func loadLexisList(danci string, mycookie []*http.Cookie) (result string) {
	respBody := geturl("http://langeasy.com.cn/loadLexisList.action?strict=1&word="+danci, mycookie)

	//var dat map[string]interface{}{}
	dat := map[string]interface{}{}
	//fmt.Println(respBody)
	if strings.Contains(respBody, "interpret") == false {
		return "false"
	} else {
		//
		json.Unmarshal([]byte(respBody), &dat)

		danci_jieshi := gojson.Json(respBody).Getindex(1).Getindex(1).Get("interpret").Tostring() //&{map[from:en to:zh]}

		//fmt.Println(danci_jieshi)
		danci_jieshi = strings.Replace(danci_jieshi, "\n", "\\n", -1)
		return danci_jieshi
	}
}
func getNewWord(danci string, mycookie []*http.Cookie) bool {
	respBody := geturl("http://langeasy.com.cn/getNewWord.action?word="+danci+"&infoidx=100", mycookie)

	if strings.Contains(respBody, "updatetime") {
		return true
	} else {
		return false
	}
	//return true
}

func tianjia_danci(danci_str string, cookie []*http.Cookie) bool {
	//logs.Logger.Info("查看 " + danci_str + " 是否在单词本")
	var isnewword = getNewWord(danci_str, cookie)
	if isnewword == false { //单词本里没有
		//开始查找单词解释
		jieshi := loadLexisList(danci_str, cookie)
		if jieshi == "false" {
			logs.Logger.Error(danci_str + " 单词不存在具体解释，请查正")
			return false
		} else {
			//logs.Logger.Info(danci_str + " 不存在单词本里，正在添加")
			tijian := post_danci(danci_str, jieshi, cookie)
			logs.Logger.Info(tijian)
			return true
		}
	} else {
		logs.Logger.Info(danci_str + " 单词已存在单词本中")
		return false
	}
}
func isNumber(s string) bool {
	s = strings.TrimSpace(s)
	n := len(s)
	if n == 0 {
		return false
	}
	if s[0] == '-' {
		s = s[1:]
	} else if s[0] == '+' {
		s = s[1:]
	}
	n = len(s)
	if n == 0 {
		return false
	}

	var isNumber = false
	i := 0
	for i < n && unicode.IsDigit(rune(s[i])) {
		i++
		isNumber = true
	}
	if i < n && s[i] == '.' {
		i++
		for i < n && unicode.IsDigit(rune(s[i])) {
			i++
			isNumber = true
		}
	}
	if isNumber && i < n && s[i] == 'e' {
		i++
		isNumber = false
		if i < n && (s[i] == '-' || s[i] == '+') {
			i++
		}
		for i < n && unicode.IsDigit(rune(s[i])) {
			i++
			isNumber = true
		}
	}
	return isNumber && i == n
}

// 对字符串数组去重
func clearRepeat(ss []string) (result []string) {
	m := make(map[string]bool)
	for _, v := range ss {
		if !m[v] {
			m[v] = true
			result = append(result, v)
		}
	}
	return result
}

//利用对比相邻单词是否一样的原理来去重
func RemoveDuplicatesAndEmpty(a []string) (ret []string) {
	a_len := len(a)
	for i := 0; i < a_len; i++ {
		if (i > 0 && a[i-1] == a[i]) || len(a[i]) == 0 {
			continue
		}
		ret = append(ret, a[i])
	}
	return
}

func HandingText(str string) []string {
	var ret []string
	str = strings.ToLower(str)                      //转小写
	str = strings.Replace(str, ".", " ", -1)        //删除点
	str = strings.Replace(str, ",", " ", -1)        //删除逗号
	str = strings.Replace(str, "\n", " ", -1)       //删除换行
	str = strings.Replace(str, "(", " ", -1)        //删除换行
	str = strings.Replace(str, ")", " ", -1)        //删除换行
	str = strings.Replace(str, ";", " ", -1)        //删除换行
	str = strings.Replace(str, "\"", " ", -1)       //删除换行
	str = strings.Replace(str, "/", " ", -1)        //删除换行
	str = strings.Replace(str, "'", " ", -1)        //删除换行
	str = strings.Replace(str, "*", " ", -1)        //删除换行
	str = strings.Replace(str, "-", " ", -1)        //删除换行
	str = strings.Replace(str, "=", " ", -1)        //删除换行
	str = strings.Replace(str, ":", " ", -1)        //删除换行
	str = strings.Replace(str, "[", " ", -1)        //删除换行
	str = strings.Replace(str, "]", " ", -1)        //删除换行
	str = strings.Replace(str, "?", " ", -1)        //删除换行
	str = strings.Replace(str, "‘", " ", -1)        //删除换行
	str = strings.Replace(str, "’", " ", -1)        //删除换行
	str = strings.Replace(str, "“", " ", -1)        //删除换行
	str = strings.Replace(str, "”", " ", -1)        //删除换行
	str = strings.Replace(str, "!", " ", -1)        //删除换行
	str = strings.Replace(str, "_", " ", -1)        //删除换行
	str = strings.Replace(str, "https://", " ", -1) //删除换行
	str = strings.Replace(str, "http://", " ", -1)  //删除换行

	all_danci := strings.Split(str, " ") //分割成数组
	//fmt.Println(all_danci)
	sort.Strings(all_danci)
	//fmt.Println(all_danci)
	//all_danci = RemoveDuplicatesAndEmpty(all_danci)
	all_danci = clearRepeat(all_danci) //map 去重复

	number := len(all_danci) //统计单词数

	for i := 0; i < number; i++ { //循环单词
		danci := all_danci[i]
		if len(danci) > 2 && !strings.Contains(danci, "@") && !isNumber(danci) { //单词大于两位的才进来
			//fmt.Fprintf(os.Stdout, "%d %v\n", i, danci)
			ret = append(ret, danci)

		}

	}
	return ret
}

func main() {
	str := `
Pompem is an open source exploit & vulnerability finder tool, designed to automate the search for Exploits and Vulnerability in the most important databases. Developed in Python, has a system of advanced search, that help the work of pen-testers and ethical hackers. In the current version, it performs searches in PacketStorm security,...

Read the full post at darknet.org.uk

    `
	t2 := time.Now()
	charlotteWeb := HandingText(str)
	//fmt.Println(charlotteWeb)
	logs.Logger.Info("单词总数：", len(charlotteWeb))
	var cookie = GetCookies()
	//fmt.Println(cookie)

	for i := 0; i < len(charlotteWeb); i++ {
		//logs.Logger.Info(charlotteWeb[i])
		tianjia_danci(charlotteWeb[i], cookie)
	}
	logs.Logger.Info("处理用时:", time.Now().Sub(t2))
}
