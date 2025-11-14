#!/usr/bin/env python3

from systemd import journal
import logging

log = logging.getLogger('journald-client-python')
log.addHandler(journal.JournalHandler())

"""
Example of client application writing logs to journald.

More documentation can be found here:
https://www.freedesktop.org/software/systemd/python-systemd/journal.html
"""


def main():
    """
    main function
    :return: None
    """
    log.setLevel(logging.DEBUG)

    log.debug("Hello, Python from debug level!")
    log.info("Hello, Python from info level!")
    log.warning("Hello, Python from warning level!")
    log.error("Hello, Python from error level!")


if __name__ == "__main__":
    main()
