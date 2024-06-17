// SPDX-License-Identifier: MIT
// SPDX-FileCopyrightText: © 2015 LabStack LLC and Echo contributors

package echo

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	staticRoutes = []*Route{
		{"GET", "/", ""},
		{"GET", "/cmd.html", ""},
		{"GET", "/code.html", ""},
		{"GET", "/contrib.html", ""},
		{"GET", "/contribute.html", ""},
		{"GET", "/debugging_with_gdb.html", ""},
		{"GET", "/docs.html", ""},
		{"GET", "/effective_go.html", ""},
		{"GET", "/files.log", ""},
		{"GET", "/gccgo_contribute.html", ""},
		{"GET", "/gccgo_install.html", ""},
		{"GET", "/go-logo-black.png", ""},
		{"GET", "/go-logo-blue.png", ""},
		{"GET", "/go-logo-white.png", ""},
		{"GET", "/go1.1.html", ""},
		{"GET", "/go1.2.html", ""},
		{"GET", "/go1.html", ""},
		{"GET", "/go1compat.html", ""},
		{"GET", "/go_faq.html", ""},
		{"GET", "/go_mem.html", ""},
		{"GET", "/go_spec.html", ""},
		{"GET", "/help.html", ""},
		{"GET", "/ie.css", ""},
		{"GET", "/install-source.html", ""},
		{"GET", "/install.html", ""},
		{"GET", "/logo-153x55.png", ""},
		{"GET", "/Makefile", ""},
		{"GET", "/root.html", ""},
		{"GET", "/share.png", ""},
		{"GET", "/sieve.gif", ""},
		{"GET", "/tos.html", ""},
		{"GET", "/articles/", ""},
		{"GET", "/articles/go_command.html", ""},
		{"GET", "/articles/index.html", ""},
		{"GET", "/articles/wiki/", ""},
		{"GET", "/articles/wiki/edit.html", ""},
		{"GET", "/articles/wiki/final-noclosure.go", ""},
		{"GET", "/articles/wiki/final-noerror.go", ""},
		{"GET", "/articles/wiki/final-parsetemplate.go", ""},
		{"GET", "/articles/wiki/final-template.go", ""},
		{"GET", "/articles/wiki/final.go", ""},
		{"GET", "/articles/wiki/get.go", ""},
		{"GET", "/articles/wiki/http-sample.go", ""},
		{"GET", "/articles/wiki/index.html", ""},
		{"GET", "/articles/wiki/Makefile", ""},
		{"GET", "/articles/wiki/notemplate.go", ""},
		{"GET", "/articles/wiki/part1-noerror.go", ""},
		{"GET", "/articles/wiki/part1.go", ""},
		{"GET", "/articles/wiki/part2.go", ""},
		{"GET", "/articles/wiki/part3-errorhandling.go", ""},
		{"GET", "/articles/wiki/part3.go", ""},
		{"GET", "/articles/wiki/test.bash", ""},
		{"GET", "/articles/wiki/test_edit.good", ""},
		{"GET", "/articles/wiki/test_Test.txt.good", ""},
		{"GET", "/articles/wiki/test_view.good", ""},
		{"GET", "/articles/wiki/view.html", ""},
		{"GET", "/codewalk/", ""},
		{"GET", "/codewalk/codewalk.css", ""},
		{"GET", "/codewalk/codewalk.js", ""},
		{"GET", "/codewalk/codewalk.xml", ""},
		{"GET", "/codewalk/functions.xml", ""},
		{"GET", "/codewalk/markov.go", ""},
		{"GET", "/codewalk/markov.xml", ""},
		{"GET", "/codewalk/pig.go", ""},
		{"GET", "/codewalk/popout.png", ""},
		{"GET", "/codewalk/run", ""},
		{"GET", "/codewalk/sharemem.xml", ""},
		{"GET", "/codewalk/urlpoll.go", ""},
		{"GET", "/devel/", ""},
		{"GET", "/devel/release.html", ""},
		{"GET", "/devel/weekly.html", ""},
		{"GET", "/gopher/", ""},
		{"GET", "/gopher/appenginegopher.jpg", ""},
		{"GET", "/gopher/appenginegophercolor.jpg", ""},
		{"GET", "/gopher/appenginelogo.gif", ""},
		{"GET", "/gopher/bumper.png", ""},
		{"GET", "/gopher/bumper192x108.png", ""},
		{"GET", "/gopher/bumper320x180.png", ""},
		{"GET", "/gopher/bumper480x270.png", ""},
		{"GET", "/gopher/bumper640x360.png", ""},
		{"GET", "/gopher/doc.png", ""},
		{"GET", "/gopher/frontpage.png", ""},
		{"GET", "/gopher/gopherbw.png", ""},
		{"GET", "/gopher/gophercolor.png", ""},
		{"GET", "/gopher/gophercolor16x16.png", ""},
		{"GET", "/gopher/help.png", ""},
		{"GET", "/gopher/pkg.png", ""},
		{"GET", "/gopher/project.png", ""},
		{"GET", "/gopher/ref.png", ""},
		{"GET", "/gopher/run.png", ""},
		{"GET", "/gopher/talks.png", ""},
		{"GET", "/gopher/pencil/", ""},
		{"GET", "/gopher/pencil/gopherhat.jpg", ""},
		{"GET", "/gopher/pencil/gopherhelmet.jpg", ""},
		{"GET", "/gopher/pencil/gophermega.jpg", ""},
		{"GET", "/gopher/pencil/gopherrunning.jpg", ""},
		{"GET", "/gopher/pencil/gopherswim.jpg", ""},
		{"GET", "/gopher/pencil/gopherswrench.jpg", ""},
		{"GET", "/play/", ""},
		{"GET", "/play/fib.go", ""},
		{"GET", "/play/hello.go", ""},
		{"GET", "/play/life.go", ""},
		{"GET", "/play/peano.go", ""},
		{"GET", "/play/pi.go", ""},
		{"GET", "/play/sieve.go", ""},
		{"GET", "/play/solitaire.go", ""},
		{"GET", "/play/tree.go", ""},
		{"GET", "/progs/", ""},
		{"GET", "/progs/cgo1.go", ""},
		{"GET", "/progs/cgo2.go", ""},
		{"GET", "/progs/cgo3.go", ""},
		{"GET", "/progs/cgo4.go", ""},
		{"GET", "/progs/defer.go", ""},
		{"GET", "/progs/defer.out", ""},
		{"GET", "/progs/defer2.go", ""},
		{"GET", "/progs/defer2.out", ""},
		{"GET", "/progs/eff_bytesize.go", ""},
		{"GET", "/progs/eff_bytesize.out", ""},
		{"GET", "/progs/eff_qr.go", ""},
		{"GET", "/progs/eff_sequence.go", ""},
		{"GET", "/progs/eff_sequence.out", ""},
		{"GET", "/progs/eff_unused1.go", ""},
		{"GET", "/progs/eff_unused2.go", ""},
		{"GET", "/progs/error.go", ""},
		{"GET", "/progs/error2.go", ""},
		{"GET", "/progs/error3.go", ""},
		{"GET", "/progs/error4.go", ""},
		{"GET", "/progs/go1.go", ""},
		{"GET", "/progs/gobs1.go", ""},
		{"GET", "/progs/gobs2.go", ""},
		{"GET", "/progs/image_draw.go", ""},
		{"GET", "/progs/image_package1.go", ""},
		{"GET", "/progs/image_package1.out", ""},
		{"GET", "/progs/image_package2.go", ""},
		{"GET", "/progs/image_package2.out", ""},
		{"GET", "/progs/image_package3.go", ""},
		{"GET", "/progs/image_package3.out", ""},
		{"GET", "/progs/image_package4.go", ""},
		{"GET", "/progs/image_package4.out", ""},
		{"GET", "/progs/image_package5.go", ""},
		{"GET", "/progs/image_package5.out", ""},
		{"GET", "/progs/image_package6.go", ""},
		{"GET", "/progs/image_package6.out", ""},
		{"GET", "/progs/interface.go", ""},
		{"GET", "/progs/interface2.go", ""},
		{"GET", "/progs/interface2.out", ""},
		{"GET", "/progs/json1.go", ""},
		{"GET", "/progs/json2.go", ""},
		{"GET", "/progs/json2.out", ""},
		{"GET", "/progs/json3.go", ""},
		{"GET", "/progs/json4.go", ""},
		{"GET", "/progs/json5.go", ""},
		{"GET", "/progs/run", ""},
		{"GET", "/progs/slices.go", ""},
		{"GET", "/progs/timeout1.go", ""},
		{"GET", "/progs/timeout2.go", ""},
		{"GET", "/progs/update.bash", ""},
	}

	gitHubAPI = []*Route{
		// OAuth Authorizations
		{"GET", "/authorizations", ""},
		{"GET", literal_4602, ""},
		{"POST", "/authorizations", ""},

		{"PUT", "/authorizations/clients/:client_id", ""},
		{"PATCH", literal_4602, ""},

		{"DELETE", literal_4602, ""},
		{"GET", "/applications/:client_id/tokens/:access_token", ""},
		{"DELETE", "/applications/:client_id/tokens", ""},
		{"DELETE", "/applications/:client_id/tokens/:access_token", ""},

		// Activity
		{"GET", "/events", ""},
		{"GET", "/repos/:owner/:repo/events", ""},
		{"GET", "/networks/:owner/:repo/events", ""},
		{"GET", "/orgs/:org/events", ""},
		{"GET", "/users/:user/received_events", ""},
		{"GET", "/users/:user/received_events/public", ""},
		{"GET", "/users/:user/events", ""},
		{"GET", "/users/:user/events/public", ""},
		{"GET", "/users/:user/events/orgs/:org", ""},
		{"GET", "/feeds", ""},
		{"GET", "/notifications", ""},
		{"GET", "/repos/:owner/:repo/notifications", ""},
		{"PUT", "/notifications", ""},
		{"PUT", "/repos/:owner/:repo/notifications", ""},
		{"GET", "/notifications/threads/:id", ""},

		{"PATCH", "/notifications/threads/:id", ""},

		{"GET", literal_7620, ""},
		{"PUT", literal_7620, ""},
		{"DELETE", literal_7620, ""},
		{"GET", "/repos/:owner/:repo/stargazers", ""},
		{"GET", "/users/:user/starred", ""},
		{"GET", "/user/starred", ""},
		{"GET", literal_5690, ""},
		{"PUT", literal_5690, ""},
		{"DELETE", literal_5690, ""},
		{"GET", "/repos/:owner/:repo/subscribers", ""},
		{"GET", "/users/:user/subscriptions", ""},
		{"GET", "/user/subscriptions", ""},
		{"GET", literal_3975, ""},
		{"PUT", literal_3975, ""},
		{"DELETE", literal_3975, ""},
		{"GET", literal_3691, ""},
		{"PUT", literal_3691, ""},
		{"DELETE", literal_3691, ""},

		// Gists
		{"GET", "/users/:user/gists", ""},
		{"GET", "/gists", ""},

		{"GET", "/gists/public", ""},
		{"GET", "/gists/starred", ""},

		{"GET", literal_3801, ""},
		{"POST", "/gists", ""},

		{"PATCH", literal_3801, ""},

		{"PUT", literal_3672, ""},
		{"DELETE", literal_3672, ""},
		{"GET", literal_3672, ""},
		{"POST", "/gists/:id/forks", ""},
		{"DELETE", literal_3801, ""},

		// Git Data
		{"GET", "/repos/:owner/:repo/git/blobs/:sha", ""},
		{"POST", "/repos/:owner/:repo/git/blobs", ""},
		{"GET", "/repos/:owner/:repo/git/commits/:sha", ""},
		{"POST", "/repos/:owner/:repo/git/commits", ""},

		{"GET", literal_2893, ""},

		{"GET", "/repos/:owner/:repo/git/refs", ""},
		{"POST", "/repos/:owner/:repo/git/refs", ""},

		{"PATCH", literal_2893, ""},
		{"DELETE", literal_2893, ""},

		{"GET", "/repos/:owner/:repo/git/tags/:sha", ""},
		{"POST", "/repos/:owner/:repo/git/tags", ""},
		{"GET", "/repos/:owner/:repo/git/trees/:sha", ""},
		{"POST", "/repos/:owner/:repo/git/trees", ""},

		// Issues
		{"GET", "/issues", ""},
		{"GET", "/user/issues", ""},
		{"GET", "/orgs/:org/issues", ""},
		{"GET", "/repos/:owner/:repo/issues", ""},
		{"GET", "/repos/:owner/:repo/issues/:number", ""},
		{"POST", "/repos/:owner/:repo/issues", ""},

		{"PATCH", "/repos/:owner/:repo/issues/:number", ""},

		{"GET", "/repos/:owner/:repo/assignees", ""},
		{"GET", "/repos/:owner/:repo/assignees/:assignee", ""},
		{"GET", "/repos/:owner/:repo/issues/:number/comments", ""},

		{"GET", "/repos/:owner/:repo/issues/comments", ""},
		{"GET", literal_7382, ""},

		{"POST", "/repos/:owner/:repo/issues/:number/comments", ""},

		{"PATCH", literal_7382, ""},
		{"DELETE", literal_7382, ""},

		{"GET", "/repos/:owner/:repo/issues/:number/events", ""},

		{"GET", "/repos/:owner/:repo/issues/events", ""},
		{"GET", "/repos/:owner/:repo/issues/events/:id", ""},

		{"GET", "/repos/:owner/:repo/labels", ""},
		{"GET", literal_5602, ""},
		{"POST", "/repos/:owner/:repo/labels", ""},

		{"PATCH", literal_5602, ""},

		{"DELETE", literal_5602, ""},
		{"GET", literal_8274, ""},
		{"POST", literal_8274, ""},
		{"DELETE", "/repos/:owner/:repo/issues/:number/labels/:name", ""},
		{"PUT", literal_8274, ""},
		{"DELETE", literal_8274, ""},
		{"GET", "/repos/:owner/:repo/milestones/:number/labels", ""},
		{"GET", "/repos/:owner/:repo/milestones", ""},
		{"GET", literal_4896, ""},
		{"POST", "/repos/:owner/:repo/milestones", ""},

		{"PATCH", literal_4896, ""},

		{"DELETE", literal_4896, ""},

		// Miscellaneous
		{"GET", "/emojis", ""},
		{"GET", "/gitignore/templates", ""},
		{"GET", "/gitignore/templates/:name", ""},
		{"POST", "/markdown", ""},
		{"POST", "/markdown/raw", ""},
		{"GET", "/meta", ""},
		{"GET", "/rate_limit", ""},

		// Organizations
		{"GET", "/users/:user/orgs", ""},
		{"GET", "/user/orgs", ""},
		{"GET", "/orgs/:org", ""},

		{"PATCH", "/orgs/:org", ""},

		{"GET", "/orgs/:org/members", ""},
		{"GET", "/orgs/:org/members/:user", ""},
		{"DELETE", "/orgs/:org/members/:user", ""},
		{"GET", "/orgs/:org/public_members", ""},
		{"GET", literal_6792, ""},
		{"PUT", literal_6792, ""},
		{"DELETE", literal_6792, ""},
		{"GET", "/orgs/:org/teams", ""},
		{"GET", literal_5497, ""},
		{"POST", "/orgs/:org/teams", ""},

		{"PATCH", literal_5497, ""},

		{"DELETE", literal_5497, ""},
		{"GET", "/teams/:id/members", ""},
		{"GET", literal_3908, ""},
		{"PUT", literal_3908, ""},
		{"DELETE", literal_3908, ""},
		{"GET", "/teams/:id/repos", ""},
		{"GET", literal_5761, ""},
		{"PUT", literal_5761, ""},
		{"DELETE", literal_5761, ""},
		{"GET", "/user/teams", ""},

		// Pull Requests
		{"GET", "/repos/:owner/:repo/pulls", ""},
		{"GET", "/repos/:owner/:repo/pulls/:number", ""},
		{"POST", "/repos/:owner/:repo/pulls", ""},

		{"PATCH", "/repos/:owner/:repo/pulls/:number", ""},

		{"GET", "/repos/:owner/:repo/pulls/:number/commits", ""},
		{"GET", "/repos/:owner/:repo/pulls/:number/files", ""},
		{"GET", "/repos/:owner/:repo/pulls/:number/merge", ""},
		{"PUT", "/repos/:owner/:repo/pulls/:number/merge", ""},
		{"GET", "/repos/:owner/:repo/pulls/:number/comments", ""},

		{"GET", "/repos/:owner/:repo/pulls/comments", ""},
		{"GET", literal_1928, ""},

		{"PUT", "/repos/:owner/:repo/pulls/:number/comments", ""},

		{"PATCH", literal_1928, ""},
		{"DELETE", literal_1928, ""},

		// Repositories
		{"GET", "/user/repos", ""},
		{"GET", "/users/:user/repos", ""},
		{"GET", "/orgs/:org/repos", ""},
		{"GET", "/repositories", ""},
		{"POST", "/user/repos", ""},
		{"POST", "/orgs/:org/repos", ""},
		{"GET", literal_8766, ""},

		{"PATCH", literal_8766, ""},

		{"GET", "/repos/:owner/:repo/contributors", ""},
		{"GET", "/repos/:owner/:repo/languages", ""},
		{"GET", "/repos/:owner/:repo/teams", ""},
		{"GET", "/repos/:owner/:repo/tags", ""},
		{"GET", "/repos/:owner/:repo/branches", ""},
		{"GET", "/repos/:owner/:repo/branches/:branch", ""},
		{"DELETE", literal_8766, ""},
		{"GET", "/repos/:owner/:repo/collaborators", ""},
		{"GET", literal_3125, ""},
		{"PUT", literal_3125, ""},
		{"DELETE", literal_3125, ""},
		{"GET", "/repos/:owner/:repo/comments", ""},
		{"GET", "/repos/:owner/:repo/commits/:sha/comments", ""},
		{"POST", "/repos/:owner/:repo/commits/:sha/comments", ""},
		{"GET", literal_1385, ""},

		{"PATCH", literal_1385, ""},

		{"DELETE", literal_1385, ""},
		{"GET", "/repos/:owner/:repo/commits", ""},
		{"GET", "/repos/:owner/:repo/commits/:sha", ""},
		{"GET", "/repos/:owner/:repo/readme", ""},

		//{"GET", "/repos/:owner/:repo/contents/*path", ""},
		//{"PUT", "/repos/:owner/:repo/contents/*path", ""},
		//{"DELETE", "/repos/:owner/:repo/contents/*path", ""},

		{"GET", "/repos/:owner/:repo/:archive_format/:ref", ""},

		{"GET", "/repos/:owner/:repo/keys", ""},
		{"GET", literal_0735, ""},
		{"POST", "/repos/:owner/:repo/keys", ""},

		{"PATCH", literal_0735, ""},

		{"DELETE", literal_0735, ""},
		{"GET", "/repos/:owner/:repo/downloads", ""},
		{"GET", "/repos/:owner/:repo/downloads/:id", ""},
		{"DELETE", "/repos/:owner/:repo/downloads/:id", ""},
		{"GET", "/repos/:owner/:repo/forks", ""},
		{"POST", "/repos/:owner/:repo/forks", ""},
		{"GET", "/repos/:owner/:repo/hooks", ""},
		{"GET", literal_1408, ""},
		{"POST", "/repos/:owner/:repo/hooks", ""},

		{"PATCH", literal_1408, ""},

		{"POST", "/repos/:owner/:repo/hooks/:id/tests", ""},
		{"DELETE", literal_1408, ""},
		{"POST", "/repos/:owner/:repo/merges", ""},
		{"GET", "/repos/:owner/:repo/releases", ""},
		{"GET", literal_3521, ""},
		{"POST", "/repos/:owner/:repo/releases", ""},

		{"PATCH", literal_3521, ""},

		{"DELETE", literal_3521, ""},
		{"GET", "/repos/:owner/:repo/releases/:id/assets", ""},
		{"GET", "/repos/:owner/:repo/stats/contributors", ""},
		{"GET", "/repos/:owner/:repo/stats/commit_activity", ""},
		{"GET", "/repos/:owner/:repo/stats/code_frequency", ""},
		{"GET", "/repos/:owner/:repo/stats/participation", ""},
		{"GET", "/repos/:owner/:repo/stats/punch_card", ""},
		{"GET", "/repos/:owner/:repo/statuses/:ref", ""},
		{"POST", "/repos/:owner/:repo/statuses/:ref", ""},

		// Search
		{"GET", "/search/repositories", ""},
		{"GET", "/search/code", ""},
		{"GET", "/search/issues", ""},
		{"GET", "/search/users", ""},
		{"GET", "/legacy/issues/search/:owner/:repository/:state/:keyword", ""},
		{"GET", "/legacy/repos/search/:keyword", ""},
		{"GET", "/legacy/user/search/:keyword", ""},
		{"GET", "/legacy/user/email/:email", ""},

		// Users
		{"GET", "/users/:user", ""},
		{"GET", "/user", ""},

		{"PATCH", "/user", ""},

		{"GET", literal_7186, ""},
		{"GET", literal_3017, ""},
		{"POST", literal_3017, ""},
		{"DELETE", literal_3017, ""},
		{"GET", "/users/:user/followers", ""},
		{"GET", "/user/followers", ""},
		{"GET", "/users/:user/following", ""},
		{"GET", "/user/following", ""},
		{"GET", literal_5296, ""},
		{"GET", "/users/:user/following/:target_user", ""},
		{"PUT", literal_5296, ""},
		{"DELETE", literal_5296, ""},
		{"GET", "/users/:user/keys", ""},
		{"GET", "/user/keys", ""},
		{"GET", literal_2071, ""},
		{"POST", "/user/keys", ""},

		{"PATCH", literal_2071, ""},

		{"DELETE", literal_2071, ""},
	}

	parseAPI = []*Route{
		// Objects
		{"POST", "/1/classes/:className", ""},
		{"GET", literal_3671, ""},
		{"PUT", literal_3671, ""},
		{"GET", "/1/classes/:className", ""},
		{"DELETE", literal_3671, ""},

		// Users
		{"POST", "/1/users", ""},
		{"GET", "/1/login", ""},
		{"GET", literal_2043, ""},
		{"PUT", literal_2043, ""},
		{"GET", "/1/users", ""},
		{"DELETE", literal_2043, ""},
		{"POST", "/1/requestPasswordReset", ""},

		// Roles
		{"POST", "/1/roles", ""},
		{"GET", literal_9741, ""},
		{"PUT", literal_9741, ""},
		{"GET", "/1/roles", ""},
		{"DELETE", literal_9741, ""},

		// Files
		{"POST", "/1/files/:fileName", ""},

		// Analytics
		{"POST", "/1/events/:eventName", ""},

		// Push Notifications
		{"POST", "/1/push", ""},

		// Installations
		{"POST", "/1/installations", ""},
		{"GET", literal_7512, ""},
		{"PUT", literal_7512, ""},
		{"GET", "/1/installations", ""},
		{"DELETE", literal_7512, ""},

		// Cloud Functions
		{"POST", "/1/functions", ""},
	}

	googlePlusAPI = []*Route{
		// People
		{"GET", "/people/:userId", ""},
		{"GET", "/people", ""},
		{"GET", "/activities/:activityId/people/:collection", ""},
		{"GET", "/people/:userId/people/:collection", ""},
		{"GET", "/people/:userId/openIdConnect", ""},

		// Activities
		{"GET", "/people/:userId/activities/:collection", ""},
		{"GET", "/activities/:activityId", ""},
		{"GET", "/activities", ""},

		// Comments
		{"GET", "/activities/:activityId/comments", ""},
		{"GET", "/comments/:commentId", ""},

		// Moments
		{"POST", "/people/:userId/moments/:collection", ""},
		{"GET", "/people/:userId/moments/:collection", ""},
		{"DELETE", "/moments/:id", ""},
	}

	paramAndAnyAPI = []*Route{
		{"GET", literal_8639, ""},
		{"GET", literal_2618, ""},
		{"GET", literal_4350, ""},
		{"GET", literal_2693, ""},
		{"GET", literal_3596, ""},
		{"GET", literal_0413, ""},

		{"POST", literal_8639, ""},
		{"POST", literal_2618, ""},
		{"POST", literal_4350, ""},
		{"POST", literal_2693, ""},
		{"POST", literal_3596, ""},
		{"POST", literal_0413, ""},

		{"PUT", literal_8639, ""},
		{"PUT", literal_2618, ""},
		{"PUT", literal_4350, ""},
		{"PUT", literal_2693, ""},
		{"PUT", literal_3596, ""},
		{"PUT", literal_0413, ""},

		{"DELETE", literal_8639, ""},
		{"DELETE", literal_2618, ""},
		{"DELETE", literal_4350, ""},
		{"DELETE", literal_2693, ""},
		{"DELETE", literal_3596, ""},
		{"DELETE", literal_0413, ""},
	}

	paramAndAnyAPIToFind = []*Route{
		{"GET", literal_8675, ""},
		{"GET", literal_3472, ""},
		{"GET", literal_7230, ""},
		{"GET", literal_9437, ""},
		{"GET", literal_3276, ""},
		{"GET", literal_3476, ""},

		{"POST", literal_8675, ""},
		{"POST", literal_3472, ""},
		{"POST", literal_7230, ""},
		{"POST", literal_9437, ""},
		{"POST", literal_3276, ""},
		{"POST", literal_3476, ""},

		{"PUT", literal_8675, ""},
		{"PUT", literal_3472, ""},
		{"PUT", literal_7230, ""},
		{"PUT", literal_9437, ""},
		{"PUT", literal_3276, ""},
		{"PUT", literal_3476, ""},

		{"DELETE", literal_8675, ""},
		{"DELETE", literal_3472, ""},
		{"DELETE", literal_7230, ""},
		{"DELETE", literal_9437, ""},
		{"DELETE", literal_3276, ""},
		{"DELETE", literal_3476, ""},
	}

	missesAPI = []*Route{
		{"GET", literal_4371, ""},
		{"GET", literal_1794, ""},
		{"GET", literal_7945, ""},
		{"GET", literal_1064, ""},

		{"POST", literal_4371, ""},
		{"POST", literal_1794, ""},
		{"POST", literal_7945, ""},
		{"POST", literal_1064, ""},

		{"PUT", literal_4371, ""},
		{"PUT", literal_1794, ""},
		{"PUT", literal_7945, ""},
		{"PUT", literal_1064, ""},

		{"DELETE", literal_4371, ""},
		{"DELETE", literal_1794, ""},
		{"DELETE", literal_7945, ""},
		{"DELETE", literal_1064, ""},
	}

	// handlerHelper created a function that will set a context key for assertion
	handlerHelper = func(key string, value int) func(c Context) error {
		return func(c Context) error {
			c.Set(key, value)
			c.Set("path", c.Path())
			return nil
		}
	}
	handlerFunc = func(c Context) error {
		c.Set("path", c.Path())
		return nil
	}
)

