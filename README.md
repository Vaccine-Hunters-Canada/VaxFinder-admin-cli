<div align="center">
    <h1>VaxFinder Admin CLI</h1>
</div>

<div align="center">
    <strong>A CLI application for administrators to manage vaccine availabilities and other data for the Vaccine Hunters Finder tool.</strong>
</div>

<br/>

<div align="center">
    <a href="">
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

## Available Commands

If you ever need help, run `vf-admin help`.

#### `vf-admin configure --key=<auth-key>`

Before running all other commands, you must add an authentication key that can be obtained through one of the project developers. Multiple profiles are supported so you can be authenticated using multiple keys. To create a profile with an authentication key, run `vf-admin configure --key=<auth-key> --profile=<named-profile>`.

**_Other commands in development._** :construction:

## Development

### Installation

```
go get .
```

### Adding New Commands

To create a top-level command, run the command below. Learn more [here](https://github.com/spf13/cobra/blob/master/cobra/README.md#cobra-add).
```
cobra add <command> --author="VaxFinder Project"
```

### Generating an OpenAPI Client

```
oapi-codegen -config=./.oapi-codegen.yaml https://vax-availability-api.azurewebsites.net/openapi.json
```

**After running this command, you may find that there is an issue.** This is discussed [here](https://github.com/deepmap/oapi-codegen/issues/343). In the meantime, manually fix the issue on line ~28 at `api/client.gen.go` by replacing `InputTypeEnum InputTypeEnum = 1` with `InputTypeEnum0 InputTypeEnum = 1`.


### Running locally

```
go run main.go help
```

You can also install the application locally which is how it is expected to be used in production. In order to do this, you must first set `GOPATH` and `GOBIN` appropriately. _If they are not set, add this to your `.bashrc` or `.bash_profile` etc. AND open new terminal._

```
go install
vf-admin help
```

## How to install

### Using `cURL`

TBA

### Using `go`

TBA

## How to uninstall

TBA
