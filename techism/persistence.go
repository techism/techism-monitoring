package techism

import (
	"appengine"
    "appengine/datastore"
    "time"
)

type Site struct {
	Title    string
    Url      string
    Status   string
    Checksum string
    Date     time.Time
}

func get_site_by_title (title string, c appengine.Context)(Site, *datastore.Key, error){
	q := datastore.NewQuery("Site").Filter("Title =", title)
	sites := make([]Site, 0, 1)
    keys, err2 := q.GetAll(c, &sites);
    site := sites[0]
    if err2 != nil {
        return site, nil, err2
    }
    key := keys[0]
    return site, key, nil
}
