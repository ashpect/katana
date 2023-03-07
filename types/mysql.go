package types

type GogsUser struct {
	Id                 int    `json:"id" bson:"id" binding:"required"`
	LowerName          string `json:"lower_name" bson:"lower_name" binding:"required"`
	Name               string `json:"name" bson:"name" binding:"required"`
	FullName           string `json:"full_name" bson:"full_name" binding:"required"`
	Email              string `json:"email" bson:"email" binding:"required"`
	Password           string `json:"passwd" bson:"passwd" binding:"required"`
	LoginType          int    `json:"login_type" bson:"login_type" binding:"required"`
	LoginSource        int    `json:"login_source" bson:"login_source" binding:"required"`
	LoginName          string `json:"login_name" bson:"login_name" binding:"required"`
	Type               int    `json:"type" bson:"type" binding:"required"`
	Location           string `json:"location" bson:"location" binding:"required"`
	Website            string `json:"website" bson:"website" binding:"required"`
	Rands              string `json:"rands" bson:"rands" binding:"required"`
	Salt               string `json:"salt" bson:"salt" binding:"required"`
	CreatedUnix        int64  `json:"created_unix" bson:"created_unix" binding:"required"`
	UpdatedUnix        int64  `json:"updated_unix" bson:"updated_unix" binding:"required"`
	LastRepoVisibility int    `json:"last_repo_visibility" bson:"last_repo_visibility" binding:"required"`
	MaxRepoCreation    int    `json:"max_repo_creation" bson:"max_repo_creation" binding:"required"`
	IsActive           bool   `json:"is_active" bson:"is_active" binding:"required"`
	IsAdmin            bool   `json:"is_admin" bson:"is_admin" binding:"required"`
	AllowGitHook       bool   `json:"allow_git_hook" bson:"allow_git_hook" binding:"required"`
	AllowImportLocal   bool   `json:"allow_import_local" bson:"allow_import_local" binding:"required"`
	ProhibitLogin      bool   `json:"prohibit_login" bson:"prohibit_login" binding:"required"`
	Avatar             string `json:"avatar" bson:"avatar" binding:"required"`
	AvatarEmail        string `json:"avatar_email" bson:"avatar_email" binding:"required"`
	UseCustomAvatar    bool   `json:"use_custom_avatar" bson:"use_custom_avatar" binding:"required"`
	NumFollowers       int    `json:"num_followers" bson:"num_followers" binding:"required"`
	NumFollowing       int    `json:"num_following" bson:"num_following" binding:"required"`
	NumStars           int    `json:"num_stars" bson:"num_stars" binding:"required"`
	NumRepos           int    `json:"num_repos" bson:"num_repos" binding:"required"`
	Description        string `json:"description" bson:"description" binding:"required"`
	NumTeams           int    `json:"num_teams" bson:"num_teams" binding:"required"`
	NumMembers         int    `json:"num_members" bson:"num_members" binding:"required"`
}

type GogsWebhook struct {
	Id           int    `json:"id" bson:"id" binding:"required"`
	RepoId       int    `json:"repo_id" bson:"repo_id" binding:"required"`
	OrgId        int    `json:"org_id" bson:"org_id" binding:"required"`
	Url          string `json:"url" bson:"url" binding:"required"`
	ContentType  string `json:"content_type" bson:"content_type" binding:"required"`
	Secret       string `json:"secret" bson:"secret" binding:"required"`
	Events       string `json:"events" bson:"events" binding:"required"`
	IsSsl        bool   `json:"is_ssl" bson:"is_ssl" binding:"required"`
	IsActive     bool   `json:"is_active" bson:"is_active" binding:"required"`
	HookTaskType int    `json:"hook_task_type" bson:"hook_task_type" binding:"required"`
	Meta         string `json:"meta" bson:"meta" binding:"required"`
	LastStatus   int    `json:"last_status" bson:"last_status" binding:"required"`
	CreatedUnix  int64  `json:"created_unix" bson:"created_unix" binding:"required"`
	UpdatedUnix  int64  `json:"updated_unix" bson:"updated_unix" binding:"required"`
}
