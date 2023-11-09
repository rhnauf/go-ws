
# Go-WS

for disclaimer, I didn't implement Supabase for the database for a couple of reason:
* I haven't used Supabase so my knowledge is pretty limited, based on the [official](https://github.com/supabase-community/supabase-go) go Supabase library and the [unofficial](https://github.com/nedpals/supabase-go) 3rd party library, It does not support the realtime db yet
* Supabase is already using Postgres based on the [docs](https://supabase.com/database), so following my previous point I'm just going to save the messages directly to the local db, rather than saving to 2 different postgres instances (I might have misunderstood the flow for this application)
* My frontend skill is pretty subpar, so if I need to utilized Supabase on the frontend side using Javascript, it might have took longer time to finish the task

so for all the reasons above, I decided to skip Supabase implementation
- - - -
run docker for postgres instance

```docker run --name postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=root -d postgres:15-alpine```

run migrations, using your own preferred tools (example using goose)

```goose postgres postgres://root:root@localhost:5432/go-ws up```

create env files & modify to your own configurations

```cp .env.example .env```

build and run

```make build-run```

open localhost:8080 on the browser

to fetch messages data manually hit endpoint below:
```
    localhost:8080/message/list [GET]
    query param: 
        username
        recipient
        start_date
        end_date
        
    example: http://localhost:8080/message/list?username=user%20a&recipient=user%20b&start_date=2023-11-08%2010%3A00%3A00&end_date=2023-11-10%2010%3A00%3A00
```

- - - -
Todo:
- [ ] implement supabase on the frontend
- [ ] implement jwt for auth
- [ ] add more validation