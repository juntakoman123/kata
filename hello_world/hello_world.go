package helloworld

import (
	"fmt"
	"io"
)

func HelloWorld(w io.Writer) {
	fmt.Fprint(w, "HelloWorld")
}
