run:
	cd src && go run . -m "Hello world"
test:
	cd src && go run . -d "./til" -e "michaelfromyeg@gmail.com" -n "Michael DeMarco" -m "This is another row to double check the date computation is working, and the page updates!" -u "https://twitter.com/michaelfromyeg"