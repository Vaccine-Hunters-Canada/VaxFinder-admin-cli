## vf-admin location add

Add a new location

```
vf-admin location add [flags]
```

### Examples

```
# Add a new active location with name "Guelph Hospital", postal code "N1E 4J4", website URL "http://www.gghorg.ca/", phone "(519) 822-5350", notes "Please call ahead to make an appointment.", tags "Guelph, Appointment", organization id 23 and address id 352.
$ vf-admin location add --active 1 --name "Guelph Hospital" --postcode "N1E4J4" --url "http://www.gghorg.ca" --phone "(519) 822-5350" --notes "Please call ahead to make an appointment." --tags "Guelph, Appointment" --organization 23 --address 352

```

### Options

```
      --active int         is this location active? 1 or 0 (default 1)
      --address int        id of address of location
  -h, --help               help for add
      --name string        name of location
      --notes string       notes about location
      --organization int   id of organization running location
      --phone string       phone number of location
      --postcode string    postal code of location
      --tags string        search tags of location
      --url string         website URL of location
```

### Options inherited from parent commands

```
      --dry-run          print the HTTP request that would be sent to the server as a cURL command
      --profile string   specifies the named profile to use for this command (default "default")
```

### SEE ALSO

* [vf-admin location](vf-admin_location.md)	 - Manage locations

