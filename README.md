# missinglicense

**missinglicense** is a simple tool that lists all of your public GitHub repositories that are missing a LICENSE file.

## Installation

Installation is simple and no different to any Go tool. The only requirement is a working [Go](https://golang.org/) install.

```
go get go.tmthrgd.dev/missinglicense
```

## Usage

Usage is simple with the `missinglicense` command requiring only the `GITHUB_TOKEN` environment variable to be set.

```
missinglicense
```

You need to set the `GITHUB_TOKEN` environment variable to a valid GitHub personal access token with `public_repo` scope. See [Authenticating with GraphQL](https://docs.github.com/en/free-pro-team@latest/graphql/guides/forming-calls-with-graphql#authenticating-with-graphql) and [Creating a personal access token](https://docs.github.com/en/free-pro-team@latest/github/authenticating-to-github/creating-a-personal-access-token) for instructions on creating a token.

**missinglicense** does not check that the LICENSE file is valid or recognised, merely that it exists. It handles both LICENSE and LICENSE.md files.

## License

[BSD 3-Clause License](LICENSE)
