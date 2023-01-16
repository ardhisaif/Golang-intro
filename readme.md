# 1. Golang Advantage
- One of the fastest back-end programming languages
- Active developer community
- Advanced concurrency techniques
- Automatic garbage collection
- Static type, Computer readable

# 2. Pointer
`pointer is a reverence or memory address`
```go
var num int = 4
// 4 is value, pointer is reverence or address where 4 was stores
```

```go
var numberA int = 4
var numberB *int = &numberA

fmt.Println("numberA (value)   :", numberA)  // 4
fmt.Println("numberA (address) :", &numberA) // 0xc20800a220

fmt.Println("numberB (value)   :", *numberB) // 4
fmt.Println("numberB (address) :", numberB)  // 0xc20800a220

// Pointer variable can only store address
// Use & to convert value to address
// Use * to convert address to value
```

# 3. Goroutine 
- thread ringan
- ukuran sangat kecilsekitar 2 kb, lebih kecil dibanding thread 1000 kb
- berjalan secara concurent didalam thread
- GOMAXPROCS(number sejumlah core CPU)
- untuk menggunakan goroutine tinggal menambahkan syntax go
