package utils

import (
	"crypto/tls"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
	"time"
)

var client http.Client
var timeout = 5 * time.Second
var keepAlive = 30 * time.Second

func init() {
	client = http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			DialContext: (&net.Dialer{
				Timeout:   timeout,
				KeepAlive: keepAlive,
			}).DialContext,
		},
	}
}

// function should be return value as follow:
// - respond http code is 200
func Get(url string) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		//logger.Error("NewRequest error", logger.Any("err", err), logger.String("url", url))
		return nil, err
	}

	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bz, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		return bz, nil
	}
	return nil, nil
}

func Forward(req *http.Request, url string) (bz []byte, err error) {
	r := forkRequest(req, url)
	res, err := client.Do(r)
	defer res.Body.Close()
	if err != nil || res.StatusCode != 200 {
		if _, err := ioutil.ReadAll(res.Body); err == nil {
			res.Body.Close()
		}
		return bz, err
	}

	bz, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return bz, err
	}
	return bz, err
}

func forkRequest(req *http.Request, url string) *http.Request {
	r, _ := http.NewRequest(req.Method, url, req.Body)
	for k, v := range req.Header {
		for _, vv := range v {
			r.Header.Add(k, vv)
		}
	}
	r.RemoteAddr = req.RemoteAddr
	r.Form = req.Form
	r.PostForm = req.PostForm
	r.MultipartForm = req.MultipartForm
	r.Proto = req.Proto
	r.ProtoMajor = req.ProtoMajor
	r.ProtoMinor = req.ProtoMinor
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	return r
}

func GetIpAddr(req *http.Request) string {
	addrs := req.Header["X-Forwarded-For"]
	if len(addrs) > 0 && "unknown" != addrs[0] {

		return addrs[0]
	}

	addrs = req.Header["Proxy-Client-IP"]
	if len(addrs) > 0 && "unknown" != addrs[0] {
		return addrs[0]
	}

	addrs = req.Header["WL-Proxy-Client-IP"]
	if len(addrs) > 0 && "unknown" != addrs[0] {
		return addrs[0]
	}

	addrs = req.Header["X-Real-IP"]
	if len(addrs) > 0 && "unknown" != addrs[0] {
		return addrs[0]
	}

	addrs = strings.Split(req.RemoteAddr, ":")
	if len(addrs) > 0 && "unknown" != addrs[0] {
		return addrs[0]
	}

	return ""
}
