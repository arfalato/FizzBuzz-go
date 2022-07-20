Fizz Buzz Test

Write a program that prints the numbers from 1 to 100.  
But for multiples of three print “Fizz” instead of the number and for the multiples of five print “Buzz”.  
For numbers which are multiples of both three and five print “FizzBuzz”.


- Run app  
```
docker build . -t fizzbuzz
``` 
and
```
docker run --rm fizzbuzz
```
output:   

```
12Fizz4BuzzFizz78FizzBuzz11Fizz1314FizzBuzz1617Fizz19BuzzFizz2223FizzBuzz26Fizz2829FizzBuzz3132Fizz34BuzzFizz3738FizzBuzz41Fizz4344FizzBuzz4647Fizz49BuzzFizz5253FizzBuzz56Fizz5859FizzBuzz6162Fizz64BuzzFizz6768FizzBuzz71Fizz7374FizzBuzz7677Fizz79BuzzFizz8283FizzBuzz86Fizz8889FizzBuzz9192Fizz94BuzzFizz9798FizzBuzz
```