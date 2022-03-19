package appservice

import (
	"github.com/javijuol/git-remind/app/cli/cliglobalopts"
	"github.com/javijuol/git-remind/domain"
	"github.com/javijuol/git-remind/infra"
)

var GetPathPatterns = domain.NewGetPathPatterns(
	domain.MultipleGetPathPatterns(
		cliglobalopts.GetPathPatterns,
		infra.GitGlobalConfigGetPathPatterns,
	),
)
var GetRepos = domain.NewGetRepos(GetPathPatterns, infra.FilesystemRepos)
var GetRepoStatuses = domain.NewGetRepoStatuses(GetRepos, infra.GetGitStatus)
var NotifyRepoStatues = domain.NewNotifyRepoStatuses(GetRepoStatuses, infra.BeeepRepoStatusNotifier)
