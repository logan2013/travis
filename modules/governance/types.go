package governance

import (
	"math/big"

	"github.com/CyberMiles/travis/utils"
	"github.com/ethereum/go-ethereum/common"
)

type Proposal struct {
	Id           string
	Proposer     common.Address
	BlockHeight  uint64
	From         common.Address
	To           common.Address
	Amount       *big.Int
	Reason       string
	CreatedAt    string
}

func NewProposal(id string, proposer common.Address, blockHeight uint64, from common.Address, to common.Address, amount *big.Int, reason string) *Proposal {
	now := utils.GetNow()
	return &Proposal {
		id,
		proposer,
		blockHeight,
		from,
		to,
		amount,
		reason,
		now,
	}
}

type Vote struct {
	ProposalId     string
	Voter          common.Address
	BlockHeight    uint64
	Answer         string
	CreatedAt      string
}

func NewVote(proposalId string, voter common.Address, blockHeight uint64, answer string) *Vote {
	now := utils.GetNow()
	return &Vote {
		proposalId,
		voter,
		blockHeight,
		answer,
		now,
	}
}
