from concurrent import futures
import logging

import grpc
import csv_parser_pb2
import csv_parser_pb2_grpc

class CSVParser(csv_parser_pb2_grpc.CSVParserServicer):
    def CountRows(self, request, context):
        return csv_parser_pb2.CSVParserResponse(rowCount=1)

def serve():
    port = "50051"
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    csv_parser_pb2_grpc.add_CSVParserServicer_to_server(CSVParser(), server)
    server.add_insecure_port("[::]:" + port)
    server.start()
    print("Server started, listening on " + port)
    server.wait_for_termination()


if __name__ == "__main__":
    logging.basicConfig()
    serve()