func checkUnusedParamValues(t *testing.T, c *context, expectParam map[string]string) {
	for i, p := range c.pnames {
		value := c.pvalues[i]
		if value != "" {
			if expectParam == nil {
				t.Errorf("pValue '%v' is set for param name '%v' but we are not expecting it with expectParam", value, p)
			} else {
				if _, ok := expectParam[p]; !ok {
					t.Errorf("pValue '%v' is set for param name '%v' but we are not expecting it with expectParam", value, p)
				}
			}
		}
	}
}

func TestRouterStatic(t *testing.T) {
	e := New()
	r := e.router
	path := "/folders/a/files/echo.gif"
	r.Add(http.MethodGet, path, handlerFunc)
	c := e.NewContext(nil, nil).(*context)

	r.Find(http.MethodGet, path, c)
	c.handler(c)

	assert.Equal(t, path, c.Get("path"))
}

func TestRouterNoRoutablePath(t *testing.T) {
	e := New()
	r := e.router
	c := e.NewContext(nil, nil).(*context)

	r.Find(http.MethodGet, "/notfound", c)
	c.handler(c)

	// No routable path, don't set Path.
	assert.Equal(t, "", c.Path())
}

func TestRouterParam(t *testing.T) {
	e := New()
	r := e.router

	r.Add(http.MethodGet, literal_8032, handlerFunc)

	var testCases = []struct {
		name        string
		whenURL     string
		expectRoute interface{}
		expectParam map[string]string
	}{
		{
			name:        "route /users/1 to /users/:id",
			whenURL:     literal_0618,
			expectRoute: literal_8032,
			expectParam: map[string]string{"id": "1"},
		},
		{
			name:        "route /users/1/ to /users/:id",
			whenURL:     "/users/1/",
			expectRoute: literal_8032,
			expectParam: map[string]string{"id": "1/"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			c := e.NewContext(nil, nil).(*context)
			r.Find(http.MethodGet, tc.whenURL, c)

			c.handler(c)
			assert.Equal(t, tc.expectRoute, c.Get("path"))
			for param, expectedValue := range tc.expectParam {
				assert.Equal(t, expectedValue, c.Param(param))
			}
			checkUnusedParamValues(t, c, tc.expectParam)
		})
	}
}

