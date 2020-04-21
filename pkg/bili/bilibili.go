package bili

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"math"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"sync"
	"time"
)

const (
	xor = 177451812
	add = 8728348608
)

var once sync.Once

var (
	table = []byte("fZodR9XQDSUm21yCkr6zBqiveYah8bt4xsWpHnJE7jL5VG3guMTKNPAwcF")
	s     = []int{11, 10, 3, 8, 4, 6}

	tr map[byte]int64
)

// transform BV to AV
func bvToAv(bv string) (av int64, err error) {
	defer func() {
		if x := recover(); x != nil {
			err = fmt.Errorf("Transform BV: %s panic: %v", bv, x)
		}
	}()
	once.Do(func() {
		tr = make(map[byte]int64)
		n := len(table)
		for i := 0; i < n; i++ {
			tr[table[i]] = int64(i)
		}
	})
	var r int64
	for i := 0; i < 6; i++ {
		r += tr[bv[s[i]]] * int64(math.Pow(float64(58), float64(i)))
	}
	av = (r - add) ^ xor
	return
}

func avToBv(av int64) (bv string, err error) {
	defer func() {
		if x := recover(); x != nil {
			err = fmt.Errorf("Transform AV: %s panic: %v", bv, x)
		}
	}()
	av = av ^ xor + add
	r := []byte("BV1  4 1 7  ")
	for i := 0; i < 6; i++ {
		p := math.Pow(float64(58), float64(i))
		v := math.Floor(float64(av) / p)
		idx := int(v) % 58
		r[s[i]] = table[idx]
	}
	bv = string(r)
	return
}

// BvToAv transform BV to AV
func BvToAv(bv string) (string, error) {
	av, err := bvToAv(bv)
	if err != nil {
		return "", err
	}
	return strconv.FormatInt(av, 10), nil
}

// AvToBv transfrom AV to BV
func AvToBv(av string) (string, error) {
	id, err := strconv.ParseInt(av, 10, 64)
	if err != nil {
		return "", err
	}
	return avToBv(id)
}

// Stat for bilibili video
type Stat struct {
	Aid        int64  `json:"aid"`
	BV         string `json:"bvid"`
	View       int64  `json:"view"`
	Danmaku    int64  `json:"danmaku"`
	Reply      int64  `json:"reply"`
	Favorite   int64  `json:"favorite"`
	Coin       int64  `json:"coin"`
	Share      int64  `json:"share"`
	Like       int64  `json:"like"`
	NowRank    int    `json:"now_rank"`
	HisRank    int    `json:"his_rank"`
	NoReprint  int    `json:"no_reprint"`
	Copyright  int    `json:"copyright"`
	Argue      string `json:"argue_msg"`
	Evaluation string `json:"evaluation"`
}

const (
	apiHost   = "api.bilibili.com"
	statURL   = "http://api.bilibili.com/x/web-interface/archive/stat"
	userAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:74.0) Gecko/20100101 Firefox/74.0"

	timeout = 5 * time.Second
)

// VideoStat returns video's statistics
func VideoStat(vid string) (*Stat, error) {
	query := make(url.Values)
	if strings.HasPrefix(strings.ToLower(vid), "bv") {
		query.Add("bvid", vid)
	} else {
		_, err := strconv.ParseInt(vid, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("Invalid video ID: %s", vid)
		}
		query.Add("aid", vid)
	}
	_url := fmt.Sprintf("%s?%s", statURL, query.Encode())
	res, err := request(http.MethodGet, _url, nil)
	if err != nil {
		return nil, err
	}
	var resp struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Data    Stat   `json:"data"`
	}
	if err = json.Unmarshal(res, &resp); err != nil {
		return nil, err
	}
	if resp.Code != 0 {
		return nil, fmt.Errorf("Request failed: %d - %s", resp.Code, resp.Message)
	}
	return &resp.Data, nil
}

func request(method, url string, body io.Reader) ([]byte, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Add("User-Agent", userAgent)
	req.Header.Add("Host", apiHost)
	client := http.Client{
		Timeout: timeout,
	}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Request bilibili api failed: %s", res.Status)
	}
	defer res.Body.Close()

	bts, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return bts, nil
}
