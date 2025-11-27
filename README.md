# BlogAggregatorGo

A small, fast aggregator for collecting and serving blog feeds written in Go. Designed for local development and lightweight deployment — fetches RSS/Atom feeds, normalizes entries, stores them locally.

## Features

- Fetch RSS/Atom feeds on a schedule
- Normalize entries (title, author, published date, content, link)

## Prerequisites

- Go 1.20+ (for building from source)

## Quick start

Set a .gatorconfig.json in your home directory with a psql connection string. Example .gatorconfig.json:

```json
{
  "db_url": "postgres://postgres:postgres@localhost:5432/gator"
}
```

Clone and run:

```bash
git clone https://example.com/you/BlogAggregatorGo.git
cd BlogAggregatorGo
go run . register <name>
go run . login <name>
go run . agg <time_between_reqs>
```
