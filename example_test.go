package bufferpool_test

import (
	"fmt"

	"github.com/winterssy/bufferpool"
)

func Example_acquireBuffer() {
	const dummyData = "hello world"
	buf := bufferpool.Get()
	defer buf.Free()
	buf.WriteString(dummyData)
	fmt.Println(buf.String())
	// Output:
	// hello world
}
