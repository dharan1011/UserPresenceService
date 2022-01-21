# User Presence System

### Problem Statement
Design a service which can tell if a user is online. </br>

### Data Access
* Service should tell if user is online / offline
* Service to handle high TPS
* Service to have low service latency
* Service to handle edge case
### Client expectation
* Clients to ping the service within the TTL to notify user presence every X seconds

### Database
* Presence entry to expire after TTL
* Database to save entry for analytics
* Database to soft delete entries