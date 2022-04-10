package crawler

import (
	"Opendulum/global"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"
)

type JingDongGoodsRes struct {
	Url    string
	Header map[string]string
	Params map[string]string
}

func NewJingDongGoodsRes(id string) *JingDongGoodsRes {
	params := map[string]string{
		"productId": id,
	}
	header := map[string]string{
		"user-agent": global.CrawlerSetting.UserAgent,
	}
	url := fmt.Sprintf("%s/%s.html", global.CrawlerSetting.JingDongGoodsHtml, params["productId"])
	return &JingDongGoodsRes{Url: url, Header: header, Params: params}
}

// Fetch 默认获取产品的前10条评论 想要获取跟多可以修改 NewJingDongCommentsRes 中的 page & pageSize 参数
func (t *JingDongGoodsRes) Fetch() []string {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", t.Url, nil)

	req.Header.Set("user-agent", t.Header["user-agent"])

	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return nil
	}
	if resp.StatusCode != 200 {
		log.Println("Http status code: ", resp.StatusCode)
		return nil
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Read error", err)
		return nil
	}
	//bodystr := mahonia.NewDecoder("utf8").ConvertString(string(body)) // 转码中文
	bodystr := strings.Replace(string(body), "\n", "", -1)                   // 替换 "\n"
	reMainBody := regexp.MustCompile(`<div class="Ptable-item">(.*?)</div>`) // 找到信息主体 包含净含量 保质期 生产许可证 产品标准号
	mainBody := reMainBody.FindString(bodystr)

	reRs := regexp.MustCompile(`<dd>(.*?)</dd>`)
	rs := reRs.FindAllString(mainBody, -1)

	for k, v := range rs {
		rs[k] = strings.TrimPrefix(strings.TrimSuffix(v, "</dd>"), "<dd>")
	}
	return rs
}
