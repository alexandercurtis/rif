# rif
rif is a little command line calculator I wrote to start learning the Go language.

The name comes from one of the three R's : readin, ritin and rifmetic

It performs arithmetic on its command parameters, in a reverse-polish fashion, where you specify the operands first then the operator. This is inspired by the Linux tool `dc`.

For example:

    rif 1 2 +
    3

    rif 10 1 -
    9

    rif 56 7 /
    8

    rif 6 7 x
    42

Note that * also works for multiplication, but will need to be escaped in most shells:

    rif 6 7 \*
    42

Numbers are pushed on to a stack, and operators work on the stack and push their result. The output of rif is the sum of the numbers on the stack.


The following operators are currently available:

+ removes all the numbers from the stack and pushes their sum, e.g.
    rif 1 2 3 4 5 6 +
    21

Since the output of rif is always the sum of the numbers of the stack, so in the above example the + can be omitted:
    rif 1 2 3 4 5 6
    21


x (or *) removes all the numbers from the stack and pushes their product, e.g.
    rif 1 2 3 4 5 6 x
    720


/ pops two numbers from the stack and pushes the result of dividing the first by the second
    rif 56 7 /
    8

- negates the number on the top of the stack and replaces it
    rif 1 -
    -1

    rif 12 4 -
    8

    rif 2 12 4 - +
    10

    rif 12 4 - 20 +
    28




