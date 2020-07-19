run:
	cd src && go run . -m "Hello world"
test:
	cd src && go run . -d "./til" -e "michaelfromyeg@gmail.com" -n "Michael DeMarco" -m "This is (yet another) row" -u "https://michaeldemar.co"