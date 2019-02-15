package lib

import (
	"errors"
	"fmt"
)

type GoTickets interface {
	Take()
	Return()
	Active() bool
	Total() uint32
	Remainder() uint32
}

type SampleTickets struct {
	total uint32
	ticketCh chan struct{}
	active bool
}

func NewSampleTickets(total uint32) (GoTickets, error) {
    ret := SampleTickets{}
    if !ret.init(total) {
    	errMsg := fmt.Sprintf("SampleTickets init failed, total=%d\n", total)
    	return nil, errors.New(errMsg)
	}
	return &ret, nil
}

func (s *SampleTickets)init(total uint32) bool {
	if s.active {
		return false
	}
	if total == 0 {
		return false
	}
	ch := make(chan struct{}, total)
	n := int(total)
	for i:=0; i<n; i++ {
		ch <- struct{}{}
	}
	s.ticketCh = ch
	s.total = total
	s.active = true
	return true
}

func (s *SampleTickets) Take() {
	<- s.ticketCh
}

func (s *SampleTickets) Return() {
	s.ticketCh <- struct{}{}
}

func (s *SampleTickets) Active() bool {
	return s.active
}

func (s *SampleTickets) Total() uint32 {
	return s.total
}

func (s *SampleTickets) Remainder() uint32 {
	return uint32(len(s.ticketCh))
}



