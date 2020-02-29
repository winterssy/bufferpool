# bufferpool

**[bufferpool](https://pkg.go.dev/github.com/winterssy/bufferpool)** provides a pool of byte buffers, inspired by **[uber-go/zap](https://github.com/uber-go/zap)** .

![Build](https://img.shields.io/github/workflow/status/winterssy/bufferpool/Test/master?logo=appveyor) [![codecov](https://codecov.io/gh/winterssy/bufferpool/branch/master/graph/badge.svg)](https://codecov.io/gh/winterssy/bufferpool) [![Go Report Card](https://goreportcard.com/badge/github.com/winterssy/bufferpool)](https://goreportcard.com/report/github.com/winterssy/bufferpool) [![GoDoc](https://img.shields.io/badge/godoc-reference-5875b0)](https://pkg.go.dev/github.com/winterssy/bufferpool) [![License](https://img.shields.io/github/license/winterssy/bufferpool.svg)](LICENSE)

## Install

```sh
go get -u github.com/winterssy/bufferpool
```

## Usage

```go
import "github.com/winterssy/bufferpool"
```

## Quick Start

```go
package main

import (
	"fmt"

	"github.com/winterssy/bufferpool"
)

func main() {
	const dummyData = "hello world"
	buf := bufferpool.Get()
 	defer buf.Free()
	buf.WriteString(dummyData)
	fmt.Println(buf.String())
	// Output:
	// hello world
}
```

## License

**[MIT](LICENSE)**
