.PHONY: all build clean

INC ?= $(wildcard *.h)
SRC ?= $(wildcard *.c)
OBJ ?= $(SRC:.c=.o)
RUN ?= aliens


all:    build

build:  $(OBJ) $(INC)
	gcc $(OBJ) -lraylib -lGL -lm -lpthread -ldl -lrt -lX11 -o $(RUN)

clean:
	rm -f $(OBJ) $(RUN)

%.o: %.c
	+gcc -c -O2 $<
