package http_module

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	Header_Set_Content_json = 0
	Header_Add_X_auth_token = 1
	Header_Set_Content_text = 2
)

const timeout_limit_sec time.Duration = 120

//https://velog.io/@roeniss/content-type%EC%97%90-%EB%94%B0%EB%A5%B8-golang-POST-%EB%8D%B0%EC%9D%B4%ED%84%B0-%EC%A0%84%EC%86%A1-%EB%B0%A9%EB%B2%95

func Post(url string, payload string, token string, headers ...int) (int, map[string][]string, string) {

	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(payload)))
	//req, err := http.NewRequest("POST", url, bytes.NewBufferString(payload))

	for _, header_const_num := range headers {
		add_header(req, token, header_const_num)
	}

	client := &http.Client{Timeout: timeout_limit_sec * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		Logger.Error(err)
		return 0, nil, ""
	}
	defer resp.Body.Close()

	rep_code := resp.StatusCode
	rep_header_map := resp.Header
	rep_body, _ := ioutil.ReadAll(resp.Body)
	rep_body_str := string(rep_body)

	return rep_code, rep_header_map, rep_body_str
}

func Get(url string, token string, headers ...int) (int, map[string][]string, string) {

	req, err := http.NewRequest("GET", url, nil)

	for _, header_const_num := range headers {
		add_header(req, token, header_const_num)
	}

	client := &http.Client{Timeout: timeout_limit_sec * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		Logger.Error(err)
		return 0, nil, ""
	}
	defer resp.Body.Close()

	rep_code := resp.StatusCode
	rep_header_map := resp.Header
	rep_body, _ := ioutil.ReadAll(resp.Body)
	rep_body_str := string(rep_body)

	return rep_code, rep_header_map, rep_body_str
}

func add_header(req *http.Request, token string, header_type int) {
	switch header_type {
	case Header_Set_Content_json:
		req.Header.Set("Content-Type", "application/json")
	case Header_Add_X_auth_token:
		req.Header.Add("X-Auth-Token", token)
	case Header_Set_Content_text:
		req.Header.Set("Content-Type", "text/html; charset=utf-8")
	}
}
