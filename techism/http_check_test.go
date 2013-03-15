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

func TestGi (t *testing.T){
    orig := "<a href=\"mhw.cgi?page=oLyoiLexf3SlrXHBzJyUm8a4xoWHosDMqIlwt7vCrLjAOVi/wbHGrJaGioWMkUlwdg\"><img src=\"/gi/images/bild/muc3.jpg\" border=\"0\" width=\"750\" alt=\"Bild\" /></a>"
    stripped := "<a href=\"mhw.cgi?page=oLyoiLexf3SlrXHBzJyUm8a4xoWHosDMqIlwt7vCrLjAOVi/wbHGrJaGioWMkUlwdg\"></a>"
    orig = clean_up_body (orig)
    check1 := calculate_checksum (orig)
    check2 := calculate_checksum (stripped)
    assertEquals (check1, check2, t)
}

func TestArch1 (t *testing.T){
    orig := "<img src=\"/lib/exe/indexer.php?id=main&amp;1362946031\" width=\"2\" height=\"1\" alt=\"\" /></div>"
    stripped := "</div>"
    orig = clean_up_body (orig)
    check1 := calculate_checksum (orig)
    check2 := calculate_checksum (stripped)
    assertEquals (check1, check2, t)
}

func TestArch2 (t *testing.T){
    orig := "</a><a href=\"/main?do=login&amp;sectok=1e436611c7452c1fe857a77b889cfa86\"  class=\"action login\" rel=\"nofollow\" title=\"Anmelden\">Anmelden</a></div>"
    stripped := "</a><a href=\"/main?do=login&amp;\"  class=\"action login\" rel=\"nofollow\" title=\"Anmelden\">Anmelden</a></div>"
    orig = clean_up_body (orig)    
    check1 := calculate_checksum (orig)
    check2 := calculate_checksum (stripped)
    assertEquals (check1, check2, t)
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
    orig = clean_up_body (orig)
    
    check1 := calculate_checksum (orig)
    check2 := calculate_checksum (stripped)
    assertEquals (check1, check2, t)
}

func TestParameterJsessionid (t *testing.T){
    orig := "<li><a href=\"/contact.html;jsessionid=BD7979B9FDA4630569C84DA2FE7B662C\" style=\"text-decoration:none;width:36px;display:block;\">Kontakt</a></li>"
    stripped := "<li><a href=\"/contact.html;\" style=\"text-decoration:none;width:36px;display:block;\">Kontakt</a></li>"
    orig = clean_up_body (orig) 
    fmt.Println (orig)   
    check1 := calculate_checksum (orig)
    check2 := calculate_checksum (stripped)
    assertEquals (check1, check2, t)
}

func assertEquals (str1 string, str2 string, t *testing.T){
    if str1 != str2 {
        t.Fatalf("Test failed: " + str1 + " != " + str2)
    }
}


