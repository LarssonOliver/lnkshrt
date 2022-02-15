# lnkshrt
A link shortener written in go.

## Environment variables
| Variable | Description | Default |
| - | - | - |
| `LNKSHRT_PORT` | Port to listen to for requests | `8080` |
| `LNKSHRT_IDSIZE` | Size of shortened URL ids in characters | `6` |
| `LNKSHRT_DBFILE` | Database file | `data/lnkshrt.json` |
| `LNKSHRT_PERSISTENT` | Store data in, and load data from, `LNKSHRT_DBFILE` | `undefined (false)` | 