rounding
========

[![Build Status](https://travis-ci.org/brandscreen/rounding.png)](https://travis-ci.org/brandscreen/rounding) [![Coverage Status](https://coveralls.io/repos/brandscreen/rounding/badge.png?branch=HEAD)](https://coveralls.io/r/brandscreen/rounding?branch=HEAD)

Rounding algorithms in [Go](http://golang.org) language.

See also:

* [http://en.wikipedia.org/wiki/Rounding](http://en.wikipedia.org/wiki/Rounding)
* [http://www.mathsisfun.com/numbers/rounding-methods.html](http://www.mathsisfun.com/numbers/rounding-methods.html)

# Requirements

* Go 1.1 or higher

# Installation

```bash
$ go get github.com/brandscreen/rounding
```

# Usage

## Import Rounding

Ensure you have imported the ```rounding``` package at the top of your source file.

```golang
import "github.com/brandscreen/rounding"
```

## Rounder objects

The first way to use ```rounding``` is to create instances of the ```Rounder``` objects that implement the algorithms.
This allows abstracting away the actual algorithm and settings used.

### Create a rounding object

```golang
// Assuming "math/rand" has been imported.
var rounder Rounder
var midpoint float64 = 0.5
rounder = rounding.NewDownRounder()
rounder = rounding.NewUpRounder()
rounder = rounding.NewHalfDownRounder()
rounder = rounding.NewHalfUpRounder()
rounder = rounding.NewHalfEvenRounder()
rounder = rounding.NewSymmetricDownRounder()
rounder = rounding.NewSymmetricUpRounder()
rounder = rounding.NewSymmetricHalfDownRounder()
rounder = rounding.NewSymmetricHalfUpRounder()
rounder = rounding.NewAlternateRounder()
rounder = rounding.NewHalfAlternateRounder()
rounder = rounding.NewSymmetricAlternateRounder()
rounder = rounding.NewSymmetricHalfAlternateRounder()
rounder = rounding.NewRandomRounder(midpoint, rand.Float64)
rounder = rounding.NewHalfRandomRounder(midpoint, rand.Float64)
rounder = rounding.NewSymmetricRandomRounder(midpoint, rand.Float64)
rounder = rounding.NewSymmetricHalfRandomRounder(midpoint, rand.Float64)
```

### Use the rounding algorithm

```golang
var roundedValue float64
value := 0.8
roundedValue = rounder.Round(value)
```

## Rounding functions

The first way to use ```rounding``` is to use the actual rounding functions that implement the algorithms directly.

```golang
var roundedValue float64
value := 0.8
roundedValue = rounding.RoundDown(value)
roundedValue = rounding.RoundUp(value)
roundedValue = rounding.RoundHalfUp(value)
roundedValue = rounding.RoundHalfDown(value)
roundedValue = rounding.RoundHalfEven(value)
roundedValue = rounding.RoundSymmetricDown(value)
roundedValue = rounding.RoundSymmetricUp(value)
roundedValue = rounding.RoundSymmetricHalfDown(value)
roundedValue = rounding.RoundSymmetricHalfUp(value)
```

# Contributing

Please see [CONTRIBUTING.md](https://github.com/brandscreen/rounding/blob/master/CONTRIBUTING.md).  If you have a bugfix or new feature that you would like to contribute, please find or open an issue about it first.

# License

Licensed under the MIT License.
