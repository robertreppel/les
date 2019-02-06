# "Consentaur" is a consenting system for medical data collection which we did at the UBC Blockchain Hackathon in 2017.

The original version ( https://github.com/robertreppel/consentaur ) was done with stickies-on-wall, group activity, think-out-loud, have all experts and stakeholders in the room. 

The version here has been turned into "green" ESL, suitable for auto-generating an event sourced API. 

* It required renaming some streams, commands, events and read models and assorted properties and parameters for uniqueness and clarity.

* Some read models which are needed to validate IDs (e.g. administratorId, projectId, ...) have been added.


![consentaur event storming](Eventstorming.png)


