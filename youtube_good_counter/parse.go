package youtube_good_counter

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
)



func err_io(err error){
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s", err)
	}
}


func Json_parse(r *http.Request) map[string]string{
	length, err := strconv.Atoi(r.Header.Get("Content-Length"))
	err_io(err)

	body := make([]byte, length)
	length, _ = r.Body.Read(body)

	var jsonBody map[string]string
	err = json.Unmarshal(body[:length], &jsonBody)
	err_io(err)

	return jsonBody
}
