## Adapter — Make incompatible interfaces work together

Goal: convert the interface of a class into another interface clients expect.

This example adapts a `Windows` computer (expects a USB connector) to be used where a `Mac` (expects a Lightning connector) is required. The `WindowsAdapter` makes the `Windows` machine look like a `Computer` with Lightning support for the `Client`.

Run:
```bash
go run .
```

Why this design:
- **Non-invasive**: no changes to `Windows`; the adapter handles translation
- **Single-responsibility**: adapter only maps calls/protocols
- **Practical**: mirrors real-world port/protocol adapters


