package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home!")
}

func readBody(w http.ResponseWriter, r *http.Request) {
	body, _ := io.ReadAll(r.Body)
	fmt.Fprintf(w, "Body: %s\n", string(body))
	body, _ = io.ReadAll(r.Body)
	/*请求体只能读取一次*/
	fmt.Fprintf(w, "Read body one more time: [%s] data length %d \n", string(body), len(body))
}

func getBody(w http.ResponseWriter, r *http.Request) {
	/*原生获取请求体的方法返回值为空*/
	if r.GetBody == nil {
		fmt.Fprintf(w, "Body from original getBody method is nil\n")
	}
}

func query(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	fmt.Fprintf(w, "Query is %v\n", values)
}

func url(w http.ResponseWriter, r *http.Request) {
	data, _ := json.Marshal(r.URL)
	fmt.Fprintf(w, "Url is %v\n", string(data))
}

func header(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Head is %v\n", r.Header)
}

func form(w http.ResponseWriter, r *http.Request) {
	/*使用表单之前需要调用该方法*/
	fmt.Fprintf(w, "Before ParseForm %v\n", r.Form)
	r.ParseForm()
	fmt.Fprintf(w, "After ParseForm %v\n", r.Form)
}

func multiForm(w http.ResponseWriter, r *http.Request) {
	/*使用表单之前需要调用该方法*/
	fmt.Fprintf(w, "Before ParseForm %v\n", r.Form)
	r.ParseMultipartForm(100)
	fmt.Fprintf(w, "After ParseForm %v\n", r.Form)
}

func oldMain() {
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
