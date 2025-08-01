package accessiblerole

import (
	"context"
	// "database/sql"
	"errors"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4"
	// "github.com/UraharaKiska/go-auth/internal/client/db"
	"github.com/UraharaKiska/go-auth/internal/repository"
	modelRepo "github.com/UraharaKiska/go-auth/internal/repository/accessibleRole/model"
	"github.com/UraharaKiska/platform-common/pkg/db"
)

const (
	tableName             = "accessible_role"
	idColumn              = "id"
	endpointColumn        = "endpoint"
	roleColumn			 = "role"
)

type repo struct {
	db db.Client
}

func NewRepository(db db.Client) repository.AccessibleRepository {
	return &repo{db: db}
}

func (r *repo) GetEndpointRole(ctx context.Context, endpoint string) (*modelRepo.EndpointRole, error) {
	builderInsert :=  sq.Select(endpointColumn, roleColumn).
		From(tableName).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{endpointColumn: endpoint})

	query, args, err := builderInsert.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "accessible_role_repository.GetEndpointRole",
		QueryRaw: query,
	}

	var endpointRole modelRepo.EndpointRole
	err = r.db.DB().ScanOneContext(ctx, &endpointRole, q, args...)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	

	return &endpointRole, nil

}