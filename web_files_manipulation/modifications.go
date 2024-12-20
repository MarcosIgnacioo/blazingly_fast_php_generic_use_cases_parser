package web_files_manipulation

var scriptsHead = `
	<link href="<?= BASE_PATH ?>js/toasty/dist/toasty.min.css" rel="stylesheet">

	<script type="text/javascript">
        var global_token = '<?= $_SESSION['token'] ?>';
        var global_url = '<?= BASE_PATH ?>';
        var le = <?= (DEVELOP_MODE) ?>;
        var lg = <?= (isset($_SESSION['client_id']) && $_SESSION['client_id'] != "") ? 'true' : 'false' ?>; 
    </script> 
	`
var scriptsBody = `
	<input type="hidden" id="g-recaptcha-response" name="g-recaptcha-response">

	<script src="https://unpkg.com/axios@1.1.2/dist/axios.min.js"></script>
	<script src="<?= BASE_PATH ?>js/toasty/dist/toasty.min.js" type="text/javascript"></script>
	<script src="<?= BASE_PATH ?>js/functions.js"></script>
    <script src="<?= BASE_PATH ?>js/shop.js"></script>
    <script src="<?= BASE_PATH ?>js/auth.js"></script>
	
	<script src='https://www.google.com/recaptcha/api.js?render=<?= PUBLIC_KEY_GG ?>'></script>
	<script>
		grecaptcha.ready(function() {
		grecaptcha.execute('<?= PUBLIC_KEY_GG ?>', {action: 'homepage'})
			.then(function(token) {
			document.getElementById('g-recaptcha-response').value=token;
		});
		});
	</script>
	`

var modifications map[string][]Modification = map[string][]Modification{
	"app":        rootDirModifications,
	"dashboard":  dashboardModifications,
	"tienda":     storeModifications,
	"cafe":       coffeeModifications,
	"merch":      merchModifications,
	"accesorios": accessoriesModifications,
	"diablo":     devilModifications,
	"sudadera":   sweaterModifications,
	"login":      loginModifications,
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
