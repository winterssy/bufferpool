package bufferpool

import (
	"bytes"
	"sync"
)

const _size = 1024 // by default, create 1 KiB buffers

var (
	_pool = New()
	// Get retrieves a buffer from the pool, creating one if necessary.
	Get = _pool.Get
)

// Pool is a type-safe wrapper around a sync.Pool.
type Pool struct {
	p *sync.Pool
}

// Buffer is a wrapper around a bytes.Buffer. It's intended to be pooled, so
// the only way to construct one is via a Pool.
type Buffer struct {
	*bytes.Buffer
	pool Pool
}

// New constructs a new Pool.
func New() Pool {
	return Pool{p: &sync.Pool{
		New: func() interface{} {
			return &Buffer{Buffer: bytes.NewBuffer(make([]byte, 0, _size))}
		}},
	}
}

// Get retrieves a Buffer from the pool, creating one if necessary.
func (p Pool) Get() *Buffer {
	buf := p.p.Get().(*Buffer)
	buf.pool = p
	return buf
}

func (p Pool) put(buf *Buffer) {
	p.p.Put(buf)
}

// Free returns the Buffer to its Pool.
//
// Callers must not retain references to the Buffer after calling Free.
func (b *Buffer) Free() {
	b.Reset()
	b.pool.put(b)
}
