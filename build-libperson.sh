gcc -o person.o -c -fPIC person.c
gcc -o libperson.so -shared person.o
file person.o libperson.so
gcc -o libperson.so -Wall -g -shared -fPIC person.c


