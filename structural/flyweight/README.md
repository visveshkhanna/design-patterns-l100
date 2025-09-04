## Flyweight — Share heavy intrinsic state

Goal: use sharing to support large numbers of fine-grained objects efficiently.

This example reuses a small set of icon flyweights (`IconFactory`) across many map markers. The extrinsic state (coordinates, label) lives on each `MapMarker`, while icon bytes are shared.

## Structure

```
flyweight/
├── flyweight/         # Core flyweight objects (intrinsic state)
│   └── icon.go       # Icon flyweight with shared data
├── factory/          # Flyweight factory for managing instances
│   └── icon_factory.go  # Factory to create and cache icons
├── models/           # Context objects (extrinsic state)
│   └── marker.go     # MapMarker that uses flyweight icons
├── main.go          # Demonstration of the pattern
└── README.md        # This file
```

Run:
```bash
go run .
```

Why this design:
- **Memory savings**: many markers, few icon instances
- **Clear split**: intrinsic (icon data) vs extrinsic (position/label)
- **Measurable**: the demo prints counts and estimated memory usage
- **Organized**: separated concerns into logical packages


