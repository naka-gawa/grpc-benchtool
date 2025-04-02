package server_test

import (
	"context"
	"io"
	"testing"
	"time"

	"github.com/naka-gawa/grpc-benchtool/internal/server"
	pb "github.com/naka-gawa/grpc-benchtool/proto/grpcbench"
	"google.golang.org/grpc"

	"github.com/stretchr/testify/assert"
)

type mockStream struct {
	grpc.ServerStream
	injectRecvErr error
	recvMsgs      []*pb.TestRequest
	recvIdx       int
	sentSummary   *pb.StreamSummary
}

func (m *mockStream) Recv() (*pb.TestRequest, error) {
	if m.recvIdx >= len(m.recvMsgs) {
		if m.injectRecvErr != nil {
			return nil, m.injectRecvErr
		}
		return nil, io.EOF
	}
	msg := m.recvMsgs[m.recvIdx]
	m.recvIdx++
	return msg, nil
}

func (m *mockStream) SendAndClose(summary *pb.StreamSummary) error {
	m.sentSummary = summary
	return nil
}

func (m *mockStream) RecvMsg(v interface{}) error {
	return nil
}

func (m *mockStream) SendMsg(v interface{}) error {
	return nil
}

func (m *mockStream) Context() context.Context {
	return context.Background()
}

func TestBenchHandler_UnaryTest(t *testing.T) {
	type args struct {
		clientID     string
		sentUnixNano int64
	}
	now := time.Now().UnixNano()

	tests := []struct {
		name     string
		args     args
		wantErr  bool
		validate func(*pb.TestResponse)
	}{
		{
			name: "normal latency within ~10ms",
			args: args{
				clientID:     "zunda",
				sentUnixNano: now - 10_000_000,
			},
			wantErr: false,
			validate: func(resp *pb.TestResponse) {
				assert.Equal(t, "test-server", resp.ServerId)
				assert.Greater(t, resp.ReceivedUnixNano, int64(0))
				assert.InDelta(t, 10_000_000, resp.LatencyNano, 5_000_000)
			},
		},
		{
			name: "sent in future (latency negative)",
			args: args{
				clientID:     "mirai",
				sentUnixNano: now + 20_000_000,
			},
			wantErr: false,
			validate: func(resp *pb.TestResponse) {
				assert.Less(t, resp.LatencyNano, int64(0))
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler := server.NewBenchHandler(&server.DefaultStrategy{ServerID: "test-server"})

			req := &pb.TestRequest{
				ClientId:     tt.args.clientID,
				SentUnixNano: tt.args.sentUnixNano,
			}
			got, err := handler.UnaryTest(context.Background(), req)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				tt.validate(got)
			}
		})
	}
}

func TestBenchHandler_StreamTest(t *testing.T) {
	now := time.Now().UnixNano()

	tests := []struct {
		name         string
		recvMsgs     []*pb.TestRequest
		injectErr    error
		expectError  bool
		expectCount  int64
		expectBytes  int64
		expectAvgMin float64 // ms 最小値だけチェック（最大や近似は誤差で unstable になる）
	}{
		{
			name: "3 messages with 10 bytes each",
			recvMsgs: []*pb.TestRequest{
				{ClientId: "z1", SentUnixNano: now - 5_000_000, Payload: make([]byte, 10)},
				{ClientId: "z2", SentUnixNano: now - 10_000_000, Payload: make([]byte, 10)},
				{ClientId: "z3", SentUnixNano: now - 15_000_000, Payload: make([]byte, 10)},
			},
			expectCount:  3,
			expectBytes:  30,
			expectAvgMin: 4.0, // ms
		},
		{
			name:         "no messages",
			recvMsgs:     []*pb.TestRequest{},
			expectCount:  0,
			expectBytes:  0,
			expectAvgMin: 0.0,
		},
		{
			name: "1 message with 100 bytes",
			recvMsgs: []*pb.TestRequest{
				{ClientId: "solo", SentUnixNano: now - 1_000_000, Payload: make([]byte, 100)},
			},
			expectCount:  1,
			expectBytes:  100,
			expectAvgMin: 0.5,
		},
		{
			name:         "stream.Recv returns unexpected error",
			recvMsgs:     []*pb.TestRequest{},
			injectErr:    io.ErrUnexpectedEOF,
			expectCount:  0,
			expectBytes:  0,
			expectAvgMin: 0.0,
			expectError:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler := server.NewBenchHandler(&server.DefaultStrategy{ServerID: "test-server"})

			stream := &mockStream{recvMsgs: tt.recvMsgs, injectRecvErr: tt.injectErr}

			err := handler.StreamTest(stream)
			if tt.expectError {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			if assert.NotNil(t, stream.sentSummary) {
				assert.Equal(t, tt.expectCount, stream.sentSummary.ReceivedCount)
				assert.Equal(t, tt.expectBytes, stream.sentSummary.TotalBytes)
				if tt.expectCount > 0 {
					assert.GreaterOrEqual(t, stream.sentSummary.AverageLatencyMs, tt.expectAvgMin)
				}
			}

			assert.NotNil(t, stream.sentSummary)
			assert.Equal(t, tt.expectCount, stream.sentSummary.ReceivedCount)
			assert.Equal(t, tt.expectBytes, stream.sentSummary.TotalBytes)

			if tt.expectCount > 0 {
				assert.GreaterOrEqual(t, stream.sentSummary.AverageLatencyMs, tt.expectAvgMin)
			}
		})
	}
}
