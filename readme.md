migrate  -database "postgres://postgres:(pass)@localhost:5432/meatwarehouse?sslmode=disable" -path db/migrations/  up

 migrate create -ext sql -dir db/migrations create_table_customers




 TODO
 report	
    -dailyExpenditureReport
	-transactionReport				                -> total_in - total_out
	-transactionAndExpenditureReport		        -> total_in - (total_out + expenditure)
	-unpaidOwnerReport (in)(kita beli daging)	    -> total hutang perusahaan
	-unpaidCustomersReport (out)(kita jual daging)	-> total hutang customers
	-paidOwnerReport (in)				            -> uang keluar perusahaan untuk daging
	-paidCustomersReport (out)		            	-> uang masuk ke perusahaan
	-