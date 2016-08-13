package main

import (
	"fmt"
	"sort"
	"strings"
)

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

func HandingText(str string) (ret []string) {
	str = strings.ToLower(str)               //转小写
	str = strings.Replace(str, ".", "", -1)  //删除点
	str = strings.Replace(str, ",", "", -1)  //删除逗号
	str = strings.Replace(str, "\n", "", -1) //删除换行

	all_danci := strings.Split(str, " ") //分割成数组
	//fmt.Println(all_danci)
	sort.Strings(all_danci)
	//fmt.Println(all_danci)
	all_danci = RemoveDuplicatesAndEmpty(all_danci)

	number := len(all_danci) //统计单词数

	for i := 0; i < number; i++ { //循环单词
		danci := all_danci[i]
		if len(danci) > 2 { //单词大于两位的才进来
			//fmt.Fprintf(os.Stdout, "%d %v\n", i, danci)
			ret = append(ret, danci)
		}

	}
	return
}
func main() {
	str := `
    I've been working on this framework for about 7 months. I've worked really hard to make it powerful, yet accessible. I set out to launch with documentation as good as CodeIgniter from day one, and I think we did. The syntax is intuitive and expressive been make day one.
    `
	fmt.Println(HandingText(str))

	fmt.Print("End ....\n\n")

}
