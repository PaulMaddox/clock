# clock

Clock is a small, portable Go module for retrieving and setting the system clock.

On Linux/Darwin based systems it will use [syscall.Gettimeofday()](http://golang.org/pkg/syscall/#Gettimeofday) [syscall.Settimeofday()](http://golang.org/pkg/syscall/#Settimeofday). On Windows based systems it will use [w32.GetSystemTime()](http://msdn.microsoft.com/en-gb/library/windows/desktop/ms724390(v=vs.85).aspx) and [w32.SetSystemTime()](http://msdn.microsoft.com/en-gb/library/windows/desktop/ms724942(v=vs.85).aspx).

## Example usage

```go
package main

import (
	"time"

	"github.com/PaulMaddox/clock"
)

func main() {

    // Get the time
	t, err := clock.Get()
	if err != nil {
		panic(err)
	}

    // Set the time to 12hrs in the future
	t2 := t.Add(12 * time.Hour)
	if err := clock.Set(t2); err != nil {
		panic(err)
	}

}
```
