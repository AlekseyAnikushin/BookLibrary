package handlers

import (
	"encoding/json"
	"net/http"
)

func writeResponse(w http.ResponseWriter, resp *response) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.Code)
	r := make(map[string]string)
	r["message"] = resp.Message
	if resp.Result != nil {
		r["result"] = ""
	}
	jsonResp, _ := json.Marshal(r)
	if resp.Result != nil {
		last := jsonResp[len(jsonResp)-1]
		jsonResp = append(jsonResp[0:len(jsonResp)-3], resp.Result...)
		jsonResp = append(jsonResp, last)
	}
	w.Write(jsonResp)
}
