package web_files_manipulation

var loginModifications = []Modification{
	Modification{
		Query:       "html",
		PrependHTML: loginHeader,
	},
	Modification{
		Query: "",
		AttributesChanges: []AttributeChange{
			NewAttributeChange(
				`.login_email input`,
				REPLACE_ATTRIBUTE,
				`name="email"`,
			),

			NewAttributeChange(
				``,
				REPLACE_ATTRIBUTE,
				``,
			),
			NewAttributeChange(
				``,
				REPLACE_ATTRIBUTE,
				``,
			),
			NewAttributeChange(
				``,
				REPLACE_ATTRIBUTE,
				``,
			),
			NewAttributeChange(
				``,
				REPLACE_ATTRIBUTE,
				``,
			),
			NewAttributeChange(
				``,
				REPLACE_ATTRIBUTE,
				``,
			),
			NewAttributeChange(
				``,
				REPLACE_ATTRIBUTE,
				``,
			),
		},
		HTMLChanges: []HTMLChange{
			HTMLChange{
				Query: "",
				Mode:  INNER_HTML,
				HTML:  ``,
			},
		},
	},
	Modification{
		Query: ".remove_item_on_update",
		AttributesChanges: []AttributeChange{
			AttributeChange{
				Query: "",
				Mode:  REPLACE_ATTRIBUTE,
				Attribute: Attribute{
					Name:  "style",
					Value: `display: none;`,
				},
			},
		},
	},
}
