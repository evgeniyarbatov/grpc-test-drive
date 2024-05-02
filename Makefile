gencode:
	protoc \
    --go_out=client \
    --go_opt=paths=source_relative \
    --go-grpc_out=client \
    --go-grpc_opt=paths=source_relative \
    protos/csv-parser.proto

	python3.12 -m grpc_tools.protoc \
    -I protos \
    --python_out=server \
    --pyi_out=server \
    --grpc_python_out=server \
    protos/csv-parser.proto