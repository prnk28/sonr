
- Cosmos SDK: v0.46.3
- Ignite CLI: v0.25.1
- Golang: 1.18.8 darwin/arm64

## Get started

[![CodeFactor](https://www.codefactor.io/repository/github/sonr-io/sonr/badge)](https://www.codefactor.io/repository/github/sonr-io/sonr)
  [![Status](https://img.shields.io/badge/status-active-success.svg)](https://sonr.io)
  [![Go Reference](https://pkg.go.dev/badge/github.com/sonr-io/sonr.svg)](https://pkg.go.dev/github.com/sonr-io/sonr)
  [![Go Report Card](https://goreportcard.com/badge/github.com/sonr-io/sonr)](https://goreportcard.com/report/github.com/sonr-io/sonr)
  [![GitHub Issues](https://img.shields.io/github/issues/sonr-io/sonr.svg)](https://github.com/sonr-io/sonr/issues)
  [![GitHub Pull Requests](https://img.shields.io/github/issues-pr/sonr-io/sonr.svg)](https://github.com/sonr-io/sonr/pulls)
  [![License](https://img.shields.io/badge/license-GPLv3-blue.svg)](/LICENSE)

</div>

---

<p align="center"> Build <strong>privacy-preserving</strong>, <strong>user-centric applications</strong>, on a robust, rapid-scaling platform designed for interoperability, and total digital autonomy.
    <br>
</p>

### Prerequisites
- Cosmos SDK: v0.46.3
- Ignite CLI: v0.25.1
- Golang: 1.18.8 darwin/arm64

**Getting Started**

```sh
ignite chain serve
```

`serve` command installs dependencies, builds, initializes, and starts your blockchain in development.

### Configure

Your blockchain in development can be configured with `config.yml`. To learn more, see the [Ignite CLI docs](https://docs.ignite.com).

### Web Frontend

Ignite CLI has scaffolded a Vue.js-based web app in the `client/vue` directory. Run the following commands to install dependencies and start the app:

```sh
cd client/vue
npm install
npm run serve
```

The frontend app is built using the `@starport/vue` and `@starport/vuex` packages. For details, see the [monorepo for Ignite front-end development](https://github.com/ignite/web).

## Release
To release a new version of your blockchain, create and push a new tag with `v` prefix. A new draft release with the configured targets will be created.

```sh
git tag v0.1
git push origin v0.1
```

After a draft release is created, make your final changes from the release page and publish it.

### Install
To install the latest version of your blockchain node's binary, execute the following command on your machine:

```sh
curl https://get.ignite.com/sonr-io/sonr@latest! | sudo bash
```
Learn more about [the install process](https://github.com/allinbits/starport-installer).

## Learn more

- [Homepage](https://snr.la/h)
- [Blog](https://snr.la/blg)
- [Sonr SDK docs](https://snr.la/docs)
- [Developer Chat](https://snr.la/dcrd)
