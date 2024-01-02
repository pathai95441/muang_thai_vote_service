package db_transaction

import (
	"context"
	"database/sql"
	"reflect"

	"github.com/hashicorp/go-multierror"
)

//go:generate mockgen -source=./db_transaction.go -destination=./mock/mock_db_transaction.go -package=mock_db_transaction
type IWithTxn interface {
	WithTransaction(ctx context.Context, statements []Statement, opts ...TransactionOpt) (result error)
	WithTransactionWithFunc(ctx context.Context, fn TxFn, opts ...TransactionOpt) (result error)
}

type WithTxn struct {
	db *sql.DB
}

func NewWithTxn(db *sql.DB) *WithTxn {
	return &WithTxn{db: db}
}

type TransactionOpt func(txOpts *sql.TxOptions)

func WithIsolationOpt(isolationLevel sql.IsolationLevel) TransactionOpt {
	return func(txOpts *sql.TxOptions) {
		if txOpts == nil {
			txOpts = &sql.TxOptions{}
		}

		txOpts.Isolation = isolationLevel
	}
}

func WithReadOnlyOpt() TransactionOpt {
	return func(txOpts *sql.TxOptions) {
		if txOpts == nil {
			txOpts = &sql.TxOptions{}
		}

		txOpts.ReadOnly = true
	}
}

type Statement struct {
	fn   interface{}
	args []interface{}
	// withTx is position of tx that being used as argument of fn. If not used, should set to -1
	withTx int
}

func NewStatement(fn interface{}, withTx int, args ...interface{}) Statement {
	return Statement{
		fn:     fn,
		withTx: withTx,
		args:   args,
	}
}

// Inspired by https://pseudomuto.com/2018/01/clean-sql-transactions-in-golang/
func (t *WithTxn) WithTransaction(ctx context.Context, statements []Statement, opts ...TransactionOpt) (result error) {
	var txOpts *sql.TxOptions
	for _, opt := range opts {
		opt(txOpts)
	}

	tx, err := t.db.BeginTx(ctx, txOpts)
	if err != nil {
		return err
	}

	for _, stmt := range statements {
		if err = t.exec(tx, stmt.fn, stmt.withTx, stmt.args...); err != nil {
			break
		}
	}

	if err != nil {
		result = multierror.Append(result, err)
		err = tx.Rollback()
		result = multierror.Append(result, err)
		return
	}

	if err := tx.Commit(); err != nil {
		result = multierror.Append(result, err)
		return
	}

	return
}

type TxFn func(*sql.Tx) error

func (t *WithTxn) WithTransactionWithFunc(ctx context.Context, fn TxFn, opts ...TransactionOpt) (result error) {
	var txOpts *sql.TxOptions
	for _, opt := range opts {
		opt(txOpts)
	}

	tx, err := t.db.BeginTx(ctx, txOpts)
	if err != nil {
		return err
	}

	err = fn(tx)
	if err != nil {
		result = multierror.Append(result, err)
		err = tx.Rollback()
		result = multierror.Append(result, err)
		return
	}

	if err := tx.Commit(); err != nil {
		result = multierror.Append(result, err)
		return
	}

	return
}

func (t *WithTxn) exec(tx interface{}, fn interface{}, withTx int, args ...interface{}) error {
	v := []reflect.Value{}
	for i, arg := range args {
		if i == withTx {
			v = append(v, reflect.ValueOf(tx))
		}

		v = append(v, reflect.ValueOf(arg))
	}
	if withTx == len(args) {
		v = append(v, reflect.ValueOf(tx))
	}

	result := reflect.ValueOf(fn).Call(v)
	for _, c := range result {
		if err, ok := c.Interface().(error); ok {
			return err
		}
	}

	return nil
}
