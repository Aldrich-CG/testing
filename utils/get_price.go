package utils

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

// 示例函数，模拟获取股票实时价格的函数
func getStockQuote(stockCode string) (map[string]float64, error) {
	// 在这里实现获取股票实时价格的逻辑
	// 这里只是简单地返回一个模拟的价格
	// 实际情况下，你需要调用真实的股票行情API来获取实时价格
	price := 50.0 + float64(len(stockCode))*0.1

	return map[string]float64{
		"当前价格": price,
	}, nil
}

func GetStockQuote(stockCode string) (map[string]float64, error) {
	url := fmt.Sprintf("http://hq.sinajs.cn/list=%s", stockCode)
	var price float64
	// 创建 HTTP 请求
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// 添加 Referer 到请求头
	req.Header.Add("Referer", "https://finance.sina.com.cn")

	// 发起 HTTP 请求
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	// 解析返回数据
	data := make(map[string]string)
	lines := string(body)
	quoteFields := strings.Split(lines, ",")
	if len(quoteFields) >= 32 {
		data["股票名称"] = quoteFields[0][strings.Index(quoteFields[0], "\"")+1:]
		price, err = strconv.ParseFloat(quoteFields[3], 64)
		data["股票代码"] = stockCode
		// 其他字段可按需添加
	}

	return map[string]float64{
		"当前价格": price,
	}, nil
}
