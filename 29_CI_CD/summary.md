# CI / CD

## System & Software Deployment 
Deployment adalah kegiatan yang bertujuan untuk menyebarkan aplikasi/produk yang telah dikerjakan oleh developer.

### Strategi Deployment 
* Big bang deployment strategy atau sering disebut replace
* Rollout Deployment strategy 
* Blue / Green Deployment Strategy
* Canary Deployment Strategy

### Big Bang Deployment / Replace
Sebagai gambaran saat kita update server ke server versi terbaru kita harus mematikan server tersebut lalu kita timpa. 
Kelebihannya :
1. Mudah untuk diimplementasikan, tinggal di replace.
2. Perubahaan kepada sistem secara langsung 100%. 
Kekurangan :
1. Terlalu beresiko downtime cukup lama.

### Rollout Deployment 
Sebagai gamabaran kita mengupdate tidak semuanya yaitu satu-satu, sampai seluruh server terupdate ke versi terbaru.
Kelebihan :
1. Lebih aman dari downtime
Kekurangan : 
1. Akan ada 2 versi aplikasi berjalan secara bersamaan sampai server ter deploy, dan bisa membuat bingung.
2. Karena sifatnya perlahan Deployment dan Rollback lebih lambat ketimbang big bang deployment.
3. Tidak ada kontrol request. Server baru terdeploy dengan aplikasi versi baru. Aplikasi baru tersebut langsung mendapatkan request yang sama banyak dari server yang lain.

### Blue Green Deployment Strategy
Gambaranya kita menyipakan server baru, lalu di isi aplikasi terbaru. Setelah itu server lama diswitch ke server baru tersebut.
Kelebihan :
1. Perubahan sangat cepat
2. Tidak ada issues perbedaan versi
Kekurangan : 
1. Resource yang diperlukan lebih banyak karena harus mempersiapkan server lagi
2. Testing benar-benar sangat diprioritaskan sebelum di switch.


