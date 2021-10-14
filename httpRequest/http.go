package httpRequest

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

func GetMap(url string) (map[string]interface{},error) {
	resp,err:=http.Get(url)
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

func PostMap(url string,req map[string]interface{})(map[string]interface{},error){
	b,err:=json.Marshal(req)
	if err!=nil{
		return nil,err
	}
	resp, err:=http.Post(url,"",strings.NewReader(string(b)))
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