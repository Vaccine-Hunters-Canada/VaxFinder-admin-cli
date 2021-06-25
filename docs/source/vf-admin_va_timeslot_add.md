## vf-admin va timeslot add

Add a new timeslot for vaccine availability

```
vf-admin va timeslot add [flags]
```

### Examples

```
# Add a new timeslot 10 for a vaccine availability c7bc794c-9905-4588-81e6-e557e1a494c4 with time "2006-01-02T15:04:05Z"
$ vf-admin va timeslot add c7bc794c-9905-4588-81e6-e557e1a494c4 --time "2006-01-02T15:04:05Z"

```

### Options

```
  -h, --help          help for add
      --time string   time of the slot (RFC 3339)
```

### Options inherited from parent commands

```
      --dry-run          print the HTTP request that would be sent to the server as a cURL command
      --profile string   specifies the named profile to use for this command (default "default")
```

### SEE ALSO

* [vf-admin va timeslot](vf-admin_va_timeslot.md)	 - Manage timeslots for vaccine availability

