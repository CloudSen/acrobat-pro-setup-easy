package main

import (
    "fmt"
    "sync"
    "time"
    "crypto/sha256"
)

var ( appVersion = "5.9" )

func jipd5ynCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H3yMNsMpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6ljJNtXRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func I8JAzK8UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GFbil7ERWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uLs99NqUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1vmJQ0mYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qV9GRXq1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1TejNeUgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KsMmx9fgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mFrIrLcJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L4DRUhHFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HLnlRspcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W7x4LRW0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sxdZIWbaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JQXt7ZQBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pIFuaNzPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func b0VMfSJsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LULwFGLDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0Sn0G3ojWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tv8dtf0VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tQquW4LhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KZCK2nLCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZB6hfm3JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func U8uoBjq1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lGNY67NaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nbjHXapZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SIwa5XoBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4ummYA73Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EEC8dZcEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wmg3Y7bQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rqNoWxGqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5OEPRx0LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LGprH4H5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0XUy74TSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func afDIqasrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0bdQQ7GcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7mQGmlxBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NBVHXZ56Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZzaLi7dBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fNy1uvx6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dBOtajHVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ATWY4FELWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 43h4ul9mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OhQhKnOjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SDLliqSlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func f9gzz8zOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rFG0oAGbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BWsO5AKPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eYUkN2h2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WJsE4BajWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sinmPOX8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func s5lRKKUgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Vdf0a18wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func J8vnI38HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cLFRZgxFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m9bAIW5EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func a1CfCFn9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fn9Qvc0VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uT0vwHQFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6KWEOyySWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5UGuq1uSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qJo4wj3cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func p6hRJMF0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fhNRbbDoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TkfH0DUsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func poJMHGqAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LAinlWMBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func byzvOt2EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func f8QoSf30Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func O7OE2XVDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func r7eQJTbQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 36bUbyGWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cEio52r3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MNz9WmD8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vUBS6pdNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PmERuhSOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VzLtlVQLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XYOc5uDwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i8CXC7lHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WLdYhhxHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2sjhLmkDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GfO5baWhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zt8k132KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JFZ2kdVQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 94D9UNA1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TiOlH8q8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yLgxcsScWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zRkhYAalWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5tZIVYxFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jaFpWpDvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MPLdsat7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7jAPQOaQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fq8SODcvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wjcXoLXtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tMOIszZaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 54QYSZhuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jCmXLBgkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FlKSyYZcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Th7Fw5t6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eW2ea9gDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func h1llRi74Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AgrddBSZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NpLsz2ANWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func X3Dy51w3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HDOCcJB3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XzgdSUy1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OdohTifeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bgJZLrWBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2fsMc9RkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ll5pbe5eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YHouvuELWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func J6QqD64zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EsZIcn1gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aloS8n74Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WeNA5olKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jLoRpH7jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8KOmwSI3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oWzazgQ4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i7ksJBe0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6Ohqe0HcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NbPEQjvhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QhTeWCEaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OEhJJdi4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tl1ovmepWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GFlGG5LWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 73lCCCb4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func df287krnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9mA1vB7EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func innrg9TeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EkvVAe1gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E6d4xzVyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func urqCvPfSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RzuNqZnUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GKMehTvmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wNRixYhXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lGZetUKBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func A7P0M6tvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mv5hMSJQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ovJtABzGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3qFV3p5IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IgRlXIMZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func C9E0Rm20Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func h3hyjVC9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jkvKUScyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func P6dsPfMjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func f05u4umWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func o80GPunsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PnVp7Ns0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jE3AvdVCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1Pz5XvLEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BoWJaYJyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func s6AQc0mGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VeaO3uYsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JEuPo9jGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func k9Q0gWGwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Vq08rEggWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func leLCieAgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ToCMCBWCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Lqh9CibZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xuaP5lgiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func n0HHph3nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FHqaLcqJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BDz23BLnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9Ge4RUGGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MCgDT6N8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CdWayq3nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ti1k06nNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func P2WWgR4AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tZPmrSk6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HY9trjxSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bd7ro1SlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 155
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L166AyNXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PFOTnxSOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7W97mu7EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8Qy6Bi6XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YK1bRpZdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fVZ8SQsyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tfXNuENBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hSQGCmVhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xXy7tm01Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func v6CVI4ysWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func h6l3piVuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MonOi366Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aaqxpSdnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tGbq8I8qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6S39kBdXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PgV8DKdIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 35u2msySWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xgQggbp3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lY5ysw6DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func B8GylpxpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HAx1r6dXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sGRFMPCgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ThpgCQfJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UFaHlYngWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PSyUgzkvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9wnrsqN0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CF3oodCuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IId5paEKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DdLn2hkmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZLqeVlNSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OapVvdEnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vSfy9tDEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rhUd828pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nwwRzq0YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tmEq4NKcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BgZMYaNWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3Q2rc7kiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qV9I7kg1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SSSxpVS8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MyA3Jr00Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GncpX60tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func drCoTj5xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4SuqJB9OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MNGei7DJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IXKvpILnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YvgSK7w3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7cmg62ErWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4gtl9zPoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8cvnfdQ2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TlcMoebTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tVzbvLmWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AMDCFtBQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func h7gw4B2IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bdof4OGoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lSj9CuzCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GU51suuIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zJlXXBwIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x8giBe7LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4HJaEbjPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lSodkaUpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func raGBCZcmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func c4Rl949PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func envF8YcDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dHq0UE9pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YqJ07BquWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xpfBg3YJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 09mkkum3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EjjhJ89bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 88ZrR1KNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TaE7qk0vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VQQqFmJvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CUPMVBlGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func up5ffNAfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 155
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IiBogmpIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AueTis9bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VM46ox6yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rxwzh3TpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qrJLwh3GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lfvjp4BBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QVoe4MbcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ABDfNN8CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6T8mr4JCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VxwOjh5YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0HUVzKrzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bUiJlLk9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CFWrMutdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Noa9LbhfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pDDUcORiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func otf2ljX7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GW6pWNgEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 66GnpNALWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OCS5n4F1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5OOtgAPcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OwO2BEGNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AkbM01X0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func u4i8VMMcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 44TgQ5GHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Eq98BqMvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YZb2QhFgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oALzZ6qPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vgFsBauNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TRFjXVgGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HMcqOOheWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iHuwkPNIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2VKN1SVVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uf4fOTWtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func t84uux2PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g7w90DiHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func z4l05PpqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 88cZwa64Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func J7fU3BCmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func chd28TY7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rcf3hiodWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Y4FYRNWYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OJ5cWgDPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qDiLJabjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4eKY4kYoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vwvXGMAAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uZsad6MNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gieuVY7DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uAk9Pbw3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 75QPrT39Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CDYzOKXxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eAYlsof6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lP8xlj19Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DsmfAjgHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7kVaqYUtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bNltsD9DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bsjP00yJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LiRflK0CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func I73HuRC9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W8iXr0lEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KfJX2VleWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wLrskNTmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fIkaNaXxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 11umimovWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kqoAFnyGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 45CAIhqGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eiBymahSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RsUp6QG1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func peImHNQcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zSBdlCBzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3utIn4rQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DXepxQP5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func w3Vzu9HWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uXJjlsWcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9ZjtTwydWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Abj3sMgKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func k0u5dQ5gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tvG0rV68Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JXjRzvFnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hjPGZknRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5pgxj2F2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lS0BgYbEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ttCVWzPnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LBMwYm24Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func htzWQCXIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func J9vSnsiqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rWFEdCvCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H9zJFlzYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func n5MEIDYoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7Y4nnKAKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6z7WYx0iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qCBlBiHdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NKkxAbUXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DALRmHc6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1OfqzOOqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SpoBKbW4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SMdXqfAJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hyObn6k7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gwDCHsWPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HMvNtyoQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Yn0hBIQmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3yZGOs7EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0wLo97lEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AGRNuZpqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eED1t9jyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ULKMTFf5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YBXv63yfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1gy8HFY7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6AsutWouWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vzrRUNREWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qq8MaQnIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zd2wjb5eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HY7IzbA1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QFTwV7M7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oG7cX2D6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func D2Gx8CkpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nM1DlgvOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QgwksLZ6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5OEn5zjKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YfFjaDV9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4mreXlOEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MA1OirKQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func D7PkJycjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ClmTNk42Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0tZnP9iQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gHEM8sR5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UwvBqzpwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NvYuIv7MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ProaPtF6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cvuyPKuoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AjKQUHR2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rJkwGbiaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XlvwMo24Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6QG12xNXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hgrFeZS9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BRvzoQbHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func t9xKvSupWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nqFYLQFdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V3HCNPjeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gS7xuYiwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2pNnTTZqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0x0RooMuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FW9of8awWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FYQHOXXoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MV4BiP2MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cxLGr5MDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gRS0NvR4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tWWUWAslWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LMBREgb5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oXodUkQuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func n4HUAOZlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qebeItXZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DS2WQERRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jleXvGv9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0ZSC8AfoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func f3CjRSRPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func llwhEpjKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MzQP1NYTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2TsNQ23PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7F3DdHC5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TzHoterUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kfrPFCpbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g3J1aFkWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wv35RPfNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qxfK1NsdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XrXbSnybWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d9LdwjGGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UvPfYuBwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VufQAGKUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4HaL2RRRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HQeThrPWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BjI80Fj5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RjyweV5hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func v5GvxKIWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EUkyrCTgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YAaIR4XGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rdHRhvq9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AJWm4qj5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QGVhnbTJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gFWGg1t2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 33EB3f3jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CNZdHDSmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 15sLlap0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 71ThYZohWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func O0TV2MUqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZzDT0DWKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d0VMQ6EcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vEjxcmBhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zOsIlNk7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1WI5dOy7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func drqZmKgFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ByNSbYUcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 86G08S8zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tnCDfTvzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4qlVWa0aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QLSDo78WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func odOue4BcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VXfylEH8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gxVzvLTSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1k5ORbpcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IZBUpYkxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EghatQ7FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qwOcuNLDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kzUkiyyrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vYz9ZxuFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H09Q7gXYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ai2Ww5Q1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xF2Gg1mMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func I4VXn9Y0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yCQw66vRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kzTUHviuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wAVLGpRbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rxiVfXyhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2Dz2LWcuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8nvGQ5FxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H9YbtCDlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Rh85M56UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gWt974r5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 155
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dbzuGOQtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AnTfjlqMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0ie4z3ZzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UeI82of7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fvKPZIGcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kMfS5KDHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hnIF2IQPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ks9JdQTIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func l3BitxnvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GxN7TANMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func seXQS2kMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func o81xs1fxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5O0ci8weWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func grv8Zz3TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g4VmjkrnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func P3cnRkIfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func joGKR0huWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2QK5ZO5kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LcBEclA6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LBZ0S6orWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JMVC8Y1BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KSRjyRYFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qCJuoA3gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H3DoCqSEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iZVAxeEBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UylM8fDmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 13tbWMWCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xVXdfmTJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zMRwSXUVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RfqEox9yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ypLqppG7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HjFxQzesWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YvrZ4iakWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KBsnBzsrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func B358QNADWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0tgMyLfcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func k0wzpyznWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4HzVtlRJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func af0wVbfBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zOaievwXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func G7OeWxpLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EndIKK38Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5wOJMV0DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JZbMSRIaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MAgFHDh3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OOaXOrfYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8ibhg1DqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ROpu4VW2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IGFPVRl5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jvpTzZeHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FqxLVlQPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CQws6YPRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func t6IQo7njWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IMUrQ30yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rBM53lT7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bP4cizGDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Fx0ffcIXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wGx9rJi2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zIMaFfGuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MeZJf6miWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KFQ6to0WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LJ3zfseNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yQ07oe8WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KP1ekhwYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func o8aYKK8lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oyQKUW6dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xlQXA1X1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OxNynBocWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dulVJS2bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sFjyJlosWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pJj7FSoHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1RFh1w79Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6WL01R9yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nmPkgkuuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 50rLyWjyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yfiQ20e9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uxi488KzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QBOPVVRfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func B7j77H7tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func k4b2v4r9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cJgvF79qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6hFKcWtlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FIadAYyVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func M5p9RvDoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1AHdDYxGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OtYEzlETWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9mMRhMtoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5zAxgP7GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uMlu8KxsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4H9eTGe7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7ALgTVpmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RzFNo8UIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FEtk1xniWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LSNZZ9ljWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aqbn8MUXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eRyXro1HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func b9TlsU6bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xP603w4VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func t8TS2LOWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m5Hbvhv4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WgzIzVnLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sjtJQ4fOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rtWZF38dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func n2dMfvLwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HPxPHpYMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RkdTdMoHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9uEShpivWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func js3qWrmwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mYVTtddIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fw3V9MJ5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 06QFm93BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NDjLoZD7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func z6H1C2WoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TlW0ZZ64Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Vc0oDWnWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yaWQmaC0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sPAW91UJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fufwI6k1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sZHHmm1EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LgP2MbwTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fu7hOYI1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func o4psNp7RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Dsf3lPwzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MTxOQZOAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xf8fJjPqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wKkgZY9kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 59kIsApNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sa9SD3IDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9H96tgKHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OehQnBxHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 82n3ErhnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qyi0vYyOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EpBBHngQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6fxbpPtfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IU9f5MjHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func e0NuaH5mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zsVB0NuEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 08qBXtkbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sipvoM8oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TYMUazXOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ubV2oWoAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func c5Y0oAE6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0rNzBaKnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oGDrwyErWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0LUPUqHZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JyUGKiz0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ECVMJ52SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BBk3LhPDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W69irh2GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9pAQgeufWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 97S4nNh6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PVc6AOYmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8k1pQhaxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ffDFwGfzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uvmmTUWCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2IZpRj5ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DMCuVUzhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AMGvnBo4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E3XlTYZ8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CQkZKYAIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rt7CNG1EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func exSn0Ce3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GdjlolD1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FeeoGKqVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sEUgzwVVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DcDAr7LqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CJhcHNSKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oCGvT0pQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BN8YMULWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AfJVvaTPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PvEx1RX8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YJOPpOpiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TtoUcYaTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3fhjFVuyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2VBKh7YrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WJqclDULWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rdZ5fBF8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Y2sARtuvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Bdcj9VULWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bILzZXB3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QOs0TAqDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func POiuvWpCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qD8neXejWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xrrm6xCiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TaQNvgE2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func p35u6T2FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3y3tHa35Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mfMOE3eYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6a4za0QzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kzMYNys8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KAU8JiSkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func z5OmdMzKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JZmwSiLmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6avl4QGvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func f1C1eYseWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7qhdcp0pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bsoEVK1dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func a6Rsr5rwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func C8C5YpFxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lWYAiA0YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fLrGwJqTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fdcDHzX8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iveURNNDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func G0HBgpyZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wGs10kRJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PEl2IFKeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LUy4LCJWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kraqnHGnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BSmrgH9XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CegP2dSaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wZY6ynq5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func t06AL0XFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func r4kHm1i6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AITyTPNmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LvjcroAkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cExX4maWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func elaqDKMXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func exomoF6JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KpAwfTUfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EJGypoGiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Uta5tbJJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lTbGGuSXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FbQQ8t9nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func agStPB6OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IwrT5HKwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ViO137v5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fpLi3l82Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func n8rnvCXFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qhoKriKuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RGBHlfjeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JUFc2sqaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KxbMZchNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tqS1Wmj8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CmhPjxvlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3ki3cF1YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pTCG3EDQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9ijRFvTCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 84N8xIfGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gVQZguQmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yk3EKW7bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xTkTjbdjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MzU9uZhlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Hz0cQXuPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func noZPClMWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func S5sTnC9tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EbkqQdKUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jQ0fYbpwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PkJWgXiDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HvXFpsEKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kan20eOiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UcORTlvdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func s0ZRbZ8SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nDc57UMrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MX7s5Yj0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NuajftA1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func elh0uS37Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VGFwTmNGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func B1vqdqwGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qqtz2uXIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cLbAqlsTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UGB6RsR3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d4zi8ZT6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SpmBYKhXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func P6BzfLhHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KaRjKjrYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func I0KSimKeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zd8BbAZzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func O453Pd29Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7E8fgujQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0qgrbZVcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 76yfpHkHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DoWvmsngWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ibx6WFeDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qGtTR5qUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IKExkpKXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5g96BlBRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i2SSZx7BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gJN3Dp9XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Sz1LJvfCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VCQxQz8wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aRJ2ggTnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bMoJMWmlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3Wscn6EUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oy7AMOi1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VX7UnEjTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func umXPVghXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fDjSp5ieWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eVgwas86Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mJAmRdEKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bkpD8rLKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fjBQhafHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sKTkAzX6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LxWXDDmgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iGyK98HWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AEMNXx16Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kD4tTnX3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8lt3x4VNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6yTedbHQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KM0VaYVyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vUtFSeMGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EWshqTYLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func egmDhm0vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Fe5hYjBwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IxTlKP4gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AFnmwMG8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DOTapEkHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xV3wzV1vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nyumUmSUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KmAyP9UvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MTHwhbu9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dQbHPuqQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xLSpuovoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FFFetUjkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KZQdlHT8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IjlNdmYVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ya0VQ41IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dEE6aS4hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XsQUbUsVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pgJ0oHxuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func R4V8V0XgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OehaHtZQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZL6sdrKuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dbOoPkhHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Dw4QMHx2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MH3hVvHpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zMxMdvwiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 92M77ESAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QOOvm3WZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2JfEqVAkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func v47ikfMiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uJs67gMtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m3q8O9AZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SQO2XOmxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XEduph9LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XymXEvcQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WoHHb23aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iTKLn4nDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mP9oayRXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2fIturCwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func B5iBDozJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func D4eLLNN0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VUBjFe0qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZDmZYsFpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4FgsKXYRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9XKOSEk5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hRgEPIyvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ss1DyWUhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2KPio7E4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func s1oC8PbsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KMX2ctjFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rP0ayI1QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nTCrskJfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wvQlwuJHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func l3dvgA0lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 26orLAewWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sbjnceU6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Lkx2cUzKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Jy73ctoNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5SN85zJcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rXZsfXJiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 62G75GcXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wB6T7iJ2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tN2pML3AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XBtNEc0oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0vHY7klYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NK03VNSzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5sqc0vGHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GZANMOZDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DeGScLEhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7yi0J8NHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GFDa6TCFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6co7HElqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ggZEdd9hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GOySPTuEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func k662rpdcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YPlQIC5gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PansmEsMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i3QMeN6NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iGhuvEoFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SeRPdtQdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

