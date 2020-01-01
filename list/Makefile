gcc_options = -std=c++17 -Wall --pedantic-errors

.PHONY : run clean

program : main.cpp
	g++ $(gcc_options) -o $@ $<

run : program
	./program

clean :
	rm -f ./program
