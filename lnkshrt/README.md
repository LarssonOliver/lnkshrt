# lnkshrt
A link shortener written in go.

## Container Image
The latest version is available as `larssonoliver/lnkshrt:latest` from DockerHub.

## Environment Variables
| Variable | Description | Default |
| - | - | - |
| `LNKSHRT_PORT` | Port to listen to for requests | `8080` |
| `LNKSHRT_IDSIZE` | Size of shortened URL ids in characters | `6` |
| `LNKSHRT_DBFILE` | Database file | `data/lnkshrt.json` |
| `LNKSHRT_PERSISTENT` | Store data in, and load data from, `LNKSHRT_DBFILE` | `undefined (false)` | 
| `LNKSHRT_ORIGINS` | Comma separated list of CORS origins to allow | `undefined (none)` | 
| `LNKSHRT_INDEX_REDIRECT` | A URL to redirect to from the `GET /` route | `undefined (none)` |