# calexpgo

evaluate expression in reverse polish notation

examples :
```shell
> calexpgo 1 2 3 + -
4
> calexpgo 10 15 * ; 15 5 /
150
3
> calexpgo 3 DUP *
9
```
[spec](https://exercism.org/tracks/go/exercises/forth)

Values are integer.
Operators : 
* +: addition,
* -: substraction,
* *: multiplication,
* /: integer division,
* %: modulo,
* dup : duplicate top of stack,
* swap : swap top and second of stack,
* drop : remove top of stack, 
* over : add second of stack in top

Opérators are case insensitive
