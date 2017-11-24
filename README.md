# Fat Interfaces
[![License](https://img.shields.io/github/license/theckman/fatinterfaces.svg)](https://github.com/theckman/fatinterfaces/blob/master/LICENSE)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/theckman/fatinterfaces)
[![Latest Git Tag](https://img.shields.io/github/tag/theckman/fatinterfaces.svg)](https://github.com/theckman/fatinterfaces/releases)
[![Travis master Build Status](https://img.shields.io/travis/theckman/fatinterfaces/master.svg?label=TravisCI)](https://travis-ci.org/theckman/fatinterfaces/branches)

Package fatinterfaces is a utility package to make it easier to consume common
types as interfaces.

This is to make it easier to provide mocked types for testing, or to replace the
functionality of your application without needing to update a bunch of
references.

This package defines a set of fat interfaces (interfaces with lots of methods)
that describes things like `*os.File` and `*net/http.Request`.

## License
This package is released to the Public Domain using The Unlicense. This allows
you to either import this package, or copy the code directly, without needing to
worry about licensing or attribution. Please see
the [LICENSE](https://github.com/theckman/fatinterfaces/blob/master/LICENSE)
file for more details.

While not required by the license, the only request is that any useful changes
be contributed back so others can make use of them.

## Usage
The purpose of this package is to allow you to replace references to things like
`*os.File` with an interface. In this example it would be something like
`fatinterfaces.OsFile`.

This would allow you to either use something like `*os.File`, or to use
something like Afero.

Please see the [GoDoc page](https://godoc.org/github.com/theckman/fatinterfaces)
for an idea of other supported interfaces.
