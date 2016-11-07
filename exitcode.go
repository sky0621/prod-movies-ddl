package moviesddl

// アプリの終了コード
const (
	ExitCodeOK int = iota
	ExitCodeArgsError
	ExitCodeLogSetupError
	ExitCodeConfigError
	ExitCodeError
)
