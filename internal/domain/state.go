package domain
// Domain / State — navigation logic (pure mathematics).

type State struct {
	Current int
	Total   int
}

func NewState(total int) (State) {
	return State{Current: 1, Total: total}
}

func (s *State) Next() {
	if s.Current < s.Total {
		s.Current++
	} else {
		s.Current = 1
	}
}

func (s *State) Prev() {
	if s.Current > 1 {
		s.Current--
	} else {
		s.Current = s.Total
	}
}

func (s *State) Open(k int) bool {
	if k < 1 || k > s.Total {
		return false
	}
	s.Current = k
	return true
}
