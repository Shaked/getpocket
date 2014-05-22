## Implementation for [GetPocket](www.getpocket.com) API using Go


[![Build Status](https://travis-ci.org/Shaked/getpocket.svg?branch=master)](https://travis-ci.org/Shaked/getpocket)

### Public Interfaces

#### [auth](http://getpocket.com/developer/docs/authentication)

##### auth.Authenticator

```
    RequestPermissions(requestToken string, w http.ResponseWriter, r *http.Request)
    Connect() (string, *AuthError)
    User() (*AuthUser, *AuthError)
```

##### auth.DeviceControllable

```
    SetForceMobile(forceMobile bool)
```

#### commands 

```
    New(user *auth.User, consumerKey string) *Commands
    Exec(command Executable) (Response, error)
```

##### commands.Executable

```
    Exec(user *auth.User, consumerKey string, request utils.HttpRequest) (Response, error)
```

##### [Add](http://getpocket.com/developer/docs/v3/add)

```
    Executable
    NewAdd(targetURL string) *Add
    SetTitle(title string) *Add
    SetTags(tags string) *Add
    SetTweetID(tweet_id string) *Add
```


#### Commands - [Retrieve](http://getpocket.com/developer/docs/v3/retrieve)

```
    Executable
    NewRetrieve() *Retrieve
    SetFavorite(favorite bool) *Retrieve
    SetTag(tag string) *Retrieve
    SetUntagged() *Retrieve
    SetContentType(contentType string) error
    SetSort(sort string) error
    SetDetailType(detailType string) error
```


### Example

```
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
        c := commands.New(user, consumerKey)
        u := "http://confreaks.com/videos/3432-gophercon2014-go-from-c-to-go"
        add := commands.NewAdd(u)
        add.SetTitle("Some cool title").SetTags("Shaked,Blog")
        resp1, e := c.Exec(add)
        if nil != e {
            fmt.Fprintf(w, "ERROR%s\n", e)
            return
        }
        retrieve := commands.NewRetrieve()
        resp2, e := c.Exec(retrieve)
        if nil != e {
            fmt.Fprintf(w, "ERROR%s\n", e)
            return
        }
        fmt.Fprintf(w, "%s, %s, %#v, %#v", user.AccessToken, user.Username, resp1, resp2)
    }

```

### Notes

- GetPocket **requires** to work with HTTPS only. 
- go version > 1.0
