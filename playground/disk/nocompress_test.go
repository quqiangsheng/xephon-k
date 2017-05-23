package disk

import (
	"bytes"
	"io/ioutil"
	"os"
	"testing"

	"github.com/xephonhq/xephon-k/pkg/storage/disk"
	//"encoding/binary"
	"fmt"
	"encoding/binary"
	"io"
	"bufio"
	"github.com/xephonhq/xephon-k/pkg/common"
	"github.com/pkg/errors"
)

const (
	magicnumber uint64 = 0x786570686F6E2D6B
)

// writing series to disk without any compression and then read it out
type fileHeader struct {
	version          uint8
	timeCompression  uint8
	valueCompression uint8
}

// NOTE: must pass a pointer of buffer
func (header *fileHeader) write(buf *bytes.Buffer) {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, magicnumber)
	buf.Write(b)
	buf.WriteByte(header.version)
	buf.WriteByte(header.timeCompression)
	buf.WriteByte(header.valueCompression)
}

func (header *fileHeader) Bytes() []byte {
	var buf bytes.Buffer
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, magicnumber)
	buf.Write(b)
	buf.WriteByte(header.version)
	buf.WriteByte(header.timeCompression)
	buf.WriteByte(header.valueCompression)
	return buf.Bytes()
}

/*

 file

 | magic | version | time compression | value compression | blocks | indexes | footer | magic |

 footer

 | offset of indexes |

 blocks
 | b1 | b2 | b3|

 block
 | t1, t2, ... | v1, v2, .... |

 indexes
 | num indexes | i1 | i2 ... |

 index
 | len | tags | num blocks |b1 offset | b1 size | b1 count | b2 .... |
 */
type blockWriter struct {
	header         fileHeader
	originalWriter io.WriteCloser
	w              *bufio.Writer
	n              int64
}

const intBlock byte = 1
const doubleBlock byte = 2

type indexEntry struct {
	blockType byte
	offset    int64
	size      int64
}

func NewBlockWriter(w io.WriteCloser) *blockWriter {
	return &blockWriter{
		originalWriter: w,
		w:              bufio.NewWriter(w),
		n:              0,
	}
}

func (w *blockWriter) WriteIntSeries(series *common.IntSeries) error {
	n := 0
	// write header if it does not exists
	if w.n == 0 {
		hbits, err := w.w.Write(w.header.Bytes())
		if err != nil {
			return err
		}
		n += hbits
	}
	// write timestamps and values separately
	var tBuf bytes.Buffer
	var vBuf bytes.Buffer
	b := make([]byte, 10)
	for i := 0; i < len(series.Points); i++ {
		written := binary.PutVarint(b, series.Points[i].T)
		tBuf.Write(b[:written])
		written = binary.PutVarint(b, series.Points[i].V)
		vBuf.Write(b[:written])
	}
	tbits, err := w.w.Write(tBuf.Bytes())
	if err != nil {
		return errors.Wrap(err, "fail writing time ")
	}
	n += tbits
	vbits, err := w.w.Write(vBuf.Bytes())
	if err != nil {
		return errors.Wrap(err, "fail writing value")
	}
	n += vbits
	// TODO: add the index

	w.n += int64(n)
	return nil
}

func (w *blockWriter) WriteIndex() error {
	// TODO: implementation
	return nil
}

func (w *blockWriter) Flush() error {
	if err := w.w.Flush(); err != nil {
		return err
	}

	if f, ok := w.originalWriter.(*os.File); ok {
		if err := f.Sync(); err != nil {
			return err
		}
	}
	return nil
}

func (w *blockWriter) Close() error {
	if err := w.Flush(); err != nil {
		return err
	}
	return w.originalWriter.Close()
}

func TestMagicNumber(t *testing.T) {
	var str = "xephon-k"
	t.Log(len([]byte(str))) // 8 byte, uint64
	t.Log([]byte(str))
	// [120 101 112 104 111 110 45 107]
	// 78 65 70 68 6F 6E 2D 6B
	t.Logf("% X", []byte(str))
	t.Logf("%X", []byte(str))
	t.Log([]byte(str)[0])

	// convert the magic number into binary
	t.Log(magicnumber)
	b := make([]byte, 10)
	// FIXME: it takes 9 byte instead of 8 byte to write a uint64 http://stackoverflow.com/questions/17289898/does-a-uint64-take-8-bytes-storage
	t.Log(binary.PutUvarint(b, magicnumber)) // 9 instead of 8
	t.Log(b)
	v, err := binary.ReadUvarint(bytes.NewReader(b))
	if err != nil {
		t.Fatal(err)
	}
	t.Log(v)
	// this use 8 byte
	binary.BigEndian.PutUint64(b, magicnumber)
	t.Log(b)
	t.Log(binary.BigEndian.Uint64(b))

	// Uvarint would use less than 8 byte for small value
	t.Log(binary.PutUvarint(b, 1))   // 1
	t.Log(binary.PutUvarint(b, 256)) // 2
}

func TestNoCompress_Header(t *testing.T) {
	header := fileHeader{version: 1, timeCompression: disk.CompressionNone, valueCompression: disk.CompressionNone}
	//header := fileHeader{version: 1, timeCompression: disk.CompressionGzip, valueCompression: disk.CompressionZlib}
	tmpfile, err := ioutil.TempFile("", "xephon-no-compress")
	if err != nil {
		t.Fatal(err)
	}

	defer os.Remove(tmpfile.Name())

	var buf bytes.Buffer
	// TODO: Endianness problem https://github.com/xephonhq/xephon-k/issues/34
	// but it seems for single uint8, this is not a problem
	//binary.Write(&buf, binary.LittleEndian, header.version)
	//binary.Write(&buf, binary.LittleEndian, header.timeCompression)
	//binary.Write(&buf, binary.LittleEndian, header.valueCompression)

	header.write(&buf)

	n, err := tmpfile.Write(buf.Bytes())
	t.Logf("written %d bytes\n", n)
	if err != nil {
		t.Fatal(err)
	}
	tmpfile.Close()

	// read stuff back
	f, err := os.Open(tmpfile.Name())
	readBuf := make([]byte, 11)
	n, err = f.Read(readBuf)
	t.Logf("read %d bytes\n", n)
	if err != nil {
		t.Fatal(err)
	}
	f.Close()
	// convert to header
	newHeader := fileHeader{}
	if binary.BigEndian.Uint64(readBuf[:8]) != magicnumber {
		t.Fatal("magic number does not match!")
	} else {
		t.Log("magic number match")
	}
	newHeader.version = uint8(readBuf[8])
	newHeader.timeCompression = uint8(readBuf[9])
	newHeader.valueCompression = uint8(readBuf[10])
	fmt.Printf("version %d, time compression %d, value compression %d\n",
		newHeader.version, newHeader.timeCompression, newHeader.valueCompression)
}

func TestNoCompress_Block(t *testing.T) {
	tmpfile, err := ioutil.TempFile("", "xephon-no-compress")
	if err != nil {
		t.Fatal(err)
	}

	defer os.Remove(tmpfile.Name())

	w := NewBlockWriter(tmpfile)
	s := common.NewIntSeries("s")
	s.Points = []common.IntPoint{{T: 1359788400000, V: 1}, {T: 1359788500000, V: 2}}
	w.WriteIntSeries(s)
	t.Log(w.n)
	w.Close()
}