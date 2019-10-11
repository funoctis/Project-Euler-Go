## [Problem 008](https://projecteuler.net/problem=8)

Find the greatest product of  consecutive digits in the  digit number.

#### Input Format

First line contains `T` that denotes the number of test cases.
First line of each test case will contain two integers `N` & `K`.
Second line of each test case will contain a `N` digit integer.


#### Output Format

Print the required answer for each test case.

#### Sample Input 

    2
    10 5
    3675356291
    10 5
    2709360626

#### Sample Output

    3150
    0
### Explanation

+ For `3675356291` and selecting `K = 5` consequetive digits, we have `36753`, `67535`, `75356`, `53562`, `35629` and `56291`. Where `6 x 7 x 5 x 3 x 5` gives maximum product as `3150` 
+ For `2709360626`, `0` lies in all selection of `5` consequetive digits hence maximum product remains `0`