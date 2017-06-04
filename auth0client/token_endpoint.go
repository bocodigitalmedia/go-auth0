package auth0client

type TokenEndpoint string

const TokenEndpointAuthMethodNative = TokenEndpoint("native")
const TokenEndpointClientSecretPost = TokenEndpoint("client_secret_post")
const TokenEndpointClientSecretBasic = TokenEndpoint("client_secret_basic")
