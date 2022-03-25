package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/HekapOo-hub/positionService/internal/config"
	"github.com/HekapOo-hub/positionService/internal/model"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

// PositionRepository describes work with position bd
type PositionRepository interface {
	Create(ctx context.Context, position model.Position) error
	Close(ctx context.Context, positionID string) (string, float64, error)
	UpdatePrice(positionID string, price model.GeneratedPrice) (float64, error)
	GetOpenByAccountID(accountID string) []model.Position
	GetAllOpen() []model.Position
	GuaranteedStopLossClose(ctx context.Context, positionID string) (string, float64, error)
	GetUpdatedByAccountID(accountID string) []model.Position
	GetAllAccountsID() map[string]bool
	Get(id string) (model.Position, error)
}

// PostgresPositionRepository implements PositionRepository interface
type PostgresPositionRepository struct {
	pool          *pgxpool.Pool
	positionCache map[string]model.Position
	updated       map[string]bool
	cacheMutex    sync.RWMutex
	updatedMutex  sync.RWMutex
}

// NewPostgresPositionRepository creates new instance of PostgresPositionRepository
func NewPostgresPositionRepository(ctx context.Context, opened map[string]chan model.Position, openedWithNewAcc chan string,
	closed map[string]chan string) (*PostgresPositionRepository, error) {
	cfg, err := config.GetPostgresConfig()
	if err != nil {
		log.Warnf("postgres config error: %v", err)
		return nil, fmt.Errorf("new postgres repository func %v", err)
	}

	pool, err := pgxpool.Connect(ctx, cfg.GetURL())
	if err != nil {
		log.Warnf("creating postgres connection pool error %v", err)
		return nil, fmt.Errorf("creating postgres connection pool error %w", err)
	}
	repo := &PostgresPositionRepository{pool: pool, positionCache: make(map[string]model.Position), cacheMutex: sync.RWMutex{},
		updated: make(map[string]bool), updatedMutex: sync.RWMutex{}}
	err = repo.reloadCache(ctx)
	if err != nil {
		return nil, fmt.Errorf("new postgres repository: %w", err)
	}
	go repo.waitForNotifications(ctx, cfg.GetURL(), opened, openedWithNewAcc, closed)
	return repo, nil
}

func (repo *PostgresPositionRepository) Create(ctx context.Context, position model.Position) error {
	query := "insert into positions (id,account_id,order_id,open_price,close_price,take_profit,stop_loss,symbol," +
		"guaranteed_stop_loss,state,quantity, leverage,side) values ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13)"
	_, err := repo.pool.Exec(ctx, query, position.ID, position.AccountID, position.OrderID, position.OpenPrice, position.ClosePrice,
		position.TakeProfit, position.StopLoss, position.Symbol, position.GuaranteedStopLoss, "OPEN",
		position.Quantity, position.Leverage, position.Side)
	if err != nil {
		return fmt.Errorf("postgres create position: %w", err)
	}

	repo.positionCache[position.ID] = position
	return nil
}

func (repo *PostgresPositionRepository) Close(ctx context.Context, positionID string) (accID string, outcome float64, err error) {
	query := "update positions set state= $1 where id= $2 and state= $3"

	tag, err := repo.pool.Exec(ctx, query, "CLOSE", positionID, "OPEN")
	if err != nil {
		return "", 0, fmt.Errorf("postgres close position: %w", err)
	}
	if tag.RowsAffected() == 0 {
		return "", 0, nil
	}

	accID, outcome = repo.positionCache[positionID].AccountID,
		repo.positionCache[positionID].ClosePrice*repo.positionCache[positionID].Quantity

	repo.cacheMutex.Lock()
	delete(repo.positionCache, positionID)
	repo.cacheMutex.Unlock()
	repo.updatedMutex.Lock()
	delete(repo.updated, positionID)
	repo.updatedMutex.Unlock()
	return accID, outcome, nil
}

