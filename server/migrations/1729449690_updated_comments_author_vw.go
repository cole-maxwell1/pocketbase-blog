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

		collection, err := dao.FindCollectionByNameOrId("p4f8mwphwdej2a0")
		if err != nil {
			return err
		}

		options := map[string]any{}
		if err := json.Unmarshal([]byte(`{
			"query": "SELECT\ncomments.id,\ncomments.` + "`" + `postId` + "`" + `,\nusers.` + "`" + `firstName` + "`" + `,\nusers.` + "`" + `lastName` + "`" + `,\ncomments.content,\ncomments.created,\ncomments.updated\nFROM comments\nINNER JOIN users on comments.` + "`" + `userId` + "`" + ` = users.id"
		}`), &options); err != nil {
			return err
		}
		collection.SetOptions(options)

		// remove
		collection.Schema.RemoveField("giumlhyt")

		// remove
		collection.Schema.RemoveField("t4rcbtns")

		// remove
		collection.Schema.RemoveField("08pwwzjp")

		// add
		new_postId := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "5azwiz1h",
			"name": "postId",
			"type": "relation",
			"required": true,
			"presentable": false,
			"unique": false,
			"options": {
				"collectionId": "vsqcudq8utc0c9k",
				"cascadeDelete": true,
				"minSelect": null,
				"maxSelect": 1,
				"displayFields": null
			}
		}`), new_postId); err != nil {
			return err
		}
		collection.Schema.AddField(new_postId)

		// add
		new_firstName := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "xycdan5y",
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
		}`), new_firstName); err != nil {
			return err
		}
		collection.Schema.AddField(new_firstName)

		// add
		new_lastName := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "qbbr1rzk",
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
		}`), new_lastName); err != nil {
			return err
		}
		collection.Schema.AddField(new_lastName)

		// add
		new_content := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "yvvkzqzc",
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
		}`), new_content); err != nil {
			return err
		}
		collection.Schema.AddField(new_content)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("p4f8mwphwdej2a0")
		if err != nil {
			return err
		}

		options := map[string]any{}
		if err := json.Unmarshal([]byte(`{
			"query": "SELECT\ncomments.id,\nusers.` + "`" + `firstName` + "`" + `,\nusers.` + "`" + `lastName` + "`" + `,\ncomments.content,\ncomments.created,\ncomments.updated\nFROM comments\nINNER JOIN users on comments.` + "`" + `userId` + "`" + ` = users.id"
		}`), &options); err != nil {
			return err
		}
		collection.SetOptions(options)

		// add
		del_firstName := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "giumlhyt",
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
		}`), del_firstName); err != nil {
			return err
		}
		collection.Schema.AddField(del_firstName)

		// add
		del_lastName := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "t4rcbtns",
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
		}`), del_lastName); err != nil {
			return err
		}
		collection.Schema.AddField(del_lastName)

		// add
		del_content := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "08pwwzjp",
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
		}`), del_content); err != nil {
			return err
		}
		collection.Schema.AddField(del_content)

		// remove
		collection.Schema.RemoveField("5azwiz1h")

		// remove
		collection.Schema.RemoveField("xycdan5y")

		// remove
		collection.Schema.RemoveField("qbbr1rzk")

		// remove
		collection.Schema.RemoveField("yvvkzqzc")

		return dao.SaveCollection(collection)
	})
}