func TestRouteraddAndMatchAllSupportedMethods(t *testing.T) {
	var testCases = []struct {
		name            string
		givenNoAddRoute bool
		whenMethod      string
		expectPath      string
		expectError     string
	}{
		{name: "ok, CONNECT", whenMethod: http.MethodConnect},
		{name: "ok, DELETE", whenMethod: http.MethodDelete},
		{name: "ok, GET", whenMethod: http.MethodGet},
		{name: "ok, HEAD", whenMethod: http.MethodHead},
		{name: "ok, OPTIONS", whenMethod: http.MethodOptions},
		{name: "ok, PATCH", whenMethod: http.MethodPatch},
		{name: "ok, POST", whenMethod: http.MethodPost},
		{name: "ok, PROPFIND", whenMethod: PROPFIND},
		{name: "ok, PUT", whenMethod: http.MethodPut},
		{name: "ok, TRACE", whenMethod: http.MethodTrace},
		{name: "ok, REPORT", whenMethod: REPORT},
		{name: "ok, NON_TRADITIONAL_METHOD", whenMethod: "NON_TRADITIONAL_METHOD"},
		{
			name:            "ok, NOT_EXISTING_METHOD",
			whenMethod:      "NOT_EXISTING_METHOD",
			givenNoAddRoute: true,
			expectPath:      "/*",
			expectError:     "code=405, message=Method Not Allowed",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			e := New()

			e.GET("/*", handlerFunc)

			if !tc.givenNoAddRoute {
				e.Add(tc.whenMethod, "/my/*", handlerFunc)
			}

			req := httptest.NewRequest(tc.whenMethod, "/my/some-url", nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec).(*context)

			e.router.Find(tc.whenMethod, "/my/some-url", c)
			err := c.handler(c)

			if tc.expectError != "" {
				assert.EqualError(t, err, tc.expectError)
			} else {
				assert.NoError(t, err)
			}

			expectPath := "/my/*"
			if tc.expectPath != "" {
				expectPath = tc.expectPath
			}
			assert.Equal(t, expectPath, c.Path())
		})
	}
}

func TestMethodNotAllowedAndNotFound(t *testing.T) {
	e := New()
	r := e.router

	// Routes
	r.Add(http.MethodGet, "/*", handlerFunc)
	r.Add(http.MethodPost, literal_8032, handlerFunc)

	var testCases = []struct {
		name              string
		whenMethod        string
		whenURL           string
		expectRoute       interface{}
		expectParam       map[string]string
		expectError       error
		expectAllowHeader string
	}{
		{
			name:        "exact match for route+method",
			whenMethod:  http.MethodPost,
			whenURL:     literal_0618,
			expectRoute: literal_8032,
			expectParam: map[string]string{"id": "1"},
		},
		{
			name:              "matches node but not method. sends 405 from best match node",
			whenMethod:        http.MethodPut,
			whenURL:           literal_0618,
			expectRoute:       nil,
			expectError:       ErrMethodNotAllowed,
			expectAllowHeader: "OPTIONS, POST",
		},
		{
			name:        "best match is any route up in tree",
			whenMethod:  http.MethodGet,
			whenURL:     literal_0618,
			expectRoute: "/*",
			expectParam: map[string]string{"*": "users/1"},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			req := httptest.NewRequest(tc.whenMethod, tc.whenURL, nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec).(*context)

			method := http.MethodGet
			if tc.whenMethod != "" {
				method = tc.whenMethod
			}
			r.Find(method, tc.whenURL, c)
			err := c.handler(c)

			if tc.expectError != nil {
				assert.Equal(t, tc.expectError, err)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tc.expectRoute, c.Get("path"))
			for param, expectedValue := range tc.expectParam {
				assert.Equal(t, expectedValue, c.Param(param))
			}
			checkUnusedParamValues(t, c, tc.expectParam)

			assert.Equal(t, tc.expectAllowHeader, c.Response().Header().Get(HeaderAllow))
		})
	}
}

func TestRouterOptionsMethodHandler(t *testing.T) {
	e := New()

	var keyInContext interface{}
	e.Use(func(next HandlerFunc) HandlerFunc {
		return func(c Context) error {
			err := next(c)
			keyInContext = c.Get(ContextKeyHeaderAllow)
			return err
		}
	})
	e.GET("/test", func(c Context) error {
		return c.String(http.StatusOK, "Echo!")
	})

	req := httptest.NewRequest(http.MethodOptions, "/test", nil)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusNoContent, rec.Code)
	assert.Equal(t, "OPTIONS, GET", rec.Header().Get(HeaderAllow))
	assert.Equal(t, "OPTIONS, GET", keyInContext)
}

func TestRouterTwoParam(t *testing.T) {
	e := New()
	r := e.router
	r.Add(http.MethodGet, literal_8670, handlerFunc)
	c := e.NewContext(nil, nil).(*context)

	r.Find(http.MethodGet, "/users/1/files/1", c)

	assert.Equal(t, "1", c.Param("uid"))
	assert.Equal(t, "1", c.Param("fid"))
}

// Issue #378
func TestRouterParamWithSlash(t *testing.T) {
	e := New()
	r := e.router

	r.Add(http.MethodGet, "/a/:b/c/d/:e", handlerFunc)
	r.Add(http.MethodGet, "/a/:b/c/:d/:f", handlerFunc)

	c := e.NewContext(nil, nil).(*context)
	r.Find(http.MethodGet, "/a/1/c/d/2/3", c) // `2/3` should mapped to path `/a/:b/c/d/:e` and into `:e`

	err := c.handler(c)
	assert.Equal(t, "/a/:b/c/d/:e", c.Get("path"))
	assert.NoError(t, err)
}

// Issue #1754 - router needs to backtrack multiple levels upwards in tree to find the matching route
// route evaluation order
//
// Routes:
// 1) /a/:b/c
// 2) /a/c/d
// 3) /a/c/df
//
// 4) /a/*/f
// 5) /:e/c/f
//
// 6) /*
//
// Searching route for "/a/c/f" should match "/a/*/f"
// When route `4) /a/*/f` is not added then request for "/a/c/f" should match literal_5382
//
//	              +----------+
//	        +-----+ "/" root +--------------------+--------------------------+
//	        |     +----------+                    |                          |
//	        |                                     |                          |
//	+-------v-------+                         +---v---------+        +-------v---+
//	| "a/" (static) +---------------+         | ":" (param) |        | "*" (any) |
//	+-+----------+--+               |         +-----------+-+        +-----------+
//	  |          |                  |                     |
//
// +---------------v+  +-- ---v------+    +------v----+          +-----v-----------+
// | "c/d" (static) |  | ":" (param) |    | "*" (any) |          | "/c/f" (static) |
// +---------+------+  +--------+----+    +----------++          +-----------------+
//
//	|                  |                    |
//	|                  |                    |
//
// +---------v----+      +------v--------+    +------v--------+
// | "f" (static) |      | "/c" (static) |    | "/f" (static) |
// +--------------+      +---------------+    +---------------+
func TestRouteMultiLevelBacktracking(t *testing.T) {
	var testCases = []struct {
		name        string
		whenURL     string
		expectRoute interface{}
		expectParam map[string]string
	}{
		{
			name:        literal_9425,
			whenURL:     literal_8765,
			expectRoute: literal_8765,
		},
		{
			name:        "route /a/x/df to /a/:b/c",
			whenURL:     "/a/x/c",
			expectRoute: literal_0972,
			expectParam: map[string]string{"b": "x"},
		},
		{
			name:        "route /a/x/f to /a/*/f",
			whenURL:     "/a/x/f",
			expectRoute: "/a/*/f",
			expectParam: map[string]string{"*": "x/f"}, // NOTE: `x` would be probably more suitable
		},
		{
			name:        "route /b/c/f to /:e/c/f",
			whenURL:     "/b/c/f",
			expectRoute: literal_5382,
			expectParam: map[string]string{"e": "b"},
		},
		{
			name:        "route /b/c/c to /*",
			whenURL:     "/b/c/c",
			expectRoute: "/*",
			expectParam: map[string]string{"*": "b/c/c"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			e := New()
			r := e.router

			r.Add(http.MethodGet, literal_0972, handlerHelper("case", 1))
			r.Add(http.MethodGet, "/a/c/d", handlerHelper("case", 2))
			r.Add(http.MethodGet, literal_8765, handlerHelper("case", 3))
			r.Add(http.MethodGet, "/a/*/f", handlerHelper("case", 4))
			r.Add(http.MethodGet, literal_5382, handlerHelper("case", 5))
			r.Add(http.MethodGet, "/*", handlerHelper("case", 6))

			c := e.NewContext(nil, nil).(*context)
			r.Find(http.MethodGet, tc.whenURL, c)

			c.handler(c)
			assert.Equal(t, tc.expectRoute, c.Get("path"))
			for param, expectedValue := range tc.expectParam {
				assert.Equal(t, expectedValue, c.Param(param))
			}
			checkUnusedParamValues(t, c, tc.expectParam)
		})
	}
}

