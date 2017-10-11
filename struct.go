/*

LICENSE:  MIT
Author:   sine
Email:    sinerwr@gmail.com

Contributors

*/

package cfg

type ConfigItems struct {
	Version     string `json:"version" redis:"version"`
	WechatToken string `json:"wechat.token" redis:"wechat.token"`
	AAAstatus   string `json:"aaa.status" redis:"aaa.status"`

	PublicTokenStatus string `json:"public.token.status" redis:"public.token.status"`
	PublicTokenExpire string `json:"public.token.expire" redis:"public.token.expire"`

	RedisConfigHost string `json:"redis.config.host" redis:"redis.config.host"`
	RedisConfigPort string `json:"redis.config.port" redis:"redis.config.port"`
	RedisConfigAuth string `json:"redis.config.auth" redis:"redis.config.auth"`

	RedisPublicHost string `json:"redis.public.host" redis:"redis.public.host"`
	RedisPublicPort string `json:"redis.public.port" redis:"redis.public.port"`
	RedisPublicAuth string `json:"redis.public.auth" redis:"redis.public.auth"`

	MongoUserAddress  string `json:"mongo.user.address" redis:"mongo.user.address"`
	MongoUserUsername string `json:"mongo.user.username" redis:"mongo.user.username"`
	MongoUserPassword string `json:"mongo.user.password" redis:"mongo.user.password"`

	MongoCloudAddress  string `json:"mongo.cloud.address" redis:"mongo.cloud.address"`
	MongoCloudUsername string `json:"mongo.cloud.username" redis:"mongo.cloud.username"`
	MongoCloudPassword string `json:"mongo.cloud.password" redis:"mongo.cloud.password"`

	MongoAssetAddress  string `json:"mongo.asset.address" redis:"mongo.asset.address"`
	MongoAssetUsername string `json:"mongo.asset.username" redis:"mongo.asset.username"`
	MongoAssetPassword string `json:"mongo.asset.password" redis:"mongo.asset.password"`

	MysqTestinglHost     string `json:"mysql.testing.host" redis:"mysql.testing.host"`
	MysqTestinglPort     string `json:"mysql.testing.port" redis:"mysql.testing.port"`
	MysqTestinglUsername string `json:"mysql.testing.username" redis:"mysql.testing.username"`
	MysqTestinglPassword string `json:"mysql.testing.password" redis:"mysql.testing.password"`

	RpcHeHost string `json:"rpc.He.host" redis:"rpc.He.host"`
	RpcHePort string `json:"rpc.He.port" redis:"rpc.He.port"`
	RpcLiHost string `json:"rpc.Li.host" redis:"rpc.Li.host"`
	RpcLiPort string `json:"rpc.Li.port" redis:"rpc.Li.port"`
	RpcBeHost string `json:"rpc.Be.host" redis:"rpc.Be.host"`
	RpcBePort string `json:"rpc.Be.port" redis:"rpc.Be.port"`
	RpcBHost  string `json:"rpc.B.host" redis:"rpc.B.host"`
	RpcBPort  string `json:"rpc.B.port" redis:"rpc.B.port"`
	RpcCHost  string `json:"rpc.C.host" redis:"rpc.C.host"`
	RpcCPort  string `json:"rpc.C.port" redis:"rpc.C.port"`
	RpcNHost  string `json:"rpc.N.host" redis:"rpc.N.host"`
	RpcNPort  string `json:"rpc.N.port" redis:"rpc.N.port"`
	RpcOHost  string `json:"rpc.O.host" redis:"rpc.O.host"`
	RpcOPort  string `json:"rpc.O.port" redis:"rpc.O.port"`
	RpcFHost  string `json:"rpc.F.host" redis:"rpc.F.host"`
	RpcFPort  string `json:"rpc.F.port" redis:"rpc.F.port"`
	RpcNeHost string `json:"rpc.Ne.host" redis:"rpc.Ne.host"`
	RpcNePort string `json:"rpc.Ne.port" redis:"rpc.Ne.port"`

	SentryHStatus  string `json:"sentry.H.status" redis:"sentry.H.status"`
	SentryHDSN     string `json:"sentry.H.dsn" redis:"sentry.H.dsn"`
	SentryHeStatus string `json:"sentry.He.status" redis:"sentry.He.status"`
	SentryHeDSN    string `json:"sentry.He.dsn" redis:"sentry.He.dsn"`
	SentryLiStatus string `json:"sentry.Li.status" redis:"sentry.Li.status"`
	SentryLiDSN    string `json:"sentry.Li.dsn" redis:"sentry.Li.dsn"`
	SentryBeStatus string `json:"sentry.Be.status" redis:"sentry.Be.status"`
	SentryBeDSN    string `json:"sentry.Be.dsn" redis:"sentry.Be.dsn"`
	SentryBStatus  string `json:"sentry.B.status" redis:"sentry.B.status"`
	SentryBDSN     string `json:"sentry.B.dsn" redis:"sentry.B.dsn"`
	SentryCStatus  string `json:"sentry.C.status" redis:"sentry.C.status"`
	SentryCDSN     string `json:"sentry.C.dsn" redis:"sentry.C.dsn"`
	SentryNStatus  string `json:"sentry.N.status" redis:"sentry.N.status"`
	SentryNDSN     string `json:"sentry.N.dsn" redis:"sentry.N.dsn"`
	SentryOStatus  string `json:"sentry.O.status" redis:"sentry.O.status"`
	SentryODSN     string `json:"sentry.O.dsn" redis:"sentry.O.dsn"`
	SentryFStatus  string `json:"sentry.F.status" redis:"sentry.F.status"`
	SentryFDSN     string `json:"sentry.F.dsn" redis:"sentry.F.dsn"`
	SentryNeStatus string `json:"sentry.Ne.status" redis:"sentry.Ne.status"`
	SentryNeDSN    string `json:"sentry.Ne.dsn" redis:"sentry.Ne.dsn"`
}
