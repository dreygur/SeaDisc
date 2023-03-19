# Opensea Collection floor price tracker and Discord Notifier

Opesea has recently introduced `stream-api` which is currently in `beta` stage and there might be some changes in future. Which emits new events when a transaction happens on a collection. Though the api may change in future, I would to like to give it a try.

Quick look: [Overview](https://docs.opensea.io/reference/stream-api-overview) |  [Sample Event Schemas](https://docs.opensea.io/reference/stream-api-event-schemas)

---
## What events can be listened for now?
- Item listed - an item listed for sale on the OpenSea marketplace
- Item sold - sale of an item on the OpenSea marketplace
- Item transferred - transfer of an item between wallets
- Item metadata update - update detected on the metadata provided in tokenURI for an item
- Item cancelled - cancellation of an order on the OpenSea marketplace
- Item received offer - offer received on an item in the OpenSea marketplace
- Item received bid - bid received on an item in the OpenSea marketplace

## What events can I listen until now?
Every single event! :fire:

Let's have a try!
