# Test

- go testing framework: https://github.com/onsi/ginkgo
- go matcher/assertion library: https://onsi.github.io/gomega/

## Commands

### Bootstrapping a Ginkgo suite for a package:
``` bash
cd path/to/packagename
ginko bootstrap
```

### Adding specs to a suite:
```bash
cd path/to/packagename
ginko generate packagename
```

### Running ginkgo in watch mode:
```bash
ginkgo watch -p
```
