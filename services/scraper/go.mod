module scraper

go 1.18

require (
	cloud.google.com/go/storage v1.27.0
	external v0.0.0-00010101000000-000000000000
	github.com/PuerkitoBio/goquery v1.8.0
	github.com/friendsofgo/errors v0.9.2
	github.com/google/go-cmp v0.5.8
	github.com/gorilla/mux v1.8.0
	github.com/lib/pq v1.2.1-0.20191011153232-f91d3411e481
	github.com/volatiletech/sqlboiler/v4 v4.13.0
	github.com/volatiletech/strmangle v0.0.4
)

replace external => ../external

require (
	cloud.google.com/go v0.104.0 // indirect
	cloud.google.com/go/compute v1.7.0 // indirect
	cloud.google.com/go/iam v0.3.0 // indirect
	cloud.google.com/go/secretmanager v1.7.0 // indirect
	github.com/andybalholm/cascadia v1.3.1 // indirect
	github.com/golang/groupcache v0.0.0-20200121045136-8c9f03a8e57e // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/googleapis/enterprise-certificate-proxy v0.1.0 // indirect
	github.com/googleapis/gax-go/v2 v2.5.1 // indirect
	github.com/spf13/cast v1.4.1 // indirect
	github.com/volatiletech/inflect v0.0.1 // indirect
	go.opencensus.io v0.23.0 // indirect
	golang.org/x/net v0.0.0-20220909164309-bea034e7d591 // indirect
	golang.org/x/oauth2 v0.0.0-20220909003341-f21342109be1 // indirect
	golang.org/x/sys v0.0.0-20220728004956-3c1f35247d10 // indirect
	golang.org/x/text v0.3.7 // indirect
	golang.org/x/xerrors v0.0.0-20220609144429-65e65417b02f // indirect
	google.golang.org/api v0.97.0 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/genproto v0.0.0-20220930163606-c98284e70a91 // indirect
	google.golang.org/grpc v1.49.0 // indirect
	google.golang.org/protobuf v1.28.1 // indirect
)