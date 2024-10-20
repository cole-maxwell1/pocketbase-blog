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

		collection, err := dao.FindCollectionByNameOrId("9z1q3a89d8jl95z")
		if err != nil {
			return err
		}

		collection.DeleteRule = types.Pointer("userId = @request.auth.id")

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("9z1q3a89d8jl95z")
		if err != nil {
			return err
		}

		collection.DeleteRule = nil

		return dao.SaveCollection(collection)
	})
}
