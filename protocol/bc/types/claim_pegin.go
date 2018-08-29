package types

import "github.com/bytom/protocol/bc"

type ClaimCommitment struct {
	bc.AssetAmount
	VMVersion      uint64
	ControlProgram []byte
}

// ClaimInput satisfies the TypedInput interface and represents a spend transaction.
type ClaimInput struct {
	SpendCommitmentSuffix []byte   // The unconsumed suffix of the output commitment
	Arguments             [][]byte // Witness
	ClaimCommitment
}

// NewClaimInputInput create a new SpendInput struct.
func NewClaimInputInput(arguments [][]byte, assetID bc.AssetID, amount uint64, controlProgram []byte) *TxInput {
	sc := ClaimCommitment{
		AssetAmount: bc.AssetAmount{
			AssetId: &assetID,
			Amount:  amount,
		},
		VMVersion:      1,
		ControlProgram: controlProgram,
	}
	return &TxInput{
		AssetVersion: 1,
		TypedInput: &ClaimInput{
			ClaimCommitment: sc,
			Arguments:       arguments,
		},
	}
}

// InputType is the interface function for return the input type.
func (si *ClaimInput) InputType() uint8 { return ClainPegin }
