# syslog-send

An example of how to use [Golang's `log/syslog` package](https://golang.org/pkg/log/syslog/) to send `syslog` messages to a particular syslog facility.

The repository also includes a sample Docker image that has an `rsyslog` server listening on port 514/udp forwarding any syslog messages sent to it to `stdout`.

## CLI

```
Usage:
  syslog-send [OPTIONS] Message...

Application Options:
      --address=   address of the syslog server (default: 127.0.0.1:514)
      --transport= transport to use (tcp|udp) (default: udp)
      --facility=  name of the syslog facility to send msgs to (default: local0)
      --severity=  severity of the message (default: emerg)

Help Options:
  -h, --help       Show this help message
```

## Example

First, create the `rsyslog` server:

```sh
# The Makefile under the root of this repository
# contain a `docker-compose` file that builds the
# image at `./image` and then runs the container
# with the right port mapping in place.
make run

# Send a message to the rsyslog installation using
# the default parameters.
syslog-send this is a message!

# Check the logs from the rsyslog container to
# verify that we indeed received the message.
docker logs rsyslog
2018-08-22T20:08:59-04:00 Ciros-MacBook-Pro.local custom-tag[61273]: this is a message!
```
