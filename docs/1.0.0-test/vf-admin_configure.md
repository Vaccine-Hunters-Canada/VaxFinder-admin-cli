## vf-admin configure

Configure a named profile with credentials

### Synopsis

Configure the default profile or simply create new named profiles with an authentication key.

A named profile is a set of settings and credentials that you can apply to a vf-admin CLI command.


```
vf-admin configure [flags]
```

### Examples

```
z
# set up authentication key for default profile
$ vf-admin configure --key 7260841f-b47a-4b5c-9830-585af07c4405

# set up authentication key for a custom profile
$ vf-admin configure --profile alvin --key 7260841f-b47a-4b5c-9830-585af07c4405

```

### Options

```
  -h, --help         help for configure
  -k, --key string   The authentication key for future requests.
```

### Options inherited from parent commands

```
      --profile string   specifies the named profile to use for this command (default "default")
```

### SEE ALSO

* [vf-admin](vf-admin.md)	 - vf-admin is a CLI to manage vaccine availabilities and other data for the Vaccine Hunters Finder tool.

