package techism

import (
    "appengine"
    "appengine/urlfetch"
    "net/http"
    "hash/fnv"
    "io/ioutil"
    "fmt"   
    "strconv"
    "time"
)

func check_site_status(value Site, r *http.Request){
    body,err := get_html_body(value.Url, r)
    if err != "" {
        value.Status = "ERROR"
    } else {
        checksum := calculate_checksum(body)
        if value.Checksum == "" {
            value.Checksum = checksum;
            value.Status = "OK"
        } else { 
            if checksum == value.Checksum {
                value.Status = "OK"
            } else {
                value.Status = "CHANGED"
            }
        }
    }
    value.Date = time.Now()
}


func get_html_body(url string, r *http.Request)(string, string) {
    c := appengine.NewContext(r)
    client := urlfetch.Client(c)
    //TODO check HTTP Status code
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