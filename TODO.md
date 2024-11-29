# TODO

- https://github.com/gobuffalo/buffalo/issues/371
- move journal to separate subcommand
- don't assume aliases by adding bunch of extra functionality to root
- make binder open root on empty
- don't automatically deal with `.md`, require being specific
- delete this file

`log`: Maybe do nothing for now?

`log journal|j`: Adds entry to journal
`log journal|j view|v [since (1d/1m etc.)]`: View entries [since], in a concated read-only file

`log binder|b`: Opens binder root
`log binder|b <path-to-file>`: Opens specific file
