from concurrent import futures
import logging

import grpc
import csvparser_pb2
import csvparser_pb2_grpc

import pandas as pd

df = pd.read_csv("data/data.csv")

class CSVParser(csvparser_pb2_grpc.CSVParserServicer):
    def CountRows(self, request, context):
        row_count = df[request.columnName].nunique()
        return csvparser_pb2.CSVParserResponse(rowCount=row_count)

def serve():
    port = "50051"
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    csvparser_pb2_grpc.add_CSVParserServicer_to_server(CSVParser(), server)
    server.add_insecure_port("[::]:" + port)
    server.start()
    print("Server started, listening on " + port)
    server.wait_for_termination()

if __name__ == "__main__":
    logging.basicConfig()
    serve()