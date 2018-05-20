package operations

//go:generate ffjson   $GOFILE

import (
	"github.com/denkhaus/bitshares/types"
	"github.com/denkhaus/bitshares/util"
	"github.com/juju/errors"
)

func init() {
	op := &WitnessCreateOperation{}
	types.OperationMap[op.Type()] = op
}

type WitnessCreateOperation struct {
	BlockSigningKey types.PublicKey   `json:"block_signing_key"`
	Fee             types.AssetAmount `json:"fee"`
	URL             string            `json:"url"`
	WitnessAccount  types.GrapheneID  `json:"witness_account"`
}

func (p *WitnessCreateOperation) ApplyFee(fee types.AssetAmount) {
	p.Fee = fee
}

func (p WitnessCreateOperation) Type() types.OperationType {
	return types.OperationTypeWitnessCreate
}

//TODO: verify order
func (p WitnessCreateOperation) Marshal(enc *util.TypeEncoder) error {
	if err := enc.Encode(int8(p.Type())); err != nil {
		return errors.Annotate(err, "encode OperationType")
	}

	if err := enc.Encode(p.Fee); err != nil {
		return errors.Annotate(err, "encode fee")
	}

	if err := enc.Encode(p.WitnessAccount); err != nil {
		return errors.Annotate(err, "encode WitnessAccount")
	}

	if err := enc.Encode(p.URL); err != nil {
		return errors.Annotate(err, "encode URL")
	}

	if err := enc.Encode(p.BlockSigningKey); err != nil {
		return errors.Annotate(err, "encode BlockSigningKey")
	}

	return nil
}

//NewWitnessCreateOperation creates a new WitnessCreateOperation
func NewWitnessCreateOperation() *WitnessCreateOperation {
	tx := WitnessCreateOperation{}
	return &tx
}