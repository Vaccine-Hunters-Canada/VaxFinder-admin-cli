## vf-admin va requirement add

Add a new requirement for vaccine availability

```
vf-admin va requirement add [flags]
```

### Examples

```
# Add a new requirement 10 for a vaccine availability c7bc794c-9905-4588-81e6-e557e1a494c4
$ vf-admin va requirement add c7bc794c-9905-4588-81e6-e557e1a494c4 --requirement 10

```

### Options

```
  -h, --help              help for add
      --requirement int   id of requirement
```

### Options inherited from parent commands

```
      --profile string   specifies the named profile to use for this command (default "default")
```

### SEE ALSO

* [vf-admin va requirement](vf-admin_va_requirement.md)	 - Manage requirements for vaccine availability

