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

		collection, err := dao.FindCollectionByNameOrId("mlem1aae4up4ahh")
		if err != nil {
			return err
		}

		options := map[string]any{}
		if err := json.Unmarshal([]byte(`{
			"query": "SELECT posts.id,\nposts.` + "`" + `userId` + "`" + `,\nusers.` + "`" + `firstName` + "`" + `,\nusers.` + "`" + `lastName` + "`" + `,\nposts.title,\nposts.content,\nposts.created,\nposts.updated\nFROM posts \nINNER JOIN users on posts.` + "`" + `userId` + "`" + `= users.id\n\n"
		}`), &options); err != nil {
			return err
		}
		collection.SetOptions(options)

		// remove
		collection.Schema.RemoveField("yodr10xd")

		// remove
		collection.Schema.RemoveField("zxovtje0")

		// remove
		collection.Schema.RemoveField("1ot7ngjk")

		// remove
		collection.Schema.RemoveField("ulqequh5")

		// remove
		collection.Schema.RemoveField("uhjzphmz")

		// remove
		collection.Schema.RemoveField("tnxx2rll")

		// add
		new_userId := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "lmatkkhv",
			"name": "userId",
			"type": "relation",
			"required": true,
			"presentable": false,
			"unique": false,
			"options": {
				"collectionId": "_pb_users_auth_",
				"cascadeDelete": true,
				"minSelect": null,
				"maxSelect": 1,
				"displayFields": null
			}
		}`), new_userId); err != nil {
			return err
		}
		collection.Schema.AddField(new_userId)

		// add
		new_firstName := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "lzsccauv",
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
			"id": "xo1zz1ca",
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
		new_title := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "1f9cj8ap",
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
		}`), new_title); err != nil {
			return err
		}
		collection.Schema.AddField(new_title)

		// add
		new_content := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "y6yex2ia",
			"name": "content",
			"type": "editor",
			"required": true,
			"presentable": false,
			"unique": false,
			"options": {
				"convertUrls": true
			}
		}`), new_content); err != nil {
			return err
		}
		collection.Schema.AddField(new_content)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("mlem1aae4up4ahh")
		if err != nil {
			return err
		}

		options := map[string]any{}
		if err := json.Unmarshal([]byte(`{
			"query": "SELECT posts.id,\nposts.` + "`" + `userId` + "`" + `,\nusers.` + "`" + `firstName` + "`" + `,\nusers.` + "`" + `lastName` + "`" + `,\nposts.title,\nposts.content,\nposts.created,\nposts.updated,\ntags.name as tag\nFROM posts \nINNER JOIN users on posts.` + "`" + `userId` + "`" + `= users.id\nLEFT JOIN tags on posts.id = tags.` + "`" + `postId` + "`" + `\n\n"
		}`), &options); err != nil {
			return err
		}
		collection.SetOptions(options)

		// add
		del_userId := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "yodr10xd",
			"name": "userId",
			"type": "relation",
			"required": true,
			"presentable": false,
			"unique": false,
			"options": {
				"collectionId": "_pb_users_auth_",
				"cascadeDelete": true,
				"minSelect": null,
				"maxSelect": 1,
				"displayFields": null
			}
		}`), del_userId); err != nil {
			return err
		}
		collection.Schema.AddField(del_userId)

		// add
		del_firstName := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "zxovtje0",
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
			"id": "1ot7ngjk",
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
		del_title := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "ulqequh5",
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
		}`), del_title); err != nil {
			return err
		}
		collection.Schema.AddField(del_title)

		// add
		del_content := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "uhjzphmz",
			"name": "content",
			"type": "editor",
			"required": true,
			"presentable": false,
			"unique": false,
			"options": {
				"convertUrls": true
			}
		}`), del_content); err != nil {
			return err
		}
		collection.Schema.AddField(del_content)

		// add
		del_tag := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "tnxx2rll",
			"name": "tag",
			"type": "text",
			"required": true,
			"presentable": false,
			"unique": false,
			"options": {
				"min": null,
				"max": 50,
				"pattern": ""
			}
		}`), del_tag); err != nil {
			return err
		}
		collection.Schema.AddField(del_tag)

		// remove
		collection.Schema.RemoveField("lmatkkhv")

		// remove
		collection.Schema.RemoveField("lzsccauv")

		// remove
		collection.Schema.RemoveField("xo1zz1ca")

		// remove
		collection.Schema.RemoveField("1f9cj8ap")

		// remove
		collection.Schema.RemoveField("y6yex2ia")

		return dao.SaveCollection(collection)
	})
}
