package web_files_manipulation

var detailsModifications = []Modification{
	Modification{
		Query:       "html",
		PrependHTML: detailsHeader,
	},
	anchorLoginModification,
	Modification{
		Query: "append_scripts_to_end_of_body",
		AppendHTML: `
		<script type="text/javascript">
            document.querySelector('.details_edit_form').addEventListener('submit', validateUpdate);
            document.querySelector('.password_form').addEventListener('submit', validateChangePassword);
		</script>
		`,
	},
	Modification{
		Query: "",
		AttributesChanges: []AttributeChange{

			NewAttributeChange(
				`.details_email input`,
				REPLACE_ATTRIBUTE,
				`name="email"`,
			),
			NewAttributeChange(
				`.details_email input`,
				REPLACE_ATTRIBUTE,
				`value="<?= $_SESSION['email'] ?>"`,
			),

			NewAttributeChange(
				`.details_name input`,
				REPLACE_ATTRIBUTE,
				`name="name"`,
			),
			NewAttributeChange(
				`.details_name input`,
				REPLACE_ATTRIBUTE,
				`value="<?= $_SESSION['name'] ?>"`,
			),

			NewAttributeChange(
				`.details_lastname input`,
				REPLACE_ATTRIBUTE,
				`name="lastname"`,
			),
			NewAttributeChange(
				`.details_lastname input`,
				REPLACE_ATTRIBUTE,
				`value="<?= $_SESSION['lastname'] ?>"`,
			),

			NewAttributeChange(
				`.details_phone input`,
				REPLACE_ATTRIBUTE,
				`phone="phone"`,
			),
			NewAttributeChange(
				`.details_phone input`,
				REPLACE_ATTRIBUTE,
				`value="<?= $_SESSION['phone'] ?>"`,
			),
			NewAttributeChange(
				`.details_phone input`,
				REPLACE_ATTRIBUTE,
				`name="phone"`,
			),

			NewAttributeChange(
				// NAME
				`.details_password_current input`,
				REPLACE_ATTRIBUTE,
				`name="password_current"`,
			),
			NewAttributeChange(
				// TYPE
				`.details_password_current input`,
				REPLACE_ATTRIBUTE,
				`type="password"`,
			),

			NewAttributeChange(
				// NAME
				`.details_password_1 input`,
				REPLACE_ATTRIBUTE,
				`name="password_1"`,
			),
			NewAttributeChange(
				// TYPE
				`.details_password_1 input`,
				REPLACE_ATTRIBUTE,
				`type="password"`,
			),

			NewAttributeChange(
				// NAME
				`.details_password_2 input`,
				REPLACE_ATTRIBUTE,
				`name="password_2"`,
			),
			NewAttributeChange(
				// TYPE
				`.details_password_2 input`,
				REPLACE_ATTRIBUTE,
				`type="password"`,
			),
		},
	},
}
