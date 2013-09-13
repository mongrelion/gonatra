package gonatra

import (
    "testing"
)

func TestSessionGet(t *testing.T) {
    var v string
    session := session{m: make(map[string]string)}
    // Test it retrieves blank values for unset keys.
    v = session.Get("foo")
    if v != "" {
        t.Errorf(`expected key "foo" to be empty but got %s`, v)
    }

    // Test it retrieves the proper values for set keys.
    session.m["lol"] = "lmao"
    v = session.Get("lol")
    if v != "lmao" {
        t.Errorf(`expected key "lol" to hold "lmao" but got %s`, v)
    }
}

func TestSessionSet(t *testing.T) {
    session := session{m: make(map[string]string)}
    key, value := "lolcat", "icanhazburger"
    session.Set(key, value)

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
