# remora

## usage

1. copy `config/envfile.sample` to `config/envfile`
2. write your <b>`MACKEREL_APIKEY`</b> in `config/envfile`
3. build image && `docker run` with `--env-file config/envfile`

### tips
- [a sample](https://github.com/a-know/mackerel-remora) wrotes `apikey` in `sample.yml`. but you can use `MACKEREL_APIKEY` env if don't set `apikey` in yml file.
    - [ref. config/config.go#L101](https://github.com/a-know/mackerel-remora/blob/master/config/config.go#L101)
