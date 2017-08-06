# gochimp
A cors solution for interacting with mailchimp from a frontend.

This is an example of how to communicate with the mailchimp api using a cors middleware.  You can put your mailchimp api key and list id in the config.toml, as well as the origin you wish to allow communications from.

You can build this using go build -o chimp cmd/gochimp/main.go, then run the resulting binary.

I usually run something like this at, support.my-domain.tld.  It can be run on the same server using nginx as a reverse proxy, which will allow you to run your front end website on the same server.

You can run this long term using systemd.

To actually get the count from the api you'll need to implement MemberCount in internal/stats/repo, associated direction in the comments of that method.
