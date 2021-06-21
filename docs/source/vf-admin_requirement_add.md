## vf-admin requirement add

Add a new requirement

```
vf-admin requirement add [flags]
```

### Examples

```
# HTTPOperation a new requirement with name "18+" and description "Any individual older than 18 years of age.".
$ vf-admin requirement add --name "18+" --description "Any individual older than 18 years of age."

```

### Options

```
      --description string   description of requirement
  -h, --help                 help for add
      --name string          name of requirement
```

### Options inherited from parent commands

```
      --profile string   specifies the named profile to use for this command (default "default")
```

### SEE ALSO

* [vf-admin requirement](vf-admin_requirement.md)	 - Manage requirements

