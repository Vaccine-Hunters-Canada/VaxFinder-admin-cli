<div align="center">
    <h1>VaxFinder Admin CLI</h1>
</div>

<div align="center">
    <strong>A CLI for administrators to manage vaccine availabilities and other data for the Vaccine Hunters Finder tool.</strong>
</div>

<br/>

<div align="center">
    <a href="https://golang.org/dl/">
        <img src="https://img.shields.io/github/go-mod/go-version/Vaccine-Hunters-Canada/VaxFinder-admin-cli" alt="Go v1.16">
    </a>
    <a href="https://goreportcard.com/report/github.com/Vaccine-Hunters-Canada/VaxFinder-admin-cli">
        <img src="https://goreportcard.com/badge/github.com/Vaccine-Hunters-Canada/VaxFinder-admin-cli" alt="Go Report Card">
    </a>
    <a href="https://discord.com/channels/822486436837326908/832366009091358731">
        <img src="https://img.shields.io/badge/-%23vax--ui--backend-7389D8?logo=discord&logoColor=ffffff&labelColor=6A7EC2" alt="Discord: #vax-ui-backend" />
    </a>
</div>

## Installation

Click [here](#how-to-install) for installation instructions.

## Commands

If you ever need help, run `vf-admin help`.

#### `vf-admin configure --key=<auth-key>`

Before running all other commands, you must add an authentication key that can be obtained through one of the developers working on the VaxFinder backend. Multiple profiles are supported so you can be authenticated using multiple keys. To create a profile with an authentication key, run `vf-admin configure --key=<auth-key> --profile=<named-profile>`.

**Go to [the `docs` folder](/docs) to view auto-generated documentation on the CLI.**

**_Note that there is no stable release for this tool yet._** :construction:

## Development

### Installation

```bash
go get .
```

### Pre-commit Hooks

Pre-commit hooks helps identify simple issues in code before it's committed into Git.

#### Install pre-commit

Follow installation instructions for pre-commit [here](https://pre-commit.com/#installation).

#### Install the git hook scripts

```bash
pre-commit install
```

#### Temporarily Disabling hooks

It's possible to disable hooks temporarily, but it isn't recommended.

```bash
$ SKIP=go-fmt,go-vet,go-lint git commit -m <message>
```

### Running locally

```bash
go run main.go help
```

You can also install the application locally which is how it is expected to be used in production. In order to do this, you must first set `GOPATH` and `GOBIN` appropriately. _If they are not set, add this to your `.bashrc` or `.bash_profile` etc. AND open new terminal._

```bash
make install
vf-admin help
```

### Generating an OpenAPI Client

```bash
make api-codegen
```

**After running this command, you may find that there is an issue.** This is discussed [here](https://github.com/deepmap/oapi-codegen/issues/343). In the meantime, manually fix the issue on line ~28 at `internal/api/client.gen.go` by replacing `InputTypeEnum InputTypeEnum = 1` with `InputTypeEnum0 InputTypeEnum = 1`.

### Generating Docs Automatically

```bash
make docs
```

## How to install

### Using `cURL`

TBA

### Using `go`

TBA

## How to uninstall

TBA
