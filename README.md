# ws: word searcher
`ws` command provides a list of Enlish words matches given conditions.

## Usage: see `ws help`
```
NAME:
   ws - search words with conditions

USAGE:
   ws [global options] command [command options] [arguments...]

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --match value, -m value    filter by characters with placeholders (placeholder: '_')
   --include value, -i value  filter by included characters
   --exclude value, -e value  filter by non-used characters
   --length -m, -l -m         filter by word length (ignored when -m used) (default: 0)
   --exclude-compounds        exclude compounds (default: false)
   --verbose                  displays some logs (default: false)
   --help, -h                 show help (default: false)
```

## Requirement
- **Go 1.17** or newer is required to build the project.
- **SQLite3** is required to fetch candidates from DB.

## Acknowledgement
[Dictionary](https://github.com/kujirahand/EJDict) is provided by [@kujirahand](https://github.com/kujirahand)
