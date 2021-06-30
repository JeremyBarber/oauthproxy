# OAuthProxy

## Description
Caching proxy service for oauth2 token authentication using password token flows
Original repo located at https://github.com/nehemming/oauthproxy
Reuploaded to gitlab.finbourne.com with permission

## Instructions
For the CLI manual, please run `oauthproxy --help`

### Build
The best way to build this is through the Dockerfile provided

### Configuration
By default, config settings are read from a YAML file at `./.oauthproxy`.
Custom locations can be provided via the command line.

### Credentials
Request credentials are expected in the standard Finbourne `secrets.json` format.

### Test
Run tests recursively using `go test -v ./...`

### Examples
To run an oauthproxy instance:

```
./oauthproxy serve --config /path/to/.oauthproxy
```

To make a request to a running oauthproxy instance:

```
./oauthproxy request /path/to/secrets.json
```