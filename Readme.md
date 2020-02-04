## Tasks

1. The Fibonacci sequence is defined as a sequence of integers starting with 1 and 1, where each
subsequent value is the sum of the preceding two. I.e.
f(0) = 1
f(1) = 1
f(n) = f(n-1) + f(n-2) where n >= 2
Write a program in a language of your choice to calculate the sum of the first 100 even-valued
Fibonacci numbers


2. Write a function in a language of your choice which, when passed two sorted arrays of integers
returns an array of any numbers which appear in both. No value should appear in the returned
array more than once.


3. Write a function in a language of your choice which, when passed a positive integer returns
true if the decimal representation of that integer contains no odd digits and otherwise returns
false.


4. Write a function in a language of your choice which, when passed a decimal digit X, returns the
value of X+XX+XXX+XXXX. E.g. if the supplied digit is 3 it should return 3702
(3+33+333+3333).


### Notes

For the first task (even Fibonnacci values) :

- It seems that in Golang the highest number that is 
available for common types such as int64 is ``9223372036854775807`` so concretely after n=30 it started returning inconsistent 
values (negative numbers)

- To remediate that I had to use a 
special implementation with Big.Int fibonnacci version documented on the github repository issues : https://github.com/golang/go/issues/30943

### Tests 

Tests were provided for each task except from the first one which involves 
storing values of the Fibonnacci sequence. I considered this was out of the scope of this test.

### How to use 

``make test``

``make build``

``./vaion``
