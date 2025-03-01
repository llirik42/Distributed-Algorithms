package grpc

import (
	"context"
	"distributed-algorithms/dto"
	pb "distributed-algorithms/generated/proto"
	"distributed-algorithms/transport"
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

func (server *Server) RequestForVote(_ context.Context, request *pb.RequestVoteRequest) (*pb.RequestVoteResponse, error) {
	result, err := server.handleRequestForVoteRequest(dto.RequestVoteRequest{
		Term:         request.Term,
		CandidateId:  request.CandidateId,
		LastLogIndex: request.LastLogIndex,
		LastLogTerm:  request.LastLogTerm,
	})

	if err != nil {
		return nil, err
	}

	return &pb.RequestVoteResponse{Term: result.Term, VoteGranted: result.VoteGranted}, nil
}

func (server *Server) AppendEntries(_ context.Context, request *pb.AppendEntriesRequest) (*pb.AppendEntriesResponse, error) {
	result, err := server.handleAppendEntriesRequest(dto.AppendEntriesRequest{
		Term:         request.Term,
		LeaderId:     request.LeaderId,
		PrevLogIndex: request.PrevLogIndex,
		PrevLogTerm:  request.PrevLogTerm,
		LeaderCommit: request.LeaderCommit,
	})

	if err != nil {
		return nil, err
	}

	return &pb.AppendEntriesResponse{Term: result.Term, Success: result.Success}, nil
}

func (server *Server) Listen() error {
	return server.gRPCServer.Serve(server.listener)
}
