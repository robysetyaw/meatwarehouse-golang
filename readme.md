migrate  -database "postgres://postgres:(pass)@localhost:5432/meatwarehouse?sslmode=disable" -path db/migrations/  up

 migrate create -ext sql -dir db/migrations create_table_customers

 https://www.postman.com/gold-escape-915482/workspace/chillabez/collection/18178220-ce916f8f-cf61-451c-aff4-82ed583dc452?action=share&creator=18178220