# Data Structure

## Array 
> Array merupakan struktur data yang memiliki group of type. Array hanya dapat didefinisikan oleh 1 jenis tipe data dan array memiliki ukuran yang fixed.

> Cara mendeklarasikannya
```go 
    var numbers [5]int = [5]int{1,2,3,4,5,6}
```
## Slice
> Slice merupakan struktur data seperti array, yang membedakannya adalah slice memilik ukuran data yang dinamis.

> Cara mendeklarasikannya
```go
    var number []int
```

## Map
> Map merupakan struktur data seperti array, yang membedakannya adalah map memiliki key dan value.

> Cara mendeklarasikannya
```go
    var number = map[int]int
```

## Function

> Function merupakan sekumpulan code yang dikelompokan dan dinamakan dengan nama tertentu.

> Cara mendeklarasikannya
1. Tanpa Parameter
```go
    func helloWorld() {

    }
```
2. Dengan Parameter
```go
    func tambah(x,y int){
        var tambah = x + y
        fmt.Println("Hasil tambah :",tambah)
    }
```
3. Mengembalikan nilai
```go
    func tambah(x,y int)int {
        var tambah = x + y
        return tambah
    }
```