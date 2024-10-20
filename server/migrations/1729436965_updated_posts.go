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

		collection, err := dao.FindCollectionByNameOrId("vsqcudq8utc0c9k")
		if err != nil {
			return err
		}

		// add
		new_tags := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "qh5ys0vf",
			"name": "tags",
			"type": "relation",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"collectionId": "33q7jo0z4z6q8dd",
				"cascadeDelete": false,
				"minSelect": null,
				"maxSelect": null,
				"displayFields": null
			}
		}`), new_tags); err != nil {
			return err
		}
		collection.Schema.AddField(new_tags)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("vsqcudq8utc0c9k")
		if err != nil {
			return err
		}

		// remove
		collection.Schema.RemoveField("qh5ys0vf")

		return dao.SaveCollection(collection)
	})
}
