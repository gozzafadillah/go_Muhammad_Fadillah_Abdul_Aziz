# Concurrrent Proggraming

## Sequential
> Sequential Program secara berurutan, jadi seblum tugas nya beres maka tidak dapat melakukan tugas selanjutnya (ada waiting task).
## Parallel
> Program dapat dikerjakan bersamaan, syarat nya harus multicore pada processor di server.
## Concurrent
> Menjalankan tugas secara bersamaan dan secara independent (walau satu core dia bekerja secara bersamaan)

### Apa itu Go routine
> Adalah proses/fungsi penggunaan concurent dengan bahasa golang.

Feature pada GO
1. Concurrent execution (Go routine)
2. Synchronization and messaging (Channel)
3. Multi-way concurrent control (Select)

### Multiple Goroutine
> Execution disegi eksekusi intinya multiple goroutine itu membagi fungsi-fungsi dalam satu core, sehingga dapat membagi bagi core lain untuk yang lain.

### Channel dan Select
> Channel untuk goroutine bisa saling berinteraksi dengan yang lain.
> Select membuat lebih mudah untuk mengontrol data komunikasi 