# graalvm-java-cgo-test

Build
---

GraalVM 22.1.0 Java11 MacOS and `native-image` are required.

    $ make target/call_from_go

To run

    $ export DYLD_LIBRARY_PATH=target/native
    $ ./target/call_from_go

