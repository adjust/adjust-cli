## Adjust Command Line Utility

This is an UNIX command line utility that streamlines your access to the Adjust
APIs. Use this tool for a quick access to all Adjust apps and KPIs data you run.

### Installation

[DISCLAIMER] This software is still in pre-release mode, and complete
installation guides are still being developed.

### Commands

The adjust tool has three commands: `adjust config`, `adjust deliverables` and `adjust cohorts`. You can see details for
each of these by typing for example: `adjust deliverables --help`. There's also general help: `adjust --help`.

Let's run through each of the commands with a few examples.

#### Config

To access the adjust APIs you'll need to authorize with your user token. You can obtain it from the accounts settings on
your adjust dashboard.

```
adjust config --user-token P1jTcnGvDbs-Nvz6ypQ2
```

This will create a file called `~/.adjustrc` containing all configuration settings for adjust. Further options to
`adjust config` include:

```
adjust config --app-token my-app-token
```

Some commands support multiple app-tokens (e.g. `adjust deliverables`). For these, you can configure multiple app-tokens
too like this:

```
adjust config --app-tokens my-app-token1,my-app-token2
```

#### Deliverables

Accessing your tracking data is done via the KPI Service API. While the adjust tool provides extensive inline help on
the `deliverables` and `cohorts` commands, you're still encouraged to keep track of the [KPI Service
docs](https://docs.adjust.com/en/kpi-service/).

```
adjust deliverables -k installs
```

This will print a nicely formatted table on your terminal with the installs for the configured app this month, by
default grouped by `networks`. If you prefer not working with a configured app, but passing an app token instead, you can:

```
adjust deliverables -k installs -a my-app-token
```

Passing more than one app token as well as changing the grouping is also supported.

```
adjust deliverables -k installs -a my-app-token1,my-app-token2 -g apps,networks
```

You can also specify start and end dates for your requests, apply filtering by country, OS name, etc.

```
adjust deliverables -a app-token-1,app-token-2 -k installs,sessions -s 2015-01-01 -e 2015-10-10 -c de,gb -g app,network
```

Check out `adjust deliverables --help` for more examples and usage details.

#### Cohorts

For cohort or event queries, run:

```
adjust cohorts -a ckrgvqfu1nx3 -k retained_users
```

Again `adjust cohorts --help` will give you more details.

### Output Format

By default the adjust tool prints results in a table-formatted output to the terminal. You can change that to printing
CSV using the `--csv` flag on every `deliverables` or `cohorts` request. This might be useful if you wish, for example,
to pipe the output for further processing.

Furthermore, you can directly save the CSV output to a file on your computer by passing `--file path/to/file`.

There's also a `--verbose` flag, which prints more details along the execution and the `--url-only` which only prints
the API URL for the request, but doesn't execute it.

### TODOs

  - Add and distribute tab completion.
  - Provide Windows support.
  - Implement basic 'smart' default groupings - e.g. group by network for single app-token requests and by
    `apps,networks` for multiple apps.
