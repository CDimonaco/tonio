//nolint lll
package formatters_test

import (
	"testing"

	"github.com/CDimonaco/tonio/internal/core"
	"github.com/CDimonaco/tonio/internal/core/formatters"
	"github.com/stretchr/testify/assert"
)

func TestJSONFormatterOutput(t *testing.T) {
	output, err := formatters.JSONMessage(core.Message{
		Body: []byte(`
		{
			"id": 192,
			"name": "Punk IPA 2007 - 2010",
			"tagline": "Post Modern Classic. Spiky. Tropical. Hoppy.",
			"first_brewed": "04/2007",
			"description": "Our flagship beer that kick started the craft beer revolution. This is James and Martin's original take on an American IPA, subverted with punchy New Zealand hops. Layered with new world hops to create an all-out riot of grapefruit, pineapple and lychee before a spiky, mouth-puckering bitter finish.",
			"image_url": "https://images.punkapi.com/v2/192.png",
			"abv": 6.0,
			"ibu": 60.0,
			"target_fg": 1010.0,
			"target_og": 1056.0,
			"ebc": 17.0,
			"srm": 8.5,
			"ph": 4.4,
			"attenuation_level": 82.14,
			"volume": {
			  "value": 20,
			  "unit": "liters"
			},
			"boil_volume": {
			  "value": 25,
			  "unit": "liters"
			},
			"method": {
			  "mash_temp": [
				{
				  "temp": {
					"value": 65,
					"unit": "celsius"
				  },
				  "duration": 75
				}
			  ],
			  "fermentation": {
				"temp": {
				  "value": 19.0,
				  "unit": "celsius"
				}
			  },
			  "twist": null
			},
			"ingredients": {
			  "malt": [
				{
				  "name": "Extra Pale",
				  "amount": {
					"value": 5.3,
					"unit": "kilograms"
				  }
				}
			  ],
			  "hops": [
				{
				  "name": "Ahtanum",
				  "amount": {
					"value": 17.5,
					"unit": "grams"
				   },
				   "add": "start",
				   "attribute": "bitter"
				 }
			  ],
			  "yeast": "Wyeast 1056 - American Ale™"
			},
			"food_pairing": [
			  "Spicy carne asada with a pico de gallo sauce",
			  "Shredded chicken tacos with a mango chilli lime salsa",
			  "Cheesecake with a passion fruit swirl sauce"
			],
			"brewers_tips": "While it may surprise you, this version of Punk IPA isn't dry hopped but still packs a punch! To make the best of the aroma hops make sure they are fully submerged and add them just before knock out for an intense hop hit.",
			"contributed_by": "Sam Mason <samjbmason>"
		  }
		`),
		ContentType: "application/json",
		Exchange:    "exchange_test",
		Queue:       "test_queue",
		RoutingKeys: []string{"routing"},
	})
	assert.NoError(t, err)
	assert.EqualValues(t, "{\n  \"abv\": 6.0,\n  \"attenuation_level\": 82.14,\n  \"boil_volume\": {\n    \"unit\": \"liters\",\n    \"value\": 25\n  },\n  \"brewers_tips\": \"While it may surprise you, this version of Punk IPA isn't dry hopped but still packs a punch! To make the best of the aroma hops make sure they are fully submerged and add them just before knock out for an intense hop hit.\",\n  \"contributed_by\": \"Sam Mason <samjbmason>\",\n  \"description\": \"Our flagship beer that kick started the craft beer revolution. This is James and Martin's original take on an American IPA, subverted with punchy New Zealand hops. Layered with new world hops to create an all-out riot of grapefruit, pineapple and lychee before a spiky, mouth-puckering bitter finish.\",\n  \"ebc\": 17.0,\n  \"first_brewed\": \"04/2007\",\n  \"food_pairing\": [\n    \"Spicy carne asada with a pico de gallo sauce\",\n    \"Shredded chicken tacos with a mango chilli lime salsa\",\n    \"Cheesecake with a passion fruit swirl sauce\"\n  ],\n  \"ibu\": 60.0,\n  \"id\": 192,\n  \"image_url\": \"https://images.punkapi.com/v2/192.png\",\n  \"ingredients\": {\n    \"hops\": [\n      {\n        \"add\": \"start\",\n        \"amount\": {\n          \"unit\": \"grams\",\n          \"value\": 17.5\n        },\n        \"attribute\": \"bitter\",\n        \"name\": \"Ahtanum\"\n      }\n    ],\n    \"malt\": [\n      {\n        \"amount\": {\n          \"unit\": \"kilograms\",\n          \"value\": 5.3\n        },\n        \"name\": \"Extra Pale\"\n      }\n    ],\n    \"yeast\": \"Wyeast 1056 - American Ale™\"\n  },\n  \"method\": {\n    \"fermentation\": {\n      \"temp\": {\n        \"unit\": \"celsius\",\n        \"value\": 19.0\n      }\n    },\n    \"mash_temp\": [\n      {\n        \"duration\": 75,\n        \"temp\": {\n          \"unit\": \"celsius\",\n          \"value\": 65\n        }\n      }\n    ],\n    \"twist\": null\n  },\n  \"name\": \"Punk IPA 2007 - 2010\",\n  \"ph\": 4.4,\n  \"srm\": 8.5,\n  \"tagline\": \"Post Modern Classic. Spiky. Tropical. Hoppy.\",\n  \"target_fg\": 1010.0,\n  \"target_og\": 1056.0,\n  \"volume\": {\n    \"unit\": \"liters\",\n    \"value\": 20\n  }\n}\n\n", output.String())
}
