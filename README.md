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

`+` pops two numbers from the stack and pushes their sum, e.g.

    1 2 +
    3

    1 2 3 +
    5

`S` removes all the numbers from the stack and pushes their sum, e.g.

    rif 1 2 3 4 5 6 +
    21

Since the output of rif is always the sum of the numbers of the stack, so in the above example the `P` can be omitted:

    rif 1 2 3 4 5 6
    21


`x` (or `*`) pops two numbers from the stack pushes their product, e.g.

    rif 2 3 x
    6

    rif 2 3 4 x
    12

`P` removes all the numbers from the stack and pushes their product, e.g.

    rif 1 2 3 4 5 6 x
    720

`/` pops two numbers from the stack and pushes the result of dividing the second one popped by the first.

    rif 56 7 /
    8

`-` pops two numbers from the stack and pushes the result of subtracting the first one popped from the second.

    rif 12 4 -
    8

    rif 12 4 - 8 +
    16

    rif 12 4 - 3 x
    24
