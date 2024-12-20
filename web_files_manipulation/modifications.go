package web_files_manipulation

var modifications map[string][]Modification = map[string][]Modification{
	"app":        rootDirModifications,
	"dashboard":  dashboardModifications,
	"tienda":     storeModifications,
	"cafe":       coffeeModifications,
	"merch":      merchModifications,
	"accesorios": accessoriesModifications,
	"diablo":     devilModifications,
	"sudadera":   sweaterModifications,
}

var IDS map[string]string = map[string]string{
	"": "",
}

var dashboardModifications = []Modification{
	Modification{
		Query:     `a[href="/login"]`,
		SelectAll: true,
		AttributesChanges: []AttributeChange{
			AttributeChange{
				Mode: REPLACE_ATTRIBUTE,
				Attribute: Attribute{
					Name:  "href",
					Value: `/clients?action=logout`,
				},
			},
		},
	},
}
