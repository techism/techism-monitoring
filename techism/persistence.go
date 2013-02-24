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

func get_sites_with_status_error_or_ok (c appengine.Context)([]Site, []*datastore.Key, error){
    //TODO also filter Status = "ERROR"
    q := datastore.NewQuery("Site").Filter("Status =", "OK").Order("-Date").Limit(10)
    sites := make([]Site, 0, 10)
    keys, err := q.GetAll(c, &sites);
    return sites, keys, err
}

func get_all_sites (c appengine.Context)([]Site, []*datastore.Key, error){
    //TODO
	q := datastore.NewQuery("Site").Order("-Title").Limit(500)
    sites := make([]Site, 0, 500)
    keys, err := q.GetAll(c, &sites);
    return sites, keys, err
}

func save_new_site (site *Site, c appengine.Context)(*datastore.Key, error){
    key, err := datastore.Put(c, datastore.NewIncompleteKey(c, "Site", nil), site);
    return key, err
}

func update_site (k *datastore.Key, site *Site, c appengine.Context)(error){
    _, err := datastore.Put(c, k, site)
    return err
}
