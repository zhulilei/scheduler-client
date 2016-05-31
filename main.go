package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
	"os"
)

func main() {

	bind := fmt.Sprintf("%s:%s", os.Getenv("OPENSHIFT_GO_IP"), os.Getenv("OPENSHIFT_GO_PORT"))
	fmt.Printf("listening on %s...", bind)
	http.HandleFunc("/job", jobExecute)
	http.HandleFunc("/", Index)
	err := http.ListenAndServe(bind, nil)
	if err != nil {
		log.Fatalln(err)
		return
	}
}

func jobExecute(response http.ResponseWriter, request *http.Request) {
	request.ParseForm()

	body, err := ioutil.ReadAll(request.Body)
	result := &JobResponse{}
	if err != nil {
		result.Message = "失败"
	} else {
		fmt.Println(string(body))
		result.Success = true
		result.Message = "执行成功"
		time.Sleep(5 * time.Second)
		result.Content = "执行成功"
	}

	bd, err := json.Marshal(result)
	response.Write(bd)
	defer request.Body.Close()
	fmt.Println(request.Form)

}
func Index(response http.ResponseWriter, request *http.Request) {

	response.Write([]byte("hello golang!"))

}
type JobResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Content string `json:"content"`
}
