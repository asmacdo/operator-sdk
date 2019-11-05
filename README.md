<img src="doc/images/operator_logo_sdk_color.svg" height="125px"></img>

[![Build Status](https://travis-ci.org/operator-framework/operator-sdk.svg?branch=master)](https://travis-ci.org/operator-framework/operator-sdk)
[![License](http://img.shields.io/:license-apache-blue.svg)](http://www.apache.org/licenses/LICENSE-2.0.html)
[![Go Report Card](https://goreportcard.com/badge/github.com/operator-framework/operator-sdk)](https://goreportcard.com/report/github.com/operator-framework/operator-sdk)

## Overview

This project is a component of the [Operator Framework][of-home], an
open source toolkit to manage Kubernetes native applications, called
Operators, in an effective, automated, and scalable way. Read more in
the [introduction blog post][of-blog].

[Operators][operator_link] make it easy to manage complex stateful
applications on top of Kubernetes. However writing an operator today can
be difficult because of challenges such as using low level APIs, writing
boilerplate, and a lack of modularity which leads to duplication.

The Operator SDK is a framework that uses the
[controller-runtime][controller_runtime] library to make writing
operators easier by providing:
- High level APIs and abstractions to write the operational logic more
- intuitively
- Tools for scaffolding and code generation to bootstrap a new project fast
- Extensions to cover common operator use cases

[of-home]: https://github.com/operator-framework
[of-blog]: https://coreos.com/blog/introducing-operator-framework
[operator_link]: https://coreos.com/operators/
[controller_runtime]: https://github.com/kubernetes-sigs/controller-runtime
