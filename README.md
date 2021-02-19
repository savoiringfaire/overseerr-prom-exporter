# Overseerr Prometheus Exporter

Prometheus exporter for basic Overseerr statistics. Currently provides Overseer total requests and pending requests.

## Getting Started

These instructions will cover usage information and for the docker container 

### Prerequisities


In order to run this container you'll need docker installed.

* [Windows](https://docs.docker.com/windows/started)
* [OS X](https://docs.docker.com/mac/started/)
* [Linux](https://docs.docker.com/linux/started/)

### Usage

```shell
docker run -p 2112:2112 savoiringfaire/overseerr-prom-exporter:1.0.0
```

#### Environment Variables

* `OVERSEERR_API_BASE_URL` - The base path for the overseerr API (e.g. `http://localhost:5055/api/v1`)
* `OVERSEERR_API_KEY` - The Overseerr API key to use to connect to Overseerr.

#### Connecting  

```yaml
  - job_name: 'overseerr'
    scrape_interval: 60s
    static_configs:
      - targets: ['10.20.0.19:2112']
```

## Find Us

* [GitHub](https://github.com/savoiringfaire/overseerr-prom-exporter)
* [Docker Hub](https://hub.docker.com/r/savoiringfaire/overseerr-prom-exporter)

## Contributing

Please read [CONTRIBUTING.md](CONTRIBUTING.md) for details on our code of conduct, and the process for submitting pull requests to us.

## Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the 
[tags on this repository](https://github.com/savoiringfaire/overseerr-prom-exporter/tags). 

## Authors

* **Savoir Faire** - *Initial work* - [SavoirFaire](https://github.com/savoiringfaire)

See also the list of [contributors](https://github.com/savoiringfaire/overseerr-prom-exporter) who 
participated in this project.

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details.
