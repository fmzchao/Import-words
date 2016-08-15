// test
package main

import (
	"io/ioutil"
	"net/http"
	"os"

	"github.com/mozillazg/request"
)

func main() {
	c := &http.Client{}
	req := request.NewRequest(c)
	resp, _ := req.Get("http://segmentfault.com/")
	defer resp.Body.Close() // **Don't forget close the response body**
	body, _ := ioutil.ReadAll(resp.Body)
	fr, _ := os.Create("request.html")
	fr.Write(body)
	res, _ := http.Get("http://segmentfault.com/")
	truebody, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	ft, _ := os.Create("get.html")
	ft.Write(truebody)

}
