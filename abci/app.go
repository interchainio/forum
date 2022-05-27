package abci

import (
	"context"

	"github.com/interchainio/forum/model"
	abci "github.com/tendermint/tendermint/abci/types"
)

var _ abci.Application = App{}

type App struct {
	DB *model.DB
}

// Info/Query Connection
// Return application info
func (App) Info(_ context.Context, _ *abci.RequestInfo) (*abci.ResponseInfo, error) {
	panic("not implemented") // TODO: Implement
}

func (App) Query(_ context.Context, _ *abci.RequestQuery) (*abci.ResponseQuery, error) {
	panic("not implemented") // TODO: Implement
}

// Mempool Connection
// Validate a tx for the mempool
func (App) CheckTx(_ context.Context, _ *abci.RequestCheckTx) (*abci.ResponseCheckTx, error) {
	panic("not implemented") // TODO: Implement
}

// Consensus Connection
// Initialize blockchain w validators/other info from TendermintCore
func (App) InitChain(_ context.Context, _ *abci.RequestInitChain) (*abci.ResponseInitChain, error) {
	panic("not implemented") // TODO: Implement
}

func (App) PrepareProposal(_ context.Context, _ *abci.RequestPrepareProposal) (*abci.ResponsePrepareProposal, error) {
	panic("not implemented") // TODO: Implement
}

func (App) ProcessProposal(_ context.Context, _ *abci.RequestProcessProposal) (*abci.ResponseProcessProposal, error) {
	panic("not implemented") // TODO: Implement
}

// Commit the state and return the application Merkle root hash
func (App) Commit(_ context.Context) (*abci.ResponseCommit, error) {
	panic("not implemented") // TODO: Implement
}

// Create application specific vote extension
func (App) ExtendVote(_ context.Context, _ *abci.RequestExtendVote) (*abci.ResponseExtendVote, error) {
	panic("not implemented") // TODO: Implement
}

// Verify application's vote extension data
func (App) VerifyVoteExtension(_ context.Context, _ *abci.RequestVerifyVoteExtension) (*abci.ResponseVerifyVoteExtension, error) {
	panic("not implemented") // TODO: Implement
}

// Deliver the decided block with its txs to the Application
func (App) FinalizeBlock(_ context.Context, _ *abci.RequestFinalizeBlock) (*abci.ResponseFinalizeBlock, error) {
	panic("not implemented") // TODO: Implement
}

// State Sync Connection
// List available snapshots
func (App) ListSnapshots(_ context.Context, _ *abci.RequestListSnapshots) (*abci.ResponseListSnapshots, error) {
	panic("not implemented") // TODO: Implement
}

func (App) OfferSnapshot(_ context.Context, _ *abci.RequestOfferSnapshot) (*abci.ResponseOfferSnapshot, error) {
	panic("not implemented") // TODO: Implement
}

func (App) LoadSnapshotChunk(_ context.Context, _ *abci.RequestLoadSnapshotChunk) (*abci.ResponseLoadSnapshotChunk, error) {
	panic("not implemented") // TODO: Implement
}

func (App) ApplySnapshotChunk(_ context.Context, _ *abci.RequestApplySnapshotChunk) (*abci.ResponseApplySnapshotChunk, error) {
	panic("not implemented") // TODO: Implement
}
