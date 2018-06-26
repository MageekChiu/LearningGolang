package main

import (
	"net/http"
	"io/ioutil"
	"fmt"
	"strings"
	"net/url"
	log "github.com/sirupsen/logrus"
)

func main()  {
	httpGet("https://www.douban.com/")
	//httpGet("https://www.sina.com.cn/")
}

/**
简单get
 */
func httpGet(url1 string) {
	resp, err := http.Get(url1)
	if err != nil {
		log.WithFields(log.Fields{
			"animal": "walrus",
			"error":err,
		}).Fatal("an error get")
	}


	defer resp.Body.Close()

	//body, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(string(body))

	log.WithFields(log.Fields{
		"header":resp.Header,
	}).Info("successfully got")
}

/**
x-www-form-urlencoded post
 */
func httpPost(url1 string) {
	resp, err := http.Post(url1,
		"application/x-www-form-urlencoded",
		strings.NewReader("name=cjb"))
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))
}

/**
表单 form
 */
func httpPostForm(url1 string) {
	resp, err := http.PostForm(url1,
		url.Values{"key": {"Value"}, "id": {"123"}})

	if err != nil {
		// handle error
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))

}

/**
包含复杂header
 */
func httpDo(url1 string) {
	client := &http.Client{}

	req, err := http.NewRequest("POST", url1, strings.NewReader("name=cjb"))
	if err != nil {
		// handle error
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Cookie", "name=anny")

	resp, err := client.Do(req)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))
}