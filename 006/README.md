## [Problem 006](https://projecteuler.net/problem=6)

The sum of the squares of the first ten natural numbers is, . The square of the sum of the first ten natural numbers is, . Hence the absolute difference between the sum of the squares of the first ten natural numbers and the square of the sum is .

Find the absolute difference between the sum of the squares of the first  natural numbers and the square of the sum.

#### Input Format

First line contains  that denotes the number of test cases. This is followed by  lines, each containing an integer, .

Constraints

#### Output Format

Print the required answer for each test case.

#### Sample Input

$2$
$3$
$10$

#### Sample Output

$22$
$2640$
### Explanation 

+ For , `N = 3`, we have
   &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; $n = (1 + 2 + 3)^2, =  36$
   &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; $m = {1^2  + 2^2 + 3^2} = 11$
   &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; $n - m = 25$

+ For , `N = 10`, we have
   &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; $n = (1 + 2 + ... + 10)^2, =  3025$
   &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; $m = {1^2  + 2^2 + ... + 10^2} = 385$
   &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; $n - m = 2640$
 