// Issue #1754 - router needs to backtrack multiple levels upwards in tree to find the matching route
// route evaluation order
//
// Request for "/a/c/f" should match literal_5382
//
//	                       +-0,7--------+
//	                       | "/" (root) |----------------------------------+
//	                       +------------+                                  |
//	                            |      |                                   |
//	                            |      |                                   |
//	        +-1,6-----------+   |      |          +-8-----------+   +------v----+
//	        | "a/" (static) +<--+      +--------->+ ":" (param) |   | "*" (any) |
//	        +---------------+                     +-------------+   +-----------+
//	           |          |                             |
//	+-2--------v-----+   +v-3,5--------+       +-9------v--------+
//	| "c/d" (static) |   | ":" (param) |       | "/c/f" (static) |
//	+----------------+   +-------------+       +-----------------+
//	                      |
//	                 +-4--v----------+
//	                 | "/c" (static) |
//	                 +---------------+
func TestRouteMultiLevelBacktracking2(t *testing.T) {
	e := New()
	r := e.router

	r.Add(http.MethodGet, literal_0972, handlerFunc)
	r.Add(http.MethodGet, "/a/c/d", handlerFunc)
	r.Add(http.MethodGet, literal_8765, handlerFunc)
	r.Add(http.MethodGet, literal_5382, handlerFunc)
	r.Add(http.MethodGet, "/*", handlerFunc)

	var testCases = []struct {
		name        string
		whenURL     string
		expectRoute string
		expectParam map[string]string
	}{
		{
			name:        literal_9425,
			whenURL:     literal_8765,
			expectRoute: literal_8765,
		},
		{
			name:        "route /a/x/df to /a/:b/c",
			whenURL:     "/a/x/c",
			expectRoute: literal_0972,
			expectParam: map[string]string{"b": "x"},
		},
		{
			name:        "route /a/c/f to /:e/c/f",
			whenURL:     "/a/c/f",
			expectRoute: literal_5382,
			expectParam: map[string]string{"e": "a"},
		},
		{
			name:        "route /b/c/f to /:e/c/f",
			whenURL:     "/b/c/f",
			expectRoute: literal_5382,
			expectParam: map[string]string{"e": "b"},
		},
		{
			name:        "route /b/c/c to /*",
			whenURL:     "/b/c/c",
			expectRoute: "/*",
			expectParam: map[string]string{"*": "b/c/c"},
		},
		{ // this traverses `/a/:b/c` and `/:e/c/f` branches and eventually backtracks to `/*`
			name:        "route /a/c/cf to /*",
			whenURL:     "/a/c/cf",
			expectRoute: "/*",
			expectParam: map[string]string{"*": "a/c/cf"},
		},
		{
			name:        "route /anyMatch to /*",
			whenURL:     "/anyMatch",
			expectRoute: "/*",
			expectParam: map[string]string{"*": "anyMatch"},
		},
		{
			name:        "route /anyMatch/withSlash to /*",
			whenURL:     "/anyMatch/withSlash",
			expectRoute: "/*",
			expectParam: map[string]string{"*": "anyMatch/withSlash"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			c := e.NewContext(nil, nil).(*context)

			r.Find(http.MethodGet, tc.whenURL, c)

			c.handler(c)
			assert.Equal(t, tc.expectRoute, c.Get("path"))
			for param, expectedValue := range tc.expectParam {
				assert.Equal(t, expectedValue, c.Param(param))
			}
			checkUnusedParamValues(t, c, tc.expectParam)
		})
	}
}

func TestRouterBacktrackingFromMultipleParamKinds(t *testing.T) {
	e := New()
	r := e.router

	r.Add(http.MethodGet, "/*", handlerFunc) // this can match only path that does not have slash in it
	r.Add(http.MethodGet, "/:1/second", handlerFunc)
	r.Add(http.MethodGet, "/:1/:2", handlerFunc) // this acts as match ANY for all routes that have at least one slash
	r.Add(http.MethodGet, "/:1/:2/third", handlerFunc)
	r.Add(http.MethodGet, "/:1/:2/:3/fourth", handlerFunc)
	r.Add(http.MethodGet, "/:1/:2/:3/:4/fifth", handlerFunc)

	c := e.NewContext(nil, nil).(*context)
	var testCases = []struct {
		name        string
		whenURL     string
		expectRoute string
		expectParam map[string]string
	}{
		{
			name:        "route /first to /*",
			whenURL:     "/first",
			expectRoute: "/*",
			expectParam: map[string]string{"*": "first"},
		},
		{
			name:        "route /first/second to /:1/second",
			whenURL:     "/first/second",
			expectRoute: "/:1/second",
			expectParam: map[string]string{"1": "first"},
		},
		{
			name:        "route /first/second-new to /:1/:2",
			whenURL:     "/first/second-new",
			expectRoute: "/:1/:2",
			expectParam: map[string]string{
				"1": "first",
				"2": "second-new",
			},
		},
		{ // : should match `/:1/:2` when backtracking in tree. this 1 level backtracking fails even with old implementation
			name:        "route /first/second/ to /:1/:2",
			whenURL:     "/first/second/",
			expectRoute: "/*",                                    // "/:1/:2",
			expectParam: map[string]string{"*": "first/second/"}, // map[string]string{"1": "first", "2": "second/"},
		},
		{ // : should match `/:1/:2`. same backtracking problem. when backtracking is at `/:1/:2` during backtracking this node should be match as it has executable handler
			name:        "route /first/second/third/fourth/fifth/nope to /:1/:2",
			whenURL:     "/first/second/third/fourth/fifth/nope",
			expectRoute: "/*",                                                           // "/:1/:2",
			expectParam: map[string]string{"*": "first/second/third/fourth/fifth/nope"}, // map[string]string{"1": "first", "2": "second/third/fourth/fifth/nope"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			r.Find(http.MethodGet, tc.whenURL, c)

			c.handler(c)
			assert.Equal(t, tc.expectRoute, c.Get("path"))
			for param, expectedValue := range tc.expectParam {
				assert.Equal(t, expectedValue, c.Param(param))
			}
			checkUnusedParamValues(t, c, tc.expectParam)
		})
	}
}

func TestNotFoundRouteAnyKind(t *testing.T) {
	var testCases = []struct {
		name        string
		whenURL     string
		expectRoute interface{}
		expectID    int
		expectParam map[string]string
	}{
		{
			name:        "route not existent /xx to not found handler /*",
			whenURL:     "/xx",
			expectRoute: "/*",
			expectID:    4,
			expectParam: map[string]string{"*": "xx"},
		},
		{
			name:        "route not existent /a/xx to not found handler /a/*",
			whenURL:     "/a/xx",
			expectRoute: "/a/*",
			expectID:    5,
			expectParam: map[string]string{"*": "xx"},
		},
		{
			name:        "route not existent /a/c/dxxx to not found handler /a/c/d*",
			whenURL:     "/a/c/dxxx",
			expectRoute: "/a/c/d*",
			expectID:    6,
			expectParam: map[string]string{"*": "xxx"},
		},
		{
			name:        literal_9425,
			whenURL:     literal_8765,
			expectRoute: literal_8765,
			expectID:    1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			e := New()
			r := e.router

			r.Add(http.MethodGet, "/", handlerHelper("ID", 0))
			r.Add(http.MethodGet, literal_8765, handlerHelper("ID", 1))
			r.Add(http.MethodGet, "/a/b*", handlerHelper("ID", 2))
			r.Add(http.MethodPut, "/*", handlerHelper("ID", 3))

			r.Add(RouteNotFound, "/a/c/d*", handlerHelper("ID", 6))
			r.Add(RouteNotFound, "/a/*", handlerHelper("ID", 5))
			r.Add(RouteNotFound, "/*", handlerHelper("ID", 4))

			c := e.NewContext(nil, nil).(*context)
			r.Find(http.MethodGet, tc.whenURL, c)

			c.handler(c)

			testValue, _ := c.Get("ID").(int)
			assert.Equal(t, tc.expectID, testValue)
			assert.Equal(t, tc.expectRoute, c.Path())
			for param, expectedValue := range tc.expectParam {
				assert.Equal(t, expectedValue, c.Param(param))
			}
			checkUnusedParamValues(t, c, tc.expectParam)
		})
	}
}

func TestNotFoundRouteParamKind(t *testing.T) {
	var testCases = []struct {
		name        string
		whenURL     string
		expectRoute interface{}
		expectID    int
		expectParam map[string]string
	}{
		{
			name:        "route not existent /xx to not found handler /:file",
			whenURL:     "/xx",
			expectRoute: "/:file",
			expectID:    4,
			expectParam: map[string]string{"file": "xx"},
		},
		{
			name:        "route not existent /a/xx to not found handler /a/:file",
			whenURL:     "/a/xx",
			expectRoute: "/a/:file",
			expectID:    5,
			expectParam: map[string]string{"file": "xx"},
		},
		{
			name:        "route not existent /a/c/dxxx to not found handler /a/c/d:file",
			whenURL:     "/a/c/dxxx",
			expectRoute: "/a/c/d:file",
			expectID:    6,
			expectParam: map[string]string{"file": "xxx"},
		},
		{
			name:        literal_9425,
			whenURL:     literal_8765,
			expectRoute: literal_8765,
			expectID:    1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			e := New()
			r := e.router

			r.Add(http.MethodGet, "/", handlerHelper("ID", 0))
			r.Add(http.MethodGet, literal_8765, handlerHelper("ID", 1))
			r.Add(http.MethodGet, "/a/b*", handlerHelper("ID", 2))
			r.Add(http.MethodPut, "/*", handlerHelper("ID", 3))

			r.Add(RouteNotFound, "/a/c/d:file", handlerHelper("ID", 6))
			r.Add(RouteNotFound, "/a/:file", handlerHelper("ID", 5))
			r.Add(RouteNotFound, "/:file", handlerHelper("ID", 4))

			c := e.NewContext(nil, nil).(*context)
			r.Find(http.MethodGet, tc.whenURL, c)

			c.handler(c)

			testValue, _ := c.Get("ID").(int)
			assert.Equal(t, tc.expectID, testValue)
			assert.Equal(t, tc.expectRoute, c.Path())
			for param, expectedValue := range tc.expectParam {
				assert.Equal(t, expectedValue, c.Param(param))
			}
			checkUnusedParamValues(t, c, tc.expectParam)
		})
	}
}

