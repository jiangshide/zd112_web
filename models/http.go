package models

import ("github.com/astaxie/beego"
 "net/http"
 // "fmt"
 "encoding/json"
 "strings"
 "io/ioutil"
 )

type Http struct{
	Author string
	Url string
	Method string
	Data string
	Model interface{}	
}

func (this *Http)Get() error{
	req, _ := http.NewRequest("GET", this.Url, nil)
	if len(this.Author) > 0{
		req.Header.Add("Authorization",this.Author)
	}
	res, _ := http.DefaultClient.Do(req)
 
	defer res.Body.Close()
	// body, _ := ioutil.ReadAll(res.Body)
 	err := json.NewDecoder(res.Body).Decode(&this.Model);
	beego.Info("resp:",res," | err:",err)

	// resp,err := http.Get(this.Url)
	// if resp.StatusCode != http.StatusOK{
	// 	return fmt.Errorf("request false: %s", resp.Status)
	// }

	// body := resp.Body
	// beego.Info("body:",body)
	// err = json.NewDecoder(body).Decode(&this.Model);
	// beego.Info("resp:",resp," | err:",err)
	return	err
}

//the temp
func (this *Http)Post()(err error,response string){		
    request, _ := http.NewRequest("POST", this.Url, strings.NewReader(this.Data))
    request.Header.Add("Content-Type", "application/json")  //添加请求头
    request.Header.Add("Authorization", this.Author)  //添加请求头
    //post数据并接收http响应
    resp,err :=http.DefaultClient.Do(request)
    if err==nil{
        respBody,_ :=ioutil.ReadAll(resp.Body)
        response = string(respBody)
    }
    return
}

//the temp
func (this *Http)Delete()(err error,response string){		
    request, _ := http.NewRequest("DELETE", this.Url, nil)
    request.Header.Add("Content-Type", "application/json")  //添加请求头
    request.Header.Add("Authorization", this.Author)  //添加请求头
    //post数据并接收http响应
    resp,err :=http.DefaultClient.Do(request)
    if err==nil{
        respBody,_ :=ioutil.ReadAll(resp.Body)
        response = string(respBody)
    }
    return
}