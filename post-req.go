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
    resp, err := http.PostForm(server+"addcrypted2",url.Values{"passwords": {"filecrypt.cc"}, "source": {"filecrypt.cc"}, "jk":{"function f(){ return '32303532393132393631363931363138';}"},
    "crypted":{"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"}})
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
    send_post_crypt2()
}
