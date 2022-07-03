package explorer

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/MaxKlaxxMiner/yacgui-web/modules/yacboard/crc64"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"server/lichess/explorer/ratings"
	"server/lichess/explorer/speeds"
	"strings"
	"time"
)

var requestLimiter = time.Millisecond * 1000

var lastRequest time.Time

func Request(fen string, speeds speeds.Speeds, ratings ratings.Ratings) *Response {
	fullUrl := fmt.Sprintf("https://explorer.lichess.ovh/lichess?variant=standard&topGames=0&recentGames=0&moves=256&speeds=%s&ratings=%s&fen=%s", speeds, ratings, url.QueryEscape(fen))
	crc := crc64.CrcStart.UpdateString(fullUrl)
	cacheFile := fmt.Sprintf("lichess-cache/%016x.json", crc)
	jsonResult, err := os.ReadFile(cacheFile)

	if err != nil {
		for time.Since(lastRequest) < requestLimiter {
			time.Sleep(time.Millisecond * 10)
		}
		lastRequest = time.Now()

		fmt.Println("get:", fullUrl)
		resp, err := http.Get(fullUrl)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		jsonResult, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}

		if strings.Contains(string(jsonResult), "429 Too Many Requests") {
			requestLimiter += time.Millisecond * 100
			fmt.Println("429 Too Many Requests ->", requestLimiter)
			time.Sleep(requestLimiter * 10)
			return Request(fen, speeds, ratings)
		}

		err = os.WriteFile(cacheFile, jsonResult, 644)
		if err != nil {
			if _, err := os.Stat("lichess-cache"); errors.Is(err, os.ErrNotExist) {
				os.Mkdir("lichess-cache", 644)
			}
		}
	}

	data := new(Response)

	err = json.Unmarshal(jsonResult, &data)
	if err != nil {
		panic(err)
	}

	return data
}
