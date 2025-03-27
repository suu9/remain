# Remain

A lightweight utility for tracking remaining tasks or items.

## Features

- Simple and intuitive interface
- Lightweight with no dependencies
- Cross-platform compatibility

## Installation

```bash
npm install remain
```

## Usage

```javascript
const remain = require('remain');

// Initialize with total items
const tracker = remain(100);

// Update remaining count
tracker.update(75);

// Get current status
console.log(tracker.status()); // "25 remaining (25%)"
```

## License

MIT