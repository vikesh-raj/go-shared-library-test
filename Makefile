.PHONY: clean

libsample.so: main.go
	go build -o libsample.so -buildmode=c-shared main.go

libsample.a: main.go
	go build -o libsample.a -buildmode=c-archive main.go

test: test.c libsample.so
	gcc -o test test.c ./libsample.so

test_a: test.c libsample.a
	gcc -o test_a test.c ./libsample.a -lpthread

clean:
	rm -rf libsample.so libsample.a libsample.h test test_a