package raft

import (
	"distributed-algorithms/raft/context"
	"distributed-algorithms/raft/domain"
	"encoding/json"
	"log"
)

type RequestHandler struct {
	ctx *context.Context
}

func NewRequestHandler(ctx *context.Context) *RequestHandler {
	return &RequestHandler{ctx: ctx}
}

func (handler *RequestHandler) HandleRequestVoteRequest(request domain.RequestVoteRequest) (*domain.RequestVoteResponse, error) {
	// TODO: add checks about candidate's log

	ctx := handler.ctx

	a, _ := json.MarshalIndent(request, "", " ")
	log.Printf("Node \"%s\" received request of vote: %s\n", ctx.GetNodeId(), a)

	currentTerm := ctx.GetCurrentTerm()
	var voteGranted bool

	if request.Term < currentTerm {
		voteGranted = false
	} else {
		voteGranted = ctx.Vote(request.CandidateId)
	}

	if voteGranted {
		ctx.SetNewRandomElectionTimeout()
	}

	return &domain.RequestVoteResponse{Term: currentTerm, VoteGranted: voteGranted}, nil
}

func (handler *RequestHandler) HandleAppendEntriesRequest(request domain.AppendEntriesRequest) (*domain.AppendEntriesResponse, error) {
	// TODO: add checks related to log entries

	ctx := handler.ctx

	a, _ := json.MarshalIndent(request, "", " ")
	log.Printf("Node \"%s\" received request of append-entries: %s\n", ctx.GetNodeId(), a)

	currentTerm := ctx.GetCurrentTerm()
	requestTerm := request.Term
	success := requestTerm >= currentTerm

	if success {
		// Stable phase started
		ctx.SetNewRandomElectionTimeout()
		ctx.BecomeFollower()
		ctx.SetCurrentTerm(requestTerm)
	}

	return &domain.AppendEntriesResponse{Term: currentTerm, Success: success}, nil
}
