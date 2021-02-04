package dependency_test

import (
	"github.com/paketo-buildpacks/dep-server/pkg/dependency"
	"github.com/paketo-buildpacks/dep-server/pkg/dependency/dependencyfakes"
	"github.com/paketo-buildpacks/dep-server/pkg/dependency/internal"
	"github.com/sclevine/spec"
	"github.com/sclevine/spec/report"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestComposer(t *testing.T) {
	spec.Run(t, "composer", testComposer, spec.Report(report.Terminal{}))
}

func testComposer(t *testing.T, when spec.G, it spec.S) {
	var (
		assert           = assert.New(t)
		require          = require.New(t)
		fakeChecksummer  *dependencyfakes.FakeChecksummer
		fakeFileSystem   *dependencyfakes.FakeFileSystem
		fakeGithubClient *dependencyfakes.FakeGithubClient
		fakeWebClient    *dependencyfakes.FakeWebClient
		composer         dependency.Dependency
	)

	it.Before(func() {
		fakeChecksummer = &dependencyfakes.FakeChecksummer{}
		fakeFileSystem = &dependencyfakes.FakeFileSystem{}
		fakeGithubClient = &dependencyfakes.FakeGithubClient{}
		fakeWebClient = &dependencyfakes.FakeWebClient{}

		var err error
		composer, err = dependency.NewCustomDependencyFactory(fakeChecksummer, fakeFileSystem, fakeGithubClient, fakeWebClient).NewDependency("composer")
		require.NoError(err)
	})

	when("GetAllVersionRefs", func() {
		it("returns all composer release versions with newest versions first", func() {
			fakeGithubClient.GetReleaseTagsReturns([]internal.GithubRelease{
				{
					TagName:       "3.0.0",
					PublishedDate: "2020-06-30T00:00:00Z",
				},
				{
					TagName:       "1.0.1",
					PublishedDate: "2020-06-29T00:00:00Z",
				},
				{
					TagName:       "2.0.0",
					PublishedDate: "2020-06-28T00:00:00Z",
				},
				{
					TagName:       "1.0.0",
					PublishedDate: "2020-06-27T00:00:00Z",
				},
			}, nil)

			versions, err := composer.GetAllVersionRefs()
			require.NoError(err)
			assert.Equal([]string{"3.0.0", "1.0.1", "2.0.0", "1.0.0"}, versions)

			orgArg, repoArg := fakeGithubClient.GetReleaseTagsArgsForCall(0)
			assert.Equal("composer", orgArg)
			assert.Equal("composer", repoArg)
		})
	})

	when("GetDependencyVersion", func() {
		it("returns the correct composer version", func() {
			fakeGithubClient.GetReleaseTagsReturns([]internal.GithubRelease{
				{
					TagName:       "3.0.0",
					PublishedDate: "2020-06-30T00:00:00Z",
				},
				{
					TagName:       "1.0.1",
					PublishedDate: "2020-06-29T00:00:00Z",
				},
				{
					TagName:       "2.0.0",
					PublishedDate: "2020-06-28T00:00:00Z",
				},
			}, nil)
			fakeWebClient.GetReturnsOnCall(0,
				[]byte(`aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa  composer.phar`), nil)

			actualDep, err := composer.GetDependencyVersion("1.0.1")
			require.NoError(err)

			expectedDep := dependency.DepVersion{
				Version:         "1.0.1",
				URI:             "https://getcomposer.org/download/1.0.1/composer.phar",
				SHA:             "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
				ReleaseDate:     "2020-06-29T00:00:00Z",
				DeprecationDate: "",
			}
			assert.Equal(expectedDep, actualDep)

			orgArg, repoArg := fakeGithubClient.GetReleaseTagsArgsForCall(0)
			assert.Equal("composer", orgArg)
			assert.Equal("composer", repoArg)
		})

		when("the SHA cannot be found", func() {
			it("returns an error", func() {
				fakeGithubClient.GetReleaseTagsReturns([]internal.GithubRelease{
					{
						TagName:       "3.0.0",
						PublishedDate: "2020-06-30T00:00:00Z",
					},
				}, nil)
				fakeWebClient.GetReturnsOnCall(0, nil, nil)

				_, err := composer.GetDependencyVersion("3.0.0")
				assert.Error(err)

				assert.Contains(err.Error(), "could not get SHA from file")
			})
		})
	})
}
