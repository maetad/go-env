# go-env
Getting environment variable with auto casting type

## Install

```
go get github.com/maetad/go-env
```

## Example

```go
package main

import (
	"time"

	goenv "github.com/maetad/go-env"
)

func main() {
	isDev := goenv.GetEnv[bool]("IS_DEV", false)
	timeout := goenv.GetEnv[int]("TIMEOUT", 5) * time.Second
}
```
