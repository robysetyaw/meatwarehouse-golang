migrate  -database "postgres://postgres:(pass)@localhost:5432/meatwarehouse?sslmode=disable" -path db/migrations/  up

 migrate create -ext sql -dir db/migrations create_table_customers
=================================================================================================================================
 TODO
 report	
 	-dailyExpenditureReport									-> daily_expenditure report OK
	-itemStock												-> keluar masuk stock		-

==================================================================================================================================

1. Laporan Penjualan (Sales Report)
	Laporan ini akan memberikan gambaran tentang total penjualan yang dilakukan dalam periode tertentu. Ini dapat dihasilkan dengan menjumlahkan total penjualan (total) dari semua entri transaksi dengan tx_type yang sesuai (misalnya, "out") dan mengelompokkannya berdasarkan tanggal (date) atau pelanggan (customer_id).(tx_type_= out, total)

2. Laporan Penerimaan (Receipt Report)
	Laporan ini akan menampilkan total penerimaan dari transaksi pembayaran yang dilakukan. Dapat dihasilkan dengan menjumlahkan total penerimaan (payment_amount) dari semua entri transaksi dengan payment_status yang sesuai (misalnya, "paid") dan mengelompokkannya berdasarkan tanggal (date) atau pelanggan (customer_id). (tx_type = out, paid, payment_amount)

3. Laporan Utang (Debt/Accounts Payable Report)
	Laporan ini akan memberikan gambaran tentang utang yang masih harus dibayar. Dapat dihasilkan dengan menjumlahkan total utang (total - payment_amount) dari semua entri transaksi dengan payment_status yang sesuai (misalnya, "unpaid") dan mengelompokkannya berdasarkan tanggal (date) atau pelanggan (customer_id). (tx_type = in, unpaid, total-payment_amount)

4. Laporan Laba Rugi (Profit and Loss Statement/Income Statement)
	Laporan ini akan menghitung laba atau rugi bersih berdasarkan total pendapatan (penjualan) dikurangi dengan total biaya atau pengeluaran. Untuk menghasilkan laporan ini, perlu melibatkan tabel lain yang berisi rincian biaya atau pengeluaran yang terkait dengan transaksi tersebut. (total_out-(total_in + daily_expenditures), tx_in, tx_out, daily_expenditures)

5. Laporan Arus Kas (Cash Flow Statement)
	Laporan ini akan memberikan informasi tentang aliran kas masuk dan keluar selama periode waktu tertentu. Dapat dihasilkan dengan menjumlahkan total penerimaan (payment_amount) dan menguranginya dengan total pengeluaran atau biaya dari transaksi dengan tx_type yang sesuai (misalnya, "in" untuk kas masuk dan "out" untuk kas keluar), kemudian mengelompokkannya berdasarkan tanggal (date).
	((payment_amount,tx_type = out) - ((payment_amount, tx_type = in) + daily_expenditures))

6. Laporan Rekapitulasi (Summary Report/Consolidated Report)
	Laporan ini memberikan ringkasan informasi tentang total penjualan, penerimaan, utang, laba, dan arus kas selama periode tertentu. Laporan ini akan menyajikan data secara agregat dengan menghitung total dari masing-masing kategori dalam satu laporan yang lengkap.

=====================================================================================================================================