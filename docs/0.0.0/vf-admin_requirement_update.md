## vf-admin requirement update

Update a requirement with a specified id

```
vf-admin requirement update [flags]
```

### Examples

```
# Update the requirement with id 8 to have name "High-Risk" and description "Highest- and High-Risk Health Conditions.".
$ vf-admin requirement update 8 --name "High-Risk" --description "Highest- and High-Risk Health Conditions."

```

### Options

```
      --description string   description of requirement
  -h, --help                 help for update
      --name string          name of requirement
```

### Options inherited from parent commands

```
      --profile string   specifies the named profile to use for this command (default "default")
```

### SEE ALSO

* [vf-admin requirement](vf-admin_requirement.md)	 - Manage requirements

