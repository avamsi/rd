```
$ go install github.com/avamsi/rd@latest
```
```shell
cd() {
	builtin cd $@ || {
		local p=$(rd $@) && builtin cd $p && print "rd: cd'ed to $p"
	}
}
```
