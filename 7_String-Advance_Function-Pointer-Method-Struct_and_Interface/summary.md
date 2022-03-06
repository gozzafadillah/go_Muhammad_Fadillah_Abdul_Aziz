# String - Advance Function - Pointer - Method - Struct - Interface

## String
> String merupakan kumpulan tipe data yang berisi character menjadi sebuah kalimat.
### Len pada string
> Karena sebuah string itu dapat disebut juga adalah array, maka string dapat dicari total dari Arraynya.

### Compare
> String pun dapat dicompare dengan string lain dengan return data boolean.

### Contains
> Sama seperti compare hanya membandingkan per char pada string.

### Sub-string
> String karena sebuah array maka string dapat dipanggil secara per index.

### Replace
> sebuah feature dalam package strings yang berfungsi mengubah string tersebut.

### Insert
> Dapat menambahkan string pada index tertentu.

## Function Lanjutan
> Variadic function digunakan untuk case dengan inputan dengan jumlah yang berbeda-beda. Membuat variadic function, contoh :
```go
    func data(number...int)bool{
        return true
    }
```

> Anonymous function merupakan function tanpa nama yang berfungsi mengelompokan code. Membuat function Annonymous ada 3 cara :

1. Cara pertama 
```go
    func (){
        // statement
    }()
```
2. Cara kedua
```go
    value := func (){
        fmt.Println("Hello world")
    }
```
3. Cara ketiga
```go
    func (sentence string){
        fmt.Println(sentence)
    } ("Hello world")
```

> Closure, merupakan salah satu penjabaran atau tindak lanjut Anonymous function pada kondisi menggunakan referensi variable diluar local anonymous function.

```go 
    func getCounter() func () {
        counter := 0
        return func () {
            counter ++
            return counter
        }
    }
```

> Defer function fungsi / statement yang dieksekusi terakhir. contoh :

```go
    func main () {
        defer func() {
                fmt.Println("Data pertama!")
            }()
        func() {
            fmt.Println("Data kedua!")
        }()
    }
```

> Panic function, pada panic function bila ada sebuah kondisi yang bila terpenuhi akan mengeksekusi panic yang berakibatkan program dikeluarkan secara paksa.

