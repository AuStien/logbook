bin/main: main.c | bin
	gcc main.c -o bin/main

bin:
	mkdir -p bin
