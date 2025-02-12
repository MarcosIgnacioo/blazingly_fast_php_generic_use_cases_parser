if [[ $1 == "b" ]]; then
   echo "BUILDING PROJECT"
   go run main.go
else
   echo "NOT BUILDING PROJECT JUST COPYING MOST RECENT COMPILED VERSION"
fi
echo "removiendo contenido de /srv/http"
sudo rm -r /srv/http/*
echo "copiando doce40BUILD a srv/http"
sudo cp -r ./doce40BUILD/*  /srv/http/
echo "removiendo htaccess de srv/http"
sudo rm /srv/http/.htaccess
echo "copiando htacces a srv/http"
sudo cp ./htaccess  /srv/http/.htaccess
