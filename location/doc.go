/*
Package location provides a standard way of recording an indexed location in
some source. For instance a line in a file.

The caller can set the source and then, as the entries are processed, the
index can be incremented. The content of the current location can be added
along with any notes. It is also possible to construct chains of locations to
reflect a series of file include directives. This chain can be checked to
ensure there are no loops.
*/
package location
