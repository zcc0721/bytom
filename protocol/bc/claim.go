package bc

import "io"

func (Claim) typ() string { return "claim1" }
func (c *Claim) writeForHash(w io.Writer) {
	mustWriteForHash(w, c.Peginwitness)
}

// SetDestination is support function for map tx
func (c *Claim) SetDestination(id *Hash, val *AssetAmount, pos uint64) {
	c.WitnessDestination = &ValueDestination{
		Ref:      id,
		Value:    val,
		Position: pos,
	}
}

/*
// NewCoinbase creates a new Coinbase.
func NewCoinbase(arbitrary []byte) *Coinbase {
	return &Coinbase{Arbitrary: arbitrary}
}
*/
func NewClaim(controlProgram *Program, ordinal uint64, peginwitness [][]byte) *Claim {
	return &Claim{
		ControlProgram: controlProgram,
		Ordinal:        ordinal,
		Peginwitness:   peginwitness,
	}
}
