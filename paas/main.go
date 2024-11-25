package paasio

import (
	"io"
	"sync"
)

// Define readCounter and writeCounter types here.
type readCounter struct {
	cntOps   int
	cntBytes int64
	mutex    sync.Mutex
	reader   io.Reader
}
type writeCounter struct {
	cntOps   int
	cntBytes int64
	mutex    sync.Mutex
	writer   io.Writer
}

// For the return of the function NewReadWriteCounter, you must also define a type that satisfies the ReadWriteCounter interface.

type readWriteCounter struct {
	readBytes int64
	readOps   int
	readMutex sync.Mutex

	writeBytes int64
	writeOps   int
	writeMutex sync.Mutex

	rw io.ReadWriter
}

func (r *readWriteCounter) Read(p []byte) (n int, err error) {
	r.readMutex.Lock()
	defer r.readMutex.Unlock()
	n, err = r.rw.Read(p) // Выполнить чтение
	if n > 0 {            // Учитывать только успешные операции
		r.readOps++
		r.readBytes += int64(n)
	}
	return n, err
}

func (r *readWriteCounter) Write(p []byte) (n int, err error) {
	r.writeMutex.Lock()
	defer r.writeMutex.Unlock()
	r.writeOps++
	n, err = r.rw.Write(p)
	r.writeBytes += int64(n)
	return n, err
}

func (r *readWriteCounter) ReadCount() (n int64, nops int) {
	r.readMutex.Lock()
	defer r.readMutex.Unlock()
	return r.readBytes, r.readOps
}

func (r *readWriteCounter) WriteCount() (n int64, nops int) {
	r.writeMutex.Lock()
	defer r.writeMutex.Unlock()
	return r.writeBytes, r.writeOps
}

func NewWriteCounter(writer io.Writer) WriteCounter {
	mtx := sync.Mutex{}
	return &writeCounter{
		cntOps:   0,
		cntBytes: 0,
		mutex:    mtx,
		writer:   writer,
	}
}

func NewReadCounter(reader io.Reader) ReadCounter {
	mtx := sync.Mutex{}
	return &readCounter{
		cntOps:   0,
		cntBytes: 0,
		mutex:    mtx,
		reader:   reader,
	}
}

func NewReadWriteCounter(readwriter io.ReadWriter) ReadWriteCounter {
	wmtx := sync.Mutex{}
	rmtx := sync.Mutex{}
	return &readWriteCounter{
		readBytes:  0,
		readOps:    0,
		readMutex:  rmtx,
		writeBytes: 0,
		writeOps:   0,
		writeMutex: wmtx,

		rw: readwriter,
	}
}

func (rc *readCounter) Read(p []byte) (int, error) {
	rc.mutex.Lock()
	defer rc.mutex.Unlock()
	n, err := rc.reader.Read(p)
	if n > 0 {
		rc.cntOps++
		rc.cntBytes += int64(n)
	}
	return n, err
}

func (rc *readCounter) ReadCount() (int64, int) {
	return rc.cntBytes, rc.cntOps
}

func (wc *writeCounter) Write(p []byte) (int, error) {
	wc.mutex.Lock()
	defer wc.mutex.Unlock()
	wc.cntOps++
	n, err := wc.writer.Write(p)
	wc.cntBytes += int64(n)
	return n, err
}

func (wc *writeCounter) WriteCount() (int64, int) {
	return wc.cntBytes, wc.cntOps
}
