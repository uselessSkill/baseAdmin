package handle

import (
	"baseAdmin/loger"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

var wg sync.WaitGroup

var detailUrl = ""

func requestToDetail(gid string, newsChannel chan string) {

	defer wg.Done()

	resp, err := http.Get(detailUrl + gid)
	if err != nil {
		loger.WriteToFile("getDetailErrorLog", err.Error())
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		loger.WriteToFile("getDetailErrorLog", resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("read from resp.Body failed, err:%v\n", err)
		return
	}

	newsChannel <- string(body)
}

type Respon struct {
	Err  int                    `json:"errno"`
	Data map[string]interface{} `json:"data"`
	Msg  string                 `json:"errmsg"`
}

// @title GetDetails
// @version 1.0
// @description 获取稿件信息
// @termsOfService http://swagger.io/terms/
// @contact.name wangxiaoning

func GetDetails(gids []string) (returnData map[string]interface{}) {

	newsChannel := make(chan string, len(gids))

	for _, gid := range gids {
		wg.Add(1)
		go requestToDetail(gid, newsChannel)
	}

	wg.Wait()

	close(newsChannel) // 不关闭的情况下，会导致channel 阻塞
	returnData = make(map[string]interface{}, len(gids))

	for {
		news, ok := <-newsChannel
		if !ok {
			break
		}
		resp := Respon{}
		if err := json.Unmarshal([]byte(news), &resp); err == nil {
			if resp.Err == 200 {
				for key, val := range resp.Data {
					returnData[key] = val
				}
			}
		} else {
			fmt.Println(err)
		}
	}
	return
}

// 不用wait

func UGetDetails(gids []string) (returnData map[string]interface{}) {

	newsChannel := make(chan string)

	for _, gid := range gids {

		go func(gid string, newsChannel chan string) {
			resp, err := http.Get(detailUrl + gid)
			if err != nil {
				loger.WriteToFile("getDetailErrorLog", err.Error())
			}

			defer resp.Body.Close()

			if resp.StatusCode != 200 {
				loger.WriteToFile("getDetailErrorLog", resp.Status)
			}

			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				fmt.Printf("read from resp.Body failed, err:%v\n", err)
				return
			}

			newsChannel <- string(body)
		}(gid, newsChannel)
	}

	returnData = make(map[string]interface{}, len(gids))

	for i := 0; i < len(gids); i++ {
		news := <-newsChannel
		resp := Respon{}
		if err := json.Unmarshal([]byte(news), &resp); err == nil {
			if resp.Err == 200 {
				for key, val := range resp.Data {
					returnData[key] = val
				}
			}
		} else {
			fmt.Println(err)
		}
	}

	return
}
