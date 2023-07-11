migrate  -database "postgres://postgres:(pass)@localhost:5432/meatwarehouse?sslmode=disable" -path db/migrations/  up

 migrate create -ext sql -dir db/migrations create_table_customers