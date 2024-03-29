// Create collection
db.createCollection('product')
db.createCollection('product_type')
db.createCollection('product_description')
db.createCollection('operators')
db.createCollection('transaction_detail')
db.createCollection('transaction')
db.createCollection('payment_methods')
db.createCollection('users')

// 1. Insert
// -- Insert 5 operators
db.operators.insert({_id: 1,name:"Telkomsel", created_at: ISODate("2022-03-19T01:29:29.965Z"), updated_at:ISODate("2022-03-19T01:29:29.965Z")})
db.operators.insert({_id: 2,name:"XL", created_at: ISODate("2022-03-19T01:29:29.965Z"), updated_at:ISODate("2022-03-19T01:29:29.965Z")})
db.operators.insert({_id: 3,name:"PLN", created_at: ISODate("2022-03-19T01:29:29.965Z"), updated_at:ISODate("2022-03-19T01:29:29.965Z")})
db.operators.insert({_id: 4,name:"Garena Voucher", created_at: ISODate("2022-03-19T01:29:29.965Z"), updated_at:ISODate("2022-03-19T01:29:29.965Z")})
db.operators.insert({_id: 5,name:"Playstore", created_at: ISODate("2022-03-19T01:29:29.965Z"), updated_at:ISODate("2022-03-19T01:29:29.965Z")})

// Insert 3 Product type
db.product_type.insert({_id: 1,name: "Pulsa",created_at: ISODate("2022-03-19T01:29:29.965Z"), updated_at:ISODate("2022-03-19T01:29:29.965Z")})
db.product_type.insert({_id: 2,name: "Paket Data",created_at: ISODate("2022-03-19T01:29:29.965Z"), updated_at:ISODate("2022-03-19T01:29:29.965Z")})
db.product_type.insert({_id: 3,name: "Voucher",created_at: ISODate("2022-03-19T01:29:29.965Z"), updated_at:ISODate("2022-03-19T01:29:29.965Z")})

// Insert 2 product dengan product type id = 1 dan operator id = 3
db.product.insert({_id: 1,product_type_id: 1, operator_id: 3, code: "PTL-100", name: "Pulsa Listrik 100", status:1, created_at: ISODate("2022-03-19T01:29:29.965Z"), updated_at: ISODate("2022-03-19T01:29:29.965Z")})
db.product.insert({_id: 2,product_type_id: 1, operator_id: 3, code: "PTL-150", name: "Pulsa Listrik 150", status:1, created_at: ISODate("2022-03-19T01:29:29.965Z"), updated_at: ISODate("2022-03-19T01:29:29.965Z")})


// Insert 3 product dengan product type = 2 dan operator id = 1
db.product.insert({_id: 3,product_type_id: 2, operator_id: 1, code: "TL-10", name: "Paket 10", status:1, created_at: ISODate("2022-03-19T01:29:29.965Z"), updated_at: ISODate("2022-03-19T01:29:29.965Z")})
db.product.insert({_id:4, product_type_id: 2, operator_id: 1, code: "TL-20", name: "Paket 20", status:1, created_at: ISODate("2022-03-19T01:29:29.965Z"), updated_at: ISODate("2022-03-19T01:29:29.965Z")})
db.product.insert({_id:5,product_type_id: 2, operator_id: 1, code: "TL-30", name: "Paket 30", status:1, created_at: ISODate("2022-03-19T01:29:29.965Z"), updated_at: ISODate("2022-03-19T01:29:29.965Z")})


// Insert 3 product dengan product type = 3 dan operator id = 4
db.product.insert({_id: 6,product_type_id: 3, operator_id: 4, code: "GA-10", name: "Voucher 1000", status:1, created_at: ISODate("2022-03-19T01:29:29.965Z"), updated_at: ISODate("2022-03-19T01:29:29.965Z")})
db.product.insert({_id:7,product_type_id: 3, operator_id: 4, code: "GA-20", name: "Voucher 2500", status:1, created_at: ISODate("2022-03-19T01:29:29.965Z"), updated_at: ISODate("2022-03-19T01:29:29.965Z")})
db.product.insert({_id:8,product_type_id: 3, operator_id: 4, code: "GA-30", name: "Voucher 4000", status:1, created_at: ISODate("2022-03-19T01:29:29.965Z"), updated_at: ISODate("2022-03-19T01:29:29.965Z")})


