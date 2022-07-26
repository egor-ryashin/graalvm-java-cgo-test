# graalvm-java-cgo-test

This demonstrates the use of GraalVM native shared library being used in Go executable. Particularly it shows how to call a Java function from Go and handle a returned string.

Build
---

GraalVM 22.1.0 Java11 MacOS and `native-image` are required.

    $ make target/call_from_go

To run
    
    $ ./target/call_from_go

