flop-lang
=========

Flop is yet another hobby programming language.

Usage
-----

Build the interpreter with `go build` and pipe your program to it

    cat examples/hello.flop | ./flop-lang

Features
--------
Currently it supports all constructs needed to interpret a hello world program, like

    fn main() {
        print("Hello, world!\n");
    }
