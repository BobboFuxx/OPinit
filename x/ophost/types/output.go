package types

import (
	bytes "bytes"
	"encoding/binary"

	"golang.org/x/crypto/sha3"
)

func (output Output) Validate() error {
	if len(output.OutputRoot) != 32 {
		return ErrInvalidHashLength.Wrap("output_root")
	}

	return nil
}

func (output Output) IsEmpty() bool {
	return len(output.OutputRoot) == 0 && output.L1BlockTime.IsZero() && output.L2BlockNumber == 0
}

func GenerateOutputRoot(version []byte, stateRoot []byte, storageRoot []byte, latestBlockHash []byte) [32]byte {
	seed := make([]byte, 32*4)
	copy(seed, version)
	copy(seed[32:], stateRoot)
	copy(seed[64:], storageRoot)
	copy(seed[96:], latestBlockHash)
	return sha3.Sum256(seed)
}

func GenerateWithdrawalHash(bridgeId uint64, l2Sequence uint64, sender string, receiver string, denom string, amount uint64) [32]byte {
	var withdrawalHash [32]byte
	seed := []byte{}
	seed = binary.BigEndian.AppendUint64(seed, bridgeId)
	seed = binary.BigEndian.AppendUint64(seed, l2Sequence)
	// variable length
	seed = append(seed, sender...) // put utf8 encoded address
	seed = append(seed, Splitter)
	// variable length
	seed = append(seed, receiver...) // put utf8 encoded address
	seed = append(seed, Splitter)
	// variable length
	seed = append(seed, denom...)
	seed = append(seed, Splitter)
	seed = binary.BigEndian.AppendUint64(seed, amount)

	// double hash the leaf node
	withdrawalHash = sha3.Sum256(seed)
	withdrawalHash = sha3.Sum256(withdrawalHash[:])

	return withdrawalHash
}

func GenerateNodeHash(a, b []byte) [32]byte {
	var data [32]byte
	switch bytes.Compare(a, b) {
	case 0, 1: // equal or greater
		data = sha3.Sum256(append(b, a...))
	case -1: // less
		data = sha3.Sum256(append(a, b...))
	}
	return data
}

func GenerateRootHashFromProofs(data [32]byte, proofs [][]byte) [32]byte {
	for _, proof := range proofs {
		data = GenerateNodeHash(data[:], proof)
	}
	return data
}
