package techism

import (
    "testing"
    "fmt"
)

const plainText = "This is a text without any tags"
const aHrefText = "<p> 7. March 2013 <br> @ <a href=\"http://www.munichjs.org\">MunichJS</a> Coworking </p>"

func TestRemoveImages(t *testing.T) {
    neu1 := remove_images(plainText)
    assertEquals (plainText, neu1, t)

    neu2 := remove_images(aHrefText)
    assertEquals (aHrefText, neu2, t)
}


func TestRemoveComments(t *testing.T) {
    neu1 := remove_comments(plainText)
    assertEquals (plainText, neu1, t)

    neu2 := remove_comments (aHrefText)
    assertEquals (aHrefText, neu2, t)
}


func TestRemoveHiddenFields(t *testing.T) {
    neu1 := remove_hidden_fields(plainText)
    assertEquals (plainText, neu1, t)

    neu2 := remove_hidden_fields (aHrefText)
    assertEquals (aHrefText, neu2, t)
}




func TestRemoveMetaFields(t *testing.T) {
    neu1 := remove_meta_fields(plainText)
    assertEquals (plainText, neu1, t)

    neu2 := remove_meta_fields (aHrefText)
    assertEquals (aHrefText, neu2, t)
}


func TestSecTok (t *testing.T){
    orig := "<div class=\"no\"><input type=\"hidden\" name=\"do\" /><input type=\"hidden\" name=\"sectok\" value=\"1234\" />"
    stripped := "<div class=\"no\">"
    check_body (orig, stripped, t)
}


func TestImg (t *testing.T){
    orig := "<a href=\"mhw.cgi?page=12345\"><img src=\"/gi/images/bild/muc3.jpg\" border=\"0\" width=\"750\" alt=\"Bild\" /></a>"
    stripped := "<a href=\"mhw.cgi?page=12345\"></a>"
    check_body (orig, stripped, t)
}


func TestArch1 (t *testing.T){
    orig := "<img src=\"/lib/exe/indexer.php?id=main&amp;1362946031\" width=\"2\" height=\"1\" alt=\"\" /></div>"
    stripped := "</div>"
    check_body (orig, stripped, t)
}


func TestArch2 (t *testing.T){
    orig := "</a><a href=\"/main?do=login&amp;sectok=12345\"  class=\"action login\" rel=\"nofollow\" title=\"Anmelden\">Anmelden</a></div>"
    stripped := "</a><a href=\"/main?do=login\"  class=\"action login\" rel=\"nofollow\" title=\"Anmelden\">Anmelden</a></div>"
    check_body (orig, stripped, t)
}


func TestIFrame (t *testing.T){
    orig := `<iframe
                            id="janrainFrame"
                            src="http://social-community/openid/embed"
                            scrolling="no"
                            frameBorder="no"
                            allowtransparency="false">
              </iframe>`
    stripped := ""
    check_body (orig, stripped, t)
}


func TestParameterJsessionid (t *testing.T){
    orig := "<li><a href=\"/contact.html;jsessionid=12345\" style=\"text-decoration:none;width:36px;display:block;\">Kontakt</a></li>"
    stripped := "<li><a href=\"/contact.html;\" style=\"text-decoration:none;width:36px;display:block;\">Kontakt</a></li>"
    check_body (orig, stripped, t)
}


func TestMeetup (t *testing.T){
    orig := "<p>Thursday, May 16 at 7:00 PM</p> <p>Attending: 13</p> <p>Details: http://www.meetup.com/Udacity-Coursera-Stammtisch/events/114948122/</p>"
    stripped := "<p>Thursday, May 16 at 7:00 PM</p>  <p>Details: http://www.meetup.com/Udacity-Coursera-Stammtisch/events/114948122/</p>"
    check_body_url (orig, stripped, "http://meetup.com", t)
}

func TestSoftwarearchitektur (t *testing.T){
    orig := `<div id="cloud">
    <a href="/tag/agiler_prozess?do=showtag&amp;tag=agiler_prozess" class="cloud3_tag2" title="agiler_prozess">agiler_prozess</a>
    <a href="/tag/anekdote?do=showtag&amp;tag=anekdote" class="cloud1_tag2" title="anekdote">anekdote</a>
    <a href="/tag/antipattern?do=showtag&amp;tag=antipattern" class="cloud1_tag2" title="antipattern">antipattern</a>
    <a href="/tag/arc42?do=showtag&amp;tag=arc42" class="cloud4_tag2" title="arc42">arc42</a>
    <a href="/tag/async?do=showtag&amp;tag=async" class="cloud3_tag2" title="async">async</a>
    <a href="/tag/balsamiq?do=showtag&amp;tag=balsamiq" class="cloud1_tag2" title="balsamiq">balsamiq</a>
    <a href="/tag/beanstalkd?do=showtag&amp;tag=beanstalkd" class="cloud1_tag2" title="beanstalkd">beanstalkd</a>
    </div><img src="/lib/exe/indexer.php?id=main&amp;1366729201" width="2" height="1" alt="" /><div>test</div>`
    stripped := "<div>test</div>"
    check_body_url (orig, stripped, "http://www.workshop-softwarearchitektur.de", t)
}

func TestMactreff (t *testing.T){
    orig := `<div id="sidebar">
        <ul>
        <li class="sidebox"><h2>Tags</h2><div style="overflow:hidden"><a href="http://mactreffmuenchen.wordpress.com/tag/adobe/" class="tag-link-29691" title="6 Artikel" style="font-size: 12.504347826087pt;">Adobe</a>
        <a href="http://mactreffmuenchen.wordpress.com/tag/workshop/" class="tag-link-19557" title="4 Artikel" style="font-size: 10.678260869565pt;">Workshop</a></div></li></ul>
        </div><div>test</div>`
    stripped := "<div>test</div>"
    check_body_url (orig, stripped, "http://mactreffmuenchen.wordpress.com", t)
}


func TestComments1 (t *testing.T){
     orig := `<div class="wrapper">
      <div id="topbar">
        <div class="fl_left">
            <!--
            <span style="color:#222">News:</span>
            <a href="/article/next_meeting_1_december.ejs">Upcoming meeting: 1. December 2011 </a>
            -->
        </div>`
    stripped := `<div class="wrapper">
      <div id="topbar">
        <div class="fl_left">
            
        </div>`
    check_body (orig, stripped, t) 
}


func TestComments2 (t *testing.T){
    orig := `<!-- 30 queries. 0.351 seconds. -->`
    stripped := ``
    check_body (orig, stripped, t)
}

func check_body (orig string, stripped string, t *testing.T){
    check_body_url (orig, stripped, "http://www.example.com", t) 
}

func check_body_url (orig string, stripped string, url string, t *testing.T){
    orig = clean_up_body (orig, url) 
    check1 := calculate_checksum (orig)
    check2 := calculate_checksum (stripped)
    assertEqualsExt (check1, check2, orig, stripped, t)
}


func assertEqualsExt (str1 string, str2 string, cleaned string, stripped string, t *testing.T){
    if str1 != str2 {
        fmt.Printf("cleaned: '%s' \n", cleaned)
        fmt.Printf("stripped: '%s' \n", stripped)
        t.Fatalf("Test failed: " + str1 + " != " + str2)
    }
}


func assertEquals (str1 string, str2 string, t *testing.T){
    if str1 != str2 {
        t.Fatalf("Test failed: " + str1 + " != " + str2)
    }
}