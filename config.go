package log

type Configuration struct {
	Path        string // 文件路径，如：./app.log
	Level       string // 日志输出的级别
	MaxFileSize int    // 日志文件大小的最大值，单位(M)
	MaxBackups  int    // 最多保留备份数
	MaxAge      int    // 日志文件保存的时间，单位(天)
	Compress    bool   // 是否压缩
	Caller      bool   // 日志是否需要显示调用位置
	Stdout      bool   // 是否输出到控制台
	SLog        bool   // 是否使用slog
}

type Option func(c *Configuration)

func SetMaxFileSize(size int) Option {
	return func(c *Configuration) {
		c.MaxFileSize = size
	}
}

func SetMaxBackups(n int) Option {
	return func(c *Configuration) {
		c.MaxBackups = n
	}
}

func SetMaxAge(age int) Option {
	return func(c *Configuration) {
		c.MaxAge = age
	}
}

func SetCompress(compress bool) Option {
	return func(c *Configuration) {
		c.Compress = compress
	}
}

func SetCaller(caller bool) Option {
	return func(c *Configuration) {
		c.Caller = caller
	}
}

func SetStdout(b bool) Option {
	return func(c *Configuration) {
		c.Stdout = b
	}
}

func SetSLog(b bool) Option {
	return func(c *Configuration) {
		c.SLog = b
	}
}
