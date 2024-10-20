package migrations

import (
	"encoding/json"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/models/schema"
	"github.com/pocketbase/pocketbase/tools/types"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("p4f8mwphwdej2a0")
		if err != nil {
			return err
		}

		collection.ListRule = types.Pointer("")

		collection.ViewRule = types.Pointer("")

		// remove
		collection.Schema.RemoveField("covhgwol")

		// remove
		collection.Schema.RemoveField("rkfoyv4n")

		// remove
		collection.Schema.RemoveField("pz2wog3z")

		// add
		new_firstName := &schema.SchemaField{}
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
		}`), new_firstName); err != nil {
			return err
		}
		collection.Schema.AddField(new_firstName)

		// add
		new_lastName := &schema.SchemaField{}
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
		}`), new_lastName); err != nil {
			return err
		}
		collection.Schema.AddField(new_lastName)

		// add
		new_content := &schema.SchemaField{}
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

		collection.ListRule = nil

		collection.ViewRule = nil

		// add
		del_firstName := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
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
		}`), del_firstName); err != nil {
			return err
		}
		collection.Schema.AddField(del_firstName)

		// add
		del_lastName := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
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
		}`), del_lastName); err != nil {
			return err
		}
		collection.Schema.AddField(del_lastName)

		// add
		del_content := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
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
		}`), del_content); err != nil {
			return err
		}
		collection.Schema.AddField(del_content)

		// remove
		collection.Schema.RemoveField("giumlhyt")

		// remove
		collection.Schema.RemoveField("t4rcbtns")

		// remove
		collection.Schema.RemoveField("08pwwzjp")

		return dao.SaveCollection(collection)
	})
}
