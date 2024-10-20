package migrations

import (
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/tools/types"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("vsqcudq8utc0c9k")
		if err != nil {
			return err
		}

		collection.ViewRule = types.Pointer("userId = @request.auth.id")

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("vsqcudq8utc0c9k")
		if err != nil {
			return err
		}

		collection.ViewRule = nil

		return dao.SaveCollection(collection)
	})
}
