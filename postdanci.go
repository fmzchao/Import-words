package main

import (
	"fmt"
	"sort"
	"strings"
	"time"
	"unicode"
	//strconv  //json += "\\u"+strconv.FormatInt(int64(rint), 16) // json
	//"io"
	//"net/http"
	// "net/http/cookiejar"
	// "net/url"
	//"os"
)

const (
	//getUrl         string = "http://jwc.sut.edu.cn/ACTIONQUERYSTUDENTPIC.APPPROCESS?ByStudentNO=null"
	login_url      string = "http://langeasy.com.cn/denglu.action"
	post_login_url string = "http://langeasy.com.cn/login.action"
	uname          string = "root@7jdg.com"
	pwd            string = "fuckyou"
)

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
	str = strings.ToLower(str)                //转小写
	str = strings.Replace(str, ".", " ", -1)  //删除点
	str = strings.Replace(str, ",", " ", -1)  //删除逗号
	str = strings.Replace(str, "\n", " ", -1) //删除换行
	str = strings.Replace(str, "(", " ", -1)  //删除换行
	str = strings.Replace(str, ")", " ", -1)  //删除换行
	str = strings.Replace(str, ";", " ", -1)  //删除换行
	str = strings.Replace(str, "\"", " ", -1) //删除换行

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
GitHub Terms of Service
By using the GitHub.com web site ("Service"), or any services of GitHub, Inc ("GitHub"), you are agreeing to be bound by the following terms and conditions ("Terms of Service"). IF YOU ARE ENTERING INTO THIS AGREEMENT ON BEHALF OF A COMPANY OR OTHER LEGAL ENTITY, YOU REPRESENT THAT YOU HAVE THE AUTHORITY TO BIND SUCH ENTITY, ITS AFFILIATES AND ALL USERS WHO ACCESS OUR SERVICES THROUGH YOUR ACCOUNT TO THESE TERMS AND CONDITIONS, IN WHICH CASE THE TERMS "YOU" OR "YOUR" SHALL REFER TO SUCH ENTITY, ITS AFFILIATES AND USERS ASSOCIATED WITH IT. IF YOU DO NOT HAVE SUCH AUTHORITY, OR IF YOU DO NOT AGREE WITH THESE TERMS AND CONDITIONS, YOU MUST NOT ACCEPT THIS AGREEMENT AND MAY NOT USE THE SERVICES.
    `
	t2 := time.Now()
	charlotteWeb := HandingText(str)
	fmt.Println(charlotteWeb)
	fmt.Println("单词总数：", len(charlotteWeb))
	fmt.Print("End ....\n\n")
	fmt.Println("去重用时:", time.Now().Sub(t2))

}
