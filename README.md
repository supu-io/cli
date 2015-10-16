supu.io : supu
=================

This project is the entry point to supu.io, it will allow you to interact with the whole platform.

## Table of contents

- [Quick start](#quick-start)
- [Documentation](#documentation)
- [Build status](#build-status)
- [Bugs and feature requests](#bugs-and-feature-requests)
- [TODO](#todo)
- [Contributing](#contributing)
- [Versioning](#versioning)
- [Creators](#creators)
- [Copyright and license](#copyright-and-license)

## Quick Start

You need *go* installed:

```
$ git clone git@github.com:supu-io/supu.git
$ cd supu
$ go build
$ go install
```

Now you're ready to use it, check command inline help with:
```
$ supu
```
You will be able to setup your repo with command
```
cd ~/my_repo/
supu st
```
this will create all necessary status on your issue tracker (Supported github atm)

When you're inside a git repo you can:

List current repo issue:
```
supu l
```

List all org issues:
```
supu l --org
```

Show isue details
```
supu s org/repo/id
```

Move to a different status on your workflow
```
supu m org/repo/id doing
```

List all valid status
```
supu m org/repo/id
```


## Documentation

Please check the whole Project Documentation repo at:
[supu.io documentation](https://github.com/supu-io/docs)

## Build status

* Branch Master : [![Build Status Master](https://travis-ci.org/supu-io/supu.svg?branch=master)](https://travis-ci.org/supu-io/supu)

## Bugs and feature requests

Have a bug or a feature request? Please first read the
[issue guidelines](https://github.com/supu-io/supu/blob/master/CONTRIBUTING.md#using-the-issue-tracker)
and search for existing and closed issues. If your problem or idea is not
addressed yet,
[please open a new issue](https://github.com/supu-io/supu/issues/new).

## TODO

In order of precendence always work on existing
[issues](https://github.com/supu-io/supu/issues) before spending hours on
new things.

If you have an idea for the future and it is not planed on the global
[roadmap](http://github.com/supu-io/docs/roadmap.md) please check the
[TODO list of ideas] on every project repo and add your idea there to be
discussed.

If you already added a new idea to one of the existing projects, go and ping
to a developer and ask him to disscuss it. Good luck! ;)

This project TODO idea list is here: [TODO.md](todo.md).

## Contributing

Please read through our
[contributing guidelines](https://github.com/supu-io/supu/blob/master/CONTRIBUTING.md).
Included are directions for opening issues, coding standards, and notes on
development.

Moreover, if your pull request contains patches or features, you must include
relevant unit tests.

## Versioning

For transparency into our release cycle and in striving to maintain backward
compatibility, supu-io/supu is maintained under
[the Semantic Versioning guidelines](http://semver.org/). Sometimes we screw
up, but we'll adhere to those rules whenever possible.

## Creators

**Adrià Cidre**

- <https://twitter.com/adriacidre>
- <https://github.com/adriacidre>

## Copyright and License

Code and documentation copyright 2015 supu.io authors.

Code released under
[the MIT license](https://github.com/supu-io/supu/blob/master/LICENSE).
