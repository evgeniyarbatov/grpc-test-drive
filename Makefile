gencode:
	protoc \
    --go_out=client \
    --go_opt=paths=source_relative \
    --go-grpc_out=client \
    --go-grpc_opt=paths=source_relative \
    csvparser/csvparser.proto

	source ~/.venv/bin/activate && (\
		python3.12 -m grpc_tools.protoc \
		-I csvparser \
		--python_out=server \
		--pyi_out=server \
		--grpc_python_out=server \
		csvparser/csvparser.proto \
	)