db.product_description.insert({_id:1, description: "Pulsa Token Listrik 75kwh", created_at: ISODate("2022-03-19T01:29:29.965Z"), updated_at: ISODate("2022-03-19T01:29:29.965Z")} )
db.product_description.insert({_id:2,description: "Pulsa Token Listrik 125kwh", created_at: ISODate("2022-03-19T01:29:29.965Z"), updated_at: ISODate("2022-03-19T01:29:29.965Z")} )
db.product_description.insert({_id:3,description: "Paket Data 1GB", created_at: ISODate("2022-03-19T01:29:29.965Z"), updated_at: ISODate("2022-03-19T01:29:29.965Z")} )
db.product_description.insert({_id:4,description: "Paket Data 2.5GB", created_at: ISODate("2022-03-19T01:29:29.965Z"), updated_at: ISODate("2022-03-19T01:29:29.965Z")} )
db.product_description.insert({_id:5,description: "Paket Data 4GB", created_at: ISODate("2022-03-19T01:29:29.965Z"), updated_at: ISODate("2022-03-19T01:29:29.965Z")} )
db.product_description.insert({_id:6,description: "Voucher Game 1000 Point", created_at: ISODate("2022-03-19T01:29:29.965Z"), updated_at: ISODate("2022-03-19T01:29:29.965Z")} )
db.product_description.insert({_id:7,description: "Voucher Game 2000 Point", created_at: ISODate("2022-03-19T01:29:29.965Z"), updated_at: ISODate("2022-03-19T01:29:29.965Z")} )
db.product_description.insert({_id:8,description: "Voucher Game 3000 Point", created_at: ISODate("2022-03-19T01:29:29.965Z"), updated_at: ISODate("2022-03-19T01:29:29.965Z")} )

// Insert 3 Payment 
db.payment_methods.insert({
    _id: 1,
    name: "Dana",
    status: 1,
    created_at: ISODate("2022-03-19T01:29:29.965Z"),
    updated_at: ISODate("2022-03-19T01:29:29.965Z")
})
db.payment_methods.insert({
    _id: 2,
    name: "ShopePay",
    status: 1,
    created_at: ISODate("2022-03-19T01:29:29.965Z"),
    updated_at: ISODate("2022-03-19T01:29:29.965Z")
})
db.payment_methods.insert({
    _id: 3,
    name: "Gerai Indomaret",
    status: 1,
    created_at: ISODate("2022-03-19T01:29:29.965Z"),
    updated_at: ISODate("2022-03-19T01:29:29.965Z")
})

// Insert 5 User
db.users.insert({
    _id:1,
    name: "Aziz",
    status: 1,
    dob: "2000-03-09",
    gender: "L",
    created_at: ISODate("2022-03-19T01:29:29.965Z"),
    updated_at: ISODate("2022-03-19T01:29:29.965Z")
})
db.users.insert({
    _id:2,
    name: "Maura",
    status: 1,
    dob: "2001-09-03",
    gender: "L",
    created_at: ISODate("2022-03-19T01:29:29.965Z"),
    updated_at: ISODate("2022-03-19T01:29:29.965Z")
})
db.users.insert({
    _id:3,
    name: "Kiana",
    status: 1,
    dob: "1999-05-13",
    gender: "P",
    created_at: ISODate("2022-03-19T01:29:29.965Z"),
    updated_at: ISODate("2022-03-19T01:29:29.965Z")
})
db.users.insert({
    _id:4,
    name: "Zikri",
    status: 1,
    dob: "2000-01-22",
    gender: "L",
    created_at: ISODate("2022-03-19T01:29:29.965Z"),
    updated_at: ISODate("2022-03-19T01:29:29.965Z")
})
db.users.insert({
    _id:5,
    name: "Natasha",
    status: 1,
    dob: "2000-11-29",
    gender: "P",
    created_at: ISODate("2022-03-19T01:29:29.965Z"),
    updated_at: ISODate("2022-03-19T01:29:29.965Z")
})

