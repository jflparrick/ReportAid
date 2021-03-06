package utils

import (

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// GetString - turn parseBytes32String into string
func GetString ( val [32]byte ) (string) {

    slice := val[:]
    trimmed := common.TrimRightZeroes(slice)
    encodeVal := hexutil.Encode(trimmed)
    decodeVal, _ := hexutil.Decode(encodeVal)

    return string(decodeVal)

}
