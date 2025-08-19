# HELLO IN GO

A package that provides a simple way to get greetings in Go.

## Install
```bash
go get -u github.com/carjosdan/greetings
```
## Use
```go
import (
	"fmt"

	"github.com/carjosdan/greetings"
)
func main() {

	message, err := greetings.Hello("Alice")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(message)

}
```