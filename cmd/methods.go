package main

type updateCashRequest struct {
	UserId uint
	Value  int
}

type giver struct {
	UserId uint
}

type receiver struct {
	UserId uint
}

type transferCashRequest struct {
	giver
	receiver
	Value int
}
