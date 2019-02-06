# Upcoming Releases

__Features & Fixes__

## 0.2.0-alpha

This version relaxes some validation rules for Event Storming Language which imposed unnecessary limits on what kind of event stormings can be used as input for code generation:

* In a real life event storming, an event can occur more than once in a workflow. This is now valid ESL.
* In real life, a command can result in more than one event. And so it is now in ESL.
* Read models can have derived or renamed properties which do not correspond to properties of events.
