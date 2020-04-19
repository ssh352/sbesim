package service

import (
	"errors"
	"sbe/entity"
	fix "sbe/sbe/iLinkBinary"
)

func createMsg(msgID int) (entity.SBEMessage, error) {
	switch msgID {
	case 500:
		return &fix.Negotiate500{}, nil
	case 501:
		return &fix.NegotiationResponse501{}, nil
	case 502:
		return &fix.NegotiationReject502{}, nil
	case 503:
		return &fix.Establish503{}, nil
	case 504:
		return &fix.EstablishmentAck504{}, nil
	case 505:
		return &fix.EstablishmentReject505{}, nil
	case 506:
		return &fix.Sequence506{}, nil
	case 507:
		return &fix.Terminate507{}, nil
	case 508:
		return &fix.RetransmitRequest508{}, nil
	case 509:
		return &fix.Retransmission509{}, nil
	case 510:
		return &fix.RetransmitReject510{}, nil
	case 513:
		return &fix.NotApplied513{}, nil
	case 514:
		return &fix.NewOrderSingle514{}, nil
	case 515:
		return &fix.OrderCancelReplaceRequest515{}, nil
	case 516:
		return &fix.OrderCancelRequest516{}, nil
	case 517:
		return &fix.MassQuote517{}, nil
	case 518:
		return &fix.PartyDetailsDefinitionRequest518{}, nil
	case 519:
		return &fix.PartyDetailsDefinitionRequestAck519{}, nil
	case 521:
		return &fix.BusinessReject521{}, nil
	case 522:
		return &fix.ExecutionReportNew522{}, nil
	case 523:
		return &fix.ExecutionReportReject523{}, nil
	case 524:
		return &fix.ExecutionReportElimination524{}, nil
	case 525:
		return &fix.ExecutionReportTradeOutright525{}, nil
	case 526:
		return &fix.ExecutionReportTradeSpread526{}, nil
	case 527:
		return &fix.ExecutionReportTradeSpreadLeg527{}, nil
	case 528:
		return &fix.QuoteCancel528{}, nil
	case 529:
		return &fix.OrderMassActionRequest529{}, nil
	case 530:
		return &fix.OrderMassStatusRequest530{}, nil
	case 531:
		return &fix.ExecutionReportModify531{}, nil
	case 532:
		return &fix.ExecutionReportStatus532{}, nil
	case 533:
		return &fix.OrderStatusRequest533{}, nil
	case 534:
		return &fix.ExecutionReportCancel534{}, nil
	case 535:
		return &fix.OrderCancelReject535{}, nil
	case 536:
		return &fix.OrderCancelReplaceReject536{}, nil
	case 537:
		return &fix.PartyDetailsListRequest537{}, nil
	case 538:
		return &fix.PartyDetailsListReport538{}, nil
	case 539:
		return &fix.ExecutionAck539{}, nil
	case 543:
		return &fix.RequestForQuote543{}, nil
	case 544:
		return &fix.NewOrderCross544{}, nil
	case 545:
		return &fix.MassQuoteAck545{}, nil
	case 546:
		return &fix.RequestForQuoteAck546{}, nil
	case 548:
		return &fix.ExecutionReportTradeAddendumOutright548{}, nil
	case 549:
		return &fix.ExecutionReportTradeAddendumSpread549{}, nil
	case 550:
		return &fix.ExecutionReportTradeAddendumSpreadLeg550{}, nil
	case 560:
		return &fix.SecurityDefinitionRequest560{}, nil
	case 561:
		return &fix.SecurityDefinitionResponse561{}, nil
	case 562:
		return &fix.OrderMassActionReport562{}, nil
	case 563:
		return &fix.QuoteCancelAck563{}, nil
	}
	return nil, errors.New("unknown msg id")
}
