// Package fatinterfaces is a utility package to make it easier to consume
// common types as interfaces. This package is released in to the public domain,
// to make it easier for users to consume the package or to copy code contained
// herein.
//
// This is to make it easier to provide mocked types for testing, or to replace
// the functionality of your application without needing to update a bunch of
// references.
//
// This package defines a set of fat interfaces (interfaces with lots of
// methods) that describes things like `*os.File` and `*net/http.Request`.
//
// Instead of having your code take an *os.File, instead you could receive a
// fatinterfaces.OsFile interface. This would allow you to either use something
// like `*os.File`, or to use something like Afero.
package fatinterfaces
