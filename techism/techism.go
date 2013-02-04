package techism

import (
    "appengine"
    "appengine/datastore"
    "text/template"
    "net/http"
    "time"
)

type Site struct {
	Title    string
    Url      string
    Status   string
    Checksum string
    Date     time.Time
}

var (
	statusTemplate = template.Must(template.ParseFiles("status.html"))
	)

func init() {
    http.HandleFunc("/", root)
    http.HandleFunc("/check", check_all)
    http.HandleFunc("/reset", reset)
}

func root(w http.ResponseWriter, r *http.Request) {

    c := appengine.NewContext(r)
    q := datastore.NewQuery("Site").Order("-Date").Limit(10)
    sites := make([]Site, 0, 10)
    if _, err := q.GetAll(c, &sites); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    if err := statusTemplate.Execute(w, sites); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

func check_all(w http.ResponseWriter, r *http.Request){
	    c := appengine.NewContext(r)
	    //TODO also filter Status = "ERROR"
	    q := datastore.NewQuery("Site").Filter("Status =", "OK").Order("-Date").Limit(10)
	    sites := make([]Site, 0, 10)
	    keys, err := q.GetAll(c, &sites);
	    if err != nil {
	        http.Error(w, err.Error(), http.StatusInternalServerError)
	        return
	    }
	     
	    for index, value := range sites {
	    	body,err := get_body(value.Url, r)
	    	if err != "" {
	    		value.Status = "ERROR"
	    	} else {
	    		checksum := calculate_checksum(body)
	    		if checksum == value.Checksum {
	    			value.Status = "OK"
	    		} else {
	    			value.Status = "CHANGED"
	    		}
	    		value.Checksum = checksum;
	    	}
			value.Date = time.Now()
			if _, err := datastore.Put(c, keys[index], &value);
	    	err != nil {
	    		http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}
		if err := statusTemplate.Execute(w, sites); err != nil {
	        http.Error(w, err.Error(), http.StatusInternalServerError)
	    }
	}

func reset(w http.ResponseWriter, r *http.Request){
	//load site
	//update checksum
	//update date
	//set status to OK
	root(w, r);
}