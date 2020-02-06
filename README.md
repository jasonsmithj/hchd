# h(health check)c(container)h(http)d(datastore)
//todo: english

## Description
K8s/ECS等のコンテナサービスにおいてアプリケーションの完成を待たずに環境構築を行う場合、に
http及びmysqlへのNWレベル疎通確認を行うコンテナです。

![](./.doc/hchd.png)

## Usage
### create .env file
`.env_sample` ファイルをコピーし、 `.env` ファイルを作成してください。

### modify .env file
`.env` ファイルを必要に応じて変更して下さい。

変数名|説明|例
----|----|----
CONTAINER_PORT|httpのポート番号|8081
ENV|環境名(任意)|dev
IMAGE|コンテナイメージ名(image_name:tag)|healthcheck-container-http-mysql:latest
DB_ENDPOINT|DBサーバのエンドポイント|hchd-db
DB_USER|DBユーザ名|healthcheck
DB_PASSWORD|DBパスワード|password
DB_NAME|DBテーブル名|healthcheck
DB_PORT|DBポート|3306

### ECS/K8s
AWS ECRやGCP GCR等にpushし、ご利用している環境にデプロイして下さい。
`command` には `./app` を指定するとサーバが起動します。

### HTTP Origin
#### HTTP healthcheck
`http://your_url/v1/healthcheck`

#### MySQL healthcheck
`http://your_url/v1/mysql/healthcheck`

#### Postgresql healthcheck
`http://your_url/v1/postgres/healthcheck`

#### REDIS healthcheck
`http://your_url/v1/redis/healthcheck`

##### Succeeded Message
`{"message":"Succeeded"}`

##### Failed Message
`{"message":"Failed"}`
