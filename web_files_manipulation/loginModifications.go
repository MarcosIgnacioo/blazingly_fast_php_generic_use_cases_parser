package web_files_manipulation

var loginModifications = []Modification{
	Modification{
		Query:       "html",
		PrependHTML: loginHeader,
	},
	Modification{
		Query: "append_scripts_to_end_of_body",
		AppendHTML: `
		<script type="text/javascript">
				document.querySelector('.login_form').addEventListener('submit', validateAccess);
				document.querySelector('.register_form').addEventListener('submit', validateRegister);
		</script>
		`,
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
				`.login_password input`,
				REPLACE_ATTRIBUTE,
				`name="password"`,
			),

			NewAttributeChange(
				`.login_password input`,
				REPLACE_ATTRIBUTE,
				`type="password"`,
			),

			NewAttributeChange(
				`.register_email input`,
				REPLACE_ATTRIBUTE,
				`name="email"`,
			),

			NewAttributeChange(
				`.register_password input`,
				REPLACE_ATTRIBUTE,
				`name="password"`,
			),
			NewAttributeChange(
				`.register_password input`,
				REPLACE_ATTRIBUTE,
				`type="password"`,
			),

			NewAttributeChange(
				`.register_password_confirmation input`,
				REPLACE_ATTRIBUTE,
				`name="password_confirmation"`,
			),
			NewAttributeChange(
				`.register_password_confirmation input`,
				REPLACE_ATTRIBUTE,
				`type="password"`,
			),
			NewAttributeChange(
				`.register_name input`,
				REPLACE_ATTRIBUTE,
				`name="name"`,
			),
			NewAttributeChange(
				`.register_phone input`,
				REPLACE_ATTRIBUTE,
				`name="phone"`,
			),
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
