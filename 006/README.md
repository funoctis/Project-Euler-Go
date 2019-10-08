## [Problem 006](https://projecteuler.net/problem=6)

The sum of the squares of the first ten natural numbers is <code>1<sup>2</sup> + 2<sup>2</sup> + ... + 10<sup>2</sup> = 385</code>. The square of the sum of the first ten natural numbers is, <code>(1 + 2 + ... + 10)<sup>2</sup> = 55<sup>2</sup> = 3025</code>. Hence the absolute difference between the sum of the squares of the first ten natural numbers and the square of the sum is `3025 - 385 = 2640`.

Find the absolute difference between the sum of the squares of the first `N` natural numbers and the square of the sum.

#### Input Format

First line contains `T` that denotes the number of test cases. This is followed by `T` lines, each containing an integer `N`.

#### Output Format

Print the required answer for each test case.

#### Sample Input

    2
    3
    10

#### Sample Output

    22
    2640

### Explanation

+ For `N = 3`, <code>(1 + 2 + 3)<sup>2</sup> - (1<sup>2</sup> + 2<sup>2</sup> + 3<sup>2</sup>) = 22</code>
+ For `N = 10`, <code>(1 + 2 + ... + 10)<sup>2</sup> - (1<sup>2</sup> + 2<sup>2</sup> + ... + 10<sup>2</sup>) = 2640</code>