> [!NOTE]
> **This is a fork of [sbondCo/Watcharr](https://github.com/sbondCo/Watcharr).** It adds one feature on top of v4.4.0: a **TMDB Language** server setting, so titles, overviews and posters are fetched from TMDB in the language you choose, with English as fallback for untranslated fields.
>
> **Credit:** the idea and the design come from [fjacques-2025](https://github.com/fjacques-2025), author of the upstream pull request [feat: configurable metadata language (TMDB_LANG) with English fallback](https://github.com/sbondCo/Watcharr/pull/1079). This fork reimplements that work on top of v4.4.0, after the TMDB code moved to `server/media/tmdb`.
>
> Things to know:
>
> - Content already on your lists keeps its previous language until you open its page again.
> - Search and discover results have no English fallback.
> - No Docker image is published for this fork; build it from source with the included `Dockerfile`.
> - Not affiliated with the Watcharr project. For the official version, use the upstream repository.

<h1 align="center">Watcharr</h1>
<p align="center"><img src="./static/logo-col.png" alt="logo" width="250" /></p>

<p align="center">
  <a href="https://github.com/sbondCo/Watcharr/pkgs/container/watcharr"><img src="https://img.shields.io/github/v/release/sbondCo/Watcharr?label=version&style=for-the-badge" /></a>
  <a href="https://beta.watcharr.app"><img src="https://img.shields.io/website?label=DEMO&style=for-the-badge&url=https%3A%2F%2Fbeta.watcharr.app" /></a>
  <a href="https://watcharr.app"><img src="https://img.shields.io/website?label=DOCS&style=for-the-badge&url=https%3A%2F%2Fwatcharr.app" /></a>
  <a href="https://github.com/sbondCo/Watcharr/issues"><img src="https://img.shields.io/github/issues-raw/sbondCo/Watcharr?label=ISSUES&style=for-the-badge" /></a>
  <a href="/LICENSE"><img src="https://img.shields.io/github/license/sbondCo/Watcharr?style=for-the-badge" /></a>
  <a href="https://matrix.to/#/#watcharr:matrix.org"><img src="https://img.shields.io/matrix/watcharr%3Amatrix.org?style=for-the-badge&logo=matrix" /></a>
</p>

I'm your new easily self-hosted content watched list. The place you store your watched (or watching, planned, etc) **movies** and **tv shows** (and **anime**), rate them and track their status.

With [some extra configuration](https://watcharr.app/docs/server_config/game-support-igdb) I can also track your **video games**.

I am built with Go and Svelte(Kit).

### Demo

Feel free to abuse this demo instance (nicely). It runs on the latest `dev` build so there may be bugs, as new features are tested on here too. The demo is a worst-case scenario for speed (which is why I like it as a testing ground), if you host it yourself the app will be snappy.

Demo: [https://beta.watcharr.app/](https://beta.watcharr.app/)

**NOTE:** There is no demo account, just type in a random username/password (smashing hand into keyboard is supported), then click `Not a user?` at the bottom and a `Sign Up` button will appear.

### Track new features

Most patches are tracked through [our project board](https://github.com/orgs/sbondCo/projects/9/views/3), though I am very unorganised so expect surprise updates (or don't, if you like surprises)!

You can also [view a list of all current features](FEATURES.md).

### Contents

- [Screenshots](#screenshots)
- [Set Up](#set-up)
- [Community Made Tools](#community-made-tools)
- [Getting Help](#getting-help)
- [License](#license)
- [Contributing](#contributing)

# Screenshots

<p align="center">

<img src="./screenshot/devices-mock.png" alt="Overview" />

| Homepage                                                       | Watched Show Hover                                                          |
| -------------------------------------------------------------- | --------------------------------------------------------------------------- |
| <img src="./screenshot/homepage.png?v=2" alt="Watched List" /> | <img src="./screenshot/homepage-poster-hover.png?v=2" alt="Watched List" /> |

| Watched Show Status Change                                                                  | Movie Details                                                                      |
| ------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| <img src="./screenshot/homepage-poster-change-status.png?v=2" alt="Changing Show Status" /> | <img src="./screenshot/content-details-page.png?v=2" alt="Content Details Page" /> |

| User Profile                                                            | Discover                                                             |
| ----------------------------------------------------------------------- | -------------------------------------------------------------------- |
| <img src="./screenshot/user-profile.png?v=2" alt="User Profile Page" /> | <img src="./screenshot/discover-page.png?v=2" alt="Discover Page" /> |

| Dark Homepage                                                              | Dark Content Details                                                                       |
| -------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------ |
| <img src="./screenshot/homepage-dark.png?v=2" alt="Dark Theme Homepage" /> | <img src="./screenshot/content-details-page-dark.png?v=2" alt="Dark Theme Content Page" /> |

</p>

# Set Up

[Checkout our documentation](https://watcharr.app/docs/category/installation) for an up to date guide on setup! If you hate manuals, but love docker, this [docker-compose.yml](./docker-compose.yml) file is your friend.

# Community Made Tools

Third-party tools made by the community for enhancing your Watcharr experience!

- [Kodi Plugin](https://github.com/airdogvan/watcharr_kodi) by [airdogvan](https://github.com/airdogvan) for automatically tracking your watched shows/movies.

Thanks to anyone that has made a script or tool for Watcharr. Feel free to add your own to the list if you have one!

**Note:** I cannot provide any assurances for these tools or stay on top of them (code review, etc), if you have any problems please open an issue in the project for the tool so that they can stay organized.

# Getting Help

If something isn't working for you or you are stuck, [creating an issue](https://github.com/sbondCo/Watcharr/issues/new) is the best way to get help! Every type of issue is accepted, so don't be afraid to ask anything!

You can also [join our space on Matrix](https://matrix.to/#/#watcharr:matrix.org) for support.

I'll do my best to reply!

# License

This project is licensed under the GPLv3 license. You should see the [LICENSE](LICENSE) file located in the root folder of this project for the full license text, if not, see <https://www.gnu.org/licenses/>.

# Contributing

Please continue to our [contributing guide](CONTRIBUTING.md).
