// Package service includes implementation of position service behavior
package service

import (
	"context"
	"fmt"
	"sync"

	"github.com/HekapOo-hub/positionService/internal/model"
	"github.com/HekapOo-hub/positionService/internal/repository"
	log "github.com/sirupsen/logrus"
)

type PositionService struct {
	positionRepository repository.PositionRepository
	cashRepository     repository.CashRepository
	channels           map[string][]chan model.GeneratedPrice // symbol - chan model.GeneratedPrice for accountID
	channelsMutex      sync.RWMutex
	ctx                context.Context
	opened             map[string]chan model.Position // accountID- openedPosition
	closed             map[string]chan string         // accountID- closedPositionID
}

func NewPositionService(ctx context.Context) (*PositionService, error) {
	// need a goroutine which will follow for new accounts created on another instances and call checkAccountPositions with this new account
	opened := make(map[string]chan model.Position)
	closed := make(map[string]chan string)
	openedWithNewAcc := make(chan string)
	positionRepository, err := repository.NewPostgresPositionRepository(ctx, opened, openedWithNewAcc, closed)
	if err != nil {
		return nil, fmt.Errorf("new position service: %w", err)
	}
	cashRepository, err := repository.NewRedisCashRepository(ctx)
	if err != nil {
		return nil, fmt.Errorf("new position service: %w", err)
	}
	service := PositionService{positionRepository: positionRepository,
		cashRepository: cashRepository, channels: make(map[string][]chan model.GeneratedPrice), channelsMutex: sync.RWMutex{}, ctx: ctx,
		opened: opened, closed: closed}
	service.channels["silver"] = make([]chan model.GeneratedPrice, 0)
	service.channels["gold"] = make([]chan model.GeneratedPrice, 0)
	service.channels["bitcoin"] = make([]chan model.GeneratedPrice, 0)
	service.channels["eBay"] = make([]chan model.GeneratedPrice, 0)
	service.channels["crude oil"] = make([]chan model.GeneratedPrice, 0)

	// for each accountID need to call checkAccountPositions
	accountsID := service.positionRepository.GetAllAccountsID()
	for id := range accountsID {
		service.opened[id] = make(chan model.Position)
		service.closed[id] = make(chan string)
		go service.checkAccountPositions(ctx, id)
	}
	go func() {
		for {
			select {
			case accountID := <-openedWithNewAcc:
				go service.checkAccountPositions(ctx, accountID)
			case <-ctx.Done():
				return
			}
		}
	}()
	positions := service.positionRepository.GetAllOpen()
	for i := range positions {
		service.opened[positions[i].AccountID] <- positions[i]
	}
	return &service, nil
}

func (s *PositionService) GetProfitLoss(accountID string) (map[string]float64, error) {
	positions := s.positionRepository.GetUpdatedByAccountID(accountID)
	profitLoss := make(map[string]float64)
	for i := range positions {
		profitLoss[positions[i].ID] = (positions[i].ClosePrice - positions[i].OpenPrice) * positions[i].Quantity
	}
	return profitLoss, nil
}

func (s *PositionService) Close(ctx context.Context, positionID string) error {
	accID, err := s.close(ctx, positionID)
	if err != nil {
		return err
	}
	s.closed[accID] <- positionID
	return nil
}
func (s *PositionService) close(ctx context.Context, positionID string) (string, error) {
	accID, price, err := s.positionRepository.Close(ctx, positionID)
	if err != nil {
		return "", fmt.Errorf("position service close: %w", err)
	}
	err = s.cashRepository.Update(accID, price)
	if err != nil {
		return "", fmt.Errorf("position service: %v", err)
	}
	return accID, nil
}
func (s *PositionService) GuaranteedStopLossClose(ctx context.Context, positionID string) error {
	accID, price, err := s.positionRepository.GuaranteedStopLossClose(ctx, positionID)
	if err != nil {
		return fmt.Errorf("position service guaranteed stop loss close: %w", err)
	}
	err = s.cashRepository.Update(accID, price)
	if err != nil {
		return fmt.Errorf("position service guaranteed stop loss close: %w", err)
	}
	return nil
}

