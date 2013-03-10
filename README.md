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


Deployment
----------

    appcfg.py --oauth2 update techism-monitoring/
    

Weblinks
--------
* http://golang.org/pkg/
* http://tour.golang.org/
* https://developers.google.com/appengine/docs/go/
* https://developers.google.com/appengine/docs/go/datastore/reference
