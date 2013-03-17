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
    "regexp"
    "strings"
)


func check_site_status(value *Site, r *http.Request){
    body,err := get_html_body(value.Url, r)
    if err != "" {
        value.Status = "ERROR"
    } else {
        fmt.Println(value.Url)
        cleaned_up_body := clean_up_body (body, value.Url)
        checksum := calculate_checksum(cleaned_up_body)
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


func clean_up_body (body string, url string) (string){
    body = remove_meta_fields (body)
    body = remove_hidden_fields (body)

    body = remove_comments (body)
    body = remove_images (body)
    body = remove_sessionids_and_csrftoken(body)
    body = remove_iframes (body)
    //specific checks
    if(strings.Contains(url, "it-szene")){
        body = remove_chars_itszene (body)
    } else if(strings.Contains(url, "freifunk")){
        body = remove_chars_freifunk (body)
    } else if(strings.Contains(url, "owasp") || strings.Contains(url, "ottobrunn")){
        body = remove_chars_owasp (body)
    } 
    return body
}


func calculate_checksum(body string) (string){
	fnv_sum := fnv.New64()
	fnv_sum.Write([]byte(body))
	checksum := fnv_sum.Sum64()
	return strconv.FormatUint(checksum, 16)
}

// ----------------
//TODO replace with exp/html as soon as it's bundled with appengine
//-----------------

func remove_meta_fields (body string) (string){    
    regex, _ := regexp.Compile("<meta .*?>")
    result := regex.ReplaceAllString(body, "")
    return result
}


func remove_hidden_fields (body string) (string){
    regex, _ := regexp.Compile("<input type=\"hidden\".*?>")
    result := regex.ReplaceAllString(body, "")
    return result
}


func remove_comments (body string) (string){
    regex, _ := regexp.Compile("(?s)<!--.*?-->")
    result := regex.ReplaceAllString(body, "")
    return result
}


func remove_iframes (body string) (string){
    regex, _ := regexp.Compile("(?s)<iframe.*?iframe>")
    result := regex.ReplaceAllString(body, "")
    return result
}


func remove_images (body string) (string){
    regex, _ := regexp.Compile("<img .*?/>")
    result := regex.ReplaceAllString(body, "")
    return result
}


func remove_sessionids_and_csrftoken (body string) (string){
    regex, _ := regexp.Compile("sectok=[a-fA-F0-9]*")
    result := regex.ReplaceAllString(body, "")

    regex2, _ := regexp.Compile("sid=[a-fA-F0-9]*")
    result2 := regex2.ReplaceAllString(result, "")

    regex3, _ := regexp.Compile("jsessionid=[a-fA-F0-9]*")
    result3 := regex3.ReplaceAllString(result2, "")

    regex4, _ := regexp.Compile("&amp;[0-9]*")
    result4 := regex4.ReplaceAllString(result3, "")

    return result4
}


func remove_chars_itszene (body string) (string){
    regex, _ := regexp.Compile("MttgSession[\\s]?=[\\s]?[a-fA-F0-9]*")
    result := regex.ReplaceAllString(body, "")
    return result
}


func remove_chars_freifunk (body string) (string){
    regex, _ := regexp.Compile("<li id=\"viewcount\">.*?</li>")
    result := regex.ReplaceAllString(body, "")
    return result
}


func remove_chars_owasp (body string) (string){
    regex, _ := regexp.Compile("<li id=\"footer-info-viewcount\">.*?</li>")
    result := regex.ReplaceAllString(body, "")
    return result
}
