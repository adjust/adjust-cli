## TODOs

  - Should have a --csv argument to save report to a file;
  - append with human readable names to the Requests;
  - make it possible through the configuration to use different instance of the
    API - e.g. api-qa-2.adjust.com
  - implement at least `app_token` in `~/.adjustrc` and maybe both
  - consider default action and aliases to the commands and default values for
    the params

## Adjust Command Line Utility

This is an UNIX command line utility that streamlines your access to the Adjust
APIs. Use this tool for a quick access to all Adjust apps and KPIs data you run.

### Installation

   ~~brew tap adjust/homebrew~~
   ~~brew install adjust/homebrew/adjust~~

   go get github.com/adjust/adjust-cli-client

### Login

Like in using the dashboard, before you access data from Adjust, you need to
login. This is done using your user token, which you can get from the dashboard.

   adjust login PsjTcnGvDbs-Nvz7ypQ4

This will save the token to an `~/.adjustrc` until you logout.

### Logout

Logging out will delete that file:

   adjust logout

### KPI Service API

The KPI Service is currently the only supported API by the client.

Very often you're only interested for a quick check on an app's KPI. Get that quickly:

   adjust ckrgvqfu1nx3 -k installs

This will deliver the installs for this month for that app broken down by tracker.

You can also set the start and end dates for the period you're interested in:

   adjust ckrgvqfu1nx3 -k installs,sessions -s '2015-01-01' -e '2015-10-10'

For more custom requests, you could add custom grouping and filtering by country, OS name, etc.

   adjust ckrgvqfu1nx3 -k installs,sessions -s '2015-01-01' -e '2015-10-10' -c 'de,gb' -g 'trackers'

For cohort or event queries, run:

   adjust cohorts ckrgvqfu1nx3 -k retained_users

Check out `adjust help` for more.
