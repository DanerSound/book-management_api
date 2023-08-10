package utils

import(
	"encoding/json"
	"io/ioutil"
	"net/http"
)
// i will recive a body structered in json and i need to parse it 
func ParseBody(r * http.Request, x interface{}){
	if body, err := ioutil.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal([]byte(body), x); err != nil{
			return
		}
	}
}