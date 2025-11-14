#!/bin/bash

echo "Hello from error level" | systemd-cat -t "journald-client-bash" -p "err"
echo "Hello from warning level" | systemd-cat -t "journald-client-bash" -p "warning"
echo "Hello from info level" | systemd-cat -t "journald-client-bash" -p "info"
echo "Hello from debug level" | systemd-cat -t "journald-client-bash" -p "debug"
