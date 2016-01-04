# Application Kill Switch

[![Build Status](https://travis-ci.org/bep/killswitch.svg)](https://travis-ci.org/bep/killswitch) [![Build status](https://ci.appveyor.com/api/projects/status/7mbikfi6sxyq7phh?svg=true)](https://ci.appveyor.com/project/bjornerik/killswitch)

This is a general-purpose and cross-platform application kill switch written in Go. The main use case is to protect applications running with the VPN when the VPN dies, and still keep general network connectivity. But it's extendable, so a time based kill switch for your develpoment IDE should also be possible: "Sorry, boss. My development environment will not start during Christmas!"

>"A kill switch, also known as an emergency stop or e-stop, is a safety mechanism used to shut off a device in an emergency situation in which it cannot be shut down in the usual manner. (Wikipedia)"

## Install

**Killswitch** is a Go application. The easiest way to intall it is via `go get`:

```bash
 go get -v github.com/bep/killswitch
```

This application has been confirmed to work fine on OS X, and Linux, and Windows. Desktop notifications implemented on OS X and Linux.

## Is it any good?

[Yes](https://news.ycombinator.com/item?id=3067434)

## Use

Wrap your sensitive application with a kill switch.

Provide a path to the program to watch and its arguments (optional), and then a conditional.

The conditional can be a built-in (see the `net` command) or a heartbeat-script
you can write yourself (see the `exec` command).


```
  -a, --args string    The program argument list
  -e, --exec string    The program to watch
      --interval int   Interval between checks in seconds (default 5)
```

## Kill Switches

### killswitch net

Will kill your executable if a given network interface vanishes.

```
killswitch net
```

```
  -n, --name string   The name of the network interface that must be present
```

To get the correct interface name to use with this command, try 

```
killswitch net list
```

## killswitch exec

Will kill your executable if your provided script exits with an error code.

The script (typically a shell script on *nix or a cmd- or bat-script on Windows) must exit with a non-0 exit-code
to signal that the application under watch should be killed.

See /testfiles for example scripts for both *nix and Windows.

If the script is not present on the PATH, the full path must be provided in name.


```
killswitch exec
```

```
  -n, --name string   The name of the script to use as heartbeat script. If not on PATH, the full path must be provided.
```