func (repo *PostgresPositionRepository) GuaranteedStopLossClose(ctx context.Context, positionID string) (accID string, outcome float64, err error) {
	query := "update positions set state=$1 where id=$2 and state=$3"
	closeStr := "STOP LOSS CLOSE"
	openStr := "OPEN"
	tag, err := repo.pool.Exec(ctx, query, closeStr, positionID, openStr)
	if err != nil {
		return "", 0, fmt.Errorf("postgres close position: %w", err)
	}
	if tag.RowsAffected() == 0 {
		return "", 0, nil
	}

	commissionForGuaranteedStopLoss := 0.95

	accID, outcome = repo.positionCache[positionID].AccountID,
		repo.positionCache[positionID].StopLoss*repo.positionCache[positionID].Quantity*commissionForGuaranteedStopLoss

	repo.cacheMutex.Lock()
	delete(repo.positionCache, positionID)
	repo.cacheMutex.Unlock()
	repo.updatedMutex.Lock()
	delete(repo.updated, positionID)
	repo.updatedMutex.Unlock()
	return accID, outcome, nil
}

func (repo *PostgresPositionRepository) UpdatePrice(positionID string, price model.GeneratedPrice) (float64, error) {
	repo.cacheMutex.RLock()
	pos, ok := repo.positionCache[positionID]
	if !ok {
		repo.cacheMutex.RUnlock()
		return 0, fmt.Errorf("postgres position repository update price: no such position in cache")
	}
	repo.cacheMutex.RUnlock()
	if pos.Side == "BUY" {
		pos.ClosePrice = price.Bid
	} else if pos.Side == "SELL" {
		pos.ClosePrice = price.Ask
	}
	repo.cacheMutex.Lock()
	repo.positionCache[positionID] = pos
	repo.cacheMutex.Unlock()
	log.Infof("postgres update price: id: %v      price: %v", pos.ID, pos.ClosePrice)
	repo.updatedMutex.Lock()
	repo.updated[positionID] = true
	repo.updatedMutex.Unlock()
	return pos.ClosePrice, nil
}

func (repo *PostgresPositionRepository) GetOpenByAccountID(accountID string) []model.Position {
	openPositions := make([]model.Position, 0)
	repo.cacheMutex.RLock()
	for id := range repo.positionCache {
		if repo.positionCache[id].AccountID == accountID {
			openPositions = append(openPositions, repo.positionCache[id])
		}
	}
	repo.cacheMutex.RUnlock()
	return openPositions
}

func (repo *PostgresPositionRepository) getAllOpen(ctx context.Context) ([]model.Position, error) {
	query := "select * from positions where state=$1"
	rows, err := repo.pool.Query(ctx, query, "OPEN")
	if err != nil {
		return nil, fmt.Errorf("position repository get all open: %w", err)
	}
	defer rows.Close()
	openPositions := make([]model.Position, 0)
	for rows.Next() {
		var p model.Position
		err = rows.Scan(&p.ID, &p.AccountID, &p.OrderID, &p.OpenPrice, &p.ClosePrice,
			&p.TakeProfit, &p.StopLoss, &p.Symbol, &p.GuaranteedStopLoss, &p.State, &p.Quantity, &p.Leverage, &p.Side)
		if err != nil {
			return nil, fmt.Errorf("position repository get all open: %w", err)
		}
		openPositions = append(openPositions, p)
	}
	return openPositions, nil
}

func (repo *PostgresPositionRepository) Get(id string) (model.Position, error) {
	repo.cacheMutex.RLock()
	defer repo.cacheMutex.RUnlock()
	position, ok := repo.positionCache[id]
	if !ok {
		return model.Position{}, fmt.Errorf("no such position in cache")
	}
	return position, nil
}

