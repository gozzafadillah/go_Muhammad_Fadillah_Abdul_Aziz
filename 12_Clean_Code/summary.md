# Clean Code

## Apa itu Clean Code
> Clean code adalah code yang mudah dibaca oleh pemograman lain disaat dilihat, dipahami, dan dimodifikasi.

Kenapa kita perlu Clean code
1. Mudah untuk berkolaborasi
2. Feature yang kita buat dapat berkualitas karena mudah mudah dilihat, dipahami, dan dimodifikasi.
3. Pembangunan program lebih cepat karena code dapat mudah dibaca, dipahami, dan dimodifikasi.

## Karakteristik Clean Code
1. Penamaan Variable mudah difahami
2. Penamaan Variable mudah dieja
3. Penamaan Variable harus singkat namun mendeskripsikan
4. Penamaan Variable harus konsisten (huruf dan penamaan)
5. Hindari Konteks yang tidak perlu
7. Gunakan komentar untuk dokumentasi
8. Pada fungsi buat parameter agar lebih pendek. Deklarasikan variable pada local bila memungkinkan.
9. Formating lebar baris 80-120 Char, 1 class 300 - 500 baris, baris code berdekatan, dekat dengan fungsi pemanggilan, deklarasi tidak jauh dari pengguna, perhatikan indeksasi, dan menggunakan prettier / formater


## KISS, Keep It So Simple
1. Fungsi atau class kecil
2. fungsi dibuat hanya satu pekerjaan
3. jangan gunakan banyak argumen pada fungsi
4. harus diperhatikan untuk mencapai kondisi seimbang, kecil, dan jumlah minimal

## DRY, Don't Repeat Youreself
> Duplikasi code sering terjadi karena sering copy paste. Untuk menghindari buatlah fungsi yang dapat digunakan secara berulang-ulang.

## Refactoring
> Refactoring adalah restruktur internal tanpa mengubah proses fungsi internal dan hanya merubah external.

### Teknik Refactoring 
1. Membuat Abstraksi 
2. Memecah code dengan fungsi
3. Perbaiki penamaan dan lokasi code
4. Deteksi code yang memiliki duplikasi

