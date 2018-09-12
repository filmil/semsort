# semsort

A program that sorts semantic versions given on the `os.Stdin` and prints them
out, one line at a time, to `os.Stdout`. Comes in handy for use from inside
bash shell scripts.

# Installation

```console
go get github.com/filmil/semsort
go install github.com/filmil/semsort/...
```

# Testing

```console
go test github.com/filmil/semsort/...
```

# Usage

```console
echo "4.5.6 1.2.3" | semsort
1.2.3
4.5.6
```

Also:

```console
echo "1.2.3 1.2.3-rc.1" | semsort
1.2.3-rc.1
1.2.3
```

And: 

```console
$ cat <<EOF | semsort
0.1.1
0.1.0
0.1.0-alpha
0.1.0-alpha.2
EOF

0.1.0-alpha
0.1.0-alpha.2
0.1.0
0.1.1
```


