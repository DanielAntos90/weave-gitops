## gitops check

Validates flux compatibility

```
gitops check [flags]
```

### Examples

```

# Validate flux and kubernetes compatibility
gitops check --pre

```

### Options

```
  -h, --help   help for check
  -p, --pre    perform only the pre-installation checks (default true)
```

### Options inherited from parent commands

```
  -e, --endpoint string            The Weave GitOps Enterprise HTTP API endpoint
      --insecure-skip-tls-verify   If true, the server's certificate will not be checked for validity. This will make your HTTPS connections insecure
      --namespace string           The namespace scope for this operation (default "flux-system")
  -v, --verbose                    Enable verbose output
```

### SEE ALSO

* [gitops](gitops.md)	 - Weave GitOps

###### Auto generated by spf13/cobra on 29-Apr-2022