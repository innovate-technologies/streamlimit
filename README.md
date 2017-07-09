Go Streamlimit
==============

This package allows you to gave a an io.Reader and io.Writer acting as a delyed io.Pipe. This is handy if you want to output data at a specific bitrate.

## How to use it

```
// This sets up a pipe that outputs 8 BYTES per second, 4 TIMES per SECOND
limiter := streamlimit.New(8, 4)

// Add data to the pipe
limiter.Write([]byte("Hello World"))

// Start the output of data
limiter.Start()

// Just read the data back out as in io.Reader
b := make([]byte,4)
limiter.Read(b)

```