bin/log: main.c | bin
	gcc main.c -o bin/log

bin:
	mkdir -p bin
