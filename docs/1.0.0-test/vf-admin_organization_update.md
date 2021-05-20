## vf-admin organization update

Update an organization with a specified id

```
vf-admin organization update [flags]
```

### Examples

```
# Update the organization with id 20 to have short name "WHO", full name "World Health Organization", description "The World Health Organization is a specialized agency of the United Nations responsible for international public health." and url "https://www.who.int/"
$ vf-admin organization update 20 --shortName WHO --fullName "World Health Organization" --description "The World Health Organization is a specialized agency of the United Nations responsible for international public health." --url "https://www.who.int/"

```

### Options

```
      --description string   description of organization
      --fullName string      full name of the organization
  -h, --help                 help for update
      --shortName string     short name of the organization
      --url string           url of organization
```

### Options inherited from parent commands

```
      --profile string   specifies the named profile to use for this command (default "default")
```

### SEE ALSO

* [vf-admin organization](vf-admin_organization.md)	 - Manage organizations

