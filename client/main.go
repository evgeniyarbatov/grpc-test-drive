package main

import (
	"context"
	"flag"
	"log"
	"time"

	pb "github.com/evgeniyarbatov/grpc-test-drive/csvparser"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	defaultColumnName = "id"
)

var (
	addr       = flag.String("addr", "localhost:50051", "the address to connect to")
	columnName = flag.String("columnName", defaultColumnName, "Column name to use")
)

func main() {
	flag.Parse()

	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewCSVParserClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.CountRows(ctx, &pb.CSVParserRequest{ColumnName: *columnName})
	if err != nil {
		log.Fatalf("Could not count rows: %v", err)
	}
	log.Printf("Greeting: %d", r.GetRowCount())
}
