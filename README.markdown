# What's this?

`log2curl` converts json-encoded go (incoming) `http.Request` objects to curl
commands. Right now, for requests with a body, only values from `Form` are
supported.

Input is read from stdin. Each request is expected to be given as a single line
of json.

# Usage example

```
echo '{"Path":"/bar", "Method":"POST", "Header":{"X-Whatever":["Nevermind"]}, "Form":{"oh":["well"]}}' | go run main.go http://foo.com
curl -X 'POST' -d 'oh=well' -H 'X-Whatever: Nevermind' 'http://foo.com/bar'
```
