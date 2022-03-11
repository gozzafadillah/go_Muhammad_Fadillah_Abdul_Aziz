# Recursive - Number Theory - Sorting and Searching

## Recursive
Merupakan sebuah fungsi yang dipanggil dalam fungsinya tersebut / dalam local state fungsi tersebut.

```go
    func loop(number int)int{
        if number == 1 {
            return 1
        } else{
            return loop(number - 1)
        }
    } 
```

## Number Theory
Mempelajari mengenai type dasar integer yang secara tidak langsung berhubungan dengan matematika.

> Contohnya case :
1. Mencari Prime Number
2. Mencari FPB
3. Mencari KPK

## Sorting and Search
* Sorting merupakan pengurutan pada sebuah array dari kecil ke besar (Asc) atau dari besar ke kecil (desc)
* Searching Merupakan melakukan pengulangan pada sebuah array sampai value nya memenuhi