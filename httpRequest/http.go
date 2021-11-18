package httpRequest

import (
	"crypto/tls"
	"encoding/json"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
	"time"
)

var h = &HttpRequest{
	Transport:&http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
			DualStack: true,
		}).DialContext,
		MaxIdleConns:          50,
		IdleConnTimeout:       90 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	},
}

type HttpRequest struct {
	Transport *http.Transport
}

func GetMap(url string)(map[string]interface{},error){
	return h.getMap(url)
}
func (self *HttpRequest)getMap(url string) (map[string]interface{},error) {
	client:=&http.Client{}
	client.Transport=self.Transport
	resp,err:=client.Get(url)

	if err!=nil{
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	var data map[string]interface{}
	err=json.Unmarshal(body,&data)
	if err!=nil{
		return nil, err
	}

	return data,nil
}

func PostMap(url string,req interface{})(map[string]interface{},error){
	return h.postMap(url ,req )
}
func (self *HttpRequest)postMap(url string,req interface{})(map[string]interface{},error){
	b,err:=json.Marshal(req)
	if err!=nil{
		return nil,err
	}

	client:=&http.Client{}
	client.Transport=self.Transport
	resp, err:=client.Post(url,"application/json",strings.NewReader(string(b)))
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	var data map[string]interface{}
	err=json.Unmarshal(body,&data)
	if err!=nil{
		return nil, err
	}
	return data,nil
}