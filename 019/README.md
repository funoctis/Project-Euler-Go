## [Problem 019](https://projecteuler.net/problem=19) 

You are given the following information, but you may prefer to do some research for yourself.

1 Jan 1900 was a Monday.
Thirty days has September,
April, June and November.
All the rest have thirty-one,
Saving February alone,
Which has twenty-eight, rain or shine.
And on leap years, twenty-nine.

A leap year occurs on any year evenly divisible by `4`, but not on a century unless it is divisible by `400`.

How many Sundays fell on the first of the month between two dates(both inclusive)?

#### Input Format

The first line contains an integer `T`, i.e., number of test cases.
Each testcase will contain two lines
<code>Y<sub>1</sub> M<sub>1</sub> D<sub>1</sub></code> on first line denoting starting date.
<code>Y<sub>2</sub> M<sub>2</sub> D<sub>2</sub></code> on second line denoting ending date.

#### Output Format

Print the values corresponding to each test case.

##### Sample Input

    2
    1900 1 1
    1910 1 1
    2000 1 1
    2020 1 1

#### Sample Output

    18
    35
### Explanation

For testcase 1, we have the following sundays :-

    1 April 1900
    1 July 1900
    1 September 1901
    1 December 1901
    1 June 1902
    1 February 1903
    1 March 1903
    1 November 1903
    1 May 1904
    1 January 1905
    1 October 1905
    1 April 1906
    1 July 1906
    1 September 1907 
    1 December 1907
    1 March 1908
    1 November 1908
    1 August 1909