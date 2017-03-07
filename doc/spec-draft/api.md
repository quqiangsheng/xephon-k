# API Specification Draft

Only support restful API for early version, maybe add a line protocol

## Info

For getting basic info of current tsdb, so the client library can make some choice

- `/info`

````json
{
    "version": "0.1.0",
    "features" : {
        "tag" : false,
        "shard" : false
    },
    "backend": "cassandra"
}
````

- `/info/version`

````json
{
    "version": "0.1.0"
}
````

- `/info/<backend-name>`

````json
{
    "name": "cassandra",
    "version" : "3.0.3",
    "schema" : "bucket"
}
````

## Insert

- `/w` (I know this is not restful, but why bother)
- [ ] KairosDB use `[1359788400000, 10]` to avoid the overhead of key in data points, don't know if golang can
handle this properly, may need to add custom unmarshal handler
- [ ] KairosDB seems to be allowing mixing int and float number, bec[1359788400000, 10]ause it store value using blob
  - [ ] we may need to split int and float into different table, I don't know if cassandra allow mixed float and int,
  and there seems to be little reason for changing from int to float for one series from time to time.

````json
[
    {
        "name": "cpu.load",
        "tags": {
            "region": "us-east",
            "os": "ubuntu"
        },
        "int_points": [
            [1359788400000, 10],
            [1359788300000, 13]
        ]
    },
    {
        "name": "mem.usage",
        "tags": {
            "region": "us-east",
            "os": "ubuntu"
        },
        "float_points": [
            [1359788400000, 10.3],
            [1359788300000, 13.2]
        ]
    }
]
````

## Read

- `/q` (I know this is not restful too)
- [ ] KairosDB support query multiple series at same time, this could be useful,
i.e. I want to query mem.total and mem.usage at same time
- [ ] do we need to support limit when start and end time is provided
- [ ] how to handle when one metric name actually have multiple series due to tags

````json
{
    "use_cache": false,
    "start_time": 1357023600000,
    "end_time": 1357077800000,
    "metrics" : [
        {
            "name": "cpu.idle"
        },
        {
            "name": "cpu.usage"
        }
    ]
}
````

response

- include a copy of the query
- [ ] same json problem as above

````json
{
    "query" : {
        "use_cache": false,
        "start_time": 1357023600000,
        "end_time": 1357077800000,
        "metrics" : [
            {
                "name": "cpu.idle"
            },
            {
                "name": "cpu.usage"
            }
        ]
    },
    "metrics" : [
        {
            "name": "cpu.idle",
            "values": [
                [1364968800000, 10],
                [1366351200000, 20]
            ]
        }
    ]
}
````