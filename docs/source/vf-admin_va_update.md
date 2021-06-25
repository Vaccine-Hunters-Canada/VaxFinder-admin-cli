## vf-admin va update

Update a vaccine availability with a specified id

```
vf-admin va update [flags]
```

### Examples

```
# Update vaccine availability id 7d7488e4-cc26-434d-85c4-b7df2f7e3171to be on 2021-05-25 with 3 available input type 1 location 1651 and tags vhc
$ vf-admin va update 7d7488e4-cc26-434d-85c4-b7df2f7e3171 --date "2021-05-25" --numberavailable 3 --inputtype 1 --location 1651 --tags vhc

```

### Options

```
      --date string           date for availability (YYYY-MM-DD)
  -h, --help                  help for update
      --inputtype int         input type
      --location int          id of the location
      --numberavailable int   number of vaccines available
      --numbertotal int       total number of vaccines
      --tags string           tags
      --vaccine int           id of the vaccine
```

### Options inherited from parent commands

```
      --dry-run          print the HTTP request that would be sent to the server as a cURL command
      --profile string   specifies the named profile to use for this command (default "default")
```

### SEE ALSO

* [vf-admin va](vf-admin_va.md)	 - Manage vaccine availabilities