func TestNotFoundRouteStaticKind(t *testing.T) {
	// note: static not found handler is quite silly thing to have but we still support it
	var testCases = []struct {
		name        string
		whenURL     string
		expectRoute interface{}
		expectID    int
		expectParam map[string]string
	}{
		{
			name:        "route not existent / to not found handler /",
			whenURL:     "/",
			expectRoute: "/",
			expectID:    3,
			expectParam: map[string]string{},
		},
		{
			name:        "route /a to /a",
			whenURL:     "/a",
			expectRoute: "/a",
			expectID:    1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			e := New()
			r := e.router

			r.Add(http.MethodPut, "/", handlerHelper("ID", 0))
			r.Add(http.MethodGet, "/a", handlerHelper("ID", 1))
			r.Add(http.MethodPut, "/*", handlerHelper("ID", 2))

			r.Add(RouteNotFound, "/", handlerHelper("ID", 3))

			c := e.NewContext(nil, nil).(*context)
			r.Find(http.MethodGet, tc.whenURL, c)

			c.handler(c)

			testValue, _ := c.Get("ID").(int)
			assert.Equal(t, tc.expectID, testValue)
			assert.Equal(t, tc.expectRoute, c.Path())
			for param, expectedValue := range tc.expectParam {
				assert.Equal(t, expectedValue, c.Param(param))
			}
			checkUnusedParamValues(t, c, tc.expectParam)
		})
	}
}

func TestRouternotFoundRouteWithNodeSplitting(t *testing.T) {
	e := New()
	r := e.router

	r.Add(http.MethodGet, "/test*", handlerHelper("ID", 0))
	r.Add(RouteNotFound, "/*", handlerHelper("ID", 1))
	r.Add(RouteNotFound, "/test", handlerHelper("ID", 2))

	// Tree before:
	// 1    `/`
	// 1.1      `*` (any) ID=1
	// 1.2      `test` (static) ID=2
	// 1.2.1        `*` (any) ID=0

	// node with path `test` has routeNotFound handler from previous Add call. Now when we insert `/te/st*` into router tree
	// This means that node `test` is split into `te` and `st` nodes and new node `/st*` is inserted.
	// On that split `/test` routeNotFound handler must not be lost.
	r.Add(http.MethodGet, "/te/st*", handlerHelper("ID", 3))
	// Tree after:
	// 1    `/`
	// 1.1      `*` (any) ID=1
	// 1.2      `te` (static)
	// 1.2.1        `st` (static) ID=2
	// 1.2.1.1          `*` (any) ID=0
	// 1.2.2        `/st` (static)
	// 1.2.2.1          `*` (any) ID=3

	c := e.NewContext(nil, nil).(*context)
	r.Find(http.MethodPut, "/test", c)

	c.handler(c)

	testValue, _ := c.Get("ID").(int)
	assert.Equal(t, 2, testValue)
	assert.Equal(t, "/test", c.Path())
}

// Issue #1509
func TestRouterParamStaticConflict(t *testing.T) {
	e := New()
	r := e.router
	handler := func(c Context) error {
		c.Set("path", c.Path())
		return nil
	}

	g := e.Group("/g")
	g.GET("/skills", handler)
	g.GET("/status", handler)
	g.GET("/:name", handler)

	var testCases = []struct {
		whenURL     string
		expectRoute interface{}
		expectParam map[string]string
	}{
		{
			whenURL:     "/g/s",
			expectRoute: "/g/:name",
			expectParam: map[string]string{"name": "s"},
		},
		{
			whenURL:     "/g/status",
			expectRoute: "/g/status",
			expectParam: map[string]string{"name": ""},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.whenURL, func(t *testing.T) {
			c := e.NewContext(nil, nil).(*context)

			r.Find(http.MethodGet, tc.whenURL, c)
			err := c.handler(c)

			assert.NoError(t, err)
			assert.Equal(t, tc.expectRoute, c.Get("path"))
			for param, expectedValue := range tc.expectParam {
				assert.Equal(t, expectedValue, c.Param(param))
			}
			checkUnusedParamValues(t, c, tc.expectParam)
		})
	}
}

func TestRouterParamescapeColon(t *testing.T) {
	// to allow Google cloud API like route paths with colon in them
	// i.e. https://service.name/v1/some/resource/name:customVerb <- that `:customVerb` is not path param. It is just a string
	e := New()

	e.POST("/files/a/long/file\\:undelete", handlerFunc)
	e.POST("/multilevel\\:undelete/second\\:something", handlerFunc)
	e.POST("/mixed/:id/second\\:something", handlerFunc)
	e.POST("/v1/some/resource/name:customVerb", handlerFunc)

	var testCases = []struct {
		whenURL     string
		expectRoute interface{}
		expectParam map[string]string
		expectError string
	}{
		{
			whenURL:     "/files/a/long/file:undelete",
			expectRoute: "/files/a/long/file\\:undelete",
			expectParam: map[string]string{},
		},
		{
			whenURL:     "/multilevel:undelete/second:something",
			expectRoute: "/multilevel\\:undelete/second\\:something",
			expectParam: map[string]string{},
		},
		{
			whenURL:     "/mixed/123/second:something",
			expectRoute: "/mixed/:id/second\\:something",
			expectParam: map[string]string{"id": "123"},
		},
		{
			whenURL:     "/files/a/long/file:notMatching",
			expectRoute: nil,
			expectError: "code=404, message=Not Found",
			expectParam: nil,
		},
		{
			whenURL:     "/v1/some/resource/name:PATCH",
			expectRoute: "/v1/some/resource/name:customVerb",
			expectParam: map[string]string{"customVerb": ":PATCH"},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.whenURL, func(t *testing.T) {
			c := e.NewContext(nil, nil).(*context)

			e.router.Find(http.MethodPost, tc.whenURL, c)
			err := c.handler(c)

			assert.Equal(t, tc.expectRoute, c.Get("path"))
			if tc.expectError != "" {
				assert.EqualError(t, err, tc.expectError)
			} else {
				assert.NoError(t, err)
			}
			for param, expectedValue := range tc.expectParam {
				assert.Equal(t, expectedValue, c.Param(param))
			}
			checkUnusedParamValues(t, c, tc.expectParam)
		})
	}
}

func TestRouterMatchAny(t *testing.T) {
	e := New()
	r := e.router

	// Routes
	r.Add(http.MethodGet, "/", handlerFunc)
	r.Add(http.MethodGet, "/*", handlerFunc)
	r.Add(http.MethodGet, literal_3792, handlerFunc)

	var testCases = []struct {
		whenURL     string
		expectRoute interface{}
		expectParam map[string]string
	}{
		{
			whenURL:     "/",
			expectRoute: "/",
			expectParam: map[string]string{"*": ""},
		},
		{
			whenURL:     "/download",
			expectRoute: "/*",
			expectParam: map[string]string{"*": "download"},
		},
		{
			whenURL:     "/users/joe",
			expectRoute: literal_3792,
			expectParam: map[string]string{"*": "joe"},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.whenURL, func(t *testing.T) {
			c := e.NewContext(nil, nil).(*context)

			r.Find(http.MethodGet, tc.whenURL, c)
			err := c.handler(c)

			assert.NoError(t, err)
			assert.Equal(t, tc.expectRoute, c.Get("path"))
			for param, expectedValue := range tc.expectParam {
				assert.Equal(t, expectedValue, c.Param(param))
			}
			checkUnusedParamValues(t, c, tc.expectParam)
		})
	}
}

// NOTE: this is to document current implementation. Last added route with `*` asterisk is always the match and no
// backtracking or more precise matching is done to find more suitable match.
//
// Current behaviour might not be correct or expected.
// But this is where we are without well defined requirements/rules how (multiple) asterisks work in route
func TestRouterAnyMatchesLastAddedAnyRoute(t *testing.T) {
	e := New()
	r := e.router

	r.Add(http.MethodGet, literal_3792, handlerHelper("case", 1))
	r.Add(http.MethodGet, "/users/*/action*", handlerHelper("case", 2))

	c := e.NewContext(nil, nil).(*context)

	r.Find(http.MethodGet, "/users/xxx/action/sea", c)
	c.handler(c)
	assert.Equal(t, "/users/*/action*", c.Get("path"))
	assert.Equal(t, "xxx/action/sea", c.Param("*"))

	// if we add another route then it is the last added and so it is matched
	r.Add(http.MethodGet, "/users/*/action/search", handlerHelper("case", 3))

	r.Find(http.MethodGet, "/users/xxx/action/sea", c)
	c.handler(c)
	assert.Equal(t, "/users/*/action/search", c.Get("path"))
	assert.Equal(t, "xxx/action/sea", c.Param("*"))
}

// Issue #1739
func TestRouterMatchAnyPrefixIssue(t *testing.T) {
	e := New()
	r := e.router

	// Routes
	r.Add(http.MethodGet, "/*", func(c Context) error {
		c.Set("path", c.Path())
		return nil
	})
	r.Add(http.MethodGet, literal_3792, func(c Context) error {
		c.Set("path", c.Path())
		return nil
	})

	var testCases = []struct {
		whenURL     string
		expectRoute interface{}
		expectParam map[string]string
	}{
		{
			whenURL:     "/",
			expectRoute: "/*",
			expectParam: map[string]string{"*": ""},
		},
		{
			whenURL:     literal_7186,
			expectRoute: "/*",
			expectParam: map[string]string{"*": "users"},
		},
		{
			whenURL:     "/users/",
			expectRoute: literal_3792,
			expectParam: map[string]string{"*": ""},
		},
		{
			whenURL:     "/users_prefix",
			expectRoute: "/*",
			expectParam: map[string]string{"*": "users_prefix"},
		},
		{
			whenURL:     "/users_prefix/",
			expectRoute: "/*",
			expectParam: map[string]string{"*": "users_prefix/"},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.whenURL, func(t *testing.T) {
			c := e.NewContext(nil, nil).(*context)

			r.Find(http.MethodGet, tc.whenURL, c)
			err := c.handler(c)

			assert.NoError(t, err)
			assert.Equal(t, tc.expectRoute, c.Get("path"))
			for param, expectedValue := range tc.expectParam {
				assert.Equal(t, expectedValue, c.Param(param))
			}
			checkUnusedParamValues(t, c, tc.expectParam)
		})
	}
}

// TestRouterMatchAnySlash shall verify finding the best route
// for any routes with trailing slash requests
func TestRouterMatchAnySlash(t *testing.T) {
	e := New()
	r := e.router

	// Routes
	r.Add(http.MethodGet, literal_7186, handlerFunc)
	r.Add(http.MethodGet, literal_3792, handlerFunc)
	r.Add(http.MethodGet, "/img/*", handlerFunc)
	r.Add(http.MethodGet, literal_7419, handlerFunc)
	r.Add(http.MethodGet, literal_7503, handlerFunc)
	r.Add(http.MethodGet, "/assets/*", handlerFunc)

	var testCases = []struct {
		whenURL     string
		expectRoute interface{}
		expectParam map[string]string
		expectError error
	}{
		{
			whenURL:     "/",
			expectRoute: nil,
			expectParam: map[string]string{"*": ""},
			expectError: ErrNotFound,
		},
		{ // Test trailing slash request for simple any route (see #1526)
			whenURL:     "/users/",
			expectRoute: literal_3792,
			expectParam: map[string]string{"*": ""},
		},
		{
			whenURL:     "/users/joe",
			expectRoute: literal_3792,
			expectParam: map[string]string{"*": "joe"},
		},
		// Test trailing slash request for nested any route (see #1526)
		{
			whenURL:     literal_7419,
			expectRoute: literal_7419,
			expectParam: map[string]string{"*": ""},
		},
		{
			whenURL:     "/img/load/",
			expectRoute: literal_7503,
			expectParam: map[string]string{"*": ""},
		},
		{
			whenURL:     "/img/load/ben",
			expectRoute: literal_7503,
			expectParam: map[string]string{"*": "ben"},
		},
		// Test /assets/* any route
		{ // ... without trailing slash must not match
			whenURL:     "/assets",
			expectRoute: nil,
			expectParam: map[string]string{"*": ""},
			expectError: ErrNotFound,
		},

		{ // ... with trailing slash must match
			whenURL:     "/assets/",
			expectRoute: "/assets/*",
			expectParam: map[string]string{"*": ""},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.whenURL, func(t *testing.T) {
			c := e.NewContext(nil, nil).(*context)

			r.Find(http.MethodGet, tc.whenURL, c)
			err := c.handler(c)

			if tc.expectError != nil {
				assert.Equal(t, tc.expectError, err)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tc.expectRoute, c.Get("path"))
			for param, expectedValue := range tc.expectParam {
				assert.Equal(t, expectedValue, c.Param(param))
			}
			checkUnusedParamValues(t, c, tc.expectParam)
		})
	}
}

