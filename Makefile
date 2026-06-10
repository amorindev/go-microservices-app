gen-categories:
	@protoc \
	--proto_path=shared/categories/api/grpc/proto \
	--go_out=shared/categories/api/grpc/gen --go_opt=paths=source_relative \
	--go-grpc_out=shared/categories/api/grpc/gen \
	--go-grpc_opt=paths=source_relative shared/categories/api/grpc/proto/*.proto

gen-orders:
	@protoc \
	--proto_path=shared/orders/api/grpc/proto \
	--go_out=shared/orders/api/grpc/gen --go_opt=paths=source_relative \
	--go-grpc_out=shared/orders/api/grpc/gen \
	--go-grpc_opt=paths=source_relative shared/orders/api/grpc/proto/*.proto

gen-inventory:
	@protoc \
	--proto_path=shared/inventory/api/grpc/proto \
	--go_out=shared/inventory/api/grpc/gen --go_opt=paths=source_relative \
	--go-grpc_out=shared/inventory/api/grpc/gen \
	--go-grpc_opt=paths=source_relative shared/inventory/api/grpc/proto/*.proto

run-orders:
	go run ./orders-srv

run-payments:
	go run ./payments-srv

run-inventory:
	go run ./inventory-srv

run-catalog:
	go run ./catalog-srv

run-gateway:
	go run ./gateway

run:
	make -j4 run-orders run-payments run-inventory run-catalog run-gateway