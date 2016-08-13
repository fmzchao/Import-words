package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	str := `I've been working on this framework for about 7 months. I've worked really hard to make it powerful, yet accessible. I set out to launch with documentation as good as CodeIgniter from day one, and I think we did. The syntax is intuitive and expressive.
    `
	str = strings.Replace(str, ".", "", -1) //删除点
	str = strings.Replace(str, ",", "", -1) //删除逗号

	all_danci := strings.Split(str, " ") //分割成数组
	number := len(all_danci)             //统计单词数

	for i := 0; i < number; i++ { //循环单词
		danci := all_danci[i]
		if len(danci) > 2 { //单词大于两位的才进来
			fmt.Fprintf(os.Stdout, "%d %v\n", i, danci)
			//
			fmt.Print("...")
		}

	}

}
