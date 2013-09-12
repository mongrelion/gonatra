package gonatra

import (
    "testing"
)

func TestGetSessionKey(t *testing.T) {
    var v string
    // Test it retrieves blank values for unset keys.
    v = GetSessionKey("foo")
    if v != "" {
        t.Errorf(`expected key "foo" to be empty but got %s`, v)
    }

    // Test it retrieves the proper values for set keys.
    session.m["lol"] = "lmao"
    v = GetSessionKey("lol")
    if v != "lmao" {
        t.Errorf(`expected key "lol" to hold "lmao" but got %s`, v)
    }
}

func TestSetSessionKey(t *testing.T) {
    key, value := "lolcat", "icanhazburger"
    SetSessionKey(key, value)

    // Test the variable is set in session var.
    val, ok := session.m[key]
    if !ok {
        t.Errorf("expected key \"%s\" to be defined in session map.", key)
    }

    // Test the key value is properly set.
    if val != value {
        t.Errorf("expected key \"%s\" key to have value \"%s\" but got %s", key, value, val)
    }
}

