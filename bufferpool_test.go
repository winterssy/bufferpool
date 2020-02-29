package bufferpool

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBufferPool(t *testing.T) {
	const dummyData = "dummy data"
	p := New()

	var wg sync.WaitGroup
	for g := 0; g < 10; g++ {
		wg.Add(1)
		go func() {
			for i := 0; i < 100; i++ {
				buf := p.Get()
				assert.Zero(t, buf.Len())
				assert.Equal(t, _size, buf.Cap())

				buf.WriteString(dummyData)
				assert.Equal(t, dummyData, buf.String())

				buf.Free()
			}
			wg.Done()
		}()
	}
	wg.Wait()
}