// Insert 3 Transaksi setiap User
db.transaction.insert({
    _id: 1,
    user_id: 1,
    payment_method_id: 3,
    status: "Belum Bayar",
    total_qty: 1,
    total_price: 10000,
    created_at: ISODate("2022-03-19T01:29:29.965Z"),
    updated_at: ISODate("2022-03-19T01:29:29.965Z") 
})
db.transaction.insert({
    _id: 2,
    user_id: 1,
    payment_method_id: 3,
    status: "Belum Bayar",
    total_qty: 2,
    total_price: 20000,
    created_at: ISODate("2022-03-19T01:29:29.965Z"),
    updated_at: ISODate("2022-03-19T01:29:29.965Z") 
})
db.transaction.insert({
    _id: 3,
    user_id: 1,
    payment_method_id: 3,
    status: "Belum Bayar",
    total_qty: 1,
    total_price: 100000,
    created_at: ISODate("2022-03-19T01:29:29.965Z"),
    updated_at: ISODate("2022-03-19T01:29:29.965Z") 
})
db.transaction.insert({
    _id: 4,
    user_id: 2,
    payment_method_id: 1,
    status: "Belum Bayar",
    total_qty: 1,
    total_price: 150000,
    created_at: ISODate("2022-03-19T01:29:29.965Z"),
    updated_at: ISODate("2022-03-19T01:29:29.965Z") 
})
db.transaction.insert({
    _id: 5,
    user_id: 2,
    payment_method_id: 1,
    status: "Belum Bayar",
    total_qty: 2,
    total_price: 40000,
    created_at: ISODate("2022-03-19T01:29:29.965Z"),
    updated_at: ISODate("2022-03-19T01:29:29.965Z") 
})
db.transaction.insert({
    _id: 6,
    user_id: 2,
    payment_method_id: 1,
    status: "Belum Bayar",
    total_qty: 1,
    total_price: 10000,
    created_at: ISODate("2022-03-19T01:29:29.965Z"),
    updated_at: ISODate("2022-03-19T01:29:29.965Z") 
})
db.transaction.insert({
    _id: 7,
    user_id: 3,
    payment_method_id: 3,
    status: "Belum Bayar",
    total_qty: 1,
    total_price: 10000,
    created_at: ISODate("2022-03-19T01:29:29.965Z"),
    updated_at: ISODate("2022-03-19T01:29:29.965Z") 
})
db.transaction.insert({
    _id: 8,
    user_id: 3,
    payment_method_id: 3,
    status: "Belum Bayar",
    total_qty: 1,
    total_price: 10000,
    created_at: ISODate("2022-03-19T01:29:29.965Z"),
    updated_at: ISODate("2022-03-19T01:29:29.965Z") 
})
db.transaction.insert({
    _id: 9,
    user_id: 3,
    payment_method_id: 3,
    status: "Belum Bayar",
    total_qty: 1,
    total_price: 100000,
    created_at: ISODate("2022-03-19T01:29:29.965Z"),
    updated_at: ISODate("2022-03-19T01:29:29.965Z") 
})
db.transaction.insert({
    _id: 10,
    user_id: 4,
    payment_method_id: 2,
    status: "Belum Bayar",
    total_qty: 1,
    total_price: 10000,
    created_at: ISODate("2022-03-19T01:29:29.965Z"),
    updated_at: ISODate("2022-03-19T01:29:29.965Z") 
})
db.transaction.insert({
    _id: 11,
    user_id: 4,
    payment_method_id: 3,
    status: "Belum Bayar",
    total_qty: 1,
    total_price: 20000,
    created_at: ISODate("2022-03-19T01:29:29.965Z"),
    updated_at: ISODate("2022-03-19T01:29:29.965Z") 
})
db.transaction.insert({
    _id: 12,
    user_id: 4,
    payment_method_id: 3,
    status: "Belum Bayar",
    total_qty: 1,
    total_price: 10000,
    created_at: ISODate("2022-03-19T01:29:29.965Z"),
    updated_at: ISODate("2022-03-19T01:29:29.965Z") 
})
db.transaction.insert({
    _id: 13,
    user_id: 5,
    payment_method_id: 3,
    status: "Belum Bayar",
    total_qty: 1,
    total_price: 10000,
    created_at: ISODate("2022-03-19T01:29:29.965Z"),
    updated_at: ISODate("2022-03-19T01:29:29.965Z") 
})
db.transaction.insert({
    _id: 14,
    user_id: 5,
    payment_method_id: 3,
    status: "Belum Bayar",
    total_qty: 3,
    total_price: 30000,
    created_at: ISODate("2022-03-19T01:29:29.965Z"),
    updated_at: ISODate("2022-03-19T01:29:29.965Z") 
})
db.transaction.insert({
    _id: 15,
    user_id: 5,
    payment_method_id: 3,
    status: "Belum Bayar",
    total_qty: 1,
    total_price: 100000,
    created_at: ISODate("2022-03-19T01:29:29.965Z"),
    updated_at: ISODate("2022-03-19T01:29:29.965Z") 
})

