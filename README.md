# Simple status prometheus exporter exporter

## Usage

```sh
go build -o monitoring-exporter
```

```sh
monitoring-exporter status --name <fqdn> --metric-name <name_of_the_metric>
```

### Show help
```sh
monitoring-exporter -h
monitoring-exporter status -h
```

## Using docker
```
docker run -p 8081:8081 ewencodes/monitoring-exporter status --name <fqdn> --metric-name <name_of_the_metric>
```