// Package repository consist of repositories for position service
package repository

import (
	"context"
	"fmt"
	"strconv"
	"sync"

	"github.com/HekapOo-hub/positionService/internal/config"
	"github.com/go-redis/redis"
	log "github.com/sirupsen/logrus"
)

const (
	startingBalance = 200
)

type CashRepository interface {
	Update(accountID string, price float64) error
	Get(accountID string) float64
}
type RedisCashRepository struct {
	client *redis.Client
	cache  map[string]float64
	mu     sync.RWMutex
}

func NewRedisCashRepository(ctx context.Context) (*RedisCashRepository, error) {
	redisCfg, err := config.NewRedisConfig()
	if err != nil {
		return nil, fmt.Errorf("new redis cash repository: %v", err)
	}

	redisClient := redis.NewClient(&redis.Options{
		Addr:     redisCfg.Addr,
		Password: redisCfg.Password,
		DB:       redisCfg.DB,
	})

	repo := &RedisCashRepository{client: redisClient, cache: make(map[string]float64), mu: sync.RWMutex{}}
	go repo.listen(ctx)
	return repo, nil
}

func (repo *RedisCashRepository) Update(accountID string, price float64) error {
	err := repo.client.XAdd(&redis.XAddArgs{
		Stream:       config.RedisCashStream,
		MaxLenApprox: 0,
		MaxLen:       0,
		ID:           "",
		Values: map[string]interface{}{
			accountID: price,
		},
	}).Err()
	if err != nil {
		return fmt.Errorf("redis cash repository open: %w", err)
	}
	return nil
}

func (repo *RedisCashRepository) Get(accountID string) float64 {
	repo.mu.RLock()
	defer repo.mu.RUnlock()
	if balance, ok := repo.cache[accountID]; ok {
		return balance
	}
	return startingBalance
}
func (repo *RedisCashRepository) listen(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			res, err := repo.client.XRead(&redis.XReadArgs{
				Block:   0,
				Count:   1,
				Streams: []string{config.RedisCashStream, "$"},
			}).Result()
			if err != nil {
				continue
			}
			if res[0].Messages == nil {
				log.Warn("message is empty")
				continue
			}
			cashMap := res[0].Messages[0].Values
			for accID, priceStr := range cashMap {
				price, err := strconv.ParseFloat(priceStr.(string), 32)
				if err != nil {
					log.Warnf("redis cash repository listen: %v", err)
					break
				}
				_, ok := repo.cache[accID]
				repo.mu.Lock()
				if !ok {
					repo.cache[accID] = startingBalance
				}
				repo.cache[accID] += price
				repo.mu.Unlock()
			}
		}
	}
}
