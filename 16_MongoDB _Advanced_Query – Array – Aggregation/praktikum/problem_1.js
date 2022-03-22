// use library

db.Books.insert([{_id:1, title: "Wawasan Pancasila", authorID: 1, publisherID: 1, price:
71200, stats: {page: 324, weight: 300}, publishedAt: new Date("2018-10-01"),
category: ["social", "politics"]},
{_id: 2, title: "Mata Air Keteladanan", authorID: 1, publisherID: 2, price:
106250, stats: {page: 672, weight: 650}, publishedAt: new
Date("2017-09-01"), category: ["social", "politics"]},
{ _id: 3, title: "Revolusi Pancasila", authorID: 1, publisherID: 1, price:
54400, stats: {page: 220, weight: 500}, publishedAt: new Date("2015-05-01"),
category: ["social", "politics"]},
{_id: 4, title: "Self Driving", authorID: 2, publisherID: 1, price: 58650,
stats: {page: 286, weight: 300}, publishedAt: new Date("2018-05-01"),
category: ["self-development"]},
{_id: 5, title: "Self Disruption", authorID: 2, publisherID: 2, price:
83300, stats: {page: 400, weight: 800}, publishedAt: new Date("2018-05-01"),
category: ["self-development"]}])


// -- 1. Menampilkan data Books dari Authors id = 1 dan id = 2
db.Books.find({authorid:{$ne: [1, 2]} }).pretty()

// -- 2. Tampilkan daftar buku dan harga author id 1
db.Books.find({authorID:1}, {_id:1, title:1, price:1})

// -- 3. Tampilkan total jumlah halaman buku author id 2
db.Books.aggregate([
  {
      $match: { authorID: 2 }
   },
  { 
  $group:{_id:2, TotalBalance:{$sum:"$stats.page"}}
  }
])

// -- 4. Tampilkan semua field Books dan Authors terkait
db.Books.aggregate([
    {
   $lookup:
     {
       from: "Authors",
       localField: "authorID",
       foreignField: "_id",
       as: "dataAuth"
     }
    }
])

// -- 5. Tampilkan semua field Books dan Publisher yang terkait
db.Books.aggregate([
    {
   $lookup:
     {
       from: "Publishers",
       localField: "publisherID",
       foreignField: "_id",
       as: "dataPublisher"
     }
    }
])

// -- 6. Tampilkan summary data Authors, Books, dan Publisher sesuai dengan Output

db.Books.aggregate([
  {
    $lookup: {
      from: "Authors",
      localField: "authorID",
      foreignField: "_id",
      as: "author"
    }
  },
  {
  $unwind: "$author"
  },
  {
    $lookup: {
      from: "Publishers",
      localField: "publisherID",
      foreignField: "_id",
      as: "publishers"
    }
  },
  {
  $unwind: "$publishers"
  },
  {
    $group: {
      _id: {$concat:["$author.firstName", " ", "$author.lastName"]},
      number_of_books : {$sum: 1},
      publisher : "$publishers.publisherName",
    }
  }
])


db.Books.aggregate([

    // Join with Author table
    {
        $lookup:{
            from: "Authors",       // other table name
            localField: "authorID",   // name of users table field
            foreignField: "_id", // name of userinfo table field
            as: "author"         // alias for userinfo table
        }
    },
    {   $unwind:"$author" },     // $unwind used for getting data in object or for one record only

    // Join with Publisher table
    {
        $lookup:{
            from: "Publishers", 
            localField: "publisherID", 
            foreignField: "_id",
            as: "publisher"
        }
    },
    {   $unwind:"$publisher" },

    // define which fields are you want to fetch
    {   
        $group:{
            _id : {$concat : ["$author.firstName", " ", "$author.lastName"]},
            number_of_books : {$sum: 1},
            list_of_books :  {$push : {$concat : ["$title", " (", "$publisher.publisherName", ")"] }}
        } 
    }
])


// -- 7. Diskon setiap Buku, diskon ditentukan harga buku
  // price < 60.000 -> diskon 1%
  // 60.000 < price < 90.000 -> diskon 2%
  // 90.000 < price -> diskon 3%
  db.Books.aggregate([
    { "$project": {
      "_id": "$_id",
      "title":"$title", 
      "price":"$price",
      "diskon" : {
        "$cond" : {
          "if" : {"$lt" : ["$price", 60000]},
          "then" : "Diskon 1%",
          "else" : {
            "$cond" : {
              "if" : {"$and" : [{"$gte": ["$price", 60000]},{"$lte": ["$price", 90000]} ]},
                  "then": "Diskon 2%",
                  "else": "Diskon 3%"
            }
          }
        } 
      }
    }}
  ])



  // -- 8. Tampilkan semua nama buku, harga, nama Authors, dan nama Publisher, urutkan dari harga termahal ke termurah.
  db.Books.aggregate([
    {
      "$lookup" : {
        from: "Authors",       // other table name
        localField: "authorID",   // name of users table field
        foreignField: "_id", // name of userinfo table field
        as: "author"         // alias for userinfo table
      }
    },
    {
      $lookup:{
          from: "Publishers", 
          localField: "publisherID", 
          foreignField: "_id",
          as: "publisher"
      }
    },
    {
      "$project" : {
        "title" : "$title",
        "price" : "$price",
        "author" : {$concat : [{ "$arrayElemAt": ["$author.firstName", 0] }, " ", { "$arrayElemAt": ["$author.lastName", 0] }]},
        "publisher" : {$concat : [{ "$arrayElemAt": ["$publisher.publisherName", 0]}]}
      }
    },
    {
      "$sort" : {"price": -1}
    }
  ])

  // -- 9. Tampilkan data nama buku harga dan Publisher, kemudian Tampilkan hanya data ke 3 dan ke 4
  db.Books.aggregate([
    {
      $match: {
        $or : [{"_id": 3 }, {"_id" : 4}]
      }
    },
    {
      $lookup:{
          from: "Publishers", 
          localField: "publisherID", 
          foreignField: "_id",
          as: "publisher"
      }
    },
    {
      $project : {
        "_id" : "$_id",
        "title": "$title",
        "price": "$price",
        "publisher" : {
          $concat : [{ "$arrayElemAt": ["$publisher.publisherName", 0]}]
      }
      }
    }
  ])