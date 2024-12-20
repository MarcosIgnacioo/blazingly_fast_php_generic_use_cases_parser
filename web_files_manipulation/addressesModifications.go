package web_files_manipulation

var addressesModifications = []Modification{
	Modification{
		Query:       "html",
		PrependHTML: addressesHeader,
	},
	anchorLoginModification,
	Modification{
		Query: ".address-data",
		PrependHTML: `
		<?php if (isset($addresses) && count($addresses)) : ?>
			<?php foreach ($addresses as $address) : ?>
		`,
		AppendHTML: `
			<?php endforeach ?>
		<?php endif ?>
		`,
		HTMLChanges: []HTMLChange{
			HTMLChange{
				Query: ".inner .ed-element",
				Mode:  INNER_HTML,
				HTML: `
				<p id="isPasted"><strong>Nombre:&nbsp;</strong></p>
				<p><span style="text-align: right; font-size: 1rem; letter-spacing: 0px;">
					<?= $address->first_name ?> <?= $address->last_name ?>
				</span></p>
				<p><strong>Calle:</strong><span style="white-space: pre;">&nbsp; &nbsp;&nbsp;</span></p>
				<p><?= $address->street_and_use_number ?></p>
				<p>
					<strong>Apartamento:</strong><span style="white-space: pre;"><strong>&nbsp; &nbsp;&nbsp;</strong></span>
				</p>
				<p><span style="white-space: pre;"><?= $address->apartment ?></span></p>
				<p>
					<strong>Provincia:</strong><span style="white-space: pre;"><strong>&nbsp;</strong> &nbsp;&nbsp;</span>
				</p>
				<p><?= $address->city ?> <?= $address->state ?></p>
				<p><strong>Código postal:</strong><span style="white-space: pre;">&nbsp; &nbsp;&nbsp;</span></p>
				<p><?= $address->postal_code ?> <?= $address->country->code ?></p>
				<p>
					<strong>Teléfono:</strong><span style="white-space: pre;"><strong>&nbsp;</strong> &nbsp;&nbsp;</span>
				</p>
                <p><span style="white-space: pre;"><?= $address->phone_number ?></span></p>
				<p><strong>Paí</strong><strong>s:</strong></p>
				<p><?= $address->references ?></p>
				`,
			},
			HTMLChange{
				Query: ".inner",
				Mode:  APPEND_HTML,
				HTML: `
				<style>
						#ed-456888664 > .inner {
							padding: 10px;
						}
				</style>
				`,
			},
		},
	},
}
