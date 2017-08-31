package main

import (
    "fmt"
    "net/http"
    "net/url"
    "io/ioutil"
)

var server = "http://localhost:9666/"

func send_post_plain(){
    resp, err := http.PostForm(server+"flash/add",url.Values{"passwords": {"password_is_secret"}, "source": {"http://test-client.de"}, "urls":{"http://test\nhttp://test2"}})
    if err != nil {
        fmt.Println(err)
        return
    }
    body, _ := ioutil.ReadAll(resp.Body)    
    fmt.Println(resp.StatusCode)
    fmt.Println(resp.Header.Get("Content-Type"))
    fmt.Println(string(body))    
}
// kept this func for later commits.
func send_post_crypt2(){
    resp, err := http.PostForm(server+"addcrypted2",url.Values{"passwords": {"myPassword"}, "source": {"http://jdownloader.org/spielwiese"}, "jk":{"31323334353637383930393837363534"}, "crypted":{"DRurBGEf2ntP7Z0WDkMP8e1ZeK7PswJGeBHCg4zEYXZSE3Qqxsbi5EF1KosgkKQ9SL8qOOUAI+eDPFypAtQS9A=="}})
    if err != nil {
        fmt.Println(err)
        return
    }
    body, _ := ioutil.ReadAll(resp.Body)
    fmt.Println(resp.StatusCode)
    fmt.Println(resp.Header.Get("Content-Type"))
    fmt.Println(string(body))
}

func main(){
    send_post_plain()
    //send_post_crypt2()
}
