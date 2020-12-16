package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTicketList_ErrorRate(t *testing.T) {
	tl := NewTicketList(`class: 1-3 or 5-7
row: 6-11 or 33-44
seat: 13-40 or 45-50

your ticket:
7,1,14

nearby tickets:
7,3,47
40,4,50
55,2,20
38,6,12`)

	require.Equal(t, 71, tl.ErrorRate())
}

func TestTicketList_MyTicket(t *testing.T) {
	tl := NewTicketList(`class: 0-1 or 4-19
row: 0-5 or 8-19
seat: 0-13 or 16-19

your ticket:
11,12,13

nearby tickets:
3,9,18
15,1,5
5,14,9`)

	mt := tl.MyTicket()
	require.Equal(t, 12, mt["class"])
	require.Equal(t, 11, mt["row"])
	require.Equal(t, 13, mt["seat"])
}
