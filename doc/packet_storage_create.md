## packet storage create

Create a volume

### Synopsis


Create a volume

```
packet storage create
```

### Options

```
      --count int           Snapshots count (default 4)
      --desc string         Description
      --facility string     ewr1 || sjc1 || ams1 (default "ewr1")
      --frequency string    Snapshot frequency (default "15min")
      --plan string         storage_1 || storage_2 (default "storage_1")
      --project-id string   Project ID
      --size int            Volume size (default 120)
```

### Options inherited from parent commands

```
  -k, --key string       Specify the api key
  -p, --profile string   Specify profile name (default "default")
  -v, --version          Show version and exit
```

### SEE ALSO
* [packet storage](packet_storage.md)	 - Manage your storages

###### Auto generated by spf13/cobra on 18-Jul-2016
