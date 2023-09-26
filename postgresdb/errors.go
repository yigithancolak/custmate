package postgresdb

import "errors"

var (
	ErrBeginTransaction    = errors.New("failed to begin database transaction")
	ErrRollbackTransaction = errors.New("failed to roll back database transaction")
	ErrInsertScanGroup     = errors.New("failed to insert and scan group")
	ErrCreateTime          = errors.New("failed to insert and scan time")
)
