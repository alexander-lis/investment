package infrastructure

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func AccessLogInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	start := time.Now()
	md, _ := metadata.FromIncomingContext(ctx)

	// Calls the handler
	reply, err := handler(ctx, req)

	var traceId, userId, userRole string
	if len(md["trace-id"]) > 0 {
		traceId = md["trace-id"][0]
	}
	if len(md["user-id"]) > 0 {
		userId = md["user-id"][0]
	}
	if len(md["user-role"]) > 0 {
		userRole = md["user-role"][0]
	}

	msg := fmt.Sprintf("Call:%v, traceId: %v, userId: %v, userRole: %v, time: %v", info.FullMethod, traceId, userId, userRole, time.Since(start))
	accesLog(msg)

	return reply, err

}

func accesLog(msg string) {
	filename := "/app/logs/access.log"
	writeLog(filename, msg)
}

func errLog(msg string) {
	filename := "/app/logs/error.log"
	writeLog(filename, msg)
}

func writeLog(filename, msg string) {
	f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("[ERR] error opening file: %v \n", err)
	}
	defer f.Close()

	log.SetOutput(f)
	log.Println(msg)
}
