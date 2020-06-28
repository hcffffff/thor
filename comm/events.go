// Copyright (c) 2018 The VeChainThor developers

// Distributed under the GNU Lesser General Public License v3.0 software license, see the accompanying
// file LICENSE or <https://www.gnu.org/licenses/lgpl-3.0.html>

package comm

import (
	"context"

	"github.com/vechain/thor/block"
	"github.com/vechain/thor/comm/proto"
)

// NewBlockEvent event emitted when received block announcement.
type NewBlockEvent struct {
	*block.Block
}

// NewBlockProposalEvent event emitted when received block proposal.
type NewBlockProposalEvent struct {
	*block.Proposal
}

// NewBackerSignatureEvent event emitted when received backer signature.
type NewBackerSignatureEvent struct {
	*proto.FullBackerSignature
}

// HandleBlockStream to handle the stream of downloaded blocks in sync process.
type HandleBlockStream func(ctx context.Context, stream <-chan *block.Block) error
