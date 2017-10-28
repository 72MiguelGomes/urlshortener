### Description

This is a small project to try Golang in:

- WebService Context

- Integration with Couchbase

### API Description

## GET /short/{url}

- Params

    **url:** Url to minify

- Response

    - 200

            {
               "result": "oXFSaHMAPg=="
            }


## GET /resolve/{short_url}

- Params

    **short_url:** Short Url

- Response

    - 302 - Redirect to the correct URL
    - 404 - Not Found