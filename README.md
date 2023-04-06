```
$ go install github.com/avamsi/rd@latest
```

```shell
cd() {
	# Keep the original cd error hidden for if rd succeeds below.
	builtin cd $@ 2>/tmp/rd-cde || {
		# Note: don't try to "fix" this by adding `local`, it doesn't work.
		# For whatever reason, the `&&` chaining behaviour is lost with it.
		d=$(rd $@) && builtin cd $d || {
			# No luck, show the original error as well.
			cat /tmp/rd-cde
		}
	}
}
```
