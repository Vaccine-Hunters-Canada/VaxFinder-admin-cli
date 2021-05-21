## vf-admin va requirement update

Update a requirement for vaccine availability with specified ids

```
vf-admin va requirement update [flags]
```

### Examples

```
# update vaccine availability c7bc794c-9905-4588-81e6-e557e1a494c4 and requirement id a9620b24-0dc4-4c7d-8ca2-a9b06d627d82 with requirement id 10 and active false
$ vf-admin va requirement update c7bc794c-9905-4588-81e6-e557e1a494c4 a9620b24-0dc4-4c7d-8ca2-a9b06d627d82 --requirement 10 --active=false

```

### Options

```
      --active            if the requirement is active or not
  -h, --help              help for update
      --requirement int   id of requirement
```

### Options inherited from parent commands

```
      --profile string   specifies the named profile to use for this command (default "default")
```

### SEE ALSO

* [vf-admin va requirement](vf-admin_va_requirement.md)	 - Manage requirements for vaccine availability

