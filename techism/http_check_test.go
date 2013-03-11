package techism

import (
    "testing"
)

const plainText = "This is a text without any tags"
const aHrefText = "<p> 7. March 2013 <br> @ <a href=\"http://www.munichjs.org\">MunichJS</a> Coworking </p>"

func TestRemoveImages(t *testing.T) {
    neu1 := remove_images(plainText)
    assertTrue (plainText, neu1, t)

    neu2 := remove_images(aHrefText)
    assertTrue (aHrefText, neu2, t)
}


func TestRemoveComments(t *testing.T) {
    neu1 := remove_comments(plainText)
    assertTrue (plainText, neu1, t)

     neu2 := remove_comments (aHrefText)
    assertTrue (aHrefText, neu2, t)
}


func TestRemoveHiddenFields(t *testing.T) {
    neu1 := remove_hidden_fields(plainText)
    assertTrue (plainText, neu1, t)

    neu2 := remove_hidden_fields (aHrefText)
    assertTrue (aHrefText, neu2, t)
}


func TestRemoveMetaFields(t *testing.T) {
    neu1 := remove_meta_fields(plainText)
    assertTrue (plainText, neu1, t)

    neu2 := remove_meta_fields (aHrefText)
    assertTrue (aHrefText, neu2, t)
}

func assertTrue (str1 string, str2 string, t *testing.T){
    if str1 != str2 {
        t.Fatalf("Test failed: " + str1 + " != " + str2)
    }
}


