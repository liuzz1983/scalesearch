package codec

import (
	"github.com/liuzz1983/ScaleSearch/core/filedb"
	"github.com/liuzz1983/ScaleSearch/core/codec"
)

type W3Codec struct {
	blockLimit  int32
	compression bool
	inlineLimit int32
}

func (codec *W3Codec) TermsReader(storage filedb.Storage, segment string) (codec.TermsReader, error) {
	return nil, nil
}

type W3TermsReader struct {
}
