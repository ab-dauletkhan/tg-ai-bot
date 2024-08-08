// package balance

// import (
// 	"context"
// 	"time"

// 	pgx "github.com/jackc/pgx/v5"
// 	"github.com/jackc/pgx/v5/pgxpool"
// )

// type Balance struct {
// 	ID        int
// 	UserID    int64
// 	Amount    float64
// 	Status    string
// 	CreatedAt time.Time
// }

// func InsertBalance(pool *pgxpool.Pool, balance *Balance) error {
// 	query := `
//     INSERT INTO balance (user_id, amount, status)
//     VALUES ($1, $2, $3)
//     RETURNING id, created_at;
//     `
// 	row := pool.QueryRow(context.Background(), query, balance.UserID, balance.Amount, balance.Status)
// 	return row.Scan(&balance.ID, &balance.CreatedAt)
// }

// func GetBalanceByUserID(pool *pgxpool.Pool, userID int64) (*Balance, error) {
// 	query := `
//     SELECT id, user_id, amount, status, created_at
//     FROM balance
//     WHERE user_id = $1::bigint;
//     `
// 	row := pool.QueryRow(context.Background(), query, userID)

// 	var balance Balance
// 	err := row.Scan(&balance.ID, &balance.UserID, &balance.Amount, &balance.Status, &balance.CreatedAt)
// 	if err != nil {
// 		if err == pgx.ErrNoRows {
// 			return nil, nil // No balance found for the user
// 		}
// 		return nil, err
// 	}

// 	return &balance, nil
// }

package balance

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Balance struct {
	UserID int64
	Amount float64
	Status string
}

func GetBalanceByUserID(dbpool *pgxpool.Pool, userID int64) (*Balance, error) {
	var b Balance
	err := dbpool.QueryRow(context.Background(), "SELECT user_id, amount, status FROM balance WHERE user_id=$1::bigint", userID).Scan(&b.UserID, &b.Amount, &b.Status)
	if err != nil {
		return nil, err
	}
	return &b, nil
}

func InsertBalance(dbpool *pgxpool.Pool, balance *Balance) error {
	_, err := dbpool.Exec(context.Background(), "INSERT INTO balance (user_id, amount, status) VALUES ($1, $2, $3)", balance.UserID, balance.Amount, balance.Status)
	return err
}