// Insert product pada masing-masing transaksi
db.transaction_detail.insert({
    _id: 1,
    transaction_id: 1,
    product_id: 3,
    status: "Belum Bayar",
    qty: 1,
    price: 10000,
    created_at: ISODate("2022-03-19T01:29:29.965Z"),
    updated_at: ISODate("2022-03-19T01:29:29.965Z") 
})
db.transaction_detail.insert({
    _id:2,
    transaction_id: 2,
    product_id: 6,
    status: "Belum Bayar",
    qty: 2,
    price: 10000,
    created_at: ISODate("2022-03-19T01:29:29.965Z"),
    updated_at: ISODate("2022-03-19T01:29:29.965Z") 
})
db.transaction_detail.insert({
    _id:3,
    transaction_id: 3,
    product_id: 1,
    status: "Belum Bayar",
    qty: 1,
    price: 100000,
    created_at: ISODate("2022-03-19T01:29:29.965Z"),
    updated_at: ISODate("2022-03-19T01:29:29.965Z") 
})
db.transaction_detail.insert({
    _id:4,
    transaction_id: 4,
    product_id: 2,
    status: "Belum Bayar",
    qty: 1,
    price: 150000,
    created_at: ISODate("2022-03-19T01:29:29.965Z"),
    updated_at: ISODate("2022-03-19T01:29:29.965Z") 
})
db.transaction_detail.insert({
    _id:5,
    transaction_id: 5,
    product_id: 2,
    status: "Belum Bayar",
    qty: 2,
    price: 20000,
    created_at: ISODate("2022-03-19T01:29:29.965Z"),
    updated_at: ISODate("2022-03-19T01:29:29.965Z") 
})
db.transaction_detail.insert({
    _id:6,
    transaction_id: 6,
    product_id: 3,
    status: "Belum Bayar",
    qty: 1,
    price: 10000,
    created_at: ISODate("2022-03-19T01:29:29.965Z"),
    updated_at: ISODate("2022-03-19T01:29:29.965Z") 
})
db.transaction_detail.insert({
    _id:7,
    transaction_id: 7,
    product_id: 6,
    status: "Belum Bayar",
    qty: 1,
    price: 10000,
    created_at: ISODate("2022-03-19T01:29:29.965Z"),
    updated_at: ISODate("2022-03-19T01:29:29.965Z") 
})
db.transaction_detail.insert({
    _id:8,
    transaction_id: 8,
    product_id: 3,
    status: "Belum Bayar",
    qty: 1,
    price: 10000,
    created_at: ISODate("2022-03-19T01:29:29.965Z"),
    updated_at: ISODate("2022-03-19T01:29:29.965Z") 
})
db.transaction_detail.insert({
    _id:9,
    transaction_id: 9,
    product_id: 1,
    status: "Belum Bayar",
    qty: 1,
    price: 100000,
    created_at: ISODate("2022-03-19T01:29:29.965Z"),
    updated_at: ISODate("2022-03-19T01:29:29.965Z") 
})
db.transaction_detail.insert({
    _id:10,
    transaction_id: 10,
    product_id: 3,
    status: "Belum Bayar",
    qty: 1,
    price: 10000,
    created_at: ISODate("2022-03-19T01:29:29.965Z"),
    updated_at: ISODate("2022-03-19T01:29:29.965Z") 
})
db.transaction_detail.insert({
    _id: 11,
    transaction_id: 11,
    product_id: 4,
    status: "Belum Bayar",
    qty: 1,
    price: 20000,
    created_at: ISODate("2022-03-19T01:29:29.965Z"),
    updated_at: ISODate("2022-03-19T01:29:29.965Z") 
})
db.transaction_detail.insert({
    _id:12,
    transaction_id: 12,
    product_id: 6,
    status: "Belum Bayar",
    qty: 1,
    price: 10000,
    created_at: ISODate("2022-03-19T01:29:29.965Z"),
    updated_at: ISODate("2022-03-19T01:29:29.965Z") 
})
db.transaction_detail.insert({
    _id:13,
    transaction_id: 13,
    product_id: 3,
    status: "Belum Bayar",
    qty: 1,
    price: 10000,
    created_at: ISODate("2022-03-19T01:29:29.965Z"),
    updated_at: ISODate("2022-03-19T01:29:29.965Z") 
})
db.transaction_detail.insert({
    _id:14,
    transaction_id: 14,
    product_id: 3,
    status: "Belum Bayar",
    qty: 3,
    price: 10000,
    created_at: ISODate("2022-03-19T01:29:29.965Z"),
    updated_at: ISODate("2022-03-19T01:29:29.965Z") 
})
db.transaction_detail.insert({
    _id:15,
    transaction_id: 15,
    product_id: 1,
    status: "Belum Bayar",
    qty: 1,
    price: 100000,
    created_at: ISODate("2022-03-19T01:29:29.965Z"),
    updated_at: ISODate("2022-03-19T01:29:29.965Z") 
})

