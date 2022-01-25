# change log of the forest package

v1.5.0 [2022-01-25]

- graphql support
- optional request line logging
- fix line nr in log

v1.4.2 [2022-01-13]

- show count in masked headers
- remove  NewTestingT
- add graphql support
- add cookie handling

v1.2.0 [2021-06-15]

- add support for GraphQL request

v1.1.1

- fixes problem with dumping a text response that contains Go format markers.

v1.1.0

- add Form to RequestConfig (thanx to Arvind Mohabir)

v1.0.0

- add VerboseOnFailure to have more information when a failure is detected.
- add Fatalf (to be used instead of t.Fatal)
- remove FailMessagePrefix package variable