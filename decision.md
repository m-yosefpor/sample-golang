
- using mongodb or sql or redis?
- embed endpoints in users or refrence: embed, because each ep only belongs to one user
- each endpoint in one goroutine? or each user only?: don't afraid to use more goroutines :D also we can have different interval-time for each ep this way
- not like health-exporter, config changes on runtime!! new goroutine for new endpoint/users??
- same set of probers for same endpoints for all users?: not good security/stability. some ep might have been registered much before and have high ep.failed, or users want different thresholds
- stat for specific call DURING TODAY!!! need tsdb? timestamp?: ignore this for now, we can see stats from the beginging


- delete urls?
- multiple urls different threshold/intervals? id?

scale multiple replicas:
- prober will face issues


veryyy frequent:
    update user.endpoint.success

frequent:
    update user.endpoint.fail
    update user.endpoint.remain


medium:
    update users.endpoints

rare:
    update users
