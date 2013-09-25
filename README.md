## What is it?

You have a directory with a bunch of Go code in it. It has some select statements in it, no doubt.

You wonder, as I often do: How many cases do those select statements have?

Fear not: `go run count.go -d <directory>` will tell you.

Running this on all the packages that were go gettable from http://godoc.org/-/index
as of Sep 25, 2013, yielded the following distribution:

```
Cases	Comm	Default	Count
0	0	0	95
1	0	1	10
1	1	0	322
2	1	1	1562
2	2	0	3018
3	2	1	97
3	3	0	464
4	3	1	16
4	4	0	197
5	4	1	2
5	5	0	64
6	6	0	30
7	7	0	14
8	8	0	10
9	9	0	6
10	10	0	1
12	12	0	1
14	14	0	2
```

* `Cases` is the total number of cases.
* `Comm` is the number of communication cases (involving channels).
* `Default` is the number of default cases (thankfully never > 1!).
* `Count` is the number of instances of such select statements.


## Acknowledgements

* Thanks to godoc.org for providing a lovely index for scraping bunches of Go.
* Thanks for @kr for `github.com/kr/fs`, vendored here.