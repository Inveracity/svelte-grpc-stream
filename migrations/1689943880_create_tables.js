migrate((db) => {
  const snapshot = [
    {
      "id": "_pb_users_auth_",
      "created": "2023-07-07 17:54:05.136Z",
      "updated": "2023-07-09 19:53:56.502Z",
      "name": "users",
      "type": "auth",
      "system": false,
      "schema": [
        {
          "system": false,
          "id": "users_name",
          "name": "name",
          "type": "text",
          "required": false,
          "unique": false,
          "options": {
            "min": null,
            "max": null,
            "pattern": ""
          }
        },
        {
          "system": false,
          "id": "users_avatar",
          "name": "avatar",
          "type": "file",
          "required": false,
          "unique": false,
          "options": {
            "maxSelect": 1,
            "maxSize": 5242880,
            "mimeTypes": [
              "image/jpeg",
              "image/png",
              "image/svg+xml",
              "image/gif",
              "image/webp"
            ],
            "thumbs": null,
            "protected": false
          }
        }
      ],
      "indexes": [],
      "listRule": "@request.auth.id != \"\"",
      "viewRule": "id = @request.auth.id",
      "createRule": null,
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
        "requireEmail": false
      }
    },
    {
      "id": "5lbk3no6oi875z1",
      "created": "2023-07-09 19:53:56.503Z",
      "updated": "2023-07-09 19:53:56.503Z",
      "name": "channels",
      "type": "base",
      "system": false,
      "schema": [
        {
          "system": false,
          "id": "mywvanqq",
          "name": "name",
          "type": "text",
          "required": true,
          "unique": false,
          "options": {
            "min": 1,
            "max": 10,
            "pattern": "^\\w+$"
          }
        },
        {
          "system": false,
          "id": "zbku29qo",
          "name": "archived",
          "type": "bool",
          "required": false,
          "unique": false,
          "options": {}
        }
      ],
      "indexes": [
        "CREATE UNIQUE INDEX `idx_xVIc4hZ` ON `channels` (`name`)"
      ],
      "listRule": "@request.auth.id != \"\" && archived != true",
      "viewRule": "@request.auth.id != \"\" && archived != true",
      "createRule": "@request.auth.id != \"\"",
      "updateRule": "@request.auth.id != \"\"",
      "deleteRule": null,
      "options": {}
    }
  ];

  const collections = snapshot.map((item) => new Collection(item));

  return Dao(db).importCollections(collections, true, null);
}, (db) => {
  return null;
})
