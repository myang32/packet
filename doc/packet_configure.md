## packet configure

Set default configs for the packet cli.

### Synopsis


Set default configs for the packet cli.

The following configurations are supported:
- default API key
  This default key will be used if "--key" flag is missing in command.
- default project ID
  This ID will be used if "--project-id" flag is missing in command.

```
packet configure
```

### Options inherited from parent commands

```
  -k, --key string       Specify the api key
  -p, --profile string   Specify profile name (default "default")
  -v, --version          Show version and exit
```

### SEE ALSO
* [packet](packet.md)	 - A unified tool to manage your packet services

###### Auto generated by spf13/cobra on 18-Jul-2016
