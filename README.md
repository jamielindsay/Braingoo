# Braingoo

A brainf\*\*\* interpreter package for Go.

It uses a tape size of 30,000 8bit signed integers to handle all standard commands, including input and output.

## Install

Install the package to your Go workspace.

```shell
go get github.com/jamielindsay/Braingoo
```

## Use

Import the package to your file.

```Go
import "braingoo"
```

Create a new interpreter object, passing it the source code as a string and input as an int8 array.

```Go
source := ",.,.,."
input := []int8{'a', 'b', 'c'}
interpreter := braingoo.NewInterpreter(source, input)
```

If you have no input then simply pass an empty array.

```Go
braingoo.NewInterpreter(source, []int8{})
```

Call the parse method to execute the source code.

```Go
interpreter.parse()
```

Access the interpreters output via the output attribute.

```Go
output = interpreter.output
```
