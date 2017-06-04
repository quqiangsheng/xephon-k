package encoding

import (
	"github.com/pkg/errors"
	"fmt"

	"github.com/xephonhq/xephon-k/pkg/util"
)

var log = util.Logger.NewEntryWithPkg("k.encoding")

const (
	_                    byte = iota
	CodecRawBigEndian
	CodecRawLittleEndian
	CodecVarInt
)

var (
	ErrTooSmall      = errors.New("data for decoding is too small")
	ErrCodecMismatch = errors.New("decoder got data encoded using other codec")
	// TODO: rename to codec not registered? we have some codec that does not support certain type, maybe
	// ErrValueTypeNotSupported, or make it an struct
	ErrCodecNotSupported = errors.New("codec is not supported")
)

var (
	registeredCodec        []byte
	registeredValueEncoder []ValueEncoder
	registeredValueDecoder []ValueDecoder
	registeredFactory      map[byte]CodecFactory = make(map[byte]CodecFactory, 4)
)

type CodecFactory interface {
	NewTimeEncoder() (TimeEncoder, error)
	NewTimeDecoder() (TimeDecoder, error)
	NewIntValueEncoder() (ValueEncoder, error)
	NewIntValueDecoder() (ValueDecoder, error)
	NewDoubleValueEncoder() (ValueEncoder, error)
	NewDoubleValueDecoder() (ValueDecoder, error)
}

type TimeEncoder interface {
	Codec() byte
	Bytes() ([]byte, error)
	Reset()
	WriteTime(t int64)
	// TODO: Once or Batch?
	//WriteTimeOnce(t []int64)
}

type TimeDecoder interface {
	Init([]byte) error
	Next() bool
	ReadTime() int64
	// TODO: ReadAll
	// TODO: Reset
}

// TODO: add SupportInt/Double/String, so we can skip some of them in test, and check if it match series type when we enforce encoding in config
type ValueEncoder interface {
	Codec() byte
	Bytes() ([]byte, error)
	Reset()
	WriteInt(v int64)
	WriteDouble(v float64)
}

type ValueDecoder interface {
	Init([]byte) error
	Next() bool
	ReadInt() int64
	ReadDouble() float64
}

func IsRegisteredCodec(codec byte) bool {
	for _, c := range registeredCodec {
		if c == codec {
			return true
		}
	}
	return false
}

func GetFactory(codec byte) (CodecFactory, error) {
	factory, ok := registeredFactory[codec]
	if !ok {
		return nil, errors.Errorf("codec %s does not have registered factory", CodecString(codec))
	}
	return factory, nil
}

func CodecString(codec byte) string {
	switch codec {
	case CodecRawBigEndian:
		return "codec: raw bigendian"
	case CodecRawLittleEndian:
		return "codec: raw littleendian"
	case CodecVarInt:
		return "codec: variable length integer"
	default:
		return fmt.Sprintf("codec: unknown %d", codec)
	}
}
