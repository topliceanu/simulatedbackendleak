// You can edit this code!
// Click here and start typing.
package main

import (
	"testing"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient/simulated"
	"go.uber.org/goleak"
)

func TestSimulatedBackendLeak(t *testing.T) {
	defer goleak.VerifyNone(t)
	genesisData := types.GenesisAlloc{}
	simulatedBackend := simulated.NewBackend(genesisData)
	defer simulatedBackend.Close()

}
