package gitlabapi

type project struct {
	Id            int    `json:"id"`
	Name          string `json:"name"`
	SshUrlToRepo  string `json:"ssh_url_to_repo"`
	HttpUrlToRepo string `json:"http_url_to_repo"`
}

type GitlabGroup struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	ParentId int    `json:"parent_id"`
	// {
	// 	"id": 4,
	// 	"name": "Twitter",
	// 	"path": "twitter",
	// 	"description": "Aliquid qui quis dignissimos distinctio ut commodi voluptas est.",
	// 	"visibility": "public",
	// 	"avatar_url": null,
	// 	"web_url": "https://gitlab.example.com/groups/twitter",
	// 	"request_access_enabled": false,
	// 	"full_name": "Twitter",
	// 	"full_path": "twitter",
	// 	"runners_token": "ba324ca7b1c77fc20bb9",
	// 	"file_template_project_id": 1,
	// 	"parent_id": null,
	// 	"created_at": "2020-01-15T12:36:29.590Z",
	// 	"shared_with_groups": [
	// 	  {
	// 		"group_id": 28,
	// 		"group_name": "H5bp",
	// 		"group_full_path": "h5bp",
	// 		"group_access_level": 20,
	// 		"expires_at": null
	// 	  }
	// 	],
	// 	"prevent_sharing_groups_outside_hierarchy": false,
	// 	"projects": [ // Deprecated and will be removed in API v5
	// 	  {
	// 		"id": 7,
	// 		"description": "Voluptas veniam qui et beatae voluptas doloremque explicabo facilis.",
	// 		"default_branch": "master",
	// 		"tag_list": [], //deprecated, use `topics` instead
	// 		"topics": [],
	// 		"archived": false,
	// 		"visibility": "public",
	// 		"ssh_url_to_repo": "git@gitlab.example.com:twitter/typeahead-js.git",
	// 		"http_url_to_repo": "https://gitlab.example.com/twitter/typeahead-js.git",
	// 		"web_url": "https://gitlab.example.com/twitter/typeahead-js",
	// 		"name": "Typeahead.Js",
	// 		"name_with_namespace": "Twitter / Typeahead.Js",
	// 		"path": "typeahead-js",
	// 		"path_with_namespace": "twitter/typeahead-js",
	// 		"issues_enabled": true,
	// 		"merge_requests_enabled": true,
	// 		"wiki_enabled": true,
	// 		"jobs_enabled": true,
	// 		"snippets_enabled": false,
	// 		"container_registry_enabled": true,
	// 		"created_at": "2016-06-17T07:47:25.578Z",
	// 		"last_activity_at": "2016-06-17T07:47:25.881Z",
	// 		"shared_runners_enabled": true,
	// 		"creator_id": 1,
	// 		"namespace": {
	// 		  "id": 4,
	// 		  "name": "Twitter",
	// 		  "path": "twitter",
	// 		  "kind": "group"
	// 		},
	// 		"avatar_url": null,
	// 		"star_count": 0,
	// 		"forks_count": 0,
	// 		"open_issues_count": 3,
	// 		"public_jobs": true,
	// 		"shared_with_groups": [],
	// 		"request_access_enabled": false
	// 	  },
	// 	  {
	// 		"id": 6,
	// 		"description": "Aspernatur omnis repudiandae qui voluptatibus eaque.",
	// 		"default_branch": "master",
	// 		"tag_list": [], //deprecated, use `topics` instead
	// 		"topics": [],
	// 		"archived": false,
	// 		"visibility": "internal",
	// 		"ssh_url_to_repo": "git@gitlab.example.com:twitter/flight.git",
	// 		"http_url_to_repo": "https://gitlab.example.com/twitter/flight.git",
	// 		"web_url": "https://gitlab.example.com/twitter/flight",
	// 		"name": "Flight",
	// 		"name_with_namespace": "Twitter / Flight",
	// 		"path": "flight",
	// 		"path_with_namespace": "twitter/flight",
	// 		"issues_enabled": true,
	// 		"merge_requests_enabled": true,
	// 		"wiki_enabled": true,
	// 		"jobs_enabled": true,
	// 		"snippets_enabled": false,
	// 		"container_registry_enabled": true,
	// 		"created_at": "2016-06-17T07:47:24.661Z",
	// 		"last_activity_at": "2016-06-17T07:47:24.838Z",
	// 		"shared_runners_enabled": true,
	// 		"creator_id": 1,
	// 		"namespace": {
	// 		  "id": 4,
	// 		  "name": "Twitter",
	// 		  "path": "twitter",
	// 		  "kind": "group"
	// 		},
	// 		"avatar_url": null,
	// 		"star_count": 0,
	// 		"forks_count": 0,
	// 		"open_issues_count": 8,
	// 		"public_jobs": true,
	// 		"shared_with_groups": [],
	// 		"request_access_enabled": false
	// 	  }
	// 	],
	// 	"shared_projects": [ // Deprecated and will be removed in API v5
	// 	  {
	// 		"id": 8,
	// 		"description": "Velit eveniet provident fugiat saepe eligendi autem.",
	// 		"default_branch": "master",
	// 		"tag_list": [], //deprecated, use `topics` instead
	// 		"topics": [],
	// 		"archived": false,
	// 		"visibility": "private",
	// 		"ssh_url_to_repo": "git@gitlab.example.com:h5bp/html5-boilerplate.git",
	// 		"http_url_to_repo": "https://gitlab.example.com/h5bp/html5-boilerplate.git",
	// 		"web_url": "https://gitlab.example.com/h5bp/html5-boilerplate",
	// 		"name": "Html5 Boilerplate",
	// 		"name_with_namespace": "H5bp / Html5 Boilerplate",
	// 		"path": "html5-boilerplate",
	// 		"path_with_namespace": "h5bp/html5-boilerplate",
	// 		"issues_enabled": true,
	// 		"merge_requests_enabled": true,
	// 		"wiki_enabled": true,
	// 		"jobs_enabled": true,
	// 		"snippets_enabled": false,
	// 		"container_registry_enabled": true,
	// 		"created_at": "2016-06-17T07:47:27.089Z",
	// 		"last_activity_at": "2016-06-17T07:47:27.310Z",
	// 		"shared_runners_enabled": true,
	// 		"creator_id": 1,
	// 		"namespace": {
	// 		  "id": 5,
	// 		  "name": "H5bp",
	// 		  "path": "h5bp",
	// 		  "kind": "group"
	// 		},
	// 		"avatar_url": null,
	// 		"star_count": 0,
	// 		"forks_count": 0,
	// 		"open_issues_count": 4,
	// 		"public_jobs": true,
	// 		"shared_with_groups": [
	// 		  {
	// 			"group_id": 4,
	// 			"group_name": "Twitter",
	// 			"group_full_path": "twitter",
	// 			"group_access_level": 30,
	// 			"expires_at": null
	// 		  },
	// 		  {
	// 			"group_id": 3,
	// 			"group_name": "Gitlab Org",
	// 			"group_full_path": "gitlab-org",
	// 			"group_access_level": 10,
	// 			"expires_at": "2018-08-14"
	// 		  }
	// 		]
	// 	  }
	// 	],
	// 	"ip_restriction_ranges": null
	//   }
}
type GitlabProject struct {
	Id        int             `json:"id"`
	Namespace GitlabNamespace `json:"namespace"`

	// {
	//     "id": 3,
	//     "description": null,
	//     "default_branch": "master",
	//     "visibility": "private",
	//     "ssh_url_to_repo": "git@example.com:diaspora/diaspora-project-site.git",
	//     "http_url_to_repo": "http://example.com/diaspora/diaspora-project-site.git",
	//     "web_url": "http://example.com/diaspora/diaspora-project-site",
	//     "readme_url": "http://example.com/diaspora/diaspora-project-site/blob/master/README.md",
	//     "tag_list": [ //deprecated, use `topics` instead
	//       "example",
	//       "disapora project"
	//     ],
	//     "topics": [
	//       "example",
	//       "disapora project"
	//     ],
	//     "owner": {
	//       "id": 3,
	//       "name": "Diaspora",
	//       "created_at": "2013-09-30T13:46:02Z"
	//     },
	//     "name": "Diaspora Project Site",
	//     "name_with_namespace": "Diaspora / Diaspora Project Site",
	//     "path": "diaspora-project-site",
	//     "path_with_namespace": "diaspora/diaspora-project-site",
	//     "issues_enabled": true,
	//     "open_issues_count": 1,
	//     "merge_requests_enabled": true,
	//     "jobs_enabled": true,
	//     "wiki_enabled": true,
	//     "snippets_enabled": false,
	//     "can_create_merge_request_in": true,
	//     "resolve_outdated_diff_discussions": false,
	//     "container_registry_enabled": false, // deprecated, use container_registry_access_level instead
	//     "container_registry_access_level": "disabled",
	//     "security_and_compliance_access_level": "disabled",
	//     "container_expiration_policy": {
	//       "cadence": "7d",
	//       "enabled": false,
	//       "keep_n": null,
	//       "older_than": null,
	//       "name_regex": null, // to be deprecated in GitLab 13.0 in favor of `name_regex_delete`
	//       "name_regex_delete": null,
	//       "name_regex_keep": null,
	//       "next_run_at": "2020-01-07T21:42:58.658Z"
	//     },
	//     "created_at": "2013-09-30T13:46:02Z",
	//     "last_activity_at": "2013-09-30T13:46:02Z",
	//     "creator_id": 3,
	//     "namespace": {
	//       "id": 3,
	//       "name": "Diaspora",
	//       "path": "diaspora",
	//       "kind": "group",
	//       "full_path": "diaspora",
	//       "avatar_url": "http://localhost:3000/uploads/group/avatar/3/foo.jpg",
	//       "web_url": "http://localhost:3000/groups/diaspora"
	//     },
	//     "import_status": "none",
	//     "import_error": null,
	//     "permissions": {
	//       "project_access": {
	//         "access_level": 10,
	//         "notification_level": 3
	//       },
	//       "group_access": {
	//         "access_level": 50,
	//         "notification_level": 3
	//       }
	//     },
	//     "archived": false,
	//     "avatar_url": "http://example.com/uploads/project/avatar/3/uploads/avatar.png",
	//     "license_url": "http://example.com/diaspora/diaspora-client/blob/master/LICENSE",
	//     "license": {
	//       "key": "lgpl-3.0",
	//       "name": "GNU Lesser General Public License v3.0",
	//       "nickname": "GNU LGPLv3",
	//       "html_url": "http://choosealicense.com/licenses/lgpl-3.0/",
	//       "source_url": "http://www.gnu.org/licenses/lgpl-3.0.txt"
	//     },
	//     "shared_runners_enabled": true,
	//     "group_runners_enabled": true,
	//     "forks_count": 0,
	//     "star_count": 0,
	//     "runners_token": "b8bc4a7a29eb76ea83cf79e4908c2b",
	//     "ci_default_git_depth": 50,
	//     "ci_forward_deployment_enabled": true,
	//     "ci_allow_fork_pipelines_to_run_in_parent_project": true,
	//     "ci_separated_caches": true,
	//     "public_jobs": true,
	//     "shared_with_groups": [
	//       {
	//         "group_id": 4,
	//         "group_name": "Twitter",
	//         "group_full_path": "twitter",
	//         "group_access_level": 30
	//       },
	//       {
	//         "group_id": 3,
	//         "group_name": "Gitlab Org",
	//         "group_full_path": "gitlab-org",
	//         "group_access_level": 10
	//       }
	//     ],
	//     "repository_storage": "default",
	//     "only_allow_merge_if_pipeline_succeeds": false,
	//     "allow_merge_on_skipped_pipeline": false,
	//     "restrict_user_defined_variables": false,
	//     "only_allow_merge_if_all_discussions_are_resolved": false,
	//     "remove_source_branch_after_merge": false,
	//     "printing_merge_requests_link_enabled": true,
	//     "request_access_enabled": false,
	//     "merge_method": "merge",
	//     "squash_option": "default_on",
	//     "auto_devops_enabled": true,
	//     "auto_devops_deploy_strategy": "continuous",
	//     "approvals_before_merge": 0,
	//     "mirror": false,
	//     "mirror_user_id": 45,
	//     "mirror_trigger_builds": false,
	//     "only_mirror_protected_branches": false,
	//     "mirror_overwrites_diverged_branches": false,
	//     "external_authorization_classification_label": null,
	//     "packages_enabled": true,
	//     "service_desk_enabled": false,
	//     "service_desk_address": null,
	//     "autoclose_referenced_issues": true,
	//     "suggestion_commit_message": null,
	//     "enforce_auth_checks_on_uploads": true,
	//     "merge_commit_template": null,
	//     "squash_commit_template": null,
	//     "issue_branch_template": "gitlab/%{id}-%{title}",
	//     "marked_for_deletion_at": "2020-04-03", // Deprecated and will be removed in API v5 in favor of marked_for_deletion_on
	//     "marked_for_deletion_on": "2020-04-03",
	//     "compliance_frameworks": [ "sox" ],
	//     "statistics": {
	//       "commit_count": 37,
	//       "storage_size": 1038090,
	//       "repository_size": 1038090,
	//       "wiki_size" : 0,
	//       "lfs_objects_size": 0,
	//       "job_artifacts_size": 0,
	//       "pipeline_artifacts_size": 0,
	//       "packages_size": 0,
	//       "snippets_size": 0,
	//       "uploads_size": 0
	//     },
	//     "container_registry_image_prefix": "registry.example.com/diaspora/diaspora-client",
	//     "_links": {
	//       "self": "http://example.com/api/v4/projects",
	//       "issues": "http://example.com/api/v4/projects/1/issues",
	//       "merge_requests": "http://example.com/api/v4/projects/1/merge_requests",
	//       "repo_branches": "http://example.com/api/v4/projects/1/repository_branches",
	//       "labels": "http://example.com/api/v4/projects/1/labels",
	//       "events": "http://example.com/api/v4/projects/1/events",
	//       "members": "http://example.com/api/v4/projects/1/members",
	//       "cluster_agents": "http://example.com/api/v4/projects/1/cluster_agents"
	//     }
	//   }
}

type GitlabNamespace struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

// "namespace": {
//       "id": 3,
//       "name": "Diaspora",
//       "path": "diaspora",
//       "kind": "group",
//       "full_path": "diaspora",
//       "avatar_url": "http://localhost:3000/uploads/group/avatar/3/foo.jpg",
//       "web_url": "http://localhost:3000/groups/diaspora"
//     },

type Variable struct {
	VariableType     string `json:"variable_type"`
	Key              string `json:"key"`
	Value            string `json:"value"`
	Protected        bool   `json:"protected"`
	Masked           bool   `json:"masked"`
	Raw              bool   `json:"raw"`
	EnvironmentScope string `json:"environment_scope"`
}
