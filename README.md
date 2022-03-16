# HTTP Mon

[![Contributor Covenant](https://img.shields.io/badge/Contributor%20Covenant-2.1-4baaaa.svg)](code_of_conduct.md)
[![Latest Release](https://img.shields.io/github/release/thanos-io/thanos.svg?style=flat-square)](https://github.com/m-yosefpor/httpmon/releases/latest)


### Github Configuration

- description
- topics (tags)
- setting > security > security alerts enabled
- setting > security > dependabot update enabled
- protect main: no one
- require resolved discussion
- required pipeline pass

- github discussions
- github pages + link in description
- release
- projects

### Gitlab Configurations

- creat API token
- CICD Variables: GITLAB_API_TOKEN, GITLAB_PRODUCTION_OKD_TOKEN, GITLAB_STAGING_OKD_TOKEN
- master: protected: no one
- pipeline must suceed
- Jira integration?

### Used

- conventional commits v1.0.0: https://github.com/conventional-commits/conventionalcommits.org
- semantic versioning v2.0.0: https://github.com/semver/semver
- CODE_OF_CONDUCT: Contributor Covenant v2.1: https://github.com/EthicalSource/contributor_covenant
- LICENSE: Apache 2.0

- golangci-lint: https://github.com/golangci/golangci-lint
- Semantic Version Util: https://github.com/caarlos0/svu
- changelog generator: https://github.com/git-chglog/git-chglog
- goreleaser: https://github.com/goreleaser/goreleaser
- distroless: https://github.com/GoogleContainerTools/distroless
- github: workflows, dependabot, issue form, pr template

TODO: description

## Development

You can use docker-compose to run a local copy of the application. Here are useful Make commands:
* `make build` builds Docker image locally.
* `make run` spins up local Docker containers so you can access via web or command line.
* `make test` runs all tests locally.
* `make rsh` spins up a temporary container based on the built image and gives you a bash inside that container.
* `make debug` spins up a temporary container based on the build image with sleep entrypoint, and gives you a bash inside that container to debug in case container stops immediately.



## Changelog

See [CHANGELOG.md](./CHANGELOG.md)

## Security

### Reporting security vulnerabilities

If you find a security vulnerability or any security related issues, please DO NOT file a public issue, instead send your report privately to TODO@gmail.com. Security reports are greatly appreciated and we will publicly thank you for it.

For more info please see [here](SECURITY.md).

## License

GNU Affero General Public License v3.0, see [LICENSE](LICENSE).