// 2. Read
// a. tampilkan nama users yang gender : "L"
db.users.find({gender:"L"})

// b. tampilkan product dengan id = 3
db.product.find({_id: 3}).pretty()

// c. Hitung jumlah user / pelanggan dengan status gender perempuan
db.users.find({gender:"P"}).count()

// d. tampilkan data users sesuai abjad (ASC)
db.users.find().sort({name:1})

// e. tampilkan 5 data product
db.product.find().limit(5)

// 3. Update 
// a. ubah nama product id 1 jadi "product dummy"
db.product.update({_id:1}, {_id: 1,product_type_id: 1, operator_id: 3, code: "PTL-100", name: "Product Dummy", status:1, created_at: ISODate("2022-03-19T01:29:29.965Z"), updated_at: ISODate("2022-03-19T01:29:29.965Z")})

// b. ubah qty = 3 pada transaction detail = 1
db.transaction_detail.update({_id:1}, {
    transaction_id: 1,
    product_id: 3,
    status: "Belum Bayar",
    qty: 3,
    price: 10000,
    created_at: ISODate("2022-03-19T01:29:29.965Z"),
    updated_at: ISODate("2022-03-19T01:29:29.965Z") 
})

// 4. Delete
// a. Delete data pada product id 1
db.product.remove({_id:1})

// b. Delete pada product dengan product_type = 1
db.product.remove({product_type_id:1})