# OAuthProxy

## Description

Caching proxy service for oauth2 token authentication using password token flows

## Instructions

For the CLI manual, please run `oauthproxy --help`

### Configuration
By default, config settings are read from a YAML file at `./.oauthproxy`.
Custom locations can be provided via the command line.

### Credentials
Request credentials are expected in the standard Finbourne `secrets.json` format.

### Examples
To run an oauthproxy instance:

```
./oauthproxy serve --config /path/to/.oauthproxy
```

To make a request to a running oauthproxy instance:

```
./oauthproxy request /path/to/secrets.json
```

## Work Required:
- Expand Readme with usage description
- Remove hardcoded imports from 'github.com/nehemming/oauthproxy/'
- Add log message for returning a cached token
- Log message for 'passing on downstream' has an incomplete path
- Consider behaviour if secrets doesnt reference the proxy, as that will still work but bypass the actual service
- It needs tests
- Add it to the image used for the tests

- Refactor service.go so it's not just a blob of a million different functions
- Clear up a lot of typos and non-descriptive comments and fixup any single-letter names