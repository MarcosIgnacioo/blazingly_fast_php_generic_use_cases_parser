package web_files_manipulation

var ordersModifications = []Modification{
	Modification{
		Query:       "html",
		PrependHTML: ordersHeader,
	},
	Modification{
		Query: ".table-text table tbody",
		InnerHTML: `
		<?php if (isset($orders) && count($orders)) : ?>
		<?php foreach ($orders as $order) : ?>
			<tr>
				<td style="width: 17.42%;"><div style="text-align: center;">#<?= $order->folio ?></div></td>
				<td style="width: 19.2363%;">
					<div style="text-align: center;"><?= (new DateTime($order->order_date))->format('F d, Y') ?> <span style="white-space: pre;" id="isPasted">&nbsp; &nbsp;&nbsp;</span>&nbsp;</div>
				</td>
				<td style="width: 19.1538%;"><div style="text-align: center;"><?= $order->order_status->name ?></div></td>
				<td style="width: 20%; text-align: center;">
					<div style="text-align: center;">
						$ <?= number_format($order->total,2) ?> <br>
						<small>
							por <?= count($order->presentations) ?> artículos
						</small>
					</div>
				</td>
				<td style="width: 24.1899%;">
					<div style="text-align: center;">
						<a href="<?= BASE_PATH ?>account/order/<?= $order->folio ?>/">
							Detalles
						</a>
					</div>
				</td>
			</tr>
		<?php endforeach ?>
		<?php endif ?>
		`,
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
