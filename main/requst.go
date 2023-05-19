package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"time"
)

type StructHoliday struct {
	Holiday bool   `json:"holiday"`
	Name    string `json:"name"`
	Wage    int    `json:"wage"`
	Date    string `json:"date"`
	Rest    int    `json:"rest"`
	T0      int64  `json:"t0"`
	T1      int64  `json:"t1"`
}

type StructResp struct {
	Code    int `json:"code"`
	Holiday map[string]*StructHoliday
}

const URL = "https://timor.tech/api/holiday/year/"

func RequestHolidayJson(year int, file string) {
	client := http.Client{}
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s%d?week=Y", URL, year), nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36")
	// 发起请求
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		os.Exit(3)
	}
	// 程序在使用完回复后必须关闭回复的主体
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		os.Exit(4)
	}
	js := &StructResp{}
	err = json.Unmarshal(data, js)
	if err != nil {
		fmt.Println(err)
		os.Exit(5)
	}
	for _, h := range js.Holiday {
		tm, _ := time.Parse(time.DateOnly, h.Date)
		h.T0 = tm.UnixMilli()
		h.T1 = tm.AddDate(0, 0, 1).UnixMilli()
	}

	os.MkdirAll(path.Dir(file), 0666)
	f, _ := os.OpenFile(file, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	d, err := json.Marshal(js.Holiday)
	if err != nil {
		fmt.Println(err)
		os.Exit(6)
	}
	_, err = f.Write(d)
	if err != nil {
		fmt.Println(err)
		os.Exit(7)
	}
}
