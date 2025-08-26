## Bridge — Decouple abstraction from implementation

Goal: separate an abstraction from its implementation so the two can vary independently.

This example bridges `Computer` (Mac/Windows) and `Printer` (HP/Epson). Computers call into a printer via the `Printer` interface. Any computer can work with any printer without combinatorial subclasses.

Run:
```bash
go run .
```

Why this design:
- **Cartesian flexibility**: mix-and-match computers × printers at runtime
- **Avoids subclass explosion**: composition over inheritance
- **Clear roles**: abstraction owns the high-level workflow; implementor does the work