func (repo *PostgresPositionRepository) GetUpdatedByAccountID(accountID string) []model.Position {
	updatedPositions := make([]model.Position, 0)
	repo.cacheMutex.RLock()
	for id := range repo.positionCache {
		repo.updatedMutex.RLock()
		updated, ok := repo.updated[id]
		if !ok {
			repo.updatedMutex.RUnlock()
			continue
		}
		pos := repo.positionCache[id]
		repo.updatedMutex.RUnlock()
		if pos.AccountID == accountID && updated {
			updatedPositions = append(updatedPositions, pos)
			repo.updatedMutex.Lock()
			repo.updated[id] = false
			repo.updatedMutex.Unlock()
		}
	}
	repo.cacheMutex.RUnlock()
	return updatedPositions
}

func (repo *PostgresPositionRepository) GetAllAccountsID() map[string]bool {
	accountsID := make(map[string]bool)
	repo.cacheMutex.RLock()
	for id := range repo.positionCache {
		accountsID[repo.positionCache[id].AccountID] = true
	}
	repo.cacheMutex.RUnlock()
	return accountsID
}

func (repo *PostgresPositionRepository) GetAllOpen() []model.Position {
	openPositions := make([]model.Position, 0)
	repo.cacheMutex.RLock()
	for id := range repo.positionCache {
		openPositions = append(openPositions, repo.positionCache[id])
	}
	repo.cacheMutex.RUnlock()
	return openPositions
}
func (repo *PostgresPositionRepository) reloadCache(ctx context.Context) error {
	openPositions, err := repo.getAllOpen(ctx)
	if err != nil {
		return fmt.Errorf("postgres position repository reload cache %w", err)
	}
	repo.cacheMutex.Lock()
	for i := range openPositions {
		repo.positionCache[openPositions[i].ID] = openPositions[i]
	}
	repo.cacheMutex.Unlock()
	return nil
}
func (repo *PostgresPositionRepository) waitForNotifications(ctx context.Context, url string, opened map[string]chan model.Position,
	openedWithNewAcc chan string, closed map[string]chan string) {
	reportErr := func(ev pq.ListenerEventType, err error) {
		if err != nil {
			log.Warnf("Failed to start listener in wait for notification: %v", err)
			return
		}
	}

	listener := pq.NewListener(url, 50*time.Second, time.Minute, reportErr)
	err := listener.Listen(config.PostgresChannel)
	if err != nil {
		log.Warnf("wait for notifications: %v", err)
		return
	}
	pidRow := repo.pool.QueryRow(ctx, "select pg_backend_pid();")
	var pid int
	err = pidRow.Scan(&pid)
	if err != nil {
		log.Warnf("wait for notification scan pid row: %v", err)
		return
	}
	for {
		select {
		case n := <-listener.Notify:
			if pid == n.BePid {
				break
			}
			var msg model.PostgresChannelMessage
			err := json.Unmarshal([]byte(n.Extra), &msg)
			if err != nil {
				log.Warnf("wait for notofication listen position data: %v", err)
				return
			}
			switch msg.Action {
			case "INSERT":
				repo.cacheMutex.Lock()
				log.Warnf("position added: %v", msg.Position)
				repo.positionCache[msg.Position.ID] = msg.Position
				repo.cacheMutex.Unlock()
				if _, ok := opened[msg.Position.AccountID]; ok {
					opened[msg.Position.AccountID] <- msg.Position
				} else {
					openedWithNewAcc <- msg.Position.AccountID
					opened[msg.Position.AccountID] = make(chan model.Position)
					closed[msg.Position.AccountID] = make(chan string)
					opened[msg.Position.AccountID] <- msg.Position
				}
			case "UPDATE":
				if msg.Position.State == "CLOSE" {
					repo.cacheMutex.Lock()
					delete(repo.positionCache, msg.Position.ID)
					repo.cacheMutex.Unlock()
					closed[msg.Position.AccountID] <- msg.Position.ID
				}
				/*case "DELETE":
				_, err = repo.pool.Exec(ctx, "delete from positions where id=$1", msg.Position.ID)
				if err != nil {
					log.Warnf("wait for notification: %v", err)
					break
				}*/
			}
		case <-ctx.Done():
			log.Warn("notify ctx done...")
			return
		}
	}
}
