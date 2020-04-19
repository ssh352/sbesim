package entity

type SBESession interface {
	Send(msg SBEMessage) error
	Serve() error
	Close() error
}
