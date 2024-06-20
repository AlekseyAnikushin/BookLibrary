package handlers

import (
	"encoding/json"
	"net/http"
)

func writeResponse(w http.ResponseWriter, message string, httpStatusCode int, result []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatusCode)
	resp := make(map[string]string)
	resp["message"] = message
	if result != nil {
		resp["result"] = ""
	}
	jsonResp, _ := json.Marshal(resp)
	if result != nil {
		last := jsonResp[len(jsonResp)-1]
		jsonResp = append(jsonResp[0:len(jsonResp)-3], result...)
		jsonResp = append(jsonResp, last)
	}
	w.Write(jsonResp)
}
