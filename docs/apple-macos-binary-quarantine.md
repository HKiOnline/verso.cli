# Apple MacOS Binary Quarantine

By default, binaries downloaded from the Internet are given an extended attribute to mark them as quarantined. This affects Verso CLI as well as long as its not (at least yet) code signed. If you feel comfortable trusting Verso, you may de-quarantine the downloaded binary following the instructions below.

1. Open your terminal emulator
2. Use xattr utility to de-quarantine the binary

```bash
xattr -dr com.apple.quarantine <path to file>/verso
```

If you are not the owner of the binary file you may need to use sudo to run xattr. However I'd recommend changing the binary ownership first.
