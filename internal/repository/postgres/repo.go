package postgres

import postgresdb "github.com/sazonovItas/mocosso/gen/go/db/postgres"

type postgresRepository struct {
	*postgresdb.Queries
}

func NewPostgresRepository(queries *postgresdb.Queries) *postgresRepository {
	return &postgresRepository{
		Queries: queries,
	}
}