func (s *PositionService) Open(ctx context.Context, position model.Position) error {
	positions := s.GetOpenByAccountID(position.AccountID)
	balance := s.cashRepository.Get(position.AccountID) - position.OpenPrice*position.Quantity
	positionsToClose := make([]model.Position, 0)
	for i := range positions {
		if balance < 0 {
			balance += positions[i].ClosePrice * positions[i].Quantity
			positionsToClose = append(positionsToClose, positions[i])
		} else {
			break
		}
	}
	if balance < 0 {
		return fmt.Errorf("not enough money to open this position")
	}
	for i := range positionsToClose {
		s.closed[positionsToClose[i].AccountID] <- positionsToClose[i].ID
		err := s.Close(ctx, positionsToClose[i].ID)
		if err != nil {
			return fmt.Errorf("position service open: %v", err)
		}
	}

	err := s.positionRepository.Create(ctx, position)
	if err != nil {
		return fmt.Errorf("position service open: %w", err)
	}
	err = s.cashRepository.Update(position.AccountID, -position.OpenPrice*position.Quantity)
	if err != nil {
		return fmt.Errorf("position service open: %w", err)
	}
	if _, ok := s.opened[position.AccountID]; !ok {
		s.opened[position.AccountID] = make(chan model.Position)
		s.closed[position.AccountID] = make(chan string)
		go s.checkAccountPositions(s.ctx, position.AccountID)
	}
	s.opened[position.AccountID] <- position
	return nil
}

func (s *PositionService) GetOpenByAccountID(accountID string) []model.Position {
	openPositions := s.positionRepository.GetOpenByAccountID(accountID)
	return openPositions
}

func (s *PositionService) UpdatePriceByPositionID(positionID string, price model.GeneratedPrice) (float64, error) {
	return s.positionRepository.UpdatePrice(positionID, price)
}

func (s *PositionService) SendPriceForUpdatePositions(price model.GeneratedPrice) {
	s.channelsMutex.Lock()
	for _, channel := range s.channels[price.Symbol] {
		channel <- price
	}
	s.channelsMutex.Unlock()
}

func (s *PositionService) checkAccountPositions(ctx context.Context, accountID string) {
	s.channelsMutex.Lock()
	channel := make(chan model.GeneratedPrice)
	for symbol := range s.channels {
		s.channels[symbol] = append(s.channels[symbol], channel)
	}
	s.channelsMutex.Unlock()
	opened := make([]model.Position, 0)
	for {
		select {
		case openedPosition := <-s.opened[accountID]:
			opened = append(opened, openedPosition)
		case price := <-channel:
			for i := range opened {
				if opened[i].Symbol != price.Symbol {
					continue
				}
				closePrice, err := s.UpdatePriceByPositionID(opened[i].ID, price)
				if err != nil {
					log.Warnf("check account position: %v", err)
					continue
				}
				if opened[i].Leverage {
					// check take profit and stop loss
					if opened[i].GuaranteedStopLoss && closePrice <= opened[i].StopLoss {
						err = s.GuaranteedStopLossClose(ctx, opened[i].ID)
						if err != nil {
							log.Warnf("check account positions: %v", err)
						}
						for j := range opened {
							if opened[j] == opened[i] {
								opened = append(opened[:j], opened[j+1:]...)
								break
							}
						}
					} else if closePrice <= opened[i].StopLoss || closePrice >= opened[i].TakeProfit {
						_, err = s.close(ctx, opened[i].ID)
						if err != nil {
							log.Warnf("check account positions: %v", err)
						}
						for j := range opened {
							if opened[j] == opened[i] {
								opened = append(opened[:j], opened[j+1:]...)
								break
							}
						}
					}
				}
			}
		case closedPositionID := <-s.closed[accountID]:
			for i := range opened {
				if opened[i].ID == closedPositionID {
					opened = append(opened[:i], opened[i+1:]...)
					break
				}
			}

		case <-ctx.Done():
			log.Info("ctx done..")
			return
		}
	}
}
