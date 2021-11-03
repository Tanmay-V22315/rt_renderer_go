clean:
	rm build/*
build:
	go build -o build/renderer .
br:
	go build -o build/renderer . 
	./build/renderer > build/test_img.ppm
	cd build && xdg-open test_img.ppm
	cd ..