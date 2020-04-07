/**
 * Author:  hushengbo
 * Email:   hushengbo@qutoutiao.net
 * Date:    2020-04-07
 */

package stack

import "container/list"

type Stack struct {
	list *list.List
}

func New() *Stack {
	return &Stack{list.New()}
}

func (s *Stack)Push(v interface{}) {
	s.list.PushBack(v)
}

func (s *Stack)Pop() interface{} {
	e := s.list.Back()
	if e != nil {
		s.list.Remove(e)
		return e.Value
	}
	return nil
}

func (s *Stack)Peak() interface{} {
	e := s.list.Back()
	if e != nil {
		return e.Value
	}
	return nil
}

func (s *Stack)Len() int {
	return s.list.Len()
}

func (s *Stack)Empty() bool {
	return s.list.Len() == 0
}
