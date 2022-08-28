# Watts-Up Power Meter API in Go

Returns an object which can read/log watts-up power meter values. Define the port that power meter is connected to (typicall ttyUSB0) and path to write results.

This API depends on the [watts-up meter executeable](https://github.com/pyrovski/watts-up). Clone the repo and build the executeable for target machine.

### Example Usage
``` go
    port := "ttyUSB0"       // port address physical meter is connected to
	command := []string{    // arguments to run with wattsup executeable
        "./wattsup", 
        port,
        "-g",
        "watts"
    }
	filename := "out.watts" // path to file for text output
	w := Wattsup{}
	if err := w.Init(port, filename, command); err != nil {
		w.Start()
	}

    // A new file should be created with `filename` and
    // output will be piped to that file.
```