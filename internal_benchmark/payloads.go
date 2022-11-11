package internal_benchmark

var PingPayload = `{
	"app_permissions": "521842576449",
	"application_id": "1033202906385625180",
	"channel_id": "1035162410912321546",
	"data": {
	  "id": "1040089416561070122",
	  "name": "ping",
	  "type": 1
	},
	"entitlement_sku_ids": [],
	"guild_id": "901426442242498650",
	"guild_locale": "en-US",
	"id": "1040472678085165088",
	"locale": "en-US",
	"member": {
	  "avatar": null,
	  "communication_disabled_until": null,
	  "deaf": false,
	  "flags": 0,
	  "is_pending": false,
	  "joined_at": "2021-10-23T11:06:51.533000+00:00",
	  "mute": false,
	  "nick": null,
	  "pending": false,
	  "permissions": "4398046511103",
	  "premium_since": "2022-10-24T12:06:54.776000+00:00",
	  "roles": [
		"1033704865764999188",
		"958593111527927869"
	  ],
	  "user": {
		"avatar": "b2e8fe6c07b71ce7df186abddb7fa792",
		"avatar_decoration": null,
		"discriminator": "0005",
		"id": "97470053615673344",
		"public_flags": 4194368,
		"username": "tristan"
	  }
	},
	"token": "brrr",
	"type": 2,
	"version": 1
  }`

type PingResponse struct {
	Type int `json:"type"`
	Data struct {
		Content string `json:"content"`
	} `json:"data"`
}
