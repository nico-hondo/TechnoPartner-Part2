create table categories(
	category_id smallint primary key,
	category_name varchar(15),
	description text
)

drop table categories

insert into categories(category_id, category_name, description)
values(1, 'Pemasukan', 'Pemasukan yang diterima user baik berupa Gaji, Tunjangan maupun Bonus'),
(2, 'Pengeluaran', 'Pengeluaran yang dilakukan oleh user baik berupa Sewa kost, Makan, Pakaian, Nonton Bioskop')

select * from categories

create table transaction(
	transact_id varchar(6) primary key,
	type_transact varchar(15),
	description text,
	nominal money,
	time_of_transaction date,
	category_id smallint
)

alter table transaction add constraint category_id_fk foreign key (category_id) 
references categories(category_id)

drop table transaction

insert into transaction(transact_id, type_transact, description, nominal, time_of_transaction, category_id)
values('t-0001', 'Gaji', 'Gaji Bulan Mei terhitung dari tanggal 28 april - 28 mei', 7500000, '2023-05-28', 1),
('t-0002', 'Bonus', 'Bonus Lembur pada periode 10 mei - 15 mei', 1000000, '2023-05-28', 1),
('t-0003', 'Sewa Kost', 'Pembayaran Sewa Kost Bulan 6', 1500000, '2023-05-30', 2),
('t-0004', 'Makan', 'Pembayaran Makan Warteg Bahari Putra', 20000, '2023-06-02', 2)

-- Menampilkan Total Pemasukan
select SUM(t.nominal) as total_income from transaction as t inner join categories as c 
on t.category_id = c.category_id where c.category_name = 'Pemasukan'

-- Menampilkan Total Pengeluaran
select SUM(t.nominal) as total_expensis from transaction as t inner join categories as c 
on t.category_id = c.category_id where c.category_name = 'Pengeluaran'

-- Menampilkan Saldo saat ini
with income_calc as(
select SUM(t.nominal) from transaction as t inner join categories as c 
on t.category_id = c.category_id where c.category_name = 'Pemasukan'
), expanses_calc as(
select SUM(t.nominal) from transaction as t inner join categories as c 
on t.category_id = c.category_id where c.category_name = 'Pengeluaran'
)
select income_calc - expanses_calc as current_saldo;


select SUM(t.nominal) from transaction as t inner join categories as c 
on t.category_id = c.category_id where c.category_name = 'Pemasukan'

select SUM(t.nominal) from transaction as t inner join categories as c 
on t.category_id = c.category_id where c.category_name = 'Pengeluaran'
