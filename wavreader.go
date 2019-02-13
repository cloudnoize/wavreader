package wavreader

import (
	"errors"
	"io"
	"os"

	"github.com/cloudnoize/conv"
)

func handleErr(err error) bool {
	if err != nil {
		println(err.Error())
		return true
	}
	return false
}

type WavHHeader struct {
	hdr []byte
}

func (this *WavHHeader) chunkId() string {
	return string(this.hdr[:4])
}

func (this *WavHHeader) chunkSize() uint32 {
	return conv.BytesToUint32(this.hdr, 4)
}

func (this *WavHHeader) format() string {
	return string(this.hdr[8:12])
}

func (this *WavHHeader) subchunkID() string {
	return string(this.hdr[12:16])
}

func (this *WavHHeader) Subchunk1Size() uint32 {
	return conv.BytesToUint32(this.hdr, 16)
}
func (this *WavHHeader) AudioFormat() uint16 {
	return conv.BytesToUint16(this.hdr, 20)
}

func (this *WavHHeader) NumChannels() uint16 {
	return conv.BytesToUint16(this.hdr, 22)
}
func (this *WavHHeader) SampleRate() uint32 {
	return conv.BytesToUint32(this.hdr, 24)
}

func (this *WavHHeader) ByteRate() uint32 {
	return conv.BytesToUint32(this.hdr, 28)
}
func (this *WavHHeader) BlockAlign() uint16 {
	return conv.BytesToUint16(this.hdr, 32)
}
func (this *WavHHeader) BitsPerSample() uint16 {
	return conv.BytesToUint16(this.hdr, 34)
}

func (this *WavHHeader) Subchunk2ID() string {
	return string(this.hdr[36:40])
}

func (this *WavHHeader) Subchunk2Size() uint32 {
	return conv.BytesToUint32(this.hdr, 40)
}

func (this *WavHHeader) String() {
	println("ChunkId ", this.chunkId())
	println("chunkSize ", this.chunkSize())
	println("format ", this.format())
	println("subchunk ID ", this.subchunkID())
	println("Subchunk1Size ", this.Subchunk1Size())
	println("AudioFormat ", this.AudioFormat())
	println("NumChannels ", this.NumChannels())
	println("SampleRate ", this.SampleRate())
	println("ByteRate ", this.ByteRate())
	println("BlockAlign ", this.BlockAlign())
	println("BitsPerSample ", this.BitsPerSample())
	println("Subchunk2ID ", this.Subchunk2ID())
	println("Subchunk2Size ", this.Subchunk2Size())
}

type Wav struct {
	WavHHeader
	io.Reader
	io.ReaderAt
	io.Closer
}

func New(path string) (*Wav, error) {
	fd, err := os.Open(path)

	if handleErr(err) {
		return nil, err
	}

	b := make([]byte, 44)

	n, err := fd.Read(b)

	if handleErr(err) {
		fd.Close()
		return nil, err
	}

	if n < 44 {
		return nil, errors.New("Failed to read 44 bytes(wav header size)")
	}

	w := &Wav{WavHHeader: WavHHeader{hdr: b}, Reader: fd, ReaderAt: fd, Closer: fd}

	return w, nil

}

func (this *Wav) PrintHeader() {
	this.String()
}
