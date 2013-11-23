# proc-stat

[![GoDoc](http://godoc.org/github.com/guillermo/go.procstat?status.png)](http://godoc.org/github.com/guillermo/go.procstat)

Package proc-stat provides an interface for /proc/:pid/stat

See Stats structure for knowing the data that you can get, but consumend cpu, io and mem is there.

```golang
  	stats := stats.Stat{Pid: os.Getpid()}
  	err := stats.Update()
  	if err != nil {
  		panic(err)
  	}
```


*NOTES*: If the comm have a space in the middle, this program will fail to read all the arguments. Look in man proc for more info.

## Docs

Visit: http://godoc.org/github.com/guillermo/go.proc-stat

## LICENSE

BSD
