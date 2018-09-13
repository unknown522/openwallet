package btcLikeTxDriver

import "errors"

type TxIn struct {
	TxID                  []byte
	Vout                  []byte
	ScriptPubkeySignature []byte
	Sequence              []byte
}

func newTxInForEmptyTrans(vin []Vin) ([]TxIn, error) {
	var ret []TxIn

	for _, v := range vin {
		txid, err := reverseHexToBytes(v.TxID)
		if err != nil || len(txid) != 32 {
			return nil, errors.New("Invalid previous txid!")
		}
		vout := uint32ToLittleEndianBytes(v.Vout)

		ret = append(ret, TxIn{txid[:], vout, nil, nil})
	}
	return ret, nil
}

func (vin *TxIn) setSequence(lockTime uint32, replaceable bool) {
	if replaceable {
		vin.Sequence = uint32ToLittleEndianBytes(SequenceMaxBip125RBF)
	} else if lockTime != 0 {
		vin.Sequence = uint32ToLittleEndianBytes(SequenceFinal - 1)
	} else {
		vin.Sequence = uint32ToLittleEndianBytes(SequenceFinal)
	}
}
