## Adapter — Make incompatible interfaces work together

Goal: convert the interface of a class into another interface clients expect.

This example adapts a `Windows` computer (expects a USB connector) to be used where a `Mac` (expects a Lightning connector) is required. The `WindowsAdapter` makes the `Windows` machine look like a `Computer` with Lightning support for the `Client`.

### Structure
- `interfaces/` - Contains the Computer interface that clients expect
- `devices/` - Contains concrete device implementations (Mac, Windows)
- `adapters/` - Contains adapters that make incompatible devices work with the interface
- `client/` - Contains client code that uses the Computer interface

Run:
```bash
go run .
```

Why this design:
- **Non-invasive**: no changes to `Windows`; the adapter handles translation
- **Single-responsibility**: adapter only maps calls/protocols
- **Practical**: mirrors real-world port/protocol adapters
- **Clear separation**: each component has its own folder for better organization


