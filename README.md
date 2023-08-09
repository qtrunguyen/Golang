# Golang

## 1. Pros and Cons

### Pros
- Simple, easy learning curve
- Faster compilation time (compared to other backend languages Javascript, PHP, ..)
- Concurrency

### Cons
- Lack of libraries

## 2. Data Types
- Static typed(int, float, bool)
- Pointers: Have pass by value/pointer (similar to C++)
- Array & Slices: Array size are fixed, Slices size can grow/shrink overtime
- Struct
- Map
- Interfaces

## 2. Goroutines and Channels

### Goroutines
- Concurrent execution of functions or code blocks
- All goroutines run when main goroutine starts, and exit when main finishes
- Example

``` go
package main

import (  
    "fmt"
)

func hello() {  
    fmt.Println("Print this line FIRST")
}
func main() {  
    go hello()
    time.Sleep(3 * time.Second)
    fmt.Println("Print this line SECOND")
}

```

### Channels
- Channels are communication pipeline between goroutines
- All code execution will be stopped until reading and writing channels finish
- Channels can be bi-direction or single-direction
```
Bidirectional channel : chan T
Send only channel: chan <- T
Reeceive only channel: <- chan T
```

- Channels can be closed
```
close(channel)
```

- Example
```go
package main

import (  
    "fmt"
    "time"
)

func hello(done chan bool) {  
    fmt.Println("Print this line FIRST")
    done <- true
}
func main() {  
    done := make(chan bool)
    go hello(done)
    <-done
    fmt.Println("Print this line SECOND")
}
```



