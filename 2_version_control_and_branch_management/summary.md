# Resume Git version, Branch, Conflict, & Collaboration Workflow

## 3 Poin yang saya dapatkan:
* Versioning, Merupakan interaksi antara project dengan tim dimana terjadinya sebuah revisi / perubahan. dalam versioning pada git bila ingin merubah versi ke sebelumnya tinggal di ubah dengan "git log --outline" disana kita dapat me reset secara soft atau hard ke versi yang salah atau tidak kepakai.
* Branching, Membuat cabang untuk mempermudah tahapan development. setiap development diwajibkan me-branch setidaknya lebih dari 2 agar file origin tersimpan aman dan bebas bug dari file yang sedang kita develop.
* Collaboration Workflow, meneruskan pemahaman saya dari branching disini kita menerapkan branching ke 4 branch berbeda.
Ada main/master, development, featureA, featureB. Setiap programmer melakukan pembuatan feature tanpa harus merubah pada di file origin.

## Beberapa Tips yang harus dihindari :
+ Biar kan branch origin tidak terdistribusikan.
+ Hindari perubahan secara langsung di branch development apalagi di master. 
+ Tidak boleh merge branch dari feature langsung ke main/master.
+ Ketika semua dicek feature nya dan tidak ada bug maka pihak developer lah yang berhak merge ke main / master.