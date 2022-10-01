package internal

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestUnmarshalFileListText(t *testing.T) {
	text := `
list([
{file:'scf20220910.log.gz',size:10965},
{file:'scf20220911.log.gz',size:1290}
]);
`
	want := []string{"{\"file\":'scf20220910.log.gz',\"size\":10965}", "{\"file\":'scf20220911.log.gz',\"size\":1290}"}
	t.Run("test", func(t *testing.T) {
		got := UnmarshalFileListText(text)
		if diff := cmp.Diff(got, want); diff != "" {
			t.Errorf("(-got+want)\n%v", diff)
		}
	})
}
