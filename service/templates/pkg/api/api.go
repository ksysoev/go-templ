package api

type Service struct{
  config Config
}

type Config struct{
	Listen: string `mapstructure:"listen"`
}
