package solver

type broadcaster[T any] struct {
	out   chan T
	count int
}

func (b *broadcaster[T]) broadcast(msg T) {
	for i := 0; i < b.count; i++ {
		b.out <- msg
	}
}
