# oscarmlage.com

## Instructions

After `git clone` you can run this site with a simple `make serve`:

```sh
$ make serve
[...]
Attaching to serve_1
serve_1  | Start building sites …
serve_1  | hugo v0.88.0-ACC5EB5B linux/amd64 BuildDate=2021-09-02T09:27:28Z VendorInfo=gohugoio
serve_1  |
serve_1  |                    | EN
serve_1  | -------------------+-----
serve_1  |   Pages            | 34
serve_1  |   Paginator pages  |  0
serve_1  |   Non-page files   |  3
serve_1  |   Static files     | 19
serve_1  |   Processed images |  0
serve_1  |   Aliases          | 14
serve_1  |   Sitemaps         |  1
serve_1  |   Cleaned          |  0
serve_1  |
serve_1  | Built in 293 ms
```

Then the site will be available in `http://localhost:1313`

## Other commands

- `make build`: Builds the static site into `public/` directory
- `make serve`: Launches a server in `http://localhost:1313`
- `make shell`: Enters in the container in order to create new contents, etc....
  Once inside the container you can run some other `hugo` commands:
  - `hugo new posts/slugified-title`: Create a new post.
  - `hugo new projects/whatever`: Create a new project called `whatever`.

