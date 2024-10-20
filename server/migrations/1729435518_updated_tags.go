package migrations

import (
	"encoding/json"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/models/schema"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("33q7jo0z4z6q8dd")
		if err != nil {
			return err
		}

		// add
		new_postId := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "tjv6lano",
			"name": "postId",
			"type": "relation",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"collectionId": "vsqcudq8utc0c9k",
				"cascadeDelete": false,
				"minSelect": null,
				"maxSelect": null,
				"displayFields": null
			}
		}`), new_postId); err != nil {
			return err
		}
		collection.Schema.AddField(new_postId)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("33q7jo0z4z6q8dd")
		if err != nil {
			return err
		}

		// remove
		collection.Schema.RemoveField("tjv6lano")

		return dao.SaveCollection(collection)
	})
}
