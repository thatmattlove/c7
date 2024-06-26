<div align="center">
  <h3><code>c7</code></h3>
  <br/>
  <a href="https://github.com/thatmattlove/c7/actions/workflows/tests.yml">
    <img alt="Test Status" src="https://img.shields.io/github/actions/workflow/status/thatmattlove/c7/tests.yml?branch=main&event=push&style=for-the-badge">
  </a>
  <br/>
  Quickly Decode (or encode) a Cisco Type 7 Password at the Command Line
</div>

## Usage

### Install using [Goblin](https://goblin.run)

```shell
curl -sf http://goblin.run/github.com/thatmattlove/c7 | sh
```

### Download the latest [release](https://github.com/thatmattlove/c7/releases/latest)

There are multiple builds of the release, for different CPU architectures/platforms. Download and unpack the release for your platform:

```shell
wget <release url>
tar xvfz <release file> c7
```

### Run the binary

```console
$ ./c7 --help

Usage: c7 <command> [flags]

Quickly Encode/Decode a Cisco Type 7 Password

Flags:
  -h, --help    Show context-sensitive help.

Commands:
  encode (e) <value> [flags]
    Encode a value to a Cisco type 7 hash

  decode (d) <value> [flags]
    Decode a Cisco type 7 hash back to plain text

Run "c7 <command> --help" for more information on a command.
```

#### Encode

```console
$ c7 encode Password1234

 Encoded
 ─────────────────────────────
  0236054818110033481F5B4A51
```

#### Decode

```console
$ c7 decode 06360E325F59060B0146405858

 Decoded
 ───────────────
  Password1234
```

![GitHub](https://img.shields.io/github/license/thatmattlove/c7?style=for-the-badge&color=black)
