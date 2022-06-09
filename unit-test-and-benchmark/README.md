# intro
Go's built-in support for unit testing makes it easier to test as you go.

# how to write unit test in go
create a file with a suffix `_test.go`, see [example](unittest)。

# how to write a test case
Arrange Act Assert, or Given When Then, see [example](aaa)

# assert
usually we would use assert in unit test, see [example](assert)
> ps1. foodpanda does not rule that you must use assert see these two real repo, [with assert](https://github.com/deliveryhero/mikro/blob/1cce665b744f51951834d6a872dfb154cac73438/pkg/apprunner/options_test.go#L40) and [without assert](https://github.com/deliveryhero/mikro/blob/1cce665b744f51951834d6a872dfb154cac73438/pkg/admin/handler/root_test.go#L62-L71)

# mock
we may talk about mock later with interface

# how to write function unit test for a function
  * a test case per unit test func vs all test cases within a unit test func, see [example](testcases)
    * long/short test func name
    * ease/hard to test specific test case
    * clean/complex code in arrange block

# parallel
when some func take too long time to run, use parallel, see [example](parallel)

# bench mark
see [example](benchmark)