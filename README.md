Stream data from server to clients. Clients should be able to join mid-stream and pick up data from when they are finally connected. This should accept arbitrary numbers of connections but each connection should be identifiable by some kind of id selected by that client.

Considerations:
  - Multiple clients with the same id.
    - Maybe combine the ip address and id
  - Have start/end signals understood by the client as to do actions.
  - Allow for sending specific pieces of data to specific clients such that each client receives completely different data than any other client.



### UPDATE
As of 9/24/17

Switching from a udp base to a tcp one. This occurs from reading more about streaming data in tcp vs udp.
