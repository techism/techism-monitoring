techism-monitoring
==================

Utility App for Techism to check the status of User Group Websites.

Prerequisites
-------------
* App Engine SDK for Go: https://developers.google.com/appengine/downloads
* PATH:


    export PATH=/home/ck/techism/google_appengine/:$PATH


Running the application
-----------------------

    dev_appserver.py techism-monitoring
    dev_appserver.py techism-monitoring/ --clear_datastore


Running Tests
--------------
* https://github.com/icub3d/appenginetesting can't be installed with GAE 1.7.5
* Brute Force Method
    * Download the full-fledged Go SDK: https://code.google.com/p/go/downloads/list
    * Copy files from goroot/src/cmd/cgo to the AppEngine SDK
    * Run: 
        * cd techism-monitoring/techism
        * go test http_check_test.go http_check.go persistence.go


Deployment
----------

    appcfg.py --oauth2 update techism-monitoring/
    

Weblinks
--------
* http://golang.org/pkg/
* http://tour.golang.org/
* https://developers.google.com/appengine/docs/go/
* https://developers.google.com/appengine/docs/go/datastore/reference

* http://talks.golang.org/2012/concurrency.slide#47
