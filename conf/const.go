package conf

// HTTP启动项
const App = "conf/app.ini"
const HttpServer = "HttpServer"
const RpcConnect = "RpcConnect"

// RPC连接
const DefaultIp = "127.0.0.1"
const DefaultHttpPort = 8080
const DefaultRpcPort = 50052

// jwt
const JWTKeyValue = "lzc_query_system"
const JWTExpireValue = 7200
const JWTFlag = "jwt_flag"        // 标识本次登录是否携带JWT
const JWTClaims = "jwt_claims"    // 存储jwt的声明字段
const JWTHeader = "Authorization" // jwt请求头标识字段

// user
const User_ID = "user_id"
const User_PWD = "user_pwd"
