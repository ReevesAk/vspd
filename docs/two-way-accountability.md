# Two-way Accountability

- When vspd is started for the first time, it generates a ed25519 keypair and
  stores it in the database. This key is used to sign all API responses, and the
  signature is included in the response header `VSP-Server-Signature`.

- Every client request which references a ticket should include a HTTP header
  `VSP-Client-Signature`. The value of this header must be a signature of the
  request body, signed with the commitment address of the referenced ticket.

## Examples

### Server does not vote ticket

- The client gets an address to pay the VSP fee, creates a
  transaction showing payment of the VSP fee  and submits it
  to the server along with voting preferences for the ticket.
  The server returns a timestamped JSON response pledging to
  the ticket and voting preferences thus providing valid
  prove that these voting preferences were sent by the clients.

- Here is an example of how the request body and JSON response
  would look like?:

  client:
  {
  	"request": "payfee",
  	"fee_tx": "serialized transaction paying fee",
  	"voting_key": "...",
  	"agenda_preferences": "..."
  }
  server:
  {
  	"receipt": "..."
  }




### Client denies changing their voting preferences

- This occurs when a malicious clients update their voting preferences, with
  VSP honoring the new changes and yhe client further uses and old receipt
  to accuse the VSP od the voting with the wrong preferences. This replay
  attack is rightly handled by making the clients sign request body with
  ticket commitment address and put the result in an HTTP header thus providing
  client with accountability and prove of ach ticket ownership.

- Below is a typical example of how the request body from the client
  should look like and the expected JSON response from the server.

  client:
  {
    "request": "setpreferences",
  	"ticket": "ticket hash",
  	"commitment": "signature using commitment key",
    "agenda_preferences": "..."
  }

  server:
  {
          "receipt": "..."
  }

