package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func main() {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.nasdaq.com/api/quote/TEF/dividends?assetclass=stocks", nil)

	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("authority", "api.nasdaq.com")
	req.Header.Set("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
	req.Header.Set("accept-language", "en-US,en;q=0.9,zh-TW;q=0.8,zh;q=0.7,und;q=0.6,zh-CN;q=0.5")
	req.Header.Set("cache-control", "max-age=0")
	req.Header.Set("sec-ch-ua", `"Not A(Brand";v="99", "Google Chrome";v="121", "Chromium";v="121"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"macOS"`)
	req.Header.Set("sec-fetch-dest", "document")
	req.Header.Set("sec-fetch-mode", "navigate")
	req.Header.Set("sec-fetch-site", "none")
	req.Header.Set("sec-fetch-user", "?1")
	req.Header.Set("upgrade-insecure-requests", "1")
	req.Header.Set("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.0.0 Safari/537.36")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	r := &Response{}
	err = json.Unmarshal(bodyText, r)
	if err != nil {
		log.Fatal(err)
	}

	exDividendData, err := time.Parse("01/02/2006", r.Data.ExDividendDate)
	if err != nil {
		log.Fatal(err)
	}
	now := time.Now()
	if int(now.Sub(exDividendData).Hours()/24) < 7 {
		fmt.Println("exDividendData: ", exDividendData)
		fmt.Println(r.Data.ExDividendDate)
		return
	}
}
