// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0
// source: clear_invoice_table.sql

package database

import (
	"context"
)

const clearInvoiceData = `-- name: ClearInvoiceData :exec

DELETE FROM invoice
`

func (q *Queries) ClearInvoiceData(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, clearInvoiceData)
	return err
}
