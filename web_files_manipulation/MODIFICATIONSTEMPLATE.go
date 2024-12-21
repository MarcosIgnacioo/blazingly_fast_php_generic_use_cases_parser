package web_files_manipulation

//
// var HEADER = ``
//
// var NAMEModifications = []Modification{
// 	Modification{
// 		Query:       "html",
// 		PrependHTML: HEADER,
// 	},
// 	Modification{
// 		Query: "",
// 		HTMLChanges: []HTMLChange{
// 			HTMLChange{
// 				Query: "",
// 				Mode:  INNER_HTML,
// 				HTML:  ``,
// 			},
// 		},
// 		AttributesChanges: []AttributeChange{
// 			NewAttributeChange(
// 				`product_item_cart`,
// 				APPEND_ATTRIBUTE,
// 				`class="item_<?= $product->id ?>_<?= str_replace(" ", "", $product->feature) ?>"`,
// 			),
// 			NewAttributeChange(
// 				``,
// 				REPLACE_ATTRIBUTE,
// 				``,
// 			),
// 		},
// 	},
// 	Modification{
// 		Query: ".remove_item_on_update",
// 		AttributesChanges: []AttributeChange{
// 			AttributeChange{
// 				Query: "",
// 				Mode:  REPLACE_ATTRIBUTE,
// 				Attribute: Attribute{
// 					Name:  "style",
// 					Value: `display: none;`,
// 				},
// 			},
// 		},
// 	},
// }
