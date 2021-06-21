## vf-admin address update

Update an address with a specified id

```
vf-admin address update [flags]
```

### Examples

```
# Update the address with id 20 to have province "Ontario", postal code "K1A0A9", latitude "45.424807" and longitude "-75.699234"
$ vf-admin address update 20 --province "Ontario" --postcode "K1A0A9" --latitude "45.424807" --longitude "-75.699234"

```

### Options

```
      --city string         city of the address
  -h, --help                help for update
      --latitude float32    latitude code of the address
      --line1 string        line 1 of the address
      --line2 string        line 2 of the address
      --longitude float32   longitude code of the address
      --postcode string     postal code of the address
      --province string     province of the address
```

### Options inherited from parent commands

```
      --profile string   specifies the named profile to use for this command (default "default")
```

### SEE ALSO

* [vf-admin address](vf-admin_address.md)	 - Manage addresses

