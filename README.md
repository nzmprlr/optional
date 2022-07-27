# optional [![GoDoc](https://godoc.org/github.com/nzmprlr/optional?status.svg)](http://godoc.org/github.com/nzmprlr/optional) [![Go Report Card](https://goreportcard.com/badge/github.com/nzmprlr/optional)](https://goreportcard.com/report/github.com/nzmprlr/optional) [![Coverage](http://gocover.io/_badge/github.com/nzmprlr/optional)](http://gocover.io/github.com/nzmprlr/optional)

## Usage

``` go
// Comparable - go primitives
optional.Comparable(0).Empty()      // true
optional.Comparable("").Else("str") // str
optional.Comparable("").Present()   // false
optional.Comparable(0.0).IfEmpty(func(f float64) {
    fmt.Println("float '0.0' is empty")
})
optional.Comparable(1).IfPresent(func(i int) {
    fmt.Println("int 1 is present")
})
optional.Comparable(1).If(func(i int) bool {
    return i == 1
}).Get() // 1

// Slice
optional.Slice([]int{}).Empty()        // true
optional.Slice([]int{}).Else([]int{1}) // int{1}
optional.Slice([]int{}).Present()      // false
optional.Slice([]int{}).Else([]int{1}) // int{1}
optional.Slice([]string{"str"}).IfPresent(func(t []string) {
    fmt.Println(t) // str
})
optional.Slice([]int{1, 2, 3}).If(func(t []int) bool {
    return len(t) == 2
}).Else([]int{1, 2}) // int{1,2}

// Map
optional.Map(map[int]string{}).Empty()              // true
optional.Map(map[int]int{}).Else(map[int]int{1: 1}) // [int]int{1:1}

// Error
var e error
optional.Error(e).IfEmpty(func(a any) {
    fmt.Println("e is nil")
})

// Interface
var i interface{}
optional.Interface(i).IfEmpty(func(a any) {
    fmt.Println("i is nil")
})
```