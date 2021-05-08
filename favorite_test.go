package favorite

import (
	"reflect"
	"testing"
)

func TestFavs(t *testing.T) {
	a := Favs()
	if !reflect.DeepEqual(a, []string{"beer", "whiskey", "golf"}) {
		t.Errorf("got %v, want %v", "pony", a)
	}

}
