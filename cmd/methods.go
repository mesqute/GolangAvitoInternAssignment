package main

type UpdateCashRequest struct {
	UserId uint
	Value  float32
}

type Giver struct {
	UserId uint
}

type Receiver struct {
	UserId uint
}

type TransferCashRequest struct {
	Giver
	Receiver
	Value float32
}

type Validate interface {
	Validate()
}
