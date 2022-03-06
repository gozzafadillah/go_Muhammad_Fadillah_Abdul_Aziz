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
### Variadic function 
> Digunakan untuk case dengan inputan dengan jumlah yang berbeda-beda. Membuat variadic function, contoh :
```go
    func data(number...int)bool{
        return true
    }
```

### Anonymous function
> merupakan function tanpa nama yang berfungsi mengelompokan code. Membuat function Annonymous ada 3 cara :

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

### Closure
> Merupakan salah satu penjabaran atau tindak lanjut Anonymous function pada kondisi menggunakan referensi variable diluar local anonymous function.

```go 
    func getCounter() func () {
        counter := 0
        return func () {
            counter ++
            return counter
        }
    }
```

### Defer function 
> fungsi / statement yang dieksekusi terakhir. contoh :

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

### Panic function, pada panic function bila ada sebuah kondisi yang bila
> terpenuhi akan mengeksekusi panic yang berakibatkan program dikeluarkan secara paksa.

## Struct
> Struct merupakan kumpulan dari collection of field atau method yang digunakan menyerupai object oriented pada bahasa pemograman lain. Contoh :

```go
    type Student struct { 
        // <nama-field> <tipe-data>
        name string
        score int
    }

    func main () {
        var person Student
        // cara pertama
        person.name = "Fadil"
        person.score = 90
        // cara kedua
        var person2 Student {
            nama : "aziz"
            score : 77
        }
    }

```
## Interface
>Merupakan kumpulan dari collection of method yang menempel pada sebuah object. Contoh : 
```go
   type Persegi interface {
       luas() int
       keliling() int
   } 

```

## Method
> Fungsi yang ditempel pada sebuah type. Lalu ada beberapa alasan mengapa kita belajar method.
1. Membantu menulis object oriented code di dalam bahasa Go.
2. Membantu menghindari permasalahan name conflict.
3. Cara ini memudahkan untuk dibaca.
