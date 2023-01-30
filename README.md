```
$ go install github.com/avamsi/rd@latest
```
```shell
cd() {
	builtin cd $@ || {
		local d=$(rd $@) && builtin cd $d && print "rd: cd'ed to $d"
	}
}
```
