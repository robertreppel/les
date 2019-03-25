# Upcoming Releases

__Features & Fixes__

## 0.10.8-alpha (probably)

This version relaxes some validation rules for Event Storming Language which imposed unnecessary limits on what kind of event stormings can be used as input for code generation:

* In a real life event storming, an event can occur more than once in a workflow. This is now valid ESL.
* In real life, a command can result in more than one event. And so it is now in ESL.
* Read models can have derived or renamed properties which do not correspond to properties of events. It should be possible, once such properties have been identified, to tell them in EML which ones of the event-derived properties are the input for such determinations or calculations. E.g. "totalRevenue" would be identified as being based on the "amount" fields of the "Sale Made" and "Sale Canceled" events. 

## After that ...

Incorporate ESL event stormings and workflows into the schema. This will allow cross-referencing to determine if there are commands/events/read models which do not belong to one or more workflows and flag them in validation. 


