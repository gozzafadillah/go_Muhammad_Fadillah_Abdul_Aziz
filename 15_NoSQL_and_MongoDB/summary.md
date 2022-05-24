# NO SQL and MongoDB
No SQL merupakan sistem management non data relation yang tidak memerlukan skema tetap. NoSQL itu mencakup dari penyimpanan data terstruktur, semi-struktur, tidak terstruktur, dan polimorfik. MongoDB salah satu database nonSQL.
# Kelebihan NonSQL dan MongoDB
1. Penyimpanan tidak memerlukan tabel (Shema less), MongoDB terstruktur dan terintegrasi dengan JSON (JavaScript Object Nation).
2. Fast Development, Karena Schema Less maka pembuatan lebih cepat.
3. Support Big Size File
4. Support Cluster
5. Tidak perlu untuk menggunakan tabel terstruktur, MongoDB berkerja secara otomatis untuk membuat struktur tabel pada saat melakukan insert secara langsung.
6. Telah terintegrasi dengan JavaScript, Kueri pada MongoDB tidaklah sama dengan bahasa SQL yang lain, akan tetapi lebih memanfaatkan pada penggunaan bahasa pemograman JavaScript.

# Basic menggunakan MongoDB
* Collection
1. Membuat DB
```mongodb
use test;
```
2. Membuat Collection
```mongodb
db.createCollection('coba');
```

* Insert
1. Single input
```mongodb
    db.coba.insert({firstName:"Fadillah", lastName:"Abdul", gender:"male"});
```
Bulk Insert (Multiple Insert)
```mongodb
    db.coba.insert({firstName:"Fadillah", lastName:"Abdul", gender:"male"}, {firstName:"Maura", lastName:"Audirra", gender:"female"});
```

* Find
1. Menggunakan function find
```mongodb
    db.coba.find(); -> raw
    db.coba.find().pretty(); -> better
```
2. Mencari dengan parameter
```mongoDB
    db.coba.find({firstName:"Muhammad"});
```
3. Mencari dokumen dalam dokumen
*   ```mongoDB
        db.coba.find({firstName:"Muhammad"}, {$set: {
            address: {street:"Sukajadi", city:"Bandung"}
        });
    ```
*   ```mongoDB
        db.coba.find({"address.city":"Bandung"});
    ```
4. Sorting
    *   ```mongodb 
            db.coba.find().sort(firstName:1); // asc
        ```
    *   ```mongodb
            db.coba.find().sort(firstName:-1); // desc
        ```
5. Limit
```mongodb
    db.coba.find().pretty().limit(1);
```
6. Count
```mongodb
    db.coba.find({gender:"male"}).count();
```

* Update
1. Update Dokumen
```mongoDB
db.coba.update({firstName:"Muhammad"}, {firstName:"Fadillah", lastName:"Abdul", gender:"male", age:17});
```
2. $set
```mongoDB
db.coba.update({firstName:"Muhammad"}, { $set:{lastName:"Abdul"});
```
3. $inc
```mongoDB
db.coba.update({firstName:"Muhammad"}, { $inc:{age:17}});
```
4. $unset
```mongoDB
db.coba.update({firstName:"Muhammad"}, { $unset:{age:""});
```
5. $rename
```mongoDB
db.coba.update({firstName:"Muhammad"}, { $rename:{gender:"sex"});
```
6. upsert (update or insert)
Update
```mongoDB
db.coba.update({firstName:"Muhammad", lastName:"Abdul"});
```
Insert
```mongoDB
db.coba.update({firstName:"Muhammad", lastName:"Abdul"}. {upsert:true});
```

* Remove
1. Hapus Dokumen
```mongoDB
    db.coba.remove({firstName:"Muhammad"})
```
2. Hapus 1 Dokumen
```mongoDB
    db.coba.remove({firstName:"Muhammad"}, {justOne:true})
```