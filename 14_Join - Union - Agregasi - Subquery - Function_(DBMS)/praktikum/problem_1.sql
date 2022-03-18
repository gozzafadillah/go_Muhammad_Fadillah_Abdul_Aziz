-- 1. Insert
-- a. insert 5 operator
INSERT INTO operators VALUES (
    '1','Telkomsel', NOW(), NOW()
),(
    '2','Indosat', NOW(), NOW()
),(
    '3','XL', NOW(), NOW()
),(
    '4','PLN', NOW(), NOW()
),(
    '5','Axis', NOW(), NOW()
);

-- b. Insert product type
INSERT INTO product_type VALUES 
    ('1', 'Pulsa', NOW(), NOW()),
    ('2', 'Paket Data', NOW(), NOW()),
    ('3', 'Token Listrik', NOW(), NOW())
;

-- c. Insert 2 product ; type id = 1 ; operator id = 3
INSERT INTO product VALUES (
    '1', '1', '3', 'XL001', 'Pulsa 10.000', 1, NOW(), NOW()
),(
    '2', '1', '3', 'XL002', 'Pulsa 20.000', 1, NOW(), NOW()
);

-- d. insert 3 product ; type id = 2 ; operator id = 1
INSERT INTO product VALUES (
    '3', '2', '1', 'TL001', 'Paket Data 1,5 GB', 1, NOW(), NOW()
),(
    '4', '2', '1', 'TL002', 'Paket Data 2,5 GB', 1, NOW(), NOW()
),(
    '5', '2', '1', 'TL003', 'Paket Data 4 GB', 1, NOW(), NOW()
);

-- d. insert 3 product ; type id = 3 ; operator id = 4
INSERT INTO product VALUES (
    '6', '3', '4', 'TL001', 'Token Listrik 100rb', 1, NOW(), NOW()
),(
    '7', '3', '4', 'TL002', 'Token Listrik 150rb', 1, NOW(), NOW()
),(
    '8', '3', '4', 'TL003', 'Token Listrik 300rb', 1, NOW(), NOW()
);

-- e. Insert Product description pada semua product
INSERT INTO product_description values (
    '1', 'Pulsa Rp10.000 RB', NOW(), NOW()
), (
    '2', 'Pulsa Rp20.000 RB', NOW(), NOW()
),(
    '3', 'Paket Data 1,5 GB 24 Jam', NOW(), NOW()
),(
    '4', 'Paket Data 2 GB 24 Jam + 500 GB Kuota Chatting seumur hidup', NOW(), NOW()
),(
    '5', 'Paket Data 3 GB 24 Jam + 1 GB Kuota Chatting seumur hidup', NOW(), NOW()
),(
    '6', 'Token Listrik untuk 75 kwh', NOW(), NOW()
),(
    '7', 'Token Listrik untuk 155 kwh', NOW(), NOW()
),(
    '8', 'Token Listrik untuk 285 kwh', NOW(), NOW()
);

-- f. Insert 3 payment method
INSERT INTO payment_method VALUES (
    '1', "Dana", 1, NOW(), NOW()
),(
    '2', 'Gerai Indomaret', 1, NOW(),NOW()
),(
    '3', 'ShopePay', 1, NOW(),NOW()
);

-- g. Insert 5 User pada table user
INSERT INTO users VALUES(
    '1', 1, "2000-03-09","L",NOW(),NOW()
),(
    '2', 1, "1991-05-19","P",NOW(),NOW()
),(
    '3', 1, "2001-12-10","L",NOW(),NOW()
),(
    '4', 1, "1999-09-09","P",NOW(),NOW()
),(
    '5', 1, "1990-01-09","L",NOW(),NOW()
);

-- h. Insert 3 Transaksi dimasing-masing user
INSERT INTO transaction VALUES(
    '1', '1', '2', 'Belum Bayar', 1, 20000, NOW(), NOW()
),(
    '2', '1', '2', 'Belum Bayar', 1, 150000, NOW(), NOW()
),(
    '3', '1', '2', 'Belum Bayar', 1, 20000, NOW(), NOW()
),(
    '4', '2', '2', 'Belum Bayar', 1, 20000, NOW(), NOW()
),(
    '5', '2', '2', 'Sudah Bayar', 1, 100000, NOW(), NOW()
),(
    '6', '2', '2', 'Belum Bayar', 1, 50000, NOW(), NOW()
),(
    '7', '3', '2', 'Belum Bayar', 1, 20000, NOW(), NOW()
),(
    '8', '3', '2', 'Belum Bayar', 1, 150000, NOW(), NOW()
),(
    '9', '3', '2', 'Belum Bayar', 1, 20000, NOW(), NOW()
),(
    '10', '4', '2', 'Sudah Bayar', 1, 20000, NOW(), NOW()
),(
    '11', '4', '2', 'Belum Bayar', 1, 300000, NOW(), NOW()
),(
    '12', '4', '2', 'Belum Bayar', 1, 20000, NOW(), NOW()
),(
    '13', '5', '2', 'Sudah Bayar', 1, 20000, NOW(), NOW()
),(
    '14', '5', '2', 'Belum Bayar', 1, 150000, NOW(), NOW()
),(
    '15', '5', '2', 'Sudah Bayar', 2, 50000, NOW(), NOW()
);

