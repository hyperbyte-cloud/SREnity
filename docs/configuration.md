## Configuration

Configuration files are defined pkl files which have 3 top level structures.

1. inputs
2. outputs
3. slos
   1. slis


### What is pkl?

PKL is a new config language open sourced by Apple. You can find it all here [Apple PKL](https://github.com/apple/pkl)

You'll probably need [PKL Referance](https://pkl-lang.org/main/current/language-reference/index.html) if you're to understand how things are defined but using the sample configuration below will probably be enough. It's not complicated for a reason.


### Inputs & Outputs
These are datasources which SREnity uses to gather data and output results to. At the moment we support InfluxDB but this has been built to support any number of datasources.

The `name` defined as a string is used to link datasources to SLO's & SLI's.

Example:

```
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
```

### SLO
An SLO defines a group of SLI's. The SLO has name, description and output attributes.

- `name`: A string name (this is sent as a label to an output)
- `description`: A string description only used in test mode.
- `output`: where to send the summary of the computed SLI's

Example:

```
slos = new {
    new {
        name = "My SLO"
        description = "My SLO description"
        output = "influx2"
    }
}
```

### SLI
An SLI defines a single indicator to be montiored.

- `name`: A string name (this is sent as a label to an output)
- `description`: A string description only used in test mode.
- `interval`: How often should this sli be evaluated
- `query`: What statement should be run when evaluating the SLI. This **must** return a single metric which is expected to be **lower** than the goal.
- `goal`: The target value the query's returned metric. **Please note that  this is assumed to be a %**

Example:

```
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
                query = "SELECT mean(field1) FROM test_measurement fill(0)"
                goal = 99.9
            }
            new {
                name = "Failing SLI"
                description = ""
                interval = 10.s
                input = "influx1"
                query = "SELECT mean(field2) FROM test_measurement fill(0)"
                goal = 99.9
            }
        }
    }
}
```