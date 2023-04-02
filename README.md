# Simple Go version string 

Requires Go 1.18+ and a local git repository with an actual commit. Go only adds vcs meta data for actual builds, not for tests. 

```go
import (
    "fmt"
    "os"
    "github.com/gwillem/go-buildversion"
)

func main() {
    fmt.Printf("Starting %s, version %s\n", os.Args[0], buildversion.String())
    // "Starting myprog, version 20230304-08d1911"
}
```