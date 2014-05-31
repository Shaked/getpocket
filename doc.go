// Copyright 2014 github.com/Shaked Author. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

/*
Package Shaked/getpocket is an implementation for Getpocket.com's API.
The Package is documented at:

	http://godoc.org/github.com/Shaked/getpocket

Let's start with authenticating to Getpocket using Shaked/getpocket package:

	type Page struct {
        auth *auth.Auth
    }

    var (
        redirectURI     = "https://localhost:10443/authcheck"
        consumerKey     = "yourconsumerkeyhere"
        ssl_certificate = "ssl/server.crt"
        ssl_key         = "ssl/server.key"
    )

    func main() {
        var e *utils.RequestError
        a, e := auth.Factory(consumerKey, redirectURI)
        if nil != e {
            log.Fatal(e)
        }
        log.Printf("Listen on 10443")
        p := &Page{auth: a}
        http.HandleFunc("/auth", p.Auth)
        http.HandleFunc("/authcheck", p.AuthCheck)

        err := http.ListenAndServeTLS(":10443", ssl_certificate, ssl_key, nil)
        if err != nil {
            log.Fatal(err)
        }
    }

    func (p *Page) Auth(w http.ResponseWriter, r *http.Request) {
        log.Println(r.Referer(), "Example Handler, Should connect and request app permissions")
        requestToken, err := p.auth.Connect()
        if nil != err {
            fmt.Fprintf(w, "Token error %s (%d)", err.Error(), err.ErrorCode())
        }

        p.auth.RequestPermissions(requestToken, w, r)
    }

    func (p *Page) AuthCheck(w http.ResponseWriter, r *http.Request) {
        log.Println(r.Referer(), "GetPocket connection check, should get the username and access token")
        requestToken := r.URL.Query().Get("requestToken")
        if "" == requestToken {
            fmt.Fprintf(w, "Request token is invalid")
            return
        }
        user, err := p.auth.User(requestToken)
        if nil != err {
            fmt.Fprintf(w, "%s (%d)", err.Error(), err.ErrorCode())
            return
        }
        fmt.Fprintf(w, "%#v", user)
    }

Later on we can apply Getpocket's commands.

Add:

	c := commands.New(user, consumerKey)
	u := "http://confreaks.com/videos/3432-gophercon2014-go-from-c-to-go"
	add := commands.NewAdd(u)
	add.SetTitle("Some cool title").SetTags("Shaked,Blog")
	resp, e := c.Exec(add)
	if nil != e {
	    fmt.Fprintf(w, "ERROR%s\n", e)
	    return
	}

Retrieve:

	retrieve := commands.NewRetrieve()
	resp, e := c.Exec(retrieve)
	if nil != e {
	    fmt.Fprintf(w, "ERROR%s\n", e)
	    return
	}

Modify:

	id := 12345678
	favorite := modify.FactoryFavorite(id)
	unfavorite := modify.FactoryUnfavorite(id)
	add := modify.FactoryAdd(id)
	add.SetTags([]string{"go", "programming", "blog"})
	actions := []modify.Action{
		add,
		favorite,
		unfavorite,
	}
	modify := commands.NewModify(actions)
	c := commands.New(user, consumerKey)
	resp, e := c.Exec(modify)
	if nil != e {
		fmt.Fprintf(w, "ERROR%s\n", e)
	}

The modify command contains the following actions:

	add
	archive
	delete
	favorite
	unfavorite
	readd
	tag_add
	tag_clear
	tag_remove
	tag_rename
	tag_replace

Check the repository for more details and examples:

	https://www.github.com/Shaked/getpocket

*/
package getpocket
