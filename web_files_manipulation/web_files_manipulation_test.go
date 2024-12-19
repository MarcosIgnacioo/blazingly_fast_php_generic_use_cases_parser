package web_files_manipulation

import (
	"os"
	"testing"

	"github.com/sunshineplan/node"
)

func TestQuerySelector(t *testing.T) {
	f, _ := os.ReadFile("./asdf.html")
	doc, _ := node.ParseHTML(string(f))

	got := QuerySelector(doc, "a[title=\"caca\"]")
	want := `<a href="/logout" class="product_name_home" title="caca">GUACHIPILnyagamer, OAX</a>`
	if got.HTML() != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
