package performance

import (
	"os"
	"testing"

	"github.com/bytom/bytom/account"
	"github.com/bytom/bytom/mining"
	"github.com/bytom/bytom/test"
	dbm "github.com/bytom/bytom/database/leveldb"
)

// Function NewBlockTemplate's benchmark - 0.05s
func BenchmarkNewBlockTpl(b *testing.B) {
	testDB := dbm.NewDB("testdb", "leveldb", "temp")
	defer os.RemoveAll("temp")

	chain, _, txPool, err := test.MockChain(testDB)
	if err != nil {
		b.Fatal(err)
	}
	accountManager := account.NewManager(testDB, chain)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		mining.NewBlockTemplate(chain, txPool, accountManager)
	}
}
