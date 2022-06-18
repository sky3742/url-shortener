# API-based URL shortener service
A service that provides functionality such as Tiny URL

---

### Build
```bash
> go build
```

### Test
```bash
> go test .\test
```

---

### API Specification

##### [POST] /create
> use this api to create a unique shortened url

[JSON] request
| field | type   | description     |
| ----  | ----   | ----            |
| url   | string | destination url |

sample request
```bash
curl -X POST http://localhost:8888/create -d "{\"url\":\"https://www.google.com\"}'
```

[JSON] reponse
| field | type   | description |
| ---- | ---- | ---- |
| url   | string | shortened url |

sample response
```json
{
    "url": "localhost:8888/d6fe7c"
}
```

---

##### [GET] /{id}
> visiting http://localhost:8888/d6fe7c in browser will be redirect to original url: https://www.google.com

---

#### [GET] /stats/{id}
> use this api to get stats of a specific id

[JSON] request
> no

sample request
```bash
curl -X GET http://localhost:8888/stats/d6fe7c
```

[JSON] reponse
| field          | type   | description                        |
| -----          | ----   | ----                               |
| id             | string | unique 6 characters alphanumeric   |
| destination    | string | destination url                    |
| redirect_count | int    | redirect count since created       |
| created_at     | int64  | creation time in unix milliseconds |

sample response
```json
{
    "id": "d6fe7c",
    "destination": "https://www.google.com",
    "redirect_count": 0,
    "created_at": 1655530287098
}
```

---

#### [GET] /stats
> use this api to get ids' stats

[JSON] request
> no

sample request
```bash
curl -X GET http://localhost:8888/stats
```

[JSON] reponse
| field          | type   | description                        |
| -----          | ----   | ----                               |
| id             | string | unique 6 characters alphanumeric   |
| destination    | string | destination url                    |
| redirect_count | int    | redirect count since created       |
| created_at     | int64  | creation time in unix milliseconds |

sample response
```json
[
    {
        "id": "d6fe7c",
        "destination": "https://www.google.com",
        "redirect_count": 0,
        "created_at": 1655530287098
    }
]
```
