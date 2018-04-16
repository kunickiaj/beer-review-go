package cmd

// JiraConfig configuration structure for JIRA
type JiraConfig struct {
	URL      string
	Username string
	Password string
}

// GerritConfig configuration structure for gerrit
type GerritConfig struct {
	URL string
}
