# Challenge pusher
The is a tool for pushing CTF challenge configurations to the [haaukins exercise database](https://github.com/aau-network-security/haaukins-exercises).
And is used in our private repositories to add/update challenges based on configurations created by the developers. 

## Example challenge config

```yml
name: Challenge Name
tag: unique_challenge_tag
category: Forensics
secret: true
solutions:
  - challengename: Sub challenge 1
    steps: |
     Step-by-step guide to challenge
  - challengename: Sub challenge 2
    steps: |
     Step-by-step guide to challenge
description: |
  A very decent description
hosts:
  - image: <path to image>
    flags:
      - tag: unique_challenge_tag-0
        name: Sub challenge 1
        static: HKN{FLAG-STRING}
        points: 20
        category: Forensics
        description: |
          Some text for participants
  - image: <path to image>
    flags:
      - tag: unique_challenge_tag-1
        name: Sub challenge 2
        static: HKN{FLAG-STRING}
        points: 20
        category: Forensics
        description: |
          Some text for participants
```

