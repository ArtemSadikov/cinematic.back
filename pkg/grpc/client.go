package grpc

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

type Client interface {
	OpenConn(ctx context.Context) *grpc.ClientConn
}

type client struct {
	host string
	conn *grpc.ClientConn
}

func (c client) OpenConn(ctx context.Context) *grpc.ClientConn {
	dialCtx, cancel := context.WithTimeout(ctx, time.Second*15)
	defer cancel()

	cc, err := grpc.DialContext(
		dialCtx,
		c.host,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	if err != nil {
		log.Fatalf("Failed to start gRPC connection: %v", err)
	}

	c.conn = cc

	return c.conn
}

func NewClient(host string) Client {
	if host == "" {
		log.Fatal("No host provided")
	}

	log.Printf("Service client created for host \"%s\"\n", host)

	return &client{host, nil}
}
