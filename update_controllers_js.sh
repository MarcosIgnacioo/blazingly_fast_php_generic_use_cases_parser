DOCE40=/home/marcig/work/doce40
BOT=/home/marcig/work/blazingly_fast_php_generic_use_cases_parser
env --chdir=$DOCE40 -S git pull
cp $DOCE40/js/auth.js $BOT/extra_js
cp $DOCE40/js/shop.js $BOT/extra_js
cp $DOCE40/js/functions.js $BOT/extra_js
rm -r $BOT/extra_js/toasty
cp -r $DOCE40/js/toasty $BOT/extra_js

cp $DOCE40/app/AuthController.php $BOT/controllers
cp $DOCE40/app/CategoriesController.php $BOT/controllers
cp $DOCE40/app/config.php $BOT/controllers
cp $DOCE40/app/ContactController.php $BOT/controllers
cp $DOCE40/app/ProductsController.php $BOT/controllers
cp $DOCE40/app/session_data.php $BOT/controllers
cp $DOCE40/app/ShopController.php $BOT/controllers
cp $DOCE40/app/ToolsController.php $BOT/controllers
