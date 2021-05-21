## vf-admin va add

Add a new vaccine availability

```
vf-admin va add [flags]
```

### Examples

```
# Add new vaccine availability on 2021-05-25 with 3 available input type 1 location 1651 and tags vhc
$ vf-admin va add --date "2021-05-25" --numberavailable 3 --inputtype 1 --location 1651 --tags vhc

```

### Options

```
      --date string           date for availability (YYYY-MM-DD)
  -h, --help                  help for add
      --inputtype int         input type
      --location int          id of the location
      --numberavailable int   number of vaccines available
      --numbertotal int       total number of vaccines
      --tags string           tags
      --vaccine int           id of the vaccine
```

### Options inherited from parent commands

```
      --profile string   specifies the named profile to use for this command (default "default")
```

### SEE ALSO

* [vf-admin va](vf-admin_va.md)	 - Manage vaccine availabilities

