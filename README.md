# gitbranch

This is a small project extracting information from a remote git repository
about heads and tags

## What does it do?

It returns all branches for a given repository

## Call

Post to `/api/branches` with the following payload:

    {
      repository: "git@gitserver:test.git"
    }

## Result

Result will look like this:

    {
      "options": [
          {
              "key": "master",
              "value": "master",
              "enabled": true,
              "image": "..."
          },
          {
              "key": "develop",
              "value": "develop",
              "enabled": true,
              "image": "..."
          },
          {
              "key": "feature/a",
              "value": "feature/a",
              "enabled": true,
              "image": "..."
          }
      ]
    }

The service is thought to be used in conjunction with https://github.com/grundic/teamcity-web-parameters

## Credits

Images from https://www.shareicon.net/author/hub