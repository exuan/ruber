package data

import (
	"context"

	"github.com/exuan/ruber/app/recipe/service/internal/conf"
	"github.com/exuan/ruber/app/recipe/service/internal/data/ent"
	"github.com/exuan/ruber/app/recipe/service/internal/data/ent/migrate"
	"github.com/go-kratos/kratos/v2/log"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/wire"
	_ "github.com/lib/pq"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewRecipeRepo)

// Data .
type Data struct {
	db *ent.Client
}

func NewData(conf *conf.Data, logger log.Logger) (*Data, func(), error) {
	log := log.NewHelper(log.With(logger, "module", "recipe-service/data"))

	client, err := ent.Open(
		conf.Database.Driver,
		conf.Database.Source,
	)
	if err != nil {
		log.Fatalf("failed opening connection to db: %v", err)
	}

	if err := client.Schema.Create(context.Background(), migrate.WithForeignKeys(false)); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	d := &Data{
		db: client,
	}
	return d, func() {
		if err := d.db.Close(); err != nil {
			log.Error(err)
		}
	}, nil
}
