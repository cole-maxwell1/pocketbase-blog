package migrations

import (
	"encoding/json"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/models"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		jsonData := `{
			"id": "33q7jo0z4z6q8dd",
			"created": "2024-10-20 14:43:01.038Z",
			"updated": "2024-10-20 14:43:01.038Z",
			"name": "tags",
			"type": "base",
			"system": false,
			"schema": [
				{
					"system": false,
					"id": "tdprgdoo",
					"name": "name",
					"type": "text",
					"required": true,
					"presentable": false,
					"unique": false,
					"options": {
						"min": null,
						"max": 50,
						"pattern": ""
					}
				}
			],
			"indexes": [],
			"listRule": "",
			"viewRule": "",
			"createRule": "",
			"updateRule": null,
			"deleteRule": null,
			"options": {}
		}`

		collection := &models.Collection{}
		if err := json.Unmarshal([]byte(jsonData), &collection); err != nil {
			return err
		}

		return daos.New(db).SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("33q7jo0z4z6q8dd")
		if err != nil {
			return err
		}

		return dao.DeleteCollection(collection)
	})
}
