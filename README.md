<div align="center">
  <h3><code>c7</code></h3>
  <br/>
  <a href="https://github.com/thatmattlove/c7/actions?query=workflow%3Agoreleaser">
    <img alt="GitHub Workflow Status" src="https://img.shields.io/github/workflow/status/thatmattlove/c7/goreleaser?color=9100fa&style=for-the-badge">
  </a>
  <br/>
  Quickly Decode (or encode) a Cisco Type 7 Password at the Command Line
</div>

## Usage

### Download the latest [release](https://github.com/thatmattlove/c7/releases/latest)

There are multiple builds of the release, for different CPU architectures/platforms:

There are multiple builds of the release, for different CPU architectures/platforms. Download and unpack the release for your platform:

```shell
wget <release url>
tar xvfz <release file> c7
```

### Run the binary

```console
$ ./c7 --help

Decode (or Encode) Cisco Type 7 Passwords

Options:

  -h, --help     show help
  -e, --encode   Encode instead of decode
```

#### Decode

```console
$ ./c7 0236054818110033481F5B4A51

Encoded: 0236054818110033481F5B4A51
Decoded: Password1234

```

#### Encode

Coming Soon!

## Creating a New Release

This project uses [GoReleaser](https://goreleaser.com/) to manage releases. After completing code changes and committing them via Git, be sure to tag the release before pushing:

```
git tag <release>
```

Once a new tag is pushed, GoReleaser will automagically create a new build & release.
