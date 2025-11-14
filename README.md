Logging using journald from CLI applications
============================================

When you run your application as a systemd service, then you needn't
use any special logging library, because systemd will take care of
it. When a process started systemd, then everything written to stdout
and stderr will be logged to journald.

When you run an application from the terminal, then you need to use a special
library to write logs to journald. This repository contains simple examples
of how to do it for Go applications, Python applications and Bash scripts.

Example of a Go application
---------------------------

To build the example of Go application run:

```console
$ go build
```

To run testing Go application run:

```console
$ ./journald-client
```

Example of a Python application
-------------------------------

To run the example of a Python application, run:

```console
$ python ./journald-client.py
```

Example of a Bash script
------------------------

To run the example of a Python application, run:

```console
$ ./journald-client.sh
```

Showing log messages
--------------------

You can see the logs in journald using the following command:

```console
$ journalctl --user -f -n 0
```

If you want to see only messages from the Go application, then use the following
command:

```console
$ journalctl --user -f -n 0 -t journald-client
```