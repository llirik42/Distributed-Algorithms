package grpc

import (
	"context"
	pb "distributed-algorithms/generated/proto"
	"distributed-algorithms/src/raft/domain"
	"distributed-algorithms/src/raft/transport"
	"fmt"
	"google.golang.org/grpc"
	"net"
)

type Server struct {
	pb.RaftServiceServer
	gRPCServer                  *grpc.Server
	listener                    net.Listener
	handleRequestForVoteRequest transport.HandleRequestForVoteRequest
	handleAppendEntriesRequest  transport.HandleAppendEntriesRequest
}

func (server *Server) RequestForVote(
	_ context.Context,
	request *pb.RequestVoteRequest,
) (*pb.RequestVoteResponse, error) {
	result, err := server.handleRequestForVoteRequest(&domain.RequestVoteRequest{
		Term:         request.Term,
		CandidateId:  request.CandidateId,
		LastLogIndex: request.LastLogIndex,
		LastLogTerm:  request.LastLogTerm,
	})

	if err != nil {
		return nil, fmt.Errorf("failed to handle request for vote %s: %w", request.String(), err)
	}

	return &pb.RequestVoteResponse{Term: result.Term, VoteGranted: result.VoteGranted}, nil
}

func (server *Server) AppendEntries(
	_ context.Context,
	request *pb.AppendEntriesRequest,
) (*pb.AppendEntriesResponse, error) {
	result, err := server.handleAppendEntriesRequest(&domain.AppendEntriesRequest{
		Term:         request.Term,
		LeaderId:     request.LeaderId,
		PrevLogIndex: request.PrevLogIndex,
		PrevLogTerm:  request.PrevLogTerm,
		LeaderCommit: request.LeaderCommit,
	})

	if err != nil {
		return nil, fmt.Errorf("failed to handle append-entries request %s: %w", request.String(), err)
	}

	return &pb.AppendEntriesResponse{Term: result.Term, Success: result.Success}, nil
}

func (server *Server) Listen() error {
	if err := server.gRPCServer.Serve(server.listener); err != nil {
		return fmt.Errorf("failed to serve gRPC-server: %w", err)
	}

	return nil
}

func (server *Server) Shutdown() error {
	server.gRPCServer.GracefulStop()
	return nil
}
