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
			"query": "SELECT posts.id,\nposts.` + "`" + `userId` + "`" + `,\nusers.` + "`" + `firstName` + "`" + `,\nusers.` + "`" + `lastName` + "`" + `,\nposts.title,\nposts.content,\nposts.tags,\nposts.created,\nposts.updated\nFROM posts \nINNER JOIN users on posts.` + "`" + `userId` + "`" + `= users.id\n\n"
		}`), &options); err != nil {
			return err
		}
		collection.SetOptions(options)

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

		// add
		new_userId := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "c4gxhbns",
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
			"id": "ksajbanh",
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
			"id": "jaepl7rf",
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
			"id": "oig9ifm7",
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
			"id": "p5sx9y5s",
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

		// add
		new_tags := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "finjkqsv",
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

		// add
		del_userId := &schema.SchemaField{}
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
		}`), del_userId); err != nil {
			return err
		}
		collection.Schema.AddField(del_userId)

		// add
		del_firstName := &schema.SchemaField{}
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
		}`), del_firstName); err != nil {
			return err
		}
		collection.Schema.AddField(del_firstName)

		// add
		del_lastName := &schema.SchemaField{}
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
		}`), del_lastName); err != nil {
			return err
		}
		collection.Schema.AddField(del_lastName)

		// add
		del_title := &schema.SchemaField{}
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
		}`), del_title); err != nil {
			return err
		}
		collection.Schema.AddField(del_title)

		// add
		del_content := &schema.SchemaField{}
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
		}`), del_content); err != nil {
			return err
		}
		collection.Schema.AddField(del_content)

		// remove
		collection.Schema.RemoveField("c4gxhbns")

		// remove
		collection.Schema.RemoveField("ksajbanh")

		// remove
		collection.Schema.RemoveField("jaepl7rf")

		// remove
		collection.Schema.RemoveField("oig9ifm7")

		// remove
		collection.Schema.RemoveField("p5sx9y5s")

		// remove
		collection.Schema.RemoveField("finjkqsv")

		return dao.SaveCollection(collection)
	})
}
