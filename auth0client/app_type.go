package auth0client

type AppType string

const AppTypeSinglePageApplication = AppType("spa")
const AppTypeNative = AppType("native")
const AppTypeWeb = AppType("regular_web")
const AppTypeNonInteractive = AppType("non_interactive")