func TestRouterMatchAnyMultiLevel(t *testing.T) {
	e := New()
	r := e.router

	// Routes
	r.Add(http.MethodGet, literal_7302, handlerFunc)
	r.Add(http.MethodGet, literal_5480, handlerFunc)
	r.Add(http.MethodGet, "/api/users/*", handlerFunc)
	r.Add(http.MethodGet, literal_9012, handlerFunc)
	r.Add(http.MethodGet, "/other/*", handlerFunc)
	r.Add(http.MethodGet, "/*", handlerFunc)

	var testCases = []struct {
		whenURL     string
		expectRoute interface{}
		expectParam map[string]string
		expectError error
	}{
		{
			whenURL:     literal_7302,
			expectRoute: literal_7302,
			expectParam: map[string]string{"*": ""},
		},
		{
			whenURL:     literal_5480,
			expectRoute: literal_5480,
			expectParam: map[string]string{"*": ""},
		},
		{
			whenURL:     "/api/users/joe",
			expectRoute: "/api/users/*",
			expectParam: map[string]string{"*": "joe"},
		},
		{
			whenURL:     "/api/nousers/joe",
			expectRoute: literal_9012,
			expectParam: map[string]string{"*": "nousers/joe"},
		},
		{
			whenURL:     "/api/none",
			expectRoute: literal_9012,
			expectParam: map[string]string{"*": "none"},
		},
		{
			whenURL:     "/api/none",
			expectRoute: literal_9012,
			expectParam: map[string]string{"*": "none"},
		},
		{
			whenURL:     "/noapi/users/jim",
			expectRoute: "/*",
			expectParam: map[string]string{"*": "noapi/users/jim"},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.whenURL, func(t *testing.T) {
			c := e.NewContext(nil, nil).(*context)

			r.Find(http.MethodGet, tc.whenURL, c)
			err := c.handler(c)

			if tc.expectError != nil {
				assert.Equal(t, tc.expectError, err)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tc.expectRoute, c.Get("path"))
			for param, expectedValue := range tc.expectParam {
				assert.Equal(t, expectedValue, c.Param(param))
			}
			checkUnusedParamValues(t, c, tc.expectParam)
		})
	}
}
func TestRouterMatchAnyMultiLevelWithPost(t *testing.T) {
	e := New()
	r := e.router

	// Routes
	e.POST(literal_3906, handlerFunc)
	e.POST("/api/auth/forgotPassword", handlerFunc)
	e.Any(literal_9012, handlerFunc)
	e.Any("/*", handlerFunc)

	var testCases = []struct {
		whenMethod  string
		whenURL     string
		expectRoute interface{}
		expectParam map[string]string
		expectError error
	}{
		{ // POST /api/auth/login shall choose login method
			whenURL:     literal_3906,
			whenMethod:  http.MethodPost,
			expectRoute: literal_3906,
			expectParam: map[string]string{"*": ""},
		},
		{ // POST /api/auth/logout shall choose nearest any route
			whenURL:     "/api/auth/logout",
			whenMethod:  http.MethodPost,
			expectRoute: literal_9012,
			expectParam: map[string]string{"*": "auth/logout"},
		},
		{ // POST to /api/other/test shall choose nearest any route
			whenURL:     "/api/other/test",
			whenMethod:  http.MethodPost,
			expectRoute: literal_9012,
			expectParam: map[string]string{"*": "other/test"},
		},
		{ // GET to /api/other/test shall choose nearest any route
			whenURL:     "/api/other/test",
			whenMethod:  http.MethodGet,
			expectRoute: literal_9012,
			expectParam: map[string]string{"*": "other/test"},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.whenURL, func(t *testing.T) {
			c := e.NewContext(nil, nil).(*context)

			method := http.MethodGet
			if tc.whenMethod != "" {
				method = tc.whenMethod
			}
			r.Find(method, tc.whenURL, c)
			err := c.handler(c)

			if tc.expectError != nil {
				assert.Equal(t, tc.expectError, err)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tc.expectRoute, c.Get("path"))
			for param, expectedValue := range tc.expectParam {
				assert.Equal(t, expectedValue, c.Param(param))
			}
			checkUnusedParamValues(t, c, tc.expectParam)
		})
	}
}

func TestRouterMicroParam(t *testing.T) {
	e := New()
	r := e.router
	r.Add(http.MethodGet, "/:a/:b/:c", func(c Context) error {
		return nil
	})
	c := e.NewContext(nil, nil).(*context)
	r.Find(http.MethodGet, "/1/2/3", c)
	assert.Equal(t, "1", c.Param("a"))
	assert.Equal(t, "2", c.Param("b"))
	assert.Equal(t, "3", c.Param("c"))
}

func TestRouterMixParamMatchAny(t *testing.T) {
	e := New()
	r := e.router

	// Route
	r.Add(http.MethodGet, "/users/:id/*", func(c Context) error {
		return nil
	})
	c := e.NewContext(nil, nil).(*context)

	r.Find(http.MethodGet, "/users/joe/comments", c)
	c.handler(c)
	assert.Equal(t, "joe", c.Param("id"))
}

func TestRouterMultiRoute(t *testing.T) {
	e := New()
	r := e.router

	// Routes
	r.Add(http.MethodGet, literal_7186, handlerFunc)
	r.Add(http.MethodGet, literal_8032, handlerFunc)

	var testCases = []struct {
		whenMethod  string
		whenURL     string
		expectRoute interface{}
		expectParam map[string]string
		expectError error
	}{
		{
			whenURL:     literal_7186,
			expectRoute: literal_7186,
			expectParam: map[string]string{"*": ""},
		},
		{
			whenURL:     literal_0618,
			expectRoute: literal_8032,
			expectParam: map[string]string{"id": "1"},
		},
		{
			whenURL:     "/user",
			expectRoute: nil,
			expectParam: map[string]string{"*": ""},
			expectError: ErrNotFound,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.whenURL, func(t *testing.T) {
			c := e.NewContext(nil, nil).(*context)

			method := http.MethodGet
			if tc.whenMethod != "" {
				method = tc.whenMethod
			}
			r.Find(method, tc.whenURL, c)
			err := c.handler(c)

			if tc.expectError != nil {
				assert.Equal(t, tc.expectError, err)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tc.expectRoute, c.Get("path"))
			for param, expectedValue := range tc.expectParam {
				assert.Equal(t, expectedValue, c.Param(param))
			}
			checkUnusedParamValues(t, c, tc.expectParam)
		})
	}
}

func TestRouterPriority(t *testing.T) {
	e := New()
	r := e.router

	// Routes
	r.Add(http.MethodGet, literal_7186, handlerFunc)
	r.Add(http.MethodGet, literal_7916, handlerFunc)
	r.Add(http.MethodGet, literal_8032, handlerFunc)
	r.Add(http.MethodGet, literal_6139, handlerFunc)
	r.Add(http.MethodGet, "/users/:id/files", handlerFunc)
	r.Add(http.MethodGet, literal_3245, handlerFunc)
	r.Add(http.MethodGet, literal_3792, handlerFunc)
	r.Add(http.MethodGet, "/users/new/*", handlerFunc)
	r.Add(http.MethodGet, "/*", handlerFunc)

	var testCases = []struct {
		whenMethod  string
		whenURL     string
		expectRoute interface{}
		expectParam map[string]string
		expectError error
	}{
		{
			whenURL:     literal_7186,
			expectRoute: literal_7186,
		},
		{
			whenURL:     literal_7916,
			expectRoute: literal_7916,
		},
		{
			whenURL:     literal_0618,
			expectRoute: literal_8032,
			expectParam: map[string]string{"id": "1"},
		},
		{
			whenURL:     literal_6139,
			expectRoute: literal_6139,
		},
		{
			whenURL:     "/users/1/files",
			expectRoute: "/users/:id/files",
			expectParam: map[string]string{"id": "1"},
		},
		{
			whenURL:     literal_7916,
			expectRoute: literal_7916,
		},
		{
			whenURL:     "/users/news",
			expectRoute: literal_8032,
			expectParam: map[string]string{"id": "news"},
		},
		{
			whenURL:     literal_3245,
			expectRoute: literal_3245,
		},
		{
			whenURL:     "/users/joe/books",
			expectRoute: literal_3792,
			expectParam: map[string]string{"*": "joe/books"},
		},
		{
			whenURL:     "/users/new/someone",
			expectRoute: "/users/new/*",
			expectParam: map[string]string{"*": "someone"},
		},
		{
			whenURL:     "/users/dew/someone",
			expectRoute: literal_3792,
			expectParam: map[string]string{"*": "dew/someone"},
		},
		{ // Route > /users/* should be matched although /users/dew exists
			whenURL:     "/users/notexists/someone",
			expectRoute: literal_3792,
			expectParam: map[string]string{"*": "notexists/someone"},
		},
		{
			whenURL:     "/nousers",
			expectRoute: "/*",
			expectParam: map[string]string{"*": "nousers"},
		},
		{
			whenURL:     "/nousers/new",
			expectRoute: "/*",
			expectParam: map[string]string{"*": "nousers/new"},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.whenURL, func(t *testing.T) {
			c := e.NewContext(nil, nil).(*context)

			method := http.MethodGet
			if tc.whenMethod != "" {
				method = tc.whenMethod
			}
			r.Find(method, tc.whenURL, c)
			err := c.handler(c)

			if tc.expectError != nil {
				assert.Equal(t, tc.expectError, err)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tc.expectRoute, c.Get("path"))
			for param, expectedValue := range tc.expectParam {
				assert.Equal(t, expectedValue, c.Param(param))
			}
			checkUnusedParamValues(t, c, tc.expectParam)
		})
	}
}

func TestRouterIssue1348(t *testing.T) {
	e := New()
	r := e.router

	r.Add(http.MethodGet, "/:lang/", func(c Context) error {
		return nil
	})
	r.Add(http.MethodGet, "/:lang/dupa", func(c Context) error {
		return nil
	})
}

// Issue #372
func TestRouterPriorityNotFound(t *testing.T) {
	e := New()
	r := e.router

	// Add
	r.Add(http.MethodGet, literal_0341, handlerFunc)
	r.Add(http.MethodGet, literal_1489, handlerFunc)

	var testCases = []struct {
		whenMethod  string
		whenURL     string
		expectRoute interface{}
		expectParam map[string]string
		expectError error
	}{
		{
			whenURL:     literal_0341,
			expectRoute: literal_0341,
		},
		{
			whenURL:     literal_1489,
			expectRoute: literal_1489,
		},
		{
			whenURL:     "/abc/def",
			expectRoute: nil,
			expectError: ErrNotFound,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.whenURL, func(t *testing.T) {
			c := e.NewContext(nil, nil).(*context)

			method := http.MethodGet
			if tc.whenMethod != "" {
				method = tc.whenMethod
			}
			r.Find(method, tc.whenURL, c)
			err := c.handler(c)

			if tc.expectError != nil {
				assert.Equal(t, tc.expectError, err)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tc.expectRoute, c.Get("path"))
			for param, expectedValue := range tc.expectParam {
				assert.Equal(t, expectedValue, c.Param(param))
			}
			checkUnusedParamValues(t, c, tc.expectParam)
		})
	}
}

