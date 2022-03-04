package models

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	Status      int         `json:"status"`
	Data        interface{} `json:"data"`
	Message     string      `json:"message"`
	contentType string
	writer      http.ResponseWriter
}

func CreateDefaultResponse(w http.ResponseWriter) *Response {
	return &Response{
		Status:      http.StatusOK,
		writer:      w,
		contentType: "application/json",
	}
}

func SendNotFound(w http.ResponseWriter) {
	response := CreateDefaultResponse(w)
	response.NotFound()
	response.Send()
}

//NotFound is a helper function to send a 404 response
func (response *Response) NotFound() {
	response.Status = http.StatusNotFound
	response.Message = "Resource Not Found"
}

func SendUnprocessableEntity(w http.ResponseWriter) {
	response := CreateDefaultResponse(w)
	response.UnprocessableEntity()
	response.Send()
}

func (response *Response) UnprocessableEntity() {
	response.Status = http.StatusUnprocessableEntity
	response.Message = "Unprocessable Entity"
}

func SendNoContent(w http.ResponseWriter) {
	response := CreateDefaultResponse(w)
	response.NoContent()
	response.Send()
}

func (response *Response) NoContent() {
	response.Status = http.StatusNoContent
	response.Message = "No Content"
}

func SendData(w http.ResponseWriter, data interface{}) {
	response := CreateDefaultResponse(w)
	response.Data = data
	response.Send()
}

func (response *Response) Send() {
	response.writer.Header().Set("Content-Type", response.contentType)
	response.writer.WriteHeader(response.Status)
	response.writer.Header().Set("Access-Control-Allow-Origin", "*")
	response.writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	output, _ := json.Marshal(&response)
	fmt.Fprintf(response.writer, "%s", output)

}
