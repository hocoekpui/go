package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Home!")
	if err != nil {
		return
	}
}

func readBody(w http.ResponseWriter, r *http.Request) {
	body, _ := io.ReadAll(r.Body)
	_, err := fmt.Fprintf(w, "Body: %s\n", string(body))
	if err != nil {
		return
	}
	body, _ = io.ReadAll(r.Body)
	/*请求体只能读取一次*/
	_, err = fmt.Fprintf(w, "Read body one more time: [%s] data length %d \n", string(body), len(body))
	if err != nil {
		return
	}
}

func getBody(w http.ResponseWriter, r *http.Request) {
	/*原生获取请求体的方法返回值为空*/
	if r.GetBody == nil {
		_, err := fmt.Fprintf(w, "Body from original getBody method is nil\n")
		if err != nil {
			return
		}
	}
}

func query(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	_, err := fmt.Fprintf(w, "Query is %v\n", values)
	if err != nil {
		return
	}
}

func url(w http.ResponseWriter, r *http.Request) {
	data, _ := json.Marshal(r.URL)
	_, err := fmt.Fprintf(w, "Url is %v\n", string(data))
	if err != nil {
		return
	}
}

func header(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Head is %v\n", r.Header)
	if err != nil {
		return
	}
}

func form(w http.ResponseWriter, r *http.Request) {
	/*使用表单之前需要调用该方法*/
	_, err := fmt.Fprintf(w, "Before ParseForm %v\n", r.Form)
	if err != nil {
		return
	}
	err = r.ParseForm()
	if err != nil {
		return
	}
	_, err = fmt.Fprintf(w, "After ParseForm %v\n", r.Form)
	if err != nil {
		return
	}
}

func multiForm(w http.ResponseWriter, r *http.Request) {
	/*使用表单之前需要调用该方法*/
	_, err := fmt.Fprintf(w, "Before ParseForm %v\n", r.Form)
	if err != nil {
		return
	}
	err = r.ParseMultipartForm(100)
	if err != nil {
		return
	}
	_, err = fmt.Fprintf(w, "After ParseForm %v\n", r.Form)
	if err != nil {
		return
	}
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/readBodyOnce", readBody)
	http.HandleFunc("/getBody", getBody)
	http.HandleFunc("/query", query)
	http.HandleFunc("/url", url)
	http.HandleFunc("/header", header)
	http.HandleFunc("/form", form)
	http.HandleFunc("/multiForm", multiForm)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
