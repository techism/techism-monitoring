package techism

import (
    "appengine"
    "appengine/user"
    "text/template"
    "net/http"
    "time"
    "fmt"
)

var (
	statusTemplate = template.Must(template.New("status.html").Funcs(fns).ParseFiles("status.html"))
	)

type appHandler func (c appengine.Context, w http.ResponseWriter, r *http.Request) error

func init() {
    http.Handle("/", appHandler(root))
    http.Handle("/check", appHandler(check_all))
    http.Handle("/reset", appHandler(reset))
    http.Handle("/add", appHandler(add))
}

func (fn appHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    c := appengine.NewContext(r)
    u := user.Current(c)
    if u == nil {
        url, _ := user.LoginURL(c, "/")
        fmt.Fprintf(w, `<a href="%s">Sign in</a>`, url)
        return
    }

    if u.Admin {
        w.Header().Set("Content-Security-Policy", "default-src 'self'")
        if err := fn(c, w, r); err != nil {
            fmt.Println(err)
            http.Error(w, err.Error(), http.StatusInternalServerError)
        }
        
    }
}

//shows all sites from the database
func root(c appengine.Context, w http.ResponseWriter, r *http.Request) error {
    sites, _, err := get_all_sites(c)
    if err != nil {
        return err
    }
    if err := statusTemplate.Execute(w, sites); err != nil {
        return err
    }
    return nil
}

//checks a sites with status = OK
func check_all(c appengine.Context, w http.ResponseWriter, r *http.Request) error {
    sites, keys, err := get_sites_with_status_error_or_ok(c)
        
    if err != nil {
        return err
    }
    for index, site := range sites {
      	check_site_status(&site, r)
        
        key := keys[index]
        err := update_site (key, &site, c)
      	if err != nil {
       		return err
    	}
    }
    return nil
}

//resets status and checksum of a given site
func reset(c appengine.Context, w http.ResponseWriter, r *http.Request) error {
   	if err1 := r.ParseForm(); err1 != nil {
   		return err1
   	}
   	title := r.FormValue("title")
    site, key, err2 := get_site_by_title (title, c)
   	if err2 != nil {
   		return err2
   	}

    html_body, err3 := get_html_body(site.Url, r) 
    if err3 != "" {
   		site.Status = "ERROR"
    } else {
        cleaned_up_body := clean_up_body (html_body)
       	site.Checksum = calculate_checksum (cleaned_up_body)
       	site.Status = "OK"
  	}
   	site.Date = time.Now()
    update_site (key, &site, c )
   	http.Redirect(w, r, "/", http.StatusFound);
    return nil
}

//add new site
func add(c appengine.Context, w http.ResponseWriter, r *http.Request) error {
    title := r.FormValue("title")
    url := r.FormValue("url")
    site := &Site{
  		Title: title,
   		Url:   url,
   	}
    check_site_status(site, r)
    if _, err := save_new_site(site, c)
    err != nil {
   		return err
   	}
   	http.Redirect(w, r, "/", http.StatusFound);
    return nil
}

var fns = template.FuncMap{
	"odd": func(x int) bool {
		return x % 2 == 1
	},
	"even": func(x int) bool {
		return x % 2 == 0
	},
}
