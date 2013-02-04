package techism

import (
    "appengine"
    "appengine/urlfetch"
    "net/http"
    "hash/fnv"
    "io/ioutil"
    "fmt"   
    "strconv"
)

func get_body(url string, r *http.Request)(string, string) {
    c := appengine.NewContext(r)
    client := urlfetch.Client(c)
    resp, err := client.Get(url)
    if err != nil {
        fmt.Printf("key: %s", err)
		return "","ERROR"
    }
    body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("key: %s", err)
		return "","ERROR"
	}
	return string(body), ""
}

func calculate_checksum(body string) (string){
	fnv_sum := fnv.New64()
	fnv_sum.Write([]byte(body))
	checksum := fnv_sum.Sum64()
	return strconv.FormatUint(checksum, 16)
}
