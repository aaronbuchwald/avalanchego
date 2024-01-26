package validators

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/ava-labs/avalanchego/ids"
)

func TestDummy(t *testing.T) {
	nodeID, err := ids.NodeIDFromString("NodeID-zbFM8qH7MnQ8uo6rm4tZaq6vhbF4cpgt")
	if err != nil {
		t.Fatal(err)
	}
	if ExpectedVdrSetQ[nodeID].Weight != 890344773985347 {
		t.Fatal("unexpected weight")
	}

	res, err := json.Marshal(ExpectedVdrSetQ)
	if err != nil {
		t.Fatal(err)
	}

	t.Fatal(fmt.Errorf("Actual response:%s\n\n\nExpected response: %s", string(res), string(res)))
}
