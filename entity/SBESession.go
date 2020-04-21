package entity

//go:generate mockgen -destination=../mock/entity/mock_SBESession.go -package=mockentity  sbe/entity SBESession
type SBESession interface {
	Send(msg SBEMessage) error
	Serve() error
	Close() error
	GetSeqNo() uint32
	SetSeqNo(uint32)
	GetUUID() uint64
}