-- i. Insert 3 product di masing-masing transaksi
INSERT INTO transaction_detail VALUES (
    '2','1', "Belum Bayar", 1, 20000, NOW(),NOW()       
),(
    '7','2', "Belum Bayar", 1, 150000, NOW(),NOW()       
),(
    '3','3', "Belum Bayar", 1, 20000, NOW(),NOW()       
),(
    '2','4', "Belum Bayar", 1, 20000, NOW(),NOW()       
),(
    '6','5', "Belum Bayar", 1, 100000, NOW(),NOW()       
),(
    '5','6', "Belum Bayar", 1, 50000, NOW(),NOW()       
),(
    '2','7', "Belum Bayar", 1, 20000, NOW(),NOW()       
),(
    '7','8', "Belum Bayar", 1, 150000, NOW(),NOW()       
),(
    '3','9', "Belum Bayar", 1, 20000, NOW(),NOW()       
),(
    '2','10', "Belum Bayar", 1, 20000, NOW(),NOW()       
),(
    '4','11', "Belum Bayar", 1, 30000, NOW(),NOW()       
),(
    '3','12', "Belum Bayar", 1, 20000, NOW(),NOW()       
),(
    '2','13', "Belum Bayar", 1, 20000, NOW(),NOW()       
),(
    '7','14', "Belum Bayar", 1, 150000, NOW(),NOW()       
),(
    '5','15', "Belum Bayar", 1, 20000, NOW(),NOW()       
);

-- 2. Select
-- a. Tampilkan nama user dengan gender Laki-laki (L)
SELECT nama FROM users WHERE gender = "L";
-- b. Tampilkan Product id = 3
SELECT * FROM product WHERE id = 3;
-- c. Tampilkan data pelanggan yang memiliki nama mengandung a dan crated_at 7 hari kebelakang
SELECT * FROM users WHERE 
NOW() > created_at AND nama LIKE '%a%';
-- d. Hitung jumlah gender perempuan
 SELECT COUNT(*) FROM users WHERE gender = "P";
--  e. Tampilkan 5 data pada product
SELECT * FROM product LIMIT 5;


-- 3. Update
-- a. Ubah data pada id = 1 dengan name = product dummy
UPDATE product SET name = "product dummy" WHERE id = 1;
-- b. ubah data pada id = 1 dengan qty = 3
UPDATE transaction_detail SET qty = 3 WHERE product_id = 1;


-- 4. Delete
-- a. Delete pada table product id = 1
DELETE FROM product WHERE id = 1;
-- b. Delete pada table product dengan product type = 1
DELETE FROM product WHERE product_type = 1;

-- JOIN, UNION, Sub-Query, and  Function
-- 1. Gabungkan data transaksi dari user id 1 dan user id 2
SELECT * FROM transaction WHERE user_id = 1 
UNION 
SELECT * FROM transaction WHERE user_id = 2;
-- 2. Tampilkan Jumlah harga transaksi user id 1
SELECT sum(total_price) as `Total harga user id = 1` FROM transaction WHERE user_id = 1;
-- 3.Tampilkan total transaksi dengan product type = 2
SELECT COUNT(*) FROM product
JOIN transaction_detail
ON transaction_detail.product_id = product.id
WHERE product.product_type_id = 2;
-- 4. Tampilkan semua field tablle product dan field name table product type yang saling berhubungan
SELECT * FROM product
INNER JOIN product_type
ON product.product_type_id = product_type.id;
-- 5. Tampilkan semua field table transaction, field name table product dan field name table user.
SELECT product.name, transaction.*, users.name  FROM transaction
INNER JOIN transaction_detail 
ON transaction_detail.transaction_id = transaction.id
INNER JOIN product 
ON transaction_detail.product_id = product.id
INNER JOIN users
ON transaction.user_id = users.id;
-- 6. buat function setelah data transaksi dihapus maka transaction detail terhapus juga dengan transaction id yang dimaksud.

-- 7. buat function setelah data transaksi detail dihapus maka data total_qty terupdate berdasarkan qty data transaction id yang dihapus

-- 8. tampilkan data product yang tidak pernah ada ditable transaction_detail dengan sub-query

