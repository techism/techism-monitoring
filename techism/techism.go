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

//shows all sites from the database
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
        return
    }
}

//checks a sites with status = OK
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
    for index, site := range sites {
    	check_site_status(site, r)
    	_, err := datastore.Put(c, keys[index], &site);
    	if err != nil {
    		http.Error(w, err.Error(), http.StatusInternalServerError)
    		return
		}
	}
	if err := statusTemplate.Execute(w, sites); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
}

//resets status and checksum of a given site
func reset(w http.ResponseWriter, r *http.Request){
	c := appengine.NewContext(r)
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//TODO replace with request parameter
	title := r.FormValue("title")
	q := datastore.NewQuery("Site").Filter("Title =", title)
	sites := make([]Site, 0, 1)
    keys, err := q.GetAll(c, &sites);
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    site := sites[0]
    html_body, err = get_html_body(site.Url, r) 
    if err != "" {
		site.Status = "ERROR"
    } else
    	site.Checksum = calculate_checksum (html_body)
    	site.Status = "OK"
	}
	site.Date = time.Now()
	root(w, r);
}