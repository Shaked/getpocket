Implementation for [GetPocket](www.getpocket.com) API using Go
===========

[![Build Status](https://travis-ci.org/Shaked/gogetpocket.svg)](https://travis-ci.org/Shaked/getpocket)

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

### Example

```
    var a *auth.Auth
    func main(){
        certificate := "ssl/server.crt"
        key := "ssl/server.key"
        a, err := auth.New(consumerKey, redirectURI)
        if nil != err {
            log.Fatal(err)
        }
        InitHandlers()
        err := http.ListenAndServeTLS(":10443", certificate, key, nil)
        if err != nil {
            log.Fatal(err)
        }
    }

    func InitHandlers() {
        http.HandleFunc("/auth", HandlerExample)
        http.HandleFunc("/authcheck", HandlerExampleGetPocketCheck)
    }

    func HandlerExample(w http.ResponseWriter, r *http.Request) {
        log.Println(r.Referer(), "Example Handler, Should connect and request app permissions")
        requestToken, err := a.Connect()
        if nil != err {
            fmt.Fprintf(w, "Token error %s (%d)", err.Error(), err.ErrorCode())
        }

        a.RequestPermissions(requestToken, w, r)
    }

    func HandlerExampleGetPocketCheck(w http.ResponseWriter, r *http.Request) {
        log.Println(r.Referer(), "GetPocket connection check, should get the username and access token")
        requestToken := r.URL.Query().Get("requestToken")
        if "" == requestToken {
            fmt.Fprintf(w, "Request token is invalid")
            return
        }
        user, err := a.User(requestToken)
        if nil != err {
            fmt.Fprintf(w, "%s (%d)", err.Error(), err.ErrorCode())
            return
        }

        fmt.Fprintf(w, "%s, %s", user.AccessToken, user.Username)
    }
```

### Notes

- GetPocket **requires** to work with HTTPS only. 
- go version > 1.0
