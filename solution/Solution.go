package solution

type Solution struct{
	Stack []int
	innerTransaction *Solution
}

func (s *Solution) Push(value int) {
	if s.innerTransaction == nil {
		s.Stack = append([]int{value}, s.Stack...)
	} else {
		for current := s.innerTransaction ; current != nil ; {
			current.Push(value)
			return
		}
	}
}

func (s *Solution) Top() int {
	if len(s.Stack) == 0 {
		return 0
	}
	if s.innerTransaction == nil {
		return s.Stack[0]
	}
	var rv int
	for current := s.innerTransaction; current != nil; current = current.innerTransaction {
		rv = current.Top()
	}
	return rv
}

func (s *Solution) Pop() {
	if len(s.Stack) == 0 {
		return
	}
	if s.innerTransaction == nil {
		s.Stack = s.Stack[1:]
		return
	}
	for current := s.innerTransaction; current != nil ; {
		current.Pop()
		return
	}
}

func (s *Solution) Begin() {
	if s.innerTransaction == nil {
		s.innerTransaction = &Solution{Stack: make([]int, 0)}
	} else {
		for current := s.innerTransaction; current != nil; current = current.innerTransaction {
			if current.innerTransaction == nil {
				current.Begin()
				return
			}
		}
	}
}

func (s *Solution) Commit() bool {
	for current := s; current.innerTransaction != nil; current = current.innerTransaction {
		if current.innerTransaction.innerTransaction == nil {
			current.Stack = append(current.innerTransaction.Stack, current.Stack...)
			current.innerTransaction = nil
			return true
		}
	}
	return false
}

func (s *Solution) Rollback() bool {
	for current := s; current.innerTransaction != nil; current = current.innerTransaction {
		if current.innerTransaction.innerTransaction == nil {
			current.innerTransaction = nil
			return true
		}
	}
	return false
}
