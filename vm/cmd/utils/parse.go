package utils

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"reflect"
	"github.com/ethereum/go-ethereum/core"
	gen "github.com/CyberMiles/travis/misc/genesis"
)

// defaultGenesisBlob is the JSON representation of the default
// genesis file in $GOPATH/src/github.com/tendermint/ethermint/setup/simulate_genesis.json
// nolint=lll
var defaultGenesisBlob = []byte(`
{
    "config": {
        "chainId": 15,
        "homesteadBlock": 0,
        "eip155Block": 0,
        "eip158Block": 0
    },
    "nonce": "0xdeadbeefdeadbeef",
    "timestamp": "0x00",
    "parentHash": "0x0000000000000000000000000000000000000000000000000000000000000000",
    "mixhash": "0x0000000000000000000000000000000000000000000000000000000000000000",
    "difficulty": "0x40",
    "gasLimit": "0xF00000000",
    "alloc": {
        "0x7eff122b94897ea5b0e2a9abf47b86337fafebdc": { "balance": "10000000000000000000000000000000000" },
	    "0x77beb894fc9b0ed41231e51f128a347043960a9d": { "balance": "10000000000000000000000000000000000" }
    }
}`)

var blankGenesis = new(core.Genesis)

var errBlankGenesis = errors.New("could not parse a valid/non-blank Genesis")

// ParseGenesisOrDefault tries to read the content from provided
// genesisPath. If the path is empty or doesn't exist, it will
// use defaultGenesisBytes as the fallback genesis source. Otherwise,
// it will open that path and if it encounters an error that doesn't
// satisfy os.IsNotExist, it returns that error.
func ParseGenesisOrDefault(genesisPath string) (*core.Genesis, error) {
	//var genesisBlob = defaultGenesisBlob[:]
	genesisBlob, err := gen.DevGenesisBlock().MarshalJSON();
	//genesisBlob, err := gen.SimulateGenesisBlock().MarshalJSON();
	if err != nil {
		return nil, err
	}
	if len(genesisPath) > 0 {
		blob, err := ioutil.ReadFile(genesisPath)
		if err != nil && !os.IsNotExist(err) {
			return nil, err
		}
		if len(blob) >= 2 { // Expecting atleast "{}"
			genesisBlob = blob
		}
	}

	genesis := new(core.Genesis)
	if err := json.Unmarshal(genesisBlob, genesis); err != nil {
		return nil, err
	}

	if reflect.DeepEqual(blankGenesis, genesis) {
		return nil, errBlankGenesis
	}

	return genesis, nil
}
