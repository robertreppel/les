# Event Modeling Language Language Compliance Test

Tests whether an API created from EML by a builder such as les-node complies with the EML specification.

* Are all the command processor end points working?
* Are the read models working?
* Are all the business rules implemented?
* Are the required validation errors returned when executing invalid commands?

## Test ESL

Tests all features which can be used frm Event Storming Language

```cd esl && make setup && sleep 1 && make test```


## Test EML

Tests all features which are supported in Event Modeling Language Language 0.10.1-alpha, but are not part of ESL.

```cd eml && make setup && sleep 1 && make test```
