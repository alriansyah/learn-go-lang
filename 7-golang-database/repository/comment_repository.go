package repository

import (
	"context"
	"golang_database/entity"
)

// Nama repository sama dengan nama tabel / entity, yaitu comment
type CommentRepository interface {
	// Jika butuh transaksi, maka implementasikan Tx
	Insert(ctx context.Context, comment entity.Comment) (entity.Comment, error)
	FindById(ctx context.Context, id int32) (entity.Comment, error)
	FindAll(ctx context.Context) ([]entity.Comment, error)
}
