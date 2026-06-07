gen:
	@protoc \
	--proto_path=shared/api/proto \
	--go_out=shared/api/proto/gen --go_opt=paths=source_relative \
	--go-grpc_out=shared/api/proto/gen \
	--go-grpc_opt=paths=source_relative shared/api/proto/*.proto