# GitHub Sponsors API
A REST API for GitHub Sponsors, if you need a GraphQL one, there is an offical API. Some links and examples can be found in [the Notes section](#notes).

## Endpoints

### /count/user/
Get the count of people who sponsor `user`.

Example query: [`https://sponsors.trnck.dev/count/filiptronicek/`](https://sponsors.trnck.dev/count/filiptronicek)

<details>
<summary>
  Example responce:
</summary>

```json
{
  "sponsors": {
    "count": 4
  }
}
```
</details>

### /sponsors/user/
Get details of people who sponsor `user`.

Example query: [`https://sponsors.trnck.dev/sponsors/filiptronicek`](https://sponsors.trnck.dev/sponsors/filiptronicek)

<details>
<summary>
  Example responce:
</summary>

```json
{
  "sponsors": [
    {
      "handle": "bdougie",
      "avatar": "https://avatars.githubusercontent.com/u/5713670?s=60&v=4",
      "profile": "https://github.com/bdougie",
      "details": {
        "login": "bdougie",
        "id": 5713670,
        "node_id": "MDQ6VXNlcjU3MTM2NzA=",
        "avatar_url": "https://avatars.githubusercontent.com/u/5713670?v=4",
        "gravatar_id": "",
        "url": "https://api.github.com/users/bdougie",
        "html_url": "https://github.com/bdougie",
        "followers_url": "https://api.github.com/users/bdougie/followers",
        "following_url": "https://api.github.com/users/bdougie/following{/other_user}",
        "gists_url": "https://api.github.com/users/bdougie/gists{/gist_id}",
        "starred_url": "https://api.github.com/users/bdougie/starred{/owner}{/repo}",
        "subscriptions_url": "https://api.github.com/users/bdougie/subscriptions",
        "organizations_url": "https://api.github.com/users/bdougie/orgs",
        "repos_url": "https://api.github.com/users/bdougie/repos",
        "events_url": "https://api.github.com/users/bdougie/events{/privacy}",
        "received_events_url": "https://api.github.com/users/bdougie/received_events",
        "type": "User",
        "site_admin": true,
        "name": "Brian Douglas",
        "company": "GitHub",
        "blog": "https://bdougie.live",
        "location": "Oakland, CA",
        "email": null,
        "hireable": null,
        "bio": "Making open-source accessible with  @open-sauced (he/him)",
        "twitter_username": "bdougieYO",
        "public_repos": 358,
        "public_gists": 28,
        "followers": 1023,
        "following": 161,
        "created_at": "2013-10-17T22:25:29Z",
        "updated_at": "2021-02-07T14:45:40Z"
      }
    },
    {
      "handle": "svobodavl",
      "avatar": "https://avatars.githubusercontent.com/u/58887042?s=60&v=4",
      "profile": "https://github.com/svobodavl",
      "details": {
        "login": "svobodavl",
        "id": 58887042,
        "node_id": "MDQ6VXNlcjU4ODg3MDQy",
        "avatar_url": "https://avatars.githubusercontent.com/u/58887042?v=4",
        "gravatar_id": "",
        "url": "https://api.github.com/users/svobodavl",
        "html_url": "https://github.com/svobodavl",
        "followers_url": "https://api.github.com/users/svobodavl/followers",
        "following_url": "https://api.github.com/users/svobodavl/following{/other_user}",
        "gists_url": "https://api.github.com/users/svobodavl/gists{/gist_id}",
        "starred_url": "https://api.github.com/users/svobodavl/starred{/owner}{/repo}",
        "subscriptions_url": "https://api.github.com/users/svobodavl/subscriptions",
        "organizations_url": "https://api.github.com/users/svobodavl/orgs",
        "repos_url": "https://api.github.com/users/svobodavl/repos",
        "events_url": "https://api.github.com/users/svobodavl/events{/privacy}",
        "received_events_url": "https://api.github.com/users/svobodavl/received_events",
        "type": "User",
        "site_admin": false,
        "name": "Vláďa Svoboda",
        "company": null,
        "blog": "vladja.itch.io/",
        "location": "Prague, Czech Republic",
        "email": null,
        "hireable": true,
        "bio": null,
        "twitter_username": "vlada_svoboda",
        "public_repos": 7,
        "public_gists": 0,
        "followers": 8,
        "following": 9,
        "created_at": "2019-12-14T16:56:34Z",
        "updated_at": "2021-02-04T15:53:59Z"
      }
    },
    {
      "handle": "kahy9",
      "avatar": "https://avatars.githubusercontent.com/u/48121432?s=60&v=4",
      "profile": "https://github.com/kahy9",
      "details": {
        "login": "kahy9",
        "id": 48121432,
        "node_id": "MDQ6VXNlcjQ4MTIxNDMy",
        "avatar_url": "https://avatars.githubusercontent.com/u/48121432?v=4",
        "gravatar_id": "",
        "url": "https://api.github.com/users/kahy9",
        "html_url": "https://github.com/kahy9",
        "followers_url": "https://api.github.com/users/kahy9/followers",
        "following_url": "https://api.github.com/users/kahy9/following{/other_user}",
        "gists_url": "https://api.github.com/users/kahy9/gists{/gist_id}",
        "starred_url": "https://api.github.com/users/kahy9/starred{/owner}{/repo}",
        "subscriptions_url": "https://api.github.com/users/kahy9/subscriptions",
        "organizations_url": "https://api.github.com/users/kahy9/orgs",
        "repos_url": "https://api.github.com/users/kahy9/repos",
        "events_url": "https://api.github.com/users/kahy9/events{/privacy}",
        "received_events_url": "https://api.github.com/users/kahy9/received_events",
        "type": "User",
        "site_admin": false,
        "name": "Josef Kahoun",
        "company": "@MicrosoftSTC",
        "blog": "",
        "location": "Czech Republic",
        "email": null,
        "hireable": null,
        "bio": null,
        "twitter_username": "kahy_dot_sh",
        "public_repos": 14,
        "public_gists": 0,
        "followers": 11,
        "following": 18,
        "created_at": "2019-03-01T09:36:36Z",
        "updated_at": "2021-02-03T19:10:30Z"
      }
    },
    {
      "handle": "kdaigle",
      "avatar": "https://avatars.githubusercontent.com/u/2501?s=60&v=4",
      "profile": "https://github.com/kdaigle",
      "details": {
        "login": "kdaigle",
        "id": 2501,
        "node_id": "MDQ6VXNlcjI1MDE=",
        "avatar_url": "https://avatars.githubusercontent.com/u/2501?v=4",
        "gravatar_id": "",
        "url": "https://api.github.com/users/kdaigle",
        "html_url": "https://github.com/kdaigle",
        "followers_url": "https://api.github.com/users/kdaigle/followers",
        "following_url": "https://api.github.com/users/kdaigle/following{/other_user}",
        "gists_url": "https://api.github.com/users/kdaigle/gists{/gist_id}",
        "starred_url": "https://api.github.com/users/kdaigle/starred{/owner}{/repo}",
        "subscriptions_url": "https://api.github.com/users/kdaigle/subscriptions",
        "organizations_url": "https://api.github.com/users/kdaigle/orgs",
        "repos_url": "https://api.github.com/users/kdaigle/repos",
        "events_url": "https://api.github.com/users/kdaigle/events{/privacy}",
        "received_events_url": "https://api.github.com/users/kdaigle/received_events",
        "type": "User",
        "site_admin": true,
        "name": "Kyle Daigle",
        "company": "@github ",
        "blog": "www.kyledaigle.com",
        "location": "Tolland, CT",
        "email": null,
        "hireable": null,
        "bio": "Senior Director, Special Projects at @github ",
        "twitter_username": "kdaigle",
        "public_repos": 67,
        "public_gists": 7,
        "followers": 385,
        "following": 6,
        "created_at": "2008-03-07T14:32:21Z",
        "updated_at": "2021-02-09T21:12:22Z"
      }
    }
  ]
}
```
</details>

### /sponsoring/user/
Get all users who are sponsored by `user`.

Example query: [`https://sponsors.trnck.dev/sponsoring/svobodavl`](https://sponsors.trnck.dev/sponsoring/svobodavl)

<details>
<summary>
  Example responce:
</summary>

```json
{
  "sponsorees": [
    {
      "handle": "filiptronicek",
      "avatar": "https://avatars.githubusercontent.com/u/29888641?s=88&u=152b134e3e6e3d003ecd55fdde31c4171144c771&v=4",
      "profile": "https://github.com/filiptronicek"
    }
  ]
}
```
</details>

## Notes
- The `/sponsors/user/` and `/count/user/` endpoints got themselves an official API (as of June of 2021)! you can use them like this: (https://github.com/github/feedback/discussions/3818)

```gql
query {
  user(login: "filiptronicek") {
    ... on Sponsorable {
      sponsors(first: 100) {
        totalCount
        nodes {
          ... on User { login }
          ... on Organization { login }
        }
      }
    }
  }
}
```
- The `/sponsoring/user/` endpoint is implemented in the GraphQL API since August of 2021, usable like this: https://github.com/github/feedback/discussions/3818#discussioncomment-1131586

```gql
query {
  user(login: "cheshire137") {
    sponsoring(first: 10) {
      totalCount
      nodes {
        ... on User { login }
        ... on Organization { login }
      }
    }
  }
}
```
