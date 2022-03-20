package crawler

import (
	"Opendulum/global"
	"encoding/json"
	"fmt"
	"github.com/axgle/mahonia"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type JinDongCommentsRes struct {
	Url    string
	Header map[string]string
	Params map[string]string
}

func NewJingDongCommentsRes(id string) *JinDongCommentsRes {
	params := map[string]string{
		"productId":   id,
		"score":       "0",
		"sortType":    "5",
		"page":        "0",
		"pageSize":    "10",
		"isShadowSku": "0",
		"rid":         "0",
		"fold":        "1",
	}
	header := map[string]string{
		"user-agent": global.CrawlerSetting.UserAgent,
		"cookie":     global.CrawlerSetting.Cookie,
		"referer":    global.CrawlerSetting.Referer,
	}
	url := fmt.Sprintf("%s&productId=%s&score=%s&sortType=%s&page=%s&pageSize=%s&isShadowSku=%s&fold=%s",
		global.CrawlerSetting.JingDongCommentUrl,
		params["productId"],
		params["score"],
		params["sortType"],
		params["page"],
		params["pageSize"],
		params["isShadowSku"],
		params["fold"])
	return &JinDongCommentsRes{Url: url, Header: header, Params: params}
}

// Fetch 默认获取产品的前10条评论 想要获取跟多可以修改 NewJingDongCommentsRes 中的 page & pageSize 参数
func (t *JinDongCommentsRes) Fetch() []string {
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
	// Parse
	bodyStr := strings.Replace(string(body), "\n", " ", -1)    // 替换 "\n"
	bodyStr = mahonia.NewDecoder("gbk").ConvertString(bodyStr) // 转码中文

	bodyStr = strings.TrimPrefix(strings.TrimSuffix(bodyStr, ");"), "fetchJSON_comment98(")

	data := make(map[string]interface{})
	err = json.Unmarshal([]byte(bodyStr), &data)
	if err != nil {
		return nil
	}
	list := data["comments"].([]interface{})

	var comments []string

	for _, v := range list {
		comments = append(comments, v.(map[string]interface{})["content"].(string))
	}
	return comments
}
