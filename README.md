# Logbook

A small C program to easily be able to log thoughts and stuff.

Will follow the following file structure:

```
2024/
  01/
    01.md
    02.md
  02/
    01.md
```

Will follow the following file structure, with headers being automatically generated:

```Markdown
# Thu 18/07/24

## 10:43

Some thoughts written down.

## 12:01

More thoughts.
```

## Setup

Requires the `LOGBOOK_HOME` envvar which is the folder where the log files should be created.

## Development

Compile the code:

```Bash
make bin/log
```

Run the code:

```Bash
./bin/log
```

## Future possibilites

- Add tooling for grouping together notes
- Add possiblity of storing files remotely using `scp`
