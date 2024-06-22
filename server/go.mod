module dataaccess.com/server

go 1.22.2

require (
	dataaccess.com/constant v0.0.0-00010101000000-000000000000
	github.com/lib/pq v1.10.9
)

replace dataaccess.com/constant => ../constant

replace dataaccess.com/model => ../model
