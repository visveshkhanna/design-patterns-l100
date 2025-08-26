## Flyweight — Share heavy intrinsic state

Goal: use sharing to support large numbers of fine-grained objects efficiently.

This example reuses a small set of icon flyweights (`IconFactory`) across many map markers. The extrinsic state (coordinates, label) lives on each `MapMarker`, while icon bytes are shared.

Run:
```bash
go run .
```

Why this design:
- **Memory savings**: many markers, few icon instances
- **Clear split**: intrinsic (icon data) vs extrinsic (position/label)
- **Measurable**: the demo prints counts and estimated memory usage


