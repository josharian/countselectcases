## What is it?

You have a directory with a bunch of Go code in it. It has some select statements in it, no doubt.

You wonder, as I often do: How many cases do those select statements have?

Fear not: `go run count.go -d <directory>` will tell you.

Running this on all the packages that were go gettable from http://godoc.org/-/index
as of Sep 25, 2013, yielded the following distribution:

```
Cases	Comm	Default	Count
0		0		0		95
1		0		1		10
1		1		0		322
2		1		1		1562
2		2		0		3018
3		2		1		97
3		3		0		464
4		3		1		16
4		4		0		197
5		4		1		2
5		5		0		64
6		6		0		30
7		7		0		14
8		8		0		10
9		9		0		6
10		10		0		1
12		12		0		1
14		14		0		2
```

* `Cases` is the total number of cases.
* `Comm` is the number of communication cases (involving channels).
* `Default` is the number of default cases (thankfully never > 1!).
* `Count` is the number of instances of such select statements.

Update: Using the godoc corpus from Jun 2014 godoc:

```
Cases	Comm	Default	Count
0	0	0	224
1	0	1	15
1	1	0	793
2	1	1	4440
2	2	0	8727
3	2	1	468
3	3	0	1672
4	3	1	270
4	4	0	492
5	4	1	233
5	5	0	190
6	5	1	134
6	6	0	87
7	6	1	57
7	7	0	40
8	7	1	16
8	8	0	22
9	8	1	3
9	9	0	13
10	10	0	6
11	11	0	2
12	12	0	2
14	14	0	5
15	15	0	4
16	16	0	1
17	17	0	2
22	22	0	1
```

## Acknowledgements

* Thanks to godoc.org for providing a lovely index for scraping bunches of Go.
* Thanks for @kr for `github.com/kr/fs`, vendored here.
