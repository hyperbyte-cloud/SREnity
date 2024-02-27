# SREnity

`SREnity` is an open-source project written in Golang that enables users to easily configure and manage Service Level Objectives (SLOs) and Service Level Indicators (SLIs). This tool allows you to set up data sources, compute metrics, and submit the results to an output for further analysis and monitoring.

## Features

- **Configuration-driven:** Easily define SLOs and SLIs using a flexible configuration system.
- **Data Sources:** Connect to various data sources to collect the necessary metrics.
- **Computation Engine:** Perform calculations on collected data to derive SLOs and SLIs.
- **Output Integration:** Submit computed metrics to an output for monitoring and analysis.

## Getting Started

SREnity is a simple tool used to monitor and report the state of a SLO using SLI's. SREnity uses a simple yaml configuration file to define what it should monitor.

    1. Define a config file
    2. Run SREnity `srenity -c <config_file> start`

### Installation

To install SREnity simply download it from the releases and install it to: `/usr/local/sbin` this will now be in your path and you'll be able to use it.

### Usage

CLI Usage:

```
NAME:
   srenity - SREniy the SLO Tool

USAGE:
   srenity [global options] command [command options] 

COMMANDS:
   start, r     Run the SLO tool
   test, t      Test the configuration
   validate, v  Validate the configuration
   help, h      Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --config FILE, -c FILE  Load configuration from FILE (default: "config.pkl")
   --help, -h              show help
```

1. Define a configuration file (`srenity-config.pkl`):

```pkl
inputs = new {
    new {
        name = "influx1"
        type = "influxdb_v1"
        config = new {
            host = "http://localhost:8086"
            database = "testdb"
        }
    }
}
outputs = new {
    new {
        name = "influx2"
        type = "influxdb_v1"
        config = new {
            host = "http://localhost:8086"
            database = "slo"
        }
    }
}
slos = new {
    new {
        name = "My SLO"
        description = "My SLO description"
        output = "influx2"
        slis {
            new {
                name = "Successful SLI"
                description = ""
                interval = 10.s
                input = "influx1"
                query = "SELECT mean(field1) FROM test_measurement WHERE time > now() - 5m GROUP BY time(1m) fill(0)"
                goal = 99.9
            }
            new {
                name = "Failing SLI"
                description = ""
                interval = 10.s
                input = "influx1"
                query = "SELECT mean(field2) FROM test_measurement WHERE time > now() - 5m GROUP BY time(1m) fill(0)"
                goal = 99.9
            }
        }
    }
}

```

2. Run SREnity

```bash
srenity start -c srenity-config.pkl
```

This command will start monitoring SLOs based on the provided configuration.

## Configuration

SREnity uses a [PKL](https://github.com/apple/pkl) configuration file to define SLOs, SLIs, data sources, and more. See [Configuration Guide](docs/configuration.md) for detailed information.

## Contributing

We welcome contributions from the Golang community! If you'd like to contribute, please follow our [Contribution Guidelines](CONTRIBUTING.md).

## License

This project is licensed under the [MIT License](LICENSE).

## Acknowledgments



