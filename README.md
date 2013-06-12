## `go-semver` — Semantic version library for Go

Everyone loves [semantic versioning](http://semver.org).  Everyone loves Go.  Now you can use the two together.  **MAGIC.**

### Usage

Versions are parsed from strings like so:

    semver.FromString("1.2.3")
    semver.FromString("1.9.0-beta1")
    // and so on...

This gives you a `Version` struct that you can then interact with and compare:

    v1 = semver.FromString("1.2.3")
    v2 = semver.FromString("1.3.5")

    v1.LessThan(v2) // true
    v1.GreaterThan(v2) // false
    v1.Equal(v2) // false
    v1.NotEqual(v2) // true
    v1.LessThanOrEqual(v2) // true
    v2.GreaterThanOrEqual(v1) // true
    v1.String() // 1.2.3

You can also do pessimistic comparisons [like RubyGems](http://www.devalot.com/articles/2012/04/gem-versions.html):

    v1 = semver.FromString("1.2.8")
    v2 = semver.FromString("1.2.0")
    v3 = semver.FromString("1.0.0")
    v4 = semver.FromString("1.8.0")

    v1.PessimisticGreaterThan(v2) // true
    v1.PessimisticGreaterThan(v3) // true
    v1.PessimisticGreaterThan(v4) // false
