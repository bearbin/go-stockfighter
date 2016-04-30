/*
Package sflib provides a client for using the Stockfighter API.

Getting Started

	client := sflib.NewClient("APIKey")

	// Check that the API is up and running.
	err := client.Heartbeat()

This package requires an API key. You can fetch the key from your Stockfighter
[account page](https://www.stockfighter.io/ui/api_keys), or, if you wish you may
use your session key, which will be destroyed when you log out.

Spoilers

This package is not intended to spoil your fun. This package simply implements the
API as documented on the [Stockfighter site](https://starfighter.readme.io), and
any special features that you require for specific levels will have to be implemented
on your own.

Your own Work

This package will no doubt be helpful with your exploration of the API and with
your progression through the levels. However, you will find it more helpful to
take this code and build on it yourself rather than simply being a "user".
*/
package sflib
