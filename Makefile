BINARY_NAME=hotelreservation
build:
	@echo "**** Building the binary of the backend service... ****"
	go build -o bin/${BINARY_NAME} cmd/main.go
	@echo "**** Application building is succeed ! ****"

run:
	./bin/${BINARY_NAME}

test:
	go test -v ./...

# For migrations
#  docker run -i -v "H:\1- freelancing path\Courses\golang stack\projects\Hotel-Reservation-Backend\db\postgres\migrations:/migrations" --network host migrate/migrate -path=/migrations/ -database "postgresql://postgres:postgres@127.0.0.1:1111/hrdb?sslmode=disable" up 1


# For test migrations 
# docker run -i -v "H:\1- freelancing path\Courses\golang stack\projects\Hotel-Reservation-Backend\db\postgres\test_migrations:/migrations" --network host migrate/migrate -path=/migrations/ -database "postgresql://postgres:postgres@127.0.0.1:2222/testhrdb?sslmode=disable" up 1

# To update a dirty schema 
# UPDATE "public"."schema_migrations" SET "dirty" = 'false' WHERE "version" = 1;