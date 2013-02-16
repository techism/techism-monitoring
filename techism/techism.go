package techism

import (
    "appengine"
    "appengine/datastore"
    "text/template"
    "net/http"
    "time"
    "fmt"
)

var (
	statusTemplate = template.Must(template.ParseFiles("status.html"))
	)

func init() {
    http.HandleFunc("/", root)
    http.HandleFunc("/check", check_all)
    http.HandleFunc("/reset", reset)
    http.HandleFunc("/add", add)
}

//shows all sites from the database
func root(w http.ResponseWriter, r *http.Request) {
    c := appengine.NewContext(r)
    sites, _, err := get_all_sites(c)

    if err != nil {
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
	if err1 := r.ParseForm(); err1 != nil {
		http.Error(w, err1.Error(), http.StatusInternalServerError)
		return
	}
	title := r.FormValue("title")
	//TODO check size
    site, key, err2 := get_site_by_title (title, c)
	if err2 != nil {
		http.Error(w, err2.Error(), http.StatusInternalServerError)
		return
	}

    html_body, err3 := get_html_body(site.Url, r) 
    if err3 != "" {
		site.Status = "ERROR"
    } else {
    	site.Checksum = calculate_checksum (html_body)
    	site.Status = "OK"
	}
	site.Date = time.Now()
	datastore.Put(c, key, &site);
	root(w, r);
}

func add(w http.ResponseWriter, r *http.Request){
	//TODO
    title := r.FormValue("title")
    url := r.FormValue("url")

	c := appengine.NewContext(r)
    g := &Site{
		Title: title,
		Url:   url,
	}

    if _, err := datastore.Put(c, datastore.NewIncompleteKey(c, "Site", nil), g);
    err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        fmt.Println(err)
		return
	}
    
    //check_site_status(*g, r)
    datastore.Put(c, datastore.NewIncompleteKey(c, "Site", nil), g);
	http.Redirect(w, r, "/", http.StatusFound);
}
