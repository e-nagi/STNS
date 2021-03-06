package stns

import (
	"reflect"
	"testing"
)

func TestGetByName(t *testing.T) {
	users := Attributes{
		"test1": &Attribute{
			ID: 1,
			User: &User{LinkUsers: []string{"foo", "bar"},
				Password: "foo",
			},
			Group: &Group{Users: []string{"foo", "bar"}},
		},
		"test3": &Attribute{},
		"test4": &Attribute{
			ID: 4,
		},
	}

	_users := users.GetByName("test1")

	for n, u := range _users {
		if u.ID != 1 {
			t.Error("ummatch user id")
		}

		if !reflect.DeepEqual(u.LinkUsers, []string{"foo", "bar"}) {
			t.Error("ummatch link user")
		}

		if !reflect.DeepEqual(u.Users, []string{"foo", "bar"}) {
			t.Error("ummatch link user")
		}

		if u.Password != "foo" {
			t.Error("ummatch password")
		}

		if n != "test1" {
			t.Error("ummatch user name")
		}
	}
	notfound := users.GetByName("test2")
	if notfound != nil {
		t.Error("ummatch user id")
	}

	t3 := users.GetByName("test3")

	if t3 != nil {
		t.Error("ummatch user id test3")
	}

	t4 := users.GetByName("test4")
	if t4 == nil {
		t.Error("ummatch user id test4")
	}
}

func TestGetByID(t *testing.T) {
	users := Attributes{
		"test1": &Attribute{
			ID: 1,
		},
		"test3": &Attribute{},
		"test4": &Attribute{
			ID: 4,
		},
	}
	_users := users.GetByID("1")
	for n, u := range _users {
		if u.ID != 1 {
			t.Error("ummatch user id")
		}

		if n != "test1" {
			t.Error("ummatch user name")
		}
	}

	notfound := users.GetByID("2")
	if notfound != nil {
		t.Error("ummatch user id")
	}

	t3 := users.GetByID("3")

	if t3 != nil {
		t.Error("ummatch user id test3")
	}

	t4 := users.GetByID("4")
	if t4 == nil {
		t.Error("ummatch user id test4")
	}
}
