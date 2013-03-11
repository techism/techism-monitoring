package techism

import (
    "testing"
)

const text = "This is a text without any tags"

func TestRemoveNoImages(t *testing.T) {
    neu := remove_images(text)
    if text != neu {
        t.Fatalf("Remove images failed")
    }
}


func TestRemoveNoComments(t *testing.T) {
    neu := remove_comments(text)
    if text != neu {
        t.Fatalf("Remove comments failed")
    }
}


func TestRemoveNoHiddenFields(t *testing.T) {
    neu := remove_hidden_fields(text)
    if text != neu {
        t.Fatalf("Remove hidden fields failed")
    }
}


func TestRemoveNoMetaFields(t *testing.T) {
    neu := remove_meta_fields(text)
    if text != neu {
        t.Fatalf("Remove meta fields failed")
    }
}


