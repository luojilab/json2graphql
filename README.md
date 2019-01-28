A script to generate graphql schema from json.

# Usage
```bash
go run main.go -h
```


    NAME:
       inspect - generate a graphql schema based on json

    USAGE:
       main [global options] command [command options] [arguments...]

    DESCRIPTION:
       inspect json and generate draft schema.graphql

    COMMANDS:
         inspect  generate a graphql schema based on json
         help, h  Shows a list of commands or help for one command

    GLOBAL OPTIONS:
       --verbose, -v             show logs
       --input value, -i value   the json filename
       --output value, -o value  the target filename to store generated schema
       --help, -h                show help

# Example

```bash
go run main.go -i example.json
```

# TODO

- [ ] build it as a web service that render schema on the fly like [json.cn](https://json.cn)
- [ ] support to read from multi json files.
- [ ] get input from http request rather than local file.
- [ ] integrate with graphql server frameworks like [gqlgen](https://github.com/99designs/gqlgen) and auto generate resolver