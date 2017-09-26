
package main
import (
    "fmt"
    "net/http"
    "net/http/httputil"
    "encoding/base64"
    "encoding/hex"
    "regexp"
    "crypto/aes"
    "crypto/cipher"
    "errors"
)

var Debug = 0

func handler(w http.ResponseWriter, r *http.Request){
    if r.Method == "POST" {
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
    method := "add plain"
    urls := r.FormValue("urls")
    source := r.FormValue("source")
    passwords := r.FormValue("passwords")
    fmt.Printf("Methode:%v\n",method)
    fmt.Printf("urls:\n%+v\n", urls)
    fmt.Printf("source:%+v\n",source)
    fmt.Printf("passwords:%+v\n",passwords)
}

func aes_decrypt(key string,data string) (decoded string){
  src :=[]byte(key)
  //encode the key
  dst := make([]byte, hex.DecodedLen(len(src)))
  n, err := hex.Decode(dst, src)
  if err != nil {
    fmt.Println(err)
  }
  //decode the data
  crypted, err := base64.StdEncoding.DecodeString(data)
    if err != nil {
      fmt.Println("decode error:", err)
      return
  }
  cipher_key := dst[:n]
  block, err := aes.NewCipher(cipher_key)
	if err != nil {
		return
  }
  if len(crypted) < aes.BlockSize {
		err = errors.New("Ciphertext block size is too short!")
		return
  }
  mode := cipher.NewCBCDecrypter(block, cipher_key)
  mode.CryptBlocks(crypted, crypted)
  return string(crypted)
}

func addcrypted2(r *http.Request){
    method := "addcryptd2"
    package_name :=r.FormValue("package")
    source := r.FormValue("source")
    passwords := r.FormValue("passwords")
    jk := r.FormValue("jk")
    crypted := r.FormValue("crypted")
    re, _ := regexp.Compile(`\{.?return.?'(\d+)'.?\}`)
    res := re.FindStringSubmatch(jk)
    var key string
    if res != nil{
      key = res[1]
    }else{
      fmt.Println("No Key found")
      return
    }
    urls := aes_decrypt(key,crypted)

    fmt.Printf("Methode:%v\n",method)
    if len(package_name) > 0 {
      fmt.Printf("package:%+v\n",package_name)
    }
    fmt.Printf("source:%+v\n",source)
    fmt.Printf("passwords:%+v\n",passwords)
    fmt.Printf("urls:\n%v\n",urls)
}

func main(){
    fmt.Println("Starting local CNL webserver on localhost:9666")
    http.HandleFunc("/", handler)
    //~ http.HandleFunc("/jdcheck.js", alive)
    //~ http.HandleFunc("/addcrypted2", addcrypted2)
    http.ListenAndServe(":9666",nil)
}
