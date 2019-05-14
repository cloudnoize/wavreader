package main

import (
	"os"

	"github.com/cloudnoize/wavreader"
)

// type streamImp struct {
// 	q *locklessq.Q
// }

// //mono float 32
// func (s *streamImp) Cbb(inputBuffer, outputBuffer unsafe.Pointer, frames uint64) {
// 	ob := (*[512]float32)(outputBuffer)
// 	for i := uint64(0); i < frames; i++ {
// 		val, ok := s.q.Pop()
// 		if ok {
// 			(*ob)[i] = val

// 		}
// 	}
// }

// func (this *streamImp) Write(b []byte) (n int, err error) {
// 	for i := 0; i < len(b)/4; i++ {
// 		f := conv.BytesToFloat32(b, i*4)
// 		this.q.Insert(f)
// 	}
// 	return len(b), nil
// }

func main() {
	// sigs := make(chan os.Signal, 1)
	// done := make(chan bool, 1)
	// signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	// go func() {
	// 	sig := <-sigs
	// 	fmt.Println()
	// 	fmt.Println(sig)
	// 	done <- true
	// }()

	f := os.Getenv("FILE")
	w, err := wavreader.New(f)
	if err != nil {
		println(err.Error())
		return
	}

	defer w.Close()

	w.String()

	// bytescount := w.DataBytesCount()
	// println("data size in bytes - ", bytescount)
	// si := &streamImp{q: locklessq.New(int32(bytescount))}

	// pa.Cba[0] = si

	// pa.Initialize()
	// s, _ := pa.OpenDefaultStream(0, 1, pa.Float32, 44100, 512, nil)
	// s.Start()
	// defer func() {
	// 	s.Stop()
	// 	s.Close()
	// }()

	// io.Copy(si, w)
	// fmt.Println("awaiting signal")
	// <-done
	// fmt.Println("exiting")

}
