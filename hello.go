package gopathless

import (
	"fmt"
	"io"
)

func HelloLibrary(w io.Writer) {
	fmt.Fprintf(w, "heji, mundus!\n")
}
