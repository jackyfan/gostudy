package testhttp

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
	"strings"
)

const http_root = "/home/haoel/coolshell.cn/"

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	fmt.Println("start")
	r.ParseForm()               //解析参数，默认是不会解析的
	fmt.Println("Form", r.Form) //这些信息是输出到服务器端的打印信息
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello astaxie!") //这个写入到w的是输出到客户端的
}

//读取一些HTTP的头
func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "rootHandler: %s\n", r.URL.Path)
	fmt.Fprintf(w, "URL: %s\n", r.URL)
	fmt.Fprintf(w, "Method: %s\n", r.Method)
	fmt.Fprintf(w, "RequestURI: %s\n", r.RequestURI)
	fmt.Fprintf(w, "Proto: %s\n", r.Proto)
	fmt.Fprintf(w, "HOST: %s\n", r.Host)
}

//特别的URL处理
func viewHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "viewHandler: %s", r.URL.Path)
}

//一个静态网页的服务示例。（在http_root的html目录下）
func htmlHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("htmlHandler: %s\n", r.URL.Path)

	filename := http_root + r.URL.Path
	fileext := filepath.Ext(filename)

	content, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("   404 Not Found!\n")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	var contype string
	switch fileext {
	case ".html", "htm":
		contype = "text/html"
	case ".css":
		contype = "text/css"
	case ".js":
		contype = "application/javascript"
	case ".png":
		contype = "image/png"
	case ".jpg", ".jpeg":
		contype = "image/jpeg"
	case ".gif":
		contype = "image/gif"
	default:
		contype = "text/plain"
	}
	fmt.Printf("ext %s, ct = %s\n", fileext, contype)

	w.Header().Set("Content-Type", contype)
	fmt.Fprintf(w, "%s", content)

}
func Testhttp() {
	http.HandleFunc("/say", sayhelloName) //设置访问的路由
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/html/", htmlHandler)
	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
