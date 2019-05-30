[![GoDoc](https://godoc.org/github.com/nickwells/location.mod?status.png)](https://godoc.org/github.com/nickwells/location.mod)


# location.mod
A standard way of recording an indexed location in some source. For instance a line in a file.

The package is called `location`.

## How to use location
The L struct gives a common way of recording information about a location
within a source. You can add details of the content and a descriptive note.

You would typically create a location before processing a stream of data and
then on every natural boundary (such as a line in a file or a row in a
database query) you would increment the location. Then if you detect a
problem you can report the location simply.

## Errors
You can construct an error from a location and you will automatically have
the location of the error recorded.

## Chains
You can create chains of locations to reflect a chain of `include` directives
and check that there are no loops in the chain.

