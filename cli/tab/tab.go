// Helper to print tabular data
package tab

import (
	"fmt"
	"os"
	"text/tabwriter"
)

// Tab output tabulated content.
type Tab struct {
	*tabwriter.Writer
}

// New creates a tabwriter
func New() (*Tab, func()) {
	t := Tab{}
	t.Writer = tabwriter.NewWriter(os.Stdout, 1, 2, 2, ' ', 0)
	fn := func() {
		t.Writer.Flush()
	}
	return &t, fn
}

// Output a line
func (w *Tab) Output(a ...interface{}) {
	for i, it := range a {
		fmt.Fprint(w, it)
		if i+1 < len(a) {
			w.Write([]byte{'\t'})
		} else {
			w.Write([]byte{'\n'})
		}
	}
}
