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

func init() {
    http.HandleFunc("/", root)
    http.HandleFunc("/check", check_all)
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
    q := datastore.NewQuery("Site").Filter("Status =", "OK").Order("-Date").Limit(10)
    sites := make([]Site, 0, 10)
    keys, err := q.GetAll(c, &sites);
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
     
    for index, value := range sites {
    	//TODO
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

var (
	statusTemplate = template.Must(template.ParseFiles("status.html"))
	)