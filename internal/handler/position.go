// Package handler includes grpc handler for working with position service
package handler

import (
	"context"
	"fmt"
	"io"

	"github.com/HekapOo-hub/positionService/internal/model"
	"github.com/HekapOo-hub/positionService/internal/proto/positionpb"
	"github.com/HekapOo-hub/positionService/internal/service"
	log "github.com/sirupsen/logrus"
)

const (
	// PositionPort is a port for position grpc server
	PositionPort = ":50005"
)

// PositionHandler is a handler for grpc server for using position service
type PositionHandler struct {
	positionpb.UnimplementedPositionServiceServer
	positionService *service.PositionService
}

// NewPositionHandler creates new instance of position handler
func NewPositionHandler(ctx context.Context) (*PositionHandler, error) {
	positionService, err := service.NewPositionService(ctx)
	if err != nil {
		log.Warnf("new position handler: %v", err)
		return nil, fmt.Errorf("new position handler: %w", err)
	}

	handler := PositionHandler{positionService: positionService}

	return &handler, nil
}

// Close is used for closing opened position
func (h *PositionHandler) Close(ctx context.Context, id *positionpb.PositionID) (*positionpb.Empty, error) {
	err := h.positionService.Close(ctx, id.Value)
	if err != nil {
		return nil, fmt.Errorf("position handler close: %w", err)
	}
	return &positionpb.Empty{}, nil
}

// Open is used for opening position
func (h *PositionHandler) Open(ctx context.Context, position *positionpb.Position) (*positionpb.Empty, error) {
	// create chan and pass it to channels
	pos := model.Position{ID: position.ID, AccountID: position.AccountID, OrderID: position.OrderID,
		OpenPrice: position.OpenPrice, ClosePrice: position.ClosePrice, TakeProfit: position.TakeProfit, Quantity: position.Quantity,
		StopLoss: position.StopLoss, Symbol: position.Symbol, State: position.State, GuaranteedStopLoss: position.GuaranteedStopLoss, Side: position.Side,
		Leverage: position.Leverage}
	err := h.positionService.Open(ctx, pos)
	if err != nil {
		return nil, fmt.Errorf("position handler open: %w", err)
	}

	return &positionpb.Empty{}, nil
}

// GetOpen returns all open position with specified accountID
func (h *PositionHandler) GetOpen(ctx context.Context, id *positionpb.AccountID) (*positionpb.Positions, error) {
	openPositions := h.positionService.GetOpenByAccountID(id.Value)
	protoOpenPositions := &positionpb.Positions{Value: make([]*positionpb.Position, 0)}
	for i := range openPositions {
		protoPosition := positionpb.Position{ID: openPositions[i].ID, AccountID: openPositions[i].AccountID, OrderID: openPositions[i].OrderID,
			OpenPrice: openPositions[i].OpenPrice, ClosePrice: openPositions[i].ClosePrice, TakeProfit: openPositions[i].TakeProfit,
			StopLoss: openPositions[i].StopLoss, Symbol: openPositions[i].Symbol, State: openPositions[i].State}
		protoOpenPositions.Value = append(protoOpenPositions.Value, &protoPosition)
	}
	return protoOpenPositions, nil
}

// UpdatePrices is called from orderService to update current price for open positions
func (h *PositionHandler) UpdatePrices(stream positionpb.PositionService_UpdatePricesServer) error {
	ctx := stream.Context()
	for {
		select {
		case <-ctx.Done():
			if err := ctx.Err(); err != nil {
				log.Warnf("position handler update prices %v", err)
				return fmt.Errorf("position handler update prices %w", err)
			}
			return nil
		default:
			price, err := stream.Recv()
			if err == io.EOF {
				return nil
			}
			if err != nil {
				log.Printf("receive error in download %v", err)
				continue
			}
			h.positionService.SendPriceForUpdatePositions(model.GeneratedPrice{Symbol: price.Symbol,
				Ask: price.Ask, Bid: price.Bid})
		}
	}
}

// GetProfitLoss send into stream account's profit and loss for every open position
func (h *PositionHandler) GetProfitLoss(accountID *positionpb.AccountID, stream positionpb.PositionService_GetProfitLossServer) error {
	ctx := stream.Context()
	for {
		select {
		case <-ctx.Done():
			if err := ctx.Err(); err != nil {
				log.Warnf("position handler get profit and loss: %v", err)
				return fmt.Errorf("position handler get profit and loss %w", err)
			}
			return nil
		default:
			profitLoss, err := h.positionService.GetProfitLoss(accountID.Value)
			if err != nil {
				return fmt.Errorf("position handler get profit and loss: %w", err)
			}
			for id, value := range profitLoss {
				if streamErr := stream.Send(&positionpb.ProfitLoss{PositionID: id, Value: value}); streamErr != nil {
					log.Warnf("sending map with profit and loss position handler get profit and loss %v", streamErr)
					return fmt.Errorf("sending map with profit and loss position handler get profit and loss %w", streamErr)
				}
			}
		}
	}
}
