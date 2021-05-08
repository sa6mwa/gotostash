# Format bytes as megabytes, mebibytes, etc...

The two initial functions in this package (`ByteCountSI` and `ByteCountIEC`)
are courtesy of CS prof Stefan Nilsson and are redistributed from
<https://yourbasic.org/golang/formatting-byte-size-to-human-readable-format/>.
Content is Copyright (C) Stefan Nilsson and distributed under [CC BY
3.0](http://creativecommons.org/licenses/by/3.0/).

## Usage

```go
package main
  
import (
  "fmt"

  "github.com/sa6mwa/gotostash/pkg/humanreadable"
)

func main() {
  fmt.Println(humanreadable.IEC(1500000))
  // Output: 1.4 MiB
}
```
