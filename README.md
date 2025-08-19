# HELLO IN GO

A package that provides a simple way to get greetings in Go.

## Install
```bash
go get -u https://github.com/carjosdan/greetings
```
## Use
```go
import (
	"github.com/carjosdan/greetings"
)
func main() {

	message, err := greetings.Hello(names)
	if err != nil {
		log.SetPrefix("greetings: ")
		log.SetFlags(0)
		log.Fatal(err)
	}
	fmt.Println(message)

}
```