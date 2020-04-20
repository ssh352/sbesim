package entity

type SBESession interface {
	Send(msg SBEMessage) error
	Serve() error
	Close() error
	GetSeqNo() uint32
	SetSeqNo(uint32)
	GetUUID() uint64
}
