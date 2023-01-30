```
$ go install github.com/avamsi/rd@latest
```
```shell
cd() {
	builtin cd $@ || builtin cd $(rd $@)
}
```
