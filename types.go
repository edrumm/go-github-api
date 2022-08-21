package go_github_api

type Owner struct {
	Login      string `json:"login"`
	Id         int    `json:"id"`
	AvatarUrl  string `json:"avatar_url"`
	GravatarId string `json:"gravatar_id"`
	Url        string `json:"url"`
	HtmlUrl    string `json:"html_url"`
}

type License struct {
	Key  string `json:"key"`
	Name string `json:"name"`
	Url  string `json:"url"`
}

type Repository struct {
	Id              int      `json:"id"`
	NodeId          string   `json:"node_id"`
	Name            string   `json:"name"`
	Owner           *Owner   `json:"owner"`
	HtmlUrl         string   `json:"html_url"`
	Description     string   `json:"description"`
	Fork            bool     `json:"fork"`
	CreatedAt       string   `json:"created_at"`
	UpdatedAt       string   `json:"updated_at"`
	PushedAt        string   `json:"pushed_at"`
	GitUrl          string   `json:"git_url"`
	SshUrl          string   `json:"ssh_url"`
	CloneUrl        string   `json:"clone_url"`
	SvnUrl          string   `json:"svn_url"`
	Homepage        string   `json:"homepage"`
	Size            int      `json:"size"`
	StargazersCount int      `json:"stargazers_count"`
	WatchersCount   int      `json:"watchers_count"`
	Language        string   `json:"language"`
	ForksCount      int      `json:"forks_count"`
	Archived        bool     `json:"archived"`
	Disabled        bool     `json:"disabled"`
	OpenIssuesCount int      `json:"open_issues_count"`
	License         *License `json:"license"`
	AllowForking    bool     `json:"allow_forking"`
	IsTemplate      bool     `json:"is_template"`
	Topics          []string `json:"topics"`
	DefaultBranch   string   `json:"default_branch"`
}
