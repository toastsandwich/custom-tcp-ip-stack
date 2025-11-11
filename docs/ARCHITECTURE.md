# Architecture Overview

This document describes the design and architecture of the Custom TCP/IP Stack project.

## Layers Implemented

- **IP Layer:** Handles addressing, packetization and reassembly, and basic routing.
- **TCP Layer:** Reliable, connection-oriented transport, including sequence numbers and retransmission.
- **Utilities:** Helpers, checksumming, and value conversion.

## High-Level Design Principles

- **Modularity:** Each protocol is in its own package/module.
- **Testability:** Layers can be unit tested independently.
- **Extensibility:** Easy to add protocols, hooks, or logging.

For deeper details, see [IP](IP.md) and [TCP](TCP.md).