func TestRouterParamNames(t *testing.T) {
	e := New()
	r := e.router

	// Routes
	r.Add(http.MethodGet, literal_7186, handlerFunc)
	r.Add(http.MethodGet, literal_8032, handlerFunc)
	r.Add(http.MethodGet, literal_8670, handlerFunc)

	var testCases = []struct {
		whenMethod  string
		whenURL     string
		expectRoute interface{}
		expectParam map[string]string
		expectError error
	}{
		{
			whenURL:     literal_7186,
			expectRoute: literal_7186,
		},
		{
			whenURL:     literal_0618,
			expectRoute: literal_8032,
			expectParam: map[string]string{"id": "1"},
		},
		{
			whenURL:     "/users/1/files/1",
			expectRoute: literal_8670,
			expectParam: map[string]string{
				"uid": "1",
				"fid": "1",
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.whenURL, func(t *testing.T) {
			c := e.NewContext(nil, nil).(*context)

			method := http.MethodGet
			if tc.whenMethod != "" {
				method = tc.whenMethod
			}
			r.Find(method, tc.whenURL, c)
			err := c.handler(c)

			if tc.expectError != nil {
				assert.Equal(t, tc.expectError, err)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tc.expectRoute, c.Get("path"))
			for param, expectedValue := range tc.expectParam {
				assert.Equal(t, expectedValue, c.Param(param))
			}
			checkUnusedParamValues(t, c, tc.expectParam)
		})
	}
}

// Issue #623 and #1406
func TestRouterStaticDynamicConflict(t *testing.T) {
	e := New()
	r := e.router

	r.Add(http.MethodGet, literal_1850, handlerHelper("a", 1))
	r.Add(http.MethodGet, literal_3028, handlerHelper("b", 2))
	r.Add(http.MethodGet, literal_7916, handlerHelper("d", 4))
	r.Add(http.MethodGet, "/users/:name", handlerHelper("e", 5))
	r.Add(http.MethodGet, literal_6954, handlerHelper("c", 3))
	r.Add(http.MethodGet, "/", handlerHelper("f", 6))

	var testCases = []struct {
		whenMethod  string
		whenURL     string
		expectRoute interface{}
		expectParam map[string]string
		expectError error
	}{
		{
			whenURL:     literal_1850,
			expectRoute: literal_1850,
			expectParam: map[string]string{"*": ""},
		},
		{
			whenURL:     "/dictionary/skillsnot",
			expectRoute: literal_3028,
			expectParam: map[string]string{"name": "skillsnot"},
		},
		{
			whenURL:     "/dictionary/type",
			expectRoute: literal_3028,
			expectParam: map[string]string{"name": "type"},
		},
		{
			whenURL:     literal_6954,
			expectRoute: literal_6954,
		},
		{
			whenURL:     literal_7916,
			expectRoute: literal_7916,
		},
		{
			whenURL:     "/users/new2",
			expectRoute: "/users/:name",
			expectParam: map[string]string{"name": "new2"},
		},
		{
			whenURL:     "/",
			expectRoute: "/",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.whenURL, func(t *testing.T) {
			c := e.NewContext(nil, nil).(*context)

			method := http.MethodGet
			if tc.whenMethod != "" {
				method = tc.whenMethod
			}
			r.Find(method, tc.whenURL, c)
			err := c.handler(c)

			if tc.expectError != nil {
				assert.Equal(t, tc.expectError, err)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tc.expectRoute, c.Get("path"))
			for param, expectedValue := range tc.expectParam {
				assert.Equal(t, expectedValue, c.Param(param))
			}
			checkUnusedParamValues(t, c, tc.expectParam)
		})
	}
}

// Issue #1348
func TestRouterParamBacktraceNotFound(t *testing.T) {
	e := New()
	r := e.router

	// Add
	r.Add(http.MethodGet, "/:param1", handlerFunc)
	r.Add(http.MethodGet, "/:param1/foo", handlerFunc)
	r.Add(http.MethodGet, "/:param1/bar", handlerFunc)
	r.Add(http.MethodGet, "/:param1/bar/:param2", handlerFunc)

	var testCases = []struct {
		name        string
		whenMethod  string
		whenURL     string
		expectRoute interface{}
		expectParam map[string]string
		expectError error
	}{
		{
			name:        "route /a to /:param1",
			whenURL:     "/a",
			expectRoute: "/:param1",
			expectParam: map[string]string{"param1": "a"},
		},
		{
			name:        "route /a/foo to /:param1/foo",
			whenURL:     literal_0341,
			expectRoute: "/:param1/foo",
			expectParam: map[string]string{"param1": "a"},
		},
		{
			name:        "route /a/bar to /:param1/bar",
			whenURL:     literal_1489,
			expectRoute: "/:param1/bar",
			expectParam: map[string]string{"param1": "a"},
		},
		{
			name:        "route /a/bar/b to /:param1/bar/:param2",
			whenURL:     "/a/bar/b",
			expectRoute: "/:param1/bar/:param2",
			expectParam: map[string]string{
				"param1": "a",
				"param2": "b",
			},
		},
		{
			name:        "route /a/bbbbb should return 404",
			whenURL:     "/a/bbbbb",
			expectRoute: nil,
			expectError: ErrNotFound,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			c := e.NewContext(nil, nil).(*context)

			method := http.MethodGet
			if tc.whenMethod != "" {
				method = tc.whenMethod
			}
			r.Find(method, tc.whenURL, c)
			err := c.handler(c)

			if tc.expectError != nil {
				assert.Equal(t, tc.expectError, err)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tc.expectRoute, c.Get("path"))
			for param, expectedValue := range tc.expectParam {
				assert.Equal(t, expectedValue, c.Param(param))
			}
			checkUnusedParamValues(t, c, tc.expectParam)
		})
	}
}

func testRouterAPI(t *testing.T, api []*Route) {
	e := New()
	r := e.router

	for _, route := range api {
		r.Add(route.Method, route.Path, func(c Context) error {
			return nil
		})
	}
	c := e.NewContext(nil, nil).(*context)
	for _, route := range api {
		t.Run(route.Path, func(t *testing.T) {
			r.Find(route.Method, route.Path, c)
			tokens := strings.Split(route.Path[1:], "/")
			for _, token := range tokens {
				if token[0] == ':' {
					assert.Equal(t, c.Param(token[1:]), token)
				}
			}
		})
	}
}

func TestRouterGitHubAPI(t *testing.T) {
	testRouterAPI(t, gitHubAPI)
}

// Issue #729
func TestRouterParamAlias(t *testing.T) {
	api := []*Route{
		{http.MethodGet, "/users/:userID/following", ""},
		{http.MethodGet, "/users/:userID/followedBy", ""},
		{http.MethodGet, "/users/:userID/follow", ""},
	}
	testRouterAPI(t, api)
}

// Issue #1052
func TestRouterParamOrdering(t *testing.T) {
	api := []*Route{
		{http.MethodGet, literal_3672, ""},
		{http.MethodGet, literal_8370, ""},
		{http.MethodGet, literal_5390, ""},
	}
	testRouterAPI(t, api)
	api2 := []*Route{
		{http.MethodGet, literal_8370, ""},
		{http.MethodGet, literal_5390, ""},
		{http.MethodGet, literal_3672, ""},
	}
	testRouterAPI(t, api2)
	api3 := []*Route{
		{http.MethodGet, literal_3672, ""},
		{http.MethodGet, literal_5390, ""},
		{http.MethodGet, literal_8370, ""},
	}
	testRouterAPI(t, api3)
}

// Issue #1139
func TestRouterMixedParams(t *testing.T) {
	api := []*Route{
		{http.MethodGet, "/teacher/:tid/room/suggestions", ""},
		{http.MethodGet, "/teacher/:id", ""},
	}
	testRouterAPI(t, api)
	api2 := []*Route{
		{http.MethodGet, "/teacher/:id", ""},
		{http.MethodGet, "/teacher/:tid/room/suggestions", ""},
	}
	testRouterAPI(t, api2)
}

// Issue #1466
func TestRouterParam1466(t *testing.T) {
	e := New()
	r := e.router

	r.Add(http.MethodPost, "/users/signup", handlerFunc)
	r.Add(http.MethodPost, "/users/signup/bulk", handlerFunc)
	r.Add(http.MethodPost, "/users/survey", handlerFunc)
	r.Add(http.MethodGet, literal_0148, handlerFunc)
	r.Add(http.MethodGet, "/interests/:name/users", handlerFunc)
	r.Add(http.MethodGet, "/skills/:name/users", handlerFunc)
	// Additional routes for Issue 1479
	r.Add(http.MethodGet, literal_9267, handlerFunc)
	r.Add(http.MethodGet, literal_3697, handlerFunc)
	r.Add(http.MethodGet, literal_7098, handlerFunc)

	var testCases = []struct {
		whenURL     string
		expectRoute interface{}
		expectParam map[string]string
	}{
		{
			whenURL:     "/users/ajitem",
			expectRoute: literal_0148,
			expectParam: map[string]string{"username": "ajitem"},
		},
		{
			whenURL:     "/users/sharewithme",
			expectRoute: literal_0148,
			expectParam: map[string]string{"username": "sharewithme"},
		},
		{ // route `/users/signup` is registered for POST. so param route `/users/:username` (lesser priority) is matched as it has GET handler
			whenURL:     "/users/signup",
			expectRoute: literal_0148,
			expectParam: map[string]string{"username": "signup"},
		},
		// Additional assertions for #1479
		{
			whenURL:     "/users/sharewithme/likes/projects/ids",
			expectRoute: literal_9267,
			expectParam: map[string]string{"username": "sharewithme"},
		},
		{
			whenURL:     "/users/ajitem/likes/projects/ids",
			expectRoute: literal_9267,
			expectParam: map[string]string{"username": "ajitem"},
		},
		{
			whenURL:     "/users/sharewithme/profile",
			expectRoute: literal_3697,
			expectParam: map[string]string{"username": "sharewithme"},
		},
		{
			whenURL:     "/users/ajitem/profile",
			expectRoute: literal_3697,
			expectParam: map[string]string{"username": "ajitem"},
		},
		{
			whenURL:     "/users/sharewithme/uploads/self",
			expectRoute: literal_7098,
			expectParam: map[string]string{
				"username": "sharewithme",
				"type":     "self",
			},
		},
		{
			whenURL:     "/users/ajitem/uploads/self",
			expectRoute: literal_7098,
			expectParam: map[string]string{
				"username": "ajitem",
				"type":     "self",
			},
		},
		{
			whenURL:     "/users/tree/free",
			expectRoute: nil, // not found
			expectParam: map[string]string{"id": ""},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.whenURL, func(t *testing.T) {
			c := e.NewContext(nil, nil).(*context)

			r.Find(http.MethodGet, tc.whenURL, c)
			c.handler(c)
			assert.Equal(t, tc.expectRoute, c.Get("path"))
			for param, expectedValue := range tc.expectParam {
				assert.Equal(t, expectedValue, c.Param(param))
			}
			checkUnusedParamValues(t, c, tc.expectParam)
		})
	}
}

// Issue #1655
func TestRouterFindNotPanicOrLoopsWhenContextSetParamValuesIsCalledWithLessValuesThanEchoMaxParam(t *testing.T) {
	e := New()
	r := e.router

	v0 := e.Group("/:version")
	v0.GET("/admin", func(c Context) error {
		c.SetParamNames("version")
		c.SetParamValues("v1")
		return nil
	})

	v0.GET("/images/view/:id", handlerHelper("iv", 1))
	v0.GET("/images/:id", handlerHelper("i", 1))
	v0.GET("/view/*", handlerHelper("v", 1))

	//If this API is called before the next two one panic the other loops ( of course without my fix ;) )
	c := e.NewContext(nil, nil)
	r.Find(http.MethodGet, "/v1/admin", c)
	c.Handler()(c)
	assert.Equal(t, "v1", c.Param("version"))

	//panic
	c = e.NewContext(nil, nil)
	r.Find(http.MethodGet, "/v1/view/same-data", c)
	c.Handler()(c)
	assert.Equal(t, "same-data", c.Param("*"))
	assert.Equal(t, 1, c.Get("v"))

	//looping
	c = e.NewContext(nil, nil)
	r.Find(http.MethodGet, "/v1/images/view", c)
	c.Handler()(c)
	assert.Equal(t, "view", c.Param("id"))
	assert.Equal(t, 1, c.Get("i"))
}

// Issue #1653
func TestRouterPanicWhenParamNoRootOnlyChildsFailsFind(t *testing.T) {
	e := New()
	r := e.router

	r.Add(http.MethodGet, literal_0948, handlerFunc)
	r.Add(http.MethodGet, "/users/:id/edit", handlerFunc)
	r.Add(http.MethodGet, "/users/:id/active", handlerFunc)

	var testCases = []struct {
		whenURL      string
		expectRoute  interface{}
		expectParam  map[string]string
		expectStatus int
	}{
		{
			whenURL:     "/users/alice/edit",
			expectRoute: "/users/:id/edit",
			expectParam: map[string]string{"id": "alice"},
		},
		{
			whenURL:     "/users/bob/active",
			expectRoute: "/users/:id/active",
			expectParam: map[string]string{"id": "bob"},
		},
		{
			whenURL:     literal_0948,
			expectRoute: literal_0948,
			expectParam: nil,
		},
		//This panic before the fix for Issue #1653
		{
			whenURL:      "/users/createNotFound",
			expectStatus: http.StatusNotFound,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.whenURL, func(t *testing.T) {
			c := e.NewContext(nil, nil).(*context)

			r.Find(http.MethodGet, tc.whenURL, c)
			err := c.handler(c)

			if tc.expectStatus != 0 {
				assert.Error(t, err)
				he := err.(*HTTPError)
				assert.Equal(t, tc.expectStatus, he.Code)
			}
			assert.Equal(t, tc.expectRoute, c.Get("path"))
			for param, expectedValue := range tc.expectParam {
				assert.Equal(t, expectedValue, c.Param(param))
			}
			checkUnusedParamValues(t, c, tc.expectParam)
		})
	}
}

