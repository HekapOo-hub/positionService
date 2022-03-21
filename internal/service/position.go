package service

import (
	"context"
	"fmt"
	"github.com/HekapOo-hub/positionService/internal/model"
	"github.com/HekapOo-hub/positionService/internal/repository"
	log "github.com/sirupsen/logrus"
	"sync"
)

type PositionService struct {
	positionRepository repository.PositionRepository
	cashRepository     repository.CashRepository
	channels           map[string][]chan model.GeneratedPrice // symbol- price from grpc stream
	mu                 sync.RWMutex
	ctx                context.Context
}

func NewPositionService(ctx context.Context) (*PositionService, error) {
	positionRepository, err := repository.NewPostgresPositionRepository(ctx)
	if err != nil {
		return nil, fmt.Errorf("new position service: %w", err)
	}
	cashRepository, err := repository.NewRedisCashRepository(ctx)
	if err != nil {
		return nil, fmt.Errorf("new position service: %w", err)
	}
	service := PositionService{positionRepository: positionRepository,
		cashRepository: cashRepository, channels: make(map[string][]chan model.GeneratedPrice), mu: sync.RWMutex{}, ctx: ctx}
	service.channels["silver"] = make([]chan model.GeneratedPrice, 0)
	service.channels["gold"] = make([]chan model.GeneratedPrice, 0)
	service.channels["bitcoin"] = make([]chan model.GeneratedPrice, 0)
	service.channels["eBay"] = make([]chan model.GeneratedPrice, 0)
	service.channels["crude oil"] = make([]chan model.GeneratedPrice, 0)
	positions := service.positionRepository.GetAllOpen()
	for _, position := range positions {
		go service.checkOpenPosition(ctx, position)
	}
	return &service, nil
}

func (s *PositionService) GetProfitLoss(accountID string) (map[string]float64, error) {
	positions := s.positionRepository.GetUpdatedByAccountID(accountID)
	profitLoss := make(map[string]float64)
	for _, position := range positions {
		profitLoss[position.ID] = (position.ClosePrice - position.OpenPrice) * position.Quantity

	}
	return profitLoss, nil
}

func (s *PositionService) Close(ctx context.Context, positionID string) error {
	accID, price, err := s.positionRepository.Close(ctx, positionID)
	if err != nil {
		return fmt.Errorf("position service close: %w", err)
	}
	err = s.cashRepository.Update(accID, price)
	if err != nil {
		return fmt.Errorf("position service: %v", err)
	}
	return nil
}
func (s *PositionService) GuaranteedStopLossClose(ctx context.Context, positionID string) error {
	accId, price, err := s.positionRepository.GuaranteedStopLossClose(ctx, positionID)
	if err != nil {
		return fmt.Errorf("position service guaranteed stop loss close: %w", err)
	}
	err = s.cashRepository.Update(accId, price)
	if err != nil {
		return fmt.Errorf("position service guaranteed stop loss close: %w", err)
	}
	return nil
}

func (s *PositionService) Open(ctx context.Context, position model.Position) error {
	err := s.positionRepository.Create(ctx, position)
	if err != nil {
		return fmt.Errorf("position service open: %w", err)
	}
	err = s.cashRepository.Update(position.AccountID, -position.OpenPrice*position.Quantity)
	if err != nil {
		return fmt.Errorf("position service open: %w", err)
	}
	go s.checkOpenPosition(s.ctx, position)
	return nil
}

func (s *PositionService) GetOpenByAccountID(ctx context.Context, accountID string) ([]model.Position, error) {
	openPositions, err := s.positionRepository.GetOpenByAccountID(ctx, accountID)
	if err != nil {
		return nil, fmt.Errorf("get open position service: %w", err)
	}
	return openPositions, nil
}

func (s *PositionService) UpdatePriceByPositionID(positionID string, price model.GeneratedPrice) (float64, error) {
	closePrice, err := s.positionRepository.UpdatePrice(positionID, price)
	if err != nil {
		return 0, fmt.Errorf("position service update price by id: %w", err)
	}
	return closePrice, nil
}

func (s *PositionService) SendPriceForUpdatePositions(price model.GeneratedPrice) {
	s.mu.Lock()
	for _, channel := range s.channels[price.Symbol] {
		channel <- price
	}
	s.mu.Unlock()
}

func (s *PositionService) checkOpenPosition(ctx context.Context, position model.Position) {
	channel := make(chan model.GeneratedPrice, 1)
	s.mu.Lock()
	s.channels[position.Symbol] = append(s.channels[position.Symbol], channel)
	s.mu.Unlock()
	for {
		//updated field:
		select {
		case price := <-channel:
			log.Infof("check open positions: %v", price)
			// need to set position.Updated=true
			closePrice, err := s.UpdatePriceByPositionID(position.ID, price)
			if err != nil {
				log.Warnf("position handler open: %v", err)
				continue
			}
			if position.Leverage {
				// check take profit and stop loss
				if position.GuaranteedStopLoss && closePrice <= position.StopLoss {
					err = s.GuaranteedStopLossClose(ctx, position.ID)
					if err != nil {
						log.Warnf("position handler open: %v", err)
					}
					// delete channel from s.channels[position.Symbol]
					close(channel)
					channels := make([]chan model.GeneratedPrice, 0)
					s.mu.Lock()
					for _, ch := range s.channels[price.Symbol] {
						if ch == channel {
							continue
						}
						channels = append(channels, ch)
					}
					s.channels[price.Symbol] = channels
					s.mu.Unlock()
					return
				} else if closePrice <= position.StopLoss || closePrice >= position.TakeProfit {
					err = s.Close(ctx, position.ID)
					if err != nil {
						log.Warnf("position handler open: %v", err)
					}
					close(channel)
					channels := make([]chan model.GeneratedPrice, 0)
					s.mu.Lock()
					for _, ch := range s.channels[price.Symbol] {
						if ch == channel {
							continue
						}
						channels = append(channels, ch)
					}
					s.channels[price.Symbol] = channels
					s.mu.Unlock()
					return
				}
			}
		case <-ctx.Done():
			log.Info("ctx done..")
			return
		}
	}
}
