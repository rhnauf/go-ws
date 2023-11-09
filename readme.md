
# Go-WS

for disclaimer, I didn't implement Supabase for the database for a couple of reason:
* I haven't used Supabase so my knowledge is pretty limited, based on the go official Supabase [library](https://github.com/supabase-community/supabase-go) and the [unofficial](https://github.com/nedpals/supabase-go) 3rd party library, It does not support the realtime db yet
* Supabase is already using Postgres based on the [docs](https://supabase.com/database), so following my previous point I'm just going to save the messages directly to the local db, rather than saving to 2 different postgres instances (I might have misunderstood the flow for this application)
* My frontend skill is pretty subpar, so if I need to utilized Supabase on the frontend side using Javascript, it might have took longer time to finish the task

so for all the reasons above, I decided to skip Supabase implementation
- - - -
run migrations, using your own preferred tools (example using goose)

```goose postgres postgres://root:admin123@localhost:5433/go-ws up```

- - - -
Todo:
- [ ] implement supabase
- [ ] implement jwt for auth
- [ ] add more validation