// Issue #1726
func TestRouterDifferentParamsInPath(t *testing.T) {
	e := New()
	r := e.router
	r.Add(http.MethodPut, "/*", func(Context) error {
		return nil
	})
	r.Add(http.MethodPut, "/users/:vid/files/:gid", func(Context) error {
		return nil
	})
	r.Add(http.MethodGet, literal_8670, func(Context) error {
		return nil
	})
	c := e.NewContext(nil, nil).(*context)

	r.Find(http.MethodGet, "/users/1/files/2", c)
	assert.Equal(t, "1", c.Param("uid"))
	assert.Equal(t, "2", c.Param("fid"))

	r.Find(http.MethodGet, "/users/1/shouldBacktrackToFirstAnyRouteAnd405", c)
	assert.Equal(t, "/*", c.Path())

	r.Find(http.MethodPut, "/users/3/files/4", c)
	assert.Equal(t, "3", c.Param("vid"))
	assert.Equal(t, "4", c.Param("gid"))
}

func TestRouterHandleMethodOptions(t *testing.T) {
	e := New()
	r := e.router

	r.Add(http.MethodGet, literal_7186, handlerFunc)
	r.Add(http.MethodPost, literal_7186, handlerFunc)
	r.Add(http.MethodPut, literal_8032, handlerFunc)
	r.Add(http.MethodGet, literal_8032, handlerFunc)

	var testCases = []struct {
		name              string
		whenMethod        string
		whenURL           string
		expectAllowHeader string
		expectStatus      int
	}{
		{
			name:              "allows GET and POST handlers",
			whenMethod:        http.MethodOptions,
			whenURL:           literal_7186,
			expectAllowHeader: "OPTIONS, GET, POST",
			expectStatus:      http.StatusNoContent,
		},
		{
			name:              "allows GET and PUT handlers",
			whenMethod:        http.MethodOptions,
			whenURL:           literal_0618,
			expectAllowHeader: "OPTIONS, GET, PUT",
			expectStatus:      http.StatusNoContent,
		},
		{
			name:              "GET does not have allows header",
			whenMethod:        http.MethodGet,
			whenURL:           literal_7186,
			expectAllowHeader: "",
			expectStatus:      http.StatusOK,
		},
		{
			name:              "path with no handlers does not set Allows header",
			whenMethod:        http.MethodOptions,
			whenURL:           "/notFound",
			expectAllowHeader: "",
			expectStatus:      http.StatusNotFound,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			req := httptest.NewRequest(tc.whenMethod, tc.whenURL, nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec).(*context)

			r.Find(tc.whenMethod, tc.whenURL, c)
			err := c.handler(c)

			if tc.expectStatus >= 400 {
				assert.Error(t, err)
				he := err.(*HTTPError)
				assert.Equal(t, tc.expectStatus, he.Code)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expectStatus, rec.Code)
			}
			assert.Equal(t, tc.expectAllowHeader, c.Response().Header().Get(HeaderAllow))
		})
	}
}

func TestRouteraddEmptyPathToSlashReverse(t *testing.T) {
	e := New()
	r := e.router
	r.add(http.MethodGet, "", "empty", handlerFunc) // emtpy path is normalized to `/`

	assert.Equal(t, "/", r.Reverse("empty"))
}

func TestRouterReverseNotFound(t *testing.T) {
	e := New()
	r := e.router
	r.add(http.MethodGet, "", "empty", handlerFunc)

	assert.Equal(t, "", r.Reverse("not-existing"))
}

func TestRouterReverse(t *testing.T) {
	e := New()
	r := e.router
	dummyHandler := func(Context) error { return nil }

	r.add(http.MethodGet, literal_5831, literal_5831, dummyHandler)
	r.add(http.MethodGet, literal_6832, literal_6832, dummyHandler)
	r.add(http.MethodGet, literal_3416, literal_3416, dummyHandler)
	r.add(http.MethodGet, literal_4207, literal_4207, dummyHandler)
	r.add(http.MethodGet, literal_9351, literal_9351, dummyHandler)

	assert.Equal(t, literal_5831, r.Reverse(literal_5831))
	assert.Equal(t, literal_5831, r.Reverse(literal_5831, "missing param"))
	assert.Equal(t, literal_6832, r.Reverse(literal_6832))
	assert.Equal(t, "/static/foo.txt", r.Reverse(literal_6832, "foo.txt"))

	assert.Equal(t, literal_3416, r.Reverse(literal_3416))
	assert.Equal(t, "/params/one", r.Reverse(literal_3416, "one"))
	assert.Equal(t, literal_4207, r.Reverse(literal_4207))
	assert.Equal(t, "/params/one/bar/:qux", r.Reverse(literal_4207, "one"))
	assert.Equal(t, "/params/one/bar/two", r.Reverse(literal_4207, "one", "two"))
	assert.Equal(t, "/params/one/bar/two/three", r.Reverse(literal_9351, "one", "two", "three"))
}

func TestRouterAllowHeaderForAnyOtherMethodType(t *testing.T) {
	e := New()
	r := e.router

	r.Add(http.MethodGet, literal_7186, handlerFunc)
	r.Add("COPY", literal_7186, handlerFunc)
	r.Add("LOCK", literal_7186, handlerFunc)

	req := httptest.NewRequest("TEST", literal_7186, nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec).(*context)

	r.Find("TEST", literal_7186, c)
	err := c.handler(c)

	assert.EqualError(t, err, "code=405, message=Method Not Allowed")
	assert.ElementsMatch(t, []string{"COPY", "GET", "LOCK", "OPTIONS"}, strings.Split(c.Response().Header().Get(HeaderAllow), ", "))
}

func benchmarkRouterRoutes(b *testing.B, routes []*Route, routesToFind []*Route) {
	e := New()
	r := e.router
	b.ReportAllocs()

	// Add routes
	for _, route := range routes {
		r.Add(route.Method, route.Path, func(c Context) error {
			return nil
		})
	}

	// Routes adding are performed just once, so it doesn't make sense to see that in the benchmark
	b.ResetTimer()

	// Find routes
	for i := 0; i < b.N; i++ {
		for _, route := range routesToFind {
			c := e.pool.Get().(*context)
			r.Find(route.Method, route.Path, c)
			e.pool.Put(c)
		}
	}
}

func BenchmarkRouterStaticRoutes(b *testing.B) {
	benchmarkRouterRoutes(b, staticRoutes, staticRoutes)
}

func BenchmarkRouterStaticRoutesMisses(b *testing.B) {
	benchmarkRouterRoutes(b, staticRoutes, missesAPI)
}

func BenchmarkRouterGitHubAPI(b *testing.B) {
	benchmarkRouterRoutes(b, gitHubAPI, gitHubAPI)
}

func BenchmarkRouterGitHubAPIMisses(b *testing.B) {
	benchmarkRouterRoutes(b, gitHubAPI, missesAPI)
}

func BenchmarkRouterParseAPI(b *testing.B) {
	benchmarkRouterRoutes(b, parseAPI, parseAPI)
}

func BenchmarkRouterParseAPIMisses(b *testing.B) {
	benchmarkRouterRoutes(b, parseAPI, missesAPI)
}

func BenchmarkRouterGooglePlusAPI(b *testing.B) {
	benchmarkRouterRoutes(b, googlePlusAPI, googlePlusAPI)
}

func BenchmarkRouterGooglePlusAPIMisses(b *testing.B) {
	benchmarkRouterRoutes(b, googlePlusAPI, missesAPI)
}

func BenchmarkRouterParamsAndAnyAPI(b *testing.B) {
	benchmarkRouterRoutes(b, paramAndAnyAPI, paramAndAnyAPIToFind)
}

const literal_4602 = "/authorizations/:id"

const literal_7620 = "/notifications/threads/:id/subscription"

const literal_5690 = "/user/starred/:owner/:repo"

const literal_3975 = "/repos/:owner/:repo/subscription"

const literal_3691 = "/user/subscriptions/:owner/:repo"

const literal_3801 = "/gists/:id"

const literal_2893 = "/repos/:owner/:repo/git/refs/*ref"

const literal_7382 = "/repos/:owner/:repo/issues/comments/:id"

const literal_5602 = "/repos/:owner/:repo/labels/:name"

const literal_8274 = "/repos/:owner/:repo/issues/:number/labels"

const literal_4896 = "/repos/:owner/:repo/milestones/:number"

const literal_6792 = "/orgs/:org/public_members/:user"

const literal_5497 = "/teams/:id"

const literal_3908 = "/teams/:id/members/:user"

const literal_5761 = "/teams/:id/repos/:owner/:repo"

const literal_1928 = "/repos/:owner/:repo/pulls/comments/:number"

const literal_3125 = "/repos/:owner/:repo/collaborators/:user"

const literal_1385 = "/repos/:owner/:repo/comments/:id"

const literal_0735 = "/repos/:owner/:repo/keys/:id"

const literal_1408 = "/repos/:owner/:repo/hooks/:id"

const literal_3521 = "/repos/:owner/:repo/releases/:id"

const literal_7186 = "/users"

const literal_3017 = "/user/emails"

const literal_5296 = "/user/following/:user"

const literal_2071 = "/user/keys/:id"

const literal_3671 = "/1/classes/:className/:objectId"

const literal_2043 = "/1/users/:objectId"

const literal_9741 = "/1/roles/:objectId"

const literal_7512 = "/1/installations/:objectId"

const literal_8639 = "/root/:first/foo/*"

const literal_2618 = "/root/:first/:second/*"

const literal_4350 = "/root/:first/bar/:second/*"

const literal_2693 = "/root/:first/qux/:second/:third/:fourth"

const literal_3596 = "/root/:first/qux/:second/:third/:fourth/*"

const literal_8675 = "/root/one/foo/after/the/asterisk"

const literal_3472 = "/root/one/foo/path/after/the/asterisk"

const literal_7230 = "/root/one/two/path/after/the/asterisk"

const literal_9437 = "/root/one/bar/two/after/the/asterisk"

const literal_3276 = "/root/one/qux/two/three/four"

const literal_3476 = "/root/one/qux/two/three/four/after/the/asterisk"

const literal_4371 = "/missOne"

const literal_1794 = "/miss/two"

const literal_7945 = "/miss/three/levels"

const literal_1064 = "/miss/four/levels/nooo"

const literal_8032 = "/users/:id"

const literal_0618 = "/users/1"

const literal_8670 = "/users/:uid/files/:fid"

const literal_9425 = "route /a/c/df to /a/c/df"

const literal_0972 = "/a/:b/c"

const literal_5382 = "/:e/c/f"

const literal_3792 = "/users/*"

const literal_7419 = "/img/load"

const literal_7503 = "/img/load/*"

const literal_7302 = "/api/users/jack"

const literal_5480 = "/api/users/jill"

const literal_9012 = "/api/*"

const literal_3906 = "/api/auth/login"

const literal_7916 = "/users/new"

const literal_6139 = "/users/dew"

const literal_3245 = "/users/newsee"

const literal_0341 = "/a/foo"

const literal_1489 = "/a/bar"

const literal_1850 = "/dictionary/skills"

const literal_3028 = "/dictionary/:name"

const literal_6954 = "/server"

const literal_3672 = "/:a/:b/:c/:id"

const literal_8370 = "/:a/:id"

const literal_5390 = "/:a/:e/:id"

const literal_0148 = "/users/:username"

const literal_9267 = "/users/:username/likes/projects/ids"

const literal_3697 = "/users/:username/profile"

const literal_7098 = "/users/:username/uploads/:type"

const literal_0948 = "/users/create"

const literal_5831 = "/static"

const literal_6832 = "/static/*"

const literal_3416 = "/params/:foo"

const literal_4207 = "/params/:foo/bar/:qux"

const literal_9351 = "/params/:foo/bar/:qux/*"

const literal_8765 = "/a/c/df"

const literal_8766 = "/repos/:owner/:repo"
