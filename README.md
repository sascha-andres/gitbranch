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
              "enabled": true
          },
          {
              "key": "develop",
              "value": "develop",
              "enabled": true
          },
          {
              "key": "feature/a",
              "value": "feature/a",
              "enabled": true
          }
      ]
    }
    
git ls-remote | awk '{print $2}' | grep refs/heads | cut -c 12-