# change log of the forest package

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