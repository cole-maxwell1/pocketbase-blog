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
			"id": "mlem1aae4up4ahh",
			"created": "2024-10-19 15:52:30.772Z",
			"updated": "2024-10-19 15:52:30.772Z",
			"name": "posts_author_vw",
			"type": "view",
			"system": false,
			"schema": [
				{
					"system": false,
					"id": "1d5fyiia",
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
					"id": "xyludfrf",
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
					"id": "wiyayc5l",
					"name": "title",
					"type": "text",
					"required": true,
					"presentable": false,
					"unique": false,
					"options": {
						"min": null,
						"max": 100,
						"pattern": ""
					}
				},
				{
					"system": false,
					"id": "igyjnlf6",
					"name": "content",
					"type": "editor",
					"required": true,
					"presentable": false,
					"unique": false,
					"options": {
						"convertUrls": true
					}
				}
			],
			"indexes": [],
			"listRule": "",
			"viewRule": "",
			"createRule": null,
			"updateRule": null,
			"deleteRule": null,
			"options": {
				"query": "SELECT posts.id,\nusers.` + "`" + `firstName` + "`" + `,\nusers.` + "`" + `lastName` + "`" + `,\nposts.title,\nposts.content,\nposts.created,\nposts.updated\nFROM posts \nINNER JOIN users on posts.` + "`" + `userId` + "`" + `= users.id\n"
			}
		}`

		collection := &models.Collection{}
		if err := json.Unmarshal([]byte(jsonData), &collection); err != nil {
			return err
		}

		return daos.New(db).SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("mlem1aae4up4ahh")
		if err != nil {
			return err
		}

		return dao.DeleteCollection(collection)
	})
}
