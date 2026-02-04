package common

type Stack struct {
	list []rune
}

func (s *Stack) Push(item rune) {
	s.list = append(s.list, item)
}

func (s *Stack) Pop() rune {
	if len(s.list) == 0 {
		return 0
	}
	item := s.list[len(s.list)-1]
	s.list = s.list[:len(s.list)-1]
	return item
}

func (s *Stack) Count() int {
	return len(s.list)
}
