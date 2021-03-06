//This file is generated by btsgen. DO NOT EDIT.
//operation sample data for OperationTypeWitnessCreate

package samples

import (
	"github.com/denkhaus/bitshares/gen/data"
	"github.com/denkhaus/bitshares/types"
)

func init() {
	data.OpSampleMap[types.OperationTypeWitnessCreate] =
		sampleDataWitnessCreateOperation
}

var sampleDataWitnessCreateOperation = `{
  "block_signing_key": "BTS6Ppd2uCY3yrrsiEsSXhCCGiSs8WotZ4ri6qWu4R68hs5sw8vTo",
  "fee": {
    "amount": 1000000000,
    "asset_id": "1.3.0"
  },
  "url": "url-to-proposal",
  "witness_account": "1.2.14002"
}`

//end of file
