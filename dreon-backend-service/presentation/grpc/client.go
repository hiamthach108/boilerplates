package grpc

import "google.golang.org/grpc"

type InternalClient struct {
	conn *grpc.ClientConn
}

func NewInternalClientFromConn(conn *grpc.ClientConn) *InternalClient {
	return &InternalClient{conn: conn}
}

func (c *InternalClient) Conn() *grpc.ClientConn {
	return c.conn
}

func (c *InternalClient) Close() error {
	if c.conn == nil {
		return nil
	}
	return c.conn.Close()
}
