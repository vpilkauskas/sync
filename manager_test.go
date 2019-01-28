package sync

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"time"
)

func TestManager_Set(t *testing.T) {
	m := New()
	m.Set(1000)

	assert.Equal(t, 1000, m.m.workAmount)
}

func BenchmarkManager_Set(b *testing.B) {
	m := New()

	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		m.Set(1)
	}

	assert.Equal(b, 1, m.m.workAmount)
}

func TestManager_Add(t *testing.T) {
	m := New()
	m.Add(10)
	m.Add(10)

	assert.Equal(t, 20, m.m.workAmount)
}

func BenchmarkManager_Add(b *testing.B) {
	m := New()

	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		m.Add(1)
	}

	assert.Equal(b, b.N, m.m.workAmount)
}

func TestManager_Done(t *testing.T) {
	m := New()
	m.Add(1)

	timer := time.NewTimer(time.Millisecond * 300)
	defer timer.Stop()

	m.Done()

	select {
	case <-timer.C:
		assert.Fail(t, "timeout")
	case <-m.D:
	}
}

func BenchmarkManager_WorkFlow(b *testing.B) {
	b.ReportAllocs()

	m := New()
	m.Add(b.N)

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		go m.Done()
	}

	<-m.D
}
