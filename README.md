# GO + NATS

Это девконтейнер со втроенным сервисом `NATS` и `NATS CLI`. А заодно небольшой гайд по использованию. 

Сервер уже добавлен в сервер по умолчанию. И можно пользоваться

```go
nats.Connect(nats.DefaultURL)
```

## CLI

Сервер уже настроет, поэтому можно сразу пользоваться без лишних флагов. Например, запустим в одном терминале:

```bash
nats sub com.hello.*
```

Теперь этот терминал будет читать все сообщения из топика `com.hello`, например `com.hello.foo` и `com.hello.bar`. Чтобы попробовать как это работает откроем другой терминал и запустим там

```bash
echo "hello on com.hello.foo" | nats pub com.hello.foo
echo "hello on com.hello.bar" | nats pub com.hello.bar
echo "hello on com.another-topic.bar" | nats pub com.another-topic.bar
```

Последнее сообщение, конечно не дошло подписчику, потому что мы на него не подписывались. И вот ещё хороший пример

```bash
echo "hello on com.hello.bar.fail" | nats pub com.hello.bar.fail
```

Это сообщение тоже не дойдет. Потому что `*` в топике отвеает за один токен между точками, если один или больше, нужно использовать `>`, чтобы прочитать сообщение из `com.hello.bar.fail` нам бы понадобился подписчик `com.hello.>`, то есть

```bash
nats sub 'com.hello.>'
```


