# Resume Basic Programming

## Basic Programming
### Apa itu bahasa Go
> Go atau Golang, merupakan salah satu bahasa yang dikembangkan oleh google yang dapat dipakai back-end bahkan diluar backend pun bisa karena Go merupakan bahasa yang general purpose. Go merupakan bahasa yang baik untuk lower level program,itulah mengapa bahasa Go sedang populer dikalangan programer.

### Karakteristik Go 
* mudah, go merupakan bahasa pemograman yang mudah untuk dipelajari untuk pemula.
* handal, karena bahasa general purpose
* simple, simmple disini dalam segi pemakaiannya sangat ringan dan singkat.

## Alasan memakai Go
1. Go simple membuat yang ngoding nya fun.
2. gabungan dari bahasa pemograman yang static (bahasa C) dan berbasis seperti oop (bahasa C++).
3. Ringan, saat di build bahasa Go memiliki resource yang ringan, bahkan di loop berkali-kali besar resource nya stabil tidak berkurang atau bertambah.
4. built in conqurency.
5. Dipakai oleh perusahaan besar.

### Variable tipe data
> variable merupakan tempat penyimpanan sementara yang berfungsi untuk keberlangsungan membuat sebuah ekspresi, dan uniknnya Go hanya dapat menentukan satu type data untuk satu variable.

### Tipe data
> Seperti halnya bahasa pemograman secara umumnya, Go memiliki tipe data yang terdiiri dari :
* Boolean
* numeric 
    + integer
        - uint
        - uint16
        - uint32
        - uint64
        - int8
        - int16
        - int64
    + float
        - float32
        - float64
    + complex
        - complex64
        - complex128
* String

### Variable declaration 
> contoh deklarasi dibahasa Go
```go

    var age int

```

### Type variable declaration 
> Ada macam untuk membuat deklarasi variable untuk bahasa Go

* long

 1. ```go
    var age int
 go```
2. ```go
    var age int = 21
 go```
3. ```go
    var age, height int 
 go```
4. ```go
    var age, height int = 21, 165
 go```

* short

1. ```go 
    age := 21
    go```

### Zero value
> saat dideklarasikan seperti ini
```go 
    var age int
```
> maka age = 0
> Contoh zero value
 * boolean = false
 * float = 0.0
 * int = 0
 * string = ""

### const 
> Const, seperti variable yang membedakannya bila telah dideklarasikan maka tidak dapat diubah kembali.
> Contoh ada 2 :
* Single : ```go  const pi = 3,14 go```
* Multiple 
```go
 const pi = 3,14 
```

## Branching
### if, if-else, else
* ``` if <kondisi> { <statement> } ```
* ``` if <kondisi> {
        statement1
    } else if <kondisi-2> {
        <statement-2>
    }   ```
* ``` if <kondisi> {
    <statement>
    } else {
        <statement-else>
    }  ```
### switch 
* ``` switch <trigger> {
        case <trigger-1>:
            <state-1>
        case <trigger-2>:
            <state-2>
         ...
         default: 
            <state-default>
} ```

### looping
* simple looping with for
```
    for (<iteration>; <kondisi_memenuhi>, <iteration> ++){
        <ekspresi / statement> 
    }
```
