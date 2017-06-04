package auth0client

type TokenEndpointAuthMethod string

const TokenEndpointAuthMethodNative = TokenEndpointAuthMethod("native")
const TokenEndpointAuthMethodClientSecretPost = TokenEndpointAuthMethod("client_secret_post")
const TokenEndpointAuthMethodClientSecretBasic = TokenEndpointAuthMethod("client_secret_basic")
