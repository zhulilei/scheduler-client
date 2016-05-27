package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {

	http.HandleFunc("/job", jobExecute)
	err := http.ListenAndServe(":8010", nil)
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

type JobResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Content string `json:"content"`
}
