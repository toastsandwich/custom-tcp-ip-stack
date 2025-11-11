# IP Layer Details

This file covers the core implementation of the IP protocol.

## Features

- **Packet Construction:** IP header creation, fragmentation (if applicable).
- **Routing Logic:** Simplistic internal routing and delivery.
- **Checksumming:** IPv4 header checksum implementation.
- **ICMP Handling:** Minimal implementation or stubs for basic tests.

These components are essential for supporting the upper layers, such as TCP.

For broader context, see [Architecture](ARCHITECTURE.md).