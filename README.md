# fakepicdar
Fake Picdar API in Go for importing images from a database dump.

Needs a file structure like:

```
./static
├── images
│   └──DB*24080701.jpg
└── xmls
    └── DB*24080701.xml
```

Where `.xml` files are are the `MogulResponse` for a single image record.

## Running

```
go run $GO_PATH/github.com/guardian/fakepicdar/cmd/main.go -basedir="~/static"
```
