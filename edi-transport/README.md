# EDI Transport

A service to receive EDI files from Tivit.

This service can run in server mode, creating a SFTP Server, or in client mode
being a SFTP Client.

### About private packages
with specific configurations, a special `script/go-mod` is provided. Use
this instead of manually running `go mod download`:

```bash
$ script/go-mod download
```

In case one prefer to use `go mod` directly, the following configuration must
be present in your environment:

1. Run `git config --global url."git@github.com:".insteadOf "https://github.com/"`

### Using Docker
For automated tools such as bootstrapping, and automated testing,
Docker and Docker Compose is required. In case one does not intend to use
Docker, skip this section and refer to **Manually**.

1. Clone this repository
2. Execute `script/bootstrap` to prepare the local environment
3. Execute `script/run` to bring up a development server.

This will create all required files and containers, and configure the repository
as needed.

### Manually

> :warning: **Warning:** Before following instructions on this section, ensure
the second part of the **About private packages** section were followed.

1. Clone this repository
6. Execute `script/go-mod download` to download project dependencies, or install
them manually, provided one read and understood the **About private packages**
section of this document.

## Hacking on this Project

This project avoids installing system-wide tools, and provides facilities that
installs binaries into the local repository as needed.

### EditorConfig
This project uses EditorConfig to ensure files are consistent across platforms
and adheres a few formatting conventions. Ensure to install its extension on
your editor of choice by visiting [editorconfig.org](https://editorconfig.org).

### Mocks
[Mockery](https://github.com/vektra/mockery) is used to provide mocks to
Repositories and Services. Those files are available in the `mocks` directory,
and are created automatically by `script/generate-mocks`. Running it will
download a Mockery binary targeted to your current operating system and execute
it. No other action is needed.

### Testing
This project attempts to keep a high test coverage in order to ensure builds
are free of previsible bugs and does not contains regressions.

### Linting
[`golangci-lint`](https://github.com/golangci/golangci-lint) is used extensively
to ensure code adheres to safe practices and stays clean. `script/lint` will
execute required tasks and output any offenses.

### Git
Whenever possible, commits must be prefixed by the component affecting it.
Examples:

```
f9d6d58 Support: Add missing popd to script/run
0939793 Support: Ensure bootstrap downloads required deps
3af4b26 CI: Update liveness and readiness probes
c8e176a Support: Add script/go-mod
f41db58 Config: Respect ISTIO_PROXY_ENABLED
e5cf335 Entrypoint: Update logging facilities
```

The following conventions must be adhered:

1. Commits must start with an uppercase letter.
2. Commits must be as granular as possible, which allows them to be easily
reversed.
3. Commits must not end with a period.
3. Commits must use the imperative voice in the subject line.
4. Branches must be rebased instead of merged into themselves.

For extra information, developers are suggested to refer to [this post](https://chris.beams.io/posts/git-commit/#seven-rules)
by Chris Beans.


## Automations

> ℹ️ **Heads up!** Automations included in this repository are considered as
beta quality. If you encounter any kind of issues runing them, [open an issue](issues/new)
or contact a contributor.

### Bootstrap
`script/bootstrap` is provided to allow developers to easily start working on
this project. It is reponsible for initialising Docker containers, creating
databases, applying migrations and creating a `.env` file, required by other
automations.
This tool will refuse to overwrite contents or re-bootstrap a bootstrapped
repository. In case to override this check, provide a `--force` flag to the
tool.

### Go Module manipulation
`script/go-mod` is a small wrapper around `go mod`. It provides a configured
environment to deal with private packages used by this project.
This tool works by providing extra Git configuration by using a special
`.gitconfig` available in this repository that defines settings and includes
the user's default `.gitconfig` (by expanding `~/.gitconfig`), setting a
`GONOSUMDB` environment variables, and passing any provided arguments to the
`go mod` tool.
For instance, to download dependencies, `script/go-mod download` should suffice.

### Linter
`script/lint` is responsible for running `golint-ci` with specific flags. It is
intended to be used by developers relying on Docker or running local servers.

### Running the Application
`script/run` sources all variables defined by `.env` and executes the
application's default entrypoint, exposing an HTTP server in the port 3000.

### Testing
`script/test` requires Docker and executes all tests in a separate environment.
It is responsible for bringing servers up, applying all migrations, and execute
all tests in this environment, tearing it down after it is done. In case tear
down phase fails, `script/test stop` or `script/test down` will remove any
created container.
