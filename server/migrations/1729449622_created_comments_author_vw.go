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
			"id": "p4f8mwphwdej2a0",
			"created": "2024-10-20 18:40:22.155Z",
			"updated": "2024-10-20 18:40:22.155Z",
			"name": "comments_author_vw",
			"type": "view",
			"system": false,
			"schema": [
				{
					"system": false,
					"id": "covhgwol",
					"name": "firstName",
					"type": "text",
					"required": true,
					"presentable": false,
					"unique": false,
					"options": {
						"min": null,
						"max": 200,
						"pattern": ""
					}
				},
				{
					"system": false,
					"id": "rkfoyv4n",
					"name": "lastName",
					"type": "text",
					"required": true,
					"presentable": false,
					"unique": false,
					"options": {
						"min": null,
						"max": 200,
						"pattern": ""
					}
				},
				{
					"system": false,
					"id": "pz2wog3z",
					"name": "content",
					"type": "text",
					"required": true,
					"presentable": false,
					"unique": false,
					"options": {
						"min": null,
						"max": 1000,
						"pattern": ""
					}
				}
			],
			"indexes": [],
			"listRule": null,
			"viewRule": null,
			"createRule": null,
			"updateRule": null,
			"deleteRule": null,
			"options": {
				"query": "SELECT\ncomments.id,\nusers.` + "`" + `firstName` + "`" + `,\nusers.` + "`" + `lastName` + "`" + `,\ncomments.content,\ncomments.created,\ncomments.updated\nFROM comments\nINNER JOIN users on comments.` + "`" + `userId` + "`" + ` = users.id"
			}
		}`

		collection := &models.Collection{}
		if err := json.Unmarshal([]byte(jsonData), &collection); err != nil {
			return err
		}

		return daos.New(db).SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("p4f8mwphwdej2a0")
		if err != nil {
			return err
		}

		return dao.DeleteCollection(collection)
	})
}
