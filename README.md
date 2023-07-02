# GitHub Sponsors API
A REST API for GitHub Sponsors, if you need a GraphQL one, there is an official API. Some links and examples can be found in [the Notes section](#notes).

## Endpoints

### /v2/count/user/
Get the count of people who sponsor `user`. Includes private sponsors.

Example query: [`https://ghs.vercel.app/v2/count/filiptronicek/`](https://ghs.vercel.app/v2/count/filiptronicek)

<details>
<summary>
  Example response:
</summary>

```json
{
  "status": "success",
  "sponsors": {
    "current": 4,
    "past": 10
  }
}
```
</details>

### /v2/sponsors/user/
Get details of people who sponsor `user`. Does **not** include private sponsors.

Example query: [`https://ghs.vercel.app/v2/sponsors/filiptronicek`](https://ghs.vercel.app/v2/sponsors/filiptronicek)

<details>
<summary>
  Example response:
</summary>

```json
{
  "status": "success",
  "sponsors": {
    "current": [
      {
        "username": "aellopos",
        "avatar": "https://avatars.githubusercontent.com/u/39790985?s=60&v=4"
      },
      {
        "username": "mosh98",
        "avatar": "https://avatars.githubusercontent.com/u/48658042?s=60&v=4"
      },
      {
        "username": "kahy9",
        "avatar": "https://avatars.githubusercontent.com/u/48121432?s=60&v=4"
      },
      {
        "username": "0ndras3k",
        "avatar": "https://avatars.githubusercontent.com/u/57116019?s=60&v=4"
      },
      {
        "username": "AdamSchinzel",
        "avatar": "https://avatars.githubusercontent.com/u/66002635?s=60&v=4"
      },
      {
        "username": "czM1K3",
        "avatar": "https://avatars.githubusercontent.com/u/45005362?s=60&v=4"
      },
      {
        "username": "svobodavl",
        "avatar": "https://avatars.githubusercontent.com/u/58887042?s=60&v=4"
      },
      {
        "username": "bigint",
        "avatar": "https://avatars.githubusercontent.com/u/69431456?s=60&v=4"
      },
      {
        "username": "anuraghazra",
        "avatar": "https://avatars.githubusercontent.com/u/35374649?s=60&v=4"
      }
    ],
    "past": [
      {
        "username": "scraptechguy",
        "avatar": "https://avatars.githubusercontent.com/u/75474651?s=60&v=4"
      },
      {
        "username": "bdougie",
        "avatar": "https://avatars.githubusercontent.com/u/5713670?s=60&v=4"
      },
      {
        "username": "kdaigle",
        "avatar": "https://avatars.githubusercontent.com/u/2501?s=60&v=4"
      }
    ]
  }
}
```
</details>

### /sponsoring/user/
Get all users who are sponsored by `user`.

Example query: [`https://ghs.vercel.app/sponsoring/svobodavl`](https://ghs.vercel.app/sponsoring/svobodavl)

<details>
<summary>
  Example response:
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
