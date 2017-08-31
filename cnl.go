
package main
import (
    "fmt"
    "net/http"
    //"reflect"
    "net/http/httputil"
    //"encoding/base64"
    //"crypto/aes"
    //"crypto/cipher"
    //"strings"
    
)

var Debug = 0




func handler(w http.ResponseWriter, r *http.Request){
    if r.Method == "POST" {
        //fmt.Println("Receive a POST Request")     
        r.ParseForm()
        
        //debug 
        if(Debug == 1){
            requestDump, err := httputil.DumpRequest(r, true)
            if err != nil {
              fmt.Println(err)
            }
            fmt.Println(string(requestDump))
       }
        
        
    }else if r.Method == "GET" {
        //fmt.Println("Receive a GET Request")
    }
    var path = r.URL.Path[1:]
    var rep = ""
    if(path == ""){
        rep = flash()
    }else if(path == "jdcheck.js"){
        rep = alive()
    }else if(path == "addcrypted2"){
        //fmt.Println(reflect.TypeOf(r.Form).String())
        addcrypted2(r)
        w.WriteHeader(200)        
    }else if(path == "flash/add"){
        add(r)
        w.WriteHeader(200)
    }else{
        fmt.Println("Unsupported path")
        w.WriteHeader(404)
        
    }
    fmt.Fprintf(w, rep)      
}

func flash() string{
    return "JDownloader"
}

func alive() string{
    return "jdownloader=true;\nvar version='10629';\n"    
}

func add(r *http.Request){
    // seperator of urls are \n
    urls := r.FormValue("urls")
    source := r.FormValue("source")
    passwords := r.FormValue("passwords")
    fmt.Printf("urls:\n%+v\n", urls)
    fmt.Printf("source:%+v\n",source)
    fmt.Printf("passwords:%+v\n",passwords)
}

func addcrypted2(r *http.Request){
    
    source := r.FormValue("source")
    passwords := r.FormValue("passwords")
    jk := r.FormValue("jk")
    crypted := r.FormValue("crypted")
    
    fmt.Printf("%+v\n", r.Form)
    fmt.Printf("source:%+v\n",source)
    fmt.Printf("passwords:%+v\n",passwords)
    fmt.Printf("jk:%+v\n",jk)
    fmt.Printf("crypted:%+v\n",crypted)
    fmt.Println("addcrypted2 not supported yet ;-)")
}

func main(){
    fmt.Println("Starting local CNL webserver on localhost:9666")
    http.HandleFunc("/", handler)    
    //~ http.HandleFunc("/jdcheck.js", alive)
    //~ http.HandleFunc("/addcrypted2", addcrypted2)   
    http.ListenAndServe(":9666",nil)
}
