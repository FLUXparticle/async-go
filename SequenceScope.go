package async

type SequenceScope[T any] struct {
	results chan T
}

func Sequence[T any](block func(*SequenceScope[T])) chan T {
	results := make(chan T)
	go func() {
		block(&SequenceScope[T]{results: results})
		close(results)
	}()
	return results
}

func (s *SequenceScope[T]) Yield(value T) {
	s.results <- value
}

func (s *SequenceScope[T]) YieldAll(values []T) {
	for _, value := range values {
		s.results <- value
	}
}
