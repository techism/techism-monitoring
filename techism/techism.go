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

var (
	statusTemplate = template.Must(template.ParseFiles("status.html"))
	)