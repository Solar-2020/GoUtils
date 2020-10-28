package common

type SharedConfig struct {
	Port                           string `envconfig:"PORT" default:"8099"`
	InterviewService			   string `envconfig:"INTERVIEW_SERVICE" default:"localhost:8099"`
	AuthServiceAddress			   string  `envconfig:"AUTH_SERVICE_ADDRESS" default:""`
	GroupServiceAddress			   string  `envconfig:"GROUP_SERVICE_ADDRESS" default:""`
	AccountServiceAddress		   string  `envconfig:"ACCOUNT_SERVICE_ADDRESS" default:""`
}
