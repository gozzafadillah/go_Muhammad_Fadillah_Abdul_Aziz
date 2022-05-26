# MongoDB Advance Query - Array Aggregation

## Export import / dump restore
1. Export
    mongoexport -u myTester --db text --collection coba --out coba.json => output berupa file json
2. Import
    mongoimport -u myTester -c coba -d test --mode upsert --file coba.json => membaca file mongo dalam bentuk json
3. Dump
    mongodump -u myTester --db test --collection coba
4. Restore
    mongorestore -u myTester --collection coba --db test dump/test/coba.json

## advance query
1. Comparison
    * $ne -> ```mongodb db.coba.find({gender:{ $ne: "male"} }) ``` 
    * $ge-$lt | $gte-$lte -> ```mongodb db.coba.find({age: { $gte:18, $lte: 20 }}) ``` (18 <= x < 20)
2. logical
    * $or
        ```mongodb
            db.coba.find({ $or: {{gender:"female"}, {age:18}} });
        ```
    * $and
        ```mongodb
            db.coba.find({ $and: {{gender:"female"}, {age:18}} });
        ```
    * $in -> mencari yang termasuk
        ```mongodb
            db.coba.find({city: {$in: ["Bandung", "Jakarta", "Surabaya"] } );
        ```
    * $nin -> mencari yang tidak termasuk
        ```mongodb
            db.coba.find({city: {$in: ["Bandung", "Jakarta", "Surabaya"] } );
        ```
    * $not -> mencari yang tidak termasuk
        ```mongodb 
            db.coba.find({city: { $not : {$in: ["Bandung", "Jakarta", "Surabaya"] } } );
        ```
    * Evaluator
        $regex
            ```mongodb
                db.coba.find({ firstName: /AZ/});
                db.coba.find({firstName: { $regex: /^A/  } });
            ```
## Array / Embedded Document
1. Find
    * $all
        ```mongodb
            db.coba.find({ firstName: { $all: ["Fadillah", "Muhammad"] } });
        ```
    * $size
        ```mongodb
            db.coba.find({ firstName: { $size: 2 } });
        ```
    * $slice
        ```mongodb
            db.coba.find({firstName: "Muhammad"}, { city: { $slice: 1 } });
            db.coba.find({firstName: "Muhammad"}, { city: { $slice: -1 } });
        ```