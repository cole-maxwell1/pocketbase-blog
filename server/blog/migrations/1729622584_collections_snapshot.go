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
		jsonData := `[
			{
				"id": "_pb_users_auth_",
				"created": "2024-10-19 15:45:14.517Z",
				"updated": "2024-10-21 20:01:41.913Z",
				"name": "users",
				"type": "auth",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "users_name",
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
						"id": "users_avatar",
						"name": "avatar",
						"type": "file",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"mimeTypes": [
								"image/jpeg",
								"image/png",
								"image/svg+xml",
								"image/gif",
								"image/webp"
							],
							"thumbs": null,
							"maxSelect": 1,
							"maxSize": 5242880,
							"protected": false
						}
					},
					{
						"system": false,
						"id": "p2ekiew2",
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
						"id": "vvehousl",
						"name": "bio",
						"type": "editor",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"convertUrls": false
						}
					}
				],
				"indexes": [],
				"listRule": "id = @request.auth.id",
				"viewRule": "id = @request.auth.id",
				"createRule": "",
				"updateRule": "id = @request.auth.id",
				"deleteRule": "id = @request.auth.id",
				"options": {
					"allowEmailAuth": true,
					"allowOAuth2Auth": true,
					"allowUsernameAuth": true,
					"exceptEmailDomains": null,
					"manageRule": null,
					"minPasswordLength": 8,
					"onlyEmailDomains": null,
					"onlyVerified": false,
					"requireEmail": false
				}
			},
			{
				"id": "vsqcudq8utc0c9k",
				"created": "2024-10-19 15:48:56.518Z",
				"updated": "2024-10-21 19:24:15.749Z",
				"name": "posts",
				"type": "base",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "yh7ztwbf",
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
					},
					{
						"system": false,
						"id": "ugmg5t0b",
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
						"id": "e2ec0ybl",
						"name": "content",
						"type": "editor",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"convertUrls": true
						}
					},
					{
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
					}
				],
				"indexes": [],
				"listRule": null,
				"viewRule": "userId = @request.auth.id",
				"createRule": "userId = @request.auth.id",
				"updateRule": "userId = @request.auth.id",
				"deleteRule": "userId = @request.auth.id",
				"options": {}
			},
			{
				"id": "mlem1aae4up4ahh",
				"created": "2024-10-19 15:52:30.772Z",
				"updated": "2024-10-21 19:24:15.771Z",
				"name": "posts_author_vw",
				"type": "view",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "ryomdivu",
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
					},
					{
						"system": false,
						"id": "jec9ltqz",
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
						"id": "lpzgyge0",
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
						"id": "gkcrgnmw",
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
						"id": "d0uinjrz",
						"name": "content",
						"type": "editor",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"convertUrls": true
						}
					},
					{
						"system": false,
						"id": "wpqkrou1",
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
					}
				],
				"indexes": [],
				"listRule": "",
				"viewRule": "",
				"createRule": null,
				"updateRule": null,
				"deleteRule": null,
				"options": {
					"query": "SELECT posts.id,\nposts.` + "`" + `userId` + "`" + `,\nusers.` + "`" + `firstName` + "`" + `,\nusers.` + "`" + `lastName` + "`" + `,\nposts.title,\nposts.content,\nposts.tags,\nposts.created,\nposts.updated\nFROM posts \nINNER JOIN users on posts.` + "`" + `userId` + "`" + `= users.id\n\n"
				}
			},
			{
				"id": "9z1q3a89d8jl95z",
				"created": "2024-10-19 18:36:06.387Z",
				"updated": "2024-10-19 18:40:31.510Z",
				"name": "comments",
				"type": "base",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "6ep2lcox",
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
					},
					{
						"system": false,
						"id": "a9ex7jjs",
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
					},
					{
						"system": false,
						"id": "gxsvel5p",
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
				"listRule": "",
				"viewRule": "",
				"createRule": "userId = @request.auth.id",
				"updateRule": "userId = @request.auth.id",
				"deleteRule": "userId = @request.auth.id",
				"options": {}
			},
			{
				"id": "33q7jo0z4z6q8dd",
				"created": "2024-10-20 14:43:01.038Z",
				"updated": "2024-10-21 19:24:15.740Z",
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
			},
			{
				"id": "p4f8mwphwdej2a0",
				"created": "2024-10-20 18:40:22.155Z",
				"updated": "2024-10-21 19:24:15.795Z",
				"name": "comments_author_vw",
				"type": "view",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "efaobcby",
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
					},
					{
						"system": false,
						"id": "fytmbmzu",
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
						"id": "cl9jon7g",
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
						"id": "ebgx6tf6",
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
				"listRule": "",
				"viewRule": "",
				"createRule": null,
				"updateRule": null,
				"deleteRule": null,
				"options": {
					"query": "SELECT\ncomments.id,\ncomments.` + "`" + `postId` + "`" + `,\nusers.` + "`" + `firstName` + "`" + `,\nusers.` + "`" + `lastName` + "`" + `,\ncomments.content,\ncomments.created,\ncomments.updated\nFROM comments\nINNER JOIN users on comments.` + "`" + `userId` + "`" + ` = users.id"
				}
			},
			{
				"id": "90kja2o29ekev0z",
				"created": "2024-10-21 20:01:04.539Z",
				"updated": "2024-10-21 20:02:18.391Z",
				"name": "authors_vw",
				"type": "view",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "wydngo5n",
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
						"id": "yafrqer1",
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
						"id": "kllhxqnu",
						"name": "username",
						"type": "text",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "h8oqbzak",
						"name": "bio",
						"type": "editor",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"convertUrls": false
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
					"query": "SELECT users.id,\nusers.` + "`" + `firstName` + "`" + `,\nusers.` + "`" + `lastName` + "`" + `,\nusers.username,\nusers.bio\nFROM users"
				}
			}
		]`

		collections := []*models.Collection{}
		if err := json.Unmarshal([]byte(jsonData), &collections); err != nil {
			return err
		}

		return daos.New(db).ImportCollections(collections, true, nil)
	}, func(db dbx.Builder) error {
		return nil
	})
}
