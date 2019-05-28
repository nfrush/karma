package config

import (
	"regexp"
	"time"
)

type alertmanagerConfig struct {
	Name    string
	URI     string
	Timeout time.Duration
	Proxy   bool
	TLS     struct {
		CA                 string
		Cert               string
		Key                string
		InsecureSkipVerify bool `yaml:"insecureSkipVerify"  mapstructure:"insecureSkipVerify"`
	}
	Headers map[string]string
}

type jiraRule struct {
	Regex string
	URI   string
}

type CustomLabelColor struct {
	Value         string         `yaml:"value" mapstructure:"value"`
	ValueRegex    string         `yaml:"value_re" mapstructure:"value_re"`
	CompiledRegex *regexp.Regexp `yaml:"-" mapstructure:"-"`
	Color         string         `yaml:"color" mapstructure:"color"`
}

type CustomLabelColors map[string][]CustomLabelColor

type configSchema struct {
	Alertmanager struct {
		Interval time.Duration
		Servers  []alertmanagerConfig
	}
	Annotations struct {
		Default struct {
			Hidden bool
		}
		Hidden  []string
		Visible []string
		Keep    []string
		Strip   []string
	}
	Custom struct {
		CSS string
		JS  string
	}
	Debug   bool
	Filters struct {
		Default []string
	}
	Grid struct {
		Sorting struct {
			Order        string
			Reverse      bool
			Label        string
			CustomValues struct {
				Labels map[string]map[string]int
			} `yaml:"customValues" mapstructure:"customValues"`
		}
	}
	Labels struct {
		Keep  []string
		Strip []string
		Color struct {
			Custom CustomLabelColors
			Static []string
			Unique []string
		}
	}
	Listen struct {
		Address string
		Port    int
		Prefix  string
	}
	Log struct {
		Config bool
		Level  string
		Format string
	}
	JIRA      []jiraRule
	Receivers struct {
		Keep  []string
		Strip []string
	}
	Sentry struct {
		Private string
		Public  string
	}
	SilenceForm struct {
		Strip struct {
			Labels []string
		}
	} `yaml:"silenceForm"  mapstructure:"silenceForm"`
}
