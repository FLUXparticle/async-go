package async

type SequenceScope struct {
	results chan any
}

func Sequence(block func(*SequenceScope)) chan any {
	results := make(chan any)
	go func() {
		block(&SequenceScope{results: results})
		close(results)
	}()
	return results
}

func (s *SequenceScope) Yield(value any) {
	s.results <- value
}

func (s *SequenceScope) YieldAll(values []any) {
	for _, value := range values {
		s.results <- value
	}
}
