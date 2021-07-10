# go-discover

Fork of [`go-ethereum/p2p/discover`](https://github.com/ethereum/go-ethereum/tree/master/p2p/discover), isolated for external usage.

The git-history is preserved with `git subtree split`.
The `go.mod`, `README.md` were added, along with import-path updates.

Logging is abstracted away with a simple `Logger` interface, avoiding the go-ethereum `Logger` entirely.


## License

LGPL v3, see [`LICENSE`](./LICENSE) file.
