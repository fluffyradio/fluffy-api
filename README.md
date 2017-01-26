# Fluffy Radio Api [![Twitter Follow](https://img.shields.io/twitter/follow/fluffy_radio.svg?style=social&label=Follow)]()
A wrapper API for the Spacial Audio Cloud API

| Linux Build | Windows Build | License |
| ----------- | ------------- | ------- |
| [![Build Status](https://img.shields.io/travis/fluffyradio/fluffy-api.svg)](https://travis-ci.org/fluffyradio/fluffy-api) | [![Build status](https://img.shields.io/appveyor/ci/brentpabst/fluffy-api.svg)](https://ci.appveyor.com/project/brentpabst/fluffy-api) | [![license](https://img.shields.io/github/license/fluffyradio/fluffy-api.svg)]() |

This API is designed to wrap the somewhat ugly and non-standard Spacial Audio Cloud API.  The wrapper is built in Go and proxies requests to Spacial Audio.  The API is publised to Pivotal Web Services.

## Getting Started
These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites
There are a few things you need to do before you begin working on the project.
1. Install Go, we've used `Go 1.7.4`
1. Install [Glide](https://glide.sh) for dependency management
1. Have access to a Spacial Audio Cloud Station ID and Token

### Installing
A few things have to be done to run the API.  Ensure you have `git clone`'d the site first.
1. Run `glide install`
1. Run `go build`
1. Run `./fluffy-api --spacial-id=SPACIALSTATIONID --spacial-token=SPACIALSTATIONTOKEN`

## Running Tests
1. Run `go test`

## Deployment
This code repository is connected and published to Pivotal Web Services, a Cloud Foundry distribution.  The Travis CI build is the source for the deployment.

## Contributing
Please read [CONTRIBUTING.md](CONTRIBUTING.md) for details on our code of conduct, and the process for submitting pull requests.

## Versioning
We use [SemVer](http://semver.org/) for versioning.

## Authors
* **Brent Pabst** - Initial work* - [brentpabst](https://github.com/brentpabst)

See also the list of [contributors](https://github.com/203solutions/203solutions.github.io/contributors) who participated in this project.

## License
This project is licensed under the MIT license - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

* [PurpleBooth](https://github.com/PurpleBooth) for sample README and CONTRIBUTING file Gists