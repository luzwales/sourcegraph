// GENERATED CODE - DO NOT EDIT!
//
// Generated by:
//
//   go run gen_context_and_mock.go -o1 context.go -o2 mockstore/mockstores.go
//
// Called via:
//
//   go generate
//

package mockstore

import (
	"sourcegraph.com/sourcegraph/sourcegraph/pkg/store"
	srcstore "sourcegraph.com/sourcegraph/srclib/store"
)

// Stores has a field for each store interface with the concrete mock type (to obviate the need for tedious type assertions in test code).
type Stores struct {
	Accounts           Accounts
	BuildLogs          BuildLogs
	Builds             Builds
	DefExamples        DefExamples
	Defs               Defs
	Directory          Directory
	ExternalAuthTokens ExternalAuthTokens
	GlobalDeps         GlobalDeps
	GlobalRefs         GlobalRefs
	Graph              srcstore.MockMultiRepoStore
	Password           Password
	Queue              Queue
	RepoConfigs        RepoConfigs
	RepoStatuses       RepoStatuses
	RepoVCS            RepoVCS
	Repos              Repos
	Users              Users
}

func (s *Stores) Stores() store.Stores {
	return store.Stores{
		Accounts:           &s.Accounts,
		BuildLogs:          &s.BuildLogs,
		Builds:             &s.Builds,
		DefExamples:        &s.DefExamples,
		Defs:               &s.Defs,
		Directory:          &s.Directory,
		ExternalAuthTokens: &s.ExternalAuthTokens,
		GlobalDeps:         &s.GlobalDeps,
		GlobalRefs:         &s.GlobalRefs,
		Graph:              &s.Graph,
		Password:           &s.Password,
		Queue:              &s.Queue,
		RepoConfigs:        &s.RepoConfigs,
		RepoStatuses:       &s.RepoStatuses,
		RepoVCS:            &s.RepoVCS,
		Repos:              &s.Repos,
		Users:              &s.Users,
	}
}
