-- create database
CREATE DATABASE `alta_online_shop`;
use alta_online_shop;

-- create table
CREATE TABLE `user` (
  id int(11) AUTO_INCREMENT primary key,
  name varchar(255),
  alamat_id int(11)
);
CREATE TABLE `alamat` (
  id int(11) primary key,
  alamat varchar(255),
  kota varchar(255),
  kode_pos varchar(55)
);
CREATE TABLE `product`(
  id int(11) auto_increment primary key,
  product varchar(255),
  product_description varchar(255),
  product_type varchar(255),
  operator varchar(55),
  payment_method_id int(11)
);
CREATE TABLE `payment_method` (
  payment_method_id int(11) primary key,
  description varchar(255)
);
CREATE TABLE `user_payment_method` (
 payment_method_id int(11),
  user_id int(11)
);
CREATE TABLE `transaction`(
  id int(11) auto_increment primary key,
  product_id int(11),
  user_id int(11),
  total float
);
CREATE TABLE `kurir`(
  id int(11) auto_increment primary key,
  name varchar(255),
  created_at datetime,
  updated_at datetime
);

-- Tambah ongkos dasar 
alter table `kurir` add column ongkos_dasar int(11) after name;

-- Rename table kurir menjadi shipping
alter table `kurir` rename to `shipping`;

-- Drop Table shipping
drop table shipping;

-- Relation 
-- one to one
ALTER TABLE `product` ADD FOREIGN KEY (`payment_method_id`) REFERENCES `payment_method` (`payment_method_id`);
-- one to many 
ALTER TABLE `user` ADD FOREIGN KEY (`alamat_id`) REFERENCES `alamat` (`id`);
-- many to many
ALTER TABLE `user_payment_method` ADD FOREIGN KEY (`user_id`) REFERENCES `user` (`id`);
ALTER TABLE `user_payment_method` ADD FOREIGN KEY (`payment_method_id`) REFERENCES `payment_method` (`payment_method_id`);